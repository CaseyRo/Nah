---
title: "Spike results: which runtime, and how many databases"
permalink: /research/spike-results/
layout: page
date: 2026-09-13
---

Measured 13 September 2026 on an Apple M4, 10 cores, 16 GB, macOS. Every
candidate ran the same fixture, the same harness and the same two endpoints.
Thresholds were written first, in `THRESHOLDS.md`, and are not restated
charitably here.

Fixture: **50 circles, one year of history, 116,800 moments**, 40 family-shaped
circles at three moments a day and 10 busy circles at twenty. Moments are opaque
blobs of about 400 bytes, because the server never reads content.

## The four candidates

| | Dart + shelf | Plain Go | Go + PocketBase | Python + FastAPI |
|---|---|---|---|---|
| Idle memory | **15.9 MB** | 18.9 MB | 54.7 MB | 278.2 MB |
| Peak under load | **56.8 MB** | 74.2 MB | 719.0 MB | 404.0 MB |
| Artifact | **6.5 MB** | 15.0 MB | 32.9 MB | 38.3 MB |
| Cold start, seeded | **34 ms** | **34 ms** | 121 ms | 432 ms |
| Feed p95, 1 reader | 0.47 ms | **0.43 ms** | 0.47 ms | 0.45 ms |
| Feed p95, 50 readers | 10.2 ms | **8.5 ms** | 12.4 ms | 37.2 ms |
| Reads/sec, 50 readers | 5,781 | **13,951** | 8,403 | 1,977 |
| Inserts/sec, 16 writers | **11,215** | 10,217 | 6,482 | 6,890 |
| Post p95 | 5.9 ms | **4.9 ms** | 9.4 ms | 6.5 ms |
| Database on disk | 54.2 MB | 58.1 MB | 112.2 MB | 80.1 MB |

Against the thresholds:

| Threshold | Dart | Plain Go | PocketBase | Python |
|---|---|---|---|---|
| Feed p95 under 50 ms | pass | pass | pass | pass |
| Post p95 under 100 ms | pass | pass | pass | pass |
| Idle under 150 MB | pass | pass | pass | **fail** |
| Peak under 400 MB | pass | pass | **fail** | **fail** |
| Cold start under 2 s | pass | pass | pass | pass |
| Artifact under 100 MB | pass | pass | pass | pass |

## Python is out, and not on speed

278 MB idle across four workers, before anyone has posted anything. A single
worker would use less and be slower still: at four workers it already serves a
seventh of plain Go's throughput and its feed p95 is four times worse.

The packaging argument from the issue held up too. The artifact is a virtual
environment rather than a file, and it needs an interpreter and a process
manager wherever it runs. That was always the objection, and the memory number
settles it independently.

## PocketBase costs 7x the memory, and I could not find out why

This is the surprise, and it overturns the framework choice in ADR-0010.

Peak memory under the read benchmark, same workload, same machine:

| | Peak | After 15 s idle |
|---|---|---|
| Plain Go | 73.5 MB | 48.5 MB |
| Dart | 56.4 MB | 47.5 MB |
| PocketBase | 705 MB | 20 MB |

Three hypotheses, all tested, all negative:

1. **Go's garbage collector being lazy.** `GOMEMLIMIT=256MiB` gave 732 MB.
   `GOMEMLIMIT=128MiB` gave 742 MB. The limit made no difference at all, which
   means this is not heap the collector can be told to reclaim.
2. **PocketBase's record layer.** Querying the database directly and skipping
   records entirely still peaked at 587 MB.
3. **Request logging**, which PocketBase persists for every call by default.
   Disabling it, with raw queries as well, still peaked at 537 MB.

Each fix helped a little, 719 to 587 to 537, so there are several contributors
and no single cause. None of them gets within 7x of plain Go.

The memory does come back when the server goes quiet, so this is churn rather
than a leak. That is a weaker defence than it sounds: the peak is what a box has
to survive, and an out-of-memory killer does not care that the memory was about
to be returned. At 5,000 requests, which is a quiet evening for one busy circle,
it still peaked at 555 MB.

**What dropping PocketBase costs.** It was bringing sign-in, file storage, an
admin interface, backups, live updates and access rules. Of those, only auth and
file storage are on the critical path, and both are bounded work on a server
whose entire job is to store blobs and return the newest thirty. Live updates and
access rules were never needed: there is no realtime requirement, and membership
is the only rule there is.

## One database per circle

Sixteen writers, each writing to a different circle:

| | One shared database | One per circle |
|---|---|---|
| Writes/sec | 16,963 | **20,202** |
| Write p50 | **0.02 ms** | 0.62 ms |
| Write p95 | **0.04 ms** | 1.47 ms |
| Write worst case | 172 ms | **3.1 ms** |
| Export one circle | 8 ms | **~0 ms**, a file copy |
| Delete one circle | 2 ms + 170 ms vacuum | **~0 ms**, `rm` |
| 200 circles open | n/a | fine, 200 connections |

The shared database is quicker at the median and far worse at the tail, because
every writer queues behind one lock and a checkpoint stalls everyone. A 172 ms
stall is invisible at this scale but it is the shape of the thing that gets worse
with more circles, not better.

The decisive numbers are the last three rows. Export, delete and move stop being
queries and become filesystem operations, which is exactly what the product
promises about a circle being able to leave. Vacuuming a shared database after
deleting a circle blocks every other circle on that server.

**Decision: one SQLite database per circle.** The cost is one backup stream per
file, which Litestream handles from a single configuration.

## Digests are a non-event

5,000 circles, 2,046 of them with something to say, all firing in the same
minute:

| | Wall time | Per-circle p95 |
|---|---|---|
| All at once | **0.70 s** | 31.2 ms |
| Spread over 10 s | 10.0 s | 1.8 ms |

Against a 60 second threshold. Digests can stay a scheduled loop inside the same
binary, with no queue and no separate worker.

The honest caveat: the push relay here is a local mock, so this measures the
database side only. The real constraint is Apple and Google, and that lives in
the relay, which is a separate service by necessity anyway. Spreading is still
worth doing, because it costs nothing and it is the relay that will thank us.

## What is not measured yet

- **Restore from backup** needs Litestream installed. Threshold 7 is open.
- **Media path and growth** (CDI-1824) is not run.
- **ARM and small-box behaviour** (CDI-1825) is not run. Go and Dart both
  cross-compile cleanly, so the expectation is mild, but it is an expectation
  rather than a measurement.
- **Encryption cost on a cheap phone** (CDI-1871) needs a physical device.
  Thresholds 10 and 11 are open, and they are the two that decide whether
  switching circles feels instant.

None of these change the runtime choice, which is why ADR-0011 does not wait for
them.

## Methodology, including what would make these numbers wrong

- The load generator is Python, so every latency includes a few milliseconds of
  client overhead. It is identical for all four candidates, so the comparison
  holds even though the absolute numbers flatter nobody.
- Python ran four uvicorn workers, which is the realistic deployment. One worker
  would use roughly a quarter of the memory and be slower again.
- Dart ran in a single isolate, which is why its throughput is half of Go's while
  its latency is fine. A production Dart server would need an isolate pool, and
  that work is not in its numbers.
- Everything ran on one machine over loopback, on an M4. A cheap ARM box will be
  slower across the board, and that matters for absolute thresholds rather than
  for the ranking.
- Reported cold starts are from a seeded database, not an empty one.
