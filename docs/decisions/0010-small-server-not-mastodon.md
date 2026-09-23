---
title: "ADR-0010 — A small server of our own, not Mastodon"
permalink: /decisions/0010-small-server-not-mastodon/
layout: page
status: Accepted
date: 2026-09-13
---

## Status

Accepted (2026-09-13). Supersedes [ADR-0002](0002-mastodon-backend.html). Amended by [ADR-0011](0011-plain-go-and-a-database-per-circle.html): the server, the single binary and SQLite all stand, but the spike measured PocketBase at seven times the memory of plain Go and it was dropped.

**Amended 2026-09-23.** Until Nah? hosts servers for other people (M7), everyone is on one server, whose address the app carries as a build setting, so the directory is not needed yet and moves to M7 (CDI-1838). The push relay, the invite-link domain and the App Review demo stand.

## Context

ADR-0002 forked vanilla Mastodon for two reasons. The practical one was that Mastodon had already solved accounts, timelines, media and OAuth. The deeper one was a commitment to build on great open source software rather than assemble a proprietary stack on rented infrastructure.

That second commitment is not in question here. What changed is everything around it.

The September 2026 audit set a new requirement: any family or friend group should be able to run a server with something close to one click. Mastodon is five processes, a Postgres, a Redis, a job runner and a streaming server, and it wants two to four gigabytes of memory before anyone posts a photo. Managed Mastodon hosting exists precisely because a person has to look after it.

Then [ADR-0009](0009-circles-not-one-circle.html) made circles the core data model, and research settled whether anything off the shelf can express them. Nothing can.

## Decision

**Build Nah? Home as a single Go binary on PocketBase used as a framework, with SQLite, and run only the smallest possible set of central services.**

- One binary, one database file, media on local disk or S3-compatible storage, Litestream streaming backups off the machine.
- Photos and voice are encoded on the phone before upload, so the server never runs FFmpeg.
- PocketBase supplies sign-in, file storage, live updates, scheduled jobs, backups and an admin screen, and has an official Dart SDK for the Flutter client. We write the social layer: circles, memberships, invitations, moments, digests, lifecycle and export.
- The project runs three central things and nothing else: a push relay, because Apple and Google only accept pushes signed with the publisher's keys; the invite-link domain with a directory mapping a circle's identifier to its current address; and a demo circle for App Review. [Sygnal](https://github.com/element-hq/sygnal), the Matrix ecosystem's push gateway, is a mature answer for the relay and can be used without adopting Matrix.
- Federation is deferred. A circle lives on one server.
- Identity belongs to a key, never to a hostname, so a circle can move hosts without losing itself. This is the lesson from Mastodon, where an instance's domain can never change.

## Considered alternatives

- **Fork Mastodon (ADR-0002 as written).** Mastodon has four visibility values and no concept of an audience group. Lists are read-side only. The request to post to selected lists was closed as "not planned" in May 2023, and the groups request has been open since November 2016. The forks that did build circle posts hold the recipient list only on the sender's side, so nobody else can verify who a post was for. Supporting circles means a new data model, a fifth visibility state through the whole delivery pipeline, per-circle fan-out, reply inheritance and notification routing, after which we would switch off public timelines, trends, discovery and follower counts, which are Mastodon's actual product.
- **GoToSocial.** One Go binary with SQLite, and the lightest Mastodon-API server. Still follow-based, still beta, and its allowlist federation mode is experimental. It saves writing a server and then fights every rule we have.
- **Bonfire.** The only project with circles and per-post boundaries as first-class primitives, and its own worked example is a party hidden from one specific person. Rejected because private groups explicitly do not federate yet, it runs on Elixir and wants roughly 8 GB of RAM, and 1.0 shipped in 2025.
- **Friendica or Hubzilla.** Both have had per-item access control with named circles for a decade, and Hubzilla can actually prove a remote visitor's identity. Rejected for tiny communities, dated interfaces and a maintenance base of very few people.
- **Matrix.** The best conceptual fit: rooms are circles, one account can join circles on other people's servers, and the Dart SDK is well maintained and encrypts custom event types for free. Rejected because push rules have no time condition at all, so the per-circle digest would be built from scratch regardless; room version upgrades tombstone a room and do not transfer membership, which is fatal for an archive-centred product; export and account portability are weaker than a zip file; and FUTO Circles, a funded team building this exact product on Matrix, concluded they would have to write their own server anyway and archived it in February 2025.
- **Serverpod or Supabase.** Serverpod would let us share Dart models with the client but requires PostgreSQL. Self-hosted Supabase is ten or more containers. Neither can be a one-click install for a family.

## Consequences

Positive:

- The whole MVP is create-read-update-delete, one scheduled job and an export. A circle's feed is one query with no fan-out, no timeline cache and no ranking.
- A circle costs cents a month to run, so a single small server can host many, and a family can run its own on hardware it already owns.
- Nothing in the stack fights the product. There are no counts to hide, no public timeline to disable and no upstream to rebase against.
- Because there is no search, no ranking and no counts, the server never needs to read a moment. Encrypting moments later becomes a client change rather than a server rewrite.
- The Dart SDK keeps the Flutter client close to the server without a second language on the client side.

Negative / acknowledged cost:

- We own every security fix, every migration and every bug that a large upstream would have caught.
- PocketBase is pre-1.0 with one maintainer and breaks compatibility between releases, so the version must be pinned and vendored. The escape hatch is that underneath it is plain Go and SQLite.
- No federation means an account per server, which is the cost ADR-0009 already accepts.
- Moderation and abuse handling are ours, including the cut-off switch that Apple expects from an app's publisher.
- We give up the fediverse. Nah? will not interoperate with anything, by choice.

## What would prove this wrong

- PocketBase breaks something expensive between releases twice in a row, or is abandoned. The response is to drop to plain Go and SQLite, not to find another framework.
- People demand circles that span servers early, which would mean federation was not safe to defer.
- A circle outgrows SQLite. This should be impossible at 150 members, so if it happens the data model is wrong, not the database.
- App Review or the push relay forces a central service we are not willing to operate, which would undermine the claim that a family can hold its own data.
