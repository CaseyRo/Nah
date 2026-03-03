# Design System Spec

Tailwind configuration, CSS custom properties, shadcn-svelte theme customization, and component primitives.

## ADDED Requirements

### Requirement: Color tokens defined as CSS custom properties

The design system SHALL define all colors as CSS custom properties, allowing runtime theming and dark mode support.

#### Scenario: Primary color usage

- **WHEN** a component uses the primary action color
- **THEN** it references `--color-primary` (#EE3423 pomegranate red)

#### Scenario: Dark mode colors

- **WHEN** the system is in dark mode
- **THEN** surface colors swap to dark variants via CSS custom properties

### Requirement: Brand typography — display font + system body

The design system SHALL use a single display font for the logo and headings, with system fonts for body/UI text.

**Display font:** Nunito (recommended) — or Outfit or Poppins as alternatives. Rounded, friendly, warm. Loaded via Google Fonts or self-hosted.

**Usage:**

```css
/* Display font — logo, headings, onboarding titles */
--font-display: 'Nunito', -apple-system, BlinkMacSystemFont, sans-serif;

/* Body/UI font — system stack for native feel and zero load time */
--font-body: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
```

**Where display font is used:**

- App logo / wordmark
- Screen titles (Feed, Friends, Profile)
- Onboarding card headings
- "You're all caught up" end-of-feed message
- Profile display name (on profile page only, not in cards)

**Where system font is used:**

- All body text, card content, timestamps, metadata
- Buttons, form inputs, navigation labels
- Comments, reaction labels

**Type scale tokens:**

```css
--text-xs: 0.75rem;    /* 12px - timestamps, metadata */
--text-sm: 0.875rem;   /* 14px - secondary text */
--text-base: 1rem;     /* 16px - body */
--text-lg: 1.125rem;   /* 18px - card titles */
--text-xl: 1.25rem;    /* 20px - section headers */
--text-2xl: 1.5rem;    /* 24px - profile names */
--text-3xl: 1.875rem;  /* 30px - onboarding headings (display font) */

--font-normal: 400;
--font-medium: 500;
--font-semibold: 600;
--font-bold: 700;      /* Display font headings only */

--letter-spacing-tight: -0.01em;   /* Display headings */
--letter-spacing-normal: 0;        /* Body text */
--letter-spacing-wide: 0.02em;     /* Uppercase labels, metadata */

--line-height-tight: 1.2;    /* Headings */
--line-height-normal: 1.5;   /* Body text */
--line-height-relaxed: 1.75; /* Long-form content */
```

#### Scenario: Body text rendering

- **WHEN** body text is displayed
- **THEN** it uses `--font-body` at `--text-base` (16px) with `--font-normal` weight

#### Scenario: Heading rendering

- **WHEN** a screen title or heading is displayed
- **THEN** it uses `--font-display` at `--text-xl` or larger with `--font-bold` weight and `--letter-spacing-tight`

#### Scenario: Display font loading

- **WHEN** the app loads
- **THEN** the display font is loaded with `font-display: swap` to prevent FOIT
- **AND** system font is shown until the display font loads

### Requirement: Spacing scale based on 4px unit

The design system SHALL define spacing tokens in 4px increments from space-1 (4px) to space-12 (48px).

#### Scenario: Component padding

- **WHEN** a card component has padding
- **THEN** it uses spacing tokens (e.g., --space-4 for 16px)

#### Scenario: Layout gaps

- **WHEN** elements are spaced in a layout
- **THEN** gaps use spacing tokens for consistency

### Requirement: Border radius tokens

The design system SHALL define border radius tokens: sm (6px), md (8px), lg (12px), and full (pill shape).

#### Scenario: Button corners

- **WHEN** a button is rendered
- **THEN** it uses --radius-sm or --radius-full depending on style

#### Scenario: Card corners

- **WHEN** a card is rendered
- **THEN** it uses --radius-md for subtle rounding

### Requirement: Animation timing tokens

The design system SHALL define animation duration and easing tokens, including spring curves for bouncy effects.

#### Scenario: Fast micro-interaction

- **WHEN** a hover state animates
- **THEN** it uses --duration-fast (150ms) with standard easing

#### Scenario: Bouncy entrance animation

- **WHEN** a modal or menu enters
- **THEN** it uses --duration-spring (500ms) with --ease-spring (overshoot curve)

### Requirement: Tailwind config exports design tokens

The Tailwind configuration SHALL extend the default theme with all custom tokens (colors, spacing, animation).

#### Scenario: Using tokens in Tailwind classes

- **WHEN** a developer writes `bg-primary` or `p-4`
- **THEN** it maps to the defined design tokens

#### Scenario: Custom animation classes

- **WHEN** a developer needs spring animation
- **THEN** they can use `animate-spring-in` class defined in Tailwind config

### Requirement: Dark mode token architecture

The design system SHALL define a complete dark mode palette scoped via `[data-theme="dark"]` or `prefers-color-scheme: dark`.

**Light palette (default):**

```css
:root {
  --color-primary: #EE3423;
  --color-primary-hover: #D42D1E;
  --color-primary-light: #FF6B5B;

  --color-surface-1: #FFFFFF;       /* Cards, modals */
  --color-surface-2: #FAFAFA;       /* Page background */
  --color-surface-3: #F5F5F5;       /* Inset areas, input backgrounds */
  --color-surface-4: #EBEBEB;       /* Dividers, borders */
  --color-surface-5: #E0E0E0;       /* Disabled backgrounds */

  --color-text-1: #1A1A1A;          /* Primary text */
  --color-text-2: #6B6B6B;          /* Secondary text */
  --color-text-3: #9E9E9E;          /* Placeholder, disabled */
  --color-text-4: #BDBDBD;          /* Hint text */
  --color-text-inverse: #FFFFFF;    /* Text on primary color */

  --color-success: #22C55E;
  --color-warning: #F59E0B;
  --color-error: #EF4444;
}
```

**Dark palette:**

```css
[data-theme="dark"] {
  --color-primary: #FF6B5B;          /* Slightly lighter red for dark bg legibility */
  --color-primary-hover: #FF8A7A;
  --color-primary-light: #EE3423;

  --color-surface-1: #1E1E1E;       /* Cards, modals */
  --color-surface-2: #141414;       /* Page background */
  --color-surface-3: #252525;       /* Inset areas */
  --color-surface-4: #333333;       /* Dividers, borders */
  --color-surface-5: #404040;       /* Disabled backgrounds */

  --color-text-1: #F5F5F5;          /* Primary text */
  --color-text-2: #ABABAB;          /* Secondary text */
  --color-text-3: #757575;          /* Placeholder, disabled */
  --color-text-4: #525252;          /* Hint text */
  --color-text-inverse: #1A1A1A;    /* Text on primary color */

  --color-success: #4ADE80;
  --color-warning: #FBBF24;
  --color-error: #F87171;
}
```

**Shadow strategy for dark mode:**

- Light mode: box-shadow for elevation (standard approach)
- Dark mode: shadows become invisible on dark surfaces. Replace with:
  - 1px borders using `--color-surface-4` for card edges
  - Subtle luminance differences between surface levels
  - No drop shadows on dark mode surfaces

#### Scenario: Automatic dark mode

- **WHEN** user's system preference is `prefers-color-scheme: dark`
- **THEN** `[data-theme="dark"]` is applied to the root element
- **AND** all color tokens resolve to dark palette values

#### Scenario: Shadow-to-border swap

- **WHEN** dark mode is active
- **THEN** card shadows are replaced with subtle borders: `border: 1px solid var(--color-surface-4)`

### Requirement: Component specs

The design system SHALL define the following component primitives:

**Avatar:**

| Size | Dimension | Usage |
|------|-----------|-------|
| xs | 24px | Reaction badges, inline mentions |
| sm | 32px | Comment authors, compact lists |
| md | 40px | Card headers, friend list items |
| lg | 56px | Profile page (compressed header) |
| xl | 80px | Profile page (expanded header) |

- Shape: circular (`border-radius: 9999px`)
- Online indicator: 10px green dot, bottom-right, with 2px white border
- Reaction badge: small custom reaction icon overlaid on bottom-right (replaces online dot contextually)
- Fallback: initials on primary-light background

**Card variants:**

- Base card: white surface, radius-md, subtle shadow (light) or 1px border (dark)
- Photo/Video card: edge-to-edge media, no content padding
- Text card: padded (space-6), optional warm tint background
- Music card: two-column layout (art left, info right)
- Location card: full-width map, overlaid venue name
- Status card: compact single-line, no content area expansion

**Bottom sheet:**

- Slides up from bottom with spring animation
- Max height: 90vh
- Drag handle: 36px wide, 4px tall, centered, radius-full
- Backdrop: semi-transparent overlay (not glassmorphism)
- Drag-to-dismiss: swipe down past threshold closes sheet

**Toast / Snackbar:**

- Position: bottom of screen, above tab bar, 16px inset
- Auto-dismiss: 5 seconds (configurable)
- Variants:
  - **Success:** green left accent bar + message
  - **Error:** red left accent bar + message + optional "Retry" action
  - **Info:** neutral accent bar + message
- Animation: slide up from bottom, spring easing
- Max 1 toast visible at a time (queue additional)

**Skeleton loader:**

- Animated shimmer (left-to-right gradient sweep)
- Matches layout of content it replaces (card skeleton, avatar skeleton, text line skeleton)
- Uses `--color-surface-3` as base, `--color-surface-4` as shimmer highlight
- Reduced-motion: static gray block (no animation)

**Empty state pattern:**

- Centered illustration (warm, friendly, brand-consistent)
- Heading in display font
- Body text explaining what goes here
- Primary CTA button
- Used for: zero-friends feed, no search results, empty profile, no notifications

**Badge:**

- Notification count: red circle, white text, positioned top-right of parent
- Supporter badge: small icon beside display name (design TBD in community-funding spec)
- Sizing: min 18px diameter, scales with count digits

**Chip:**

- Rounded pill shape (radius-full)
- Used for: privacy selector ("Close Friends" / "All Friends"), tags, filter controls
- States: default (outline), selected (filled primary), disabled (muted)
- Height: 32px, horizontal padding: space-3

### Requirement: Z-index scale

The design system SHALL define a z-index scale to prevent stacking context conflicts.

```css
--z-base: 0;
--z-card: 1;
--z-sticky: 10;         /* Sticky headers, tab bar */
--z-fab: 20;             /* Floating action button */
--z-overlay: 30;         /* Backdrop overlays */
--z-bottom-sheet: 40;    /* Bottom sheets */
--z-modal: 50;           /* Modal dialogs */
--z-toast: 60;           /* Toast notifications (always on top) */
--z-radial-menu: 55;     /* Radial menu (above modals, below toasts) */
```

#### Scenario: FAB above tab bar

- **WHEN** FAB and tab bar are both visible
- **THEN** FAB renders above tab bar (`--z-fab` > `--z-sticky`)

#### Scenario: Toast above everything

- **WHEN** a toast appears while a modal is open
- **THEN** toast renders above the modal (`--z-toast` > `--z-modal`)

### Requirement: Status indicator tokens

The design system SHALL define status indicator colors for presence states.

```css
--color-status-online: #22C55E;    /* Green dot */
--color-status-away: #F59E0B;      /* Amber dot */
--color-status-dnd: #EF4444;       /* Red dot */
--color-status-offline: #9E9E9E;   /* Gray dot */
```

#### Scenario: Online friend indicator

- **WHEN** a friend is currently active
- **THEN** their avatar shows a green online dot using `--color-status-online`

### Requirement: Consider oklch color definitions

For future iterations, the design system SHOULD consider defining colors in oklch color space for programmatic dark mode derivation and perceptually uniform color manipulation. This is a v2 consideration — v1 uses hex values as defined above.
