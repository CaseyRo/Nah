# Nah? app

A private home for your closest people.

Flutter, pinned to 3.47.4 in the repo's `mise.toml`. Cubit for state, Dio for
HTTP, `cryptography` for the device's Ed25519 key, and `flutter_secure_storage`
to keep that key in the keychain.

## Run it against a local server

```bash
(cd ../server && go run ./cmd/server) &
(cd ../server && go run ./cmd/server invite)
flutter run
```

The second line prints a single-use invitation for the first person on the
server: paste it into the app when it asks. Everyone after that joins with an
invitation from someone already in.

The server address is a build setting, `--dart-define=NAH_SERVER=...`. For
real phones it is `https://nah.casey.berlin`, the dev server on the fleet. It
defaults to `http://127.0.0.1:8080`, which the iOS simulator reaches. An Android
emulator needs `http://10.0.2.2:8080`.

## What M1 does

- **Signs in with a key, not a password** ([ADR-0015](../../docs/decisions/0015-no-passwords-a-key-on-the-device.md)).
  On first run the app makes an Ed25519 key, keeps it in the keychain on this
  device only, and asks for an invitation, because Nah? is by invitation.
  Joining with it registers the person, connects them to whoever sent it, and
  signs a challenge. A 401 signs in again without anyone noticing.
- **Your name, asked first** (CDI-1896). Straight after joining, the app asks
  what the people closest to you call you, and every moment in the feed names
  who posted it. The name travels as a profile, unsealed like moments for now.
- **One page of feed, newest first** (CDI-1833). Pull down to reload it; it
  never loads more.
- **Text moments** (CDI-1834), written as an [ADR-0016](../../docs/decisions/0016-the-moment-envelope.md)
  envelope and read back as hostile input. They travel **unencrypted** until
  CDI-1863 gives each person a content key. The blob's first byte says how it
  is sealed, so what is posted now stays readable when encryption arrives.
- **Connecting by pasting an invitation**, standing in for the touch (CDI-1840)
  and the link (CDI-1839), which are M2.

## Checks

```bash
flutter analyze && flutter test
```

`test/two_people_test.dart` is CDI-1835 with the phones taken out: two
identities against a real server, one joining through the other's invitation,
each posting, each reading both. It is skipped unless `NAH_SERVER` points at a
running `apps/server` and `NAH_INVITE` holds a fresh invitation from
`server invite`; the comment at the top of the test has the command.
