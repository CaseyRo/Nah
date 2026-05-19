---
title: "ADR-0005 — Ritual onboarding"
permalink: /decisions/0005-ritual-onboarding/
layout: page
status: Accepted
date: 2026-04-26
---

## Status

Accepted (2026-04-26).

## Context

The MVP wow lands in the first 30 seconds (ADR-0008). The hard part is that the headline ("no counts anywhere") is a negative feature: nothing visible communicates it on a first-open empty screen. The user just sees... a quiet space. Maybe their inviter's name. Maybe nothing. Without intentional design, "no counts" reads as "broken Instagram."

So the onboarding has to do positive work: communicate the texture of the place, not just the absence of metrics.

The risk is overcorrection: ritual that performs deliberateness is worse than no ritual at all. "We made our login slow on purpose" reads great in a memo and infuriates people in practice.

## Decision

Four-beat ritual onboarding, ~60-90 seconds from install to in-app:

1. **One slow beat.** A single sentence on a calm background, fading in over 3-4 seconds. No skip button. The slowness is the message. Something like *"Nah? is for fewer people, less often."*
2. **One real question.** Asked seriously, answered seriously. Not "what's your name." Something the answer to which signals what this place is for. (Open question: what exactly. Candidate: *"Who are you, in two sentences, for the people who'll see you here?"*)
3. **Felt arrival.** A single soft haptic tap when you cross the threshold. No confetti. No "Welcome to Nah?!" banner. The haptic is the welcome.
4. **The quiet space.** You arrive. If you were invited, you see the inviter's last 1-2 moments (real or seeded). If those don't exist, you see the inviter's name and "It's quiet" with confidence.

The story of Nah? lives with the inviter, not the app. The inviter explains Nah? in their own voice ("I want us to share without strangers watching") in the chat thread they already share with the invitee. The app supplies the structural sentence on first open, not a marketing pitch.

## Considered alternatives

- **Standard social-app onboarding (email/password, profile photo, find your friends, allow notifications).** Rejected. Every step is a tax. None of them communicate what Nah? is, and most of them are tax-shaped in ways Nah? specifically refuses (find-friends asks for contact permissions, allow-notifications defaults to yes).
- **No onboarding at all (you tap the invite link and you're in).** Rejected. The empty room on first open is too austere. We need to land the texture of the place before showing the empty default state.
- **A multi-screen tutorial explaining the philosophy.** Rejected. Apps that over-explain don't trust the user. Trusting the user is itself a value statement; we keep it short and confident.
- **A single-screen onboarding (one combined screen, no slow beat).** Rejected. The slow beat is doing work that one screen can't: it signals "this app doesn't respect your urgency, because urgency is what we're against."

## Consequences

Positive:

- Communicates the texture of the place before any feature surface is visible.
- The "story" being delivered by the human inviter, not the app, is on-brand for the deliberate-sharing philosophy.
- Cheap to build (four screens, one haptic call, no auth-tier-onboarding flows like contact import).
- Sets the rhythm. A user who arrives slowly is more likely to use slowly.

Negative / acknowledged cost:

- Friends who agreed to install because Casey asked them, not because they're seeking a meditative app experience, may roll their eyes at the slow beat. Risk: it reads as pretentious. Mitigation: the four beats are short. One slow beat ≠ a slow app.
- The "one real question" is harder to design than the slow beat. If the question is wrong, the whole onboarding lands as gimmick. The question hasn't been finalized.
- Seeded posts on first open (ADR-pending, currently bundled here) require either real content from the inviter or mocked content attributed to them. The mock path has product-quality risks: a friend sees Casey "saying" something Casey didn't say and is confused.

## What would prove this wrong

- We ship and friends consistently skip past or speed-read the slow beat. (Concrete measure: at least 60% of new users spend the intended 3-4 seconds on the slow-beat screen.)
- The "one real question" gets dishonest answers ("fine," "idk") routinely, suggesting it's read as a hoop rather than a real prompt.
- The seeded mock posts confuse the invitee badly enough that they message the inviter to ask about the fake content.
