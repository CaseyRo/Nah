# Core identity tasks

Specification only this round. Nothing below is started.

## 1. Design system

- [ ] 1.1 Light and dark themes in `apps/mobile` from the values in `DESIGN.md`, following the system appearance
- [ ] 1.2 Bundle Nunito as an asset; platform sans for everything that is not brand voice
- [ ] 1.3 Spring timings from `DESIGN.md`, with an instant or fade path when reduce motion is on
- [ ] 1.4 One glyph-button component for primary actions (Pomegranate circle, white glyph, label beside it), used by the +, send and join; check both palettes against WCAG AA

## 2. Radial menu

- [ ] 2.1 The + fixed bottom-right on the feed, mirrored by a left-hand setting
- [ ] 2.2 Bloom of text, voice, photo and music; each opens its composer with no step in between
- [ ] 2.3 Close by the close mark, a tap outside, or system back
- [ ] 2.4 Screen-reader names, expanded state, focus into and back out of the menu; 44 point targets

## 3. App shell

- [ ] 3.1 Replace the app bar with a feed header that opens your circle and you; no tab bar, no badge (the words already say "your circle", since CDI-1896)
- [ ] 3.2 Load on launch, foreground and pull to refresh; keep the last feed on the device and show it offline
- [ ] 3.3 Display name "Nah?" on both platforms; native splash on the brand background; system bars follow the theme
- [ ] 3.4 Explain before asking for notification permission; a tapped notification opens the feed
- [ ] 3.5 A CI check that fails if an analytics, crash-reporting, advertising, attribution or AI package enters the app's dependency tree (CDI-1860)
- [x] 3.6 Keep `works.cdit.nah` as the bundle identifier (ruled 2026-09-15)

## 4. Brand assets

- [ ] 4.1 Logo, as an SVG master (Casey)
- [ ] 4.2 iOS icon set and Android adaptive icon from the master
- [ ] 4.3 Site favicon and one default link-preview image
- [ ] 4.4 Check that an invite link's preview shows no person

## 5. Verification

- [ ] 5.1 VoiceOver and TalkBack pass over the feed, the + menu and a composer
- [ ] 5.2 Cold start to a scrollable stored feed within 2 seconds on iPhone 13 and Pixel 6 class devices
- [ ] 5.3 Capture a full session's traffic and confirm every request goes to a Nah? address
