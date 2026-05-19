---
title: "ADR-0001 — Flutter over PWA"
permalink: /decisions/0001-flutter-over-pwa/
layout: page
status: Accepted
date: 2026-04-26
---

## Status

Accepted (2026-04-26). Repo materially pivoted 2026-05-19. See also the [Flutter decision blog post]({{ '/2026/05/19/the-flutter-decision/' | relative_url }}).

## Context

The original Nah? framing was SvelteKit + Mastodon, shipped as a PWA. That choice was made in early 2026 when the question "Native vs Flutter vs PWA" (CDI-775 on Linear) was still open. PWA was picked as the act-small default.

In an April session, the question came back. The trigger was a voice memo listing features that PWA can't do: BLE proximity, business-card exchange, music share-in, full programmable haptics, OS-level integration. Plus the underlying anxiety about iOS Safari being the bottleneck for everything.

The PWA capability matrix was filed first, showing exactly which features survive iOS Safari and which don't. The honest read was: BLE proximity, background location, Web Share Target (inbound), and full haptic patterns are genuinely off the table. EU iOS users lose push entirely under DMA.

Then the Flutter data came in. Flutter 3.41 supports iOS 13+ (2019) and Android API 24 (Android 7.0, 2016). Compare to PWA push, which requires iOS 16.4+ (March 2023) and is disabled in the EU. Flutter actually reaches further back than feature-complete PWA on iOS.

## Decision

Flutter from day one. One Dart codebase, iOS + Android, native install via the App Store and Play Store. Mastodon backend stays.

## Considered alternatives

- **Stay PWA (the previous decision).** Rejected. Kills BLE proximity (a future signature feature), kills EU iOS push, kills home-screen widgets and Live Activities (the OS-level integration that fits the "deliberate sharing" philosophy). The features the PWA matrix put in the "defer" bucket are the same features Casey kept returning to in audio captures.
- **Native iOS (Swift) + Android (Kotlin).** Rejected. Best UX, two codebases. Solo project will not finish two parallel apps. Three to five times the maintenance burden of a Flutter codebase for the same surface area.
- **React Native.** Considered. Flutter wins on rendering consistency (Impeller 2.0) and animation performance for the Path-inspired motion language. JS ecosystem advantage is irrelevant when the team is one Dart-friendly developer. Hiring is easier with RN, but solo project, so the hiring factor is moot.
- **Flutter Web (PWA via Flutter).** Rejected as a primary target. Flutter Web exists and is improving but isn't how you ship a polished mobile experience yet. Web/desktop reconsidered for v2+.

## Consequences

Positive:

- BLE proximity becomes a real Phase 2 feature, not a defer-forever dream.
- Reaches iOS 13+ and Android 7+, wider than feature-complete PWA.
- App Store and Play Store presence (trust signal for early users).
- OS-level integration (widgets, Live Activities, Focus modes, Action Button, CoreHaptics) becomes available, even if we don't use it all in MVP.
- Single Dart codebase; ships to two platforms with one effort.

Negative / acknowledged cost:

- Bundle size: ~30-50MB iOS IPA vs ~5-15MB native, ~1-2MB PWA.
- $99/year Apple + $25 one-time Google. PWA is $0.
- Cannot push updates without store review (Shorebird closes most of this gap, but it's a paid service).
- Hiring is harder than RN/Web (6-8 week fill vs 3-4 week). Doesn't bite at solo scale; revisit later.
- iOS BLE proximity (specifically iBeacon) requires CoreLocation via platform channels, so a small amount of Swift is unavoidable.

## What would prove this wrong

- We ship MVP and the bundle size becomes a barrier to invitation acceptance (a friend installs and bounces because "too big"). Concrete trigger: >10% of invited friends fail to complete first-open after install.
- Flutter's Google-commitment posture changes materially. Fuchsia-shaped sunset patterns, leadership shakeups, or release cadence slowdown beyond two cycles in a row would force a revisit.
- The OS-level features (widgets, Live Activities, Focus integration) prove to be features nobody actually uses, in which case the Flutter premium over PWA didn't buy what we thought.
