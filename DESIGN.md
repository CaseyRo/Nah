---
name: Nah
description: A private home for your closest people, designed as a walled garden.
colors:
  pomegranate: "#EE3423"
  pomegranate-hover: "#D42D1E"
  pomegranate-light: "#FF6B5B"
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
  moment-card:
    backgroundColor: "{colors.surface-1}"
    textColor: "{colors.text-1}"
    rounded: "{rounded.md}"
    padding: "0"
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

Color is Restrained: warm whites carry 90% of every screen, pomegranate red carries the remaining 10%. The neutral palette stays cool-warm-balanced (a faint pull toward the pomegranate hue is the long-term direction, currently unrealized as pure greys). Type is split between Nunito for moments of brand voice (logo, screen titles, onboarding) and the platform default sans for everything else (body, labels, data, controls). The platform body font matters: Nah should feel native on iOS and native on Android, not translated.

**Key Characteristics:**

- Light-primary identity, warm white surfaces, pomegranate accent ≤10% of pixels
- Soft corners (8–12px), generous padding, handmade-adjacent geometry
- Spring physics on motion, never bounce or elastic
- One display font (Nunito) for brand voice, platform default for everything else
- Glassmorphism scoped to exactly two surfaces (profile header compression, radial menu backdrop), forbidden elsewhere
- Four distinct moment card layouts under one card shell
- 48px floating analog timeline clock as a signature element

## 2. Colors

A restrained warm palette. The pomegranate carries the brand; the warm whites carry the room.

### Primary

- **Pomegranate** (`#EE3423`): the warm point in the room. Used for the radial FAB, primary action buttons and the analog clock's second hand. Never used for content backgrounds, never as a gradient, never tinted onto more than 10% of any screen.
- **Pomegranate Hover** (`#D42D1E`): pressed/active state of Pomegranate. Used only as a momentary state, never at rest.
- **Pomegranate Light** (`#FF6B5B`): the dark-mode primary (slightly lighter for legibility on dark surfaces). Also used very sparingly as a soft background tint on text-moment cards.

### Neutral

- **Surface 1** (`#FFFFFF`): card and modal surfaces. The room's white walls.
- **Surface 2** (`#FAFAFA`): page background. Subtle warmth distinguishes from Surface 1.
- **Surface 3** (`#F5F5F5`): inset surfaces, search input backgrounds, skeleton bases.
- **Surface 4** (`#EBEBEB`): dividers, card borders, skeleton shimmer highlight.
- **Surface 5** (`#E0E0E0`): disabled control backgrounds.
- **Text 1** (`#1A1A1A`): primary content text. Never pure black.
- **Text 2** (`#6B6B6B`): timestamps, metadata, secondary captions.
- **Text 3** (`#9E9E9E`): placeholder text, disabled labels.
- **Text 4** (`#BDBDBD`): hint text, decorative dividers.

### Semantic

- **Success** (`#22C55E`): post-success toast accent.
- **Warning** (`#F59E0B`): soft-warning toast accent.
- **Error** (`#EF4444`): destructive confirm buttons, error toast accent, validation messages.

### Named Rules (Colors)

**The One Voice Rule.** Pomegranate appears on ≤10% of any rendered surface. Its rarity is the point. When two pomegranate elements are visible at once (e.g., the FAB and a primary button), neither is decorative — both are doing work.

**The No-Black Rule.** Pure black (`#000000`) is forbidden. Text uses `#1A1A1A`. Future palette refinements should tint neutrals toward the pomegranate hue (chroma 0.005–0.01 in OKLCH) rather than expanding the grey scale.

**The No-Counts Rule.** *(Added 2026-04-26; see [ADR-0004](docs/decisions/0004-no-counts-anywhere.md).)* No numeric counts appear anywhere in the rendered UI. No like count, no view count, no follower count, no read count, no reaction count, no comment count. Where we need to communicate "someone did this," we render names and avatars: *Maya*, never *3 people*. The limit of 150 is communicated in words, never as a meter or progress bar. This rule is universal across surfaces; there is no exception screen.

## 3. Typography

**Display Font:** Nunito (fallback: `-apple-system, BlinkMacSystemFont, sans-serif`)
**Body Font:** Platform default sans (`-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif`)

**Character:** Nunito is rounded, friendly, warm — the visual equivalent of the brand voice. It carries the logo and screen titles, where Nah wants to say "this is a place, not an app." Everything functional (body, labels, buttons, data, timestamps) uses the platform's own sans, so the app feels native on iOS and native on Android. The split is intentional: brand surfaces breathe with Nunito, working surfaces disappear into the OS.

### Hierarchy

- **Display** (Nunito, 30px, 700, 1.2, letter-spacing -0.01em): onboarding card titles, hero moments. Maximum two uses per screen.
- **Headline** (Nunito, 24px, 700, 1.2, letter-spacing -0.01em): screen titles, profile display names on the profile page only.
- **Title** (Nunito, 20px, 600, 1.3): section headers within a screen.
- **Body** (platform sans, 16px, 400, 1.5): moment text, card content, prose. Cap line length at 65–75ch where prose is long-form.
- **Label** (platform sans, 14px, 500, 1.4, letter-spacing 0.02em): button text, chip text, navigation.
- **Caption** (platform sans, 12px, 400, 1.4): timestamps, metadata, hint text.

### Named Rules (Typography)

**The Display-Off Rule.** Nunito never appears on functional UI elements. Buttons use Label. Form labels use Label. Data uses Body or smaller. Display fonts on buttons read as marketing copy, which Nah is not.

**The System-Body Rule.** Body text always uses the platform default sans, never Nunito. The trade is intentional: a small loss of cross-platform pixel-identity, a large gain in feeling-like-the-OS.

## 4. Elevation

Nah is flat by default. Cards rest on the page without shadow at rest. Depth is conveyed through surface tonality (Surface 1 over Surface 2) and 1px borders in `Surface 4`, not through ambient drop shadows. This is the garden's quiet: nothing floats unless it's doing something.

Shadows appear only as a response to state: an active radial FAB carries a subtle elevation, a bottom sheet casts a soft drop, a toast lifts slightly above the feed. None of these are decorative; each marks a thing that is happening *now*.

In dark mode, shadows are invisible on dark surfaces. Cards switch to 1px borders in `Surface 4` (dark variant) and rely entirely on tonal stepping between surface levels. This is the same architecture inverted, not a separate visual language.

### Shadow Vocabulary

- **FAB ambient** (`box-shadow: 0 4px 12px rgba(238, 52, 35, 0.15)`): the radial FAB's at-rest elevation. Tinted toward pomegranate so the shadow extends the warm point rather than darkening it.
- **Bottom sheet** (`box-shadow: 0 -4px 24px rgba(0, 0, 0, 0.08)`): the soft lift of the composer above the page.
- **Toast** (`box-shadow: 0 4px 16px rgba(0, 0, 0, 0.10)`): the brief lift of a toast notification above the feed.

### Named Rules (Elevation)

**The Flat-By-Default Rule.** Surfaces are flat at rest. Shadows appear only on active or transient elements (FAB, bottom sheet, toast, modal). A card never has a shadow. A button never has a shadow.

**The Border-Replaces-Shadow Rule.** In dark mode, every place that uses a shadow in light mode uses a 1px border instead. The rule cascades automatically; widgets check the theme brightness and swap.

## 5. Components

Warm and considered. Soft corners, generous padding, handmade-adjacent geometry. Spring physics on motion.

### Buttons

- **Shape:** pill (`rounded-full`, 9999px) for text buttons, circle for glyph buttons. Pill rather than rectangle because Nah? is a place, not a console.
- **Primary:** a 48px Pomegranate circle carrying a white 20px glyph, with its label beside it in Label typography and Text 1. No text is ever set on Pomegranate: white on `#EE3423` measures 4.08:1, under WCAG AA's 4.5:1 for text, while a glyph needs only 3:1. Pressed swaps the circle to Pomegranate Hover with a 150ms ease-out transition. The radial FAB and the send button are this component at their own sizes.
- **Ghost:** Surface 1 background, Text 1 text, same shape and padding. Used for secondary actions in composer and onboarding. Pressed state shifts background to Surface 2.
- **Destructive:** the primary shape with an Error circle, because white on `#EF4444` measures 3.76:1 and the same rule applies. Used only for "Discard" and "Leave Nah?" confirmations.
- **Focus:** 2px Pomegranate outline at 4px offset from the button edge. Visible only on keyboard focus.

### Floating Action Button (Radial FAB)

- **Size:** 56px circle.
- **Color:** Pomegranate background, white "+" glyph at 24px.
- **Position:** fixed bottom-right, 24px inset from screen edges. Left-hand mode mirrors to bottom-left.
- **At rest:** carries the FAB ambient shadow (pomegranate-tinted).
- **Active (menu open):** the "+" rotates 45° to "×" over 250ms, menu items fan outward with 30ms stagger and spring entrance (400ms total).

### Moment Cards

- **Corner Style:** 8px radius (`md`). Soft, considered, not razor-sharp.
- **Background:** Surface 1 default. Text moments may opt into a soft Pomegranate Light tint (≤8% saturation) as a background-color choice in the composer.
- **Shadow Strategy:** none at rest (flat-by-default rule).
- **Border:** none in light mode. 1px Surface 4 in dark mode.
- **Internal Padding:** zero for photo moments (edge-to-edge media). 24px for text moments. Custom for voice and music.
- **Header:** 40px avatar + display name (Title) + timestamp (Caption) in a horizontal row at top of card.

Four moment variants share this card shell ([ADR-0006](docs/decisions/0006-three-moment-types.md), amended 2026-09-15):

1. **Photo:** edge-to-edge media, no internal padding around the media block. Aspect-ratio-preserving. One photo per moment.
2. **Text:** padded 24px content area, Body typography sized up to 18px, optional warm Pomegranate Light tint background.
3. **Voice:** a recording of up to about sixty seconds with a waveform, played inline with a tap. Its visual treatment is not designed yet.
4. **Music:** two-column content area. 64px album-art square on the left, song title and artist in stacked Body and Caption on the right. Tapping opens the song in a music app the reader chooses.

### Chips

- **Style:** pill shape, 32px height, 12px horizontal padding, Label typography.
- **Default state:** Surface 2 background, Text 2 text.
- **Selected state:** Pomegranate background, white text. Smooth 200ms transition between states.
- **Disabled state:** Surface 5 background, Text 3 text.

### Reaction Button

- **Size:** 32px circle, inside a 44px tap area.
- **At rest:** Surface 1 background, Text 2 outline-heart icon.
- **Reacted:** the reaction the person gave replaces the outline heart, in full colour. Only the poster sees other people's reactions, by name and face, never as a count ([ADR-0018](docs/decisions/0018-reactions-the-poster-sees.md)).
- **Tap:** a picker slides in above the button with the five reactions, Smile, Wink, Sad, Wow and Love, in a horizontal row, with a 250ms spring entrance.
- **Illustrations:** Nah?'s own, not system emoji. Simple and rounded, legible at 24px, with Pomegranate as an accent where it helps.

### Bottom Sheet (Composer)

- **Shape:** rounded only at the top (`rounded-t-lg`, 12px), flat at the bottom edge of the screen.
- **Background:** Surface 1. Semi-transparent dim backdrop (not blur) behind the sheet.
- **Drag handle:** 36×4px pill in Surface 4, centered, 12px from the top edge.
- **Animation:** slides up from below with 400ms spring entrance, drag-to-dismiss past 1/3 swipe threshold.
- **Max height:** 90vh. Content scrolls internally if it exceeds.

### Floating Timeline Clock (signature component)

A 48px analog clock face fixed to the left edge of the viewport during scroll. Hour and minute hands smoothly animate to match the timestamp of the moment nearest the viewport's vertical center. Second hand is rendered in Pomegranate. Date text ("Mon, Mar 3") sits below the clock face in Caption type. Appears on scroll-start, fades after 2 seconds of scroll inactivity. Under `prefers-reduced-motion`, falls back to a static digital time display.

This is the single most distinctive Nah component. It is not optional.

### No Tab Bar

There is no tab bar. The feed is the whole screen, and your circle and you open from its header. The timeline clock and the radial FAB are the only chrome over the feed, and nothing carries a badge.

### Avatar

- **Sizes:** xs 24px, sm 32px, md 40px (card headers, your circle), lg 56px (compressed profile header), xl 80px (expanded profile header).
- **Shape:** circular.
- **Fallback:** initials on Pomegranate Light background, white text.

## 6. Do's and Don'ts

### Do

- **Do** keep Pomegranate at ≤10% of any rendered screen. It is a focal point, not a fill.
- **Do** use Nunito for the logo, screen titles, and onboarding headings. Use the platform default sans everywhere else.
- **Do** use 8px corner radius on cards, 12px on bottom sheets, full-radius pills on buttons and chips. Never sharp 90° corners on touchable surfaces.
- **Do** ship spring physics on every entrance, press, and dismiss. 400ms with overshoot for entrances, 250ms ease-out for exits.
- **Do** use the 48px floating timeline clock as the signature scroll companion. It is the single most recognizable Nah element.
- **Do** treat the limit of 150 as a designed boundary, never a limitation. If it is ever drawn, draw faces, never something that fills towards a number.
- **Do** keep the Feed chronological, full stop. The visual hierarchy must never imply algorithmic ranking, suggestions, or "you might like."
- **Do** use Pomegranate Light as a soft text-moment background tint, sparingly, opt-in.
- **Do** show empty states with warm illustrations, never blank screens. "Welcome home" framing on first run, "You're all caught up" framing at end of feed.

### Don't

- **Don't** look like Instagram, TikTok, or X. No Stories rail at the top. No suggested-content interleaving. No infinite-scroll engagement bait. No algorithmic feed under any circumstance.
- **Don't** look like LinkedIn or Threads. No follower counts. No public profiles. No professional-self-curation chrome. No mutual-connection paths shown publicly.
- **Don't** display numeric counts of anything, anywhere. No like count, no view count, no follower count, no read count, no reaction count, no comment count. Names and avatars instead. The No-Counts Rule is universal; there is no exception screen.
- **Don't** look like Discord or Slack. No aggressive unread counts as the primary navigation signal. No multi-server sidebars. No "typing" indicators in v1.
- **Don't** look like a VC-SaaS landing page. No gradient hero text. No three-feature card grids on any product surface. No hero-metric templates. No "X is faster than Y" comparison tables. No glassmorphism dashboards.
- **Don't** use side-stripe borders (colored `border-left` or `border-right` greater than 1px). Never. On any card, list item, callout, or alert.
- **Don't** use gradient text (`background-clip: text` with a gradient background). Solid colors only. Emphasis through weight or size.
- **Don't** use glassmorphism outside the two approved surfaces: profile header compression overlay and radial menu backdrop. Forbidden everywhere else, including modals, sheets, and notification toasts.
- **Don't** use display fonts (Nunito) on buttons, form labels, data, timestamps, or any other functional UI element. Display fonts on buttons read as marketing copy.
- **Don't** use pure black (`#000000`) or pure white (`#FFFFFF`) for text. Text uses `#1A1A1A`. Future iterations should tint neutrals toward the pomegranate hue.
- **Don't** use bounce or elastic easing curves. Spring physics are warm; bounce and elastic are loud.
- **Don't** add card grids of identical-size feature blocks. Each moment card variant has its own treatment; the Feed is varied by design.
- **Don't** introduce a modal as a first-thought solution. Exhaust inline, progressive, and bottom-sheet alternatives first. Bottom sheets are the Nah default for any "secondary surface" need.
