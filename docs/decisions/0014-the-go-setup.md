---
title: "ADR-0014 — The Go setup, and how it ships"
permalink: /decisions/0014-the-go-setup/
layout: page
status: Accepted
date: 2026-09-13
---

## Status

Accepted (2026-09-13). Follows [ADR-0011](0011-plain-go-and-a-database-per-circle.html), which chose plain Go and left every library question open.

## Context

Dropping PocketBase removed a framework and, with it, a set of default answers about routing, storage, migrations, sign-in and configuration. Those answers now have to be chosen deliberately.

The spike also produced a warning worth generalising. PocketBase failed on memory for reasons three experiments could not identify, in code nobody here had read. Every dependency is a promise to debug someone else's work at three in the morning, and the server's entire job is to store opaque blobs and return the newest thirty of them.

## Decision

**A dependency budget of five direct dependencies through M1**, and the standard library everywhere it is good enough. It is good enough almost everywhere here.

| Concern | Choice | Instead of |
|---|---|---|
| HTTP and routing | `net/http`, with method and path patterns | chi, echo, gin |
| SQLite driver | `modernc.org/sqlite`, no cgo | mattn/go-sqlite3 |
| Migrations | `PRAGMA user_version`, stepped on open | golang-migrate, goose |
| Queries | `database/sql` and written-out SQL | sqlc, any ORM |
| IDs | UUIDv7, hand-rolled | google/uuid, ULID libraries |
| Logging | `log/slog`, JSON to stdout | zap, zerolog |
| Configuration | Environment variables into a struct at start | viper, koanf |
| Rate limiting | `golang.org/x/time/rate` | a middleware framework |
| Tests | `testing` and `net/http/httptest`, with `-race` | testify |
| Error reporting | Structured logs, and nothing that phones home | Sentry as a service |

That is two direct dependencies today: the SQLite driver, and `x/time/rate` when the auth endpoints exist.

**Why no cgo.** `CGO_ENABLED=0` gives one static binary per architecture, cross-compiled from anything, with no C toolchain in CI and no glibc surprises in a container. The spike measured the pure-Go driver at 73.5 MB peak and 13,951 reads a second, which is far more headroom than 150 people need. If raw write throughput ever binds, the cgo driver is the escape hatch, and the cost of taking it is exactly the cross-compilation we are buying here.

**Why migrations are hand-rolled.** With one database per circle there is no single database to migrate. A migration tool that assumes one file and one version table is the wrong shape for thousands of files. Stepping `user_version` when a circle's database is opened is about forty lines, runs per file, and is the only approach that degrades sensibly when a server holds circles that have not been opened in months.

**Why no code generation for queries.** There are roughly ten queries in the whole product. A generation step earns its place at fifty.

**Error reporting cannot be a service.** The project's own guardrail forbids anything that phones home, which rules out Sentry as a hosted product despite it appearing in the original technical design. Structured logs to stdout, and a self-hosted collector later if logs stop being enough.

### How it ships

- **Build:** `CGO_ENABLED=0` for `linux/amd64` and `linux/arm64`. Home hardware is no longer a target, but ARM cloud instances are cheaper than x86 ones and cost one line in the matrix.
- **Checks on every push:** `go vet`, `staticcheck`, `go test -race`, and `govulncheck`.
- **Release:** a plain build matrix producing a container image to the registry. Goreleaser is the right tool for shipping binaries to strangers, which is M7's problem, not M1's.
- **Version numbers belong to the release workflow.** The fleet convention bumps the patch version and tags on every push to main, so the version is never edited by hand and a feature branch is rebased before it is fast-forwarded.
- **Deploy is a Komodo git webhook**, and the webhook secret must be set and identical on both the GitHub side and the Komodo stack. An empty secret produces a 401 on every delivery, with no deploy and no visible error, which is the fleet's most expensive silent failure.
- **A green push is not a release.** Check that the tag appeared and the deploy event fired.

## Considered alternatives

- **A web framework, for the middleware.** Rejected. About ten routes need logging, authentication and rate limiting, which is roughly thirty lines of handler wrapping. We just removed a framework for costing seven times the memory of the thing underneath it.
- **sqlc.** Genuinely good, and the type safety is real. Rejected for now because it adds a generation step to a codebase with ten queries, and because hand-written SQL in one place is easy to read when the queries are this small. Worth revisiting rather than dismissed.
- **A migration library.** Rejected on shape, not quality: per-circle databases break the assumption every one of them makes.
- **cgo SQLite for speed.** Rejected because it trades cross-compilation, CI simplicity and static linking for throughput we measured we do not need.
- **Goreleaser now.** Rejected as premature. It becomes correct the moment a stranger installs the server themselves.

## Consequences

Positive:

- Two direct dependencies means almost nothing to audit, upgrade or be surprised by, and `go build` needs no toolchain beyond Go.
- One static binary per architecture, which is what makes both container deployment and the eventual self-hosted rung simple.
- The code that matters is the code we wrote, which is the lesson the spike paid for.

Negative / acknowledged cost:

- Sign-in, file storage and any admin tooling are ours to write, with no upstream to inherit fixes from.
- Hand-rolled migrations are a small piece of infrastructure that has to be right, because it runs against every circle's database.
- A dependency budget is a rule that will feel wrong at some point. It is meant to make adding a dependency a decision rather than a reflex, not to make it impossible.

## What would prove this wrong

- The hand-rolled migration stepper corrupts or blocks a circle in a way a library would have prevented.
- Hand-written SQL becomes the place bugs live once the query count grows past about thirty.
- The pure-Go SQLite driver becomes the bottleneck under real write load, which the spike says is far away but did not test on the eventual hosting hardware.
- Writing sign-in takes materially longer than the week or two assumed, which would retrospectively make PocketBase's memory cost worth paying.
