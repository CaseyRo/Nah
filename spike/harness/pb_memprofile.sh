#!/usr/bin/env bash
# Does PocketBase's peak scale with how many requests we send, and does it come
# back afterwards? Allocation churn that the collector lags behind looks very
# different from memory the process actually needs.
set -uo pipefail
here=$(cd "$(dirname "$0")" && pwd)
out="$here/../results"
run_one() { # <dir> <startcmd> <port> <requests> <label>
  cd "$1" || exit 1
  eval "exec $2" >/tmp/memprof.log 2>&1 & pid=$!
  for _ in $(seq 1 300); do curl -sf -o /dev/null "http://127.0.0.1:$3/feed?circle_id=fam000&limit=1" && break; sleep 0.1; done
  sleep 3
  idle=$(ps -o rss= -p $pid | tr -d ' ')
  "$here/measure.sh" "$pid" "/tmp/memprof-$5-$4.txt" & mpid=$!
  python3 "$here/harness.py" --port "$3" bench --concurrency 50 --requests "$4" >/dev/null 2>&1
  after=$(ps -o rss= -p $pid | tr -d ' ')
  sleep 15   # does it hand memory back when it goes quiet?
  settled=$(ps -o rss= -p $pid | tr -d ' ')
  kill $mpid 2>/dev/null; kill $pid 2>/dev/null; wait $pid 2>/dev/null; sleep 1
  python3 - "$5" "$4" "$idle" "$after" "$settled" "/tmp/memprof-$5-$4.txt" "$out/memory_profile.json" <<'PY'
import json, os, sys
label, reqs, idle, after, settled, rssf, res = sys.argv[1:]
s = [int(x) for x in open(rssf) if x.strip()]
row = {"candidate": label, "requests": int(reqs), "idle_mb": round(int(idle)/1024,1),
       "peak_mb": round(max(s)/1024,1) if s else None,
       "right_after_mb": round(int(after)/1024,1),
       "after_15s_idle_mb": round(int(settled)/1024,1)}
all = json.load(open(res)) if os.path.exists(res) else []
all.append(row); json.dump(all, open(res,"w"), indent=2); print(json.dumps(row))
PY
}
for n in 5000 20000 50000; do
  run_one "$here/../candidates/go-pocketbase" "./pbspike serve --http=127.0.0.1:8090" 8090 "$n" go-pocketbase
done
run_one "$here/../candidates/go-plain" "env DB=data.db ADDR=127.0.0.1:8091 ./plainspike" 8091 50000 go-plain
run_one "$here/../candidates/dart-shelf" "env DB=data.db ADDR=127.0.0.1:8093 ./dartspike" 8093 50000 dart-shelf
