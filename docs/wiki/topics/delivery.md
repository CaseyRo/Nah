# Delivery

## Purpose [coverage: medium -- 4 sources]

Sources span 2026-09-13 to 2026-09-14.

Delivery is how a change gets checked, built and published: hooks before a commit, two GitHub workflows after a push, the [server](server.md)'s container image, and the pin that fixes the [app](mobile.md)'s toolchain. Since 2026-09-25 every push to `main` deploys the server to nebula-1 at `nah.casey.berlin`, the dev and dogfooding server, which is also staging until production exists.

## Architecture [coverage: medium -- 5 sources]

- [`.pre-commit-config.yaml`](../../../.pre-commit-config.yaml): markdown and yaml linting and file hygiene on every commit. Markdown is linted with `--fix`, so a commit can rewrite what you wrote, and generated platform code under `apps/mobile` is skipped.
- [`.github/workflows/server.yml`](../../../.github/workflows/server.yml): on changes under `apps/server`, vet, staticcheck, race-enabled tests and govulncheck, then static binaries for linux amd64 and arm64 kept as artifacts for seven days.
- [`.github/workflows/pages.yml`](../../../.github/workflows/pages.yml): builds the Jekyll site in `docs` and deploys it to GitHub Pages on changes under `docs`, except `docs/wiki`, which the site also excludes.
- [`apps/server/Dockerfile`](../../../apps/server/Dockerfile): a static build copied onto an empty base image, running as uid 65532 with `/data` as its volume and no shell inside.
- [`apps/server/compose.yaml`](../../../apps/server/compose.yaml): the fleet deploy. Komodo builds the image in place on nebula-1, keeps `/data` in a named volume, gives a stop 30 seconds to checkpoint every file, and binds the port to nebula-1's Tailscale address.
- [`mise.toml`](../../../mise.toml): pins Flutter, so the app's toolchain belongs to the repository rather than to whichever machine builds it.

## Talks To [coverage: medium -- 3 sources]

- **GitHub Actions and GitHub Pages**, through the two workflows above.
- **Komodo**, stack `git-nah-nebula` on nebula-1, by a GitHub push webhook ([ADR-0014](../../decisions/0014-the-go-setup.md)).
- **The fleet's Cloudflare tunnel**, which serves `nah.casey.berlin` with no Cloudflare Access in front, because the app signs in with its own key, and **Gatus**, which checks `/healthz` as `nah-dev` in `CDiT-dev/fleet-stacks`.
- **No container registry.** The fleet builds from git; a registry image waits for M7 (the 2026-09-25 note on [ADR-0014](../../decisions/0014-the-go-setup.md)).

## Key Decisions [coverage: medium -- 4 sources]

Newest first.

- **2026-09-25:** the dev server runs on nebula-1, built from git in place, at `nah.casey.berlin`, with no backups until Litestream, an accepted risk (CDI-1831, the note on [ADR-0014](../../decisions/0014-the-go-setup.md)).

- **2026-09-14:** the wiki lives in `docs/wiki` but is excluded from the published site and from the site's build trigger, so compiling it rebuilds nothing.
- **2026-09-13:** how the server ships: static binaries for amd64 and arm64, checks on every push, a Komodo webhook deploy, and version numbers owned by a release workflow rather than edited by hand ([ADR-0014](../../decisions/0014-the-go-setup.md)).
- **2026-09-13:** deploys must be invisible to people using the app ([ADR-0015](../../decisions/0015-no-passwords-a-key-on-the-device.md)). See [invisible deploys](../concepts/invisible-deploys.md).
- **2026-09-13:** the server workflow was written before the server existed, so the gate was in place before the first line of Go ([server.yml](../../../.github/workflows/server.yml)).

## Running It [coverage: medium -- 3 sources]

Hooks, once per clone and then on demand:

```bash
pre-commit install
pre-commit run --all-files
```

The server's checks and container commands are in the [server README](../../../apps/server/README.md), the app's checks in the [app README](../../../apps/mobile/README.md). To preview the site, run `bundle install` and then `bundle exec jekyll serve` from `docs`, which needs Ruby 3.3 or later.

## Gotchas [coverage: high -- 5 sources]

- **A commit can stop on its first attempt.** When the end-of-file or markdown hook rewrites a file, the commit aborts; stage the rewritten files and commit again.
- **Only server and site changes run anything on push.** App changes are checked locally or not at all.
- **Every push to `main` redeploys the server**, including pushes that only touch docs. The webhook secret is identical on GitHub and in the Komodo stack; if either is changed alone, every delivery answers 401 and deploys nothing, silently ([ADR-0014](../../decisions/0014-the-go-setup.md)).
- **The tunnel has connectors on nebula-1 and werkstatt-1**, so it reaches the server at nebula-1's Tailscale address. A port bound to localhost would fail about half the time.
- **Nothing backs up the data yet.** Losing the `nah-data` volume loses the circle, until Litestream.
- **A green push is not a deploy.** Check the webhook's recent delivery on GitHub and the stack's update in Komodo ([ADR-0014](../../decisions/0014-the-go-setup.md)).
- **The container has no shell and runs as uid 65532.** Health checks must come from outside, against `/healthz`, and a mounted data directory must be writable by that uid.
- **The server workflow's own comment records a red run** from the push that added it before any Go code existed.

## Sources

- [.pre-commit-config.yaml](../../../.pre-commit-config.yaml)
- [.github/workflows/server.yml](../../../.github/workflows/server.yml)
- [.github/workflows/pages.yml](../../../.github/workflows/pages.yml)
- [apps/server/Dockerfile](../../../apps/server/Dockerfile)
- [mise.toml](../../../mise.toml)
- [apps/server/README.md](../../../apps/server/README.md)
- [apps/mobile/README.md](../../../apps/mobile/README.md)
- [README.md](../../../README.md)
- [docs/decisions/0010-small-server-not-mastodon.md](../../decisions/0010-small-server-not-mastodon.md)
- [docs/decisions/0014-the-go-setup.md](../../decisions/0014-the-go-setup.md)
- [docs/decisions/0015-no-passwords-a-key-on-the-device.md](../../decisions/0015-no-passwords-a-key-on-the-device.md)
