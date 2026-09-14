# Nah?

Read the codebase wiki before exploring source: [docs/wiki/quickstart.md](docs/wiki/quickstart.md). It covers how the server, the app, the spike and delivery fit together, how to run each part, and where the gotchas are, and its Map links every page. Go to the code for exact behaviour, types and syntax.

## Keeping the wiki true

- The reasoning behind the architecture lives in the decision records in `docs/decisions/`. The wiki links to them rather than repeating them.
- When a file listed in `.wiki-compiler.json` changes (a README, a decision record, a workflow, the Dockerfile), run `/wiki-compile` so the wiki follows.
- The wiki follows the wiki doctrine: `quickstart.md` is the only entry point, pages explain why and how to operate but never restate code, each idea has one home, and there are no stubs.
