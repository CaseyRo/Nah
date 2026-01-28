---
title: "Why We're Forking Vanilla Mastodon (Not Hometown)"
date: 2025-01-28 10:00:00 +0100
categories: [Architecture, Decisions]
tags: [mastodon, backend, open-source]
---

When we started designing Nah, the obvious choice seemed to be [Hometown](https://github.com/hometown-fork/hometown) — a Mastodon fork specifically built for private, local-only communities. But after thinking it through, we decided to fork vanilla Mastodon instead. Here's why.

## The Appeal of Hometown

Hometown adds a killer feature to Mastodon: **local-only posting**. In a federated network, this lets you post something that stays on your server and never gets shared with the wider fediverse. It's perfect for inside jokes, community-specific discussions, or anything you don't want leaving your cozy corner of the internet.

Darius Kazemi built Hometown because Mastodon's maintainers are philosophically opposed to features that create "walled gardens." They believe in open federation. Hometown exists for people who want something more private.

So why didn't we use it?

## The Realization: We Don't Need Local-Only

Here's the thing: **if you disable federation entirely, every post is automatically local.**

Let me explain. "Local-only" is a per-post setting that says "don't send this to other servers." But Nah is designed as a single, closed instance. There are no other servers. Federation is completely off. So the distinction between "local" and "federated" posts becomes meaningless — they're all local by definition.

Hometown solves a problem we don't have.

## The Maintenance Risk

Beyond the technical redundancy, Hometown has a practical problem: it's in maintenance limbo.

As of early 2025, Darius Kazemi (the sole maintainer) has stated he wants to hand off governance. For the past year, the project received almost no maintenance except for critical security patches. At one point, it was 9 months behind Mastodon's releases.

Forking a fork that's struggling to keep up with upstream seemed risky.

## What About Glitch-soc?

We also considered [Glitch-soc](https://glitch-soc.github.io/docs/), another Mastodon fork with local-only posting and many other features. It's more actively maintained than Hometown, but it's described as a "kitchen sink" of experimental features. More complexity than we need.

## Our Approach: Surgical Modifications

By forking vanilla Mastodon directly, we get:

- **Best maintenance path** — security patches flow from the largest community
- **Cleanest codebase** — less to understand and maintain
- **Full control** — we're not waiting on someone else's governance decisions

Our modifications are surgical:

1. **Disable federation** — single closed instance
2. **Hide local/public timelines** — privacy by architecture
3. **Default posts to followers-only** — your friends see your posts, nobody else
4. **Add 150-friend limit** — Dunbar's number, enforced at the database level
5. **Extend reactions** — beyond simple "favorites"

That's it. We're not rebuilding Mastodon; we're configuring it for privacy and adding a few features on top.

## The Philosophical Tension

There's an irony here. Mastodon's maintainers don't want private features because they believe in open federation. We're using their software to build the opposite — a closed, intimate space.

But that's the beauty of open source. We don't need their permission or buy-in. We fork, we modify, we build what we need. They maintain the 95% of the code we're using as-is, and we maintain our small additions.

Everyone wins.

## The Bottom Line

Hometown and Glitch-soc exist because people wanted features Mastodon wouldn't add. But those features solve for "private posts on an otherwise federated server." We want something simpler: a completely private server.

For that, vanilla Mastodon with federation turned off is the cleanest foundation.

---

*This is part of an ongoing series documenting the decisions behind Nah. Follow along as we build a private social network for your closest 150 people.*
