# Nah

**A private home for your closest people.**

> *"Share with your circle, not the world."*

Nah ("Not Alone Here") is a private social network built out of circles — a family circle, a close-friends circle, a neighbourhood circle — each one capped at Dunbar's number. Inspired by [Path](https://en.wikipedia.org/wiki/Path_(social_network)), built on modern open-source infrastructure, and designed to never compromise on intimacy.

**Viral? Nah. Vital.**

---

## Core Principles

- **Small by design** — every circle is capped at Dunbar's number, and you belong to a few of them
- **Private by default** — no public timelines, no algorithmic feeds
- **Real friends, real moments** — mutual friendships, rich moment sharing
- **No algorithmic theater** — chronological feed only, no engagement optimization
- **Open source** — community-funded, transparent, yours to verify

---

## Built With

Nah stands on the shoulders of giants. We're grateful to these open source projects:

| Project | What we use it for |
|---------|-------------------|
| [@benbjohnson/litestream](https://github.com/benbjohnson/litestream) | Continuous SQLite backup |
| [@flutter/flutter](https://github.com/flutter/flutter) | Cross-platform mobile app framework (iOS + Android) |
| [@dart-lang/sdk](https://github.com/dart-lang/sdk) | Dart language and SDK |
| [@bloclibrary/bloc](https://github.com/felangel/bloc) | State management |
| [@dio-package/dio](https://github.com/cfug/dio) | HTTP client |
| [@modernc-org/sqlite](https://gitlab.com/cznic/sqlite) | Pure-Go SQLite driver, so the server needs no cgo |

### Tooling

| Project | What we use it for |
|---------|-------------------|
| [OpenSpec](https://openspec.dev) · [@openspec](https://github.com/openspec) | Specification-driven development workflow |
| [@caddyserver/caddy](https://github.com/caddyserver/caddy) | Web server for marketing site (planned) |

---

## Tech Stack

| Layer | Technology |
|-------|------------|
| **Mobile app** | Flutter 3.41+ / Dart 3.11+ (iOS + Android) |
| **State** | Bloc / Cubit |
| **Networking** | Dio + the circle server's own API |
| **Backend** | Nah? Home — one static Go binary, `net/http`, no framework |
| **Database** | SQLite, one file per circle, streamed off-box with Litestream |
| **Storage** | Local disk or S3-compatible — ciphertext either way |

### Stack evolution

| Date | Decision | Why |
|---|---|---|
| 2026-01 | ~~Svelte 5 + SvelteKit (PWA), Tailwind v4, shadcn-svelte, TanStack Query~~ | Act small, ship a PWA MVP. Mobile-first web. |
| 2026-05-19 | Flutter 3.41+ / Dart, Bloc, Dio, Hive/Drift, Mastodon backend unchanged | Pivot to native via Flutter. See [ADR-0001 — Flutter over PWA](docs/decisions/0001-flutter-over-pwa.md) and the [Flutter decision blog post](https://caseyro.github.io/Nah/2026/05/19/the-flutter-decision/). |
| 2026-09-13 | ~~Mastodon fork~~ → one Go binary on PocketBase with SQLite | Circles became the core data model, and nothing off the shelf has them. See [ADR-0009](docs/decisions/0009-circles-not-one-circle.md), [ADR-0010](docs/decisions/0010-small-server-not-mastodon.md) and the [circles research](docs/research/circles-and-the-backend.md). |
| 2026-09-13 | ~~PocketBase~~ → plain Go, one SQLite file per circle | Measured rather than argued: PocketBase peaked at 705 MB against plain Go's 73.5 MB, unexplained by three experiments. See [ADR-0011](docs/decisions/0011-plain-go-and-a-database-per-circle.md), [ADR-0014](docs/decisions/0014-the-go-setup.md) and the [spike results](https://caseyro.github.io/Nah/research/spike-results/). |
| 2026-09-13 | Content encrypted on the device; sign-in is an Ed25519 key, not a password | The server has no use for plaintext, so it should not hold any. See [ADR-0012](docs/decisions/0012-encrypted-on-device.md) and [ADR-0015](docs/decisions/0015-no-passwords-a-key-on-the-device.md). |

The PWA framing is preserved struck-through in the OpenSpec history (`openspec/changes/cap-01-core-identity/specs/pwa-shell/spec.md`) so the evolution is visible, not erased.

See [design.md](openspec/changes/nah-vision/design.md) for full architecture details.

---

## Project Structure

```text
/apps
  /mobile           # Flutter app — iOS + Android (coming soon)
  /server           # Nah? Home — Go + SQLite, one database per circle
/packages
  /ui               # Nah design system: ThemeData, widgets, tokens (coming soon)
/docs
  /_posts           # Build-in-public blog posts
/openspec
  /changes          # Feature specifications
    /nah-vision     # Core vision (complete)
    /cap-01-*       # Capability implementations (scaffolded)
```

---

## Documentation

- **[Vision & Specs](openspec/changes/nah-vision/)** — Why, what, and how
- **[Research](docs/research/)** — Concept audit and the retention investigation
- **[Decisions](docs/decisions/)** — Architecture decision records
- **[Reflections](https://caseyro.github.io/Nah/)** — Build-in-public blog

---

## Status

🚧 **Pre-alpha** — Vision documented, implementation starting.

### Progress

- [x] Vision proposal
- [x] Technical design
- [x] Feature specifications (8 capabilities)
- [ ] Project scaffolding
- [ ] Core implementation
- [ ] Alpha release

---

## Development

### Prerequisites

- Flutter 3.41+ (stable channel)
- Dart 3.11+
- Xcode 16+ (for iOS builds)
- Android Studio / Android SDK (for Android builds)
- Docker (for Mastodon backend)
- Ruby 3.3+ (for local docs preview)

### Pre-commit Hooks

We use [pre-commit](https://pre-commit.com/) for code quality:

```bash
# Install pre-commit
brew install pre-commit  # or pip install pre-commit

# Install hooks
pre-commit install

# Run manually
pre-commit run --all-files
```

---

## Contributing

Nah is open source and community-driven. We're not ready for contributions yet, but we will be soon.

**Watch this repo** to follow along as we build.

---

## License

[AGPL-3.0](LICENSE) — Same as Mastodon.

---

## Support

Nah is community-funded. No ads, no data sales, no VC.

Support options coming soon via Open Collective and GitHub Sponsors.

---

*"Nah is a private home for your closest people — to share life without performing for the internet."*
