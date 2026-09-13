package nah

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// device is one phone: a keypair and a session token.
type device struct {
	pub   ed25519.PublicKey
	priv  ed25519.PrivateKey
	token string
}

func newDevice(t *testing.T) *device {
	t.Helper()
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	return &device{pub: pub, priv: priv}
}

func testKey(t *testing.T) []byte {
	t.Helper()
	key, err := SessionKey(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	return key
}

func newTestServer(t *testing.T) (*httptest.Server, *Store) {
	t.Helper()
	dir := t.TempDir()
	store, err := NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	key, err := SessionKey(dir)
	if err != nil {
		t.Fatal(err)
	}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	srv := httptest.NewServer(NewServer(store, NewAuth(key), log).Handler())
	t.Cleanup(func() {
		srv.Close()
		store.Close()
	})
	return srv, store
}

func call(t *testing.T, srv *httptest.Server, method, path, token string, body any) (int, []byte) {
	t.Helper()
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		r = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, srv.URL+path, r)
	if err != nil {
		t.Fatal(err)
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	res, err := srv.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	out, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	return res.StatusCode, out
}

// signIn runs the whole of ADR-0015 from the client side: ask for a challenge,
// sign it, exchange the signature for a token.
func (d *device) signIn(t *testing.T, srv *httptest.Server, circle string) {
	t.Helper()
	status, body := call(t, srv, "POST", "/v1/circles/"+circle+"/challenge", "",
		map[string]any{"public_key": d.pub})
	if status != http.StatusOK {
		t.Fatalf("challenge: status %d: %s", status, body)
	}
	var ch struct {
		Challenge string `json:"challenge"`
	}
	if err := json.Unmarshal(body, &ch); err != nil {
		t.Fatal(err)
	}

	sig := ed25519.Sign(d.priv, SignedMessage(circle, ch.Challenge))
	status, body = call(t, srv, "POST", "/v1/circles/"+circle+"/session", "", map[string]any{
		"public_key": d.pub,
		"challenge":  ch.Challenge,
		"signature":  sig,
	})
	if status != http.StatusOK {
		t.Fatalf("session: status %d: %s", status, body)
	}
	var s struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(body, &s); err != nil {
		t.Fatal(err)
	}
	d.token = s.Token
}

// TestTwoDevicesOneCircle is CDI-1835 without the phones: two devices join one
// circle, each posts a moment, and each sees both.
func TestTwoDevicesOneCircle(t *testing.T) {
	srv, store := newTestServer(t)
	circle, invite, err := store.Create()
	if err != nil {
		t.Fatal(err)
	}

	a, b := newDevice(t), newDevice(t)
	for _, d := range []*device{a, b} {
		if status, body := call(t, srv, "POST", "/v1/circles/"+circle+"/join", "",
			map[string]any{"invite": invite, "public_key": d.pub}); status != http.StatusNoContent {
			t.Fatalf("join: status %d: %s", status, body)
		}
		d.signIn(t, srv, circle)
	}

	// The server stores ciphertext. These bytes stand in for it; nothing here
	// ever tries to read them.
	for _, tc := range []struct {
		d    *device
		blob []byte
	}{{a, []byte("ciphertext from A")}, {b, []byte("ciphertext from B")}} {
		if status, body := call(t, srv, "POST", "/v1/circles/"+circle+"/moments", tc.d.token,
			map[string]any{"blob": tc.blob}); status != http.StatusCreated {
			t.Fatalf("post: status %d: %s", status, body)
		}
	}

	for name, d := range map[string]*device{"A": a, "B": b} {
		status, body := call(t, srv, "GET", "/v1/circles/"+circle+"/moments", d.token, nil)
		if status != http.StatusOK {
			t.Fatalf("%s feed: status %d: %s", name, status, body)
		}
		var feed []Moment
		if err := json.Unmarshal(body, &feed); err != nil {
			t.Fatal(err)
		}
		if len(feed) != 2 {
			t.Fatalf("%s sees %d moments, want 2", name, len(feed))
		}
		// Newest first.
		if !bytes.Equal(feed[0].Blob, []byte("ciphertext from B")) {
			t.Errorf("%s: feed is not newest first: %q", name, feed[0].Blob)
		}
		if feed[0].AuthorID == feed[1].AuthorID {
			t.Errorf("%s: both moments claim the same author", name)
		}
	}
}

func TestStrangerCannotRead(t *testing.T) {
	srv, store := newTestServer(t)
	circle, invite, err := store.Create()
	if err != nil {
		t.Fatal(err)
	}
	member := newDevice(t)
	call(t, srv, "POST", "/v1/circles/"+circle+"/join", "",
		map[string]any{"invite": invite, "public_key": member.pub})
	member.signIn(t, srv, circle)

	if status, _ := call(t, srv, "GET", "/v1/circles/"+circle+"/moments", "", nil); status != http.StatusUnauthorized {
		t.Errorf("no token: status %d, want 401", status)
	}
	if status, _ := call(t, srv, "GET", "/v1/circles/"+circle+"/moments", "not-a-token", nil); status != http.StatusUnauthorized {
		t.Errorf("bogus token: status %d, want 401", status)
	}

	// A session is bound to the circle it was issued for.
	other, _, err := store.Create()
	if err != nil {
		t.Fatal(err)
	}
	if status, _ := call(t, srv, "GET", "/v1/circles/"+other+"/moments", member.token, nil); status != http.StatusUnauthorized {
		t.Errorf("token reused on another circle: status %d, want 401", status)
	}
}

func TestChallengeIsSingleUse(t *testing.T) {
	auth := NewAuth(testKey(t))
	d := newDevice(t)
	ch := auth.Challenge("c", d.pub)
	sig := ed25519.Sign(d.priv, SignedMessage("c", ch))

	if _, _, err := auth.Session("c", ch, d.pub, sig); err != nil {
		t.Fatalf("first use: %v", err)
	}
	if _, _, err := auth.Session("c", ch, d.pub, sig); err == nil {
		t.Error("a replayed challenge was accepted")
	}
}

func TestSignatureMustVerify(t *testing.T) {
	auth := NewAuth(testKey(t))
	d, impostor := newDevice(t), newDevice(t)

	ch := auth.Challenge("c", d.pub)
	if _, _, err := auth.Session("c", ch, d.pub, ed25519.Sign(impostor.priv, SignedMessage("c", ch))); err == nil {
		t.Error("a signature from the wrong key was accepted")
	}

	// A signature over a different circle's message must not open this one.
	ch = auth.Challenge("c", d.pub)
	if _, _, err := auth.Session("c", ch, d.pub, ed25519.Sign(d.priv, SignedMessage("elsewhere", ch))); err == nil {
		t.Error("a signature bound to another circle was accepted")
	}
}

func TestMemberCap(t *testing.T) {
	store, err := NewStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	circle, invite, err := store.Create()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < MemberCap; i++ {
		pub, _, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			t.Fatal(err)
		}
		if err := store.Join(circle, invite, pub); err != nil {
			t.Fatalf("member %d: %v", i, err)
		}
	}
	pub, _, _ := ed25519.GenerateKey(rand.Reader)
	if err := store.Join(circle, invite, pub); err != ErrCircleFull {
		t.Fatalf("member %d: got %v, want ErrCircleFull", MemberCap+1, err)
	}
}

func TestCircleIDIsNotAPath(t *testing.T) {
	for _, id := range []string{
		"../../etc/passwd",
		"..",
		"",
		"0192f0a1-1c2d-7e3f-8a4b-5c6d7e8f9a0",   // one short
		"0192f0a1-1c2d-7e3f-8a4b-5c6d7e8f9a0bb", // one long
		"0192f0a1-1c2d-7e3f-8a4b-5c6d7e8f9aZb",  // not hex
		"0192F0A1-1C2D-7E3F-8A4B-5C6D7E8F9A0B",  // upper case: filenames are case-sensitive
		"0192f0a1_1c2d_7e3f_8a4b_5c6d7e8f9a0b",
	} {
		if ValidID(id) {
			t.Errorf("ValidID(%q) = true, want false", id)
		}
	}
	if id := NewID(); !ValidID(id) {
		t.Errorf("ValidID(NewID()) = false for %q", id)
	}

	store, err := NewStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	if _, err := store.Feed("../../etc/passwd", 10); err != ErrBadCircleID {
		t.Errorf("Feed with a traversal id: got %v, want ErrBadCircleID", err)
	}
}

func TestMigrationsAreIdempotent(t *testing.T) {
	dir := t.TempDir()
	store, err := NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	circle, invite, err := store.Create()
	if err != nil {
		t.Fatal(err)
	}
	d := newDevice(t)
	if err := store.Join(circle, invite, d.pub); err != nil {
		t.Fatal(err)
	}
	if _, err := store.Post(circle, d.pub, []byte("ciphertext")); err != nil {
		t.Fatal(err)
	}
	store.Close()

	// Reopening steps user_version, finds it already at the end, and changes nothing.
	reopened, err := NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer reopened.Close()
	feed, err := reopened.Feed(circle, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(feed) != 1 {
		t.Fatalf("after reopen: %d moments, want 1", len(feed))
	}
}

func TestBadInvite(t *testing.T) {
	store, err := NewStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	circle, _, err := store.Create()
	if err != nil {
		t.Fatal(err)
	}
	d := newDevice(t)
	if err := store.Join(circle, "not-an-invite", d.pub); err != ErrBadInvite {
		t.Fatalf("got %v, want ErrBadInvite", err)
	}
}

// TestSessionSurvivesARestart is the point of ADR-0015's revision: deploys are
// frequent, and a token minted before one has to still work after it.
func TestSessionSurvivesARestart(t *testing.T) {
	dir := t.TempDir()
	key, err := SessionKey(dir)
	if err != nil {
		t.Fatal(err)
	}
	d := newDevice(t)

	before := NewAuth(key)
	ch := before.Challenge("c", d.pub)
	token, _, err := before.Session("c", ch, d.pub, ed25519.Sign(d.priv, SignedMessage("c", ch)))
	if err != nil {
		t.Fatal(err)
	}

	// A fresh process, reading the same key off the same volume.
	reread, err := SessionKey(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(key, reread) {
		t.Fatal("SessionKey generated a second key instead of reading the first")
	}
	after := NewAuth(reread)
	pub, err := after.Lookup("c", token)
	if err != nil {
		t.Fatalf("token did not survive the restart: %v", err)
	}
	if !pub.Equal(d.pub) {
		t.Error("token came back as the wrong device")
	}

	// Still bound to its circle, and worthless under a different key.
	if _, err := after.Lookup("elsewhere", token); err == nil {
		t.Error("token opened another circle")
	}
	if _, err := NewAuth(testKey(t)).Lookup("c", token); err == nil {
		t.Error("token verified under a different signing key")
	}
}

func TestTamperedTokenIsRejected(t *testing.T) {
	auth := NewAuth(testKey(t))
	d, impostor := newDevice(t), newDevice(t)
	ch := auth.Challenge("c", d.pub)
	token, _, err := auth.Session("c", ch, d.pub, ed25519.Sign(d.priv, SignedMessage("c", ch)))
	if err != nil {
		t.Fatal(err)
	}

	parts := strings.SplitN(token, ".", 3)
	if len(parts) != 3 {
		t.Fatalf("token is not v1.claims.mac: %q", token)
	}
	claims, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		t.Fatal(err)
	}

	// Swap in another device's key, keeping the original MAC.
	forged := append([]byte{}, claims[:8]...)
	forged = append(forged, impostor.pub...)
	forged = append(forged, claims[8+ed25519.PublicKeySize:]...)
	swapped := parts[0] + "." + base64.RawURLEncoding.EncodeToString(forged) + "." + parts[2]
	if _, err := auth.Lookup("c", swapped); err == nil {
		t.Error("a token with someone else's key was accepted")
	}

	// Push the expiry out by a century, keeping the original MAC.
	extended := append([]byte{}, claims...)
	binary.BigEndian.PutUint64(extended[:8], uint64(time.Now().Add(100*365*24*time.Hour).UnixMilli()))
	longer := parts[0] + "." + base64.RawURLEncoding.EncodeToString(extended) + "." + parts[2]
	if _, err := auth.Lookup("c", longer); err == nil {
		t.Error("a token with a stretched expiry was accepted")
	}

	for _, bad := range []string{"", ".", "v1", "v1.", "v2." + parts[1] + "." + parts[2], parts[1] + "." + parts[2]} {
		if _, err := auth.Lookup("c", bad); err == nil {
			t.Errorf("malformed token %q was accepted", bad)
		}
	}
}

func TestExpiredTokenIsRejected(t *testing.T) {
	auth := NewAuth(testKey(t))
	d := newDevice(t)
	ch := auth.Challenge("c", d.pub)
	token, expires, err := auth.Session("c", ch, d.pub, ed25519.Sign(d.priv, SignedMessage("c", ch)))
	if err != nil {
		t.Fatal(err)
	}
	auth.now = func() time.Time { return expires.Add(time.Second) }
	if _, err := auth.Lookup("c", token); err == nil {
		t.Error("an expired token was accepted")
	}
}
