---
title: "ADR-0006 — Three moment types"
permalink: /decisions/0006-three-moment-types/
layout: page
status: Accepted
date: 2026-04-26
---

## Status

Accepted (2026-04-26). Acknowledged most-expensive MVP option.

## Context

A moment is the fundamental unit of Nah?. What types of moments can a user share? The earlier nah-vision spec listed five (photo/video, text, music, location, sleep/wake). The April session re-cut this down to MVP scale.

The question: which compose flows ship for MVP? Each type is a separate UX surface (composer, card variant, edit/delete affordances, edge cases), so each type is a real cost.

## Decision

Three moment types for MVP: **text, voice, photo**. User picks which to compose. Each type has a distinct composer flow and a distinct card variant.

- **Text.** Padded typographic card. Casual prose, not posts-with-headlines.
- **Voice.** Audio recording with a waveform render. Up to ~60 seconds. Plays inline with a tap-to-play.
- **Photo.** Edge-to-edge media card. Single photo (not carousel) in MVP. Aspect-ratio preserving.

Video, music share, location, and status are explicitly not in MVP.

## Considered alternatives

- **Voice-only MVP.** Rejected. The most opinionated cut, the most distinctive choice (no one else has voice as the primary medium). Strong philosophy fit because voice is hard to fake, hard to perform, hard to dash off. Rejected because Casey wanted three options to test which his actual circle would reach for; locking voice-only is a strong opinion masquerading as scope discipline.
- **Text-only MVP.** Rejected. Too austere. Doesn't differentiate from a group chat. Photo is the lowest-friction option for non-writers in the circle.
- **Five moment types (text, voice, photo, music share, location).** Rejected as scope. Music-share and location each carry meaningful implementation cost (Spotify/Apple Music integration for music; map tile rendering, geocoding, privacy precision controls for location). Both deferred without prejudice.
- **Rich-post (text + photo + voice attachments in a single moment).** Rejected. Mixes mediums and creates a "create-a-post" mental model identical to Instagram/Facebook. We want each moment to be one thing, said in one medium.

## Consequences

Positive:

- Three composer flows is enough variety to learn what the actual circle reaches for in real usage. Casey's circle of 5-10 friends provides real data.
- Each type can have a distinct visual treatment, which makes the feed feel varied (anti-card-grid-disease).
- Voice as a co-equal medium is a real departure from Instagram-shaped apps. Worth shipping for the differentiation even if usage skews to photo and text.

Negative / acknowledged cost:

- Triples the surface area vs. picking one type. Three composer designs, three card variants, three edge-case sets, three sets of error states.
- Voice has hard implementation problems: background recording on iOS without jank, waveform rendering at 60fps, offline-queue with audio blobs, transcript for accessibility.
- Three options up-front splits attention. A simpler "one composer, one medium" might land cleaner.

## What would prove this wrong

- Casey's MVP circle reaches for one type ≥80% of the time and barely uses the other two. (We learn what to cut for v2.)
- The voice composer ships at 60% polish because it carried the most implementation weight, and friends bounce off it specifically.
- A friend describes Nah? as "yet another text + photo app," which means the three-type cut didn't communicate variety even though it was technically present.
