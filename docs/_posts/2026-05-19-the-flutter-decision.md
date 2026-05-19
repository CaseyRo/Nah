---
title: "The Flutter Decision (or: how I picked the worst of three)"
date: 2026-05-19 12:00:00 +0100
categories: [Architecture, Decisions]
tags: [flutter, pwa, native, tech-stack, decision-log]
---

I've decided Nah will be built in Flutter.

I want to write this down honestly, because the path to that decision was *not* a triumphant march. It was three months of circling the same question, two AI assistants giving me opposite answers from the same audio capture, and a quiet realization that I'd been avoiding the choice because none of the options felt good.

## A history with all three

Before I tell you why Flutter, let me tell you why I kept *not* picking Flutter.

**PWA never felt good.** I've shipped PWAs. I've installed PWAs on my iPhone. Every time, there's a moment — usually the second or third launch — where the app does something slightly off. The status bar is the wrong colour. A pull-to-refresh fires when I didn't mean it. A back gesture closes the whole thing. Safari's storage quota eats my cached timeline after a week of inactivity. iOS push notifications require gymnastics that don't exist on Android. EU users on iOS can't even get push at all under DMA. PWA is *fine*, but Nah is supposed to feel like coming home, and PWAs always feel like crashing on a friend's couch.

**Native is too expensive.** I'm one person with a finite life. Building two parallel apps — Swift for iOS, Kotlin for Android, two codebases, two release cycles, two ecosystems to learn deeply, two sets of bugs to chase — is not a project I can ship. It's a project I can *begin*. Native would give me the best experience, and it would also give me a half-finished iOS app and a never-started Android app eighteen months from now. I know this about myself.

**Flutter sounds and feels bloated.** Every time I open a Flutter app on my phone, I can *tell*. The animations are smooth, almost too smooth — the kind of smooth that says "we're rendering our own UI, not yours." The app is 47MB to do what a Swift app does in 8MB. The text rendering is subtly wrong. The keyboard feels imported. Flutter ships its own engine, which is brilliant engineering and also exactly what makes the apps feel like Flutter apps, not iPhone apps or Android apps.

**Worst of three?** That's where I was. PWA: hollow. Native: impossible. Flutter: bloated. And the project sat.

## The April 25 fork

On April 25, I recorded an audio capture while thinking out loud. I was driving, probably, listing every question I had: BLE proximity for nearby discovery, location services, music sharing, the OneSec-inspired "deliberate friction" idea, business-card exchange like the old Pokken games. I was asking myself, *if Flutter, does the SDK actually cover this?*

That capture got processed by two different AI systems on the same day.

One of them — the Things one — read me carefully and said: *start with a Flutter spike. Validate the BLE and location story, then decide.* That advice matched what I was actually asking.

The other one — the SiYuan PWA capability matrix — read the iOS Safari constraints and said: *lock PWA for MVP, defer native to Phase 2.* It produced a beautiful table. It made tidy recommendations. It was completely confident.

I read the SiYuan doc and the confident framing made me think the decision was already made. I let it sit. I didn't ship anything. The repo stayed on its March-3 Svelte+SvelteKit framing — itself a decision I'd made *months* before the April thinking, and never re-examined.

Today I caught it. The two AI outputs from the same source audio reached opposite conclusions, and I'd been letting the louder one win without checking which one was right.

## Why Flutter wins on a re-read

Here's what changed when I went back to the actual content of the audio instead of the AI summary of it.

I want **BLE proximity**. Not as a nice-to-have. As a USP. The "people near you can exchange a moment" thing — that's the kind of feature that turns a social app into a *place*. PWAs can't do it on iOS. They can't even pretend. Web Bluetooth is Chrome-only, Android-only. If I go PWA, that feature is dead, not deferred. It's gone.

I want **deliberate-sharing native feel**. OneSec-style friction. Haptic confirmation when you complete a friction gate. Real launch from the home screen, real lock-screen notifications, real share-sheet integration both directions. Some of this PWA does badly. Some of this PWA doesn't do at all. Flutter does all of it, in one codebase, with plugins that already exist.

I want **to actually ship**. Native is two codebases. I will not finish two codebases.

The bloat criticism stays. Flutter apps are bigger than native, and they look like Flutter apps if you know what you're looking for. But: Nah is being installed by 150 people who already know the host. They're not browsing the App Store comparing it to other apps. The 30MB cost buys me a feature set that PWA can't deliver and a codebase that native won't let me finish.

That's the trade. Worst of three, picked deliberately.

## What I checked about Flutter in 2026

Because the bloat criticism was the one that made me hesitate, I went looking for the case *against* Flutter today.

- **Maintenance:** Flutter 3.41 / Dart 3.11 shipped in February as a "maturity milestone" — 825 commits, mostly stability. Google committed to four stable releases per year, 18-month LTS, 99.9% SLA. Google Pay and Google Ads ride on it. It is not Wave 2 of Google's killed-project graveyard. ([State of Flutter 2026](https://devnewsletter.com/p/state-of-flutter-2026/))
- **Cross-platform share:** ~46% of cross-platform devs use Flutter vs ~38% on React Native (Statista 2026). 32,000+ packages on pub.dev. Not dying. ([Flutter vs React Native 2026](https://tech-insider.org/flutter-vs-react-native-2026/))
- **Bundle size:** real. Empty app ~12-15MB IPA; real app with plugins lands 30-50MB on iOS. Native Swift lands ~5-15MB. PWA: ~1-2MB. Worst case for Nah: 50MB. I'll take it. ([Flutter app size benchmarks](https://medium.com/@expertappdevs/how-to-reduce-app-bundle-size-in-react-native-and-flutter-without-losing-features-2026-guide-018a4062a1b4))
- **BLE on iOS:** `flutter_blue_plus`, `universal_ble`, `flutter_splendid_ble` cover BLE Central. iBeacon proximity specifically still needs platform channels into CoreLocation — meaning some Swift. Less Swift than native. More Swift than zero. ([Flutter Bluetooth packages](https://fluttergems.dev/bluetooth-nfc-beacon/))
- **Hiring:** harder than React Native (6-8 week fill vs 3-4 week). Doesn't apply to a solo project, but I'm noting it for the day Nah becomes more than me. ([Why Big Tech Won't Hire Flutter Developers in 2026](https://dev.to/eira-wexford/why-big-tech-wont-hire-flutter-developers-2i88))
- **Performance:** Impeller 2.0 (current rendering engine) does 60-120fps. Path-style spring physics, radial menu blooms, the floating timeline clock — Flutter is genuinely good here. ([Flutter vs React Native 2026 architecture](https://www.bolderapps.com/blog-posts/flutter-vs-react-native-in-2026-why-the-new-architecture-and-impeller-2-0-changed-everything))

## What this changes in the repo

The README still says Svelte 5 + SvelteKit. The nah-vision design doc still lists "Mobile-first PWA" as Decision #2. The cap-01-core-identity change has a whole spec called `pwa-shell` with manifest scenarios and service worker requirements. The cap-01 task list has a section labeled "4. PWA Configuration".

All of that is wrong now. I'll be updating the specs in a separate pass — the architecture is the same in spirit (Mastodon backend, Dunbar-150 enforcement, custom reactions, light-primary identity, all of it stands), but the client is a Flutter app, not a SvelteKit PWA. The `pwa-shell` spec becomes an `app-shell` spec. The repo structure shifts: `/apps/web` becomes `/apps/mobile`. Tailwind tokens become Flutter ThemeData. shadcn-svelte primitives become Flutter widgets.

The backend doesn't move. Mastodon doesn't care what's talking to it.

## What didn't change

The soul of Nah. 150 friends maximum. Chronological feed. No algorithms. Mutual friendships. Custom reaction icons, not emoji. Pomegranate red on warm white. The radial "+" menu. The big analog timeline clock. The "coming home" feeling.

I'm not changing the destination. I'm changing the car.

## A note on the back-and-forth

If you've followed Nah's blog so far, you've watched me write decisive-sounding posts about every other decision. "Light mode is the primary identity." "Messaging is officially v1.5." "Glassmorphism scoped to exactly two places."

This one isn't decisive-sounding. This one was hard.

I think it's worth saying that out loud, because building in public means letting people see the parts where you wobble. The choice between PWA, native, and Flutter is genuinely a choice between three imperfect options for a small private social app. None of them was obviously right. The PWA AI thought PWA was right. The Things AI thought Flutter was right. Both were defending their conclusions confidently. I had to actually read the original audio capture, remember what I'd asked, and pick.

The lesson — for me, for anyone building with AI tools — is that "the doc with the tidy table" isn't the same as "the right answer." Decisive presentation feels like decisiveness. It isn't. I let a confident document make a decision for me for three weeks. I won't do that again.

Picking Flutter is picking the worst of three. I just think it's the worst that ships.

nah!
