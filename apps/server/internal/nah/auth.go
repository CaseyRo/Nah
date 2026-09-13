package nah

import (
	"crypto/ed25519"
	"errors"
	"sync"
	"time"
)

// Sign-in per ADR-0015: no passwords. The device holds an Ed25519 key, the
// server hands out a challenge, and a signature over that challenge buys a
// session token.
//
// Challenges and sessions live in memory and die with the process. A restart
// costs each client one silent round trip, because the device still holds the
// key and can simply sign again.

const (
	challengeTTL = 2 * time.Minute
	sessionTTL   = 30 * 24 * time.Hour

	// signPrefix domain-separates this signature from anything else the same
	// identity key might ever sign, and binds it to one circle so a signature
	// captured on one circle cannot open another.
	signPrefix = "nah-auth-v1:"
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

type grant struct {
	circleID string
	pub      ed25519.PublicKey
	expires  time.Time
}

// Auth holds the short-lived challenges and the live sessions.
type Auth struct {
	mu         sync.Mutex
	challenges map[string]grant
	sessions   map[string]grant
	now        func() time.Time // swapped in tests
}

func NewAuth() *Auth {
	return &Auth{
		challenges: map[string]grant{},
		sessions:   map[string]grant{},
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
	a.challenges[c] = grant{circleID: circleID, pub: pub, expires: a.now().Add(challengeTTL)}
	return c
}

// Session consumes a challenge and, if the signature verifies, returns a
// bearer token and its expiry.
func (a *Auth) Session(circleID, challenge string, pub ed25519.PublicKey, sig []byte) (string, time.Time, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	g, ok := a.challenges[challenge]
	// Burn it whether or not it verifies: one challenge, one attempt.
	delete(a.challenges, challenge)
	if !ok || a.now().After(g.expires) || g.circleID != circleID || !g.pub.Equal(pub) {
		return "", time.Time{}, ErrBadChallenge
	}
	if !ed25519.Verify(pub, SignedMessage(circleID, challenge), sig) {
		return "", time.Time{}, ErrBadSignature
	}

	a.sweep()
	token := NewToken()
	expires := a.now().Add(sessionTTL)
	a.sessions[token] = grant{circleID: circleID, pub: pub, expires: expires}
	return token, expires, nil
}

// Lookup returns the public key a token stands for on this circle. A session
// is bound to the circle it was issued for and is worthless on any other.
func (a *Auth) Lookup(circleID, token string) (ed25519.PublicKey, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	g, ok := a.sessions[token]
	if !ok || g.circleID != circleID {
		return nil, ErrNoSession
	}
	if a.now().After(g.expires) {
		delete(a.sessions, token)
		return nil, ErrNoSession
	}
	return g.pub, nil
}

// sweep drops expired entries. Callers hold the lock.
//
// ponytail: O(n) over both maps on every issue. n is members per host, which is
// tens. Give it a clock and a heap when that stops being true.
func (a *Auth) sweep() {
	now := a.now()
	for k, g := range a.challenges {
		if now.After(g.expires) {
			delete(a.challenges, k)
		}
	}
	for k, g := range a.sessions {
		if now.After(g.expires) {
			delete(a.sessions, k)
		}
	}
}
