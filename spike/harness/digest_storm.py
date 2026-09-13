#!/usr/bin/env python3
"""CDI-1823: what happens when thousands of digests fire in the same minute.

Every other request in this product arrives when a human opens an app. Digests
arrive at times people chose, and people choose round numbers. This is the only
synchronised load in the system, so it is the one worth measuring.

Each circle costs one "did anything happen today" query plus one call to a push
relay, which is mocked locally so the network hop is included but Apple is not.
"""
import json, random, sqlite3, statistics, threading, time, http.server, socketserver, urllib.request
from concurrent.futures import ThreadPoolExecutor

CIRCLES = 5000
ACTIVE_SHARE = 0.4      # the rest posted nothing today and must be suppressed
RELAY_PORT = 8099

class Relay(http.server.BaseHTTPRequestHandler):
    def do_POST(self):
        self.rfile.read(int(self.headers.get("content-length", 0)))
        self.send_response(200); self.end_headers(); self.wfile.write(b"ok")
    def log_message(self, *a): pass

def serve_relay():
    socketserver.TCPServer.allow_reuse_address = True
    srv = socketserver.ThreadingTCPServer(("127.0.0.1", RELAY_PORT), Relay)
    threading.Thread(target=srv.serve_forever, daemon=True).start()
    return srv

def build(path):
    random.seed(3)
    c = sqlite3.connect(path); c.execute("PRAGMA journal_mode=WAL")
    c.execute("""CREATE TABLE moments (id INTEGER PRIMARY KEY AUTOINCREMENT,
        circle_id TEXT, created_at INTEGER)""")
    c.execute("CREATE INDEX idx ON moments (circle_id, created_at DESC)")
    now = int(time.time())
    rows = []
    for i in range(CIRCLES):
        cid = f"c{i:05d}"
        # history, plus something today for the active share
        for d in range(1, 30):
            rows.append((cid, now - d * 86400))
        if random.random() < ACTIVE_SHARE:
            rows.append((cid, now - 3600))
    c.executemany("INSERT INTO moments (circle_id, created_at) VALUES (?,?)", rows)
    c.commit(); c.close()

def run(path, workers, spread_s):
    """spread_s > 0 releases the circles evenly across that many seconds, which
    is what jitter actually does. Sleeping inside a worker would just serialise
    the sleeps, which is a harness bug rather than a finding."""
    cutoff = int(time.time()) - 86400
    local = threading.local()
    sent = [0]
    lock = threading.Lock()
    def one(args):
        cid, release_at = args
        if release_at:
            delay = release_at - time.monotonic()
            if delay > 0:
                time.sleep(delay)
        c = getattr(local, "c", None)
        if c is None:
            c = sqlite3.connect(path, check_same_thread=False); local.c = c
        t0 = time.perf_counter()
        hit = c.execute("SELECT 1 FROM moments WHERE circle_id=? AND created_at>? LIMIT 1",
                        (cid, cutoff)).fetchone()
        if hit:
            body = json.dumps({"circle": cid, "text": "Your circle posted today"}).encode()
            req = urllib.request.Request(f"http://127.0.0.1:{RELAY_PORT}/push", body,
                                         {"Content-Type": "application/json"})
            urllib.request.urlopen(req, timeout=10).read()
            with lock: sent[0] += 1
        return (time.perf_counter() - t0) * 1000
    t0 = time.perf_counter()
    base = time.monotonic()
    work = [(f"c{i:05d}", base + (i / CIRCLES) * spread_s if spread_s else 0)
            for i in range(CIRCLES)]
    with ThreadPoolExecutor(max_workers=workers) as ex:
        lat = list(ex.map(one, work))
    dt = time.perf_counter() - t0
    s = sorted(lat)
    return {"circles": CIRCLES, "pushes_sent": sent[0], "workers": workers,
            "spread_seconds": spread_s, "wall_seconds": round(dt, 2),
            "per_circle_p50_ms": round(statistics.median(s), 2),
            "per_circle_p95_ms": round(s[int(len(s) * .95)], 2)}

if __name__ == "__main__":
    import os, tempfile
    d = tempfile.mkdtemp(prefix="nah-digest-")
    path = os.path.join(d, "digest.db")
    build(path)
    srv = serve_relay()
    out = {"all_at_once": run(path, 32, 0), "spread_over_10s": run(path, 32, 10)}
    print(json.dumps(out, indent=2))
    srv.shutdown()
