# Core identity design

## Context

The January design for this change was a full visual specification, written before most of Nah?'s decisions existed, and `DESIGN.md` was later generated from it. Since then ADR-0004 removed counts, ADR-0006 cut moments to three types, ADR-0008 cut reactions and comments, ADR-0010 removed Mastodon, ADR-0015 removed passwords, and ADR-0017 removed circles and the audience picker. The app today is the M1 skeleton on Flutter's defaults.

This document records only the decisions that shape cap-01. Values live in `DESIGN.md`; reasons live in `docs/decisions/`.

## Goals / Non-Goals

**Goals**

- One home for every design value, so a spec and the app cannot disagree with it.
- A shell and a + menu that fit one network with one audience.
- An app that makes requests only to Nah?'s own addresses.

**Non-Goals**

- Moment cards, the timeline clock, onboarding, reactions and presence. Each belongs to another change (see the table below).
- Building anything. This round is specification only.

## Decisions

### `DESIGN.md` holds the values, the spec holds behaviour

Ruled 2026-09-15. The two copies had already drifted: the January spec used an elastic curve that `DESIGN.md` forbids. Hex codes, the type scale, spacing, radii, shadows and motion timings live in `DESIGN.md` only. The design-system spec says what must be observably true and points there for numbers.

Rejected: keeping values in both and syncing them by hand, which is how they drifted.

### No tab bar

Ruled 2026-09-15, following ADR-0008 ("floating timeline clock as the only navigation chrome"). With one network and one audience there are three places: the feed, your people, and you. The feed is the app, and the other two open from its header. A tab bar would be the only thing on screen that is not the feed, and its badges were already ruled out by ADR-0007.

Rejected: Feed / People / You tabs as `DESIGN.md` had them. That costs a permanent strip of chrome for two destinations people rarely visit.

### Three items in the + menu, and no audience step

ADR-0006 sets text, voice and photo. ADR-0017 removed the audience picker, so choosing an item opens its composer and nothing sits between them. The radial bloom stays for three items because it is the product's signature gesture, not because three items need a menu.

### The display font ships in the app

Nunito is bundled rather than fetched at runtime. A runtime fetch is a request to Google from every fresh install, which breaks the no-trackers guarantee (CDI-1860), and it fails offline on first launch.

### No design package

The January plan put tokens in `/packages/nah_ui`. There is one app and no second consumer, so the theme lives in `apps/mobile`. Split it out when something else needs it.

### Nothing is written on Pomegranate

Ruled 2026-09-15, after four options were mocked in the Nah? Paper file. White 14px text on Pomegranate measures 4.08:1, under the 4.5:1 WCAG AA asks of text, and Pomegranate is not changing. So a primary action is a Pomegranate circle carrying a white glyph, which needs only 3:1, with its label beside it. The + and the send button already work this way, and the rule makes them one component. Destructive actions follow the same shape on Error, where white measures 3.76:1.

Rejected: a 19px bold label, which passes as large text but gives buttons a type size of their own; a dark ink label, which passes at 4.63:1 but reads like a warning; and judging by APCA, which scores the pair Lc 72 but still fails the WCAG 2 checks that audits and EU rules apply.

### Staying current without a socket

The streaming connection left with Mastodon. The app loads the newest page on launch, on return to the foreground and on pull to refresh, never loads a second page (CDI-1833), and never polls in the background. The server's 30 second recheck of each person's newest moment is internal to the feed query, and the app cannot see it.

### Push through the relay

Notifications arrive through Nah?'s own relay, which holds the Apple and Google keys (ADR-0010, CDI-1848, M4). What a notification says and when it is sent belongs to ADR-0007.

## Where the January scope went

| January item | Now |
|---|---|
| `pwa-shell` spec | Removed. Superseded by `app-shell` (ADR-0001); in git history at 655cfb7. |
| Email and password sign-up, verification, forgotten password | Removed. A key on the device (ADR-0015) and joining by invitation. |
| Four-step onboarding and the "150 friends max" explainer | cap-02, as joining, following ADR-0005's ritual. |
| Friend-limit dialog and inactive-friend suggestions | cap-02. A full network is said in words (ADR-0004), with no suggestions. |
| Reaction icons and the heart button | cap-04, waiting on CDI-1856. |
| Comment count, notification count, pending-request dot | Removed (ADR-0004, ADR-0007). |
| Close Friends / All Friends chip | Removed (ADR-0017). |
| Presence dots and status colours | cap-06. |
| View-receipt threshold | Removed (ADR-0004, ADR-0008). |
| Music, location and status moments | Out for now (ADR-0006). cap-03 owns moment types. |
| Moment card variants, timeline clock, end of feed | cap-03. |
| Offline queue and media cache | cap-03 (CDI-1847, M3). |

## Risks / Trade-offs

- **Your people sit one tap deeper without a tab bar.** ADR-0017 makes connecting mostly physical, so the list should be visited rarely. If people keep looking for it, revisit.
- **Android push may need Google's messaging client in the app.** The no-reporting requirement covers requests the app makes itself; M4 has to check what that client sends before it is added.
- **`DESIGN.md` loses the sections that described a richer product.** They stay in git history, and the decision records say why they went.

## Open Questions

- **The logo**, which does not exist yet.
- **The final bundle identifier**, which cannot change after first submission.
- **Whether "you" gets a profile header.** That decides whether the second glassmorphism surface in `DESIGN.md` survives.
- **Whether the feed stored on the device needs encryption** beyond what the operating system provides, once moments are sealed (CDI-1863).
