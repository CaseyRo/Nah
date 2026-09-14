---
concept: Invisible deploys
last_compiled: 2026-09-14
topics_connected: [server, mobile, delivery]
status: active
---

# Invisible deploys

## Pattern

The server is meant to be deployed constantly, and nobody using the app should ever be able to tell. That single requirement turns up in places that do not look related: how sessions are signed, how the feed remembers who posted last, how migrations are written, what the app does with a 401, and what an old app does with a moment written by a newer one. None of them is about deployment on its face.

## Instances

- **2026-09-14** in [server](../topics/server.md): the feed keeps each person's newest-moment time in memory and rereads it from their file every 30 seconds, so a moment taken by a second instance on the same volume cannot stay hidden for longer ([spike results](../../../spike/RESULTS.md)).
- **2026-09-14** in [mobile](../topics/mobile.md): a 401 signs the app in again without anyone noticing, and every blob starts with a byte naming its seal, so moments posted before encryption stay readable after it ([app README](../../../apps/mobile/README.md)).
- **2026-09-13** in [server](../topics/server.md): the first sign-in design kept sessions in memory and accepted logging everyone out on each deploy. It was revised the same day to tokens that carry their own claims, signed by a key kept on the volume ([ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md)).
- **2026-09-13** in [mobile](../topics/mobile.md): the posting app writes a plain-text fallback into every envelope, because phones update when they feel like it and an old app still has to show something sensible ([ADR-0016](../../decisions/0016-the-moment-envelope.md)).
- **2026-09-13** in [delivery](../topics/delivery.md): deploys arrive by webhook, and a green push is explicitly not a release ([ADR-0014](../../decisions/0014-the-go-setup.md)).

## What This Means

The requirement is cheap only while there is one instance and one app version. Each choice above cost tens of lines because it was made before a second instance or a second app version existed. Retrofitting any of them once people depend on the service is where the expense is.

It also points at the next bug before it ships: any state held in one process's memory breaks as soon as two processes share a volume. Sign-in challenges are still held that way ([ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md)), and two instances on one volume have only been simulated, inside a single test process, never run as two servers. Until they are, invisibility is a design target rather than a measured property, and a brief stop and start is the conservative way to deploy.

## Sources

- [topics/server.md](../topics/server.md)
- [topics/mobile.md](../topics/mobile.md)
- [topics/delivery.md](../topics/delivery.md)
