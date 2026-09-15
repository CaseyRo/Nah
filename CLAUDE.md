# Nah?

Read the codebase wiki before exploring source: [docs/wiki/quickstart.md](docs/wiki/quickstart.md). It covers how the server, the app, the spike and delivery fit together, how to run each part, and where the gotchas are, and its Map links every page. Go to the code for exact behaviour, types and syntax.

## OpenSpec and Linear are 1:1

- Every change in `openspec/changes/` has exactly one Linear issue in the Nah? project, and carries its number: `cdi-<number>-<slug>`, lowercase, the same as the issue's git branch (`casey/cdi-<number>-<slug>`). The issue's title is what produces that slug.
- Issue first: create or find the issue before the change. Never a change without an issue, and never development on an issue without a change. A change found without one gets an issue backfilled.
- The line under a proposal's title links its issue: `Linear: [CDI-<number>](https://linear.app/cdit/issue/CDI-<number>)`.
- Two gates close the pair. The change is archived when its work is deployed on staging, and the issue moves to In Review then. The issue moves to Done when the work is deployed on production. Until production exists, staging is the fleet server (CDI-1831) and issues wait in In Review.
- A cancelled or out-of-scope change cancels its issue too.
- Build issues are sub-issues of their change's issue (for example CDI-1836 to CDI-1842 under CDI-1883), and are developed under that change.

## Keeping the wiki true

- The reasoning behind the architecture lives in the decision records in `docs/decisions/`. The wiki links to them rather than repeating them.
- When a file listed in `.wiki-compiler.json` changes (a README, a decision record, a workflow, the Dockerfile), run `/wiki-compile` so the wiki follows.
- The wiki follows the wiki doctrine: `quickstart.md` is the only entry point, pages explain why and how to operate but never restate code, each idea has one home, and there are no stubs.
