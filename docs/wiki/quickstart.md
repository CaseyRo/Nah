# Nah? codebase wiki

Nah? is a private social network in which each person has one network of up to 150 people. Phones run a Flutter app that talks to one small Go server, and the server stores what it does not read. The product is described in the [root README](../../README.md) and the [decision records](../decisions/index.md); this wiki explains how the code fits together and how to run it.

## Run it

In two terminals, from the repository root:

```bash
cd apps/server && go run ./cmd/server
cd apps/mobile && flutter run
```

The server listens on `:8080`, which is where the app looks by default. The checks are `go vet ./... && go test -race ./...` in `apps/server` and `flutter analyze && flutter test` in `apps/mobile`.

## Where things are

- `apps/server`: the [server](topics/server.md).
- `apps/mobile`: the [app](topics/mobile.md).
- `spike`: the [measurements](topics/spike.md) that decided the server's shape. None of it ships.
- `docs`: the published site, the decision records, and this wiki, which the site does not publish.
- `.github`, `.pre-commit-config.yaml` and `mise.toml`: [delivery](topics/delivery.md).

## Where it stands, 2026-09-14

M1, the walking skeleton, works on one machine: two people can register, connect, post and read each other's moments, through the real app and the real server. What remains of M1 is two physical phones on a deployed server, and that deploy does not exist yet.

## Reading this wiki

Every section heading carries a coverage tag. `high` means the section can be trusted on its own, `medium` means the linked sources hold more detail, and `low` means read the sources. For exact behaviour, types or syntax, read the code: this wiki explains why things are the way they are and how to operate them, and it never restates what the code says.

## Map

- [quickstart.md](quickstart.md): this page.
- [topics/server.md](topics/server.md): the Go server, person files, sign-in and the feed.
- [topics/mobile.md](topics/mobile.md): the Flutter app, the device identity, moments and connecting.
- [topics/spike.md](topics/spike.md): the measurement harness, its thresholds and its results.
- [topics/delivery.md](topics/delivery.md): hooks, workflows, the container, the toolchain pin and the deploy still to come.
- [concepts/invisible-deploys.md](concepts/invisible-deploys.md): why a deploy must never be noticed, and everything that shapes.
- [concepts/opaque-moments.md](concepts/opaque-moments.md): why the server never reads a moment, and what that costs.
- [schema.md](schema.md): topics, sections and conventions, for whoever edits or recompiles this wiki.
- [log.md](log.md): what each compile changed.
