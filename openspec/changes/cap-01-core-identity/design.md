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

**Decision:** System fonts with humanist feel, no custom web fonts

**Scale:**

```css
/* Font family - native feel, fast load */
--font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;

/* Scale (mobile-first) */
--text-xs: 0.75rem;    /* 12px - timestamps */
--text-sm: 0.875rem;   /* 14px - secondary text */
--text-base: 1rem;     /* 16px - body */
--text-lg: 1.125rem;   /* 18px - card titles */
--text-xl: 1.25rem;    /* 20px - section headers */
--text-2xl: 1.5rem;    /* 24px - profile names */

/* Weights */
--font-normal: 400;
--font-medium: 500;    /* Titles, emphasis */
--font-semibold: 600;  /* Buttons, key UI */
```

**Rationale:** System fonts feel native and load instantly. San Francisco (iOS) and Roboto (Android) are humanist sans-serifs — friendly, readable, not cold. Custom fonts add load time and rarely justify the tradeoff for a PWA.

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

**Decision:** Bottom-left FAB with 7-item radial fan

**Specifications:**

- Position: Fixed, bottom-left corner (24px inset)
- Size: 56px diameter (touch-friendly)
- Color: Primary red with white "+" icon
- Shadow: Subtle elevation (0 4px 12px rgba(0,0,0,0.15))

**Menu items (clockwise from top):**

1. 📷 Photo
2. 🎬 Video
3. 📍 Location
4. 🎵 Music
5. 💬 Text
6. 👥 With (tag friends)
7. 🌙 Sleep/Wake

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
│ 😊 😢 😲  •  12 comments           │
└────────────────────────────────────┘
```

**Floating time indicator:**

- Appears on left edge during scroll
- Shows clock face with animated hands
- Date text below clock
- Fades in/out based on scroll velocity
- Smoothly interpolates between post timestamps

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

**Decision:** Emoticon sticker bubbles, not generic likes

**Reaction set:**

- 😊 Smile (default positive)
- 😉 Wink (strong like / inside joke)
- 😢 Sad (sympathy)
- 😲 Wow (surprised)
- ❤️ Love (deep appreciation)

**Interaction:**

- Tap heart icon → reaction picker appears
- Tap reaction → bubble animates to post
- Animation: Scale from 0 with spring, slight wobble
- Reactions show as row of small circles with friend avatars

**Display:**

- Show up to 3 unique reaction types as icons
- "+N more" for additional
- Tap to see full list with names

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

### 10. Onboarding & Modals

**Decision:** Stacked modal cards with consistent animation

**Welcome screen:**

- Full-screen primary red background
- Logo centered
- "Sign In" / "Sign Up" buttons (white, rounded)

**Onboarding cards:**

- Centered modal (max 85% width)
- Rounded corners (radius-lg)
- Enter: scale from 0.9 + fade, spring easing
- Exit: scale to 0.9 + fade, ease-out
- Stack effect: Previous card visible behind (smaller, dimmed)

**First-time hints:**

- Tooltip anchored to radial menu
- "Create a moment" with arrow
- Dismisses after first use or tap

## Risks / Trade-offs

| Decision | Trade-off |
|----------|-----------|
| System fonts | Less brand distinction, but faster load and native feel |
| Spring animations | More complex to implement, requires animation library |
| Floating time indicator | Scroll performance must be optimized (60fps) |
| Frosted glass | Not supported in all browsers, needs fallback |
| Custom reactions | More UI complexity than simple likes |

## Open Questions

1. **Dark mode:** Full dark theme or just respect system preference for now?
2. **Haptics:** Add subtle vibration feedback on mobile for key actions?
3. **Sound:** Any UI sounds (message sent, reaction added)?
4. **Accessibility:** How to handle motion-reduced preferences while keeping personality?

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
