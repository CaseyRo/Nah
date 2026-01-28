# Nah Vision Tasks

This vision change establishes the foundation. Implementation is broken into 8 capability changes (listed in section 4).

## 1. Project Foundation

- [ ] 1.1 Initialize monorepo with pnpm workspaces (`/apps/web`, `/apps/server`, `/packages/ui`)
- [ ] 1.2 Set up SvelteKit project in `/apps/web` with TypeScript
- [ ] 1.3 Configure Tailwind CSS v4 and shadcn-svelte
- [ ] 1.4 Set up Mastodon fork in `/apps/server` (Docker-based)
- [ ] 1.5 Configure Mastodon for closed instance (federation OFF, followers-only default)
- [ ] 1.6 Set up development environment (Docker Compose for local dev)

## 2. Documentation & Brand

- [ ] 2.1 Finalize privacy policy (plain language, transparent about admin access)
- [ ] 2.2 Create landing page content (micro-manifesto, value proposition)
- [ ] 2.3 Publish GitHub Pages docs site
- [ ] 2.4 Set up community channels (GitHub Discussions, Discord, or similar)

## 3. Infrastructure Scaffolding

- [ ] 3.1 Set up S3-compatible storage for media
- [ ] 3.2 Configure CDN for media delivery
- [ ] 3.3 Set up Sentry for error tracking
- [ ] 3.4 Configure basic monitoring (Prometheus + Grafana)
- [ ] 3.5 Set up CI/CD pipeline (GitHub Actions)

## 4. Capability Implementation (Separate Changes)

Each capability has its own OpenSpec change. Work through in order:

- [ ] 4.1 Complete `cap-01-core-identity` — App identity, branding, design language, radial menu
- [ ] 4.2 Complete `cap-02-friend-circles` — 150-friend limit, mutual friendships, inner circle
- [ ] 4.3 Complete `cap-03-moments` — Timeline, moment types (photo, text, music, location, sleep/wake)
- [ ] 4.4 Complete `cap-04-reactions` — Emoji reactions, read receipts, who-viewed
- [ ] 4.5 Complete `cap-05-messaging` — Private DMs, group chat, ephemeral messages
- [ ] 4.6 Complete `cap-06-ambient-presence` — Opt-in status sharing (music, location, battery)
- [ ] 4.7 Complete `cap-07-user-ownership` — Data export, account deletion, privacy controls
- [ ] 4.8 Complete `cap-08-community-funding` — Donation model, supporter perks, transparent finances

## 5. Launch Preparation

- [ ] 5.1 Internal alpha testing (invite-only)
- [ ] 5.2 Security review of Mastodon modifications
- [ ] 5.3 Load testing for expected initial user base
- [ ] 5.4 Finalize backup and recovery procedures
- [ ] 5.5 Prepare launch communications

## 6. Post-Launch

- [ ] 6.1 Set up community feedback collection
- [ ] 6.2 Establish contribution guidelines for open source
- [ ] 6.3 Create public roadmap
- [ ] 6.4 Set up donation infrastructure (Open Collective, GitHub Sponsors)
