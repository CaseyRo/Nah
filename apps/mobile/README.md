# Nah? app

A private home for your closest people.

Flutter, pinned to 3.47.4 in the repo's `mise.toml`. Cubit for state, Dio for
HTTP, `cryptography` for the device's Ed25519 key, and `flutter_secure_storage`
to keep that key in the keychain.

## Run it against a local server

```bash
(cd ../server && go run ./cmd/server) &
flutter run
```

The server address is a build setting, `--dart-define=NAH_SERVER=...`, and
defaults to `http://127.0.0.1:8080`, which the iOS simulator reaches. An Android
emulator needs `http://10.0.2.2:8080`.

## What M1 does

- **Signs in with a key, not a password** ([ADR-0015](../../docs/decisions/0015-no-passwords-a-key-on-the-device.md)).
  On first run the app makes an Ed25519 key, keeps it in the keychain on this
  device only, registers as a person and signs a challenge. A 401 signs in
  again without anyone noticing.
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
identities against a real server, connected, each posting, each reading both.
It is skipped unless `NAH_SERVER` points at a running `apps/server`.
