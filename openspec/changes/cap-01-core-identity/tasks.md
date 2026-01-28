# Core Identity Tasks

## 1. Project Setup

- [ ] 1.1 Initialize `/packages/ui` with pnpm workspace
- [ ] 1.2 Set up SvelteKit app in `/apps/web`
- [ ] 1.3 Install and configure Tailwind CSS v4
- [ ] 1.4 Install and configure shadcn-svelte
- [ ] 1.5 Set up Motion One (or svelte-motion) for animations

## 2. Design System Foundation

- [ ] 2.1 Create CSS custom properties file with all color tokens
- [ ] 2.2 Configure Tailwind theme extension (colors, spacing, radii)
- [ ] 2.3 Define typography scale in Tailwind config
- [ ] 2.4 Create animation keyframes and timing tokens
- [ ] 2.5 Define spring easing curves (--ease-spring)
- [ ] 2.6 Create `animate-spring-in`, `animate-bounce-in` utility classes
- [ ] 2.7 Document design tokens in Storybook or similar

## 3. Radial Menu Component

- [ ] 3.1 Create RadialMenu.svelte component shell
- [ ] 3.2 Implement FAB trigger button (56px, fixed bottom-left)
- [ ] 3.3 Create menu item components with icons
- [ ] 3.4 Implement fan layout positioning (radial coordinates)
- [ ] 3.5 Add spring animation on open (staggered items)
- [ ] 3.6 Add rotation animation on trigger ("+" to "×")
- [ ] 3.7 Implement close behaviors (X tap, outside tap, item select)
- [ ] 3.8 Add press feedback on menu items (scale, dim)
- [ ] 3.9 Ensure 44px minimum touch targets
- [ ] 3.10 Test on iOS Safari and Chrome Android

## 4. PWA Configuration

- [ ] 4.1 Create manifest.json with app identity
- [ ] 4.2 Configure SvelteKit adapter for PWA
- [ ] 4.3 Set up service worker (vite-plugin-pwa or custom)
- [ ] 4.4 Implement cache-first strategy for static assets
- [ ] 4.5 Implement network-first strategy for API calls
- [ ] 4.6 Add offline fallback page
- [ ] 4.7 Configure iOS splash screen meta tags
- [ ] 4.8 Configure Android theme-color meta tag
- [ ] 4.9 Test "Add to Home Screen" on iOS
- [ ] 4.10 Test installation prompt on Android
- [ ] 4.11 Verify offline mode works with cached content

## 5. Brand Assets

- [ ] 5.1 Finalize logo design (wordmark "nah" or symbol)
- [ ] 5.2 Export logo as SVG
- [ ] 5.3 Generate PNG versions at required sizes
- [ ] 5.4 Create app icons (192x192, 512x512)
- [ ] 5.5 Create Apple touch icon (180x180)
- [ ] 5.6 Create favicon (32x32, 16x16, ICO)
- [ ] 5.7 Create Open Graph image (1200x630)
- [ ] 5.8 Create iOS splash screen images
- [ ] 5.9 Verify color consistency across all assets
- [ ] 5.10 Add all assets to `/apps/web/static`

## 6. Integration & Testing

- [ ] 6.1 Wire up radial menu to placeholder composer views
- [ ] 6.2 Test design tokens in light mode
- [ ] 6.3 Test design tokens in dark mode (if implemented)
- [ ] 6.4 Verify all animations run at 60fps
- [ ] 6.5 Test on low-end Android device
- [ ] 6.6 Run Lighthouse PWA audit (target 90+)
- [ ] 6.7 Verify accessibility (keyboard nav, screen reader)

## 7. Documentation

- [ ] 7.1 Document design token usage in README
- [ ] 7.2 Create component usage examples
- [ ] 7.3 Update project README with setup instructions
