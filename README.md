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

## Built With

Nah stands on the shoulders of giants. We're grateful to these open source projects:

| Project | What we use it for |
|---------|-------------------|
| [@mastodon/mastodon](https://github.com/mastodon/mastodon) | Backend social networking engine (forked) |
| [@sveltejs/svelte](https://github.com/sveltejs/svelte) | Frontend framework (Svelte 5) |
| [@sveltejs/kit](https://github.com/sveltejs/kit) | App framework, routing, SSR |
| [@huntabyte/shadcn-svelte](https://github.com/huntabyte/shadcn-svelte) | UI component primitives |
| [@tailwindlabs/tailwindcss](https://github.com/tailwindlabs/tailwindcss) | Utility-first CSS |
| [@TanStack/query](https://github.com/TanStack/query) | Data fetching & caching |

### Tooling

| Project | What we use it for |
|---------|-------------------|
| [OpenSpec](https://openspec.dev) · [@openspec](https://github.com/openspec) | Specification-driven development workflow |
| [@caddyserver/caddy](https://github.com/caddyserver/caddy) | Web server (planned) |
| [@getsentry/sentry](https://github.com/getsentry/sentry) | Error tracking |

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

```text
/apps
  /web              # SvelteKit PWA (coming soon)
  /server           # Mastodon fork (coming soon)
/packages
  /ui               # Nah design system (coming soon)
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

- Node.js 20+
- pnpm
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
