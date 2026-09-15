---
title: "ADR-0016 — The moment envelope, and why it is not HTML"
permalink: /decisions/0016-the-moment-envelope/
layout: page
status: Accepted
date: 2026-09-13
---

## Status

Accepted (2026-09-14), after being proposed on 2026-09-13. Follows [ADR-0012](0012-encrypted-on-device.html), which made a moment opaque to the server, and [ADR-0006](0006-three-moment-types.html), which says a moment is one thing said in one medium.

**Amended 2026-09-15.** A music moment is the one exception to the rule against URLs outside the author's server. It carries the song's links on music services, which Nah? never fetches or previews, and which open only when a reader taps one. Its artwork is fetched once on the posting phone and travels sealed as an attachment, like any image.

It landed before the app wrote its first moment, which was the point. The app has written version 1 envelopes since 2026-09-14, unencrypted until CDI-1863, with the first byte of every blob saying how it is sealed. Where this record first said circle, [ADR-0017](0017-one-network-of-a-hundred-and-fifty.html) now means the people a moment's author is connected to.

## Context

[ADR-0015](0015-no-passwords-a-key-on-the-device.html) settled the server half of version skew, and the rule there is simple: the server deploys constantly, so the server may never require a client change.

The client half has no such rule available, because it has no server in the middle. A moment is encrypted on one device and decrypted on another, and the thing between them is a blob the server cannot read, translate, validate or migrate. So the format of a moment is a contract between phones.

That contract has to survive an unusually wide spread. Phones update when app stores and their owners feel like it. This product is explicitly for families, which means a grandparent on a two-year-old phone and a teenager on this week's build are connected to each other, reading each other. The old phone will be handed a moment written by an app that did not exist when it was installed, and it has to do something sensible.

It is also the last thing that is cheap now. Every moment written before this decision would be unversioned and unlabelled forever, and there is no server-side migration available to fix them, because the server cannot open them.

## Decision

**A moment is a versioned, self-describing envelope in JSON, and it always carries a plain-text fallback written by the client that posted it.**

```json
{
  "v": 1,
  "type": "voice",
  "fallback": "Anna sent a voice moment.",
  "sent_at": 1789310709329,
  "body": { "duration_ms": 14200, "waveform": [ 3, 9, 14 ] },
  "attachments": [ { "id": "…", "bytes": 151002, "media_type": "audio/opus" } ]
}
```

- **The fallback is written by the newer client, not guessed by the older one.** This is the whole mechanism. A future app that posts a type nobody has shipped a renderer for also writes one plain sentence describing it. An old client does not need to know what the new type is; it needs to know how to display a string, which it already does. Forward compatibility stops depending on predicting the future.
- **A closed set of types, extended only by addition.** `text`, `voice`, `photo` today, per ADR-0006. A client that meets an unknown `type` renders the fallback in place, in the feed, without disturbing anything around it.
- **Unknown fields are ignored, never rejected.** Fields are added; they are never renamed, repurposed or removed. `v` exists for the change that cannot be made additively, and the ambition is that it stays at 1 forever.
- **No envelope may contain a URL to anything outside the author's own server.** Media is an attachment: an id for a blob held by that server, fetched with the reader's session and decrypted on the device. There is no field anywhere that can point at the open internet.
- **Decrypted is not trusted.** The ciphertext proves the author held their content key. It proves nothing about whether the envelope inside is well-formed or benign, and the server cannot check it because the server cannot read it. Every reading client validates an envelope as hostile input — bounds on every length, no unbounded allocation, and a malformed envelope renders as a plainly-worded unreadable moment rather than a crash.
- **Reading survives much longer than writing.** A client too old to render a new type shows the fallback, indefinitely. A client too old to *write* a correct envelope is told to update, because a moment nobody else can read is worse than a person being asked to visit the app store.
- **JSON**, because `encoding/json` and `dart:convert` are both standard library on the two sides that matter, and because being able to read a decrypted moment with human eyes is worth a great deal in a product where every other debugging avenue is closed by design. Binary never goes inline; it goes in attachments.

## Considered alternatives

- **HTML as the payload.** The tempting one, and the instinct behind it is right: it is the one genuinely version-agnostic content format, unknown tags degrade instead of exploding, and a new moment type would need no client change at all. Rejected on three counts, any one of which is enough.

  *It does not solve the problem, it moves it.* Flutter has no HTML renderer. Rendering it means a package that maps some subset of HTML to widgets, so "any client can render anything" becomes "whatever that package version supports" — the same version skew, now with a third-party maintainer inside it and a much larger surface for it to go wrong in.

  *It lets content out.* HTML can reference external resources. One `<img>` pointing at a stranger's host tells that stranger the reader's IP address and the moment they opened it. That is content leaving the people it was posted to by construction, not by policy, and it is the one guarantee the entire architecture exists to make. A sanitiser that must be perfect forever is not a guarantee.

  *It is the wrong product.* ADR-0006 says each moment is one thing said in one medium, and the plainness is the feature. A format that can express anything will eventually be used to express a headline in forty-point bold, and the moment it can, not doing so becomes a choice each person makes rather than something the product has already decided.

  The fallback string is this idea with the attack surface removed: universally renderable, written by the author's client, incapable of reaching outside the people it was posted to.

- **A WebView per moment.** The version of the above that really would render anything. Rejected outright: it is a script engine pointed at bytes someone else wrote, decrypted on your device, in the one place this product promises nothing bad can happen. It is also ruinous for a scrolling feed.

- **Protobuf or CBOR.** Smaller, and schema evolution is their entire subject. Rejected because the saving is irrelevant — a 550 KB photo does not care about a few hundred bytes of envelope — and each costs a dependency on both sides against a budget of five. Revisit only if envelopes ever stop being small.

- **No envelope: put the type in a column and encrypt only the content.** Simpler, and it is what a normal app would do. Rejected because the type of every moment then becomes metadata the server holds, so we would know that someone sent a voice moment at eleven at night. ADR-0012 already concedes more metadata than is comfortable; this would add to it for nothing, and it would put the server back in the business of knowing what a moment is.

- **Ask the server what the reader's client can handle.** A negotiation, as HTTP does with content types. Impossible here and worth stating plainly: the author encrypts once, for everybody, possibly days before anyone reads it. There is no moment at which the server knows both what is inside a blob and who is about to open it, and there never will be.

## Consequences

Positive:

- A new moment type ships to half of someone's people without breaking the other half. That is what makes shipping to families tolerable at all.
- The guarantee that content never leaves the people it was posted to becomes a property of the format rather than a rule people have to keep remembering.
- A decrypted moment is readable by a person, which is the only debugging that will be available when something goes wrong inside the encryption.
- Costs nothing today: it is a few extra fields written before the first moment exists.

Negative / acknowledged cost:

- **Every new type costs a sentence of English, in every language shipped.** The fallback is user-facing text, so it lands in the translation surface rather than being a developer detail.
- **A fallback written badly is a bad experience,** and the client writing it is the one client that cannot see how it looks on the old app that will render it. This needs a review habit, not a mechanism.
- **We cannot enforce any of it.** The server cannot validate an envelope it cannot read, so a broken or hostile client can write anything into someone's feed. The reading client is the only defence, which is why it must treat every envelope as hostile.
- **JSON is verbose,** and every envelope is encrypted, so it does not even compress across moments. Accepted: they are small next to media.
- **A closed type set means a slower kind of change.** Adding a type is a release on both ends, not a content decision. That is the intended friction.

## What would prove this wrong

- A fallback shows up in a real feed and reads like a machine wrote it, in a product whose whole argument is that it is not a machine.
- A type turns out to need a field that cannot be added without breaking older readers, forcing `v` to 2 inside the first year. That would mean the envelope was drawn too tightly, and the shape worth stealing is whatever the change wanted to be.
- Someone finds a way for an envelope to cause a network request out of the app despite the no-URL rule — a media type handed to a platform decoder that resolves something, most likely. That would mean the rule needs enforcement in code, not just in review.
- Envelopes stop being small, because a type arrives that wants structure rather than an attachment, and the JSON overhead starts showing up in what a person's moments cost to store.
