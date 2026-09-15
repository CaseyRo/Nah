# Capability evolution

How the capability changes have moved since they were first written on 2026-01-28, and where each stands now.

Scope moved twice: on 2026-04-26 with the MVP cut ([ADR-0008](../../docs/decisions/0008-mvp-scope.md)), and on 2026-09-13 with a server of our own ([ADR-0010](../../docs/decisions/0010-small-server-not-mastodon.md)) and one network of 150 with no circles ([ADR-0017](../../docs/decisions/0017-one-network-of-a-hundred-and-fifty.md)). From 2026-09-15 the changes are restated one at a time against where Nah? stands, detailing specs and building nothing. A restated change passes `openspec validate --strict`. A change not yet restated keeps its January scope beside it as `from-nah-vision.md`, outside `specs/`.

## Status

| Change | January scope | Now | Validates |
|---|---|---|---|
| **nah-vision** | A Path successor on a Mastodon fork, with nine capability specs | **Restated 2026-09-15.** Principles only: one spec, `product-principles`, which every capability keeps. The January specs moved to their caps, and design.md is now a map to the decision records and the wiki. | Yes |
| **cap-01-core-identity** | App identity, design system, radial menu, PWA shell, brand assets | **Restated 2026-09-15.** `DESIGN.md` holds every value. No tab bar. The + offers text, voice and photo. No text is set on Pomegranate. Fonts ship in the app, and no reporting SDK is allowed. `pwa-shell` was deleted and is in git history at 655cfb7. Onboarding moved to cap-02; cards, the clock and the offline queue to cap-03; reaction icons to cap-04; "no numbers about people" to nah-vision. | Yes |
| **cap-02-friend-circles** | 150-friend cap, friend requests, an inner circle | Not restated yet. ADR-0017 replaces circles with one network per person, so the inner circle goes and the name will change. | No |
| **cap-03-moments** | Photo and video, text, music, location, sleep and wake | Not restated yet. Three types, text, voice and photo ([ADR-0006](../../docs/decisions/0006-three-moment-types.md)). It also takes moment cards, the timeline clock and the offline queue from cap-01. | No |
| **cap-04-reactions** | Five illustrated reactions, view receipts, who viewed | Not restated yet. Cut from the MVP (ADR-0008); the acknowledgement question is CDI-1856. View receipts contradict [ADR-0004](../../docs/decisions/0004-no-counts-anywhere.md). | No |
| **cap-05-messaging** | One-to-one and group chat, ephemeral messages | Not restated yet. Deferred. | No |
| **cap-06-ambient-presence** | Now playing, battery, transit | Not restated yet. Deferred, and now playing is blocked by the platforms. | No |
| **cap-07-user-ownership** | Export, deletion, a plain privacy policy | Not restated yet. Export stays a baseline. The January spec's "admins can read the database" contradicts [ADR-0012](../../docs/decisions/0012-encrypted-on-device.md). | No |
| **cap-08-community-funding** | Donations, supporter perks, a crowdfunded roadmap | Not restated yet. No monetisation in the MVP. | No |
| **cap-09-comments** | Single-level replies (a nah-vision spec in January) | **Created 2026-09-15** as its own future change, split out of nah-vision. Cut from the MVP (ADR-0008). | No |

## Why the January scope is kept

Decisions evolve, and the evolution is part of the record. Until a capability is restated, its January spec stays beside it as source material, so the original ambition is visible until it is deliberately kept, changed or cut. Once a change is restated, the January text lives in git history, and the decision records say why it changed.

## Pointers

- [Decision records](../../docs/decisions/index.md)
- [Codebase wiki](../../docs/wiki/quickstart.md)
- [SiYuan: Nah? — April Decisions (2026-04-25 to 27)](siyuan://blocks/20260520000446-8bsmo59), the April session context
