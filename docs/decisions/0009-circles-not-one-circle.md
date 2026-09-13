---
title: "ADR-0009 — Circles, not one circle"
permalink: /decisions/0009-circles-not-one-circle/
layout: page
status: Accepted
date: 2026-09-13
---

## Status

Accepted (2026-09-13). Supersedes [ADR-0003](0003-connection-model.html). Amends [ADR-0007](0007-respectful-notifications.html) and [ADR-0008](0008-mvp-scope.html).

## Context

Nah? was designed around one circle per person: your closest 150, one room, one feed. ADR-0003 made invitation the act of connecting inside that single circle.

Real life is not shaped that way. One person belongs to a family circle, a close-friends circle they keep away from their mother, a neighbourhood circle for the street party, a school-parents circle for pickups, a circle for people who own the same dog breed, and a puppy-training circle they see once a week. Forcing all of those into one audience is the exact reason people stop posting: the moment you have one room, you write for the most conservative person in it.

Two other things forced the question. The September 2026 concept audit proposed one server per group, which does not survive six bubbles per person, because nobody runs a server for a puppy class and nobody pays six subscriptions. And the research into private social apps found that the products which survive a decade are the ones where one person posting is enough, which changes what a circle has to support.

## Decision

**A circle is the unit of audience, a person belongs to many, and a server hosts many circles.**

- Every moment is addressed to exactly one circle. "Who can see this?" always has a one-word answer.
- Invitation is joining. The spirit of ADR-0003 survives: one motion, no separate reciprocation step, no slot economy. What changes is that you are invited to a circle rather than to a person's life.
- Server and circle are decoupled. One server hosts many circles, so creating a circle is free and creating a server is rare.
- A person may hold accounts on more than one server, and the app merges their circles into a single view. Cross-server circle membership is federation, and it is deferred until it hurts.
- Dunbar's cap moves from the person to the circle. A circle has an upper bound, communicated as language, never as a meter, per [ADR-0004](0004-no-counts-anywhere.html).
- **Digest rhythm is per membership**, extending ADR-0007. Each person sets each circle to silent, a daily time, a weekly day and time, or a fixed interval before that circle's recurring event. The respectful default stands; what changes is that the default is now per circle.
- **Circles have a lifecycle.** A circle that never gets going expires quietly, because nothing is lost. A circle that had a life and goes quiet warns its members, can be revived by anyone posting, and is archived with an export for every member before it closes. Silent deletion never happens.
- **The archive is two objects.** The personal archive is everything you have posted anywhere, permanently yours, exportable, and outliving the circles it was posted to. The circle archive is the group's shared memory and ends with the circle. Resurfacing is personal first; circle-wide resurfacing is something a circle opts into.

Circles come in two shapes, and only one of them is the product. **Memory circles** (family, close friends, the dog-breed group) accumulate and are worth re-reading in a year. **Logistics circles** (the school run, the street party, the puppy class) want to know who is collecting whom. Nah? is built for memory circles. Logistics circles may exist, but we do not add chat to serve them.

## Considered alternatives

- **Keep one circle per person (ADR-0003 as written).** Simplest data model, and it is what the repo says today. Rejected because a single audience is the thing people route around by not posting, and because the many-bubbles case is the actual life of the intended user.
- **One server per group, where a circle and a server are the same thing.** This was the September audit's recommendation. Rejected because six circles would mean six servers, six accounts and six subscriptions. Nobody administers a server for a dog-training class.
- **Per-post audience selection from a list of people.** Google+ circles, essentially. Rejected because it makes every post a permissions exercise, and because ADR-0003 already rejected the reasoning that leads there. A circle is a place you post into, not a checkbox list you assemble each time.
- **Two-degree visibility, where you see your circles and their circles.** Rejected for the same reason as in ADR-0003: it reintroduces a discovery surface and makes the audience impossible to reason about.

## Consequences

Positive:

- The audience is a place, not a setting. This is how people already organise, in named group chats.
- One motivated person brings a whole group, which softens the cold start that killed most products in this category.
- Per-circle rhythms become possible, which nothing in the fediverse or Matrix can express today.
- Self-hosting stays meaningful, because a family or a community can run a server and put several of their circles on it.
- A quiet circle is no longer a dead app. Your other circles, and your own archive, still have something in them.

Negative / acknowledged cost:

- An account per server until federation exists. The app hides it, but the seam is real when you join a circle on someone else's machine.
- A moment belongs to one circle, so posting the same thing to two circles is a deliberate act rather than a default.
- The personal archive becomes a first-class surface that must be built, not a by-product.
- More moving parts than one circle: membership, per-circle settings, lifecycle states and two archive scopes.
- ADR-0008's MVP scope is written for one circle and needs re-cutting. In particular, the no-reactions decision deserves revisiting, because posting into silence is the documented way small circles die.

## What would prove this wrong

- People create one circle and never make a second one. If the median user has exactly one circle after six weeks, the whole model is overhead and ADR-0003 was right.
- People make logistics circles, ask for chat, and stop using Nah? when it is not there. That means the two shapes cannot live in one product.
- The account-per-server seam blocks real joins. Concrete trigger: someone invited to a circle on a friend's self-hosted server gives up during signup.
- Per-circle digest settings go untouched. If almost everyone leaves every circle on the default, the flexibility was imagined rather than needed.
