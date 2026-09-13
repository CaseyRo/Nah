#!/usr/bin/env python3
"""CDI-1822: one SQLite database per circle, or one for the whole server?

Cheap to test now, expensive to change later. Builds both layouts from the same
fixture and measures the operations that actually differ: concurrent writes
across circles, exporting one circle, deleting one circle, and how many file
handles a server ends up holding.
"""
import json, os, random, shutil, sqlite3, statistics, string, sys, tempfile, time
from concurrent.futures import ThreadPoolExecutor

CIRCLES = 50
MOMENTS_PER_CIRCLE = 2000
BLOB = 400
SCHEMA = """CREATE TABLE IF NOT EXISTS moments (
  id INTEGER PRIMARY KEY AUTOINCREMENT, circle_id TEXT NOT NULL,
  author_id TEXT NOT NULL, created_at INTEGER NOT NULL, blob TEXT NOT NULL);
CREATE INDEX IF NOT EXISTS idx_moments_circle_created ON moments (circle_id, created_at DESC);"""

def connect(path):
    c = sqlite3.connect(path, check_same_thread=False, timeout=10)
    c.execute("PRAGMA journal_mode=WAL")
    c.execute("PRAGMA synchronous=NORMAL")
    c.execute("PRAGMA busy_timeout=10000")
    return c

def blob():
    return "".join(random.choices(string.ascii_letters, k=BLOB))

def build(root):
    random.seed(1)
    now = int(time.time())
    shared = os.path.join(root, "shared.db")
    c = connect(shared); c.executescript(SCHEMA)
    rows = []
    for i in range(CIRCLES):
        cid = f"c{i:03d}"
        for j in range(MOMENTS_PER_CIRCLE):
            rows.append((cid, f"u{j%20:02d}", now - j * 600, blob()))
    c.executemany("INSERT INTO moments (circle_id,author_id,created_at,blob) VALUES (?,?,?,?)", rows)
    c.commit(); c.close()

    per = os.path.join(root, "per"); os.makedirs(per, exist_ok=True)
    for i in range(CIRCLES):
        cid = f"c{i:03d}"
        cc = connect(os.path.join(per, f"{cid}.db")); cc.executescript(SCHEMA)
        cc.executemany("INSERT INTO moments (circle_id,author_id,created_at,blob) VALUES (?,?,?,?)",
                       [r for r in rows if r[0] == cid])
        cc.commit(); cc.close()
    return shared, per

def writers(paths, per_writer=200):
    """Each thread writes to its own circle. Shared: one file. Per-circle: many."""
    def work(arg):
        cid, path = arg
        c = connect(path)
        lat = []
        for k in range(per_writer):
            t0 = time.perf_counter()
            c.execute("INSERT INTO moments (circle_id,author_id,created_at,blob) VALUES (?,?,?,?)",
                      (cid, "u00", int(time.time()), blob()))
            c.commit()
            lat.append((time.perf_counter() - t0) * 1000)
        c.close()
        return lat
    t0 = time.perf_counter()
    with ThreadPoolExecutor(max_workers=len(paths)) as ex:
        lat = [x for part in ex.map(work, paths) for x in part]
    dt = time.perf_counter() - t0
    s = sorted(lat)
    return {"writers": len(paths), "writes": len(lat), "seconds": round(dt, 2),
            "writes_per_sec": round(len(lat) / dt),
            "write_p50_ms": round(statistics.median(s), 2),
            "write_p95_ms": round(s[int(len(s) * .95)], 2),
            "write_max_ms": round(max(s), 2)}

def main():
    root = tempfile.mkdtemp(prefix="nah-topology-")
    try:
        t0 = time.perf_counter()
        shared, per = build(root)
        build_s = round(time.perf_counter() - t0, 1)

        n = 16
        shared_w = writers([(f"c{i:03d}", shared) for i in range(n)])
        per_w = writers([(f"c{i:03d}", os.path.join(per, f"c{i:03d}.db")) for i in range(n)])

        # export one circle
        t0 = time.perf_counter()
        c = connect(shared)
        rows = c.execute("SELECT * FROM moments WHERE circle_id=?", ("c007",)).fetchall()
        out = os.path.join(root, "export_shared.db")
        e = connect(out); e.executescript(SCHEMA)
        e.executemany("INSERT INTO moments (id,circle_id,author_id,created_at,blob) VALUES (?,?,?,?,?)", rows)
        e.commit(); e.close(); c.close()
        export_shared = round((time.perf_counter() - t0) * 1000)

        t0 = time.perf_counter()
        shutil.copy(os.path.join(per, "c007.db"), os.path.join(root, "export_per.db"))
        export_per = round((time.perf_counter() - t0) * 1000)

        # delete one circle
        c = connect(shared)
        t0 = time.perf_counter()
        c.execute("DELETE FROM moments WHERE circle_id=?", ("c008",)); c.commit()
        delete_shared = round((time.perf_counter() - t0) * 1000)
        t0 = time.perf_counter()
        c.execute("VACUUM"); c.close()
        vacuum_shared = round((time.perf_counter() - t0) * 1000)

        t0 = time.perf_counter()
        for suffix in ("", "-wal", "-shm"):
            p = os.path.join(per, "c008.db" + suffix)
            if os.path.exists(p): os.remove(p)
        delete_per = round((time.perf_counter() - t0) * 1000)

        # what a server holds open at 200 circles
        handles = []
        try:
            for i in range(200):
                p = os.path.join(per, f"h{i:03d}.db")
                cc = connect(p); cc.executescript(SCHEMA); handles.append(cc)
            open_ok = len(handles)
        except Exception as exc:
            open_ok = f"failed at {len(handles)}: {exc}"
        finally:
            for h in handles: h.close()

        print(json.dumps({
            "fixture": {"circles": CIRCLES, "moments_per_circle": MOMENTS_PER_CIRCLE, "build_seconds": build_s},
            "shared_db": {"concurrent_writes": shared_w,
                          "export_one_circle_ms": export_shared,
                          "delete_one_circle_ms": delete_shared,
                          "vacuum_after_delete_ms": vacuum_shared},
            "db_per_circle": {"concurrent_writes": per_w,
                              "export_one_circle_ms": export_per,
                              "delete_one_circle_ms": delete_per,
                              "open_connections_at_200_circles": open_ok},
        }, indent=2))
    finally:
        shutil.rmtree(root, ignore_errors=True)

if __name__ == "__main__":
    main()
