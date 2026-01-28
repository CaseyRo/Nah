# Nah Vision Proposal

## Why

Social media has become performance theater. We broadcast to hundreds of "friends" we barely know, curate ourselves for algorithmic approval, and measure worth in viral moments. The result: loneliness dressed as connection.

Path showed us another way. From 2010-2018, it proved people *want* intimate digital spaces — a place to share life's real moments with their actual close circle. Path reached 50M users and turned down a $100M Google offer. But it chased growth, diluted its intimacy (raising friend limits from 50 to 500), and eventually died in Facebook's shadow.

**Nah** ("Not Alone Here") is the spiritual successor to Path — built on its lessons, powered by modern open-source infrastructure, and designed to never compromise on intimacy. The opportunity is now: social media fatigue is at an all-time high, and people are hungry for spaces that feel like home, not a stage.

**Viral? Nah. Vital.**

## What Changes

This is a greenfield project. We're creating a new social network from scratch with these foundational principles:

- **Small by design**: Hard cap of 150 friends (Dunbar's number) — the cognitive limit for stable relationships
- **Private by default**: No public timeline, no discoverability, no algorithmic feed
- **Real friends, real moments**: Mutual friendship model (both must accept), rich moment sharing (photos, music, location, sleep/wake)
- **No algorithmic theater**: Chronological feed only, no engagement optimization, no ads
- **Open-source & community-funded**: Transparent code, donation-supported, no data mining
- **Modern infrastructure**: Mastodon/ActivityPub backend with custom Path-inspired frontend

## Capabilities

### New Capabilities

- `core-identity`: App identity, branding, and design language — the "coming home" aesthetic inspired by Path's warm, polished UI
- `friend-circles`: Dunbar-limited connections (150 max), mutual friendship model, inner circle support for selective sharing
- `moments`: Timeline of life moments — photos, text, music, location, sleep/wake status, auto-location check-ins
- `reactions`: Expressive emoji reactions (not just "like"), read receipts, "who viewed" transparency
- `messaging`: Private 1:1 and small group chat with optional ephemeral messages (Path Talk inspired)
- `ambient-presence`: Opt-in status sharing — what you're listening to, battery level, transit status — subtle cues that deepen presence
- `user-ownership`: Data export, account portability, transparent privacy controls
- `community-funding`: Donation model, optional supporter perks (badges, themes), crowdfunded feature roadmap

### Modified Capabilities

*None — this is a new project.*

## Impact

### Technical Foundation
- Backend: Mastodon instance (Ruby on Rails, PostgreSQL, Redis) configured for closed/private operation
- Frontend: Modern web app (React/Vue/Svelte) as PWA, mobile-first design
- Storage: S3-compatible media storage with CDN
- Protocol: ActivityPub foundation (federation disabled initially, but architecture supports future expansion)

### Design Direction
- Path's signature radial "+" menu for posting
- Warm color palette, smooth animations, "homey" feel
- Clock-style timeline, cover photos, profile fade effects
- Playful reaction stickers with friend avatars

### Business Model
- No ads, no data sales
- Community donations (Patreon, Open Collective, in-app)
- Optional digital goods (sticker packs, themes) with proceeds to project
- Transparent financial reporting

### Philosophical Commitments
- Never raise the friend limit above 150
- Never add public/viral features
- Never optimize for "engagement" over connection
- Always remain open-source

---

*"Nah is a private home for your closest people — to share life without performing for the internet."*
