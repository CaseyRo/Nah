# Nah?

**A private home for your closest people.**

> *"Viral? Nah. Vital."*

Nah? ("Not Alone Here") is a private social network for the people you are actually close to. Each person has one network of up to 150 people, connected in person first, with one chronological feed and nobody to perform for. Inspired by [Path](https://en.wikipedia.org/wiki/Path_(social_network)), built on open source software, and designed never to compromise on intimacy.

---

## Core Principles

- **Small by design**: your network is capped at Dunbar's number, said in words and never shown as a count
- **Private by default**: no public timelines, no algorithmic feeds, and a server that does not read what you share
- **By invitation**: nobody joins without an invitation from someone already there
- **Real friends, real moments**: every connection is mutual, and made in person first
- **No algorithmic theater**: one chronological feed, no engagement optimization
- **Open source**: community-funded, transparent, yours to verify

---

## Built With

Nah? stands on the shoulders of giants. We're grateful to these open source projects:

| Project | What we use it for |
|---------|-------------------|
| [@flutter/flutter](https://github.com/flutter/flutter) | Cross-platform mobile app framework (iOS and Android) |
| [@dart-lang/sdk](https://github.com/dart-lang/sdk) | Dart language and SDK |
| [@bloclibrary/bloc](https://github.com/felangel/bloc) | State management |
| [@dio-package/dio](https://github.com/cfug/dio) | HTTP client |
| [@dint-dev/cryptography](https://github.com/dint-dev/cryptography) | The Ed25519 key each phone signs in with |
| [@juliansteenbakker/flutter_secure_storage](https://github.com/juliansteenbakker/flutter_secure_storage) | Keeping that key in the device's keychain |
| [@modernc-org/sqlite](https://gitlab.com/cznic/sqlite) | Pure-Go SQLite driver, so the server needs no cgo |
| [@benbjohnson/litestream](https://github.com/benbjohnson/litestream) | Continuous SQLite backup (planned) |

### Tooling

| Project | What we use it for |
|---------|-------------------|
| [OpenSpec](https://openspec.dev) · [@openspec](https://github.com/openspec) | Specification-driven development workflow |
| [@jdx/mise](https://github.com/jdx/mise) | Pinning the Flutter version |
| [@caddyserver/caddy](https://github.com/caddyserver/caddy) | Web server for marketing site (planned) |

---

## Tech Stack

| Layer | Technology |
|-------|------------|
| **Mobile app** | Flutter 3.47.4, pinned in `mise.toml` (iOS and Android) |
| **State** | Bloc / Cubit |
| **Networking** | Dio and the Nah? server's own API |
| **Backend** | Nah? Home: one static Go binary, `net/http`, no framework |
| **Database** | SQLite, one file per person; streaming backups with Litestream are planned |
| **Storage** | Local disk or S3-compatible for photos and voice, ciphertext either way (planned) |

### Stack evolution

| Date | Decision | Why |
|---|---|---|
| 2026-01 | ~~Svelte 5 + SvelteKit (PWA), Tailwind v4, shadcn-svelte, TanStack Query~~ | Act small, ship a PWA MVP. Mobile-first web. |
| 2026-05-19 | Flutter 3.41+ / Dart, Bloc, Dio, Hive/Drift, Mastodon backend unchanged | Pivot to native via Flutter. See [ADR-0001: Flutter over PWA](docs/decisions/0001-flutter-over-pwa.md) and the [Flutter decision blog post](https://caseyro.github.io/Nah/2026/05/19/the-flutter-decision/). |
| 2026-09-13 | ~~Mastodon fork~~ → one Go binary on PocketBase with SQLite | Circles became the core data model, and nothing off the shelf has them. See [ADR-0009](docs/decisions/0009-circles-not-one-circle.md), [ADR-0010](docs/decisions/0010-small-server-not-mastodon.md) and the [circles research](docs/research/circles-and-the-backend.md). |
| 2026-09-13 | ~~PocketBase~~ → plain Go, one SQLite file per circle | Measured rather than argued: PocketBase peaked at 705 MB against plain Go's 73.5 MB, unexplained by three experiments. See [ADR-0011](docs/decisions/0011-plain-go-and-a-database-per-circle.md), [ADR-0014](docs/decisions/0014-the-go-setup.md) and the [spike results](https://caseyro.github.io/Nah/research/spike-results/). |
| 2026-09-13 | Content encrypted on the device; sign-in is an Ed25519 key, not a password | The server has no use for plaintext, so it should not hold any. See [ADR-0012](docs/decisions/0012-encrypted-on-device.md) and [ADR-0015](docs/decisions/0015-no-passwords-a-key-on-the-device.md). |
| 2026-09-14 | ~~One SQLite file per circle~~ → one per person, the feed read newest poster first | Circles gave way to one network of 150 per person, which made the feed a fan-in. Measured before it was built. See [ADR-0017](docs/decisions/0017-one-network-of-a-hundred-and-fifty.md), the amended [ADR-0011](docs/decisions/0011-plain-go-and-a-database-per-circle.md) and the [spike results](https://caseyro.github.io/Nah/research/spike-results/). |

The PWA framing is preserved in git history (the `pwa-shell` spec, deleted at 655cfb7), and [CAP_EVOLUTION.md](openspec/changes/CAP_EVOLUTION.md) records how every capability moved, so the evolution is visible, not erased.

How the code fits together today is in the [codebase wiki](docs/wiki/quickstart.md). The product principles and every capability are in [OpenSpec](openspec/changes/), starting with [nah-vision](openspec/changes/cdi-1881-nah-vision/).

---

## Project Structure

```text
/apps
  /mobile           # Flutter app for iOS and Android
  /server           # Nah? Home: Go and SQLite, one database per person
/docs
  /_posts           # Build-in-public blog posts
  /decisions        # Architecture decision records
  /research         # The research behind those decisions
  /wiki             # Codebase wiki, not published on the site
/spike              # Measurements that shaped the server; none of it ships
/openspec
  /changes          # One change per capability, named after its Linear issue
    /cdi-1881-nah-vision   # The product principles
    /cdi-*          # Capability changes
```

---

## Documentation

- **[Codebase wiki](docs/wiki/quickstart.md)**: how the code fits together, and how to run it
- **[Decisions](docs/decisions/)**: architecture decision records
- **[Research](docs/research/)**: the concept audit, the retention investigation and the spike results
- **[Vision & Specs](openspec/changes/)**: the product principles and every capability
- **[Reflections](https://caseyro.github.io/Nah/)**: build-in-public blog

---

## Status

🚧 **Pre-alpha.** The M1 walking skeleton runs on one machine: two people join by invitation, connect, post text moments and read each other's feed. Next is the same on two real phones against a deployed server.

### Progress

- [x] Vision proposal
- [x] Technical design
- [x] Feature specifications
- [x] Project scaffolding
- [ ] Walking skeleton on two phones (M1)
- [ ] Alpha release

---

## Development

### Prerequisites

- [mise](https://mise.jdx.dev), which installs the pinned Flutter: run `mise install` in the repository
- Go 1.27 or newer, for `apps/server`
- Xcode for iOS builds and the Android SDK for Android builds; `flutter doctor` lists what is missing
- [pre-commit](https://pre-commit.com/), for the hooks below
- Ruby 3.3 and Bundler, only to preview the site in `docs/`
- Docker, only to build the server's container image

How to run the server and the app is in the [codebase wiki](docs/wiki/quickstart.md).

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

Nah? is open source and community-driven. We're not ready for contributions yet, but we will be soon.

**Watch this repo** to follow along as we build.

---

## License

[AGPL-3.0-or-later](LICENSE), with one additional term in [NOTICE](NOTICE): every copy and every fork keeps the attribution "Based on Nah?, <https://github.com/CaseyRo/Nah>", so it can always be traced back to the original.

---

## Support

Nah? is community-funded. No ads, no data sales, no VC.

Support options coming soon via Open Collective and GitHub Sponsors.

---

*"Nah? is a private home for your closest people, to share life without performing for the internet."*
