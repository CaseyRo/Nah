// Candidate A: PocketBase used as a Go framework (ADR-0010's incumbent).
// Two endpoints only: POST /moments and GET /feed.
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type momentIn struct {
	CircleID  string `json:"circle_id"`
	AuthorID  string `json:"author_id"`
	CreatedAt int64  `json:"created_at"`
	Blob      string `json:"blob"`
}

func main() {
	app := pocketbase.New()

	app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		if err := e.Next(); err != nil {
			return err
		}
		// PocketBase persists a log entry for every request by default. At
		// 50,000 requests that is 50,000 rows and a lot of garbage, so find out
		// whether it is what the memory is being spent on.
		if os.Getenv("PB_NOLOGS") == "1" {
			st := e.App.Settings()
			st.Logs.MaxDays = 0
			if err := e.App.Save(st); err != nil {
				return err
			}
		}
		if _, err := e.App.FindCollectionByNameOrId("moments"); err == nil {
			return nil
		}
		c := core.NewBaseCollection("moments")
		c.Fields.Add(
			&core.TextField{Name: "circle_id", Required: true, Max: 64},
			&core.TextField{Name: "author_id", Required: true, Max: 64},
			&core.NumberField{Name: "created_at", Required: true},
			&core.TextField{Name: "blob", Required: true, Max: 100000},
		)
		c.AddIndex("idx_moments_circle_created", false, "circle_id, created_at DESC", "")
		return e.App.Save(c)
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/moments", func(e *core.RequestEvent) error {
			var in momentIn
			if err := e.BindBody(&in); err != nil {
				return e.BadRequestError("bad body", err)
			}
			col, err := e.App.FindCollectionByNameOrId("moments")
			if err != nil {
				return err
			}
			r := core.NewRecord(col)
			r.Set("circle_id", in.CircleID)
			r.Set("author_id", in.AuthorID)
			r.Set("created_at", in.CreatedAt)
			r.Set("blob", in.Blob)
			if err := e.App.SaveNoValidate(r); err != nil {
				return err
			}
			return e.JSON(http.StatusCreated, map[string]bool{"ok": true})
		})

		se.Router.GET("/feed", func(e *core.RequestEvent) error {
			q := e.Request.URL.Query()
			limit, _ := strconv.Atoi(q.Get("limit"))
			if limit <= 0 || limit > 200 {
				limit = 30
			}
			// PB_RAW=1 skips PocketBase's record layer and queries the database
			// directly, to find out whether the record API is what allocates.
			if os.Getenv("PB_RAW") == "1" {
				rows := []struct {
					ID        string `db:"id" json:"id"`
					CircleID  string `db:"circle_id" json:"circle_id"`
					AuthorID  string `db:"author_id" json:"author_id"`
					CreatedAt int64  `db:"created_at" json:"created_at"`
					Blob      string `db:"blob" json:"blob"`
				}{}
				err := e.App.DB().
					NewQuery("SELECT id, circle_id, author_id, created_at, blob FROM moments WHERE circle_id={:cid} ORDER BY created_at DESC LIMIT {:lim}").
					Bind(dbx.Params{"cid": q.Get("circle_id"), "lim": limit}).
					All(&rows)
				if err != nil {
					return err
				}
				return e.JSON(http.StatusOK, rows)
			}
			records := []*core.Record{}
			err := e.App.RecordQuery("moments").
				AndWhere(dbx.HashExp{"circle_id": q.Get("circle_id")}).
				OrderBy("created_at DESC").
				Limit(int64(limit)).
				All(&records)
			if err != nil {
				return err
			}
			return e.JSON(http.StatusOK, records)
		})
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
