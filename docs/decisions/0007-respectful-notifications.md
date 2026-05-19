---
title: "ADR-0007 — Respectful notifications default"
permalink: /decisions/0007-respectful-notifications/
layout: page
status: Accepted
date: 2026-04-26
---

## Status

Accepted (2026-04-26).

## Context

Every social app's notification system is a moral question. Defaults shape behavior. The standard default is "send a push for each event" (each like, each comment, each follow), then expect users to opt OUT once they're annoyed. That default treats user attention as cheap and assumes the user's job is to defend their own attention against the app.

Nah? rejects that posture. The brand is "give people back the time they invest in what's really important." If the app's first action on a new install is to interrupt the user's day, the brand is already broken.

## Decision

The notification default is **the most respectful option available**, with users opting INTO more, never the other direction.

Concrete shape:

- **Default for new users:** one daily digest. Time-of-day user-configurable, suggested ~7pm local. Message uses vague-quantity wording ("Your circle posted today") with no counts. Tap deep-links to the feed.
- **Zero-moments suppression.** If your circle posted nothing today, no notification arrives. The daily digest exists when there's something to say, not as a heartbeat.
- **Push for per-moment is opt-in, never the default.** A user can enable "notify me for every moment from Wife" or "every moment, ever." But the default is the digest.
- **No in-app red dot for unread moments.** Unread state exists in the feed view (a moment looks unread until you scroll past it), but no tab-bar badge, no app-icon badge.

## Considered alternatives

- **Per-moment push by default (industry standard).** Rejected. Treats user attention as cheap. The opposite of the brand.
- **Push for nothing, pull-only.** Considered. Most respectful possible. Rejected as MVP default because friends-of-Casey will install and then check Nah? once, find nothing, never return. The daily digest is the lightest possible reminder that something is alive in there, without being noisy.
- **Weekly digest by default.** Rejected. Too infrequent for a 5-10-friend circle where any individual posting cadence is irregular. Weekly creates "nothing today" feeling even when the circle was active.
- **Smart-timing notifications (ML-determined optimal moment).** Rejected categorically. Smart timing optimizes for the wrong thing (open-rate, engagement). We don't optimize for those.
- **User picks frequency at onboarding.** Rejected as MVP default-setter. Choice-paralysis on first install. The default carries the philosophy; the settings screen lets people opt INTO more.

## Consequences

Positive:

- Default behavior carries the brand. New users experience the philosophy in action on day one, not as marketing copy.
- Lower notification volume than any comparable social app. Lower battery cost. Lower interruption cost.
- Zero-moments suppression prevents the "nothing today" feeling that would otherwise erode trust in the daily digest.

Negative / acknowledged cost:

- Lower engagement metrics than competing apps. If a future version of Nah? is judged by DAU, this default looks bad on paper. (Counter: we don't judge by DAU.)
- Users who want more frequent notifications have to find the setting. Discoverability of opt-IN matters.
- The daily digest depends on push, which has the EU iOS DMA reality (push gated). Means EU iOS users get an even quieter experience than non-EU. Acceptable; the rest of the app still works.

## What would prove this wrong

- A user mistakes the silence for the app being broken. Specifically: they ask Casey "is Nah? working?" because they've installed and never heard from it.
- The opt-IN settings screen is so undiscoverable that users who would have wanted more notifications never find them.
- The daily digest sends too little (zero-moments-too-often) and friends silently stop opening the app because it never has anything.
