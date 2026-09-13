package nah

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"time"
)

// NewID is a UUIDv7: 48 bits of millisecond timestamp, then randomness.
// Hand-rolled per ADR-0014 — this is the whole library.
//
// Time-ordered matters here because a circle id is also a filename, and
// a directory listing that sorts by creation time is worth forty lines.
func NewID() string {
	var b [16]byte
	mustRandom(b[:])
	ms := time.Now().UnixMilli()
	b[0] = byte(ms >> 40)
	b[1] = byte(ms >> 32)
	b[2] = byte(ms >> 24)
	b[3] = byte(ms >> 16)
	b[4] = byte(ms >> 8)
	b[5] = byte(ms)
	b[6] = (b[6] & 0x0f) | 0x70 // version 7
	b[8] = (b[8] & 0x3f) | 0x80 // RFC 4122 variant
	h := hex.EncodeToString(b[:])
	return h[0:8] + "-" + h[8:12] + "-" + h[12:16] + "-" + h[16:20] + "-" + h[20:32]
}

// ValidID is the guard on every circle id that arrives from a URL. A circle id
// becomes a path, so anything that is not exactly this shape is refused before
// it reaches the filesystem.
func ValidID(s string) bool {
	if len(s) != 36 {
		return false
	}
	for i := 0; i < 36; i++ {
		c := s[i]
		switch i {
		case 8, 13, 18, 23:
			if c != '-' {
				return false
			}
		default:
			isHex := (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')
			if !isHex {
				return false
			}
		}
	}
	return true
}

// NewToken is 32 random bytes, URL-safe. Used for invites, challenges and
// session tokens. At 256 bits there is nothing to guess.
func NewToken() string {
	var b [32]byte
	mustRandom(b[:])
	return base64.RawURLEncoding.EncodeToString(b[:])
}

func mustRandom(b []byte) {
	if _, err := rand.Read(b); err != nil {
		// crypto/rand does not fail on any platform we ship to, and a server
		// that cannot generate a key must not keep serving.
		panic("nah: no randomness available: " + err.Error())
	}
}
