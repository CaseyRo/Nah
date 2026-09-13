"""Candidate C: Python, FastAPI on uvicorn, SQLite.

Sync endpoints on purpose: FastAPI runs them in a thread pool, which is the
idiomatic way to use a blocking database driver here.
"""
import os, sqlite3, threading
from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse

DB = os.environ.get("DB", "data.db")
app = FastAPI()
_write_lock = threading.Lock()
_local = threading.local()


def conn():
    c = getattr(_local, "c", None)
    if c is None:
        c = sqlite3.connect(DB, check_same_thread=False, timeout=5.0)
        c.execute("PRAGMA journal_mode=WAL")
        c.execute("PRAGMA synchronous=NORMAL")
        c.execute("PRAGMA busy_timeout=5000")
        _local.c = c
    return c


with sqlite3.connect(DB) as c:
    c.execute("PRAGMA journal_mode=WAL")
    c.execute("""CREATE TABLE IF NOT EXISTS moments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        circle_id TEXT NOT NULL, author_id TEXT NOT NULL,
        created_at INTEGER NOT NULL, blob TEXT NOT NULL)""")
    c.execute("CREATE INDEX IF NOT EXISTS idx_moments_circle_created ON moments (circle_id, created_at DESC)")


@app.post("/moments")
async def post_moment(request: Request):
    d = await request.json()
    def write():
        with _write_lock:
            c = conn()
            c.execute("INSERT INTO moments (circle_id, author_id, created_at, blob) VALUES (?,?,?,?)",
                      (d["circle_id"], d["author_id"], d["created_at"], d["blob"]))
            c.commit()
    import anyio
    await anyio.to_thread.run_sync(write)
    return JSONResponse({"ok": True}, status_code=201)


@app.get("/feed")
def feed(circle_id: str, limit: int = 30):
    limit = 30 if limit <= 0 or limit > 200 else limit
    rows = conn().execute(
        "SELECT id, circle_id, author_id, created_at, blob FROM moments "
        "WHERE circle_id = ? ORDER BY created_at DESC LIMIT ?", (circle_id, limit)).fetchall()
    return [{"id": r[0], "circle_id": r[1], "author_id": r[2], "created_at": r[3], "blob": r[4]} for r in rows]
