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
	mux.HandleFunc("POST /v1/circles/{circle}/join", s.throttle(s.handleJoin))
	mux.HandleFunc("POST /v1/circles/{circle}/challenge", s.throttle(s.handleChallenge))
	mux.HandleFunc("POST /v1/circles/{circle}/session", s.throttle(s.handleSession))
	mux.HandleFunc("GET /v1/circles/{circle}/moments", s.handleFeed)
	mux.HandleFunc("POST /v1/circles/{circle}/moments", s.handlePost)
	return s.logging(mux)
}

// --- handlers ---

func (s *Server) handleJoin(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Invite    string `json:"invite"`
		PublicKey []byte `json:"public_key"`
	}
	if !decode(w, r, &in) {
		return
	}
	pub, ok := publicKey(w, in.PublicKey)
	if !ok {
		return
	}
	switch err := s.store.Join(r.PathValue("circle"), in.Invite, pub); {
	case err == nil:
		w.WriteHeader(http.StatusNoContent)
	case errors.Is(err, ErrCircleFull):
		// Language, never a number (ADR-0004).
		fail(w, http.StatusConflict, "This circle is full.")
	case errors.Is(err, ErrBadInvite):
		fail(w, http.StatusForbidden, "This invitation is no longer valid.")
	default:
		s.oops(w, r, err)
	}
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
	circle := r.PathValue("circle")
	// A challenge is issued without checking membership on purpose: answering
	// differently for a stranger would turn this into a membership oracle.
	if !ValidID(circle) {
		fail(w, http.StatusNotFound, "No such circle.")
		return
	}
	write(w, http.StatusOK, map[string]any{
		"challenge": s.auth.Challenge(circle, pub),
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
	circle := r.PathValue("circle")
	member, err := s.store.IsMember(circle, pub)
	if err != nil {
		s.storeErr(w, r, err)
		return
	}
	if !member {
		fail(w, http.StatusForbidden, "This device is not a member of that circle.")
		return
	}
	token, expires, err := s.auth.Session(circle, in.Challenge, pub, in.Signature)
	if err != nil {
		fail(w, http.StatusUnauthorized, "That did not verify. Ask for a new challenge and try again.")
		return
	}
	write(w, http.StatusOK, map[string]any{
		"token":      token,
		"expires_at": expires.UnixMilli(),
	})
}

func (s *Server) handleFeed(w http.ResponseWriter, r *http.Request) {
	circle, _, ok := s.authorize(w, r)
	if !ok {
		return
	}
	limit := defaultFeedLimit
	if n, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil && n > 0 {
		limit = min(n, maxFeedLimit)
	}
	moments, err := s.store.Feed(circle, limit)
	if err != nil {
		s.storeErr(w, r, err)
		return
	}
	write(w, http.StatusOK, moments)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	circle, pub, ok := s.authorize(w, r)
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
	m, err := s.store.Post(circle, pub, in.Blob)
	if err != nil {
		s.storeErr(w, r, err)
		return
	}
	write(w, http.StatusCreated, m)
}

// --- plumbing ---

// authorize resolves the bearer token to a member of the circle in the path.
func (s *Server) authorize(w http.ResponseWriter, r *http.Request) (circle string, pub ed25519.PublicKey, ok bool) {
	circle = r.PathValue("circle")
	token, found := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !found || token == "" {
		fail(w, http.StatusUnauthorized, "Sign in first.")
		return "", nil, false
	}
	pub, err := s.auth.Lookup(circle, token)
	if err != nil {
		fail(w, http.StatusUnauthorized, "That session has expired. Sign in again.")
		return "", nil, false
	}
	// Checked per request rather than at sign-in, so that removing a member
	// takes effect immediately once M5 can remove one.
	member, err := s.store.IsMember(circle, pub)
	if err != nil {
		s.storeErr(w, r, err)
		return "", nil, false
	}
	if !member {
		fail(w, http.StatusForbidden, "This device is not a member of that circle.")
		return "", nil, false
	}
	return circle, pub, true
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
		// Paths carry circle ids and nothing else. No body, ever: it is ciphertext
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
	if errors.Is(err, ErrNoCircle) || errors.Is(err, ErrBadCircleID) {
		fail(w, http.StatusNotFound, "No such circle.")
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
