# Nah? vision design

## Context

In January this document designed a Mastodon fork, and none of it was built. Since 2026-09-13 Nah? runs on a small server of its own, and the reasons are in the decision records. This page maps each part of Nah? to where it is decided and where it is explained. It holds no decisions of its own.

## The architecture

| Part | What it is now | Decided in | Explained in |
|---|---|---|---|
| The app | Flutter, for iOS and Android | ADR-0001 | `docs/wiki/topics/mobile.md` |
| The server | One small Go server of our own, not Mastodon | ADR-0010, ADR-0014 | `docs/wiki/topics/server.md` |
| Storage | SQLite, one database file per person | ADR-0011, amended 2026-09-14 | `docs/wiki/topics/spike.md` |
| Your circle | One network of up to 150 per person, called your circle, read as a fan-in feed | ADR-0017 | `docs/wiki/topics/server.md` |
| Sign-in | An Ed25519 key on the device, and stateless session tokens | ADR-0015 | `docs/wiki/topics/mobile.md` |
| Moments | Sealed on the device in a versioned envelope; the server stores bytes | ADR-0012, ADR-0016 | `docs/wiki/concepts/opaque-moments.md` |
| Run centrally | A push relay, the invite-link domain, and a demo for App Review; the directory from M7 (CDI-1838) | ADR-0010 | `docs/wiki/topics/delivery.md` |

## What the January design became

| January | Now | Where |
|---|---|---|
| Vanilla Mastodon fork on Rails, PostgreSQL and Redis | Plain Go and SQLite | ADR-0010, ADR-0011 |
| ActivityPub, with federation off for now | No federation | ADR-0008, ADR-0010 |
| Followers-only posts, enforced by the API | Moments the server cannot read | ADR-0012 |
| Friend requests and an inner circle | One mutual circle per person, with no subsets | ADR-0017 |
| Email and password accounts | A key on the device | ADR-0015 |
| Mastodon's streaming API | Load on launch, on foreground and on pull | CDI-1882 `app-shell` |
| Hive or Drift for offline use | The stored feed is specified in CDI-1882, the offline queue in CDI-1884 | CDI-1882, CDI-1884 |
| Sentry, Prometheus and Grafana | No error-reporting service; nothing reports home | ADR-0014 |
| S3-compatible storage with a CDN | Not decided; the media path is still to be measured | CDI-1824 |
| `/packages/nah_ui` and `/packages/nah_api` | `apps/mobile` and `apps/server` only | CDI-1882 design.md |
| Push, badges and a notification for every event | A daily digest by default, and no badges | ADR-0007 |
| Messaging deferred to v1.5 | No chat in Nah? | `product-principles` |

## Open Questions

The January open questions, answered or moved:

- **Music integration and location precision:** out, with those moment types (ADR-0006).
- **Finding first friends without discovery:** invitation, and the touch (ADR-0017).
- **Moderation in a circle of 150:** CDI-1867, reporting and blocking when the server cannot read anything.
- **Media storage costs:** CDI-1824.
- **A widget gallery:** not needed until the design system exists (CDI-1882).
