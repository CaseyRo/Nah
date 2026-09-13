#!/usr/bin/env bash
# Measure one candidate end to end, in the order the thresholds are written.
# The server is tracked by PID, never by matching command lines, because a
# pattern also matches this script and killing yourself makes for poor data.
# usage: run_candidate.sh <name> <port> <workdir> <start-cmd> <reset-cmd> <artifact> <db-glob>
set -uo pipefail
name=$1; port=$2; workdir=$3; start=$4; reset=$5; artifact=$6; dbglob=$7
here=$(cd "$(dirname "$0")" && pwd)
out="$here/../results"; mkdir -p "$out"
srvpid=""

start_server() { eval "exec $start" >>/tmp/$name.log 2>&1 & srvpid=$!; }
stop_server() {
  [ -n "$srvpid" ] || return 0
  pkill -P "$srvpid" 2>/dev/null; kill "$srvpid" 2>/dev/null; wait "$srvpid" 2>/dev/null
  sleep 1; srvpid=""
}
rss_kb() {
  local pids="$srvpid $(pgrep -P "$srvpid" 2>/dev/null | tr "\n" " ")"
  ps -o rss= -p $pids 2>/dev/null | awk "{s+=\$1} END {print s+0}"
}
wait_up() {
  for _ in $(seq 1 600); do
    curl -sf -o /dev/null "http://127.0.0.1:$port/feed?circle_id=fam000&limit=1" && return 0
    sleep 0.1
  done
  echo "$name: server never came up"; tail -5 /tmp/$name.log; exit 1
}

cd "$workdir" || exit 1
echo "== $name: reset =="; eval "$reset"; : > /tmp/$name.log

echo "== $name: start, then seed =="
start_server; wait_up
python3 "$here/harness.py" --port "$port" seed --concurrency 16 > "$out/$name.seed.json" 2>/dev/null
stop_server

echo "== $name: cold start against a seeded database =="
python3 "$here/harness.py" --port "$port" coldstart --cmd "$start" > "$out/$name.cold.json" 2>/dev/null
sleep 1

echo "== $name: idle memory, then two benches =="
start_server; wait_up; sleep 4
idle=$(rss_kb)
"$here/measure.sh" "$srvpid" "$out/$name.rss.txt" &
mpid=$!
python3 "$here/harness.py" --port "$port" bench --concurrency 1  --requests 1000  > "$out/$name.b1.json" 2>/dev/null
python3 "$here/harness.py" --port "$port" bench --concurrency 50 --requests 50000 > "$out/$name.b50.json" 2>/dev/null
kill $mpid 2>/dev/null
stop_server

dbbytes=$(du -sk $dbglob 2>/dev/null | awk "{s+=\$1} END {print s*1024}")
artbytes=$(du -sk "$artifact" 2>/dev/null | awk "{print \$1*1024}")
python3 - "$out/$name.json" "$name" "${idle:-0}" "${artbytes:-0}" "${dbbytes:-0}" "$out" <<PYEOF
import json, sys
res, name, idle, art, db, out = sys.argv[1:]
def load(p):
    try: return json.load(open(p))
    except Exception: return {}
try:
    samples = [int(x) for x in open(f"{out}/{name}.rss.txt") if x.strip()]
except Exception:
    samples = []
d = {"candidate": name,
     "idle_rss_mb": round(int(idle)/1024, 1),
     "peak_rss_mb": round(max(samples)/1024, 1) if samples else None,
     "artifact_mb": round(int(art)/1048576, 1),
     "db_mb": round(int(db or 0)/1048576, 1)}
for f in ("seed", "cold", "b1", "b50"):
    d.update(load(f"{out}/{name}.{f}.json"))
d.pop("concurrency", None)
json.dump(d, open(res, "w"), indent=2)
print(json.dumps(d, indent=2))
PYEOF
