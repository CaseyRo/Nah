package nah

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

// maxBody caps every request. Text moments are small and media does not go
// through this path yet; raise it deliberately when M3 arrives.
const maxBody = 1 << 20 // 1 MiB

const (
	defaultFeedLimit = 30
	maxFeedLimit     = 200
)

type Server struct {
	store *Store
	auth  *Auth
	log   *slog.Logger

	// ponytail: one limiter for the whole server, not per IP. A family server
	// behind a tunnel has no noisy neighbours; give it a per-IP map when it does.
	unauthed *rate.Limiter
}

func NewServer(store *Store, auth *Auth, log *slog.Logger) *Server {
	return &Server{
		store:    store,
		auth:     auth,
		log:      log,
		unauthed: rate.NewLimiter(20, 40),
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("POST /v1/people", s.throttle(s.handleRegister))
	mux.HandleFunc("POST /v1/people/{person}/challenge", s.throttle(s.handleChallenge))
	mux.HandleFunc("POST /v1/people/{person}/session", s.throttle(s.handleSession))
	mux.HandleFunc("POST /v1/people/{person}/invites", s.handleInvite)
	mux.HandleFunc("POST /v1/people/{person}/connections", s.handleConnect)
	mux.HandleFunc("GET /v1/people/{person}/feed", s.handleFeed)
	mux.HandleFunc("POST /v1/people/{person}/moments", s.handlePost)
	return s.logging(mux)
}

// --- handlers ---

// handleRegister makes a person for a device. It is open, like installing the
// app: a person with no connections can read nothing and reach nobody, and
// connecting still takes a touch or an invite.
func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	var in struct {
		PublicKey []byte `json:"public_key"`
	}
	if !decode(w, r, &in) {
		return
	}
	pub, ok := publicKey(w, in.PublicKey)
	if !ok {
		return
	}
	id, err := s.store.Register(pub)
	if err != nil {
		s.oops(w, r, err)
		return
	}
	write(w, http.StatusCreated, map[string]string{"id": id})
}

func (s *Server) handleChallenge(w http.ResponseWriter, r *http.Request) {
	var in struct {
		PublicKey []byte `json:"public_key"`
	}
	if !decode(w, r, &in) {
		return
	}
	pub, ok := publicKey(w, in.PublicKey)
	if !ok {
		return
	}
	person := r.PathValue("person")
	// A challenge is issued without checking whose key this is on purpose:
	// answering differently for a stranger would turn this into an oracle.
	if !ValidID(person) {
		fail(w, http.StatusNotFound, "No such person.")
		return
	}
	write(w, http.StatusOK, map[string]any{
		"challenge": s.auth.Challenge(person, pub),
	})
}

func (s *Server) handleSession(w http.ResponseWriter, r *http.Request) {
	var in struct {
		PublicKey []byte `json:"public_key"`
		Challenge string `json:"challenge"`
		Signature []byte `json:"signature"`
	}
	if !decode(w, r, &in) {
		return
	}
	pub, ok := publicKey(w, in.PublicKey)
	if !ok {
		return
	}
	person := r.PathValue("person")
	owns, err := s.store.Owns(person, pub)
	if err != nil {
		s.storeErr(w, r, err)
		return
	}
	if !owns {
		fail(w, http.StatusForbidden, "This device does not belong to that person.")
		return
	}
	token, expires, err := s.auth.Session(person, in.Challenge, pub, in.Signature)
	if err != nil {
		fail(w, http.StatusUnauthorized, "That did not verify. Ask for a new challenge and try again.")
		return
	}
	write(w, http.StatusOK, map[string]any{
		"token":      token,
		"expires_at": expires.UnixMilli(),
	})
}

func (s *Server) handleInvite(w http.ResponseWriter, r *http.Request) {
	person, ok := s.authorize(w, r)
	if !ok {
		return
	}
	token, err := s.store.Invite(person)
	if err != nil {
		s.storeErr(w, r, err)
		return
	}
	write(w, http.StatusCreated, map[string]string{"invite": token})
}

// handleConnect redeems an invite for the person in the path. The body names
// who made it; a touch and a link both deliver exactly these two things.
func (s *Server) handleConnect(w http.ResponseWriter, r *http.Request) {
	person, ok := s.authorize(w, r)
	if !ok {
		return
	}
	var in struct {
		Person string `json:"person"`
		Invite string `json:"invite"`
	}
	if !decode(w, r, &in) {
		return
	}
	switch err := s.store.Connect(person, in.Person, in.Invite); {
	case err == nil:
		w.WriteHeader(http.StatusNoContent)
	case errors.Is(err, ErrNetworkFull):
		// Language, never a number (ADR-0004). Either side can be the full one;
		// CDI-1842 owns the real sentence, and this placeholder fits both.
		fail(w, http.StatusConflict, "There is no room for this connection right now.")
	case errors.Is(err, ErrBadInvite):
		fail(w, http.StatusForbidden, "This invitation is no longer valid.")
	case errors.Is(err, ErrOwnInvite):
		fail(w, http.StatusBadRequest, "That is your own invitation.")
	default:
		s.storeErr(w, r, err)
	}
}

func (s *Server) handleFeed(w http.ResponseWriter, r *http.Request) {
	person, ok := s.authorize(w, r)
	if !ok {
		return
	}
	limit := defaultFeedLimit
	if n, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil && n > 0 {
		limit = min(n, maxFeedLimit)
	}
	moments, err := s.store.Feed(person, limit)
	if err != nil {
		s.storeErr(w, r, err)
		return
	}
	write(w, http.StatusOK, moments)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	person, ok := s.authorize(w, r)
	if !ok {
		return
	}
	var in struct {
		Blob []byte `json:"blob"`
	}
	if !decode(w, r, &in) {
		return
	}
	if len(in.Blob) == 0 {
		fail(w, http.StatusBadRequest, "A moment cannot be empty.")
		return
	}
	m, err := s.store.Post(person, in.Blob)
	if err != nil {
		s.storeErr(w, r, err)
		return
	}
	write(w, http.StatusCreated, m)
}

// --- plumbing ---

// authorize resolves the bearer token to the person in the path.
func (s *Server) authorize(w http.ResponseWriter, r *http.Request) (person string, ok bool) {
	person = r.PathValue("person")
	token, found := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !found || token == "" {
		fail(w, http.StatusUnauthorized, "Sign in first.")
		return "", false
	}
	pub, err := s.auth.Lookup(person, token)
	if err != nil {
		fail(w, http.StatusUnauthorized, "That session has expired. Sign in again.")
		return "", false
	}
	// Checked per request rather than at sign-in, so that replacing a lost
	// phone's key takes effect immediately once M5 can do that.
	owns, err := s.store.Owns(person, pub)
	if err != nil {
		s.storeErr(w, r, err)
		return "", false
	}
	if !owns {
		fail(w, http.StatusForbidden, "This device does not belong to that person.")
		return "", false
	}
	return person, true
}

func (s *Server) throttle(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.unauthed.Allow() {
			fail(w, http.StatusTooManyRequests, "Too many attempts. Try again in a moment.")
			return
		}
		h(w, r)
	}
}

func (s *Server) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		// Paths carry person ids and nothing else. No body, ever: it is ciphertext
		// and it is not ours.
		s.log.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.status,
			"ms", time.Since(start).Milliseconds())
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (s *Server) storeErr(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, ErrNoPerson) || errors.Is(err, ErrBadPersonID) {
		fail(w, http.StatusNotFound, "No such person.")
		return
	}
	s.oops(w, r, err)
}

func (s *Server) oops(w http.ResponseWriter, r *http.Request, err error) {
	s.log.Error("request failed", "path", r.URL.Path, "err", err)
	fail(w, http.StatusInternalServerError, "Something went wrong here. It is not your fault.")
}

func decode(w http.ResponseWriter, r *http.Request, v any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, maxBody)
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		fail(w, http.StatusBadRequest, "That request did not make sense.")
		return false
	}
	return true
}

func publicKey(w http.ResponseWriter, b []byte) (ed25519.PublicKey, bool) {
	if len(b) != ed25519.PublicKeySize {
		fail(w, http.StatusBadRequest, "That is not a valid device key.")
		return nil, false
	}
	return ed25519.PublicKey(b), true
}

func write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func fail(w http.ResponseWriter, status int, msg string) {
	write(w, status, map[string]string{"error": msg})
}
