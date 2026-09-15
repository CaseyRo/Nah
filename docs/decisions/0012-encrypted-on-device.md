---
title: "ADR-0012 — Content is encrypted on the device, and we hold no keys"
permalink: /decisions/0012-encrypted-on-device/
layout: page
status: Accepted
date: 2026-09-13
---

## Status

Accepted (2026-09-14), after being proposed on 2026-09-13. Extends [ADR-0010](0010-small-server-not-mastodon.html), and moves self-hosting out of the first release.

**Amended 2026-09-15.** The content key no longer travels in the invite. An invite carries a one-time secret after the `#`; the key, with its history, reaches the new connection only after the connection completes, sealed to that person's device key and checked against the secret, so a screenshot of a used invite unlocks nothing. A person's key rotates whenever one of their connections ends or they block someone, and the new key is sealed to the device keys their phone received from each remaining connection, never to keys the server supplies. Earlier moments stay readable to those who could read them. Recovery and one device per person are in the `cdi-1865-recovery-and-key-rotation` OpenSpec change.

Restated for [ADR-0017](0017-one-network-of-a-hundred-and-fifty.html), which replaced circles with one network of up to 150 per person. The guarantee is unchanged; the noun is not. Where this record first said a circle has a content key, a person now has one, handed over when two people connect.

The M1 walking skeleton sends moments unencrypted until each person has a content key (CDI-1863). The first byte of every blob already says how it is sealed, so nothing posted before encryption arrives needs migrating.

## Context

The September audit noted that Nah? never searches, ranks or counts anything, so its server never needs to read a moment. That made encryption a client change rather than a server rewrite, and parked it as a later possibility.

The proposal is to stop parking it. Encrypt on the device from the start, so privacy is structural rather than promised, and simplify the first release by having the project host everyone rather than asking families to run servers.

One correction to the framing before the decision. Encryption here cannot be per device, because a moment has to be readable by everyone its author is connected to. A person has a key, and the people they are connected to hold it. "Only the author's device can read it" would describe a diary, not a network.

## Decision

**The server stores opaque blobs, and the project holds boxes rather than keys.**

- Moments and media are encrypted on the device before upload and decrypted on the device after download. The server can store, order and deliver them. It cannot read them.
- **A person has a content key**, and it reaches someone when the two of them connect: in the touch, or in the invite link, in the part after the `#` that browsers never send to a server. The key therefore travels to the people who should hold it without the project ever seeing it, using the mechanism that already connects people.
- **The project hosts the first release.** Self-hosting moves to a later phase rather than out of the product. People still register centrally so they can be found by the people they connect with, and so a person can move hosts later.
- **Key management ships in stages**, because this is where family products die:
  - *First release.* One key per person, handed over when two people connect. No rotation. An explicit, plainly worded warning that moments can only be read on devices holding the key, so if every such device is lost, that history is gone.
  - *Later.* Rotation when a connection is removed, support for a person's second device, and social recovery, where someone you are connected to approves a new device rather than a recovery phrase being typed by someone's grandmother.

## Considered alternatives

- **Plaintext on the server, protected by a policy.** The status quo, and what every competitor does. Rejected because the project's stated principle is that guarantees come from the code, and because holding other families' photographs in readable form is a liability with no upside once the server has no use for the plaintext.
- **Full key management in the first release.** Device verification, cross-signing, key backup, history sharing for late joiners. Rejected on evidence rather than taste: Element's unable-to-decrypt tracking issue has been open since 2022, history for late joiners took years and two CVEs to land, and FUTO Circles built exactly this product for exactly this audience and archived it.
- **Per-device encryption, as first proposed.** Rejected because it does not describe a network. People who are connected but cannot read each other are not connected.
- **Encrypt media only, leave text readable.** Rejected as the worst of both: all of the key management, none of the guarantee.

## Consequences

Positive:

- Privacy stops being a promise and becomes a property. The honest sentence in the privacy policy changes from "we do not read your moments" to "we cannot".
- Legal exposure drops sharply. A breach yields ciphertext. A demand for content can only be answered with metadata, because the content is not ours to produce.
- Hosting other people's networks becomes a much smaller thing to take on, which is what makes the paid, project-hosted model in this decision safe to offer at all.
- Nothing is lost, because the product already refused every feature that would need plaintext: no search, no ranking, no counts, no server-side thumbnails.

Negative / acknowledged cost:

- **Metadata remains.** We know who is connected to whom, and when they posted. That is not nothing, and the privacy policy must say so in the same plain words it uses for everything else.
- **Lose every device holding a key and what it opened is gone.** There is no password reset for content, and the first release has no recovery path beyond someone you are connected to still holding the key.
- **Removing someone is forward-looking only.** They keep what they could already read, which is also true of every conversation anyone has ever had.
- **We cannot inspect reported content**, so a report has to carry the reporter's own decrypted copy. Apple still requires the reporting mechanism to exist, so this has to be designed rather than waved at.
- **Export can only be produced by a device holding the key.** The server cannot hand anyone their archive.
- **A leaked invite link is a leaked key** until that invite is revoked, because the link is what connects people. Making the touch the ordinary way to connect ([ADR-0017](0017-one-network-of-a-hundred-and-fifty.html)) makes the link the exception rather than the rule.

## What would prove this wrong

- The first ten real people using it generate recurring "I lost my photos" incidents. That would mean staged key management was too thin to ship.
- Somebody connects on a second device and cannot do it without help from the developer.
- A hosting customer needs content removed that we cannot remove, and it becomes a legal problem rather than a policy statement.
- The invite-carries-the-key mechanism turns out to leak in practice, for example because links survive in group chats far longer than anyone expects.
