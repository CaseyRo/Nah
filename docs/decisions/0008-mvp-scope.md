---
title: "ADR-0008 — MVP scope"
permalink: /decisions/0008-mvp-scope/
layout: page
status: Accepted
date: 2026-04-26
---

## Status

Accepted (2026-04-26).

## Context

"MVP" is one of the most misused words in product. It often means "everything we could think of, minus the things that broke last sprint." For Nah?, MVP needs to be radically constrained, because:

- The audience is small (5-10 close friends, not a public beta).
- The headline (no counts) is a negative feature that needs space to land.
- Three-type composer (ADR-0006) is already the most expensive MVP option in its dimension.
- We pivoted to Flutter (ADR-0001), which is more setup cost than PWA. The compensation is shipping less in MVP, not more.

## Decision

**MVP audience:** Casey + 5-10 close friends. Real usage. Not a public beta. Not optimized for stranger-conversion. Not pitch-ready for investors.

**MVP capabilities** (the five-capability cut, drafted as cap-mvp):

1. **Onboarding** — Slow beat, one real question, felt arrival, quiet space arrival. Seeded posts from inviter (real or mocked).
2. **Circle** — Invitation IS connection. Soft cap as language. No discovery, no search, no public anything.
3. **Moments** — Three types: text, voice, photo. Immutable after post (no edit). One medium per moment.
4. **Feed** — Chronological. Quiet space empty state. No infinite scroll. No read tracking visible. Floating timeline clock as the only navigation chrome.
5. **Notifications** — Daily digest default. Vague-quantity wording. Zero-moments suppression. Opt-IN to more.

**Not in MVP** (explicit list):

- No reactions (no like, no emoji, no Path-style five-icon set)
- No comments / replies / threading
- No DMs / messaging
- No public profiles
- No discovery / search / "people you may know"
- No federation
- No stories
- No drafts (Compose-or-post, no in-between state)
- No edit-after-post (moments are immutable)
- No counts of anything anywhere (ADR-0004 makes this universal)
- No federation
- No widgets, Live Activities, Focus integration, or Action Button shortcuts (Phase 2)
- No touch-to-connect BLE (Phase 2)
- No music share, location, status — i.e. the cap-XX scaffolds in nah-vision/specs/ beyond the first five (Phase 2 or later)

## Considered alternatives

- **Wider MVP (everything in nah-vision).** Rejected. Eight capabilities × three moment types × full reactions × messaging would be 6-12 months of solo work before testing with the actual friends. The point of the small audience is to test fast.
- **Narrower MVP (text-only, one moment type, two friends).** Considered. Strong opinion. Rejected because the three-type cut is what tells us *what people reach for* — and we don't yet know.
- **No explicit not-in-MVP list.** Rejected. Without the list, scope creeps invisibly. The list is the contract.

## Consequences

Positive:

- Realistic build time: 2-3 months focused. Not the 6+ months the wider nah-vision would have taken.
- The actual circle gets to use Nah? while it's small enough to feel hand-built, which is the whole point.
- The not-in-MVP list is recoverable. Each item is a future cap-XX change. Nothing is rejected forever; everything is deferred deliberately.
- Less to commit means less to revisit if the headline doesn't land.

Negative / acknowledged cost:

- Friends who try Nah? at MVP will notice missing features (no reactions, no comments). The pitch has to land on what *is* there, not despite what isn't.
- The explicit not-in-MVP list creates expectation management overhead. Future "when will you add reactions?" questions are inevitable.
- The decision to lock cap-XX-style multi-capability progression into a single cap-mvp bundle (rather than the existing eight scaffolds) creates reconciliation work against the current repo.

## What would prove this wrong

- The MVP ships, gets used for 6 weeks by the actual circle, and the answer to "did the people I invited get time back" is *no*.
- A specific not-in-MVP feature blocks usage early (e.g., absence of comments leaves moments feeling like dead-letters and the circle stops posting).
- The five-capability cut is internally inconsistent in ways we don't see yet (e.g., onboarding's seeded-posts mechanic depends on something in moments-spec that doesn't exist).
