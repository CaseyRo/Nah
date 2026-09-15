---
title: "ADR-0018 — Reactions the poster sees, and nobody counts"
permalink: /decisions/0018-reactions-the-poster-sees/
layout: page
status: Accepted
date: 2026-09-15
---

## Status

Accepted (2026-09-15). Amends [ADR-0008](0008-mvp-scope.html), which shipped no reactions. Resolves CDI-1856. Specified in the `cap-04-reactions` OpenSpec change.

## Context

ADR-0008 left reactions out of the first version entirely. The retention research (`docs/research/what-makes-people-come-back.md`) calls that the most expensive rule in the product: in a small network, a moment met with silence is how the person who posts stops, and everyone else drifts out after them. CDI-1856 asked for the question to be settled in a decision record before anything is built, because it touches the headline of [ADR-0004](0004-no-counts-anywhere.html): no counts anywhere.

The January design had five illustrated reactions shown to everyone with avatars, a double-tap shortcut, view receipts on by default, and a "who viewed" list with a count. Since then [ADR-0017](0017-one-network-of-a-hundred-and-fifty.html) has made each person's circle private, so a reaction shown to others would also reveal who is connected to whom.

## Decision

**A person can react to a moment with one of five illustrated reactions, Smile, Wink, Sad, Wow and Love, and only the person who posted the moment sees them, by name and face, never as a number.**

- One reaction per person per moment, which they can change or take back.
- The illustrations are Nah?'s own, not system emoji, so they look the same on every phone.
- A reaction is sealed like a moment: the server knows that someone reacted and when, not which reaction.
- Nothing records who saw a moment. View receipts stay out, as ADR-0008 had them.

## Considered alternatives

- **One quiet mark that says "received".** The smallest thing that ends the silence, and what CDI-1856 proposed. Rejected because it can say that a moment arrived but not how it landed, and a sad moment and a happy one deserve different answers.
- **Reactions shown to everyone who sees the moment.** Rejected. It gives the circle something to compare and perform for, and a name on someone else's moment tells readers who else is in that person's circle, which nobody may see.
- **Likes, or any count of reactions.** Rejected by ADR-0004.
- **Nothing, as ADR-0008 decided.** Rejected on the research above.
- **View receipts, or a list of who saw a moment.** Rejected. It is the quiet version of a count, and ADR-0008 rules out visible read tracking.

## Consequences

Positive:

- A poster hears back, which the research says keeps a small network alive.
- "No counts anywhere" survives: forty reactions are forty faces, never "40".
- Nobody can perform for an audience of reactions, because there is none.

Negative / acknowledged cost:

- Five illustrations have to be designed, and have to stay legible at 24px.
- The server learns that one person reacted to another's moment. It already knows the two are connected.
- Without a visible crowd there is less pull to react. That is deliberate.
- Silence still happens. Reactions make it rarer, not impossible.

## What would prove this wrong

- People who post in their first weeks still stop at the rate the research predicts for silence, even though reactions exist.
- People repeatedly ask to see how others reacted to someone else's moment. That is the audience asking to exist.
- Reactions start being reported in words as a tally, "everyone loved it", which would mean the count found its way back without a number.
