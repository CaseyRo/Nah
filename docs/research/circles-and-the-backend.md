---
title: "Circles and the backend question"
permalink: /research/circles-and-the-backend/
layout: page
date: 2026-09-13
---

*Can Mastodon, the wider fediverse, or Matrix give Nah? many private circles per
person? And what do per-circle digests, dying circles and archives look like?*

## The question

Nah? was designed around one circle per person: your closest 150, one room, one
feed. Real life is not shaped like that. A single person belongs to a family
circle, a close-friends circle they keep away from their mother, a neighbourhood
circle for organising the street party, a school-parents circle for pickups and
playdates, a circle for people who own the same dog breed, and a puppy-training
circle they see once a week.

If that is the product, the backend question reopens. Mastodon was chosen partly
to build on and give back to great open source software. That stance is not in
question. What is in question is whether Mastodon, or anything else that already
exists, can express circles.

## First, two different shapes of circle

This is analysis rather than research, and it matters more than the technology
choice.

**Memory circles** are family, close friends, the dog-breed group. They are
long-lived, they accumulate, and their value grows with time. A moment posted
there is worth re-reading in a year. This is the Nah? thesis exactly.

**Logistics circles** are the school run, the street party, the puppy class.
They are short-lived or utility-shaped. What they need is who is picking up
whom, when, and a reminder two hours before. A chronological feed of moments is
not what that job wants, and a group chat already does it well.

Serving both doubles the product. Worse, the research on this category is
consistent that the products which survive own one job the group chat does
badly, and the ones that tried to be a nicer place to hang out all died. The
safe cut is to build Nah? for memory circles, let logistics circles exist if
people make them, and resist the pull to add chat when they feel thin.

## Can Mastodon do this? No

Not "not yet". The vocabulary does not exist.

- Mastodon has exactly four visibility values: public, quiet public, followers,
  and private mention. Followers means every follower, never a subset. Private
  mention re-lists every recipient in the post body each time and lands as a
  notification rather than in a feed.
- Lists are read-side only. They group the people you follow, for your own
  timeline. They have no effect on who can see what you post.
- The request to post to selected lists was opened in 2018 with a mockup of this
  exact feature and **closed as "not planned" in May 2023**. The request for
  groups has been open since **November 2016**. A 2025 request to support the
  group federation spec sits unassigned and unanswered.
- A Mastodon maintainer drew the distinction precisely in that thread: a list
  categorises the people you follow, while a circle categorises your followers,
  and the second one "should be a new feature. Maybe call it Audiences." It was
  never built.
- The Japanese forks that did build circle posts (Fedibird, kmyblue) prove the
  cost. Sending had to be written from scratch, replies needed their own new
  visibility mode, and the recipient list is held **only by the sender**. Nobody
  else can verify or reconstruct who a post was for.

To get six overlapping private circles you would add a circles table, a per-user
membership model, a fifth visibility state carried through the whole delivery
pipeline, per-circle fan-out, reply inheritance, notification routing,
moderation paths for content no moderator can see, and export. Then you would
switch off public timelines, trends, discovery and follower counts, which are
Mastodon's actual product. That is a new data model plus a product inversion,
not a patch.

## The rest of the fediverse

**Bonfire** is the one project with this as a first-class primitive, and it is
genuinely good. Circles are named sets of contacts, boundaries are reusable
permission sets, and roles run from read to caretaker with explicit denies that
always win. Its own documentation example is a surprise birthday party where
friends can read and reply, family can also invite, and the birthday girl is
denied sight of it. That is Casey's "things you don't share with your mother",
already shipping. The catches: boundaries are enforced locally only, **private
federated groups are explicitly deferred** because the fediverse has not agreed
an interoperable spec, it runs on Elixir and wants roughly 8 GB of RAM, and 1.0
shipped in 2025.

**Friendica** and **Hubzilla** have had this model for a decade. Every item
carries an access control list of people and circles, with allow and deny, and a
contact can sit in many circles. Friendica pre-sets the composer's audience from
the circle you are viewing, which is a lovely detail worth stealing. Hubzilla is
the only thing in the fediverse that can actually prove who a remote visitor is.
Both are small, both look dated, and Friendica's own documentation admits
private photos cannot be shared outside Friendica and deletions may never
propagate.

**The protocol itself has no answer.** The group federation spec is final, but it
describes public forums and contains no discussion of private or
restricted-audience groups at all. You can address twelve people in an activity.
You cannot make any receiving server honour that. Private posts are delivered in
full to every server hosting a follower, replies fan out to the replier's
audience rather than the original's, and deletes are best effort.

## Matrix

Matrix fits the shape better than anything else. Rooms really are circles, with
their own membership, history and power levels. One account can join circles
hosted on other people's servers, which is the single thing a custom server
cannot do at all. The Dart SDK behind FluffyChat is professionally maintained,
sends arbitrary custom event types, and encrypts them without extra work.

It still is not the answer.

- **Push rules cannot express time.** The complete condition list has no time or
  schedule condition. The proposal to add one has sat unmerged for three years.
  Nah?'s single most distinctive rule, one digest per circle at a time the member
  picks, would be written entirely from scratch on top.
- **Custom events fall out of the notification system** unless encrypted, at
  which point every event looks identical to the push layer.
- **E2EE is a support burden for non-technical families.** Element's own
  unable-to-decrypt tracking issue has been open since 2022 and is still being
  updated. Lose the phone and the recovery key and encrypted history is gone.
  History for people who join later was fixed, reverted after two CVEs in 2024,
  and re-landed only in April 2026.
- **Archives are weaker than a zip file.** Export is one room at a time, media is
  never cascade-deleted, retention defaults to keeping everything forever, and
  account portability proposals have been open since 2018 without implementation.
- **Room identity is not permanent.** Room version upgrades work by tombstoning
  into a new room, and membership does not transfer. Every family member rejoins.
- **No homeserver is one-click for a family.** Synapse needs Postgres and is
  heavier than the Mastodon stack already rejected. The light Rust servers are
  two mutually hostile forks of an archived project.
- **The precedent is unambiguous.** FUTO Circles was this exact product, E2EE
  Matrix social for families, with a funded team. They forked the server, the
  media repo, the client SDK and the deploy tooling, wrote a bespoke auth
  service, and concluded: *"Getting good performance out of Matrix servers is
  something of a dark art… We expected that we would need to write our own
  server if we wanted to grow the product."* Archived February 2025.

Two things are worth taking from Matrix without adopting it. **Sygnal**, its
push gateway, is a mature answer to the relay Nah? needs whatever it builds. And
its history-visibility vocabulary, invited versus joined versus shared, is a
ready-made model for what a new circle member can see, which is a question that
has to be answered anyway.

## What this means

**The open-source commitment stands. It was never a commitment to Mastodon.**
SQLite, Litestream, Go, Flutter and Dart are all open source, all maintained,
and all things a single small binary can be built from.

**Circles are Nah?'s core data model, and nothing off the shelf has them in the
shape needed.** Buying someone else's core data model that lacks your core
concept is not a shortcut.

**Decouple the server from the circle.** The earlier recommendation of one server
per group does not survive six bubbles per person. Nobody runs a server for a
puppy class, and nobody pays six subscriptions. A server hosts many circles. A
person may have accounts on a few servers, and the app merges them into one
view. Cross-server circle membership is federation, which is the expensive part,
and it should be deferred until it hurts.

## Per-circle digests

Every membership gets its own rhythm: silent, daily at a chosen time, weekly on
a chosen day, or a set time before the circle's recurring event. Family silent,
friends on Sunday morning, puppy training two hours before class.

This is cheap to build. It is one schedule row per membership and a job that
ticks. It is also the right thing by the evidence: anchoring a behaviour to
something already in someone's day is the one habit mechanism with trial support
behind it, and a predictable time beats a clever one. Nothing in the fediverse or
Matrix does this, so it is genuinely Nah?'s own.

## When a circle dies

Two different events, which deserve different treatment.

**A circle that never got going** should close itself quietly. Nextdoor deletes a
new neighbourhood that has not reached ten members in three weeks, and it is the
only product in this category with a hard mechanic against empty rooms. Nothing
of value is lost, and an app full of dead rooms teaches everyone the product is
dead.

**A circle that had a life and went quiet** must never be deleted silently. Warn
the members, let anyone revive it by posting, and if nobody does, hand every
member the archive before closing it. The bandwidth saving is real but small. The
reason to do it is that a living app should not be mostly ghosts.

## Whose archive is it?

Both, and they are different objects.

A moment is written by a person and addressed to a circle, so there are two
claims on it.

**The personal archive** is everything you have posted, across every circle,
yours permanently, exportable, and surviving the circles themselves. This is the
thing people in this category actually pay for, and it is the answer to the
failure that killed BeReal: something worth opening on a day when nobody else
posted.

**The circle archive** is the shared memory of that group. It belongs to the
members collectively, it can be exported by any member while they are one, and
it ends when the circle ends.

Resurfacing should be personal first. Showing you your own moment from a year
ago is always safe. Showing you someone else's, a year later, is a decision they
did not make, so circle-wide resurfacing should be something a circle opts into.
When you leave a circle you keep your own moments and nobody else's.

## The warning

This is the fourth backend decision in eight months. The research above is
conclusive enough to end the question rather than continue it: no existing
open-source social backend has private circles in the shape Nah? needs, and the
one that comes closest cannot federate them.

The next artifact should be a decision record and a walking skeleton, not more
research. The most expensive thing about evaluating Matrix was the week not
spent shipping.
