// Package nah is the circle server: one SQLite file per circle, holding
// membership, invites and the ciphertext of moments.
//
// The server never reads a moment. It stores blobs, orders them by time and
// hands back the newest few. Every product feature that would need plaintext
// (search, ranking, counts, thumbnails) is refused by design, which is what
// makes ADR-0012 possible at all.
package nah

import (
	"crypto/ed25519"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

// MemberCap is Dunbar's number, applied per circle (CDI-1830). It is a
// server-side rule, not a client check, and a full circle is told so in words.
const MemberCap = 150

var (
	ErrNoCircle    = errors.New("no such circle")
	ErrBadInvite   = errors.New("invite is not valid")
	ErrCircleFull  = errors.New("circle is full")
	ErrNotAMember  = errors.New("not a member of this circle")
	ErrBadCircleID = errors.New("malformed circle id")
)

// migrations are applied in order, stepping PRAGMA user_version, when a
// circle's file is opened. Per ADR-0014 this is hand-rolled: with one database
// per circle there is no single database for a migration tool to own.
//
// Append only. Never edit a migration that has shipped.
var migrations = []string{`
CREATE TABLE circle (
	id          TEXT PRIMARY KEY,
	created_at  INTEGER NOT NULL,
	secret_key  BLOB NOT NULL      -- the circle's own identity, per ADR-0010:
);                                 -- a key, never a hostname. M2's directory checks it.

CREATE TABLE members (
	public_key  BLOB PRIMARY KEY,  -- ed25519, 32 bytes, generated on the device (ADR-0015)
	joined_at   INTEGER NOT NULL
);

CREATE TABLE invites (
	token       TEXT PRIMARY KEY,
	created_at  INTEGER NOT NULL,
	revoked_at  INTEGER
);

CREATE TABLE moments (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	circle_id   TEXT NOT NULL,
	author_id   TEXT NOT NULL,
	created_at  INTEGER NOT NULL,
	blob        BLOB NOT NULL      -- ciphertext, never read by the server (ADR-0012)
);

CREATE INDEX idx_moments_circle_created ON moments (circle_id, created_at DESC);
`}

// Store owns the circle files under one directory.
type Store struct {
	dir string

	mu   sync.Mutex
	open map[string]*circleDB
}

// circleDB keeps the split the spike measured: one writer, because SQLite
// serialises writes anyway, and several readers, which WAL allows.
type circleDB struct {
	read  *sql.DB
	write *sql.DB
}

func NewStore(dir string) (*Store, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, err
	}
	return &Store{dir: dir, open: map[string]*circleDB{}}, nil
}

func (s *Store) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, c := range s.open {
		c.read.Close()
		c.write.Close()
		delete(s.open, id)
	}
}

// Moment is one row of a circle's feed. Blob is ciphertext.
type Moment struct {
	ID        int64  `json:"id"`
	CircleID  string `json:"circle_id"`
	AuthorID  string `json:"author_id"`
	CreatedAt int64  `json:"created_at"`
	Blob      []byte `json:"blob"`
}

// Create makes a new circle: its own file, its own identity key, and one
// invite to get the first member in.
func (s *Store) Create() (id, invite string, err error) {
	id = NewID()
	_, secret, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}
	c, err := s.openCircle(id, true)
	if err != nil {
		return "", "", err
	}
	invite = NewToken()
	now := time.Now().UnixMilli()
	tx, err := c.write.Begin()
	if err != nil {
		return "", "", err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(`INSERT INTO circle (id, created_at, secret_key) VALUES (?, ?, ?)`,
		id, now, []byte(secret)); err != nil {
		return "", "", err
	}
	if _, err := tx.Exec(`INSERT INTO invites (token, created_at) VALUES (?, ?)`, invite, now); err != nil {
		return "", "", err
	}
	return id, invite, tx.Commit()
}

// Join redeems an invite and registers a device's public key as a membership.
// Invitation is joining, per ADR-0009: there is nothing for anyone to approve
// afterwards.
func (s *Store) Join(circleID, invite string, pub ed25519.PublicKey) error {
	c, err := s.circle(circleID)
	if err != nil {
		return err
	}
	tx, err := c.write.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var revoked sql.NullInt64
	err = tx.QueryRow(`SELECT revoked_at FROM invites WHERE token = ?`, invite).Scan(&revoked)
	if errors.Is(err, sql.ErrNoRows) || (err == nil && revoked.Valid) {
		return ErrBadInvite
	}
	if err != nil {
		return err
	}

	// Already a member: joining twice on the same device is a no-op, not an error.
	var exists int
	if err := tx.QueryRow(`SELECT 1 FROM members WHERE public_key = ?`, []byte(pub)).Scan(&exists); err == nil {
		return tx.Commit()
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// The cap is counted here and nowhere else. No count ever leaves the server.
	var n int
	if err := tx.QueryRow(`SELECT COUNT(*) FROM members`).Scan(&n); err != nil {
		return err
	}
	if n >= MemberCap {
		return ErrCircleFull
	}
	if _, err := tx.Exec(`INSERT INTO members (public_key, joined_at) VALUES (?, ?)`,
		[]byte(pub), time.Now().UnixMilli()); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *Store) IsMember(circleID string, pub ed25519.PublicKey) (bool, error) {
	c, err := s.circle(circleID)
	if err != nil {
		return false, err
	}
	var one int
	err = c.read.QueryRow(`SELECT 1 FROM members WHERE public_key = ?`, []byte(pub)).Scan(&one)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return err == nil, err
}

// Post stores one moment's ciphertext.
func (s *Store) Post(circleID string, pub ed25519.PublicKey, blob []byte) (Moment, error) {
	c, err := s.circle(circleID)
	if err != nil {
		return Moment{}, err
	}
	m := Moment{
		CircleID:  circleID,
		AuthorID:  hex.EncodeToString(pub),
		CreatedAt: time.Now().UnixMilli(),
		Blob:      blob,
	}
	res, err := c.write.Exec(
		`INSERT INTO moments (circle_id, author_id, created_at, blob) VALUES (?, ?, ?, ?)`,
		m.CircleID, m.AuthorID, m.CreatedAt, m.Blob)
	if err != nil {
		return Moment{}, err
	}
	m.ID, err = res.LastInsertId()
	return m, err
}

// Feed is the whole read path: this circle's moments, newest first, limited.
// No fan-out, no timeline cache, no ranking, and no cursor — a page loads and
// stops, per CDI-1833.
func (s *Store) Feed(circleID string, limit int) ([]Moment, error) {
	c, err := s.circle(circleID)
	if err != nil {
		return nil, err
	}
	rows, err := c.read.Query(
		`SELECT id, circle_id, author_id, created_at, blob FROM moments
		 WHERE circle_id = ? ORDER BY created_at DESC, id DESC LIMIT ?`, circleID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]Moment, 0, limit)
	for rows.Next() {
		var m Moment
		if err := rows.Scan(&m.ID, &m.CircleID, &m.AuthorID, &m.CreatedAt, &m.Blob); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// circle returns an existing circle's handles, or ErrNoCircle.
func (s *Store) circle(id string) (*circleDB, error) {
	return s.openCircle(id, false)
}

func (s *Store) openCircle(id string, create bool) (*circleDB, error) {
	if !ValidID(id) {
		return nil, ErrBadCircleID
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if c, ok := s.open[id]; ok {
		return c, nil
	}
	path := filepath.Join(s.dir, id+".db")
	if !create {
		if _, err := os.Stat(path); err != nil {
			return nil, ErrNoCircle
		}
	}
	// ponytail: handles are cached forever and never evicted. ADR-0011 flagged
	// file-handle pressure at thousands of circles per host; add an LRU when a
	// host actually holds that many.
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
	c := &circleDB{read: read, write: write}
	s.open[id] = c
	return c, nil
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
// against one circle's file, which is the only shape that degrades sensibly
// when a host holds circles nobody has opened in months.
func migrate(db *sql.DB) error {
	var v int
	if err := db.QueryRow(`PRAGMA user_version`).Scan(&v); err != nil {
		return err
	}
	if v > len(migrations) {
		return fmt.Errorf("circle is at schema %d, this server only knows %d", v, len(migrations))
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
