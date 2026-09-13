---
title: "The Spike (or: the framework that wanted 700 megabytes)"
date: 2026-09-13 12:00:00 +0200
categories: [Architecture, Decisions]
tags: [spike, go, sqlite, benchmarks, decision-log]
---

I said the next post would have a screenshot in it. It doesn't. It has numbers, and the numbers changed what Nah? runs on, so I'm writing them down while I still remember what they mean.

## Limits first

The first commit of the spike wasn't code. It was a list of limits with nothing next to them: how much memory the server may use idle and under load, how fast a feed query may be, how big the binary may be, how long a cold start may take. Plus a tie-break and a two-week time box.

Writing them first is the only way I know to stop a benchmark turning into a hobby. If the numbers arrive before the limits, you argue yourself into whatever you already wanted.

Then four servers, two endpoints each: post a moment, list a circle's newest thirty. Same fixture for all of them. Fifty circles, a year of history, 116,800 moments.

## What came back

| | Dart | Plain Go | Go + PocketBase | Python |
|---|---|---|---|---|
| Idle memory | 15.9 MB | 18.9 MB | 54.7 MB | 278.2 MB |
| Peak under load | 56.8 MB | 74.2 MB | 719 MB | 404 MB |
| Binary or bundle | 6.5 MB | 15 MB | 32.9 MB | 38.3 MB |
| Cold start | 34 ms | 34 ms | 121 ms | 432 ms |
| Feed p95, 50 readers | 10.2 ms | 8.5 ms | 12.4 ms | 37.2 ms |
| Reads per second | 5,781 | 13,951 | 8,403 | 1,977 |

Python went first. 278 MB sitting idle across four workers against a limit of 150, before anyone has posted anything, and a seventh of Go's throughput. I had expected it to lose on packaging, since a virtual environment is a worse thing to hand someone than a file. It didn't get that far.

Then PocketBase, which I had picked three days earlier.

## 705 megabytes

PocketBase peaked at 705 MB serving feed queries. Plain Go, same machine, same workload, same database engine underneath, peaked at 73.5 MB.

I spent the rest of the day trying to explain that. Go has a knob for making the collector work harder, so I used it. At `GOMEMLIMIT=256MiB` the peak was 732 MB. At 128 MiB it was 742 MB. No effect, which rules out the comfortable answer.

Next guess was the record layer, the part that turns rows into schema-aware objects. I skipped it and queried the database directly from inside the same server: 587 MB. Then the request log, since PocketBase writes a row for every request by default and I was sending fifty thousand of them. Turned it off: 537 MB.

Three attempts, no answer. Each one shaved a bit off, so there are several contributors and no single cause, and none of them gets near 73.

The memory does come back when the server goes quiet, so it's churn rather than a leak. That matters less than it sounds, because the peak is what the machine has to survive. At five thousand requests, which is a quiet evening for one circle, it still hit 555 MB.

So PocketBase is out, and sign-in and file storage are mine to write now. I'd rather pay that than build on a number nobody can account for.

## Dart

Dart had the smallest binary, the lowest idle memory and the fastest writes, and it tied Go on cold start. On the numbers it's the best of the four.

It lost to the tie-break I wrote before measuring, which says that when two candidates both pass, take Go. So Go didn't win. It tied and got the benefit of a rule. The case for Dart, one language across the server and the app, is better supported now than it was before I measured. Its lower throughput is a single-isolate thing rather than a language thing, and fixing it means an isolate pool that isn't in those numbers.

I'm writing that down because in six months I'll want to remember it was close.

## One file per circle

The other question was whether each circle gets its own SQLite file or they all share one.

Shared is faster in the middle and worse at the edges. Sixteen writers on one file gave a worst case of 172 ms, because they queue behind a single lock and a checkpoint stalls everyone. Separate files: 3.1 ms.

What actually decided it is duller than that. Exporting a circle out of a shared database is a query; exporting a per-circle file is a copy. Deleting from a shared database needs a vacuum that blocks every other circle on the machine for 170 ms; deleting a file is `rm`. Nah? promises a circle can pick up and leave, and that promise is worth whatever those two operations are worth.

## Digests

Digests were the thing I was worried about, because they're the only load that isn't spread out. Everything else happens when someone opens the app. Digests happen at times people picked, and people pick round numbers.

Five thousand circles, two thousand of them with something to say, all firing in the same minute: 0.70 seconds against a limit of sixty. So they stay a loop inside the same binary. The real constraint is Apple and Google, and that lives in the push relay.

## The rest of it

With the framework gone I had to make the choices it had been making. The rule I landed on is a budget: five direct dependencies, standard library wherever it's good enough. Two are in use.

Router is `net/http`. Logging is `log/slog`. Configuration is environment variables read into a struct at startup. Queries are written out by hand, because there are about ten of them. Migrations are hand-rolled, stepping a version number when a circle's file is opened, since every migration library assumes one database and we have thousands of them.

The SQLite driver is the pure-Go one, so the server compiles to a static binary with no C toolchain anywhere in the pipeline. There's no error reporting service either, which follows from a rule the project already had.

## Hardware

One decision that came from me rather than from the spike: no Raspberry Pi, no NAS. That follows from hosting the circles myself to start with.

Not targeting home hardware isn't the same as not targeting ARM, though. ARM cloud boxes are cheaper than x86 ones and Go cross-compiles to them for free, so the build keeps both.

## Still no screenshot

A day of work, four servers that get deleted, one that isn't the product, and a folder of JSON. It also overturned a decision I'd made three days earlier with confidence and no evidence.

Next is the walking skeleton: two phones, one circle, text moments, on a server I run. Screenshot in that one.

nah!
