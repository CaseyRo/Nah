# Nah? circle server

One static Go binary. One SQLite file per circle. It stores ciphertext, orders
it by time, and hands back the newest few. It never reads a moment.

The decisions behind it: [ADR-0011](../../docs/decisions/0011-plain-go-and-a-database-per-circle.md)
(plain Go, a database per circle), [ADR-0012](../../docs/decisions/0012-encrypted-on-device.md)
(encrypted on the device), [ADR-0014](../../docs/decisions/0014-the-go-setup.md)
(the libraries, and how it ships) and [ADR-0015](../../docs/decisions/0015-no-passwords-a-key-on-the-device.md)
(sign-in).

## Run it locally, with a circle in it

```bash
go run ./cmd/server create-circle && go run ./cmd/server
```

The first half prints a circle id and an invite token. The second half serves on
`:8080`. Both halves read the same two environment variables:

| Variable | Default | What it is |
|---|---|---|
| `NAH_ADDR` | `:8080` | listen address |
| `NAH_DATA_DIR` | `./data` | where the circle files and `session.key` live |

In a container:

```bash
docker build -t nah-server .
docker run --rm -v "$PWD/data:/data" nah-server create-circle
docker run --rm -p 8080:8080 -v "$PWD/data:/data" nah-server
```

## The API

Every route is scoped to one circle. There is no route that spans circles,
because there is no aggregated anything ([ADR-0013](../../docs/decisions/0013-no-aggregated-feed.md)).

| | |
|---|---|
| `GET /healthz` | is it up |
| `POST /v1/circles/{circle}/join` | `{invite, public_key}` → 204 |
| `POST /v1/circles/{circle}/challenge` | `{public_key}` → `{challenge}` |
| `POST /v1/circles/{circle}/session` | `{public_key, challenge, signature}` → `{token, expires_at}` |
| `GET /v1/circles/{circle}/moments?limit=30` | newest first |
| `POST /v1/circles/{circle}/moments` | `{blob}` → the stored moment |

Byte fields (`public_key`, `signature`, `blob`) are base64 in JSON, which is
what Go's `encoding/json` does with `[]byte` and what Dart's `base64` produces.

`blob` is the ciphertext of a moment envelope ([ADR-0016](../../docs/decisions/0016-the-moment-envelope.md)).
Nothing on this side knows or may ever learn what is in it — not the type, not
the length of the text, not whether it has a photo.

The two moment routes want `Authorization: Bearer <token>`.

### Signing in

No passwords. The device holds an Ed25519 keypair, asks for a challenge, and
signs this exact string:

```text
nah-auth-v1:<circle id>:<challenge>
```

The prefix keeps the signature from meaning anything anywhere else, and the
circle id keeps a signature captured on one circle from opening another.

A session token carries its own claims and an HMAC over them, so it survives a
restart and works across two instances sharing the data directory. Deploys are
meant to be frequent and invisible, so nothing about a session is stored.

The signing key is `session.key` in the data directory, generated on first run
at `0600`. **It is the one file in there that is not ciphertext.** Lose it and
everyone signs in again; leak it and anyone can mint a session for any key.

The client should still re-authenticate on a 401 rather than showing a login
screen, but it is no longer a per-deploy event.

## What is not here yet

- **Media.** Text moments only; the request body is capped at 1 MiB. Photo and
  voice are M3, and they get an upload path rather than this one.
- **The invite link.** The server issues the token in the path half. The content
  key lives after the `#` and the server never sees it (M2, CDI-1836).
- **Creating a circle over HTTP.** It is a subcommand, so there is no admin
  endpoint to defend. M2's directory work replaces it.
- **Names.** Nobody's display name is on the server. What a member is called
  travels inside the ciphertext, where it belongs.
- **Removing a member.** M5. Membership is checked on every request so that it
  will take effect immediately when it arrives.

## Checks

```bash
go vet ./... && go test -race ./...
go run honnef.co/go/tools/cmd/staticcheck@latest ./...
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

`TestTwoDevicesOneCircle` is CDI-1835 with the phones taken out: two devices
join one circle, each posts, and each sees both.
