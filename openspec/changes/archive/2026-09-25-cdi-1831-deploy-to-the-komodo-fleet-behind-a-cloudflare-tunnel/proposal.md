# Deploy to the Komodo fleet behind a Cloudflare tunnel

Linear: [CDI-1831](https://linear.app/cdit/issue/CDI-1831)

## Why

M1 ends with two real phones on a deployed server (CDI-1835), and no server is deployed. This puts the one Nah? server on infrastructure that already exists, as the development and dogfooding server for Casey's own circle. Until production exists it is also staging, the gate for archiving changes (`CLAUDE.md`).

## What Changes

- `apps/server/compose.yaml`: one service built in place from the repository, a named volume for `/data`, a 30 second stop grace period so every person file is checkpointed on a deploy, and the port bound to nebula-1's Tailscale address for the tunnel.
- A Komodo stack on nebula-1 that deploys on every push to `main` through a GitHub webhook, with the secret set on both sides (ADR-0014).
- `nah.casey.berlin` through the fleet's Cloudflare tunnel, with no Cloudflare Access in front: the app signs in with its own key (ADR-0015), and Access would lock the phones out.
- A Gatus endpoint on `/healthz`.
- ADR-0014 gains a dated note: Komodo builds from git, as the rest of the fleet does, and a registry waits for M7.
- No specs change. `skip_specs: true` keeps validation passing; this change delivers infrastructure, not behaviour.

## Decided on 2026-09-25 (Casey)

- nebula-1, not werkstatt-1, whose disk is 86% full.
- Build from git in place, not a registry image.
- `nah.casey.berlin`.
- No backups until Litestream. Accepted risk: a lost volume loses the dogfooding circle. CDI-1826 remains the drill.

## Impact

- The app is built with `--dart-define=NAH_SERVER=https://nah.casey.berlin` for real phones.
- `docs/wiki/topics/delivery.md` stops saying the deploy does not exist.
