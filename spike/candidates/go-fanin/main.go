// Candidate for CDI-1879: plain Go, one SQLite file per person, and a feed read
// fan-in from the files of everyone the reader is connected to.
//
// ADR-0017 replaced circles with one network of up to 150 per person, so the
// feed stopped being the single query M0 measured. The store copies apps/server
// on purpose — the same pragmas, one writer and eight readers per file, handles
// opened lazily under one lock and never evicted — so the numbers describe the
// server that would actually run rather than a tuned toy.
//
//	POST /moments  {"author_id","created_at","blob"}  -> 201, into the author's file
//	POST /connect  {"a","b"}                          -> 201, into both files
//	GET  /feed?person=<id>&limit=30                   -> newest first, across connections
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof" // to find out why, when a number is bad; idle until asked
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE IF NOT EXISTS connections (peer TEXT PRIMARY KEY);
CREATE TABLE IF NOT EXISTS moments (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at INTEGER NOT NULL,
	blob TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_moments_created ON moments (created_at DESC);`

type person struct {
	read, write *sql.DB

	// newest is the created_at of this person's newest moment, and seen is when
	// it was last read from the file. Posts through this process raise it as they
	// happen; rereading after newestTTL catches posts made by another process on
	// the same volume, which is what a rolling deploy is.
	newest, seen atomic.Int64
}

func (p *person) newestMoment() (int64, error) {
	now := time.Now().UnixMilli()
	if now-p.seen.Load() < newestTTL {
		return p.newest.Load(), nil
	}
	var n int64
	if err := p.read.QueryRow(`SELECT COALESCE(MAX(created_at), 0) FROM moments`).Scan(&n); err != nil {
		return 0, err
	}
	raise(&p.newest, n)
	p.seen.Store(now)
	return p.newest.Load(), nil
}

// raise moves v forward to n. Moments are never deleted, so newest never moves back.
func raise(v *atomic.Int64, n int64) {
	for old := v.Load(); n > old && !v.CompareAndSwap(old, n); old = v.Load() {
	}
}

type moment struct {
	ID        int64  `json:"id"`
	AuthorID  string `json:"author_id"`
	CreatedAt int64  `json:"created_at"`
	Blob      string `json:"blob"`
}

var (
	dir          = env("DATA", "data")
	readConns, _ = strconv.Atoi(env("READ_CONNS", "8"))
	feedMode     = env("FEED", "blobs") // blobs, keys or newest
	keysFirst    = feedMode != "blobs"
	newestFirst  = feedMode == "newest"
	newestTTL, _ = strconv.ParseInt(env("NEWEST_TTL_MS", "30000"), 10, 64)

	mu     sync.Mutex
	people = map[string]*person{}
)

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func openDB(path string, maxConns int) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=synchronous(NORMAL)&_pragma=foreign_keys(1)")
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

// get opens a person's file on first use and keeps it, as apps/server's
// Store.openCircle does.
func get(id string) (*person, error) {
	if id == "" || strings.ContainsAny(id, `/\.`) {
		return nil, os.ErrInvalid
	}
	mu.Lock()
	defer mu.Unlock()
	if p, ok := people[id]; ok {
		return p, nil
	}
	path := filepath.Join(dir, id+".db")
	write, err := openDB(path, 1)
	if err != nil {
		return nil, err
	}
	if _, err := write.Exec(schema); err != nil {
		write.Close()
		return nil, err
	}
	read, err := openDB(path, readConns)
	if err != nil {
		write.Close()
		return nil, err
	}
	p := &person{read: read, write: write}
	people[id] = p
	return p, nil
}

// feed reads the newest `limit` moments from each connection's file, one file
// after another, and keeps the newest `limit` of all of them.
func feed(reader string, limit int) ([]moment, error) {
	me, err := get(reader)
	if err != nil {
		return nil, err
	}
	rows, err := me.read.Query(`SELECT peer FROM connections`)
	if err != nil {
		return nil, err
	}
	var peers []string
	for rows.Next() {
		var peer string
		if err := rows.Scan(&peer); err != nil {
			rows.Close()
			return nil, err
		}
		peers = append(peers, peer)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// FEED=keys reads only (id, created_at), which the index covers, so no table
	// row and no blob is touched for a moment that does not make the page.
	cols, n := "id, created_at, blob", 3
	if keysFirst {
		cols, n = "id, created_at", 2
	}

	// FEED=newest visits connections most recent poster first, and stops at the
	// first whose newest moment is older than the page already collected. A feed
	// then reads at most limit+1 files, however many connections there are.
	var newest map[string]int64
	size := limit * len(peers)
	if newestFirst {
		newest, size = make(map[string]int64, len(peers)), 2*limit
		for _, peer := range peers {
			p, err := get(peer)
			if err != nil {
				return nil, err
			}
			if newest[peer], err = p.newestMoment(); err != nil {
				return nil, err
			}
		}
		sort.Slice(peers, func(i, j int) bool { return newest[peers[i]] > newest[peers[j]] })
	}

	all := make([]moment, 0, size)
	for _, peer := range peers {
		if newestFirst && len(all) >= limit && newest[peer] < all[limit-1].CreatedAt {
			break
		}
		p, err := get(peer)
		if err != nil {
			return nil, err
		}
		rows, err := p.read.Query(`SELECT `+cols+` FROM moments ORDER BY created_at DESC LIMIT ?`, limit)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			m := moment{AuthorID: peer}
			if err := rows.Scan([]any{&m.ID, &m.CreatedAt, &m.Blob}[:n]...); err != nil {
				rows.Close()
				return nil, err
			}
			all = append(all, m)
		}
		rows.Close()
		if err := rows.Err(); err != nil {
			return nil, err
		}
		if newestFirst {
			newestOrder(all)
			all = all[:min(limit, len(all))]
		}
	}
	newestOrder(all)
	page := all[:min(limit, len(all))]
	if keysFirst {
		return page, fillBlobs(page)
	}
	return page, nil
}

func newestOrder(ms []moment) {
	sort.Slice(ms, func(i, j int) bool {
		if ms[i].CreatedAt != ms[j].CreatedAt {
			return ms[i].CreatedAt > ms[j].CreatedAt
		}
		return ms[i].AuthorID > ms[j].AuthorID
	})
}

// fillBlobs fetches the blobs for one page of moments, one query per author.
func fillBlobs(page []moment) error {
	byAuthor := map[string]map[int64]*moment{}
	for i := range page {
		m := &page[i]
		if byAuthor[m.AuthorID] == nil {
			byAuthor[m.AuthorID] = map[int64]*moment{}
		}
		byAuthor[m.AuthorID][m.ID] = m
	}
	for author, byID := range byAuthor {
		p, err := get(author)
		if err != nil {
			return err
		}
		args := make([]any, 0, len(byID))
		for id := range byID {
			args = append(args, id)
		}
		rows, err := p.read.Query(`SELECT id, blob FROM moments WHERE id IN (?`+strings.Repeat(",?", len(args)-1)+`)`, args...)
		if err != nil {
			return err
		}
		for rows.Next() {
			var id int64
			var blob string
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

func main() {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		log.Fatal(err)
	}
	if os.Getenv("PROFILE") != "" {
		runtime.SetMutexProfileFraction(10)
	}

	http.HandleFunc("POST /moments", func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			AuthorID  string `json:"author_id"`
			CreatedAt int64  `json:"created_at"`
			Blob      string `json:"blob"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		p, err := get(in.AuthorID)
		if err == nil {
			_, err = p.write.Exec(`INSERT INTO moments (created_at, blob) VALUES (?, ?)`, in.CreatedAt, in.Blob)
		}
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		raise(&p.newest, in.CreatedAt)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"ok":true}`))
	})

	// Connection is mutual, so it is written into both files. Not atomic across
	// the two, which is a real server's problem and not this measurement's.
	http.HandleFunc("POST /connect", func(w http.ResponseWriter, r *http.Request) {
		var in struct{ A, B string }
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		for _, pair := range [][2]string{{in.A, in.B}, {in.B, in.A}} {
			p, err := get(pair[0])
			if err == nil {
				_, err = p.write.Exec(`INSERT OR IGNORE INTO connections (peer) VALUES (?)`, pair[1])
			}
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"ok":true}`))
	})

	http.HandleFunc("GET /feed", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		limit, _ := strconv.Atoi(q.Get("limit"))
		if limit <= 0 || limit > 200 {
			limit = 30
		}
		out, err := feed(q.Get("person"), limit)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(out)
	})

	log.Fatal(http.ListenAndServe(env("ADDR", "127.0.0.1:8092"), nil))
}
