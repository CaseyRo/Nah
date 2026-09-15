---
title: "ADR-0017 — One network of a hundred and fifty, not circles"
permalink: /decisions/0017-one-network-of-a-hundred-and-fifty/
layout: page
status: Accepted
date: 2026-09-13
---

## Status

Accepted (2026-09-13). Supersedes [ADR-0009](0009-circles-not-one-circle.html) and [ADR-0013](0013-no-aggregated-feed.html), both accepted earlier the same day. Amends [ADR-0012](0012-encrypted-on-device.html), where "a circle has a content key" becomes "a person has one". Restores the framing [ADR-0003](0003-connection-model.html) had before ADR-0009 replaced it.

**A note on naming, 2026-09-15.** The product calls this one network *your circle*. That is not the circles this record rejects: a person's circle is drawn around them alone, nobody shares it or enters it, and each person has exactly one. The title keeps "not circles" because it records what was decided on 2026-09-13.

## Context

ADR-0009 made circles the core object: a family circle, a close-friends circle, a neighbourhood circle, each capped at Dunbar's number. Everything decided since has been built on it, including the server that now exists.

Two problems surfaced on the same day, and they turned out to be one problem.

**The cap stopped meaning anything.** Dunbar's 150 is an ego-network figure — the number of stable relationships one person can hold. Before ADR-0009 that was exactly what the product used it for: one circle was your whole social world, so 150 was a claim about you. When one circle became many, the number was copied onto each circle rather than re-derived. `PRODUCT.md` still carries the move in a single sentence — "each capped at Dunbar's number" — and six circles of 150 is 900 people, which is six times what the number says is possible. Dunbar's number is not, and never was, a statement about how big a group can be.

**A circle is a WhatsApp group with the chat removed.** This is the harder problem and it is a product one. If a family already has a group, a family circle is that same group minus the thing people use it for, which leaves no reason to post. Circles were organising something people already have a place for.

What nobody has a place for is the audience in between: the new job, the birthday cake, the holiday plans, the separation. Too much for a public feed, too diffuse for any one group. That audience is not a category, it is a closeness — and it does not sort into rooms, because it is CEOs and neighbours and a school friend all at once.

## Decision

**A person has one network of up to 150 people, and that is the whole model. There are no circles.**

- **The 150 is yours, not a group's.** Dunbar's number goes back to describing a person, which is what it describes. Membership is by closeness, not by context; there is no expectation that the people in it know each other.
- **You post to your people.** One audience. There is nothing to choose between when composing, which is the point: choosing an audience is the moment a person starts performing.
- **You read one chronological feed** of moments from the people you are connected to. ADR-0013 refused a merged feed because merging several audiences back into one stream undoes what circles did. With one audience there is nothing to merge, so the objection does not survive its premise. Chronological, finite, no ranking, no infinite scroll — all of that stands and none of it depended on rooms.
- **Connection is mutual, and the primary way to make one is physical.** Two phones touched together, the gesture Poken made its whole product out of. A shared invite link remains, as the exception rather than the default.
- **The cap becomes enforceable again.** A person's network is a server-side object, so 150 is a rule rather than a hope. It is still expressed in language and never as a meter, per [ADR-0004](0004-no-counts-anywhere.html).

### Why physical first

It is not nostalgia for a Swiss keychain. Requiring proximity does three things nothing else does as cheaply.

It makes "Viral? Nah" structural instead of aspirational — a network you have to build in person cannot be grown by a growth loop, and no amount of later product pressure can quietly turn it into one.

It demotes the leaky path. ADR-0012's worst failure is an invite link that survives in a group chat long after it should, and that link carries the key. When the ordinary way to connect is a touch, the link becomes the exception, used rarely, revoked easily.

And it matches how these relationships actually begin. The people who belong in someone's 150 are overwhelmingly people they have stood next to.

## Considered alternatives

- **Circles, as ADR-0009 decided.** What we are leaving. Rejected on the argument above: it duplicates group chat without the chat, and it broke the cap on the way past. The reasoning in ADR-0009 was not wrong about audiences being plural — it was wrong that the plurality wants rooms.
- **Circles as subsets of the 150.** The compromise: keep rooms, but you may only put your own people in them. It fixes the cap and costs almost nothing to build, which made it tempting. Rejected because it keeps the WhatsApp-without-chat problem entirely intact, and adds a sorting chore on top of it. The cap was the smaller of the two faults.
- **One network, but no cap at all.** Rejected. The cap is the product. Without it there is nothing structurally preventing this from becoming a small Instagram, and the promise that there is nobody to perform for rests on the audience being people you actually know.
- **Keeping a room for events — a wedding, a festival, a trip.** Genuinely lost here, and worth naming rather than quietly dropping. CDI-1866 existed for it. Rejected because an event is a group of people who are not close to each other, which is the definition of what this product is not for, and because carving out one exception reintroduces the whole circles apparatus for a case the 150 was never meant to serve.

## Consequences

Positive:

- Dunbar's number means what it means, and the cap is a server-side rule rather than an unenforceable guide.
- The product now occupies a gap instead of competing with a group chat that already exists and is better at being a group chat.
- Composing gets simpler in the way that matters: there is no audience picker, so there is no moment where a person edits themselves for a room.
- Everything ADR-0013 protected survives — chronological, finite, unranked — without the caching and prefetching machinery it needed to make switching feel instant. There is nothing to switch.
- A pile of scope disappears: circle directories, circle lifecycles, circle switching, circles that expire, circles that end on purpose.

Negative / acknowledged cost:

- **The feed is now fan-in, and the spike did not measure that.** ADR-0011's numbers were taken against one room and one query. Reading a feed assembled from 150 people's moments is a different shape, and "the feed is one query, no fan-out" is no longer true. This is the one consequence that could change how the server is built rather than what it does, and it needs measuring before it is assumed.
- **ADR-0012 needs restating rather than rewriting.** A key per circle becomes a key per person, handed over when someone connects. The guarantee is unchanged; the noun is not.
- **One audience means one audience.** Everything a person posts goes to all 150, so the thing they would only tell three people has no home here. That is a real loss and it is deliberate: the alternative is an audience picker, and an audience picker is the beginning of performing.
- **Physical connection is a real constraint on people who are not in the same room,** including exactly the distant family this product keeps citing. The invite link has to work properly, not as an afterthought.
- **Two decisions made this morning are superseded by this evening.** Worth noting honestly: both were made before the cap was examined, and neither was wrong about the problem.

## What would prove this wrong

- People fill their 150 and then want to say something to a subset of it, repeatedly and out loud. That is the audience picker asking to exist, and refusing it twice would be stubbornness rather than design.
- The fan-in feed measures badly enough that serving 150 people's moments needs a timeline cache, which would put back exactly the machinery ADR-0013 was written to avoid.
- Physical-first connection turns out to strand the distant relationships people most want here, so that in practice everybody uses the link and the touch is decoration.
- Nobody reaches anything like 150, and the real number is thirty — which would mean the cap is not doing the work claimed for it and the product is a different, smaller thing.
