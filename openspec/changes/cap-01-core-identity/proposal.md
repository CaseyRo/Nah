# Core Identity Proposal

## Why

Nah needs a visual identity that immediately communicates "this is different from other social apps." The design must evoke the feeling of coming home — warm, intimate, personal. Path succeeded at this through careful attention to color, animation, and interaction design. We need to establish our design system foundation before building any features.

## What Changes

This is the first implementation capability. We're creating:

- **Design system foundation** — Color tokens, typography scale, spacing system, component primitives
- **Radial menu component** — The signature Path-inspired "+" menu with bloom animation
- **PWA shell** — App manifest, service worker, offline support, install prompts
- **Brand assets** — Logo, app icons, splash screens, favicon

## Capabilities

### New Capabilities

- `design-system`: Tailwind config, CSS custom properties, shadcn-svelte theme customization. Defines the warm color palette (coral/pink primary from mockup), typography (friendly, readable), spacing, and base component styles.

- `radial-menu`: The iconic radial content creation menu. Svelte component with Framer Motion (or svelte/motion) animations. Blooms outward on tap, contains icons for moment types.

- `pwa-shell`: SvelteKit PWA configuration. Web app manifest, service worker for offline caching, install prompt handling, splash screens for iOS/Android.

- `brand-assets`: Logo files (SVG, PNG at multiple sizes), app icons (192x192, 512x512), Apple touch icons, favicon, Open Graph images for social sharing.

### Modified Capabilities

*None — this is initial implementation.*

## Impact

### Code

- `/packages/ui` — New package for design system and shared components
- `/apps/web` — SvelteKit app shell with PWA configuration
- `/apps/web/src/lib/components` — Radial menu and core UI components

### Dependencies

- `tailwindcss` ^4.0
- `shadcn-svelte` (latest)
- `bits-ui` (shadcn-svelte dependency)
- `@sveltejs/kit` with adapter-static or adapter-auto
- `vite-plugin-pwa` or SvelteKit service worker

### Design Decisions Needed

- Final color palette (coral/pink from mockup as starting point)
- Typography choices (system fonts vs. custom web font)
- Animation library (svelte/motion, motion one, or CSS-only)
- Icon set (Lucide, Phosphor, or custom)
