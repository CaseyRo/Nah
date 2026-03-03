# Core Identity Design

## Context

This document establishes the visual and interaction design foundation for Nah. We're not just building a social app — we're crafting a *feeling*. Path succeeded because every pixel, every animation, every micro-interaction said "someone cared about this." That's the bar.

The detailed Path analysis (Brian Lovin's dissection, techwalls reviews) gives us a blueprint. Our job is to honor that craft while making it our own.

## Goals / Non-Goals

**Goals:**

- Establish a complete design token system (colors, typography, spacing, motion)
- Define the signature interactions (radial menu, timeline clock, reactions)
- Create consistent animation language across all components
- Make the app feel "handcrafted, not generic"

**Non-Goals:**

- Pixel-perfect Path recreation (we take inspiration, not screenshots)
- Desktop-first design (mobile is primary, desktop is bonus)
- Feature implementation (this is foundation only)

## Decisions

### 1. Color System

**Decision:** Pomegranate red primary with warm neutrals

**Palette:**

```css
/* Primary */
--color-primary: #EE3423;        /* Pomegranate red - FAB, app bar, key actions */
--color-primary-hover: #D42D1E;  /* Darker for hover states */
--color-primary-light: #FF6B5B;  /* Lighter for backgrounds, highlights */

/* Surfaces */
--color-surface: #FFFFFF;        /* Cards, modals */
--color-surface-dim: #FAFAFA;    /* Subtle background variation */
--color-surface-dark: #1A1A1A;   /* Dark mode surfaces */

/* Text */
--color-text-primary: #1A1A1A;   /* Main text */
--color-text-secondary: #6B6B6B; /* Timestamps, metadata */
--color-text-inverse: #FFFFFF;   /* Text on primary color */

/* Semantic */
--color-success: #22C55E;
--color-warning: #F59E0B;
--color-error: #EF4444;
```

**Rationale:** Path's red (#EE3423) is warm, not aggressive. It says "energy" and "life" without screaming. The neutral palette keeps focus on content (photos, faces) while the red draws attention to actions.

### 2. Typography

**Decision:** Display font for brand identity + system fonts for body/UI

**Font families:**

```css
/* Display font — logo, headings, onboarding titles */
--font-display: 'Nunito', -apple-system, BlinkMacSystemFont, sans-serif;

/* Body/UI font — system stack, native feel, zero load time */
--font-body: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
```

**Scale:**

```css
/* Scale (mobile-first) */
--text-xs: 0.75rem;    /* 12px - timestamps */
--text-sm: 0.875rem;   /* 14px - secondary text */
--text-base: 1rem;     /* 16px - body */
--text-lg: 1.125rem;   /* 18px - card titles */
--text-xl: 1.25rem;    /* 20px - section headers */
--text-2xl: 1.5rem;    /* 24px - profile names */
--text-3xl: 1.875rem;  /* 30px - onboarding headings */

/* Weights */
--font-normal: 400;
--font-medium: 500;    /* Titles, emphasis */
--font-semibold: 600;  /* Buttons, key UI */
--font-bold: 700;      /* Display font headings only */

/* Spacing */
--letter-spacing-tight: -0.01em;   /* Display headings */
--letter-spacing-normal: 0;        /* Body text */
--letter-spacing-wide: 0.02em;     /* Uppercase labels */

--line-height-tight: 1.2;    /* Headings */
--line-height-normal: 1.5;   /* Body text */
--line-height-relaxed: 1.75; /* Long-form content */
```

**Rationale:** Nunito (rounded, friendly, warm) gives Nah brand distinction in headings and the logo without slowing down body text rendering. System fonts for body/UI text feel native and load instantly. The display font loads with `font-display: swap` — system font shows until it's ready. Alternatives considered: Outfit (geometric, modern), Poppins (geometric, popular). Nunito's rounded terminals best match the warm, "coming home" aesthetic.

### 3. Spacing & Layout

**Decision:** 4px base unit, generous whitespace

```css
--space-1: 0.25rem;   /* 4px */
--space-2: 0.5rem;    /* 8px */
--space-3: 0.75rem;   /* 12px */
--space-4: 1rem;      /* 16px - base padding */
--space-5: 1.25rem;   /* 20px */
--space-6: 1.5rem;    /* 24px */
--space-8: 2rem;      /* 32px */
--space-10: 2.5rem;   /* 40px */
--space-12: 3rem;     /* 48px */

/* Radii */
--radius-sm: 0.375rem;  /* 6px - small buttons */
--radius-md: 0.5rem;    /* 8px - cards */
--radius-lg: 0.75rem;   /* 12px - modals */
--radius-full: 9999px;  /* Pills, avatars */
```

**Rationale:** Generous spacing creates breathing room. Path felt "airy" despite being content-dense. The 4px base ensures everything aligns on a consistent grid.

### 4. Motion & Animation

**Decision:** Springy, bouncy physics with consistent easing

**Core principles:**

- Everything animates (nothing "pops" into existence)
- Bouncy spring physics for playful feel
- Consistent timing across similar interactions
- Animations serve purpose (direct attention, confirm action)

**Timing tokens:**

```css
/* Durations */
--duration-fast: 150ms;    /* Hovers, micro-feedback */
--duration-normal: 250ms;  /* Most transitions */
--duration-slow: 400ms;    /* Modal enter/exit */
--duration-spring: 500ms;  /* Bouncy animations */

/* Easings */
--ease-out: cubic-bezier(0.16, 1, 0.3, 1);      /* Decelerate in */
--ease-in-out: cubic-bezier(0.65, 0, 0.35, 1);  /* Smooth both */
--ease-spring: cubic-bezier(0.34, 1.56, 0.64, 1); /* Overshoot bounce */
```

**Key animations:**

| Element | Animation | Timing |
|---------|-----------|--------|
| Radial menu open | Fan out with rotation + scale | 400ms spring |
| Radial menu close | Collapse back with fade | 250ms ease-out |
| Modal enter | Scale from 0.95 + fade in | 300ms spring |
| Modal exit | Scale to 0.95 + fade out | 200ms ease-out |
| Reaction bubble | Pop in with overshoot | 300ms spring |
| Pull-to-refresh | Custom fluid animation | Gesture-driven |
| Timeline clock | Smooth interpolation | 60fps on scroll |

**Rationale:** Path's animations were "delightful" because they had *personality*. The spring easing creates that bouncy, alive feeling. Consistent timing makes the app feel cohesive — like one designer touched every interaction.

### 5. Radial Menu Design

**Decision:** Bottom-right FAB with 5-item radial fan, left-hand setting available

**Specifications:**

- Position: Fixed, bottom-right corner (24px inset), above bottom tab bar
- Left-hand setting: user preference moves FAB to bottom-left
- Size: 56px diameter (touch-friendly)
- Color: Primary red with white "+" icon
- Shadow: Subtle elevation (0 4px 12px rgba(0,0,0,0.15))

**Menu items (5 items, counter-clockwise from top):**

1. Photo (camera icon)
2. Text (quote/note icon)
3. Music (music note icon)
4. Location (pin icon)
5. Status (moon/sun icon — covers sleep/wake)

**Removed from radial menu:**

- "With" (friend tagging) → moved to composer toolbar option on any moment type
- "Video" → combined with Photo (camera icon opens media picker for both)
- Sleep/Wake → combined into "Status" (single entry, picker inside composer)

**Animation sequence:**

1. "+" rotates 45° to become "×"
2. Items fan out with staggered delay (30ms each)
3. Each item scales from 0 with spring physics
4. Slight rotation on each item creates "bloom" effect
5. Backdrop dims slightly (optional)

**Dismissal:**

- Tap "×" button
- Tap outside menu
- Tap any item (then transition to composer)

### 6. Timeline Design

**Decision:** Full-width cards with left-edge time axis

**Card anatomy:**

```text
┌────────────────────────────────────┐
│ [Avatar] Name                  3h  │
├────────────────────────────────────┤
│                                    │
│         [Content Area]             │
│    (photo/text/music/location)     │
│                                    │
├────────────────────────────────────┤
│ [heart] [reactions]  •  12 comments│
└────────────────────────────────────┘
```

- `[heart]` = dedicated reaction button (tap to open picker)
- `[reactions]` = row of custom reaction icons with friend avatars
- `12 comments` = threaded comment count (tappable, opens comment sheet)

**Moment type card variants:**

Each moment type has a distinct card layout within the shared card shell (header + footer stay consistent):

| Moment Type | Content Area Treatment |
|-------------|----------------------|
| **Photo/Video** | Edge-to-edge media, no padding. Minimal chrome — let the image breathe. Video shows play button overlay. |
| **Text** | Padded content area (space-6). Typographic treatment: larger text size (--text-lg), medium weight. Background can be a subtle warm tint or user-selected color. |
| **Music** | Album art thumbnail (left) + song title, artist, album (right). Waveform or audio visualizer accent below. Optional: tappable to open in music service. |
| **Location** | Map thumbnail (static, from map tiles API) spanning full width. Venue/city name overlaid at bottom with semi-transparent backdrop. |
| **Status (sleep/wake)** | Minimal single-line with icon. Moon icon for sleep, sun icon for wake. Timestamp prominent. No content area padding — compact card. |

**Floating timeline clock:**

The timeline clock is a **prominent** design element, not a subtle indicator. It's one of Nah's signature differentiators from other social apps.

- **Position:** Fixed on left edge of viewport during scroll
- **Size:** 48px diameter clock face — large enough to read at a glance
- **Appearance:** Analog clock face with hour and minute hands, pomegranate red second hand
- **Date text:** Below clock, showing "Mon, Mar 3" format
- **Behavior:**
  - Appears when user begins scrolling
  - Clock hands smoothly animate to match the timestamp of the nearest visible moment
  - Date text updates as user scrolls through days
  - Fades out after 2 seconds of scroll inactivity
  - Smoothly interpolates between post timestamps (not jumpy)
- **Reduced-motion fallback:** Static digital time display (e.g., "2:30 PM") instead of animated clock hands. Still shows on left edge during scroll.

**End of feed:**

- Subtle ornamental illustration (low contrast)
- "You're all caught up" message
- Gentle, non-intrusive

### 7. Profile & Cover Photo

**Decision:** Facebook Timeline-style header with scroll compression

**Behavior:**

1. Full cover photo visible at top (16:9 aspect)
2. Profile photo overlaps cover (bottom-center)
3. Name and bio below profile photo
4. As user scrolls down:
   - Cover photo parallax scrolls faster
   - Profile photo shrinks and moves to top-left
   - Header compresses to minimal "Me" bar
   - Smooth, physics-based transition

**Frosted glass effect:**

- Used on overlays when cover photo shows through
- CSS: `backdrop-filter: blur(20px) saturate(180%)`
- Fallback: Semi-transparent solid color

### 8. Reactions System

**Decision:** Dedicated heart/reaction icon button on every card, custom illustrated icons

Each moment card has a dedicated heart icon button in the card footer — this is the primary reaction entry point. More discoverable and accessible than long-press alone.

**Reaction set (custom SVG icons, part of brand-assets deliverable):**

- **Smile** — default positive reaction
- **Wink** — strong like, inside joke energy
- **Sad** — sympathy, support
- **Wow** — surprise, amazement
- **Love** — deep appreciation, heart-eyes style

**Icon style:**

- Simple, rounded, friendly illustration style
- Consistent stroke weight with app's icon language
- Pomegranate red as accent color where appropriate
- Work at small sizes (24px) without losing clarity

**Interaction:**

- Tap heart icon on card footer → reaction picker appears above the button
- Double-tap card → quick "Love" reaction (shortcut, skips picker)
- Long-press card → reaction picker appears (alternative path)
- Tap reaction → bubble animates to post with spring + slight wobble
- Reactions show as row of small custom icons with friend avatars

**Display:**

- Show up to 3 unique reaction types as custom icons
- "+N more" for additional
- Tap to see full list with names and avatars

### 9. Composer Flow

**Decision:** Bottom-sheet modal with consistent structure

**Layout:**

```text
┌─────────────────────────────────────┐
│ Cancel          [Type]        Post  │
├─────────────────────────────────────┤
│                                     │
│          [Content Input]            │
│    (text field / media preview)     │
│                                     │
├─────────────────────────────────────┤
│ 🔒 Close friends  📍 Location  ...  │
└─────────────────────────────────────┘
```

**Photo composer additions:**

- Filter strip at bottom (horizontal scroll)
- Filters shown as thumbnails with preview

**Animations:**

- Slides up from bottom with spring
- Content area fades in after sheet settles
- Privacy toggles animate as chips

### 10. Onboarding Flow

**Decision:** Email + password authentication, 4-step onboarding, warm empty state

**Authentication model:** Email + password. Simple, self-hosted. No OAuth/social login for v1.

**Step 1: Welcome screen**

- Full-screen primary red background
- Nah logo centered (display font)
- Tagline: "Not Alone Here"
- "Sign In" / "Create Account" buttons (white, rounded)

**Step 2: Account creation**

- Email address input
- Password input (with strength indicator)
- Display name input
- Submit → email verification sent

**Step 3: Profile setup (post-verification)**

- Upload profile photo (camera/gallery picker, with crop)
- Optional: add a short bio
- Skip option available (can complete later)

**Step 4: "How Nah Works" explainer**

- 3 swipeable cards (stacked modal style):
  1. "Share real moments with your closest people" — illustration of moment types
  2. "150 friends max — quality over quantity" — illustration of friend circle
  3. "No algorithms, no ads, just your friends" — illustration of chronological feed
- "Get Started" button on final card

**Zero-friends empty state:**

- Warm illustration (not a blank screen)
- Heading: "Welcome home"
- Body: "Nah is better with friends. Invite someone you care about."
- Primary CTA: "Invite a Friend" button
- Secondary: "Explore your profile" link
- The radial FAB is visible but has a pulsing tooltip: "Share your first moment"

**Returning user (sign in):**

- Email + password fields
- "Forgot password" link
- Error state: "Email or password incorrect" (never reveal which)

**Modal animation (reused throughout):**

- Centered modal (max 85% width)
- Rounded corners (radius-lg)
- Enter: scale from 0.9 + fade, spring easing
- Exit: scale to 0.9 + fade, ease-out
- Stack effect: Previous card visible behind (smaller, dimmed)

### 11. Primary Navigation Model

**Decision:** Bottom tab bar with 3 tabs (v1), FAB above tabs

**Layout:**

```text
┌─────────────────────────────────────┐
│                                     │
│          [Screen Content]           │
│                                     │
│                              [FAB]  │
├─────────────────────────────────────┤
│   Feed     Friends     Profile      │
└─────────────────────────────────────┘
```

**Tabs (v1):**

1. **Feed** — home timeline (chronological moments from friends)
2. **Friends** — friend list, pending requests, invite flow
3. **Profile** — own profile/journal, settings access

**v1.5 addition:** Messages tab inserted between Friends and Profile.

**Tab bar behavior:**

- Fixed at bottom of viewport
- Active tab: pomegranate red icon + label
- Inactive tabs: secondary text color icon + label
- Badge indicators: red dot for pending friend requests, count badge for unread notifications
- Tab bar hides on scroll down, shows on scroll up (reclaim screen space)

**FAB position relative to tabs:**

- FAB sits 16px above the tab bar, bottom-right (or bottom-left in left-hand mode)
- FAB (closed) z-index is above tab bar but below modals/sheets; radial menu (open) renders above modals but below toasts
- No hamburger menu — everything is accessible via tabs + profile settings

### 12. Error Handling Patterns

**Decision:** Consistent error patterns across all user-facing flows

**Composer errors:**

- **Discard confirmation:** If user has entered content and taps Cancel/back, show confirmation dialog: "Discard this moment?" with "Keep Editing" (primary) and "Discard" (destructive)
- **Loading state:** Post button shows spinner, input is disabled during submission
- **Submission failure:** Toast notification: "Couldn't post your moment. Tap to retry." Moment stays in composer (not lost)
- **Media upload failure:** Inline error on the media thumbnail: "Upload failed" with retry icon

**Network errors:**

- **Toast pattern:** Bottom-of-screen toast, above tab bar, auto-dismiss after 5s
- **Retry:** Toast includes "Retry" action button for failed operations
- **Persistent offline:** Banner at top of screen: "You're offline. Viewing cached content." Banner dismisses automatically when connection restores.

**Friend limit error:**

- **WHEN** user attempts to add 151st friend
- **THEN** show friendly dialog: "You've reached 150 friends — Nah's limit for meaningful connections. To add someone new, you might consider reviewing your friend list."
- Link to friend curation tools (inactive friend suggestions)

**Offline queue:**

- **WHEN** user creates a moment while offline
- **THEN** moment is queued locally with a pending indicator (clock icon overlay)
- **AND** user can cancel queued moments before they sync
- **WHEN** connection restores
- **THEN** queued moments auto-submit with subtle success toast

### 13. Accessibility & Reduced Motion

**Decision:** Full `prefers-reduced-motion` support, WCAG 2.1 AA target

**Reduced motion strategy:**

- **WHEN** `prefers-reduced-motion: reduce` is set
- **THEN** all spring animations are replaced with instant transitions or simple opacity fades
- **AND** radial menu items appear simultaneously (no stagger)
- **AND** timeline clock shows static time display instead of animated hands
- **AND** pull-to-refresh uses simple spinner instead of fluid animation
- **AND** reaction bubbles appear without bounce/wobble

**Keyboard & screen reader:**

- All interactive elements are focusable and operable via keyboard
- Radial menu: ARIA roles defined in radial-menu spec
- Tab bar: `role="tablist"` with `role="tab"` children
- Moment cards: `role="article"` with semantic heading structure
- Reaction picker: `role="listbox"` with `role="option"` children

**View threshold definition:**

- A "view" is registered when 50% of the moment card is visible in the viewport for 2+ continuous seconds
- Scrolling past quickly does not trigger a view
- This threshold applies to view receipt generation (for users who have receipts enabled)

## Risks / Trade-offs

| Decision | Trade-off |
|----------|-----------|
| System fonts (body) + display font (headings) | Adds one font load, but gives brand distinction where it matters |
| Spring animations | More complex to implement, requires animation library; reduced-motion fallback required |
| Floating time indicator | Scroll performance must be optimized (60fps); static fallback for reduced-motion |
| Frosted glass (2 contexts only) | Not supported in all browsers, needs fallback; scoped to minimize performance impact |
| Custom reaction icons | More design work than system emoji, but cross-platform consistency |
| Email + password auth | Simple but no social login convenience; good enough for v1 |

## Resolved Questions

1. **Dark mode:** Respects system preference. Light is primary identity. Full dark token architecture in design-system spec.
2. **Haptics:** Deferred to post-v1. Focus on visual/motion design first.
3. **Sound:** Deferred to post-v1. No UI sounds in v1.
4. **Accessibility:** `prefers-reduced-motion` strategy defined above. WCAG 2.1 AA target.

---

## Implementation Notes

### Tailwind Config

```javascript
// tailwind.config.js (excerpt)
export default {
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#EE3423',
          hover: '#D42D1E',
          light: '#FF6B5B',
        },
      },
      animation: {
        'spring-in': 'springIn 400ms var(--ease-spring)',
        'spring-out': 'springOut 250ms var(--ease-out)',
        'bounce-in': 'bounceIn 300ms var(--ease-spring)',
      },
    },
  },
}
```

### Animation Library

Recommend **Motion One** or **Svelte Motion** for:

- Spring physics
- Gesture handling (radial menu, pull-to-refresh)
- Scroll-linked animations (timeline clock)

CSS-only won't achieve the spring physics we need.

---

*This design system is the soul of Nah. Every animation, every color choice, every spacing decision contributes to the feeling of "coming home." Shortcuts here will be felt by users, even if they can't articulate why.*
