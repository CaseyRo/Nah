# Product

## Register

product

## Users

Adults who've outgrown algorithmic social media and miss what Path felt like in 2012. Mostly iPhone-leaning, design-literate, privacy-aware. They have between 30 and 150 people they actually care about, and no good place to share the ordinary moments of life with that specific circle.

**Context of use:** in bed before sleep, on the couch on a Sunday, during a quiet coffee. Not commuting, not bored in line, not killing time. Nah is opened deliberately, not reflexively.

**Job to be done:** share a real moment with the small group of people who should see it, and see what they're up to, without performing for anyone outside that circle.

## Product Purpose

Nah ("Not Alone Here") is a private social network capped at 150 friends (Dunbar's number), the spiritual successor to Path. It exists because every other social network treats connection as engagement and engagement as attention to be sold. Nah's whole architecture refuses that trade: federation off, public timeline hidden, mutual friendships, chronological feed, no algorithm, no ads, no data sale, open source, community-funded.

**Success looks like:** 150-person networks where people post the kind of small daily moments they'd never put on Instagram, because there's nobody to perform for. Sustained use measured in years, not months. Quiet, present, returning. The success metric is time given back to relationships that matter, not engagement captured inside the app.

## Brand Personality

Warm and quiet. Deliberate, not frictionless. Private by architecture, not by promise. Slow as a positive: the opposite of fast, infinite, optimized. Honest about its limits (150 friends, no public timeline, no virality, ever). Alive in the small craft details: a floating analog clock that hands match scrolled timestamps; custom illustrated reactions instead of system emoji; spring-physics animations that say someone touched every interaction.

If the three-word frame helps: **warm, deliberate, vital**. Or **intimate, slow, honest**. Both are right.

## Anti-references

What Nah is not, in concrete terms:

- **Algorithmic engagement feeds (Instagram, TikTok, X).** No Stories rail at the top. No interleaved suggestions. No "you might also like." No engagement-optimized ranking. Chronological, full stop.
- **Broadcast or performance social (LinkedIn, Threads).** No follower counts. No public profiles. No professional self-curation. The whole model is mutual: every connection requires both sides to agree.
- **Notification-noisy chat apps (Discord, Slack).** No aggressive unread counts. No "X is typing." Reactions surface as badges, not push. Notifications are restrained by default and tunable.
- **VC-SaaS landing-page aesthetic.** No gradient hero text, no big sans-serif growth claims, no three-feature card grids, no glassmorphism dashboards, no hero-metric templates. The product surfaces use light, warm, restrained color; the marketing site is build-in-public and reflective, not a pitch deck.

## Design Principles

1. **Viral? Nah. Vital.** The tagline is the test. If a design decision optimizes for reach, growth, engagement, or session length, it is the wrong decision. Always optimize for the quality of connection inside the 150-person circle.

2. **Privacy by architecture, not policy.** Guarantees come from the code: federation disabled, followers-only by default, 150-cap enforced at the database, "who viewed" stored locally and never exposed publicly. We never ship a "trust us" pattern when an architectural pattern exists.

3. **Deliberate sharing, not frictionless.** Friction is a feature. Pre-post pauses, no infinite scroll, no autoplay, no algorithmic surfacing, restrained notifications. The OneSec lineage: make the user choose, every time.

4. **No counts anywhere.** *(Added 2026-04-26; the MVP headline. See [ADR-0004](docs/decisions/0004-no-counts-anywhere.md).)* No like count, no view count, no follower count, no read count. Where we need to communicate "someone did this," we use names and faces: *Maya reacted*, never *3 people reacted*. The 150 soft cap is communicated as language ("your circle is full"), never as a meter. Counts are the architecture of performance, and Nah? refuses performance.

5. **Design that says someone cared.** Path's signature was craft. Every animation, every empty state, every error message earns its space. If a screen could be replaced by stock SaaS components without the user noticing, the screen has failed. Shortcuts in this layer are felt, even when users can't name them.

## Accessibility & Inclusion

- **Target:** WCAG 2.1 AA across the app. Color contrast verified for both light and dark themes.
- **Motion:** Full `prefers-reduced-motion` support. Spring animations become instant transitions or simple opacity fades. The floating timeline clock becomes a static digital time display. Pull-to-refresh becomes a simple spinner. Reaction bubbles appear without bounce.
- **Touch:** All interactive elements meet 44pt iOS / 48dp Android minimum touch targets. Radial menu items at 44px minimum. FAB at 56px with adequate inset from screen edges.
- **Hand preference:** Left-hand mode setting moves the FAB to bottom-left, mirroring the radial menu accordingly.
- **Platform a11y:** Semantic widgets on both platforms (VoiceOver on iOS, TalkBack on Android). Moment cards use `Semantics(role: article)` patterns. Radial menu uses `role: menu` with `menuitem` children. Tab bar uses `role: tablist` with `role: tab` children.
- **Haptics:** Used sparingly and only as feedback for deliberate actions (post confirmation, friction-gate completion). Not for arbitrary decoration. iOS limited to single haptic taps per platform reality (no arbitrary pattern support).
- **EU iOS DMA reality:** Push notifications gracefully missing for EU iOS users. Educated once on first run; never blocking. The product remains fully functional without push.
