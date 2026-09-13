---
title: "Measuring Instead of Arguing (or: the framework that wanted 700 megabytes)"
date: 2026-09-13 12:00:00 +0200
categories: [Architecture, Decisions]
tags: [spike, go, sqlite, benchmarks, decision-log]
---

Last post ended with a promise: the next one has a screenshot in it, or it doesn't count.

This one has no screenshot. What it has instead is a set of numbers that changed what Nah? is built on, which I think is the acceptable version of breaking that promise. There is running code now. It just isn't code you can look at yet.

## Thresholds before measurements

The thing I wanted to avoid was a benchmark that turns into a hobby. So the first commit of the spike wasn't code, it was a list of numbers with no results next to them: how much memory the server may use when idle, how much under load, how fast a feed query may be, how big the deployable thing may be, how long a cold start may take.

Written first, so the tests could only confirm or refute them rather than produce something interesting that I then reasoned my way around. There was also a tie-break agreed in advance, and a two-week time box, because the documented failure mode of this project is research quietly replacing code.

Then four servers, each implementing exactly two endpoints: post a moment, and list a circle's newest thirty. Same fixture for all of them, fifty circles and a year of history, 116,800 moments. Same load generator.

## The results

| | Dart | Plain Go | Go + PocketBase | Python |
|---|---|---|---|---|
| Idle memory | 15.9 MB | 18.9 MB | 54.7 MB | 278.2 MB |
| Peak under load | 56.8 MB | 74.2 MB | **719 MB** | 404 MB |
| Binary or bundle | 6.5 MB | 15 MB | 32.9 MB | 38.3 MB |
| Cold start | 34 ms | 34 ms | 121 ms | 432 ms |
| Feed p95, 50 readers | 10.2 ms | 8.5 ms | 12.4 ms | 37.2 ms |
| Reads per second | 5,781 | 13,951 | 8,403 | 1,977 |

Python was out on the first threshold it touched. 278 MB sitting idle across four workers, before anybody has posted anything, against a limit of 150. It was also seven times slower than plain Go at serving a feed. The packaging objection I'd written down beforehand held too: the artifact is a virtual environment rather than a file, and it needs an interpreter and a process manager wherever it lands.

And then PocketBase, which was the thing I had already decided on.

## Seven hundred megabytes, and I still don't know why

PocketBase peaked at 705 MB serving feed queries. Plain Go, same machine, same workload, same database engine underneath, peaked at 73.5 MB.

I spent the rest of the spike trying to explain that, because the honest thing to do with a number that surprises you is to attack it rather than publish it.

**Maybe it's Go's garbage collector being lazy about returning memory.** Go has a knob for exactly this, so I set it. `GOMEMLIMIT=256MiB` gave a peak of 732 MB. Setting it to 128 MiB gave 742 MB. The limit made no measurable difference at all, which rules out the comfortable explanation: this isn't heap the collector can be told to hand back.

**Maybe it's the record layer**, the part of PocketBase that turns database rows into schema-aware objects. So I bypassed it entirely and queried the database directly from inside the same server. Still 587 MB.

**Maybe it's the request log**, since PocketBase writes a row for every request by default and I was sending fifty thousand of them. Turned it off, with raw queries as well. Still 537 MB.

Three hypotheses, three tests, no explanation. Each fix shaved something off, so there are several contributors and no single culprit, and none of them gets within seven times of the thing underneath.

The memory does come back when the server goes quiet, so this is churn rather than a leak. That's a weaker defence than it sounds. The peak is what a machine has to survive, and the process that gets killed for using too much memory doesn't get to explain that it was about to give it back. At five thousand requests, which is a quiet evening for one busy circle, it still peaked at 555 MB.

So PocketBase is out, and what it was bringing, sign-in and file storage, is now mine to write. That's a real cost and I'd rather pay it than build a decade on a number nobody can account for.

## The one that lost on a coin flip

Dart had the smallest binary, the lowest idle memory, the fastest writes, and a cold start tied with Go. It is, on the numbers, the best of the four.

It lost to the tie-break I wrote down before measuring, which says that when two candidates both pass every threshold, take Go.

I want to be straight about that, because it would be easy to write this up as though Go won. It didn't win. It tied, and a rule I'd set in advance broke the tie. The argument for Dart, one language across the server and the Flutter app, is now better supported than it was before I measured, not worse. If I'd found Dart slower I'd have had a tidy reason to dismiss it. Instead I have a rule, and the reason for having rules like that is precisely so they hold on the days you'd rather they didn't.

Its lower throughput is a single-isolate artifact, not a language limit. Fixing it would mean building an isolate pool, which is work that isn't in those numbers.

## One database per circle

The other question the spike settled: does each circle get its own SQLite file, or do they all share one?

Shared is faster at the median and much worse at the tail. Sixteen writers on one file produced a worst case of 172 ms, because every writer queues behind one lock and a checkpoint stalls everyone. Per-circle files: 3.1 ms.

But the numbers that actually decided it are the boring ones. Exporting a circle from a shared database is a query. Exporting a per-circle database is a file copy. Deleting a circle from a shared database needs a vacuum that blocks every other circle on that server for 170 ms. Deleting a per-circle database is `rm`.

Every promise Nah? makes about a circle being able to pick up and leave rests on those operations being trivial. So: one file per circle.

## Digests turned out to be a non-event

The thing I was most worried about operationally was digests. Every other request in this product arrives when a human opens an app, nicely spread out. Digests fire at times people choose, and people choose round numbers, so they all land at once.

Five thousand circles, two thousand of them with something to say, all firing in the same minute: 0.70 seconds. Against a threshold of sixty.

So digests stay a loop inside the same binary. No queue, no worker, no broker. The real constraint is Apple and Google at the other end, which lives in the push relay and is a separate problem.

## What I decided while I was in there

With the framework gone, a pile of defaults came back on the table. So: a dependency budget of five, and the standard library everywhere it's good enough, which turns out to be almost everywhere.

The router is `net/http`. Logging is `log/slog`. Configuration is environment variables read into a struct. Queries are written-out SQL, because there are about ten of them and a code generator earns its place at fifty. Migrations are hand-rolled, stepping a version number when a circle's database is opened, because every migration library assumes one database and we have thousands.

The SQLite driver is the pure-Go one, so the whole thing compiles to one static binary with no C toolchain anywhere in the pipeline. That's two direct dependencies today.

There's no error reporting service either, which is less a technical decision than a rule the project already has: nothing in Nah? phones home, and that includes phoning home about itself.

## Home hardware is out

One more decision, from Casey rather than from a measurement: the server will not target a Raspberry Pi or a NAS. That follows from moving self-hosting out of the first release. The project hosts the circles to begin with.

Worth separating two things that sound the same, though. Not targeting home hardware isn't the same as not targeting ARM. Cheap ARM cloud instances are cheaper than their x86 equivalents, and Go cross-compiles to them for nothing, so the build matrix keeps both. It costs one line and keeps a lever available for when hosting other people's circles becomes real.

## Still no screenshot

The spike cost a day and produced a server that is not the product, four servers that will be deleted, and a folder of JSON. It also overturned a decision I'd made three days earlier with confidence and no evidence, which is the entire reason for doing it.

The next milestone hasn't changed: two phones, one circle, text moments, on a server I run. That's the one with a screenshot in it.

nah!
