---
title: "ADR-0012 — Content is encrypted on the device, and we hold no keys"
permalink: /decisions/0012-encrypted-on-device/
layout: page
status: Proposed
date: 2026-09-13
---

## Status

Proposed (2026-09-13). Extends [ADR-0010](0010-small-server-not-mastodon.html). Narrows the MVP scope of [ADR-0009](0009-circles-not-one-circle.html) by moving self-hosting out of the first release.

## Context

The September audit noted that Nah? never searches, ranks or counts anything, so its server never needs to read a moment. That made encryption a client change rather than a server rewrite, and parked it as a later possibility.

The proposal is to stop parking it. Encrypt on the device from the start, so privacy is structural rather than promised, and simplify the first release by having the project host every circle rather than asking families to run servers.

One correction to the framing before the decision. Encryption here cannot be per person and per device, because a moment has to be readable by everyone in the circle. This is group encryption: a circle has a key, and the people in the circle hold it. "Only the author's device can read it" would describe a diary, not a circle.

## Decision

**The server stores opaque blobs, and the project holds boxes rather than keys.**

- Moments and media are encrypted on the device before upload and decrypted on the device after download. The server can store, order and deliver them. It cannot read them.
- **A circle has a content key**, and it reaches a new member through the invite link, in the part after the `#` that browsers never send to a server. The key therefore travels to members without the project ever seeing it, using the mechanism that already admits people to a circle.
- **The project hosts circles for the first release.** Self-hosting moves to a later phase rather than out of the product. Circles still register centrally so people can find and join them, and so a circle can move hosts later.
- **Key management ships in stages**, because this is where family products die:
  - *First release.* One key per circle, delivered by the invite. No rotation. An explicit, plainly worded warning that if every device in a circle is lost, that circle's history is gone.
  - *Later.* Rotation when someone is removed, support for a person's second device, and social recovery, where an existing member approves a new device rather than a recovery phrase being typed by someone's grandmother.

## Considered alternatives

- **Plaintext on the server, protected by a policy.** The status quo, and what every competitor does. Rejected because the project's stated principle is that guarantees come from the code, and because holding other families' photographs in readable form is a liability with no upside once the server has no use for the plaintext.
- **Full key management in the first release.** Device verification, cross-signing, key backup, history sharing for late joiners. Rejected on evidence rather than taste: Element's unable-to-decrypt tracking issue has been open since 2022, history for late joiners took years and two CVEs to land, and FUTO Circles built exactly this product for exactly this audience and archived it.
- **Per-person or per-device encryption, as first proposed.** Rejected because it does not describe a group. A circle whose members cannot read each other is not a circle.
- **Encrypt media only, leave text readable.** Rejected as the worst of both: all of the key management, none of the guarantee.

## Consequences

Positive:

- Privacy stops being a promise and becomes a property. The honest sentence in the privacy policy changes from "we do not read your moments" to "we cannot".
- Legal exposure drops sharply. A breach yields ciphertext. A demand for content can only be answered with metadata, because the content is not ours to produce.
- Hosting other people's circles becomes a much smaller thing to take on, which is what makes the paid, project-hosted model in this decision safe to offer at all.
- Nothing is lost, because the product already refused every feature that would need plaintext: no search, no ranking, no counts, no server-side thumbnails.

Negative / acknowledged cost:

- **Metadata remains.** We know who is in which circle, and when they posted. That is not nothing, and the privacy policy must say so in the same plain words it uses for everything else.
- **Lose every device in a circle and its history is gone.** There is no password reset for content, and the first release has no recovery path beyond another member still holding the key.
- **Removing someone is forward-looking only.** They keep what they could already read, which is also true of every conversation anyone has ever had.
- **We cannot inspect reported content**, so a report has to carry the reporter's own decrypted copy. Apple still requires the reporting mechanism to exist, so this has to be designed rather than waved at.
- **Export can only be produced by a device holding the key.** The server cannot hand anyone their archive.
- **A leaked invite link is a leaked key** until that invite is revoked, because the link is what admits people.

## What would prove this wrong

- The first real circle of ten people generates recurring "I lost my photos" incidents. That would mean staged key management was too thin to ship.
- Somebody joins on a second device and cannot do it without help from the developer.
- A hosting customer needs content removed that we cannot remove, and it becomes a legal problem rather than a policy statement.
- The invite-carries-the-key mechanism turns out to leak in practice, for example because links survive in group chats far longer than anyone expects.
