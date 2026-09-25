// Package nah is the Nah? server: one SQLite file per person, holding that
// person's device key, their connections, their invites and the ciphertext of
// what they posted (ADR-0017).
//
// The server never reads a moment. It stores blobs, orders them by time and
// hands back the newest few. Every product feature that would need plaintext
// (search, ranking, counts, thumbnails) is refused by design, which is what
// makes ADR-0012 possible at all.
package nah

import (
	"crypto/ed25519"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	_ "modernc.org/sqlite"
)

// NetworkCap is Dunbar's number, applied to a person, which is what it describes
// (ADR-0017). It is a server-side rule, not a client check, counted on both
// sides of a connection, and a full network is told so in words.
const NetworkCap = 150

// newestRecheck is how long Feed trusts what it remembers about when someone
// last posted before reading it from their file again. Posts through this
// process update it at once; the recheck is for posts taken by a second process
// on the same volume, which is what a rolling deploy is (CDI-1876). A moment
// posted to the old instance reaches the new one's feeds within this long.
const newestRecheck = 30 * time.Second

// operatorInvites is the directory, inside the data directory, that holds the
// invites `server invite` makes for a server's first person.
const operatorInvites = "invites"

var (
	ErrNoPerson    = errors.New("no such person")
	ErrBadPersonID = errors.New("malformed person id")
	ErrNoInvite    = errors.New("registration needs an invite")
	ErrBadInvite   = errors.New("invite is not valid")
	ErrUsedInvite  = errors.New("invite has already been used")
	ErrOwnInvite   = errors.New("invite belongs to the person redeeming it")
	ErrNetworkFull = errors.New("network is full")
)

// migrations are applied in order, stepping PRAGMA user_version, when a
// person's file is opened. Per ADR-0014 this is hand-rolled: with one database
// per person there is no single database for a migration tool to own.
//
// Append only, and additive only: during a rolling deploy the old instance still
// reads files the new one has migrated (CDI-1876). Never edit a migration that
// has shipped.
var migrations = []string{`
CREATE TABLE person (
	id          TEXT PRIMARY KEY,
	created_at  INTEGER NOT NULL,
	public_key  BLOB NOT NULL        -- ed25519, 32 bytes, generated on the device (ADR-0015)
);

CREATE TABLE connections (
	person_id     TEXT PRIMARY KEY,  -- mutual: that person's file holds this one
	connected_at  INTEGER NOT NULL
);

CREATE TABLE invites (
	token       TEXT PRIMARY KEY,
	created_at  INTEGER NOT NULL,
	revoked_at  INTEGER
);

CREATE TABLE moments (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at  INTEGER NOT NULL,
	blob        BLOB NOT NULL        -- ciphertext, never read by the server (ADR-0012)
);

CREATE INDEX idx_moments_created ON moments (created_at DESC);
`, `
-- A person's profile: their name and, later, how they show themselves
-- (CDI-1895). Ciphertext like a moment, never read by the server.
ALTER TABLE person ADD COLUMN profile BLOB;

-- Every invitation works once (CDI-1883): who redeemed it, and when.
ALTER TABLE invites ADD COLUMN used_at INTEGER;
ALTER TABLE invites ADD COLUMN used_by TEXT;
`}

// Store owns the person files under one directory.
type Store struct {
	dir string
	now func() time.Time // swapped in tests

	mu   sync.Mutex
	open map[string]*personDB
}

// personDB keeps the split the spike measured: one writer, because SQLite
// serialises writes anyway, and several readers, which WAL allows.
type personDB struct {
	read  *sql.DB
	write *sql.DB

	// newest is the created_at of this person's newest moment, and checked is
	// when that was last read from the file. Both are Unix milliseconds.
	newest, checked atomic.Int64
}

func NewStore(dir string) (*Store, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, err
	}
	return &Store{dir: dir, now: time.Now, open: map[string]*personDB{}}, nil
}

func (s *Store) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, p := range s.open {
		p.read.Close()
		p.write.Close()
		delete(s.open, id)
	}
}

// Moment is one moment in a feed. Blob is ciphertext. ID is unique among its
// author's moments, not across a feed; (AuthorID, ID) is. AuthorProfile is the
// author's profile ciphertext, so a feed can name who posted (CDI-1895).
type Moment struct {
	ID            int64  `json:"id"`
	AuthorID      string `json:"author_id"`
	CreatedAt     int64  `json:"created_at"`
	Blob          []byte `json:"blob"`
	AuthorProfile []byte `json:"author_profile,omitempty"`
}

// Register makes a new person, and nobody arrives uninvited. The invite comes
// either from someone already here, which connects the two of them in the same
// step (invitation is connection, ADR-0017), or, when no inviter is named, from
// the operator, which is how the first person on a server arrives.
//
// The key is what signs in (ADR-0015). The id is what the file is called and
// what other people connect to, so it can outlive a phone.
func (s *Store) Register(pub ed25519.PublicKey, inviterID, invite string) (string, error) {
	if invite == "" {
		return "", ErrNoInvite
	}
	if inviterID == "" {
		// Spent before the person exists: a failure after this costs the operator
		// one more invite, and never leaves a person nobody invited.
		if err := os.Remove(filepath.Join(s.dir, operatorInvites, tokenHash(invite))); err != nil {
			return "", ErrBadInvite
		}
		return s.newPerson(pub)
	}

	id, err := s.newPerson(pub)
	if err != nil {
		return "", err
	}
	if err := s.Connect(id, inviterID, invite); err != nil {
		// A refused invite or a full network must not leave behind a person who
		// is connected to nobody.
		if errors.Is(err, ErrNoPerson) || errors.Is(err, ErrBadPersonID) {
			err = ErrBadInvite
		}
		return "", errors.Join(err, s.remove(id))
	}
	return id, nil
}

// OperatorInvite makes a single-use invite that belongs to nobody, for the first
// person on a server; everyone after them joins through someone already here.
// It is kept as an empty file named by the invite's hash, so the invite itself
// is never written down, and redeeming it is one os.Remove, which only one
// caller can win.
func (s *Store) OperatorInvite() (string, error) {
	dir := filepath.Join(s.dir, operatorInvites)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}
	token := NewToken()
	f, err := os.OpenFile(filepath.Join(dir, tokenHash(token)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return "", err
	}
	return token, f.Close()
}

func tokenHash(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func (s *Store) newPerson(pub ed25519.PublicKey) (string, error) {
	id := NewID()
	p, err := s.openPerson(id, true)
	if err != nil {
		return "", err
	}
	if _, err := p.write.Exec(`INSERT INTO person (id, created_at, public_key) VALUES (?, ?, ?)`,
		id, s.now().UnixMilli(), []byte(pub)); err != nil {
		return "", errors.Join(err, s.remove(id))
	}
	return id, nil
}

// remove closes a person's file and deletes it. It exists only to undo a
// registration that did not finish; nothing deletes a real person yet.
func (s *Store) remove(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if p, ok := s.open[id]; ok {
		p.read.Close()
		p.write.Close()
		delete(s.open, id)
	}
	var errs []error
	for _, suffix := range []string{"", "-wal", "-shm"} {
		if err := os.Remove(filepath.Join(s.dir, id+".db"+suffix)); err != nil && !errors.Is(err, os.ErrNotExist) {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

// Owns reports whether pub is this person's device key.
func (s *Store) Owns(personID string, pub ed25519.PublicKey) (bool, error) {
	p, err := s.person(personID)
	if err != nil {
		return false, err
	}
	var one int
	err = p.read.QueryRow(`SELECT 1 FROM person WHERE public_key = ?`, []byte(pub)).Scan(&one)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return err == nil, err
}

// SetProfile replaces a person's profile ciphertext. The server stores it and
// never reads it, exactly as it does a moment.
func (s *Store) SetProfile(personID string, blob []byte) error {
	p, err := s.person(personID)
	if err != nil {
		return err
	}
	_, err = p.write.Exec(`UPDATE person SET profile = ?`, blob)
	return err
}

// Invite mints a token that connects whoever redeems it to this person. A touch
// and a link carry the same token (CDI-1839, CDI-1840); how it travels is the
// app's business.
func (s *Store) Invite(personID string) (string, error) {
	p, err := s.person(personID)
	if err != nil {
		return "", err
	}
	token := NewToken()
	_, err = p.write.Exec(`INSERT INTO invites (token, created_at) VALUES (?, ?)`, token, s.now().UnixMilli())
	return token, err
}

// Connect redeems inviterID's invite and connects the two people both ways.
// Invitation is connection (ADR-0017): nobody approves anything afterwards. An
// invite works once; a refused connection leaves it unused.
//
// The two files are written one after the other, because SQLite cannot commit
// across files in WAL mode. If the second write fails the first is undone,
// invite included, and both are idempotent, so retrying the whole call is
// always safe.
func (s *Store) Connect(personID, inviterID, invite string) error {
	if personID == inviterID {
		return ErrOwnInvite
	}
	me, err := s.person(personID)
	if err != nil {
		return err
	}
	them, err := s.person(inviterID)
	if err != nil {
		return err
	}
	added, err := s.connect(them, personID, invite)
	if err != nil {
		return err
	}
	if _, err := s.connect(me, inviterID, ""); err != nil {
		if added {
			if _, undo := them.write.Exec(`DELETE FROM connections WHERE person_id = ?`, personID); undo != nil {
				return errors.Join(err, undo)
			}
			if _, undo := them.write.Exec(`UPDATE invites SET used_at = NULL, used_by = NULL WHERE token = ?`, invite); undo != nil {
				return errors.Join(err, undo)
			}
		}
		return err
	}
	return nil
}

// connect adds other to p's connections, checking the invite first when there
// is one and spending it in the same transaction. The cap is counted in that
// transaction too, on the only connection that writes this file, so two people
// cannot take the last place and a refusal leaves the invite unused.
func (s *Store) connect(p *personDB, other, invite string) (added bool, err error) {
	tx, err := p.write.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	// Already connected: connecting twice is a no-op, not an error.
	var one int
	connected := true
	if err := tx.QueryRow(`SELECT 1 FROM connections WHERE person_id = ?`, other).Scan(&one); errors.Is(err, sql.ErrNoRows) {
		connected = false
	} else if err != nil {
		return false, err
	}

	if invite != "" {
		var revoked, used sql.NullInt64
		var usedBy sql.NullString
		err := tx.QueryRow(`SELECT revoked_at, used_at, used_by FROM invites WHERE token = ?`, invite).
			Scan(&revoked, &used, &usedBy)
		switch {
		case errors.Is(err, sql.ErrNoRows) || (err == nil && revoked.Valid):
			return false, ErrBadInvite
		case err != nil:
			return false, err
		case used.Valid && connected && usedBy.String == other:
			return false, tx.Commit() // a retry of the connection this invite made
		case used.Valid:
			return false, ErrUsedInvite
		}
	}

	if connected {
		return false, tx.Commit()
	}

	// The cap is counted here and nowhere else. No count ever leaves the server.
	var n int
	if err := tx.QueryRow(`SELECT COUNT(*) FROM connections`).Scan(&n); err != nil {
		return false, err
	}
	if n >= NetworkCap {
		return false, ErrNetworkFull
	}
	now := s.now().UnixMilli()
	if _, err := tx.Exec(`INSERT INTO connections (person_id, connected_at) VALUES (?, ?)`, other, now); err != nil {
		return false, err
	}
	if invite != "" {
		if _, err := tx.Exec(`UPDATE invites SET used_at = ?, used_by = ? WHERE token = ?`, now, other, invite); err != nil {
			return false, err
		}
	}
	return true, tx.Commit()
}

// Post stores one moment's ciphertext in its author's file.
func (s *Store) Post(personID string, blob []byte) (Moment, error) {
	p, err := s.person(personID)
	if err != nil {
		return Moment{}, err
	}
	m := Moment{AuthorID: personID, CreatedAt: s.now().UnixMilli(), Blob: blob}
	res, err := p.write.Exec(`INSERT INTO moments (created_at, blob) VALUES (?, ?)`, m.CreatedAt, m.Blob)
	if err != nil {
		return Moment{}, err
	}
	raise(&p.newest, m.CreatedAt)
	m.ID, err = res.LastInsertId()
	return m, err
}

// Feed is the whole read path: the newest moments from this person and the
// people they are connected to, newest first, limited. No ranking and no cursor
// — a page loads and stops, per CDI-1833.
//
// It reads fan-in from each person's own file, which keeps export, delete and
// move as file operations (ADR-0011). Reading every file costs one SQLite
// transaction per connection, and at 150 that failed the latency threshold
// (CDI-1879, spike/RESULTS.md). So it visits people most recent poster first,
// reads only (id, created_at), which the index covers, and stops at the first
// person whose newest moment is older than the page it already holds. However
// many connections there are, a feed reads at most limit+1 files, then fetches
// the ciphertext of the moments that made the page.
func (s *Store) Feed(personID string, limit int) ([]Moment, error) {
	me, err := s.person(personID)
	if err != nil {
		return nil, err
	}
	rows, err := me.read.Query(`SELECT person_id FROM connections`)
	if err != nil {
		return nil, err
	}
	ids := []string{personID} // your own moments are in your feed too
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			rows.Close()
			return nil, err
		}
		ids = append(ids, id)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}

	type author struct {
		id     string
		db     *personDB
		newest int64
	}
	authors := make([]author, 0, len(ids))
	for _, id := range ids {
		p, err := s.person(id)
		if err != nil {
			return nil, err
		}
		newest, err := s.newest(p)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author{id: id, db: p, newest: newest})
	}
	sort.Slice(authors, func(i, j int) bool { return authors[i].newest > authors[j].newest })

	page := make([]Moment, 0, 2*limit)
	for _, a := range authors {
		if len(page) == limit && a.newest < page[limit-1].CreatedAt {
			break // nobody from here on has anything newer than the page
		}
		rows, err := a.db.read.Query(`SELECT id, created_at FROM moments ORDER BY created_at DESC LIMIT ?`, limit)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			m := Moment{AuthorID: a.id}
			if err := rows.Scan(&m.ID, &m.CreatedAt); err != nil {
				rows.Close()
				return nil, err
			}
			page = append(page, m)
		}
		rows.Close()
		if err := rows.Err(); err != nil {
			return nil, err
		}
		sort.Slice(page, func(i, j int) bool { return newer(page[i], page[j]) })
		page = page[:min(limit, len(page))]
	}

	if err := s.fillBlobs(page); err != nil {
		return nil, err
	}
	return page, nil
}

// newer orders a feed: newest first, with ties broken on author and id so a page
// is the same page every time it is read.
func newer(a, b Moment) bool {
	if a.CreatedAt != b.CreatedAt {
		return a.CreatedAt > b.CreatedAt
	}
	if a.AuthorID != b.AuthorID {
		return a.AuthorID > b.AuthorID
	}
	return a.ID > b.ID
}

// newest returns when a person last posted: from memory if it was read from
// their file within newestRecheck, and from the file otherwise.
func (s *Store) newest(p *personDB) (int64, error) {
	now := s.now().UnixMilli()
	if checked := p.checked.Load(); checked != 0 && now-checked < newestRecheck.Milliseconds() {
		return p.newest.Load(), nil
	}
	var n int64
	if err := p.read.QueryRow(`SELECT COALESCE(MAX(created_at), 0) FROM moments`).Scan(&n); err != nil {
		return 0, err
	}
	raise(&p.newest, n)
	p.checked.Store(now)
	return p.newest.Load(), nil
}

// raise moves v forward to n. Moments are never deleted, so newest never moves
// back; whatever deletes one later has to reset it.
func raise(v *atomic.Int64, n int64) {
	for old := v.Load(); n > old && !v.CompareAndSwap(old, n); old = v.Load() {
	}
}

// fillBlobs reads the ciphertext for a page of moments and each author's
// profile, from each author's file: per author on the page, never per
// connection, which is the cost the feed was measured on (spike/RESULTS.md).
func (s *Store) fillBlobs(page []Moment) error {
	byAuthor := map[string][]*Moment{}
	for i := range page {
		byAuthor[page[i].AuthorID] = append(byAuthor[page[i].AuthorID], &page[i])
	}
	for author, moments := range byAuthor {
		p, err := s.person(author)
		if err != nil {
			return err
		}
		var profile []byte
		if err := p.read.QueryRow(`SELECT profile FROM person`).Scan(&profile); err != nil {
			return err
		}
		for _, m := range moments {
			m.AuthorProfile = profile
		}
		byID := make(map[int64]*Moment, len(moments))
		args := make([]any, len(moments))
		for i, m := range moments {
			byID[m.ID], args[i] = m, m.ID
		}
		rows, err := p.read.Query(
			`SELECT id, blob FROM moments WHERE id IN (?`+strings.Repeat(", ?", len(moments)-1)+`)`, args...)
		if err != nil {
			return err
		}
		for rows.Next() {
			var id int64
			var blob []byte
			if err := rows.Scan(&id, &blob); err != nil {
				rows.Close()
				return err
			}
			byID[id].Blob = blob
		}
		rows.Close()
		if err := rows.Err(); err != nil {
			return err
		}
	}
	return nil
}

// person returns an existing person's handles, or ErrNoPerson.
func (s *Store) person(id string) (*personDB, error) {
	return s.openPerson(id, false)
}

func (s *Store) openPerson(id string, create bool) (*personDB, error) {
	if !ValidID(id) {
		return nil, ErrBadPersonID
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if p, ok := s.open[id]; ok {
		return p, nil
	}
	path := filepath.Join(s.dir, id+".db")
	if !create {
		if _, err := os.Stat(path); err != nil {
			return nil, ErrNoPerson
		}
	}
	// ponytail: handles are cached forever and never evicted. The fan-in spike
	// held seven to nine open files and about a quarter of a megabyte per person
	// with a file open (CDI-1879); add an LRU when one host holds thousands.
	write, err := openDB(path, 1)
	if err != nil {
		return nil, err
	}
	if err := migrate(write); err != nil {
		write.Close()
		return nil, err
	}
	read, err := openDB(path, 8)
	if err != nil {
		write.Close()
		return nil, err
	}
	p := &personDB{read: read, write: write}
	s.open[id] = p
	return p, nil
}

func openDB(path string, maxConns int) (*sql.DB, error) {
	dsn := path + "?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=synchronous(NORMAL)&_pragma=foreign_keys(1)"
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxConns)
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// migrate steps PRAGMA user_version to the end of the migrations slice. It runs
// against one person's file, which is the only shape that degrades sensibly
// when a host holds people nobody has opened in months.
func migrate(db *sql.DB) error {
	var v int
	if err := db.QueryRow(`PRAGMA user_version`).Scan(&v); err != nil {
		return err
	}
	if v > len(migrations) {
		return fmt.Errorf("person file is at schema %d, this server only knows %d", v, len(migrations))
	}
	for i := v; i < len(migrations); i++ {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		if _, err := tx.Exec(migrations[i]); err != nil {
			tx.Rollback()
			return fmt.Errorf("migration %d: %w", i+1, err)
		}
		// PRAGMA takes no parameters; i is an int, so this is not injectable.
		if _, err := tx.Exec(fmt.Sprintf(`PRAGMA user_version = %d`, i+1)); err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}
