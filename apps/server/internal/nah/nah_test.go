package nah

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	mathrand "math/rand/v2"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"sort"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// device is one phone: a keypair, the person it registered as, and a session.
type device struct {
	pub    ed25519.PublicKey
	priv   ed25519.PrivateKey
	person string
	token  string
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

func newStore(t *testing.T) *Store {
	t.Helper()
	store, err := NewStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(store.Close)
	return store
}

// register makes a person with an operator's invite, the way the first person on
// a server arrives, for tests that are about the store rather than the routes.
func register(t *testing.T, store *Store) string {
	t.Helper()
	invite, err := store.OperatorInvite()
	if err != nil {
		t.Fatal(err)
	}
	id, err := store.Register(newDevice(t).pub, "", invite)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func invite(t *testing.T, store *Store, person string) string {
	t.Helper()
	token, err := store.Invite(person)
	if err != nil {
		t.Fatal(err)
	}
	return token
}

func mustPost(t *testing.T, store *Store, person, blob string) {
	t.Helper()
	if _, err := store.Post(person, []byte(blob)); err != nil {
		t.Fatal(err)
	}
}

// personFiles counts the person files in the store's directory.
func personFiles(t *testing.T, store *Store) int {
	t.Helper()
	files, err := filepath.Glob(filepath.Join(store.dir, "*.db"))
	if err != nil {
		t.Fatal(err)
	}
	return len(files)
}

// ticking is a clock that moves on a second every time it is read, so no two
// moments in a test share a time and newest first is never decided by a tie.
func ticking() func() time.Time {
	var n atomic.Int64
	start := time.Date(2026, 9, 14, 12, 0, 0, 0, time.UTC)
	return func() time.Time { return start.Add(time.Duration(n.Add(1)) * time.Second) }
}

func newTestServer(t *testing.T) (*httptest.Server, *Store) {
	t.Helper()
	store := newStore(t)
	store.now = ticking()
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	srv := httptest.NewServer(NewServer(store, NewAuth(testKey(t)), log).Handler())
	t.Cleanup(srv.Close)
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

// join is a phone's first minute: register with an invitation, then sign in the
// way ADR-0015 describes. With no inviter the invitation is an operator's, as
// for the first person on a server; with one it is theirs, which connects the
// two of them in the same step.
func join(t *testing.T, srv *httptest.Server, store *Store, inviter *device) *device {
	t.Helper()
	d := newDevice(t)
	body := map[string]any{"public_key": d.pub}
	if inviter == nil {
		token, err := store.OperatorInvite()
		if err != nil {
			t.Fatal(err)
		}
		body["invite"] = token
	} else {
		body["person"], body["invite"] = inviter.person, inviter.invite(t, srv)
	}
	status, resp := call(t, srv, "POST", "/v1/people", "", body)
	if status != http.StatusCreated {
		t.Fatalf("register: status %d: %s", status, resp)
	}
	var reg struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(resp, &reg); err != nil {
		t.Fatal(err)
	}
	d.person = reg.ID

	status, resp = call(t, srv, "POST", "/v1/people/"+d.person+"/challenge", "",
		map[string]any{"public_key": d.pub})
	if status != http.StatusOK {
		t.Fatalf("challenge: status %d: %s", status, resp)
	}
	var ch struct {
		Challenge string `json:"challenge"`
	}
	if err := json.Unmarshal(resp, &ch); err != nil {
		t.Fatal(err)
	}

	sig := ed25519.Sign(d.priv, SignedMessage(d.person, ch.Challenge))
	status, resp = call(t, srv, "POST", "/v1/people/"+d.person+"/session", "", map[string]any{
		"public_key": d.pub,
		"challenge":  ch.Challenge,
		"signature":  sig,
	})
	if status != http.StatusOK {
		t.Fatalf("session: status %d: %s", status, resp)
	}
	var s struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(resp, &s); err != nil {
		t.Fatal(err)
	}
	d.token = s.Token
	return d
}

// invite asks the server for an invite from d.
func (d *device) invite(t *testing.T, srv *httptest.Server) string {
	t.Helper()
	status, body := call(t, srv, "POST", "/v1/people/"+d.person+"/invites", d.token, nil)
	if status != http.StatusCreated {
		t.Fatalf("invite: status %d: %s", status, body)
	}
	var inv struct {
		Invite string `json:"invite"`
	}
	if err := json.Unmarshal(body, &inv); err != nil {
		t.Fatal(err)
	}
	return inv.Invite
}

// connect is what a touch or a link carries between two people who are both
// already here: a makes an invite and b redeems it.
func connect(t *testing.T, srv *httptest.Server, a, b *device) {
	t.Helper()
	if status, body := call(t, srv, "POST", "/v1/people/"+b.person+"/connections", b.token,
		map[string]any{"person": a.person, "invite": a.invite(t, srv)}); status != http.StatusNoContent {
		t.Fatalf("connect: status %d: %s", status, body)
	}
}

func (d *device) post(t *testing.T, srv *httptest.Server, blob string) {
	t.Helper()
	if status, body := call(t, srv, "POST", "/v1/people/"+d.person+"/moments", d.token,
		map[string]any{"blob": []byte(blob)}); status != http.StatusCreated {
		t.Fatalf("post: status %d: %s", status, body)
	}
}

func (d *device) feed(t *testing.T, srv *httptest.Server) []Moment {
	t.Helper()
	status, body := call(t, srv, "GET", "/v1/people/"+d.person+"/feed", d.token, nil)
	if status != http.StatusOK {
		t.Fatalf("feed: status %d: %s", status, body)
	}
	var feed []Moment
	if err := json.Unmarshal(body, &feed); err != nil {
		t.Fatal(err)
	}
	return feed
}

// TestTwoPeopleConnected is CDI-1835 without the phones: one person joins
// through the other's invitation, each posts a moment, and each sees both.
func TestTwoPeopleConnected(t *testing.T) {
	srv, store := newTestServer(t)
	a := join(t, srv, store, nil)
	b := join(t, srv, store, a)

	// The server stores ciphertext. These bytes stand in for it; nothing here
	// ever tries to read them.
	a.post(t, srv, "ciphertext from A")
	b.post(t, srv, "ciphertext from B")

	for name, d := range map[string]*device{"A": a, "B": b} {
		feed := d.feed(t, srv)
		if len(feed) != 2 {
			t.Fatalf("%s sees %d moments, want 2", name, len(feed))
		}
		// Newest first.
		if string(feed[0].Blob) != "ciphertext from B" {
			t.Errorf("%s: feed is not newest first: %q", name, feed[0].Blob)
		}
		if feed[0].AuthorID != b.person || feed[1].AuthorID != a.person {
			t.Errorf("%s: authors are %s and %s, want B then A", name, feed[0].AuthorID, feed[1].AuthorID)
		}
	}
}

// TestPeopleAlreadyHereConnect covers the other way in: two people who have
// both joined connect through /connections.
func TestPeopleAlreadyHereConnect(t *testing.T) {
	srv, store := newTestServer(t)
	a, b := join(t, srv, store, nil), join(t, srv, store, nil)
	a.post(t, srv, "before they met")
	connect(t, srv, a, b)

	if feed := b.feed(t, srv); len(feed) != 1 || string(feed[0].Blob) != "before they met" {
		t.Errorf("after connecting, B's feed is %d moments, want A's one", len(feed))
	}
}

// TestRegistrationIsByInvitation holds the rule that nobody arrives uninvited,
// and that a refused registration leaves no person behind.
func TestRegistrationIsByInvitation(t *testing.T) {
	srv, store := newTestServer(t)
	try := func(body map[string]any) int {
		t.Helper()
		body["public_key"] = newDevice(t).pub
		status, _ := call(t, srv, "POST", "/v1/people", "", body)
		return status
	}

	if status := try(map[string]any{}); status != http.StatusBadRequest {
		t.Errorf("no invite: status %d, want 400", status)
	}
	if status := try(map[string]any{"invite": "made-up"}); status != http.StatusForbidden {
		t.Errorf("made-up operator invite: status %d, want 403", status)
	}

	token, err := store.OperatorInvite()
	if err != nil {
		t.Fatal(err)
	}
	if status := try(map[string]any{"invite": token}); status != http.StatusCreated {
		t.Fatalf("operator invite: status %d, want 201", status)
	}
	if status := try(map[string]any{"invite": token}); status != http.StatusForbidden {
		t.Errorf("operator invite used twice: status %d, want 403", status)
	}

	a := join(t, srv, store, nil)
	before := personFiles(t, store)
	if status := try(map[string]any{"person": a.person, "invite": "made-up"}); status != http.StatusForbidden {
		t.Errorf("made-up invite from a person: status %d, want 403", status)
	}
	if status := try(map[string]any{"person": NewID(), "invite": "made-up"}); status != http.StatusForbidden {
		t.Errorf("invite from nobody: status %d, want 403", status)
	}
	if after := personFiles(t, store); after != before {
		t.Errorf("refused registrations left %d person files behind", after-before)
	}
}

func TestStrangerCannotRead(t *testing.T) {
	srv, store := newTestServer(t)
	a, stranger := join(t, srv, store, nil), join(t, srv, store, nil)
	a.post(t, srv, "only for the people A is connected to")

	feed := "/v1/people/" + a.person + "/feed"
	if status, _ := call(t, srv, "GET", feed, "", nil); status != http.StatusUnauthorized {
		t.Errorf("no token: status %d, want 401", status)
	}
	if status, _ := call(t, srv, "GET", feed, "not-a-token", nil); status != http.StatusUnauthorized {
		t.Errorf("bogus token: status %d, want 401", status)
	}

	// A session is bound to the person it was issued for.
	if status, _ := call(t, srv, "GET", feed, stranger.token, nil); status != http.StatusUnauthorized {
		t.Errorf("someone else's token: status %d, want 401", status)
	}
	if status, _ := call(t, srv, "POST", "/v1/people/"+a.person+"/moments", stranger.token,
		map[string]any{"blob": []byte("forged")}); status != http.StatusUnauthorized {
		t.Errorf("posting as someone else: status %d, want 401", status)
	}

	// And a feed holds only the people its reader is connected to.
	if got := stranger.feed(t, srv); len(got) != 0 {
		t.Errorf("a stranger's own feed shows %d moments from someone they never connected to", len(got))
	}
}

// TestFeedMatchesReadingEveryFile holds Feed's shortcut to account. It stops
// reading once nobody left can have anything newer than its page (CDI-1879),
// and that has to give exactly the page that reading everybody's file would.
func TestFeedMatchesReadingEveryFile(t *testing.T) {
	store := newStore(t)
	rng := mathrand.New(mathrand.NewPCG(1, 2))
	var at time.Time
	store.now = func() time.Time { return at }

	const people, most = 40, 40
	persons := make([]string, people)
	for i := range persons {
		persons[i] = register(t, store)
	}

	// Every moment gets a second of its own, posted in no particular order, so a
	// file's row ids say nothing about time.
	base := time.Date(2025, 9, 14, 0, 0, 0, 0, time.UTC)
	seconds := rng.Perm(people * most)
	posted := map[string][]Moment{}
	for i, p := range persons {
		for j := range rng.IntN(most + 1) {
			at = base.Add(time.Duration(seconds[i*most+j]) * time.Second)
			m, err := store.Post(p, fmt.Appendf(nil, "moment %d by %s", j, p))
			if err != nil {
				t.Fatal(err)
			}
			posted[p] = append(posted[p], m)
		}
	}

	connections := map[string][]string{}
	for i, a := range persons {
		for _, b := range persons[i+1:] {
			if rng.IntN(3) == 0 {
				if err := store.Connect(b, a, invite(t, store, a)); err != nil {
					t.Fatal(err)
				}
				connections[a] = append(connections[a], b)
				connections[b] = append(connections[b], a)
			}
		}
	}

	for _, limit := range []int{1, 10, 200} {
		for _, reader := range persons {
			want := append([]Moment{}, posted[reader]...)
			for _, c := range connections[reader] {
				want = append(want, posted[c]...)
			}
			sort.Slice(want, func(i, j int) bool { return newer(want[i], want[j]) })
			want = want[:min(limit, len(want))]

			got, err := store.Feed(reader, limit)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != len(want) {
				t.Fatalf("limit %d, %s: feed has %d moments, reading every file gives %d", limit, reader, len(got), len(want))
			}
			for i := range want {
				g, w := got[i], want[i]
				if g.AuthorID != w.AuthorID || g.ID != w.ID || g.CreatedAt != w.CreatedAt || !bytes.Equal(g.Blob, w.Blob) {
					t.Fatalf("limit %d, %s: moment %d is %s/%d, reading every file gives %s/%d",
						limit, reader, i, g.AuthorID, g.ID, w.AuthorID, w.ID)
				}
			}
		}
	}
}

// TestPostFromAnotherInstanceArrives is a rolling deploy in miniature: two
// stores on one volume, a moment taken by the old one, and a feed read from the
// new one. Feed trusts what it remembers about who posted last for
// newestRecheck, so after that the moment has to be there.
func TestPostFromAnotherInstanceArrives(t *testing.T) {
	dir := t.TempDir()
	stores := make([]*Store, 2)
	for i := range stores {
		s, err := NewStore(dir)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(s.Close)
		stores[i] = s
	}
	old, neu := stores[0], stores[1]
	now := time.Date(2026, 9, 14, 12, 0, 0, 0, time.UTC)
	old.now = func() time.Time { return now }
	neu.now = old.now

	reader, author, other := register(t, neu), register(t, neu), register(t, neu)
	for _, p := range []string{author, other} {
		if err := neu.Connect(reader, p, invite(t, neu, p)); err != nil {
			t.Fatal(err)
		}
	}
	mustPost(t, neu, author, "before the deploy")
	now = now.Add(500 * time.Millisecond)
	mustPost(t, neu, other, "a little later")

	// Reading a feed teaches the new instance who posted last.
	if _, err := neu.Feed(reader, 1); err != nil {
		t.Fatal(err)
	}

	// The old instance, still draining, takes a newer moment from author.
	now = now.Add(500 * time.Millisecond)
	mustPost(t, old, author, "during the deploy")

	now = now.Add(newestRecheck)
	feed, err := neu.Feed(reader, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(feed) != 1 {
		t.Fatalf("feed has %d moments, want 1", len(feed))
	}
	if got := string(feed[0].Blob); got != "during the deploy" {
		t.Fatalf("a recheck later, the new instance's feed starts with %q; want the moment the old one took", got)
	}
}

// TestNetworkCap is ADR-0017's cap: 150, counted by the server, on both sides.
func TestNetworkCap(t *testing.T) {
	store := newStore(t)
	hub := register(t, store)
	for i := range NetworkCap {
		if err := store.Connect(register(t, store), hub, invite(t, store, hub)); err != nil {
			t.Fatalf("connection %d: %v", i+1, err)
		}
	}
	mustPost(t, store, hub, "from the hub")

	// Full as the one who made the invite: nothing is written on either side.
	late, hubInvite := register(t, store), invite(t, store, hub)
	if err := store.Connect(late, hub, hubInvite); !errors.Is(err, ErrNetworkFull) {
		t.Fatalf("connection %d: got %v, want ErrNetworkFull", NetworkCap+1, err)
	}
	assertUnused(t, store, hub, hubInvite)

	// Joining through a full network's invite is refused, and leaves no person.
	before := personFiles(t, store)
	if _, err := store.Register(newDevice(t).pub, hub, hubInvite); !errors.Is(err, ErrNetworkFull) {
		t.Fatalf("joining through a full network: got %v, want ErrNetworkFull", err)
	}
	if after := personFiles(t, store); after != before {
		t.Errorf("a refused registration left %d person files behind", after-before)
	}

	// Full as the one redeeming: the inviter's side is written first and has to
	// be undone, or the inviter would be reading the hub one-way.
	inviter := register(t, store)
	fromInviter := invite(t, store, inviter)
	if err := store.Connect(hub, inviter, fromInviter); !errors.Is(err, ErrNetworkFull) {
		t.Fatalf("redeeming while full: got %v, want ErrNetworkFull", err)
	}
	assertUnused(t, store, inviter, fromInviter)

	for name, p := range map[string]string{"late": late, "inviter": inviter} {
		feed, err := store.Feed(p, 10)
		if err != nil {
			t.Fatal(err)
		}
		if len(feed) != 0 {
			t.Errorf("%s can read the hub after the connection was refused", name)
		}
	}
}

// assertUnused fails unless a refused connection left the invite unspent.
func assertUnused(t *testing.T, store *Store, owner, token string) {
	t.Helper()
	p, err := store.person(owner)
	if err != nil {
		t.Fatal(err)
	}
	var used sql.NullInt64
	if err := p.read.QueryRow(`SELECT used_at FROM invites WHERE token = ?`, token).Scan(&used); err != nil {
		t.Fatal(err)
	}
	if used.Valid {
		t.Errorf("a refused connection spent the invite")
	}
}

// TestInviteWorksOnce holds CDI-1883's rule: every invitation works exactly
// once, a retry of the connection it made is harmless, and nobody else can use
// it afterwards, whether to connect or to arrive.
func TestInviteWorksOnce(t *testing.T) {
	store := newStore(t)
	a, b, c := register(t, store), register(t, store), register(t, store)
	token := invite(t, store, a)
	if err := store.Connect(b, a, token); err != nil {
		t.Fatal(err)
	}
	if err := store.Connect(b, a, token); err != nil {
		t.Errorf("retrying the same connection: got %v, want nil", err)
	}
	if err := store.Connect(c, a, token); !errors.Is(err, ErrUsedInvite) {
		t.Errorf("a second person with a used invite: got %v, want ErrUsedInvite", err)
	}
	before := personFiles(t, store)
	if _, err := store.Register(newDevice(t).pub, a, token); !errors.Is(err, ErrUsedInvite) {
		t.Errorf("arriving with a used invite: got %v, want ErrUsedInvite", err)
	}
	if after := personFiles(t, store); after != before {
		t.Errorf("a refused registration left %d person files behind", after-before)
	}
}

// TestFeedNamesAuthors is CDI-1896 on the server: a profile set by its owner
// comes back, byte for byte and unread, on every moment of theirs in a feed.
func TestFeedNamesAuthors(t *testing.T) {
	srv, store := newTestServer(t)
	a := join(t, srv, store, nil)
	b := join(t, srv, store, a)
	if status, body := call(t, srv, "PUT", "/v1/people/"+a.person+"/profile", a.token,
		map[string]any{"blob": []byte("sealed: Maya")}); status != http.StatusNoContent {
		t.Fatalf("profile: status %d: %s", status, body)
	}
	a.post(t, srv, "from a")
	b.post(t, srv, "from b")
	for _, m := range b.feed(t, srv) {
		want := ""
		if m.AuthorID == a.person {
			want = "sealed: Maya"
		}
		if string(m.AuthorProfile) != want {
			t.Errorf("%s's moment carries profile %q, want %q", m.AuthorID, m.AuthorProfile, want)
		}
	}
	if status, _ := call(t, srv, "PUT", "/v1/people/"+a.person+"/profile", b.token,
		map[string]any{"blob": []byte("not yours")}); status != http.StatusUnauthorized {
		t.Errorf("setting someone else's profile: status %d, want 401", status)
	}
}

func TestBadInvite(t *testing.T) {
	store := newStore(t)
	a, b, c := register(t, store), register(t, store), register(t, store)
	if err := store.Connect(b, a, "not-an-invite"); !errors.Is(err, ErrBadInvite) {
		t.Errorf("made-up invite: got %v, want ErrBadInvite", err)
	}
	fromA := invite(t, store, a)
	if err := store.Connect(b, c, fromA); !errors.Is(err, ErrBadInvite) {
		t.Errorf("A's invite presented as C's: got %v, want ErrBadInvite", err)
	}
	if err := store.Connect(a, a, fromA); !errors.Is(err, ErrOwnInvite) {
		t.Errorf("own invite: got %v, want ErrOwnInvite", err)
	}
	if _, err := store.Register(newDevice(t).pub, "", ""); !errors.Is(err, ErrNoInvite) {
		t.Errorf("registering with no invite: got %v, want ErrNoInvite", err)
	}
	if _, err := store.Register(newDevice(t).pub, a, "not-an-invite"); !errors.Is(err, ErrBadInvite) {
		t.Errorf("registering with a made-up invite: got %v, want ErrBadInvite", err)
	}
}

func TestChallengeIsSingleUse(t *testing.T) {
	auth := NewAuth(testKey(t))
	d := newDevice(t)
	ch := auth.Challenge("p", d.pub)
	sig := ed25519.Sign(d.priv, SignedMessage("p", ch))

	if _, _, err := auth.Session("p", ch, d.pub, sig); err != nil {
		t.Fatalf("first use: %v", err)
	}
	if _, _, err := auth.Session("p", ch, d.pub, sig); err == nil {
		t.Error("a replayed challenge was accepted")
	}
}

func TestSignatureMustVerify(t *testing.T) {
	auth := NewAuth(testKey(t))
	d, impostor := newDevice(t), newDevice(t)

	ch := auth.Challenge("p", d.pub)
	if _, _, err := auth.Session("p", ch, d.pub, ed25519.Sign(impostor.priv, SignedMessage("p", ch))); err == nil {
		t.Error("a signature from the wrong key was accepted")
	}

	// A signature made for a different person must not sign in as this one.
	ch = auth.Challenge("p", d.pub)
	if _, _, err := auth.Session("p", ch, d.pub, ed25519.Sign(d.priv, SignedMessage("someone-else", ch))); err == nil {
		t.Error("a signature bound to another person was accepted")
	}
}

func TestPersonIDIsNotAPath(t *testing.T) {
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

	if _, err := newStore(t).Feed("../../etc/passwd", 10); !errors.Is(err, ErrBadPersonID) {
		t.Errorf("Feed with a traversal id: got %v, want ErrBadPersonID", err)
	}
}

func TestMigrationsAreIdempotent(t *testing.T) {
	dir := t.TempDir()
	store, err := NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	p := register(t, store)
	mustPost(t, store, p, "ciphertext")
	store.Close()

	// Reopening steps user_version, finds it already at the end, and changes nothing.
	reopened, err := NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer reopened.Close()
	feed, err := reopened.Feed(p, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(feed) != 1 {
		t.Fatalf("after reopen: %d moments, want 1", len(feed))
	}
}

// TestOldFilesMigrate opens a person file written before the profile and
// single-use invites existed, as every file on a server from M1 is, and checks
// the additive migration brings it forward with nothing lost.
func TestOldFilesMigrate(t *testing.T) {
	dir := t.TempDir()
	id := NewID()
	old, err := openDB(filepath.Join(dir, id+".db"), 1)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := old.Exec(migrations[0] + `PRAGMA user_version = 1;`); err != nil {
		t.Fatal(err)
	}
	if _, err := old.Exec(`INSERT INTO person (id, created_at, public_key) VALUES (?, 0, ?)`,
		id, []byte(newDevice(t).pub)); err != nil {
		t.Fatal(err)
	}
	if _, err := old.Exec(`INSERT INTO moments (created_at, blob) VALUES (1, 'kept')`); err != nil {
		t.Fatal(err)
	}
	old.Close()

	store, err := NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	if err := store.SetProfile(id, []byte("sealed: Maya")); err != nil {
		t.Fatalf("setting a profile on a migrated file: %v", err)
	}
	feed, err := store.Feed(id, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(feed) != 1 || string(feed[0].Blob) != "kept" || string(feed[0].AuthorProfile) != "sealed: Maya" {
		t.Fatalf("after migrating: %+v", feed)
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
	ch := before.Challenge("p", d.pub)
	token, _, err := before.Session("p", ch, d.pub, ed25519.Sign(d.priv, SignedMessage("p", ch)))
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
	pub, err := after.Lookup("p", token)
	if err != nil {
		t.Fatalf("token did not survive the restart: %v", err)
	}
	if !pub.Equal(d.pub) {
		t.Error("token came back as the wrong device")
	}

	// Still bound to its person, and worthless under a different key.
	if _, err := after.Lookup("someone-else", token); err == nil {
		t.Error("token signed in as another person")
	}
	if _, err := NewAuth(testKey(t)).Lookup("p", token); err == nil {
		t.Error("token verified under a different signing key")
	}
}

func TestTamperedTokenIsRejected(t *testing.T) {
	auth := NewAuth(testKey(t))
	d, impostor := newDevice(t), newDevice(t)
	ch := auth.Challenge("p", d.pub)
	token, _, err := auth.Session("p", ch, d.pub, ed25519.Sign(d.priv, SignedMessage("p", ch)))
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
	if _, err := auth.Lookup("p", swapped); err == nil {
		t.Error("a token with someone else's key was accepted")
	}

	// Push the expiry out by a century, keeping the original MAC.
	extended := append([]byte{}, claims...)
	binary.BigEndian.PutUint64(extended[:8], uint64(time.Now().Add(100*365*24*time.Hour).UnixMilli()))
	longer := parts[0] + "." + base64.RawURLEncoding.EncodeToString(extended) + "." + parts[2]
	if _, err := auth.Lookup("p", longer); err == nil {
		t.Error("a token with a stretched expiry was accepted")
	}

	for _, bad := range []string{"", ".", "v1", "v1.", "v2." + parts[1] + "." + parts[2], parts[1] + "." + parts[2]} {
		if _, err := auth.Lookup("p", bad); err == nil {
			t.Errorf("malformed token %q was accepted", bad)
		}
	}
}

func TestExpiredTokenIsRejected(t *testing.T) {
	auth := NewAuth(testKey(t))
	d := newDevice(t)
	ch := auth.Challenge("p", d.pub)
	token, expires, err := auth.Session("p", ch, d.pub, ed25519.Sign(d.priv, SignedMessage("p", ch)))
	if err != nil {
		t.Fatal(err)
	}
	auth.now = func() time.Time { return expires.Add(time.Second) }
	if _, err := auth.Lookup("p", token); err == nil {
		t.Error("an expired token was accepted")
	}
}
