#!/usr/bin/env bash
# PocketBase peaked at 719 MB under load with Go's default GC target. Is that
# tunable, or is it the framework? Re-run the same bench with GOMEMLIMIT set.
set -uo pipefail
here=$(cd "$(dirname "$0")" && pwd)
cd "$here/../candidates/go-pocketbase" || exit 1
out="$here/../results"
for limit in 256MiB 128MiB; do
  GOMEMLIMIT=$limit ./pbspike serve --http=127.0.0.1:8090 >/tmp/pb-mem.log 2>&1 &
  pid=$!
  for _ in $(seq 1 300); do curl -sf -o /dev/null "http://127.0.0.1:8090/feed?circle_id=fam000&limit=1" && break; sleep 0.1; done
  sleep 3
  idle=$(ps -o rss= -p $pid | tr -d ' ')
  "$here/measure.sh" "$pid" "/tmp/pb-mem-$limit.txt" & mpid=$!
  python3 "$here/harness.py" --port 8090 bench --concurrency 50 --requests 50000 > "/tmp/pb-bench-$limit.json" 2>/dev/null
  kill $mpid 2>/dev/null; kill $pid 2>/dev/null; wait $pid 2>/dev/null; sleep 1
  python3 - "$limit" "$idle" "/tmp/pb-mem-$limit.txt" "/tmp/pb-bench-$limit.json" "$out/go-pocketbase-memlimit.json" <<'PY'
import json, sys, os
limit, idle, rssf, benchf, res = sys.argv[1:]
s = [int(x) for x in open(rssf) if x.strip()]
b = json.load(open(benchf))
row = {"GOMEMLIMIT": limit, "idle_rss_mb": round(int(idle)/1024,1),
       "peak_rss_mb": round(max(s)/1024,1) if s else None,
       "feed_p95_ms_c50": b.get("feed_p95_ms_c50"), "rps_c50": b.get("rps_c50")}
all = json.load(open(res)) if os.path.exists(res) else []
all.append(row); json.dump(all, open(res,"w"), indent=2); print(json.dumps(row))
PY
done
