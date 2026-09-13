#!/usr/bin/env bash
# Append total resident memory in KB for a server process and its children,
# so multi-worker servers are counted honestly. Runs until the parent is gone.
pid=$1; out=$2
: > "$out"
while kill -0 "$pid" 2>/dev/null; do
  pids="$pid $(pgrep -P "$pid" 2>/dev/null | tr '\n' ' ')"
  ps -o rss= -p $pids 2>/dev/null | awk '{s+=$1} END {if (s>0) print s}' >> "$out"
  sleep 0.2
done
