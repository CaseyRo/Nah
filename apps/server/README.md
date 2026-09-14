# Nah? server

One static Go binary. One SQLite file per person. It stores ciphertext, orders
it by time, and hands back the newest few. It never reads a moment.

The decisions behind it: [ADR-0011](../../docs/decisions/0011-plain-go-and-a-database-per-circle.md)
(plain Go, a database per file, amended for people), [ADR-0012](../../docs/decisions/0012-encrypted-on-device.md)
(encrypted on the device), [ADR-0014](../../docs/decisions/0014-the-go-setup.md)
(the libraries, and how it ships), [ADR-0015](../../docs/decisions/0015-no-passwords-a-key-on-the-device.md)
(sign-in) and [ADR-0017](../../docs/decisions/0017-one-network-of-a-hundred-and-fifty.md)
(one network of 150 per person).

## Run it locally

```bash
go run ./cmd/server
```

It serves on `:8080`. There is nothing to create first: a person registers from
the app. It reads two environment variables:

| Variable | Default | What it is |
|---|---|---|
| `NAH_ADDR` | `:8080` | listen address |
| `NAH_DATA_DIR` | `./data` | where the person files and `session.key` live |

In a container:

```bash
docker build -t nah-server .
docker run --rm -p 8080:8080 -v "$PWD/data:/data" nah-server
```

## The API

A person is a file. Every route after registration acts as the person in the
path, and the feed is the only one that reads anybody else's file.

| | |
|---|---|
| `GET /healthz` | is it up |
| `POST /v1/people` | `{public_key}` → `{id}` |
| `POST /v1/people/{person}/challenge` | `{public_key}` → `{challenge}` |
| `POST /v1/people/{person}/session` | `{public_key, challenge, signature}` → `{token, expires_at}` |
| `POST /v1/people/{person}/invites` | → `{invite}` |
| `POST /v1/people/{person}/connections` | `{person, invite}` → 204, connected both ways |
| `GET /v1/people/{person}/feed?limit=30` | you and your people, newest first |
| `POST /v1/people/{person}/moments` | `{blob}` → the stored moment |

The last four want `Authorization: Bearer <token>` for the person in the path.

Byte fields (`public_key`, `signature`, `blob`) are base64 in JSON, which is
what Go's `encoding/json` does with `[]byte` and what Dart's `base64` produces.

`blob` is the ciphertext of a moment envelope ([ADR-0016](../../docs/decisions/0016-the-moment-envelope.md)).
Nothing on this side knows or may ever learn what is in it — not the type, not
the length of the text, not whether it has a photo.

A moment's `id` is unique among its author's moments, not across a feed;
`(author_id, id)` is.

### Registering and connecting

Registration is open, like installing the app, and shares the unauthenticated
rate limit with sign-in. A person with no connections can read nothing and reach
nobody.

To connect, one person makes an invite and the other redeems it: `{person}` in
the path is whoever redeems, and the body names who made it. A touch and a link
carry the same two things, so there is one route for both and the server cannot
tell them apart ([CDI-1839](https://linear.app/cdit/issue/CDI-1839), [CDI-1840](https://linear.app/cdit/issue/CDI-1840)).
Connection is mutual and immediate. The cap of 150 is counted on both sides, and
a full network is refused in words, never a number ([ADR-0004](../../docs/decisions/0004-no-counts-anywhere.md)).

### Signing in

No passwords. The device holds an Ed25519 keypair, asks for a challenge, and
signs this exact string:

```text
nah-auth-v1:<person id>:<challenge>
```

The prefix keeps the signature from meaning anything anywhere else, and the
person id keeps a signature made for one person from signing in as another.

A session token carries its own claims and an HMAC over them, so it survives a
restart and works across two instances sharing the data directory. Deploys are
meant to be frequent and invisible, so nothing about a session is stored.

The signing key is `session.key` in the data directory, generated on first run
at `0600`. **It is the one file in there that is not ciphertext.** Lose it and
everyone signs in again; leak it and anyone can mint a session for any key.

The client should still re-authenticate on a 401 rather than showing a login
screen, but it is no longer a per-deploy event.

### The feed

A feed reads fan-in from the files of the people you are connected to. That was
measured before it was built ([CDI-1879](https://linear.app/cdit/issue/CDI-1879),
[spike results](../../spike/RESULTS.md)): reading all 150 files on every request
failed the latency threshold, because SQLite takes a file lock per read. So
`Feed` visits people most recent poster first, reads only ids and times, and
stops as soon as nobody left can have anything newer than the page. It reads at
most `limit + 1` files, however many connections there are.

What it remembers about who posted last is read again from the file after 30
seconds. A moment posted to another instance on the same volume, as during a
rolling deploy, is in everyone's feed within that.

## What is not here yet

- **Media.** Text moments only; the request body is capped at 1 MiB. Photo and
  voice are M3, and they get an upload path rather than this one.
- **Invite states.** An invite does not expire, is not single-use and cannot be
  revoked yet (CDI-1842). The link format, with the content key after the `#`
  where the server never sees it, is CDI-1836.
- **Names.** Nobody's display name is on the server. What a person is called
  travels inside the ciphertext, where it belongs.
- **A lost phone, or a second one.** M5 (CDI-1865). The key is checked on every
  request so that replacing it takes effect immediately.
- **Disconnecting.** Nothing removes a connection yet.

## Checks

```bash
go vet ./... && go test -race ./...
go run honnef.co/go/tools/cmd/staticcheck@latest ./...
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

`TestTwoPeopleConnected` is CDI-1835 with the phones taken out: two people
connect, each posts, and each sees both. `TestFeedMatchesReadingEveryFile` holds
the feed's shortcut to exactly the page that reading every file would give.
