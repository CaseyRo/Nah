---
title: "Decisions"
permalink: /decisions/
layout: page
---

Architecture Decision Records for Nah?. Each ADR captures one decision: why the question came up, what we picked, what we considered and rejected, and what would prove us wrong.

ADRs are not specs. They're the *why* behind the spec. Future contributors (or future-Casey) can read these to understand the reasoning even when the code no longer reflects the original tradeoff.

## Index

| # | Decision | Status |
|---|---|---|
| [0001](0001-flutter-over-pwa.html) | Flutter over PWA for the mobile client | Accepted (2026-04-26) |
| [0002](0002-mastodon-backend.html) | Mastodon as the backend foundation | Accepted (2026-01) |
| [0003](0003-connection-model.html) | Invitation is connection, no slot economy | Accepted (2026-04-26) |
| [0004](0004-no-counts-anywhere.html) | No counts anywhere as the MVP headline | Accepted (2026-04-26) |
| [0005](0005-ritual-onboarding.html) | Ritual onboarding (slow beat, one question, felt arrival, quiet space) | Accepted (2026-04-26) |
| [0006](0006-three-moment-types.html) | Three moment types (text, voice, photo) | Accepted (2026-04-26) |
| [0007](0007-respectful-notifications.html) | Notifications default to most respectful | Accepted (2026-04-26) |
| [0008](0008-mvp-scope.html) | MVP scope: 5-10 friends, real usage, explicit not-in-MVP list | Accepted (2026-04-26) |

## How to read these

Each ADR is short (one page). Sections in order: status, context, decision, considered alternatives with rejection reasons, consequences (positive and negative), what would prove this wrong.

The "what would prove this wrong" section is the most important and the easiest to skip. We keep it because decisions made in confident moods can be wrong in concrete circumstances, and the trigger that would force a revisit is worth writing down while we still remember it.

## How to add a new ADR

1. Copy `_template.md` to `NNNN-short-slug.md` where NNNN is the next number.
2. Fill in the front matter (`title`, `permalink: /decisions/NNNN-short-slug/`, `status`, `date`).
3. Write the decision. Keep it to one page.
4. Add a row to the index above.
5. Commit.

If a new ADR supersedes an existing one, update the older ADR's status to "Superseded by ADR-NNNN" and add a link.
