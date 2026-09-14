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
	"database/sql"
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

var (
	ErrNoPerson    = errors.New("no such person")
	ErrBadPersonID = errors.New("malformed person id")
	ErrBadInvite   = errors.New("invite is not valid")
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
// author's moments, not across a feed; (AuthorID, ID) is.
type Moment struct {
	ID        int64  `json:"id"`
	AuthorID  string `json:"author_id"`
	CreatedAt int64  `json:"created_at"`
	Blob      []byte `json:"blob"`
}

// Register makes a new person with a file of their own, holding the public key
// of the device that asked. The key is what signs in (ADR-0015); the id is what
// the file is called and what other people connect to, so it can outlive a phone.
func (s *Store) Register(pub ed25519.PublicKey) (string, error) {
	id := NewID()
	p, err := s.openPerson(id, true)
	if err != nil {
		return "", err
	}
	_, err = p.write.Exec(`INSERT INTO person (id, created_at, public_key) VALUES (?, ?, ?)`,
		id, s.now().UnixMilli(), []byte(pub))
	return id, err
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
// Invitation is connection (ADR-0017): nobody approves anything afterwards.
//
// The two files are written one after the other, because SQLite cannot commit
// across files in WAL mode. If the second write fails the first is undone, and
// both are idempotent, so retrying the whole call is always safe.
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
		}
		return err
	}
	return nil
}

// connect adds other to p's connections, checking the invite first when there
// is one. The cap is counted in the same transaction as the insert, on the only
// connection that writes this file, so two people cannot take the last place.
func (s *Store) connect(p *personDB, other, invite string) (added bool, err error) {
	tx, err := p.write.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	if invite != "" {
		var revoked sql.NullInt64
		err := tx.QueryRow(`SELECT revoked_at FROM invites WHERE token = ?`, invite).Scan(&revoked)
		if errors.Is(err, sql.ErrNoRows) || (err == nil && revoked.Valid) {
			return false, ErrBadInvite
		}
		if err != nil {
			return false, err
		}
	}

	// Already connected: connecting twice is a no-op, not an error.
	var one int
	if err := tx.QueryRow(`SELECT 1 FROM connections WHERE person_id = ?`, other).Scan(&one); err == nil {
		return false, tx.Commit()
	} else if !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	// The cap is counted here and nowhere else. No count ever leaves the server.
	var n int
	if err := tx.QueryRow(`SELECT COUNT(*) FROM connections`).Scan(&n); err != nil {
		return false, err
	}
	if n >= NetworkCap {
		return false, ErrNetworkFull
	}
	if _, err := tx.Exec(`INSERT INTO connections (person_id, connected_at) VALUES (?, ?)`,
		other, s.now().UnixMilli()); err != nil {
		return false, err
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

// fillBlobs reads the ciphertext for a page of moments, one query per author.
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
