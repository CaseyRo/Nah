# Core identity

## Why

The M1 walking skeleton runs on Flutter's defaults on purpose (CDI-1832): nothing about how Nah? looks was on the path to two phones talking. The January version of this change was written for a different product, with Mastodon underneath, email and password, reactions, comment counts, five moment types, a Close Friends picker and a tab bar, so nothing could be built from it as it stood. This restatement describes the look, the shell and the brand for Nah? as it stands now (ADR-0017: one circle of up to 150 per person and never a group, by invitation only, moments the server cannot read), so the screens M2 and M3 add are built once.

## What Changes

- **design-system**: how the app looks and moves. `DESIGN.md` is the one home for every value (colour, type, spacing, radius, shadow, motion). The spec states only behaviour that can be checked: light and dark follow the system, the display font ships inside the app, reduced motion is honoured, text is readable and operable by everyone, and no text is ever set on Pomegranate.
- **radial-menu**: the + that blooms into four moment types, text, voice, photo and music (ADR-0006, amended 2026-09-15). Choosing one opens its composer directly, with no audience step (ADR-0017).
- **app-shell**: a native iOS and Android app called "Nah?". The feed is the whole screen: no tab bar, and the timeline clock and the + are its only chrome (ADR-0008, ruled again 2026-09-15). It opens on the feed or asks for an invitation, shows the last feed offline, asks for notification permission only with a reason, and contains nothing that reports on the person.
- **brand-assets**: logo, app icons, and the site's favicon and link preview. A link preview never shows a person.
- **Removed from the January version**:
  - `pwa-shell`, superseded by `app-shell` in May (ADR-0001)
  - email and password sign-up (ADR-0015)
  - reaction icons, which wait on CDI-1856 and belong to cap-04
  - comment counts and every badge (ADR-0004, ADR-0007)
  - the Close Friends chip (ADR-0017)
  - presence dots (cap-06)
  - view receipts (ADR-0004)
  - the tab bar
  - Firebase and a streaming socket (ADR-0010)
  - a separate `nah_ui` package
- **Moved to other changes**:
  - onboarding and the full-circle message go to cap-02, where joining is connecting
  - moment card variants, the timeline clock, the offline queue and the media cache go to cap-03

## Capabilities

### New Capabilities

- `design-system`: behaviour of the theme, type, motion and accessibility, with every value held in `DESIGN.md`.
- `radial-menu`: the + and its bloom of three moment types.
- `app-shell`: the installed app, what it opens on, how it stays current, and what it never talks to.
- `brand-assets`: logo, app icons, favicon and link previews.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- `apps/mobile`: a theme, a bundled font and the feed's chrome replace Flutter's defaults; the icons and splash replace Flutter's placeholders in `ios/` and `android/`. No new package and no new dependency.
- `DESIGN.md`: brought in line with the decision records in the same pass. Reactions, comment counts, privacy chips, the tab bar, presence and the music, location and status cards are gone from it.
- Linear: no issue covers this yet. The skeleton (CDI-1832) left it out deliberately.
- Not decided here: the logo itself, and the final bundle identifier (`works.cdit.nah` is a placeholder and becomes permanent at first store submission).
