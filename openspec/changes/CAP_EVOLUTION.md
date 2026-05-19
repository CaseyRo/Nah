# Capability Evolution

State of the eight capability scaffolds, and how they've evolved since the original 2026-01 framing.

This document exists because **scope evolved 2026-04-26** ([ADR-0008](../../docs/decisions/0008-mvp-scope.md)) but the original cap-XX directories were not touched. The scaffolds are preserved so the original ambition is visible, not erased.

## Status per cap

| Cap | Original scope (2026-01) | Current state | Evolution |
|---|---|---|---|
| **cap-01-core-identity** | App identity, design system, radial menu, ~~PWA shell~~ → app shell, brand assets | **Active.** Pivoted client stack 2026-05-19 ([ADR-0001](../../docs/decisions/0001-flutter-over-pwa.md)). `pwa-shell` superseded by `app-shell`. | Stack changed, scope unchanged. |
| **cap-02-friend-circles** | 150-friend cap, mutual friendship, inner circle for selective sharing | **Scope absorbed into MVP** as the `circle` sub-capability of cap-mvp (drafted in conversation 2026-04-26, not yet committed). The slot-economy mechanic was rejected; invitation IS connection ([ADR-0003](../../docs/decisions/0003-connection-model.md)). Inner-circle / selective-sharing is post-MVP. | Empty scaffold. Scope re-cut, not deleted. |
| **cap-03-moments** | Timeline of life moments: photo, text, music, location, sleep/wake | **Scope reduced to MVP** ([ADR-0006](../../docs/decisions/0006-three-moment-types.md)). MVP cut is three types: text, voice, photo. Music share, location, status are post-MVP. Video deferred. | Empty scaffold. Scope reduced for MVP. |
| **cap-04-reactions** | Custom illustrated reaction icons (Smile, Wink, Sad, Wow, Love) | **Cut from MVP** ([ADR-0008](../../docs/decisions/0008-mvp-scope.md)). No reactions in MVP at all. Will return in a later change. | Empty scaffold. Deferred. |
| **cap-05-messaging** | Private 1:1 and small group chat (Path Talk inspired) | **Cut from MVP** (already flagged as v1.5 in nah-vision; the April decision keeps it cut). | Empty scaffold. Deferred. |
| **cap-06-ambient-presence** | Opt-in status sharing (what you're listening to, battery, transit) | **Cut from MVP.** The "what's playing" idea is also off the table technically (no platform exposes other apps' Now Playing to third parties). Phase 2+. | Empty scaffold. Deferred, partly platform-blocked. |
| **cap-07-user-ownership** | Data export, account portability, transparent privacy controls | **Partially absorbed into MVP** as a baseline (data export from day one). Full account-portability features deferred. | Empty scaffold. Reduced scope absorbed. |
| **cap-08-community-funding** | Donation model, supporter perks, crowdfunded roadmap | **Cut from MVP.** No monetization layer in MVP at all. The conversation flagged "potential monetary interesting avenues" but explicitly post-MVP. | Empty scaffold. Deferred. |

## Why this isn't a clean "archive these and start fresh"

Per the project's principle (~~things are never perfect from the start~~ — they aren't, that's the point): decisions evolve, and the evolution is part of the artifact. Deleting the cap-02..08 scaffolds would erase the original ambition. Keeping them with this evolution note preserves it.

When implementation begins on any of these post-MVP scopes, the relevant cap-XX directory will be populated with a proper proposal/design/tasks set, and this row in the table will get updated. The ADR layer documents the *why*.

## Pointers

- [ADR-0001 — Flutter over PWA](../../docs/decisions/0001-flutter-over-pwa.md)
- [ADR-0003 — Invitation is connection](../../docs/decisions/0003-connection-model.md)
- [ADR-0006 — Three moment types](../../docs/decisions/0006-three-moment-types.md)
- [ADR-0008 — MVP scope](../../docs/decisions/0008-mvp-scope.md)
- [SiYuan: Nah? — April Decisions (2026-04-25 to 27)](siyuan://blocks/20260520000446-8bsmo59) — the full session context
