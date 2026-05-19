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
- Must feel native on mobile (Flutter cross-platform, single codebase for iOS + Android)
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
- Web/desktop clients (Flutter mobile app first, web/desktop reconsidered for v2+)

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

### 2. Client: Flutter (iOS + Android)

> **Decision changed 2026-05-19** (was: SvelteKit PWA, decided 2026-01). See [ADR-0001 — Flutter over PWA](../../../docs/decisions/0001-flutter-over-pwa.md). The original PWA rationale is preserved at the end of this section so the evolution is visible.

**Decision:** Flutter 3.41+ / Dart 3.11+ as the mobile client framework

**Rationale:**

- **One codebase, two platforms** — single Dart codebase ships to both iOS and Android, the only way a solo project realistically reaches both stores
- **Native-feeling UX** — Flutter's Impeller engine renders Path-style spring physics, the radial menu bloom, and the floating analog timeline clock at 60-120fps
- **Plugin coverage for Nah's signature features** — `flutter_blue_plus` / `universal_ble` for BLE proximity, geolocator for location tags, audio_session/MediaSession for music share-out, biometric/haptic feedback plugins; iOS-specific gaps (iBeacon proximity via CoreLocation, iOS Share Extension) fill with thin Swift platform channels
- **Mature in 2026** — Flutter 3.41 / Dart 3.11 shipped as a stability-focused release; Google committed to four stable releases/year, 18-month LTS, 99.9% SLA; Google Pay, Google Ads, and parts of Assistant run on Flutter
- **Strong typing + hot reload** keeps a small team productive

**Alternatives considered (and why not):**

- *SvelteKit PWA*: original framing for nah-vision. Rejected — iOS Safari blocks BLE, background sync, Web Share Target (inbound), and gates push behind install; EU iOS users lose push entirely under DMA. BLE-based proximity is a Nah USP, not a nice-to-have. See [the Flutter decision blog post](../../../docs/_posts/2026-05-19-the-flutter-decision.md) for the full back-and-forth.
- *Native iOS (Swift) + Android (Kotlin)*: best UX, two codebases. Solo project will not finish two parallel apps.
- *React Native*: viable. Flutter wins on rendering consistency and animation performance for our Path-inspired motion language; JS ecosystem advantage is irrelevant when the team is one Dart-friendly developer.

**Trade-offs accepted:**

- Bundle size: ~30-50MB iOS IPA vs ~5-15MB native Swift, ~1-2MB PWA. Tolerable for a 150-person private network.
- Hiring is harder than RN/Web (6-8 week fill vs 3-4 week). Doesn't bite at solo scale; revisit if the project grows a team.
- Flutter apps can be visually identifiable as Flutter. Mitigated by Nah's bespoke design language — when every animation is custom, "Flutter-feel" disappears.

**~~Prior decision (2026-01-28), kept for evolution context:~~**

> ~~**Decision:** Svelte 5 + SvelteKit as the frontend framework. Mobile-first PWA that feels like a native app without app store gatekeeping. SvelteKit provides app router, SSR for fast first paint, and excellent PWA support. Svelte's reactivity model is simpler than React's virtual DOM. Smaller bundle sizes. Service workers for offline read.~~
>
> ~~**Alternatives considered (and rejected):** React (heavier bundles, more boilerplate); Vue.js (Svelte's compiler approach more aligned with performance goals); React Native / Flutter (doubles development effort and delays launch).~~
>
> Superseded once we confirmed Flutter 3.41 reaches further back than feature-complete PWA on iOS (iOS 13+ vs iOS 16.4+ for push), and once BLE proximity moved from defer-forever to a real Phase 2 product feature.

### 3. UI System: Flutter Material 3 + Cupertino + custom Nah widgets

**Decision:** Build on Flutter's Material 3 + Cupertino widget foundations, layered with Nah's custom design tokens and widgets

**Rationale:**

- **Material 3 / Cupertino** give us accessible primitives (buttons, sheets, navigation) for free, with platform-correct behavior on Android vs iOS
- **Custom Nah widgets** (`RadialMenu`, `MomentCard`, `TimelineClock`, `ReactionPicker`, etc.) live in `/packages/nah_ui` — owned by us, not a third-party theme dependency
- **ThemeData + custom Theme extensions** carry our design tokens (pomegranate red, warm neutrals, Nunito display font, spring curves) into every widget
- Path-style "handcrafted" feel comes from the custom layer, not from fighting a rigid component kit

**Component ownership model:**

- Build atop Flutter's standard widget library
- Implement Nah-branded widgets in `/packages/nah_ui`
- Tokens (colors, typography, spacing, motion) defined once in a `nah_theme.dart` module

### 4. State Management & Data Fetching

**Decision:** Bloc / Cubit for app state + Dio for HTTP + native WebSocket for Mastodon streaming

**Rationale:**

- **Bloc / Cubit** provides predictable unidirectional state flow with explicit events and states — fits well with Mastodon's REST + Streaming model
- **Dio** is the standard mature HTTP client in Dart; interceptors for auth, retry, and offline queueing
- **`web_socket_channel`** for the Mastodon Streaming API
- **Hive** or **Drift (SQLite)** for offline timeline caching and queued posts

**Data flow:**

- Mastodon REST API for CRUD operations (via Dio)
- Mastodon Streaming API (WebSocket) for real-time timeline updates
- Local Hive/Drift cache for offline viewing and fast re-open
- Queued moments persisted locally, flushed when connectivity returns

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

**Decision:** Extend Mastodon posts with structured metadata

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

**Decision:** Monorepo with mixed-language layout — Flutter app in `/apps/mobile`, Mastodon backend in `/apps/server`, Dart packages in `/packages/`

**Structure:**

```text
/apps/mobile        # Flutter app — iOS + Android (main client)
/apps/server        # Mastodon fork + Nah-specific config
/packages/nah_ui    # Nah design system: ThemeData, widgets, tokens
/packages/nah_api   # Mastodon API client wrapper (typed Dart models)
```

**Rationale:**

- Single repo simplifies development and deployment
- Shared Dart packages keep design tokens and API models DRY
- Flutter workspace via `pubspec.yaml` workspaces (Dart 3.5+) or `melos`
- Dart on the client, Ruby on the server — boundaries are clean

### 11. Observability & Operations

**Decision:** Standard open-source observability stack

**Components:**

- **Sentry**: Error tracking for frontend + backend
- **Prometheus + Grafana**: Metrics and dashboards for Mastodon
- **Structured logging**: JSON logs, PII excluded

**Infrastructure (v0 → v1):**

- Docker Compose for local dev and initial deployment
- Single VM + managed Postgres to start
- Nightly Postgres dumps + media bucket versioning for backups
- One-click user data export (trust feature)

### 12. Visual Identity: Light-Primary

**Decision:** Warm light surfaces as primary identity; dark mode is secondary

**Rationale:**

- Path's identity was warm, light, and inviting — dark mode didn't define the brand
- Light surfaces let photos and content pop; pomegranate red accents read best on white
- Dark mode respects system preference but is not the primary design target
- Glassmorphism is scoped to exactly 2 contexts: profile header compression overlay and radial menu backdrop — nowhere else

**Implementation:**

- Default: light theme (`--color-surface: #FFFFFF`, warm neutrals)
- Dark mode: follows `prefers-color-scheme: dark` system setting
- All design work and mockups start with light theme
- Dark mode token architecture defined in design-system spec

### 13. Notification Architecture

**Decision:** Layered notification system — push, in-app badge, and silent

**Push notifications (device-level):**

| Event | Push | Badge | Silent |
|-------|------|-------|--------|
| Friend request received | Yes | Yes | — |
| Friend request accepted | Yes | Yes | — |
| Reaction on your moment | — | Yes | Yes |
| Comment on your moment | Yes | Yes | — |
| Reply to your comment | Yes | Yes | — |
| First moment from new friend | Yes | — | — |
| Message received (v1.5) | Yes | Yes | — |

**In-app badge:**

- Red dot on Friends tab for pending requests
- Count badge on notification bell (if present) or tab
- Cleared when user views the relevant screen

**Notification preferences screen:**

- Master toggle: all push notifications on/off
- Per-category toggles: reactions, comments, friend requests
- Quiet hours: schedule no-push window (e.g., 10pm–8am)

**Deep-link destinations:**

| Notification type | Opens |
|-------------------|-------|
| Friend request | Friends tab → pending requests |
| Reaction | The specific moment in timeline |
| Comment | The specific moment, scrolled to comments |
| Friend accepted | New friend's profile |

### 14. Messaging Scope

**Decision:** Messaging is explicitly deferred to v1.5

Mastodon's basic DM support (direct visibility posts) is available at the infrastructure level but the full Path Talk experience (ephemeral messages, typing indicators, ambient status) is v1.5 scope. The v1 navigation model does not include a Messages tab.

**v1 navigation:** Feed, Friends, Profile (3 tabs)
**v1.5 addition:** Messages tab added as 4th tab

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| **Mastodon fork maintenance burden** | Stay close to upstream, only modify what's necessary. Our changes are surgical (config + friend limits + reactions). Rebase regularly on Mastodon releases. |
| **Single instance = single point of failure** | Standard infrastructure practices: backups, monitoring, redundancy. Federation architecture exists if we ever need to scale horizontally. |
| **Flutter bundle size (~30-50MB iOS)** | Acceptable for a private 150-person network where install friction is low. Optimize with deferred components and asset compression closer to v1. |
| **Flutter hiring harder than RN/Web (6-8w vs 3-4w)** | Solo project initially. Revisit when team grows; Dart is approachable for any strong dev. |
| **150-limit frustrates users** | Frame positively ("your closest 150"), provide tools to curate (inactive friend suggestions). |
| **Competing with nostalgia** | Focus on the *feeling* Path created, not pixel-perfect recreation. Modernize where appropriate. |
| **Community funding sustainability** | Transparent finances, clear value proposition, optional supporter perks. Start lean. |
| **iOS BLE proximity needs platform channels (CoreLocation)** | Accept ~20-50 lines of Swift for iBeacon proximity. Still vastly less than a full native iOS app. |
| **Admin database access** | Transparent privacy policy. Standard for any hosted service. E2EE for DMs considered for v2. |

## Open Questions

1. **Music integration**: Direct Spotify/Apple Music API integration, or manual song entry only? (Privacy vs. convenience trade-off)
2. **Location precision**: How granular should location sharing be? City-level default with optional venue?
3. **Onboarding**: How do users find their first friends without discoverability? Contact import? Invite codes?
4. **Moderation**: In a 150-person network, who moderates? Is community self-policing sufficient?
5. **Media storage costs**: S3 + CDN adds up. Offer reduced media quality for cost savings? Limit per-user storage?
6. **Widget gallery**: Add a Flutter widget catalog (`widgetbook`, `dashbook`, or hand-rolled gallery route) once design system stabilizes, or defer?

---

## Architecture Summary

```text
┌─────────────────────────────────────────────────────────────┐
│                       MOBILE CLIENT                          │
│  Flutter 3.41+ / Dart 3.11+ (iOS + Android)                 │
│  ├── Material 3 + Cupertino + custom Nah widgets            │
│  ├── Bloc / Cubit (state management)                        │
│  ├── Dio (HTTP) + web_socket_channel (streaming)            │
│  ├── Hive / Drift (offline timeline + queued moments)       │
│  └── Platform channels (BLE proximity, iOS Share Extension) │
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
