# Spike

## Purpose [coverage: high -- 5 sources]

Sources span 2026-09-13 to 2026-09-14.

The spike is throwaway measurement code, and none of it ships. It exists so that architecture questions are settled by numbers written down before anything was measured: which runtime the [server](server.md) is written in, whether storage is one database or many, and, once [ADR-0017](../../decisions/0017-one-network-of-a-hundred-and-fifty.md) arrived, what a feed read across 150 people costs. Its outcomes are [ADR-0011](../../decisions/0011-plain-go-and-a-database-per-circle.md) and that record's 2026-09-14 amendment.

## Architecture [coverage: medium -- 3 sources]

- [`THRESHOLDS.md`](../../../spike/THRESHOLDS.md): the pass and fail numbers, written before any measurement, and the tie-break agreed in advance.
- [`harness/harness.py`](../../../spike/harness/harness.py): seeds a fixture and benchmarks it. `--degree N` swaps the circle fixture for 1,000 people each connected to N others.
- [`harness/run_candidate.sh`](../../../spike/harness/run_candidate.sh): measures one candidate end to end, in the order the thresholds are written. Its first lines give the usage.
- [`harness/topology.py`](../../../spike/harness/topology.py) and [`harness/digest_storm.py`](../../../spike/harness/digest_storm.py): the storage-shape and digest measurements.
- [`candidates/`](../../../spike/candidates/): one minimal server per option. `go-fanin` copies the server's store on purpose and switches feed shape with the `FEED` environment variable.
- [`results/`](../../../spike/results/): every run's numbers. [`RESULTS.md`](../../../spike/RESULTS.md) is the write-up, and a copy is published on the site.

## Talks To [coverage: medium -- 2 sources]

- **Nothing outside one machine.** Every candidate runs over loopback against the same harness.
- **The server, by imitation.** `go-fanin` mirrors the server's pragmas, connection pools and file handling, so its numbers describe the code that actually runs.
- **The decision records**, which carry the conclusions: [ADR-0011](../../decisions/0011-plain-go-and-a-database-per-circle.md), and through it [ADR-0014](../../decisions/0014-the-go-setup.md), which took its dependency caution from the PocketBase result.

## Key Decisions [coverage: high -- 6 sources]

Newest first. [`RESULTS.md`](../../../spike/RESULTS.md) is the canonical home for every number.

- **2026-09-14:** reading every connection's file per feed fails (196.6 ms at the 95th percentile with 50 readers, against a 50 ms threshold); reading newest poster first passes (23.6 ms, 254 MB peak). See the amendment to [ADR-0011](../../decisions/0011-plain-go-and-a-database-per-circle.md).
- **2026-09-13:** plain Go. PocketBase peaked at 705 MB against 73.5 MB for reasons three experiments did not explain, Dart passed everything and lost only the tie-break, and Python failed idle memory ([ADR-0011](../../decisions/0011-plain-go-and-a-database-per-circle.md)).
- **2026-09-13:** a database per unit rather than one shared file: a worst-case write of 3.1 ms against 172 ms, and export and delete become file operations.
- **2026-09-13:** digests stay a loop inside the server binary, because 5,000 firing at once finished in 0.70 s.
- **2026-09-13:** thresholds, the Go tie-break and a two-week time box were written before anything ran ([THRESHOLDS.md](../../../spike/THRESHOLDS.md)), against a replaced plan to build on PocketBase ([ADR-0010](../../decisions/0010-small-server-not-mastodon.md)).

## Running It [coverage: medium -- 2 sources]

From `spike`, rerun a known candidate first to check the machine still gives the old numbers:

```bash
(cd candidates/go-plain && go build -o plainspike .)
harness/run_candidate.sh go-plain-recheck 8091 candidates/go-plain "./plainspike" \
  "rm -f data.db data.db-wal data.db-shm" plainspike "data.db*"
```

The feed read across people, at 150 connections each:

```bash
(cd candidates/go-fanin && go build -o faninspike .)
HARNESS_ARGS="--degree 150" PROBE="/feed?person=p0000&limit=30" \
  harness/run_candidate.sh go-fanin-150-newest 8092 candidates/go-fanin \
  "env FEED=newest ./faninspike" "rm -rf data" faninspike data
```

## Gotchas [coverage: medium -- 3 sources]

- **Numbers only compare on the same machine.** They come from one Apple M4 over loopback, and the Python load generator's overhead is inside every latency.
- **Give every run a new name.** A run writes `results/<name>.*` and overwrites whatever had that name.
- **Fixtures are ignored by git but large.** The people fixture left 2.5 GB of write-ahead logs behind, because the harness kills servers rather than stopping them.
- **The first commit of new results can abort.** The result files have no trailing newline, the end-of-file hook adds one, and the commit stops; stage them again and commit ([delivery](delivery.md)).
- **Still unmeasured:** restoring from backup, the media path, arm64 hardware, and encryption on a cheap phone.

## Sources

- [spike/README.md](../../../spike/README.md)
- [spike/THRESHOLDS.md](../../../spike/THRESHOLDS.md)
- [spike/RESULTS.md](../../../spike/RESULTS.md)
- [.pre-commit-config.yaml](../../../.pre-commit-config.yaml)
- [docs/decisions/0010-small-server-not-mastodon.md](../../decisions/0010-small-server-not-mastodon.md)
- [docs/decisions/0011-plain-go-and-a-database-per-circle.md](../../decisions/0011-plain-go-and-a-database-per-circle.md)
- [docs/decisions/0014-the-go-setup.md](../../decisions/0014-the-go-setup.md)
- [docs/decisions/0017-one-network-of-a-hundred-and-fifty.md](../../decisions/0017-one-network-of-a-hundred-and-fifty.md)
