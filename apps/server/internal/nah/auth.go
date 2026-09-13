package nah

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Sign-in per ADR-0015: no passwords. The device holds an Ed25519 key, the
// server hands out a challenge, and a signature over that challenge buys a
// session token.
//
// A session token carries its own claims and a MAC over them, so it survives a
// restart. Deploys are going to be frequent and nobody should be able to tell
// one happened. The signing key lives in the data directory, which means it
// moves with the volume and is shared by two instances during a rolling deploy.

const (
	challengeTTL = 2 * time.Minute
	sessionTTL   = 30 * 24 * time.Hour

	// signPrefix domain-separates this signature from anything else the same
	// identity key might ever sign, and binds it to one circle so a signature
	// captured on one circle cannot open another.
	signPrefix = "nah-auth-v1:"

	// tokenVersion prefixes every session token so the format can change later
	// without anyone having to guess what an old token is.
	tokenVersion = "v1"

	sessionKeyFile = "session.key"
	sessionKeySize = 32
)

var (
	ErrBadChallenge = errors.New("challenge is unknown or expired")
	ErrBadSignature = errors.New("signature does not verify")
	ErrNoSession    = errors.New("no valid session")
)

// SignedMessage is what a device signs to prove it holds the key. The client
// must build exactly this string.
func SignedMessage(circleID, challenge string) []byte {
	return []byte(signPrefix + circleID + ":" + challenge)
}

// SessionKey reads the key that signs session tokens, generating it on first
// run. It sits in the data directory rather than in the environment so that a
// deploy cannot forget it, and so a restored volume restores the sessions with
// it.
func SessionKey(dir string) ([]byte, error) {
	path := filepath.Join(dir, sessionKeyFile)
	key, err := os.ReadFile(path)
	switch {
	case err == nil && len(key) == sessionKeySize:
		return key, nil
	case err == nil:
		return nil, errors.New("session key file is the wrong size; move it aside to have a new one made")
	case !errors.Is(err, os.ErrNotExist):
		return nil, err
	}
	key = make([]byte, sessionKeySize)
	mustRandom(key)
	// O_EXCL so two instances starting together cannot both write one.
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return SessionKey(dir) // the other instance won; read theirs
		}
		return nil, err
	}
	defer f.Close()
	if _, err := f.Write(key); err != nil {
		return nil, err
	}
	return key, f.Sync()
}

type challenge struct {
	circleID string
	pub      ed25519.PublicKey
	expires  time.Time
}

// Auth issues challenges and mints session tokens. The only state it holds is
// the challenges, which live for two minutes and are cheap to lose.
type Auth struct {
	key []byte

	mu         sync.Mutex
	challenges map[string]challenge
	now        func() time.Time // swapped in tests
}

func NewAuth(key []byte) *Auth {
	return &Auth{
		key:        key,
		challenges: map[string]challenge{},
		now:        time.Now,
	}
}

// Challenge issues a nonce for one public key on one circle. It is accepted
// exactly once; ADR-0015 rejects a client-chosen timestamp for this reason.
func (a *Auth) Challenge(circleID string, pub ed25519.PublicKey) string {
	c := NewToken()
	a.mu.Lock()
	defer a.mu.Unlock()
	a.sweep()
	a.challenges[c] = challenge{circleID: circleID, pub: pub, expires: a.now().Add(challengeTTL)}
	return c
}

// Session consumes a challenge and, if the signature verifies, returns a
// bearer token and its expiry.
func (a *Auth) Session(circleID, chal string, pub ed25519.PublicKey, sig []byte) (string, time.Time, error) {
	a.mu.Lock()
	c, ok := a.challenges[chal]
	// Burn it whether or not it verifies: one challenge, one attempt.
	delete(a.challenges, chal)
	now := a.now()
	a.mu.Unlock()

	if !ok || now.After(c.expires) || c.circleID != circleID || !c.pub.Equal(pub) {
		return "", time.Time{}, ErrBadChallenge
	}
	if !ed25519.Verify(pub, SignedMessage(circleID, chal), sig) {
		return "", time.Time{}, ErrBadSignature
	}

	expires := now.Add(sessionTTL)
	return a.mint(circleID, pub, expires), expires, nil
}

// Lookup returns the public key a token stands for on this circle.
//
// The token carries its claims and is trusted only because the MAC verifies, so
// this survives a restart and works across instances sharing the data
// directory. The cost is that one token cannot be revoked before it expires;
// removing a member is what actually revokes access, and http.go checks
// membership on every request for exactly that reason.
func (a *Auth) Lookup(circleID, token string) (ed25519.PublicKey, error) {
	version, rest, ok := strings.Cut(token, ".")
	if !ok || version != tokenVersion {
		return nil, ErrNoSession
	}
	body, mac, ok := strings.Cut(rest, ".")
	if !ok {
		return nil, ErrNoSession
	}
	claims, err := base64.RawURLEncoding.DecodeString(body)
	if err != nil {
		return nil, ErrNoSession
	}
	got, err := base64.RawURLEncoding.DecodeString(mac)
	if err != nil {
		return nil, ErrNoSession
	}
	if !hmac.Equal(got, a.mac(claims)) {
		return nil, ErrNoSession
	}

	// claims: expiry (8 bytes) | public key (32 bytes) | circle id
	if len(claims) < 8+ed25519.PublicKeySize {
		return nil, ErrNoSession
	}
	expires := time.UnixMilli(int64(binary.BigEndian.Uint64(claims[:8])))
	pub := ed25519.PublicKey(claims[8 : 8+ed25519.PublicKeySize])
	circle := string(claims[8+ed25519.PublicKeySize:])

	if circle != circleID || a.now().After(expires) {
		return nil, ErrNoSession
	}
	return pub, nil
}

func (a *Auth) mint(circleID string, pub ed25519.PublicKey, expires time.Time) string {
	claims := make([]byte, 0, 8+ed25519.PublicKeySize+len(circleID))
	claims = binary.BigEndian.AppendUint64(claims, uint64(expires.UnixMilli()))
	claims = append(claims, pub...)
	claims = append(claims, circleID...)
	return tokenVersion + "." +
		base64.RawURLEncoding.EncodeToString(claims) + "." +
		base64.RawURLEncoding.EncodeToString(a.mac(claims))
}

func (a *Auth) mac(claims []byte) []byte {
	h := hmac.New(sha256.New, a.key)
	h.Write(claims)
	return h.Sum(nil)
}

// sweep drops expired challenges. Callers hold the lock.
//
// ponytail: O(n) on every challenge issued. n is people signing in within two
// minutes, which is single digits. Give it a heap if that stops being true.
func (a *Auth) sweep() {
	now := a.now()
	for k, c := range a.challenges {
		if now.After(c.expires) {
			delete(a.challenges, k)
		}
	}
}
