# Server

## Purpose [coverage: high -- 6 sources]

Sources span 2026-01-28 to 2026-09-14. Everything below describes the state on 2026-09-14 unless it carries its own date.

The server is the one part of Nah? that runs somewhere other than a phone: a single static Go binary that keeps one SQLite file per person. It stores the ciphertext of moments, orders it by time and hands back the newest page, and it never reads a moment (see [opaque moments](../concepts/opaque-moments.md)). Every copy of the [app](mobile.md) talks to it, and nothing else does yet.

It was built on 2026-09-13 around circles and reworked the next day for [one network of up to 150 per person](../../decisions/0017-one-network-of-a-hundred-and-fifty.md), after the feed that change implied had been [measured](spike.md). The root README calls it Nah? Home.

## Architecture [coverage: medium -- 3 sources]

The [server README](../../../apps/server/README.md) is the canonical description of the routes, the exact string a device signs, and the data directory. The code is small enough to read directly:

- [`cmd/server/main.go`](../../../apps/server/cmd/server/main.go): configuration from two environment variables, and a graceful shutdown that closes every open file.
- [`internal/nah/store.go`](../../../apps/server/internal/nah/store.go): person files, their migrations, connecting two people, and the feed.
- [`internal/nah/http.go`](../../../apps/server/internal/nah/http.go): routes, request logging, and the sentences a person reads when something is refused.
- [`internal/nah/auth.go`](../../../apps/server/internal/nah/auth.go): challenges, and session tokens that survive a restart.
- [`internal/nah/id.go`](../../../apps/server/internal/nah/id.go): time-ordered ids, and the guard that stops an id from reaching the filesystem as anything but a file name.
- [`internal/nah/nah_test.go`](../../../apps/server/internal/nah/nah_test.go): the tests, including two people connecting end to end and a moment crossing between two instances.

On disk the data directory holds one `<person id>.db` per person, plus `session.key`. The image that runs it is described under [delivery](delivery.md).

## Talks To [coverage: medium -- 4 sources]

- **The app**, over versioned routes under `/v1/people`. The app registers a device key with an invitation, signs a challenge with it, and uses the resulting session for invites, connections, the feed and posting.
- **Two Go modules and the standard library**: the pure-Go SQLite driver and `golang.org/x/time/rate`, inside a budget of five ([ADR-0014](../../decisions/0014-the-go-setup.md)).
- **Nothing central yet.** [ADR-0010](../../decisions/0010-small-server-not-mastodon.md) names the only central services the project intends to run: a push relay, the invite-link domain with a directory from an identifier to its current address, and an account for App Review. None of them exists.

Streaming backups with Litestream are planned. Nothing in `apps/server` runs them yet, and restoring from a backup has not been measured.

## Key Decisions [coverage: high -- 13 sources]

Newest first. Each decision record is the canonical home for its reasoning.

- **2026-09-15:** reactions will be stored sealed, so the server knows who reacted to which moment and when but not how, and returns them only to the poster ([ADR-0018](../../decisions/0018-reactions-the-poster-sees.md)). Not built.
- **2026-09-14:** registration is by invitation only. An invite from someone already here also connects the two people, and the first person on a server uses a single-use operator invite; the [server README](../../../apps/server/README.md) describes both.
- **2026-09-14:** the feed reads fan-in from each connection's file, newest poster first, and stops once nobody left can beat the page. Reading every file failed the latency threshold at 150 connections. See the 2026-09-14 amendment to [ADR-0011](../../decisions/0011-plain-go-and-a-database-per-circle.md) and the [spike results](../../../spike/RESULTS.md).
- **2026-09-13:** one network of up to 150 per person, no circles and one feed ([ADR-0017](../../decisions/0017-one-network-of-a-hundred-and-fifty.md)). It supersedes circles ([ADR-0009](../../decisions/0009-circles-not-one-circle.md)) and rooms ([ADR-0013](../../decisions/0013-no-aggregated-feed.md)), and brings back invitation as connection from [ADR-0003](../../decisions/0003-connection-model.md).
- **2026-09-13:** sign-in is an Ed25519 key on the device answering a server challenge, and sessions are self-contained tokens signed by `session.key`, so a restart logs nobody out ([ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md)).
- **2026-09-13:** a moment is a versioned envelope with a fallback sentence, not markup, and the server never looks inside it ([ADR-0016](../../decisions/0016-the-moment-envelope.md)).
- **2026-09-13:** content is encrypted on the device and the project holds no keys ([ADR-0012](../../decisions/0012-encrypted-on-device.md), not built until CDI-1863).
- **2026-09-13:** a dependency budget of five, the standard library wherever it is good enough, no cgo, and migrations stepped per file ([ADR-0014](../../decisions/0014-the-go-setup.md)).
- **2026-09-13:** plain Go with a SQLite file per unit of data, after PocketBase measured at seven times the memory ([ADR-0011](../../decisions/0011-plain-go-and-a-database-per-circle.md)).
- **2026-09-13:** a small server of our own rather than Mastodon ([ADR-0010](../../decisions/0010-small-server-not-mastodon.md)), superseding the Mastodon fork chosen on 2026-01-28 ([ADR-0002](../../decisions/0002-mastodon-backend.md)).

## Running It [coverage: medium -- 3 sources]

From `apps/server`:

```bash
go run ./cmd/server
```

It serves on `:8080` and keeps its files in `./data`; `NAH_ADDR` and `NAH_DATA_DIR` change either. Nah? is by invitation, so the first person on a server needs `go run ./cmd/server invite`, which prints a single-use invitation to paste into the app; everyone after them joins through someone already there. The container build and the full set of checks are in the [server README](../../../apps/server/README.md), and the same checks run on every push that touches the server ([delivery](delivery.md)).

## Gotchas [coverage: high -- 6 sources]

- **`session.key` is the one file in the data directory that is not ciphertext.** Lose it and everyone signs in again; leak it and anyone can mint a session for any key.
- **A session cannot be revoked before its 30 days are up.** Replacing the device key is what revokes access, which is why the key is checked on every request ([ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md)).
- **Challenges live in one process's memory.** With two instances behind one address, a challenge issued by one and answered at the other fails ([ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md)). See [invisible deploys](../concepts/invisible-deploys.md).
- **Do not add a query per connection to the feed.** The measured cost of a feed is SQLite's lock around every read transaction, not the queries themselves, so one more read per connection undoes the win ([spike results](../../../spike/RESULTS.md)).
- **Open files stay open for the life of the process.** That is seven to nine descriptors and about a quarter of a megabyte per person with a file open. A host holding thousands of people needs an eviction policy first.
- **Stop it with a signal, never a kill.** A graceful shutdown closes and checkpoints every file; a killed process leaves up to 4 MB of write-ahead log behind per person.
- **Migrations are additive only.** During a rolling deploy the old instance still reads files the new one has migrated; the rule is written on the migrations in `store.go`.
- **A person's invites do not expire, are not single-use and cannot be revoked yet**, and registration shares one server-wide rate limit with sign-in ([server README](../../../apps/server/README.md)).
- **An operator invite is spent the moment it is redeemed**, even if registering fails afterwards. Make another with `server invite`; a refused registration never leaves a person behind.
- **ADR-0011's title still says "per circle".** The amendment at its end is the current state.

## Sources

- [apps/server/README.md](../../../apps/server/README.md)
- [apps/server/Dockerfile](../../../apps/server/Dockerfile)
- [README.md](../../../README.md)
- [spike/RESULTS.md](../../../spike/RESULTS.md)
- [.github/workflows/server.yml](../../../.github/workflows/server.yml)
- [docs/decisions/0002-mastodon-backend.md](../../decisions/0002-mastodon-backend.md)
- [docs/decisions/0003-connection-model.md](../../decisions/0003-connection-model.md)
- [docs/decisions/0009-circles-not-one-circle.md](../../decisions/0009-circles-not-one-circle.md)
- [docs/decisions/0010-small-server-not-mastodon.md](../../decisions/0010-small-server-not-mastodon.md)
- [docs/decisions/0011-plain-go-and-a-database-per-circle.md](../../decisions/0011-plain-go-and-a-database-per-circle.md)
- [docs/decisions/0012-encrypted-on-device.md](../../decisions/0012-encrypted-on-device.md)
- [docs/decisions/0013-no-aggregated-feed.md](../../decisions/0013-no-aggregated-feed.md)
- [docs/decisions/0014-the-go-setup.md](../../decisions/0014-the-go-setup.md)
- [docs/decisions/0015-no-passwords-a-key-on-the-device.md](../../decisions/0015-no-passwords-a-key-on-the-device.md)
- [docs/decisions/0016-the-moment-envelope.md](../../decisions/0016-the-moment-envelope.md)
- [docs/decisions/0017-one-network-of-a-hundred-and-fifty.md](../../decisions/0017-one-network-of-a-hundred-and-fifty.md)
- [docs/decisions/0018-reactions-the-poster-sees.md](../../decisions/0018-reactions-the-poster-sees.md)
