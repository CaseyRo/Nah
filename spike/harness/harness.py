#!/usr/bin/env python3
"""Seed and benchmark a candidate circle server.

Every candidate implements exactly two endpoints so the comparison is a
comparison:

    POST /moments   {"circle_id","author_id","created_at","blob"}  -> 201
    GET  /feed?circle_id=<id>&limit=30                             -> newest first

Blobs stand in for ciphertext: the server never reads them, so their content is
irrelevant and only their size matters.
"""
import argparse, json, os, random, statistics, string, sys, time
import http.client
from concurrent.futures import ThreadPoolExecutor

FAMILY_CIRCLES = 40      # 12-20 members, ~3 moments a day
BUSY_CIRCLES = 10        # 150 members, ~20 moments a day
DAYS = 365
BLOB_BYTES = 400         # a text moment's ciphertext, or a photo moment's metadata

def circles():
    out = []
    for i in range(FAMILY_CIRCLES):
        out.append({"id": f"fam{i:03d}", "members": random.randint(12, 20), "per_day": 3})
    for i in range(BUSY_CIRCLES):
        out.append({"id": f"busy{i:03d}", "members": 150, "per_day": 20})
    return out

def blob(n=BLOB_BYTES):
    return "".join(random.choices(string.ascii_letters + string.digits, k=n))

def conn(host, port):
    return http.client.HTTPConnection(host, port, timeout=30)

def post_batch(host, port, items):
    c = conn(host, port)
    lat = []
    for it in items:
        body = json.dumps(it)
        t0 = time.perf_counter()
        c.request("POST", "/moments", body, {"Content-Type": "application/json"})
        r = c.getresponse(); r.read()
        lat.append((time.perf_counter() - t0) * 1000)
        if r.status not in (200, 201):
            raise SystemExit(f"POST /moments -> {r.status}")
    c.close()
    return lat

def cmd_seed(a):
    random.seed(42)
    now = int(time.time())
    items = []
    for circ in circles():
        for day in range(DAYS):
            for _ in range(circ["per_day"]):
                ts = now - day * 86400 - random.randint(0, 86399)
                items.append({
                    "circle_id": circ["id"],
                    "author_id": f"u{random.randint(0, circ['members'] - 1):03d}",
                    "created_at": ts,
                    "blob": blob(),
                })
    random.shuffle(items)
    print(f"seeding {len(items):,} moments across {FAMILY_CIRCLES + BUSY_CIRCLES} circles", flush=True, file=sys.stderr)
    chunks = [items[i::a.concurrency] for i in range(a.concurrency)]
    t0 = time.perf_counter()
    with ThreadPoolExecutor(max_workers=a.concurrency) as ex:
        lat = [x for part in ex.map(lambda c: post_batch(a.host, a.port, c), chunks) for x in part]
    dt = time.perf_counter() - t0
    print(json.dumps({
        "seeded": len(items), "seconds": round(dt, 1),
        "inserts_per_sec": round(len(items) / dt),
        "post_p50_ms": round(statistics.median(lat), 2),
        "post_p95_ms": round(pct(lat, 95), 2),
    }, indent=2))

def pct(xs, p):
    xs = sorted(xs)
    return xs[min(len(xs) - 1, int(len(xs) * p / 100))]

def read_worker(host, port, ids, n):
    c = conn(host, port)
    lat = []
    for _ in range(n):
        cid = random.choice(ids)
        t0 = time.perf_counter()
        c.request("GET", f"/feed?circle_id={cid}&limit=30")
        r = c.getresponse(); body = r.read()
        lat.append((time.perf_counter() - t0) * 1000)
        if r.status != 200:
            raise SystemExit(f"GET /feed -> {r.status}: {body[:200]}")
    c.close()
    return lat

def cmd_bench(a):
    random.seed(7)
    ids = [c["id"] for c in circles()]
    # warm the page cache so we measure the server, not the first read
    read_worker(a.host, a.port, ids, 20)
    per = max(1, a.requests // a.concurrency)
    t0 = time.perf_counter()
    with ThreadPoolExecutor(max_workers=a.concurrency) as ex:
        lat = [x for part in ex.map(lambda _: read_worker(a.host, a.port, ids, per), range(a.concurrency)) for x in part]
    dt = time.perf_counter() - t0
    key = f"c{a.concurrency}"
    out = {
        f"requests_{key}": len(lat), "concurrency": a.concurrency,
        f"rps_{key}": round(len(lat) / dt),
        f"feed_p50_ms_{key}": round(statistics.median(lat), 2),
        f"feed_p95_ms_{key}": round(pct(lat, 95), 2),
        f"feed_p99_ms_{key}": round(pct(lat, 99), 2),
    }
    print(json.dumps(out, indent=2))
    if a.out:
        open(a.out, "w").write(json.dumps(out, indent=2))

def cmd_coldstart(a):
    """Time from launching the server to its first successful response."""
    import subprocess
    t0 = time.perf_counter()
    p = subprocess.Popen(a.cmd, shell=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    while True:
        if time.perf_counter() - t0 > 30:
            p.terminate(); raise SystemExit("no response within 30s")
        try:
            c = conn(a.host, a.port)
            c.request("GET", "/feed?circle_id=fam000&limit=1")
            if c.getresponse().status == 200:
                break
        except OSError:
            time.sleep(0.02)
    dt = time.perf_counter() - t0
    print(json.dumps({"cold_start_ms": round(dt * 1000)}, indent=2))
    p.terminate()

if __name__ == "__main__":
    ap = argparse.ArgumentParser()
    ap.add_argument("--host", default="127.0.0.1")
    ap.add_argument("--port", type=int, default=8090)
    sub = ap.add_subparsers(dest="cmd", required=True)
    s = sub.add_parser("seed"); s.add_argument("--concurrency", type=int, default=16); s.set_defaults(fn=cmd_seed)
    b = sub.add_parser("bench"); b.add_argument("--concurrency", type=int, default=50)
    b.add_argument("--requests", type=int, default=5000); b.add_argument("--out"); b.set_defaults(fn=cmd_bench)
    c = sub.add_parser("coldstart"); c.add_argument("--cmd", required=True); c.set_defaults(fn=cmd_coldstart)
    a = ap.parse_args(); a.fn(a)
