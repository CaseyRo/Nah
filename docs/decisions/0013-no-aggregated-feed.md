---
title: "ADR-0013 — No aggregated feed. A circle is a room you enter"
permalink: /decisions/0013-no-aggregated-feed/
layout: page
status: Proposed
date: 2026-09-13
---

## Status

Proposed (2026-09-13). Corrects the client design assumed by [ADR-0009](0009-circles-not-one-circle.html), which described the app merging a person's circles into one view.

## Context

Once a person belongs to several circles, the obvious move is to merge them into a single stream. Every social application does this, and the client package was specified to do it.

It is the wrong move here, and it took proposing it out loud to see why.

## Decision

**There is no aggregated feed. A circle is a room you enter, and you are only ever in one at a time.**

The app opens on the circle you were last in, or on the list of your circles. Switching between them is a deliberate act, and it must feel instant. That is a caching problem rather than a merging problem.

Making switching fast, without merging:

- Each circle's recent moments are cached on the device, so entering a room renders from local storage first and refreshes behind it.
- A little prefetching in the background, so the room you are about to enter is already warm.
- A mark, never a number, showing that a circle has something new since you last looked. A dot is compatible with no counts. A badge that says four is not.

## Considered alternatives

- **One merged chronological feed across all circles.** Rejected. It recreates exactly the scroll this product exists to refuse, and it flattens audiences back into a single stream, which is the problem circles just solved. It would also mean one place holding every credential a person has, which the no-middle-tier decision already rejected for good reason.
- **A merged feed as an option, off by default.** Rejected. An option that contradicts the product's premise is still in the product, and defaults are not where this belongs. The same reasoning that rejected opt-in counts applies.
- **A digest screen summarising all circles.** Rejected as the merged feed wearing a hat.

## Consequences

Positive:

- The client gets simpler rather than harder. There is no cross-circle ordering, no dedup, no reconciling clocks between servers.
- "Who can see this?" stays answerable, because you are always somewhere specific when you post.
- Each circle keeps its own texture. The family room and the puppy-training room should not feel like one stream with labels.

Negative / acknowledged cost:

- Someone in six circles has to visit six rooms, and the app has to make that cheap or they will stop.
- There is no single place that answers "what happened today", which is a real convenience being given up on purpose.
- Prefetching several circles quietly costs battery and data, and needs to stay modest.

## What would prove this wrong

- People repeatedly ask for one view of everything, in those words.
- Switching circles is where people drop off in use, measured as sessions that open one room and never enter a second.
- The number of circles per person turns out to be small enough that the room metaphor is overhead rather than clarity.
