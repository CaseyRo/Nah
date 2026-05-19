---
title: "ADR-0004 — No counts anywhere"
permalink: /decisions/0004-no-counts-anywhere/
layout: page
status: Accepted
date: 2026-04-26
---

## Status

Accepted (2026-04-26). The MVP headline. Sentence-describable to a friend.

## Context

"What makes Nah? amazing and fresh in the first 30 seconds, in one sentence?" The candidates were haptics-as-ritual, voice-first, daily windows, home-screen widgets, slow-loading-by-design, and no-counts. Of those, no-counts was the only one that's structural (not gimmicky), visible (you notice the absence immediately), philosophical (it's the thing Instagram needs to live), and free to implement (it's mostly *not* shipping features).

## Decision

**No visible counts of anything, anywhere.** No like count. No view count. No follower count. No reaction count. No read count. No daily-active-friend count. No streak count. The numbers don't exist in the UI.

Where we genuinely need to communicate "someone did this," we use names and faces: *Maya reacted*, not *3 people reacted*. *Wife and 2 close friends have read this* never; either *Wife has read this* or nothing.

The soft 150 cap (ADR-0003) follows the same rule: communicated as language ("your circle is full"), never as a meter.

## Considered alternatives

- **Selective counts (likes off, but show "read by N friends").** Rejected. Once any number exists, attention starts to optimize for it. Read counts become a performance metric just as cleanly as like counts. The discipline is the rule's totality.
- **Counts visible only to the poster, not viewers.** Rejected. Even private counts shape posting behavior. Casey checks his own posts to see "did anyone see this?" — and that's the same gravitational pull as public counts.
- **Counts hidden by default, opt-in to see.** Rejected. Opt-in is the wrong default direction. The architecture should make the right thing easy and the wrong thing intentional, not the inverse. Also: opt-in counts that "the design-conscious users disable" become a signal of in-group savvy, which is its own kind of performance.
- **Status quo: ship some counts because they're standard.** Rejected. The headline disappears.

## Consequences

Positive:

- The sentence-describable difference from Instagram lands instantly. Friends describe Nah? to other friends in one sentence: "no counts of anything, anywhere."
- Architecturally simpler: we don't ship a counters service. We don't think about caching counts, eventual consistency on counts, or migration paths when counts get bigger than 32-bit ints.
- Removes an entire class of design temptations downstream. Engagement-optimization decisions become much harder to even *propose* when there's nothing to optimize for.

Negative / acknowledged cost:

- "No counts" is a negative feature. The first 30 seconds of an app where things are absent doesn't feel amazing on its own — it feels empty. ADR-0005 (ritual onboarding) and the seeded-posts mechanic exist to fill the space the absence creates.
- We probably *will* track engagement signals server-side (for the daily digest, for soft-cap warnings, for debugging). Discipline: they exist for the system, never for the user.
- Some standard features (sorting feeds by popularity, "trending," recommendations) become structurally impossible. That's intended, but it removes options if the product needs to flex later.

## What would prove this wrong

- We ship MVP and the absence of counts is invisible — friends don't notice or describe it as the thing that's different. That means the headline doesn't actually do work and we picked the wrong sentence.
- A friend bounces because they "need to know who actually saw this" for a moment that mattered to them, and the absence of a read indicator feels like the app withholding information.
- The seeded-posts and ritual-onboarding mechanics (the positive features the absence is meant to make space for) don't carry the wow on their own, and the experience flattens.
