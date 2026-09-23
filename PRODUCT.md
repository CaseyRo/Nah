# Product

## Register

product

## Platform

Native app for iOS and Android, built with Flutter. Design in the platform's conventions: the app should feel native on each, not translated.

## Users

Adults who've outgrown algorithmic social media and miss what Path felt like in 2012. Mostly iPhone-leaning, design-literate, privacy-aware. They have between 30 and 150 people they actually care about, and no good place to share the ordinary moments of life with that specific circle.

**Context of use:** in bed before sleep, on the couch on a Sunday, during a quiet coffee. Not commuting, not bored in line, not killing time. Nah? is opened deliberately, not reflexively.

**Job to be done:** share a real moment with the small group of people who should see it, and see what they're up to, without performing for anyone outside that circle.

## Product Purpose

Nah? ("Not Alone Here") is a private social network for the people you are actually close to, the spiritual successor to Path. Each person has one circle of up to 150 people (Dunbar's number), joined by invitation only and connected in person first. It exists because every other social network treats connection as engagement and engagement as attention to be sold. Nah?'s whole architecture refuses that trade: no public anything, no federation, mutual connections, one chronological feed, no algorithm, no ads, no data sale, moments sealed on the phone so the server cannot read them, open source, community-funded by donations.

**Success looks like:** circles where people post the kind of small daily moments they'd never put on Instagram, because there's nobody to perform for. Sustained use measured in years, not months. Quiet, present, returning. The success metric is time given back to relationships that matter, not engagement captured inside the app.

**Where the model stands (2026-09-23):** one circle per person, drawn around that person alone; nobody shares or enters it, and there are no groups or rooms ([ADR-0017](docs/decisions/0017-one-network-of-a-hundred-and-fifty.md), which superseded the many-circles model of ADR-0009 the same day). The product calls it *your circle*. Every person is known by the name they gave, and may show their gender, pronouns and sexual orientation to their circle, or rather not tell ([CDI-1895](https://linear.app/cdit/issue/CDI-1895), 2026-09-23). Four moment types: text, voice, photo and music, and a text can be a full written post with a title, because people were reflective on Path. Reactions exist again, seen only by the poster, by name and face. No chat, no AI, no counts. One phone per person; a new phone comes through two people in the circle vouching. The rules people live by are on the site's [Rules of engagement](docs/_tabs/rules-of-engagement.md), and the specs are in `openspec/changes/`.

## Brand Personality

Warm and quiet. Deliberate, not frictionless. Private by architecture, not by promise. Slow as a positive: the opposite of fast, infinite, optimized. Honest about its limits (one circle of up to 150, nothing public, no virality, ever). Alive in the small craft details: a floating analog clock whose hands match the moment scrolling past; Nah?'s own illustrated reactions instead of system emoji; spring-physics motion that says someone touched every interaction.

If the three-word frame helps: **warm, deliberate, vital**. Or **intimate, slow, honest**. Both are right.

## Anti-references

What Nah? is not, in concrete terms:

- **Algorithmic engagement feeds (Instagram, TikTok, X).** No Stories rail at the top. No interleaved suggestions. No "you might also like." No engagement-optimized ranking. Chronological, full stop, and the page ends.
- **Broadcast or performance social (LinkedIn, Threads).** No follower counts. No public profiles. No professional self-curation. The whole model is mutual: every connection exists for both people or for neither.
- **Notification-noisy chat apps (Discord, Slack, WhatsApp).** No chat at all. No badges anywhere, on any screen or on the icon. One daily digest by default, nothing on a quiet day. Reactions reach only the poster, in that digest.
- **VC-SaaS landing-page aesthetic.** No gradient hero text, no big sans-serif growth claims, no three-feature card grids, no glassmorphism dashboards, no hero-metric templates. The product surfaces use light, warm, restrained color; the marketing site is build-in-public and reflective, not a pitch deck.

## Design Principles

1. **Viral? Nah. Vital.** The tagline is the test. If a design decision optimizes for reach, growth, engagement, or session length, it is the wrong decision. Always optimize for the quality of connection inside a person's circle.

2. **Privacy by architecture, not policy.** Guarantees come from the code: moments sealed on the phone with keys the project never holds, the 150 enforced by the server, joining only by invitation, no record of who saw a moment, no analytics or reporting software in the app. We never ship a "trust us" pattern when an architectural pattern exists.

3. **Deliberate sharing, not frictionless.** Friction is a feature. Connecting in person, no infinite scroll, no autoplay, no algorithmic surfacing, restrained notifications, no drafts. The OneSec lineage: make the user choose, every time.

4. **No counts anywhere.** *(Added 2026-04-26; the MVP headline. See [ADR-0004](docs/decisions/0004-no-counts-anywhere.md).)* No like count, no view count, no connection count, no read count, no reaction count. Where we need to communicate "someone did this," we use names and faces: *Maya reacted*, never *3 people reacted*. A full circle is said in words ("Maya's circle is full"), never as a meter, and nobody is shown how many people they have. Counts are the architecture of performance, and Nah? refuses performance.

5. **Design that says someone cared.** Path's signature was craft. Every animation, every empty state, every error message earns its space. If a screen could be replaced by stock SaaS components without the user noticing, the screen has failed. Shortcuts in this layer are felt, even when users can't name them.

## Accessibility & Inclusion

- **Target:** WCAG 2.1 AA across the app. Color contrast verified for both light and dark themes. No text is set on Pomegranate: a primary action is a Pomegranate circle carrying a glyph at 3:1 or better, its label beside it.
- **Motion:** Full reduce-motion support. Spring animations become instant transitions or simple opacity fades. The floating timeline clock becomes a static digital time display. Pull-to-refresh becomes a simple spinner. The + menu's items appear together, with no stagger, rotation or overshoot.
- **Touch:** All interactive elements meet 44pt iOS / 48dp Android minimum touch targets. Radial menu items at 44px minimum. The + at 56px with adequate inset from screen edges.
- **Hand preference:** Left-hand mode setting moves the + to bottom-left, mirroring the radial menu accordingly.
- **Platform a11y:** Semantic widgets on both platforms (VoiceOver on iOS, TalkBack on Android). Every control announced with a name and a role, in visual order. Moments use `Semantics(role: article)` patterns. The radial menu uses `role: menu` with `menuitem` children. There is no tab bar. Text scales with the system text size.
- **Haptics:** Used sparingly and only as feedback for deliberate actions (post confirmation, arriving after the first run). Not for arbitrary decoration. iOS limited to single haptic taps per platform reality (no arbitrary pattern support).
- **Without push:** a person who declines notifications, or whose phone cannot receive them, is never blocked; the product remains fully functional without push. (Corrected 2026-09-23: this line once said EU iPhones lose push under the DMA. Native apps were never affected, and Apple reversed the change for web apps on 2024-03-01.)
