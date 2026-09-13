---
title: "ADR-0011 — Plain Go, and one database per circle"
permalink: /decisions/0011-plain-go-and-a-database-per-circle/
layout: page
status: Accepted
date: 2026-09-13
---

## Status

Accepted (2026-09-13). Amends [ADR-0010](0010-small-server-not-mastodon.html), which chose Go but named PocketBase as the framework. The server, the single binary and SQLite all stand. PocketBase does not.

## Context

ADR-0010 settled that Nah? needs its own small server rather than Mastodon, Matrix or anything else off the shelf, and proposed building it on PocketBase because it brings sign-in, file storage, backups and an admin interface for free.

The spike existed to test that by measurement rather than argument. Thresholds were written first. Four candidates implemented the same two endpoints, seeded with the same fixture of 50 circles and 116,800 moments, and were measured on the same machine. The full numbers are in [the spike results]({{ '/research/spike-results/' | relative_url }}).

## Decision

**Plain Go with SQLite, no framework, and one database file per circle.**

- Go, on the numbers and on the pre-registered tie-break. It passed every threshold, has the best throughput of the four, cross-compiles to a NAS, and ships as one file.
- **No PocketBase.** It failed the peak memory threshold by a wide margin, for reasons three separate experiments did not explain.
- **One SQLite database per circle**, rather than one for the whole server.
- Sign-in and file storage, which PocketBase was going to supply, get written. Neither is large on a server whose entire job is to store opaque blobs and return the newest thirty.

## Considered alternatives

- **Go on PocketBase, as ADR-0010 proposed.** Rejected on memory. Under the read benchmark it peaked at 705 MB where plain Go peaked at 73.5 MB, against a threshold of 400 MB. `GOMEMLIMIT` at 256 MiB and again at 128 MiB changed nothing, so it is not heap the collector can be told to release. Bypassing the record layer with raw queries still peaked at 587 MB. Turning off the per-request logging it persists by default still peaked at 537 MB. The memory does return when the server goes quiet, so it is churn rather than a leak, but the peak is what a small box has to survive, and at only 5,000 requests it still peaked at 555 MB. Giving it up costs sign-in, file storage, an admin screen, live updates and access rules. Only the first two are on the critical path, and membership is the only access rule this product has.
- **Dart with shelf.** The closest thing to a winner and genuinely surprising: the smallest binary at 6.5 MB, the lowest idle memory at 15.9 MB, the fastest writes, and a cold start tied with Go. One language across server and Flutter client is worth real money to a solo developer, and this measurement makes that argument stronger rather than weaker. Rejected only because it passed no better than Go did and the tie-break was written in advance: when two candidates both pass, take Go. Its throughput was half of Go's, which is a single-isolate artifact rather than a language limit, and closing it would mean building an isolate pool that is not in these numbers.
- **Python with FastAPI.** Out on memory before anything else: 278 MB idle across four workers, against a 150 MB threshold, and 404 MB under load. It also served a seventh of Go's throughput with four times the feed latency, and its deployable artifact is a virtual environment rather than a file.
- **One shared database for all circles.** Quicker at the median, far worse at the tail: a worst-case write of 172 ms against 3.1 ms, because every writer queues behind one lock. More importantly, exporting a circle becomes a query rather than a file copy, deleting one needs a vacuum that blocks every other circle on the server, and moving one becomes a migration. Per-circle files turn all three into filesystem operations, which is what the product promises about a circle being able to leave.

## Consequences

Positive:

- The server is one file, 15 MB, starting in 34 milliseconds with a year of data on disk.
- Nothing in it is unexplained. A framework that peaks at seven times the memory for reasons nobody can name is a bad thing to build a decade on.
- Export, delete and move a circle are file operations. So is a per-circle backup stream.
- Digests are settled as a side effect: 5,000 circles firing at once completed in 0.70 seconds, so they stay a scheduled loop in the same binary with no queue.
- The escape hatch ADR-0010 named has been taken deliberately and early, rather than in a panic later.

Negative / acknowledged cost:

- **Sign-in and file storage are now ours to write**, along with whatever admin tooling is genuinely needed, which is probably a command and not a screen.
- We own every security fix in code we wrote, with no upstream to inherit patches from.
- Per-circle databases mean many small files to back up and many open handles. Two hundred connections were fine; ten thousand circles on one host is a question this spike did not ask.
- Dart was not beaten on merit. If a second language ever becomes a burden rather than a choice, this decision deserves revisiting rather than defending.

## What would prove this wrong

- Writing sign-in and file storage takes materially longer than the week or two assumed here. That would mean PocketBase's memory cost was worth paying.
- A cheap ARM box behaves differently enough to change the ranking. Not yet measured.
- The client-side encryption thresholds fail on a real phone, which would move the bottleneck off the server entirely and make this whole comparison a footnote.
- Per-circle databases become an operational problem at a scale this spike did not reach, for example thousands of circles on one host exhausting file handles or making backup unmanageable.
