# Capability evolution

How the capability changes have moved since they were first written on 2026-01-28, and where each stands now.

Scope moved twice: on 2026-04-26 with the MVP cut ([ADR-0008](../../docs/decisions/0008-mvp-scope.md)), and on 2026-09-13 with a server of our own ([ADR-0010](../../docs/decisions/0010-small-server-not-mastodon.md)) and one network of up to 150 per person instead of shared circles ([ADR-0017](../../docs/decisions/0017-one-network-of-a-hundred-and-fifty.md)). Since 2026-09-15 the product calls that network *your circle*: one per person, never a group or a room. From 2026-09-15 the changes are restated one at a time against where Nah? stands, detailing specs and building nothing. A restated change passes `openspec validate --strict`. A change not yet restated keeps its January scope beside it as `from-nah-vision.md`, outside `specs/`.

## Status

| Change | January scope | Now | Validates |
|---|---|---|---|
| **nah-vision** | A Path successor on a Mastodon fork, with nine capability specs | **Restated 2026-09-15.** Principles only: one spec, `product-principles`, which every capability keeps. The January specs moved to their caps, and design.md is now a map to the decision records and the wiki. | Yes |
| **cap-01-core-identity** | App identity, design system, radial menu, PWA shell, brand assets | **Restated 2026-09-15.** `DESIGN.md` holds every value. No tab bar. The + offers text, voice, photo and music. No text is set on Pomegranate. Fonts ship in the app, and no reporting SDK is allowed. `pwa-shell` was deleted and is in git history at 655cfb7. Onboarding moved to cap-02; cards, the clock and the offline queue to cap-03; reaction icons to cap-04; "no numbers about people" to nah-vision. | Yes |
| **cap-02-your-circle** | 150-friend cap, friend requests, an inner circle, invite links, curation suggestions | **Restated 2026-09-15**, renamed from `cap-02-friend-circles`. Joining by invitation, with ADR-0005's first run and only real moments. The touch first, with a scanned code as fallback, and the link second. Every invitation works once, never expires, and can be withdrawn. A full circle is named, never numbered. Ending a connection is silent. Friend requests, the inner circle, the "X/150" ring and curation suggestions are gone. | Yes |
| **cap-03-moments** | Photo and video, text, music, location, sleep and wake, tagging, a journal on every profile | **Restated 2026-09-15.** Text, voice, photo and music; a song is shared from any music app and its links are looked up on the phone, and ADR-0016 carries a note for them. An optional short line on photo and voice, and an optional place on any moment, from the phone's place lookup or typed by hand, with no location kept; ADR-0006 and ADR-0008 carry notes. No editing, and deleting leaves a marker. Cards, the timeline clock, honest placeholders, a plain end to the feed, and a page of one person's moments open to their circle. Video, sleep and wake, tagging and maps are gone. | Yes |
| **cap-04-reactions** | Five illustrated reactions, view receipts, who viewed | **Restated 2026-09-15.** [ADR-0018](../../docs/decisions/0018-reactions-the-poster-sees.md) brings back the five illustrated reactions, Smile, Wink, Sad, Wow and Love. Only the poster sees them, by name and face and never as a count. The kind is sealed, and each person gives one reaction and can change it. View receipts, the who-viewed list and reactions shown to everyone are gone. `DESIGN.md` has its reaction button back. | Yes |
| **cap-05-messaging** | One-to-one and group chat, ephemeral messages | **Removed 2026-09-15.** Nah? has no chat: people already have one, and ADR-0017 puts Nah? in the gap next to it. `product-principles` says so, and the January spec is in git history at f725474. | Removed |
| **cap-06-ambient-presence** | Now playing, battery, transit | **Removed 2026-09-15.** Music, the part worth keeping, is now a moment type in cap-03, shared deliberately. Battery and transit are dropped, and arriving somewhere is a text moment or a place. The January spec is in git history at f725474. | Removed |
| **cap-07-user-ownership** | Export, deletion, a plain privacy policy | Not restated yet. Export stays a baseline. The January spec's "admins can read the database" contradicts [ADR-0012](../../docs/decisions/0012-encrypted-on-device.md). | No |
| **cap-08-community-funding** | Donations, supporter perks, a crowdfunded roadmap | Not restated yet. No monetisation in the MVP. | No |
| **cap-09-comments** | Single-level replies (a nah-vision spec in January) | **Created 2026-09-15** as its own future change, split out of nah-vision. Cut from the MVP (ADR-0008). | No |

## Why the January scope is kept

Decisions evolve, and the evolution is part of the record. Until a capability is restated, its January spec stays beside it as source material, so the original ambition is visible until it is deliberately kept, changed or cut. Once a change is restated, the January text lives in git history, and the decision records say why it changed.

## Pointers

- [Decision records](../../docs/decisions/index.md)
- [Codebase wiki](../../docs/wiki/quickstart.md)
- [SiYuan: Nah? — April Decisions (2026-04-25 to 27)](siyuan://blocks/20260520000446-8bsmo59), the April session context
