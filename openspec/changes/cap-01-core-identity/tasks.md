# Core Identity Tasks

## 1. Project Setup

- [ ] 1.1 Create `/packages/nah_ui` Dart package (pubspec, lib/ scaffolding)
- [ ] 1.2 Initialize Flutter app in `/apps/mobile` (Flutter 3.41+, Dart 3.11+)
- [ ] 1.3 Add core deps: `flutter_bloc`, `dio`, `web_socket_channel`, `hive` (or `drift`)
- [ ] 1.4 Add animation + UI deps: `flutter_animate`, `google_fonts`, icon package
- [ ] 1.5 Configure iOS + Android build targets (bundle IDs, signing placeholders, min OS versions)

## 2. Design System Foundation

- [ ] 2.1 Define color tokens as Dart `const Color` values in `nah_ui/lib/tokens/colors.dart`
- [ ] 2.2 Build `nahLightTheme` and `nahDarkTheme` as Flutter `ThemeData` instances
- [ ] 2.3 Define typography scale via `TextTheme` (xs/sm/base/lg/xl/2xl/3xl)
- [ ] 2.4 Add custom Theme extension for app-specific tokens (motion curves, radii, status colors)
- [ ] 2.5 Define spring curves as Dart constants (`Curves.elasticOut`-style configs)
- [ ] 2.6 Build reusable spring animation helpers (`SpringIn`, `BounceIn` widget wrappers)
- [ ] 2.7 Document token usage in `nah_ui/README.md` with code samples

## 3. Radial Menu Component

- [ ] 3.1 Create `RadialMenu` widget shell in `nah_ui/lib/widgets/radial_menu/`
- [ ] 3.2 Implement FAB trigger button (56px, fixed bottom-right via `Positioned`)
- [ ] 3.3 Create `RadialMenuItem` widget with icon + label
- [ ] 3.4 Implement fan layout positioning (polar coordinates math in build())
- [ ] 3.5 Add spring `AnimationController` on open (staggered item entry via `Interval`)
- [ ] 3.6 Animate trigger glyph ("+" to "×") with rotation
- [ ] 3.7 Implement close behaviors (X tap, barrier tap, item select)
- [ ] 3.8 Add press feedback on menu items (scale, opacity via `AnimatedScale`)
- [ ] 3.9 Ensure 44px minimum touch targets (`InkWell` / `GestureDetector` size)
- [ ] 3.10 Test on iOS Simulator, Android emulator, and physical devices

## 4. App Shell Configuration

- [ ] 4.1 Configure iOS `Info.plist` (display name, status bar style, splash screen)
- [ ] 4.2 Configure Android `AndroidManifest.xml` (label, theme, launch mode)
- [ ] 4.3 Set up iOS app icon set (1024px master + auto-generated variants via `flutter_launcher_icons`)
- [ ] 4.4 Set up Android adaptive icon (foreground + background layers)
- [ ] 4.5 Configure iOS LaunchScreen.storyboard with brand colors
- [ ] 4.6 Configure Android `splash_screen.xml` with brand colors
- [ ] 4.7 Initialize offline cache (Hive box or Drift database) in app bootstrap
- [ ] 4.8 Wire up `Bloc`/`Cubit` providers at app root
- [ ] 4.9 Set up app lifecycle observer (`WidgetsBindingObserver`) for foreground/background
- [ ] 4.10 Add deep-link handling scaffold (`go_router` or platform `MethodChannel`)
- [ ] 4.11 Verify cold-start and warm-start performance targets (< 2s cold)

## 5. Brand Assets

- [ ] 5.1 Finalize logo design (wordmark "nah" or symbol)
- [ ] 5.2 Export logo as SVG
- [ ] 5.3 Generate PNG versions at required sizes
- [ ] 5.4 Create iOS app icon master (1024x1024)
- [ ] 5.5 Create Android adaptive icon layers (foreground + background SVG/PNG)
- [ ] 5.6 Generate iOS icon variants via `flutter_launcher_icons` (or Xcode asset catalog)
- [ ] 5.7 Create marketing site favicon (32x32, 16x16, ICO) and Open Graph image (1200x630)
- [ ] 5.8 Create iOS launch image set + Android splash drawable
- [ ] 5.9 Verify color consistency across all assets
- [ ] 5.10 Add app-bundled assets to `/apps/mobile/assets/` and configure `pubspec.yaml`

## 6. Integration & Testing

- [ ] 6.1 Wire up radial menu to placeholder composer routes
- [ ] 6.2 Test design tokens in light mode (`Brightness.light`)
- [ ] 6.3 Test design tokens in dark mode (`Brightness.dark`)
- [ ] 6.4 Verify all animations run at 60-120fps (Flutter DevTools performance overlay)
- [ ] 6.5 Test on low-end Android device (target API 24+)
- [ ] 6.6 Build release IPA/AAB and confirm bundle sizes within budget (< 50MB iOS, < 25MB Android base)
- [ ] 6.7 Verify accessibility (TalkBack on Android, VoiceOver on iOS, semantic widgets)

## 7. Documentation

- [ ] 7.1 Document design token usage in README
- [ ] 7.2 Create component usage examples
- [ ] 7.3 Update project README with setup instructions
