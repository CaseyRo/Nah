# Core Identity Proposal

## Why

Nah needs a visual identity that immediately communicates "this is different from other social apps." The design must evoke the feeling of coming home — warm, intimate, personal. Path succeeded at this through careful attention to color, animation, and interaction design. We need to establish our design system foundation before building any features.

## What Changes

This is the first implementation capability. We're creating:

- **Design system foundation** — Color tokens, typography scale, spacing system, widget primitives
- **Radial menu component** — The signature Path-inspired "+" menu with bloom animation
- **App shell** — Flutter app entry point, navigation scaffold, offline-cache wiring, lifecycle handling, install prompts via App Store / Play Store
- **Brand assets** — Logo, app icons (iOS + Android), splash screens, marketing site favicon

## Capabilities

### New Capabilities

- `design-system`: Flutter `ThemeData` + custom Theme extensions, Dart token constants in `/packages/nah_ui`. Defines the warm color palette (pomegranate red primary), typography (Nunito display + system body), spacing, motion tokens, and base widget styles.

- `radial-menu`: The iconic radial content creation widget. Flutter widget with `AnimationController` + spring physics (e.g. `flutter_animate` or hand-rolled `SpringSimulation`). Blooms outward on tap, contains icons for moment types.

- `app-shell`: Flutter app shell — entry widget, route structure, `Bloc`/`Cubit` providers, offline cache initialization (Hive/Drift), push notification setup, deep-link handling, iOS + Android lifecycle.

- `brand-assets`: Logo files (SVG, PNG at multiple sizes), iOS app icon set (1024px master + auto-generated variants), Android adaptive icon (foreground + background layers), iOS/Android splash screens, marketing site favicon + Open Graph images.

### Modified Capabilities

*None — this is initial implementation.*

## Impact

### Code

- `/packages/nah_ui` — New Dart package for design system tokens and shared widgets
- `/apps/mobile` — Flutter app shell, entry point, routing
- `/apps/mobile/lib/widgets` — Radial menu and core UI widgets

### Dependencies (Flutter / Dart)

- `flutter_bloc` — state management
- `dio` — HTTP client for Mastodon REST API
- `web_socket_channel` — Mastodon Streaming API
- `hive` or `drift` — offline cache and queued moments
- `flutter_animate` (or hand-rolled `AnimationController` + `SpringSimulation`) — bloom and spring animations
- `flutter_local_notifications` + `firebase_messaging` (or APNs direct) — push notifications
- `geolocator`, `flutter_blue_plus` (later capabilities) — location, BLE proximity
- `google_fonts` (or self-hosted) — Nunito display font

### Design Decisions Needed

- Final color palette (pomegranate red #EE3423 confirmed from nah-vision)
- Animation approach (`flutter_animate` package vs. hand-rolled spring simulations)
- Icon set (Phosphor for Flutter, Lucide port, or custom-drawn SVG-to-widget)
