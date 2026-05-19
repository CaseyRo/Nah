---
title: "ADR-0003 — Invitation is connection"
permalink: /decisions/0003-connection-model/
layout: page
status: Accepted
date: 2026-04-26
---

## Status

Accepted (2026-04-26).

## Context

Nah? is invitation-only, organic-growth-only, one-circle-per-user. The question was how the invitation mechanism, the slot mechanic, and the connection model actually interact.

The first sketch was a mutual-only graph with a slot economy: each connection costs both sides one slot from a finite pool (~150 per user). Spending a slot was meant to make invitations feel like a real cost. Casey pushed back: the slot bookkeeping created sharp edges (empty quiet-space after first accept, starter-slot allocations, pending-invite limbo, removal mechanics that need notifications), and the mechanic ended up feeling like SaaS-economy thinking rather than friendship.

## Decision

**Invitation is connection.** When Casey invites Wife and Wife accepts, both are connected in one motion. No separate "now reciprocate" step. No slot ledger.

When Wife invites Maya and Maya accepts, Maya is in Wife's circle automatically. Maya is **not** in Casey's circle. If Casey wants Maya in his, Casey invites Maya separately (or the Phase 2 touch-to-connect ritual handles it).

The soft 150 cap is communicated as language ("your circle is full"), not as a count. It bites when you bump into it, not as a hovering progress bar.

## Considered alternatives

- **Mutual-only graph with slot economy.** Rejected. Created a starter-slot bookkeeping problem (does Casey gift Maya her first slot back? does Maya arrive with zero slots?). Made the first-open experience structurally empty until reciprocation. Slot scarcity as information conflicts with "no counts anywhere" (ADR-0004).
- **Asymmetric follow (like Twitter, Instagram).** Rejected categorically. Asymmetric relationships are the substrate of parasocial dynamics, follower-count games, and the entire social model Nah? rejects. Every Nah? edge is bidirectional and equal.
- **Two-degree visibility (you see your circle plus their circles).** Rejected. Reintroduces a discovery surface and makes "who can see this moment" much harder to reason about. Conflicts with "no public anything."
- **Direct-only without reciprocity required.** Rejected. Same asymmetric-follow problem in a closed wrapper. If invite-without-reciprocation is enough to see, the inviter's choice unilaterally enrolls the invitee into a follower relationship.

## Consequences

Positive:

- One mechanism handles both onboarding new users and connecting existing users. The system can recognize whether an invitee has an account and route accordingly, transparently.
- No empty-room-after-accept failure mode. Accept-and-arrive in one step.
- Soft cap stays felt-not-measured, consistent with ADR-0004.
- The wife-overlap example works without bookkeeping. Each connection is one decision between the two people involved.

Negative / acknowledged cost:

- Existing-users-finding-each-other in MVP relies on out-of-band coordination (one of them sends the other an invite via WhatsApp, iMessage, whatever). No in-app search.
- The cap-related friction lands as a surprise the first time someone hits it, rather than being predicted by a visible meter.
- No mechanism today to break a connection silently and bilaterally. Needs design work in cap-mvp or follow-up spec.

## What would prove this wrong

- A user tries to connect with someone they actually know on Nah? and gives up because the out-of-band coordination is too friction-heavy.
- The lack of a slot meter leads to a user discovering they're "full" mid-invitation, in a moment that feels broken rather than felt.
- Phase 2 touch-to-connect doesn't ship and the in-app connection mechanism for existing users (sending each other invite links) proves to be too clunky for sustained use.
