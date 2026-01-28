# Nah Technical Design

## Context

Nah is a greenfield private social network inspired by Path. We're starting from zero with the goal of creating an intimate, community-owned space for close friends. The technical approach must:

1. Enable rapid development without reinventing social networking primitives
2. Ensure privacy and data ownership are architectural guarantees, not afterthoughts
3. Support the "coming home" UX vision with modern, responsive interfaces
4. Remain sustainable through community funding, not VC-driven growth

**Current state:** No existing codebase. We're defining the foundational architecture.

**Key constraints:**

- Must be open-source (transparency, trust, community contribution)
- Must support 150-friend limit as a first-class concept
- Must work well on mobile (PWA minimum, native apps later)
- Must be deployable on modest infrastructure (community-funded)

## Goals / Non-Goals

**Goals:**

- Establish a clear, layered architecture (backend / API / frontend)
- Choose proven, maintainable technologies over bleeding-edge
- Design for privacy-by-default at every layer
- Enable the Path-inspired UX (radial menu, rich moments, warm aesthetic)
- Support future extensibility (federation, native apps) without requiring it now

**Non-Goals:**

- Full technical specifications for each feature (that's what specs are for)
- UI/UX design details (separate design system work)
- Deployment/DevOps specifics (infrastructure planning comes later)
- Token economics or crypto integration (out of scope for v1)
- Native mobile apps (PWA first, native considered for v2+)

## Decisions

### 1. Backend: Vanilla Mastodon Fork

**Decision:** Fork vanilla Mastodon directly as the backend foundation

**Rationale:**

- Mastodon provides battle-tested social networking primitives: accounts, posts, timelines, notifications, media handling, OAuth2
- Ruby on Rails + PostgreSQL + Redis is a mature, well-understood stack
- **Best maintenance path** — security patches flow from the largest community
- Cleaner codebase than Hometown/Glitch forks — less complexity to understand
- ActivityPub protocol gives us federation capabilities if ever needed later

**Why not Hometown or Glitch?**

- Hometown's "local-only posting" is redundant when federation is disabled — all posts are already local
- Hometown has maintenance concerns (sole maintainer seeking to hand off, 9+ months behind upstream)
- Glitch adds complexity we don't need ("kitchen sink" of features)
- Mastodon upstream is philosophically opposed to private features, but we don't need their buy-in — we're forking

**Modifications to Mastodon:**

- Disable federation completely (closed instance mode)
- Disable/hide local timeline and public timeline (privacy enforcement)
- Default all posts to "followers-only" visibility
- Enforce 150-follower limit at the database/API level
- Convert follow system to mutual friendship model (both must accept)
- Add custom fields for moment types (music, location, sleep/wake)
- Extend reaction system beyond simple favorites

### 2. Frontend: Svelte 5 + SvelteKit PWA

**Decision:** Svelte 5 + SvelteKit as the frontend framework

**Rationale:**

- **Mobile-first PWA** that feels like a native app without app store gatekeeping
- SvelteKit provides app router, SSR for fast first paint, and excellent PWA support
- Svelte's reactivity model is simpler and more performant than React's virtual DOM
- Smaller bundle sizes = faster load times on mobile
- Service workers via SvelteKit adapter for offline read, fast re-open, background sync

**Alternatives considered:**

- *React*: Larger ecosystem, but heavier bundles, more boilerplate, virtual DOM overhead
- *Vue.js*: Good option, but Svelte's compiler approach is more aligned with performance goals
- *React Native / Flutter*: True native, but doubles development effort and delays launch

### 3. UI System: Tailwind CSS v4 + shadcn-svelte

**Decision:** Tailwind CSS v4 for styling, shadcn-svelte for component primitives

**Rationale:**

- **Tailwind CSS v4** provides utility classes + design tokens for consistent styling
- **shadcn-svelte** is a code-ownership component generator, not a locked dependency
- Components are copied into your repo — you own and modify them freely
- Perfect for Nah's "not too abstract, keep it fun" goal
- Built on Bits UI + Tailwind, providing accessible primitives we can customize
- Enables Path-style playful UI without fighting a rigid component library

**Component ownership model:**

- Generate base components from shadcn-svelte
- Customize heavily for Nah's warm, Path-inspired aesthetic
- Store in `/packages/ui` as Nah's design system

### 4. State Management & Data Fetching

**Decision:** TanStack Query (Svelte adapter) + SvelteKit's built-in load functions

**Rationale:**

- TanStack Query handles caching, background refetching, and optimistic updates
- SvelteKit's `load` functions provide server-side data fetching with streaming
- Svelte stores for local UI state
- IndexedDB (via `idb-keyval`) for offline timeline caching

**Data flow:**

- Mastodon REST API for CRUD operations
- Mastodon Streaming API (WebSocket) for real-time timeline updates
- Local IndexedDB cache for offline viewing and fast re-open

### 5. Data Privacy Architecture

**Decision:** Privacy-first by architecture, not policy

**How privacy works (the X/Y/Z model):**

```text
X (user) ←—friends—→ Y (user) ←—friends—→ Z (user)

X sees: Y's posts only
Y sees: X's posts + Z's posts
Z sees: Y's posts only
X cannot see Z (they're not friends)
```

All posts default to **followers-only** visibility. The API layer enforces access control — users only receive posts from people they follow. The database stores everything, but unauthorized content is never returned.

**Implementation:**

- Federation completely disabled (single closed instance)
- Local timeline and public timeline hidden/disabled in UI
- All posts default to "followers-only" visibility (server-side default)
- All accounts locked by default (follow = friend request requiring approval)
- Custom API middleware to enforce 150-friend limit
- "Who viewed" tracking stored locally, never exposed publicly
- No analytics, no tracking, no third-party scripts
- User data export always available (GDPR-style)

**Transparent privacy policy:**

- Simple, human-readable privacy policy (not legal jargon)
- Clearly state: admins can technically access database (standard for any hosted service)
- No data sales, no ads, no third-party sharing
- Users own their data — export anytime, delete anytime

### 6. Friend Limit Enforcement

**Decision:** Hard limit of 150, enforced at API level

**Rationale:**

- Path's mistake was raising limits under growth pressure
- Making it a database constraint, not a UI suggestion, ensures it can't be "accidentally" changed
- 150 aligns with Dunbar's number — the cognitive limit for meaningful relationships
- Smaller networks mean less moderation burden, more trust

**Implementation:**

- PostgreSQL constraint on follow count
- API returns error if user attempts to add 151st friend
- UI shows "X/150 friends" counter as positive framing, not limitation
- Optional "inner circle" (subset of friends) for selective sharing

### 7. Moment Types & Rich Content

**Decision:** Extend Hometown posts with structured metadata

**Rationale:**

- Path's magic was in rich moment types: music, location, sleep/wake, "with" tagging
- Mastodon posts already support attachments and custom fields
- We can store moment metadata in post JSON and render it specially in our frontend

**Moment types (v1):**

- **Photo/Video**: Standard media attachments (client-side resize/compression before upload)
- **Text**: Simple status updates
- **Music**: Song/album info (Spotify/Apple Music integration or manual entry)
- **Location**: Venue/city with optional map (privacy-conscious — user controls precision)
- **Sleep/Wake**: Status indicators with timestamps
- **With**: Tag friends who are present

**Technical approach:**

- Store moment type and metadata in structured text patterns or ActivityPub extensions
- Frontend parses and renders specialized UI per type
- Backend validates structure but doesn't need special handling

### 8. Reactions System

**Decision:** Custom reaction types stored as specialized favorites

**Rationale:**

- Path's emoji reactions (smile, frown, gasp, laugh) were warmer than "likes"
- Mastodon only has "favorite" — we need to extend this

**Implementation:**

- Custom emoji reactions stored in a separate table linked to post
- Cleanest separation, doesn't pollute reply threads
- Frontend aggregates and displays reactions with friend avatars

### 9. Messaging (Path Talk)

**Decision:** Defer to v1.5, use Mastodon DMs initially

**Rationale:**

- Mastodon has basic DM support (direct visibility posts)
- Path Talk's ephemeral messaging and ambient status are valuable but complex
- Better to ship core experience first, then enhance messaging

**Future direction:**

- Matrix protocol integration for E2EE messaging
- Ephemeral (24-hour) message option
- Ambient status sharing (listening to, in transit, etc.)

### 10. Repository Structure

**Decision:** Monorepo with pnpm workspaces

**Structure:**

```text
/apps/web          # SvelteKit PWA (main client)
/apps/server       # Mastodon fork + Nah-specific config
/packages/ui       # Nah design system (shadcn-svelte generated + customized)
```

**Rationale:**

- Single repo simplifies development and deployment
- Shared UI package enables consistency
- pnpm workspaces provide efficient dependency management
- TypeScript everywhere (frontend strongly typed; backend remains Ruby)

### 11. Observability & Operations

**Decision:** Standard open-source observability stack

**Components:**

- **Sentry**: Error tracking for frontend + backend
- **Prometheus + Grafana**: Metrics and dashboards for Hometown
- **Structured logging**: JSON logs, PII excluded

**Infrastructure (v0 → v1):**

- Docker Compose for local dev and initial deployment
- Single VM + managed Postgres to start
- Nightly Postgres dumps + media bucket versioning for backups
- One-click user data export (trust feature)

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| **Mastodon fork maintenance burden** | Stay close to upstream, only modify what's necessary. Our changes are surgical (config + friend limits + reactions). Rebase regularly on Mastodon releases. |
| **Single instance = single point of failure** | Standard infrastructure practices: backups, monitoring, redundancy. Federation architecture exists if we ever need to scale horizontally. |
| **Svelte ecosystem smaller than React** | Svelte 5 is mature. shadcn-svelte provides solid component foundation. TanStack Query has Svelte adapter. |
| **150-limit frustrates users** | Frame positively ("your closest 150"), provide tools to curate (inactive friend suggestions). |
| **Competing with nostalgia** | Focus on the *feeling* Path created, not pixel-perfect recreation. Modernize where appropriate. |
| **Community funding sustainability** | Transparent finances, clear value proposition, optional supporter perks. Start lean. |
| **No native apps hurts adoption** | PWA provides 80% of native experience. Native wrapper (Capacitor) possible later. |
| **Admin database access** | Transparent privacy policy. Standard for any hosted service. E2EE for DMs considered for v2. |

## Open Questions

1. **Music integration**: Direct Spotify/Apple Music API integration, or manual song entry only? (Privacy vs. convenience trade-off)
2. **Location precision**: How granular should location sharing be? City-level default with optional venue?
3. **Onboarding**: How do users find their first friends without discoverability? Contact import? Invite codes?
4. **Moderation**: In a 150-person network, who moderates? Is community self-policing sufficient?
5. **Media storage costs**: S3 + CDN adds up. Offer reduced media quality for cost savings? Limit per-user storage?
6. **Storybook**: Add Storybook for Svelte once design system stabilizes, or defer?

---

## Architecture Summary

```text
┌─────────────────────────────────────────────────────────────┐
│                        FRONTEND                              │
│  Svelte 5 + SvelteKit (PWA)                                 │
│  ├── Tailwind CSS v4 + shadcn-svelte                        │
│  ├── TanStack Query (caching, optimistic updates)           │
│  └── IndexedDB (offline timeline)                           │
└─────────────────────────────────────────────────────────────┘
                              │
                    REST API + Streaming
                              │
┌─────────────────────────────────────────────────────────────┐
│                        BACKEND                               │
│  Mastodon (vanilla fork, lightly modified)                  │
│  ├── Ruby on Rails                                          │
│  ├── PostgreSQL (accounts, posts, relationships)            │
│  ├── Redis (caching, background jobs)                       │
│  ├── Federation OFF, local/public timelines hidden          │
│  └── All posts default to followers-only                    │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                        STORAGE                               │
│  S3-compatible object storage + CDN                         │
│  (photos, videos, media)                                    │
└─────────────────────────────────────────────────────────────┘
```

---

*This design establishes the technical foundation. Detailed specifications for each capability follow in the specs artifacts.*
