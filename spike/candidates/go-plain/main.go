// Candidate B: plain Go, net/http, SQLite via the pure-Go driver.
// This is ADR-0010's escape hatch measured directly: what is underneath
// PocketBase once PocketBase is taken away.
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
)

var readDB, writeDB *sql.DB

type momentIn struct {
	CircleID  string `json:"circle_id"`
	AuthorID  string `json:"author_id"`
	CreatedAt int64  `json:"created_at"`
	Blob      string `json:"blob"`
}

func open(path string, maxConns int) *sql.DB {
	db, err := sql.Open("sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=synchronous(NORMAL)")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(maxConns)
	return db
}

func main() {
	path := os.Getenv("DB")
	if path == "" {
		path = "data.db"
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1:8091"
	}
	// One writer, because SQLite serialises writes anyway; many readers, which WAL allows.
	writeDB = open(path, 1)
	readDB = open(path, 8)
	if _, err := writeDB.Exec(`
		CREATE TABLE IF NOT EXISTS moments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			circle_id TEXT NOT NULL,
			author_id TEXT NOT NULL,
			created_at INTEGER NOT NULL,
			blob TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_moments_circle_created ON moments (circle_id, created_at DESC);
	`); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/moments", func(w http.ResponseWriter, r *http.Request) {
		var in momentIn
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if _, err := writeDB.Exec(
			`INSERT INTO moments (circle_id, author_id, created_at, blob) VALUES (?, ?, ?, ?)`,
			in.CircleID, in.AuthorID, in.CreatedAt, in.Blob); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"ok":true}`))
	})

	http.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		limit, _ := strconv.Atoi(q.Get("limit"))
		if limit <= 0 || limit > 200 {
			limit = 30
		}
		rows, err := readDB.Query(
			`SELECT id, circle_id, author_id, created_at, blob FROM moments
			 WHERE circle_id = ? ORDER BY created_at DESC LIMIT ?`, q.Get("circle_id"), limit)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()
		out := make([]map[string]any, 0, limit)
		for rows.Next() {
			var id, created int64
			var circle, author, blob string
			if err := rows.Scan(&id, &circle, &author, &created, &blob); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			out = append(out, map[string]any{"id": id, "circle_id": circle,
				"author_id": author, "created_at": created, "blob": blob})
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(out)
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}
