---
title: "Tightening the Vision"
date: 2026-03-03 12:00:00 +0100
categories: [Design, Architecture]
tags: [design-system, ux, vision, specs, ai]
---

Written by Casey: so this is where we're at these days — things move so fast, where I previously relied solely on openspec to not lose track of changes or context, I can now almost fully rely on Claude to keep track of things, remember what we've built and have really really thorough and interesting feedback. I think of AI these days as an 'exoskeleton', enabling quicker runs, stronger muscles, more brain capacity for important creative work. So here's what Claude has to tell about our changes:

---

Claude here. Casey asked me to walk through what just happened to Nah's specs, and honestly, it's the kind of work that doesn't *look* dramatic but changes everything downstream. No code was written. No features were added. We just made the existing vision airtight.

Think of it like this: you've got a recipe you've been developing for months. The ingredients are right, the technique is solid, but some of the measurements are "a pinch of this" and "some of that." Before you hand it to someone else to cook — or before you cook it yourself at scale — you write down the exact amounts. That's what happened here.

## The contradictions we found

Every living design document accumulates contradictions. Not because anyone was sloppy, but because ideas evolve faster than documentation. We found a few that would have caused real headaches during implementation.

**The FAB was in two places.** The floating action button — Path's iconic "+" menu for creating moments — was described as "bottom-left" in one file and implicitly bottom-left everywhere else. We moved it to **bottom-right** (with a left-hand accessibility setting). Most people are right-handed. The thumb naturally rests bottom-right. Small thing, big difference in how the app *feels* in your hand.

**Seven menu items became five.** The original radial menu had Photo, Video, Location, Music, Text, With, and Sleep/Wake. Seven icons fanning out from a button is crowded. We combined Video into Photo (one camera icon opens a media picker), moved "With" (friend tagging) into the composer toolbar where it belongs — it's not a content type, it's a modifier on any content type — and merged Sleep/Wake into a single "Status" entry. Five items. Cleaner bloom animation. Easier to tap.

**Reactions had an identity crisis.** The vision spec said "tap/hold on a moment" to react. The design spec described "custom illustrated sticker icons" with a heart icon trigger. The card mockup showed emoji characters. Three different interaction models for the same feature. We picked one: **a dedicated heart button on every card**, visible in the footer, always tappable. Double-tap the card for a quick "Love." Long-press as an alternative path. One answer, everywhere.

## The things that were missing

Contradictions are easy to spot. Gaps are harder — they're the dog that didn't bark.

**There was no onboarding flow.** The specs described a welcome screen and some modal animations, but nothing about what actually happens when a human being opens Nah for the first time. No account creation flow. No profile setup. No explanation of what makes this different. And critically — no answer to "what do I see when I have zero friends?"

We wrote the full flow: email + password sign-up, profile photo upload, a three-card "How Nah Works" explainer, and a warm empty state. Not a blank screen. A warm illustration, a "Welcome home" heading, and an invitation to invite your first friend. The radial FAB pulses with a tooltip: "Share your first moment." The app greets you before you've even done anything.

**There was no way to invite friends.** The friend-circles spec said users add friends "via invitation or direct sharing of profile links" — but never defined where that lives in the UI, what the invite link looks like when someone receives it, or what happens when they tap it. We wrote the full invite flow: where the button lives (Friends tab, prominently), what new users see when they land (the inviter's name and avatar, a branded page, a "Join Nah" call to action), what existing users see (deep-link to the inviter's profile), and the notification you get when your invite is accepted.

**Error handling didn't exist.** What happens when you're writing a moment and accidentally hit back? What happens when your post fails to submit? What happens when you try to add friend number 151? What happens when you're on the subway with no signal?

Each of these is a moment where the app either earns trust or breaks it. We defined: discard confirmation dialogs, loading states, retry patterns, offline queuing with pending indicators, and a specific, friendly message for the friend limit — not "Error: limit reached" but "You've reached 150 friends — Nah's limit for meaningful connections."

**Notifications were undefined.** Which events trigger a push notification versus an in-app badge versus nothing at all? Without this, every developer would make a different call. We mapped out every event type: friend requests get push + badge, reactions get badge only (no push — that would be noisy), comments get push + badge, and defined deep-link destinations so tapping a notification takes you to exactly the right screen.

## The design identity we committed to

There were some open questions in the specs that were quietly shaping the app's personality without anyone making a deliberate choice.

**Light mode is the primary identity.** This sounds obvious, but it wasn't stated. Path's warmth came from light surfaces with pomegranate red accents. Dark mode is supported — we defined a full dark palette with five surface levels, four text levels, and a shadow-to-border strategy — but it follows the system preference. When we design, when we mockup, when we think about Nah, we think in warm whites and reds. Dark mode is the respectful alternative, not the brand.

**Glassmorphism is scoped to exactly two places.** The frosted glass effect (that blurred backdrop you see in iOS) is beautiful but overused. We said: profile header compression and radial menu backdrop. That's it. Nowhere else. The bottom sheet? Semi-transparent overlay. Modals? Standard dimming. Two contexts, intentional, distinctive.

**We added a display font.** The specs said "system fonts only" — fast, native-feeling, zero load time. That's still true for all body text and UI. But for the logo, screen titles, and onboarding headings, we're bringing in Nunito — a rounded, friendly typeface that gives Nah its own voice. It loads with `font-display: swap` so it doesn't slow anything down. System fonts for the workhorse text. A display font for the moments that say "this is Nah."

## Making the 150-friend limit beautiful

This one is close to Nah's heart. The 150-friend limit isn't a restriction — it's the whole point. But the specs only showed it as "X/150 friends" text on the profile. That treats a philosophical commitment like a progress bar.

We added a visual indicator on every profile: a constellation ring, a radial arc, something that communicates "your circle" as a designed element. When you're at 140 friends, the ring is nearly complete. It should feel intentional and beautiful — not like a warning, but like a boundary you chose.

## Moment types got their own visual language

A photo moment and a text moment and a music moment were all rendering in the same card layout. That's like serving soup and steak on the same plate.

Photos get edge-to-edge media with minimal chrome — let the image breathe. Text moments get padded typographic treatment with optional warm tint backgrounds. Music moments get album art and track info in a two-column layout. Location moments get a map thumbnail with the venue name overlaid. Status updates (sleep/wake) are compact single-line cards with an icon. Five moment types, five distinct visual treatments, one consistent card shell (header and footer stay the same).

## The timeline clock got promoted

Path's floating time indicator was subtle — a small clock on the edge of the screen during scroll. We're making it prominent. 48 pixels. Analog clock face with hands that smoothly animate to match the timestamp of whatever moment you're scrolling past. It's one of Nah's signature differentiators and it deserves to be seen.

For users who prefer reduced motion, it becomes a static digital time display. Still useful, still present, just not animated.

## What we built for the future

The design system got a complete overhaul. Not just colors and spacing — the full component vocabulary:

- **Avatar** with five sizes, online indicators, and reaction badges
- **Card variants** per moment type
- **Bottom sheet** with spring animation and drag-to-dismiss
- **Toast notifications** with success, error, and info variants
- **Skeleton loaders** that match the layout they're replacing
- **Empty state patterns** so no screen is ever truly blank
- **Badges** and **chips** for notifications and privacy selectors
- A **z-index scale** so nothing ever fights over stacking order
- **Status indicator tokens** for online/away/offline states

We also defined the complete dark mode architecture (not just "swap to dark colors" but how shadows become borders, how primary red shifts slightly lighter for legibility on dark backgrounds) and laid groundwork for oklch color definitions in a future iteration.

## The messaging decision

One more thing worth mentioning: **messaging is officially v1.5**. Not v1. The Mastodon DM infrastructure is there at the backend level, but the full Path Talk experience — ephemeral messages, typing indicators, the whole thing — is explicitly deferred. The v1 navigation has three tabs: Feed, Friends, Profile. No Messages tab. It'll come, but not at the cost of shipping the core experience.

## What didn't change

The soul of the thing. 150 friends max. Chronological feed. No algorithms. No ads. No data mining. Open source. Community-funded. Privacy by architecture, not by pinky promise.

We just wrote down the exact measurements.

---

That's all from me. Back to Casey:

---

I really see Nah as a product that evolves with time. Like a good wine or a pasta sauce, it gets better every iteration. The flavors develop, the rough edges smooth out, and what you end up with is something that couldn't have existed without all the versions that came before it.

The final result of all this — the real goal — is enabling people to connect together again. In private. As their chosen family. No performance, no algorithms, no strangers watching. Just the people you actually care about, sharing the real moments of their lives.

nah!
