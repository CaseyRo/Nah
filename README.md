# Nah

**A private home for your closest people.**

> *"Share with your circle, not the world."*

Nah ("Not Alone Here") is a private social network limited to 150 friends — inspired by [Path](https://en.wikipedia.org/wiki/Path_(social_network)), built on modern open-source infrastructure, and designed to never compromise on intimacy.

**Viral? Nah. Vital.**

---

## Core Principles

- **Small by design** — 150 friends max (Dunbar's number)
- **Private by default** — no public timelines, no algorithmic feeds
- **Real friends, real moments** — mutual friendships, rich moment sharing
- **No algorithmic theater** — chronological feed only, no engagement optimization
- **Open source** — community-funded, transparent, yours to verify

---

## Tech Stack

| Layer | Technology |
|-------|------------|
| **Frontend** | Svelte 5 + SvelteKit (PWA) |
| **UI** | Tailwind CSS v4 + shadcn-svelte |
| **Backend** | Mastodon (vanilla fork, lightly modified) |
| **Database** | PostgreSQL + Redis |
| **Storage** | S3-compatible + CDN |

See [design.md](openspec/changes/nah-vision/design.md) for full architecture details.

---

## Project Structure

```
/apps
  /web              # SvelteKit PWA (coming soon)
  /server           # Mastodon fork (coming soon)
/packages
  /ui               # Nah design system (coming soon)
/docs
  /_reflections     # Build-in-public blog posts
/openspec
  /changes          # Feature specifications
    /nah-vision     # Core vision (complete)
    /cap-01-*       # Capability implementations (scaffolded)
```

---

## Documentation

- **[Vision & Specs](openspec/changes/nah-vision/)** — Why, what, and how
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
