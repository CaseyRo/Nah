---
name: Nah
description: A private home for your closest people, designed as a walled garden.
colors:
  pomegranate: "#EE3423"
  pomegranate-hover: "#D42D1E"
  pomegranate-light: "#FF6B5B"
  pomegranate-tint: "#FFD9D4"
  surface-1: "#FFFFFF"
  surface-2: "#FAFAFA"
  surface-3: "#F5F5F5"
  surface-4: "#EBEBEB"
  surface-5: "#E0E0E0"
  text-1: "#1A1A1A"
  text-2: "#6B6B6B"
  text-3: "#9E9E9E"
  text-4: "#BDBDBD"
  text-inverse: "#FFFFFF"
  surface-1-dark: "#201E1C"
  surface-2-dark: "#171514"
  surface-3-dark: "#292624"
  surface-4-dark: "#363230"
  text-1-dark: "#F2EEEB"
  text-2-dark: "#A7A19D"
  text-4-dark: "#5A5450"
  success: "#22C55E"
  warning: "#F59E0B"
  error: "#EF4444"
typography:
  display:
    fontFamily: "Nunito, -apple-system, BlinkMacSystemFont, sans-serif"
    fontSize: "1.875rem"
    fontWeight: 700
    lineHeight: 1.2
    letterSpacing: "-0.01em"
  headline:
    fontFamily: "Nunito, -apple-system, BlinkMacSystemFont, sans-serif"
    fontSize: "1.5rem"
    fontWeight: 700
    lineHeight: 1.2
    letterSpacing: "-0.01em"
  title:
    fontFamily: "Nunito, -apple-system, BlinkMacSystemFont, sans-serif"
    fontSize: "1.25rem"
    fontWeight: 600
    lineHeight: 1.3
    letterSpacing: "normal"
  name:
    fontFamily: "-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif"
    fontSize: "1rem"
    fontWeight: 600
    lineHeight: 1.375
    letterSpacing: "normal"
  moment:
    fontFamily: "-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif"
    fontSize: "1.0625rem"
    fontWeight: 400
    lineHeight: "25px"
    letterSpacing: "normal"
  body:
    fontFamily: "-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif"
    fontSize: "1rem"
    fontWeight: 400
    lineHeight: 1.5
    letterSpacing: "normal"
  label:
    fontFamily: "-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif"
    fontSize: "0.875rem"
    fontWeight: 500
    lineHeight: 1.4
    letterSpacing: "0.02em"
  caption:
    fontFamily: "-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif"
    fontSize: "0.8125rem"
    fontWeight: 500
    lineHeight: 1.4
    letterSpacing: "normal"
    fontFeature: "tnum"
rounded:
  sm: "6px"
  md: "8px"
  lg: "12px"
  full: "9999px"
spacing:
  "1": "4px"
  "2": "8px"
  "3": "12px"
  "4": "16px"
  "5": "20px"
  "6": "24px"
  "8": "32px"
  "10": "40px"
  "12": "48px"
  lane: "64px"
  fab-inset: "24px"
components:
  button-primary:
    backgroundColor: "{colors.pomegranate}"
    textColor: "{colors.text-inverse}"
    rounded: "{rounded.full}"
    size: "48px"
  button-primary-label:
    textColor: "{colors.text-1}"
    typography: "{typography.label}"
  button-primary-hover:
    backgroundColor: "{colors.pomegranate-hover}"
    textColor: "{colors.text-inverse}"
    rounded: "{rounded.full}"
  button-ghost:
    backgroundColor: "{colors.surface-1}"
    textColor: "{colors.text-1}"
    typography: "{typography.label}"
    rounded: "{rounded.full}"
    padding: "12px 24px"
  fab-radial:
    backgroundColor: "{colors.pomegranate}"
    textColor: "{colors.text-inverse}"
    rounded: "{rounded.full}"
    size: "56px"
  moment-row:
    backgroundColor: "{colors.surface-2}"
    textColor: "{colors.text-1}"
    typography: "{typography.moment}"
    padding: "16px 16px 20px"
  moment-lane:
    textColor: "{colors.text-2}"
    typography: "{typography.caption}"
    width: "{spacing.lane}"
    padding: "16px 12px 0 0"
  timeline-clock:
    backgroundColor: "{colors.surface-1}"
    textColor: "{colors.text-1}"
    rounded: "{rounded.full}"
    size: "48px"
  chip-default:
    backgroundColor: "{colors.surface-2}"
    textColor: "{colors.text-2}"
    typography: "{typography.label}"
    rounded: "{rounded.full}"
    padding: "6px 12px"
    height: "32px"
  chip-selected:
    backgroundColor: "{colors.pomegranate}"
    textColor: "{colors.text-inverse}"
    typography: "{typography.label}"
    rounded: "{rounded.full}"
    padding: "6px 12px"
    height: "32px"
  reaction-button:
    backgroundColor: "{colors.surface-1}"
    textColor: "{colors.text-2}"
    rounded: "{rounded.full}"
    size: "32px"
  avatar-fallback:
    backgroundColor: "{colors.pomegranate-light}"
    textColor: "{colors.text-1}"
    rounded: "{rounded.full}"
  bottom-sheet:
    backgroundColor: "{colors.surface-1}"
    textColor: "{colors.text-1}"
    rounded: "{rounded.lg}"
    padding: "24px"
---

## 1. Overview

**Creative North Star: "The Private Garden"**

Nah is a walled garden. Mutual-entry only, hard-capped at 150 people, tended slowly over years rather than scrolled through in minutes. Every visual choice answers to that frame. Surfaces are warm and uncluttered, the way a garden is when you've just swept the path. Color is restrained, used the way a single flowering plant draws the eye in a green room. Motion is unhurried but alive, the way leaves move when you walk past. The system rejects the noise of public social media in its bones: no broadcast feeds, no algorithmic interleaving, no engagement-bait notification surfaces, no SaaS hero metrics. The garden is for the people already inside it.

The components are warm and considered. Soft 8–12px corners across the board, generous internal padding, handmade-adjacent rather than razor-precise. Spring physics on press and entrance, never sharp linear transitions. Every interactive surface should feel touched, not assembled. This is the antidote to dense product UI: nothing is crowded, nothing demands attention, everything earns it.

Color is Restrained: warm whites carry 90% of every screen, pomegranate red carries the remaining 10%. The light neutrals are still pure greys; the dark palette (set 2026-09-15) is the first to realize the long-standing direction of tinting neutrals toward the pomegranate hue, and never reaches pure black. Type is split between Nunito for moments of brand voice (logo, screen titles, onboarding) and the platform default sans for everything else (body, labels, data, controls). The platform body font matters: Nah should feel native on iOS and native on Android, not translated.

The feed (designed 2026-09-15) is a day read down a time axis, not a stack of posts. Time runs down a narrow left lane, people and their moments sit on the ground to its right, and the analog clock rides the lane as the cursor. There are no card boxes on the feed.

**Key Characteristics:**

- Light-primary identity, warm white surfaces, pomegranate accent ≤10% of pixels
- Dark mode is the same world inverted: warm-tinted dark neutrals, borders where light mode has shadows, the same red
- Soft corners (8–12px), generous padding, handmade-adjacent geometry
- Spring physics on motion, never bounce or elastic
- One display font (Nunito) for brand voice, platform default for everything else
- Glassmorphism scoped to one surface, the radial menu backdrop, and forbidden elsewhere
- Two-lane feed: a 64pt time lane with a 1pt rail, and four moment treatments set directly on the ground
- 48px analog timeline clock riding the time lane as a signature element

## 2. Colors

A restrained warm palette. The pomegranate carries the brand; the warm whites carry the room.

### Primary

- **Pomegranate** (`#EE3423`): the warm point in the room. Used for the radial FAB, primary action buttons and the analog clock's second hand and centre dot. It stays the + colour in dark mode too (2026-09-15): a white glyph on it measures 4.08:1, and it measures about 4.4:1 against the dark ground. Never used for content backgrounds, never as a gradient, never tinted onto more than 10% of any screen.
- **Pomegranate Hover** (`#D42D1E`): pressed/active state of Pomegranate. Used only as a momentary state, never at rest.
- **Pomegranate Light** (`#FF6B5B`): the avatar-fallback background, with initials in Text 1, in both modes. Also used very sparingly as a soft background tint on text-moment cards. Retired as a dark-mode primary (2026-09-15): a white glyph on it measures about 2.8:1, under the 3:1 a glyph needs.
- **Pomegranate Tint** (`#FFD9D4`): the fill inside the reaction illustrations. Identical in light and dark.

### Neutral

- **Surface 1** (`#FFFFFF`): card and modal surfaces, the clock face, the reaction button. The room's white walls.
- **Surface 2** (`#FAFAFA`): page background and the feed's ground. Subtle warmth distinguishes from Surface 1.
- **Surface 3** (`#F5F5F5`): inset surfaces, search input backgrounds, skeleton bases, the voice moment's play disc.
- **Surface 4** (`#EBEBEB`): dividers, card borders, the feed rail and the header hairline, skeleton shimmer highlight.
- **Surface 5** (`#E0E0E0`): disabled control backgrounds.
- **Text 1** (`#1A1A1A`): primary content text. Never pure black.
- **Text 2** (`#6B6B6B`): timestamps, places, metadata, secondary captions, the waveform.
- **Text 3** (`#9E9E9E`): placeholder text, disabled labels.
- **Text 4** (`#BDBDBD`): hint text, decorative dividers, the clock's four ticks.

### Dark palette

Set 2026-09-15 from the feed frames. Warm-tinted toward the pomegranate hue, never pure black. The status bar content is light. Each role keeps its light-mode job.

- **Surface 2 (dark)** (`#171514`): the ground.
- **Surface 1 (dark)** (`#201E1C`): the clock face, the reaction button, sheets and modals.
- **Surface 3 (dark)** (`#292624`): inset surfaces, the play disc.
- **Surface 4 (dark)** (`#363230`): rails, borders, the header hairline, and every place that was a shadow in light mode.
- **Text 1 (dark)** (`#F2EEEB`): primary content text.
- **Text 2 (dark)** (`#A7A19D`): about 7.1:1 on the ground.
- **Text 4 (dark)** (`#5A5450`): the clock's ticks and hint text.

Pomegranate, Pomegranate Light and Pomegranate Tint do not change in dark mode.

### Semantic

- **Success** (`#22C55E`): post-success toast accent.
- **Warning** (`#F59E0B`): soft-warning toast accent.
- **Error** (`#EF4444`): destructive confirm buttons, error toast accent, validation messages.

### Named Rules (Colors)

**The One Voice Rule.** Pomegranate appears on ≤10% of any rendered surface. Its rarity is the point. When two pomegranate elements are visible at once (e.g., the FAB and a primary button), neither is decorative — both are doing work.

**The No-Black Rule.** Pure black (`#000000`) is forbidden. Text uses `#1A1A1A`. Future palette refinements should tint neutrals toward the pomegranate hue (chroma 0.005–0.01 in OKLCH) rather than expanding the grey scale.

**The Same Red Rule.** *(Added 2026-09-15.)* Pomegranate `#EE3423` is the one red in both modes. Dark mode does not lighten it; it would fail the glyph's 3:1 and the light-mode red already passes on the dark ground.

**The No-Counts Rule.** *(Added 2026-04-26; see [ADR-0004](docs/decisions/0004-no-counts-anywhere.md).)* No numeric counts appear anywhere in the rendered UI. No like count, no view count, no follower count, no read count, no reaction count, no comment count. Where we need to communicate "someone did this," we render names and avatars: *Maya*, never *3 people*. The limit of 150 is communicated in words, never as a meter or progress bar. This rule is universal across surfaces; there is no exception screen.

## 3. Typography

**Display Font:** Nunito (fallback: `-apple-system, BlinkMacSystemFont, sans-serif`)
**Body Font:** Platform default sans (`-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif`)

**Character:** Nunito is rounded, friendly, warm — the visual equivalent of the brand voice. It carries the logo and screen titles, where Nah wants to say "this is a place, not an app." Everything functional (body, labels, buttons, data, timestamps) uses the platform's own sans, so the app feels native on iOS and native on Android. The split is intentional: brand surfaces breathe with Nunito, working surfaces disappear into the OS.

In the Paper frames Inter stands in for SF Pro; the app uses the platform sans.

### Hierarchy

- **Display** (Nunito, 30px, 700, 1.2, letter-spacing -0.01em): onboarding card titles, hero moments. Maximum two uses per screen.
- **Headline** (Nunito, 24px, 700, 1.2, letter-spacing -0.01em): screen titles, the wordmark in the feed header, profile display names on the profile page only.
- **Title** (Nunito, 20px, 600, 1.3): section headers within a screen.
- **Name** (platform sans, 16px, 600, 22px): a person's name on a moment, a song title. Corrected 2026-09-15: names were recorded as Nunito Title, but Nunito is limited to the wordmark, screen titles and onboarding.
- **Moment** (platform sans, 17px, 400, 25px): the text of a text moment.
- **Body** (platform sans, 16px, 400, 1.5): prose, the short line under a photo or voice moment (16px/22px), card content. Cap line length at 65–75ch where prose is long-form.
- **Label** (platform sans, 14px, 500, 1.4, letter-spacing 0.02em): button text, chip text, navigation, a song's artist (14px, Text 2).
- **Caption** (platform sans, 13px, 500, 1.4, tabular figures): the time in the feed lane, a voice moment's duration, the place under a moment, the names beside reactions on your own moment. The clock's date drops to 12px/500. Changed 2026-09-15 from 12px/400.
- **Marker** (platform sans, 15px, 400, Text 2): the deleted marker and the end of the feed.

### Named Rules (Typography)

**The Display-Off Rule.** Nunito never appears on functional UI elements. Buttons use Label. Form labels use Label. Data uses Body or smaller. Display fonts on buttons read as marketing copy, which Nah is not.

**The System-Body Rule.** Body text always uses the platform default sans, never Nunito. The trade is intentional: a small loss of cross-platform pixel-identity, a large gain in feeling-like-the-OS.

**The Wordmark-Only Rule.** *(Added 2026-09-15.)* On the feed, Nunito appears once: the wordmark. Names, times, places and moment text are the platform sans.

## 4. Layout

*(Section added 2026-09-15 from the feed frames, 390x844 at 2x.)*

The feed is two lanes on the Surface 2 ground, with no card boxes. Moments sit directly on the ground.

- **Time lane:** 64pt wide on the left. Each moment's time sits at the top of its row in Caption, Text 2, tabular figures, right-aligned with 12pt padding from the lane's right edge. A 1pt Surface 4 rail runs along the lane's right edge for the length of the feed.
- **Content column:** everything to the right of the rail, padded 16pt left and right. Rows are padded 16pt above and 20pt below, with 8 to 10pt between name, body and foot.
- **Row order:** name, then the moment's body, then the foot line. The foot line starts with the 32pt reaction button, then the place, 12pt apart. The right end of the foot stays free so the + never covers a control.
- **Photos** ignore the right padding and bleed to the right screen edge.
- **Markers keep their place in the lane.** A deleted moment keeps its time in the lane and reads "Maya deleted this moment." in Marker type in the content column.
- **The end of the feed** sits outside the lane structure: "You're all caught up." in Marker type at the content column's left edge, with no rail beside it. The rail ends with the last moment.
- **Header:** three overlapping 24pt faces on the left (opens your circle), the wordmark centred in Headline, your 32pt avatar on the right, and a 1pt Surface 4 hairline below. Nothing else sits above the feed.
- **Chrome:** the timeline clock in the lane and the + at bottom-right are the only chrome over the feed. There is no tab bar.

**The Two-Lane Rule.** Time on the left, people on the right. The lane, the rail and the clock are what make it the feed; every moment treatment fits inside the content column and never draws a box around itself.

## 5. Elevation

Nah is flat by default. Cards rest on the page without shadow at rest. Depth is conveyed through surface tonality (Surface 1 over Surface 2) and 1px borders in `Surface 4`, not through ambient drop shadows. This is the garden's quiet: nothing floats unless it's doing something.

Shadows appear only as a response to state: an active radial FAB carries a subtle elevation, a bottom sheet casts a soft drop, a toast lifts slightly above the feed. None of these are decorative; each marks a thing that is happening *now*.

In dark mode, shadows are invisible on dark surfaces. Cards switch to 1px borders in `Surface 4` (dark variant) and rely entirely on tonal stepping between surface levels. This is the same architecture inverted, not a separate visual language.

### Shadow Vocabulary

- **FAB ambient** (`box-shadow: 0 4px 12px rgba(238, 52, 35, 0.15)`): the radial FAB's at-rest elevation. Tinted toward pomegranate so the shadow extends the warm point rather than darkening it.
- **Bottom sheet** (`box-shadow: 0 -4px 24px rgba(0, 0, 0, 0.08)`): the soft lift of the composer above the page.
- **Toast** (`box-shadow: 0 4px 16px rgba(0, 0, 0, 0.10)`): the brief lift of a toast notification above the feed.
- **Clock** (`box-shadow: 0 4px 12px rgba(0, 0, 0, 0.10)`): the timeline clock's lift off the lane in light mode. In dark mode there is no shadow; its 1pt Surface 4 stroke is the border. *(Added 2026-09-15.)*

### Named Rules (Elevation)

**The Flat-By-Default Rule.** Surfaces are flat at rest. Shadows appear only on active or transient elements (FAB, bottom sheet, toast, modal, the clock while it rides). A card never has a shadow. A button never has a shadow.

**The Border-Replaces-Shadow Rule.** In dark mode, every place that uses a shadow in light mode uses a 1px border instead. The rule cascades automatically; widgets check the theme brightness and swap.

## 6. Components

Warm and considered. Soft corners, generous padding, handmade-adjacent geometry. Spring physics on motion.

### Buttons

- **Shape:** pill (`rounded-full`, 9999px) for text buttons, circle for glyph buttons. Pill rather than rectangle because Nah? is a place, not a console.
- **Primary:** a 48px Pomegranate circle carrying a white 20px glyph, with its label beside it in Label typography and Text 1. No text is ever set on Pomegranate: white on `#EE3423` measures 4.08:1, under WCAG AA's 4.5:1 for text, while a glyph needs only 3:1. Pressed swaps the circle to Pomegranate Hover with a 150ms ease-out transition. The radial FAB and the send button are this component at their own sizes.
- **Ghost:** Surface 1 background, Text 1 text, same shape and padding. Used for secondary actions in composer and onboarding. Pressed state shifts background to Surface 2.
- **Destructive:** the primary shape with an Error circle, because white on `#EF4444` measures 3.76:1 and the same rule applies. Used only for "Discard" and "Leave Nah?" confirmations.
- **Focus:** 2px Pomegranate outline at 4px offset from the button edge. Visible only on keyboard focus.

### Floating Action Button (Radial FAB)

- **Size:** 56px circle.
- **Color:** Pomegranate background, white "+" glyph at 24px. The same Pomegranate in dark mode (The Same Red Rule).
- **Position:** fixed bottom-right, 24pt from the right edge and 24pt above the safe-area bottom, which is 58pt from the screen edge on a home-indicator phone. Corrected 2026-09-15 from "24px inset from screen edges". Left-hand mode mirrors to bottom-left.
- **At rest:** carries the FAB ambient shadow (pomegranate-tinted).
- **Active (menu open):** the "+" rotates 45° to "×" over 250ms, menu items fan outward with 30ms stagger and spring entrance (400ms total).

### Moments in the feed

*(Rewritten 2026-09-15. The feed has no card boxes; see Layout for the lane structure every moment sits in. The card shell below this section no longer describes the feed.)*

Every moment is name, body, foot inside the content column. The four types differ only in the body ([ADR-0006](docs/decisions/0006-three-moment-types.md)):

1. **Text:** the text in Moment type (17px/25px), Text 1. Foot: reaction button, place.
2. **Photo:** the photo bleeds from the content column's left edge to the right screen edge, 4:3 shown at 206pt tall, with an 8px radius on the left corners only and square corners at the screen edge. A short line in Body (16px/22px) below, then the foot with the place.
3. **Voice:** a 40pt Surface 3 disc carrying a Text 1 play glyph on the left, a waveform of 2pt bars in Text 2 across the middle, the duration in Caption on the right. A short line in Body and the foot below. *(Designed 2026-09-15; it was unresolved before.)*
4. **Music:** 64pt artwork with a 6px radius on the left; song title in Name and artist in 14px Text 2 stacked on the right. Tapping opens the song in a music app the reader chooses.

**Your own moment** shows no reaction button. Its foot is the reactions it received: each as a 24pt face, the 22pt reaction illustration, and the name in Caption, in a row.

**Deleted marker:** "Maya deleted this moment." in Marker type, in the lane structure, with its time kept in the lane.

### Moment Cards (outside the feed)

The card shell is kept for surfaces that still box a moment (a single moment opened on its own, a person's page). It is not used on the feed.

- **Corner Style:** 8px radius (`md`). Soft, considered, not razor-sharp.
- **Background:** Surface 1 default. Text moments may opt into a soft Pomegranate Light tint (≤8% saturation) as a background-color choice in the composer.
- **Shadow Strategy:** none at rest (flat-by-default rule).
- **Border:** none in light mode. 1px Surface 4 in dark mode.
- **Internal Padding:** zero for photo moments (edge-to-edge media). 24px for text moments. Custom for voice and music.
- **Header:** 40px avatar + display name (Name) + timestamp (Caption) in a horizontal row at top of card.

### Chips

- **Style:** pill shape, 32px height, 12px horizontal padding, Label typography.
- **Default state:** Surface 2 background, Text 2 text.
- **Selected state:** Pomegranate background, white text. Smooth 200ms transition between states.
- **Disabled state:** Surface 5 background, Text 3 text.

### Reaction Button

- **Size:** 32px circle, inside a 44px tap area.
- **At rest:** Surface 1 background with a 1px Surface 4 border, Text 2 outline-heart icon. The border was added 2026-09-15 because the button sits on the ground, not on a card; without it Surface 1 on Surface 2 disappears.
- **Reacted:** the reaction the person gave replaces the outline heart, in full colour. Only the poster sees other people's reactions, by name and face, never as a count ([ADR-0018](docs/decisions/0018-reactions-the-poster-sees.md)).
- **Tap:** a picker slides in above the button with the five reactions, Smile, Wink, Sad, Wow and Love, in a horizontal row, with a 250ms spring entrance.
- **Illustrations:** Nah?'s own, not system emoji. One drawn hand, 22px, a Text 1 outline at 1.5, fills in Pomegranate Tint, Pomegranate as the accent. Identical in light and dark.

### Bottom Sheet (Composer)

- **Shape:** rounded only at the top (`rounded-t-lg`, 12px), flat at the bottom edge of the screen.
- **Background:** Surface 1. Semi-transparent dim backdrop (not blur) behind the sheet.
- **Drag handle:** 36×4px pill in Surface 4, centered, 12px from the top edge.
- **Animation:** slides up from below with 400ms spring entrance, drag-to-dismiss past 1/3 swipe threshold.
- **Max height:** 90vh. Content scrolls internally if it exceeds.

### Timeline Clock (signature component)

*(Rewritten 2026-09-15.)* A 48px analog clock face that rides the feed's time lane. It is laid over the lane label of the moment nearest the viewport's vertical centre, so the lane's time and the clock's hands say the same thing; it never floats over content. The face is Surface 1 with a 1px Surface 4 stroke, four ticks in Text 4, hour and minute hands in Text 1, the second hand and centre dot in Pomegranate. The date ("Sat 13 Sep") sits centred below the face in 12px/500 Text 2. In light mode the face carries the Clock shadow; in dark mode there is no shadow and the stroke is the border. Hour and minute hands animate smoothly as the nearest moment changes. Appears on scroll-start, fades after 2 seconds of scroll inactivity. Under `prefers-reduced-motion`, falls back to a static digital time display.

This is the single most distinctive Nah component. It is not optional.

### No Tab Bar

There is no tab bar. The feed is the whole screen, and your circle and you open from its header. The timeline clock and the radial FAB are the only chrome over the feed, and nothing carries a badge.

### Avatar

- **Sizes:** xs 24px (the header's circle faces, reactions on your own moment), sm 32px (the header's own avatar), md 40px (card headers, your circle), lg 56px, xl 80px (the top of a person's page, which has no cover photo).
- **Shape:** circular.
- **Fallback:** initials in Text 1 on a Pomegranate Light background, in both modes. White measures about 2.8:1 there, under WCAG AA; Text 1 measures about 6.2:1.

## 7. Do's and Don'ts

### Do

- **Do** keep Pomegranate at ≤10% of any rendered screen. It is a focal point, not a fill.
- **Do** keep Pomegranate `#EE3423` as the one red in both modes. Dark mode changes the neutrals, never the red.
- **Do** use Nunito for the logo, screen titles, and onboarding headings. Use the platform default sans everywhere else, names included.
- **Do** set feed moments directly on the ground in the two-lane structure: 64pt time lane, 1pt rail, content column padded 16pt. No box around a moment on the feed.
- **Do** use 8px corner radius on cards and on the left corners of feed photos, 6px on album artwork, 12px on bottom sheets, full-radius pills on buttons and chips. Never sharp 90° corners on touchable surfaces.
- **Do** ship spring physics on every entrance, press, and dismiss. 400ms with overshoot for entrances, 250ms ease-out for exits.
- **Do** use the 48px timeline clock riding the time lane as the signature scroll companion. It is the single most recognizable Nah element.
- **Do** treat the limit of 150 as a designed boundary, never a limitation. If it is ever drawn, draw faces, never something that fills towards a number.
- **Do** keep the Feed chronological, full stop. The visual hierarchy must never imply algorithmic ranking, suggestions, or "you might like."
- **Do** use Pomegranate Light as a soft text-moment background tint, sparingly, opt-in.
- **Do** show empty states with warm illustrations, never blank screens. "Welcome home" framing on first run, "You're all caught up" framing at end of feed, in Marker type outside the lane.

### Don't

- **Don't** look like Instagram, TikTok, or X. No Stories rail at the top. No suggested-content interleaving. No infinite-scroll engagement bait. No algorithmic feed under any circumstance.
- **Don't** look like LinkedIn or Threads. No follower counts. No public profiles. No professional-self-curation chrome. No mutual-connection paths shown publicly.
- **Don't** display numeric counts of anything, anywhere. No like count, no view count, no follower count, no read count, no reaction count, no comment count. Names and avatars instead. The No-Counts Rule is universal; there is no exception screen.
- **Don't** look like Discord or Slack. No aggressive unread counts as the primary navigation signal. No multi-server sidebars. No "typing" indicators in v1.
- **Don't** look like a VC-SaaS landing page. No gradient hero text. No three-feature card grids on any product surface. No hero-metric templates. No "X is faster than Y" comparison tables. No glassmorphism dashboards.
- **Don't** use side-stripe borders (colored `border-left` or `border-right` greater than 1px). Never. On any card, list item, callout, or alert.
- **Don't** use gradient text (`background-clip: text` with a gradient background). Solid colors only. Emphasis through weight or size.
- **Don't** use glassmorphism anywhere but the radial menu backdrop. Forbidden everywhere else, including a person's page, modals, sheets, and notification toasts.
- **Don't** use display fonts (Nunito) on buttons, form labels, data, timestamps, names, or any other functional UI element. Display fonts on buttons read as marketing copy.
- **Don't** use pure black (`#000000`) or pure white (`#FFFFFF`) for text, in either mode. Text uses `#1A1A1A` in light and `#F2EEEB` in dark; the dark ground is `#171514`, never black.
- **Don't** put a reaction button on your own moment. It shows the reactions it received, by face and name.
- **Don't** use bounce or elastic easing curves. Spring physics are warm; bounce and elastic are loud.
- **Don't** add card grids of identical-size feature blocks. Each moment type has its own treatment; the Feed is varied by design.
- **Don't** introduce a modal as a first-thought solution. Exhaust inline, progressive, and bottom-sheet alternatives first. Bottom sheets are the Nah default for any "secondary surface" need.
