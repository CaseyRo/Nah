---
concept: Opaque moments
last_compiled: 2026-09-14
topics_connected: [server, mobile, spike]
status: active
---

# Opaque moments

## Pattern

The server stores, orders and returns moments without ever reading one. That began as a consequence of the product rather than as a privacy feature: Nah? has no search, no ranking and no counts, so the server never had a reason to read anything ([ADR-0010](../../decisions/0010-small-server-not-mastodon.md)). Once that was true, encrypting moments on the device became a client change instead of a server rewrite ([ADR-0012](../../decisions/0012-encrypted-on-device.md)), and each later piece was built to keep it true.

## Instances

- **2026-09-14** in [mobile](../topics/mobile.md): moments travel unencrypted until each person has a content key (CDI-1863), yet the app already treats every envelope it reads as hostile input, because nothing upstream can check it ([app README](../../../apps/mobile/README.md)).
- **2026-09-14** in [server](../topics/server.md): the feed shape was chosen on numbers alone. With opaque blobs the only questions are how many files and how many transactions, never what is inside ([spike results](../../../spike/RESULTS.md)).
- **2026-09-13** in [spike](../topics/spike.md): the fixture's moments are 400-byte blobs whose content is irrelevant, which is what let one fixture stand for every candidate ([thresholds](../../../spike/THRESHOLDS.md)).
- **2026-09-13** in [server](../topics/server.md): a moment's type lives inside the envelope rather than in a column, so the server does not even learn that someone sent a voice moment at eleven at night ([ADR-0016](../../decisions/0016-the-moment-envelope.md)).

## What This Means

The refusals and the privacy are one decision seen from two sides. Every feature that would need the server to read, such as counts ([ADR-0004](../../decisions/0004-no-counts-anywhere.md)), ranking, search or server-side thumbnails, is also a feature that breaks the promise that the project cannot read what people share. Proposing any of them means reopening ADR-0012.

The cost lands on the phone and on moderation. Validation happens only on the reading phone, a report has to carry the reporter's own decrypted copy, and metadata such as who is connected and when they posted stays visible to the server ([ADR-0012](../../decisions/0012-encrypted-on-device.md)).

For now the guarantee is that the server does not read, not yet that it cannot: until CDI-1863 lands, anyone holding the data directory can open a moment.

## Sources

- [topics/server.md](../topics/server.md)
- [topics/mobile.md](../topics/mobile.md)
- [topics/spike.md](../topics/spike.md)
