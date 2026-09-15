# Mobile app

## Purpose [coverage: high -- 5 sources]

Sources span 2026-04-26 to 2026-09-15. The app was scaffolded on 2026-09-13 and reached the M1 walking skeleton on 2026-09-14.

The app is how a person uses Nah?: a Flutter app for iOS and Android that talks only to the [server](server.md). As of M1 it gives the phone an identity of its own, joins by invitation and signs in, shows one page of the feed from you and the people you are connected to, posts text moments, and connects two people by pasting an invitation.

Much of what has been decided about the app is not built yet: ritual onboarding, voice, photo and music moments, reactions, the daily digest, and connecting by touching phones.

## Architecture [coverage: medium -- 2 sources]

The [app README](../../../apps/mobile/README.md) is the canonical description of what M1 does. The code:

- [`lib/main.dart`](../../../apps/mobile/lib/main.dart): the one screen: joining by invitation, then the composer, the feed and the connect sheet.
- [`lib/feed_cubit.dart`](../../../apps/mobile/lib/feed_cubit.dart): that screen's states, with every failure turned into a sentence.
- [`lib/server.dart`](../../../apps/mobile/lib/server.dart): the device identity, sign-in and the calls to the server. It is deliberately thin until a client package the app owns replaces it.
- [`lib/moment.dart`](../../../apps/mobile/lib/moment.dart): writing a moment's envelope and reading one back.
- [`test/`](../../../apps/mobile/test/): widget tests against a fake server, envelope tests full of hostile input, and a two-person test against a real server.

The Flutter version is pinned in [`mise.toml`](../../../mise.toml) at the repository root.

## Talks To [coverage: medium -- 3 sources]

- **The server**, at an address fixed when the app is built (`--dart-define=NAH_SERVER`), never typed by a person.
- **The keychain**, through `flutter_secure_storage`, which holds the device key's seed and the person id on this device only and out of backups ([ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md)).
- **`package:cryptography`** for the Ed25519 key, **`flutter_bloc`** for state and **`dio`** for requests.

## Key Decisions [coverage: high -- 13 sources]

Newest first. Each decision record is the canonical home for its reasoning.

- **2026-09-15:** one device per person, and a new phone arrives only through recovery that two people in the circle vouch for. An invite carries a one-time secret instead of the content key, which rotates when a connection ends. The first run asks "What do you want to share with the ones closest to you right now?" and offers an avatar photo. See the notes on [ADR-0005](../../decisions/0005-ritual-onboarding.md), [ADR-0012](../../decisions/0012-encrypted-on-device.md) and [ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md). Not built.
- **2026-09-15:** reactions return as five illustrations that only the poster sees, sealed and never counted ([ADR-0018](../../decisions/0018-reactions-the-poster-sees.md)). Not built.
- **2026-09-15:** music is a fourth moment type, shared from any music app and looked up on the phone. Photo and voice may carry a short line, any moment may carry a place kept only as words, and the poster can delete a moment, which leaves a marker. See the notes on [ADR-0006](../../decisions/0006-three-moment-types.md), [ADR-0008](../../decisions/0008-mvp-scope.md) and [ADR-0016](../../decisions/0016-the-moment-envelope.md). Not built.
- **2026-09-15:** the app calls a person's 150 their circle (the note on [ADR-0017](../../decisions/0017-one-network-of-a-hundred-and-fifty.md)). The M1 screens still say "your people".
- **2026-09-14:** moments travel unencrypted until each person has a content key (CDI-1863). The first byte of every blob names its seal, so what is posted now stays readable once encryption lands ([app README](../../../apps/mobile/README.md)).
- **2026-09-13:** one feed and one audience, with no audience picker; connecting is physical first and the link is the exception ([ADR-0017](../../decisions/0017-one-network-of-a-hundred-and-fifty.md)). M1 stands in for both with a pasted invitation.
- **2026-09-13:** a moment is a versioned envelope whose fallback sentence is written by the app that posts it, and every envelope read back is hostile input ([ADR-0016](../../decisions/0016-the-moment-envelope.md)).
- **2026-09-13:** identity is a key made on the device, and the identity key is not the content key ([ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md)).
- **2026-09-13:** content is encrypted on the device ([ADR-0012](../../decisions/0012-encrypted-on-device.md), not built until CDI-1863).
- **2026-04-26:** the first release has no comments, drafts or editing ([ADR-0008](../../decisions/0008-mvp-scope.md)). It had no reactions either, until ADR-0018 brought them back.
- **2026-04-26:** notifications default to one daily digest, with nothing sent on a day with nothing in it ([ADR-0007](../../decisions/0007-respectful-notifications.md)). Not built.
- **2026-04-26:** three moment types, text, voice and photo ([ADR-0006](../../decisions/0006-three-moment-types.md)), with music added on 2026-09-15. Only text exists so far.
- **2026-04-26:** onboarding is a four-beat ritual ([ADR-0005](../../decisions/0005-ritual-onboarding.md)). Not built.
- **2026-04-26:** no counts anywhere ([ADR-0004](../../decisions/0004-no-counts-anywhere.md)), which is why the composer's length limit shows no counter.
- **2026-04-26:** Flutter rather than a web app, with the repository pivoting on 2026-05-19 ([ADR-0001](../../decisions/0001-flutter-over-pwa.md)).

## Running It [coverage: medium -- 2 sources]

From `apps/mobile`, with a server running on the same machine and, for the first person on it, an invitation from `go run ./cmd/server invite` in `apps/server`:

```bash
flutter run
flutter analyze && flutter test
```

If Flutter is not on your path, prefix either with `mise exec --`, which uses the pinned version. The two-person test is skipped unless `NAH_SERVER` points at a running server; the [app README](../../../apps/mobile/README.md) has the full command.

## Gotchas [coverage: medium -- 4 sources]

- **The server address is compiled in.** The iOS simulator reaches `127.0.0.1`; an Android emulator needs `10.0.2.2`.
- **Moments are not encrypted yet.** Until CDI-1863, anyone holding the server's data directory can read them. See [opaque moments](../concepts/opaque-moments.md).
- **The connect sheet shows a raw `person#invite` string.** It is a stand-in for the touch and the link, not the intended experience.
- **A server that no longer knows this phone's person sends the phone back to asking for an invitation.** That is what a wiped development server looks like. See `start` in `lib/server.dart`.
- **The two-person test needs a fresh operator invitation for every run**, because each one is single-use. The command is at the top of `test/two_people_test.dart`.
- **No workflow checks the app on push.** Its checks run locally ([delivery](delivery.md)).

## Sources

- [apps/mobile/README.md](../../../apps/mobile/README.md)
- [mise.toml](../../../mise.toml)
- [README.md](../../../README.md)
- [docs/decisions/0001-flutter-over-pwa.md](../../decisions/0001-flutter-over-pwa.md)
- [docs/decisions/0004-no-counts-anywhere.md](../../decisions/0004-no-counts-anywhere.md)
- [docs/decisions/0005-ritual-onboarding.md](../../decisions/0005-ritual-onboarding.md)
- [docs/decisions/0006-three-moment-types.md](../../decisions/0006-three-moment-types.md)
- [docs/decisions/0007-respectful-notifications.md](../../decisions/0007-respectful-notifications.md)
- [docs/decisions/0008-mvp-scope.md](../../decisions/0008-mvp-scope.md)
- [docs/decisions/0012-encrypted-on-device.md](../../decisions/0012-encrypted-on-device.md)
- [docs/decisions/0015-no-passwords-a-key-on-the-device.md](../../decisions/0015-no-passwords-a-key-on-the-device.md)
- [docs/decisions/0016-the-moment-envelope.md](../../decisions/0016-the-moment-envelope.md)
- [docs/decisions/0017-one-network-of-a-hundred-and-fifty.md](../../decisions/0017-one-network-of-a-hundred-and-fifty.md)
- [docs/decisions/0018-reactions-the-poster-sees.md](../../decisions/0018-reactions-the-poster-sees.md)
