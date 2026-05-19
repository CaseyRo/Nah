---
title: "ADR-0002 — Mastodon as the backend foundation"
permalink: /decisions/0002-mastodon-backend/
layout: page
status: Accepted
date: 2026-01-28
---

## Status

Accepted (2026-01-28). Unchanged through the April 2026 Flutter pivot.

## Context

Nah? needs a backend that handles social-networking primitives: accounts, posts, timelines, follow/friend relationships, notifications, media handling, OAuth2. Building these from scratch is months of undifferentiated work. The question was whether to start from a battle-tested foundation or from a thin BaaS, and if a foundation, which one.

## Decision

Fork vanilla Mastodon. Run as a closed single instance: federation off, public timeline hidden, followers-only as the visibility default, friend-limit enforced at the database layer. Convert the follow system to mutual friendship (both must accept). Build Nah? as a Flutter client that talks to Mastodon's REST + Streaming APIs.

## Considered alternatives

- **Build from scratch.** Rejected. Three to six months of undifferentiated effort before reaching anything user-facing. We'd be reimplementing primitives Mastodon has solved well, with worse security posture from day one.
- **BaaS (Firebase, Supabase, Convex).** Rejected. Vendor lock-in on the most replaceable layer of the product. Privacy-by-architecture commitment conflicts with cloud providers we don't audit. The 150-friend network doesn't have the scale that BaaS unlocks; we don't need autoscaling for a circle of 150.
- **Hometown (Mastodon fork with local-only posting).** Rejected. Local-only posting is redundant when federation is disabled. Sole maintainer, ~9 months behind upstream — maintenance risk.
- **Glitch (another Mastodon fork).** Rejected. Kitchen-sink feature additions we don't need. More code to audit, more drift from upstream security patches.
- **ActivityPub-from-scratch.** Rejected. Same problem as build-from-scratch but with extra protocol surface. The protocol is a future option, not a starting point.

## Consequences

Positive:

- Battle-tested primitives. Years of bug fixes and security patches flow from upstream.
- Cleanest codebase among Mastodon-family options.
- ActivityPub is there if we ever want federation later. It's gated off, not removed.
- Standard Ruby on Rails + PostgreSQL + Redis. Operationally well-understood.

Negative / acknowledged cost:

- Mastodon upstream is philosophically opposed to private-network features. Our changes won't be upstreamed; we maintain a permanent fork.
- Rebase burden. Every Mastodon release needs a rebase pass. Surgical fork (small set of changes) keeps this manageable.
- The Mastodon team is opinionated about visibility models in ways that occasionally fight our closed-instance posture.

## What would prove this wrong

- Rebase cost becomes prohibitive (>2 days/quarter to absorb upstream releases).
- A security CVE lands in Mastodon that we can't patch in time because our fork has diverged too far on the relevant code path.
- We discover a fundamental architectural mismatch (e.g., Mastodon assumes a follower model that fights our mutual-friend model irrecoverably at the database level).
