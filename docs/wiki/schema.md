# Wiki schema

The structure and conventions of this wiki. It was generated on the first compile. Edit it by hand to rename, merge or add topics, and the compiler follows it on the next run. The compiler adds what it discovers and never removes a topic or a concept without approval.

## Topics

- `server`: the Go server in `apps/server`.
- `mobile`: the Flutter app in `apps/mobile`.
- `spike`: the measurement harness in `spike`.
- `delivery`: hooks, workflows, the container image and the toolchain pin.

## Concepts

- `invisible-deploys`: a deploy must never be noticed by a person using the app. Connects server, mobile and delivery.
- `opaque-moments`: the server never reads a moment. Connects server, mobile and spike.

## Article structure

Topic pages use these sections, in order:

1. Purpose (required): what the part does, who depends on it, and where it sits.
2. Architecture: key files and entry points, pointing into the code rather than restating it.
3. Talks To: what it depends on and what depends on it.
4. Key Decisions: newest first, each line dated and linked to its decision record.
5. Running It: how to run, test, build and deploy it.
6. Gotchas: known issues, edge cases, failure modes and open questions.
7. Sources (required): every file that contributed.

Every heading except Sources carries a coverage tag, `[coverage: level -- N sources]`. The level is high for five or more sources, medium for two to four, and low for one or none.

Concept pages use Pattern, Instances, What This Means and Sources, with front matter naming the topics they connect. A concept exists only when it connects three or more topics.

## Conventions

- `quickstart.md` is the only entry point, and its Map lists every page. There is no `INDEX.md` or `CONTEXT.md`.
- The wiki explains why, how things fit together and how to operate them. It links to code, READMEs and decision records instead of restating them, and each idea is explained in exactly one place.
- A folder exists only while it holds several substantive pages. There are no stubs.
- Slugs are lowercase-kebab-case, links are relative markdown links, and dates are written year-month-day.
- No em-dashes and no all-caps words.
- The sources are the knowledge files listed in `.wiki-compiler.json`. Blog posts, research and openspec changes are left out on purpose, because that knowledge is kept elsewhere.

## Evolution log

- 2026-09-14: initial schema from four topics and two concepts. `quickstart.md` replaces the compiler's `INDEX.md` and `CONTEXT.md`, and Running It replaces its `API Surface` and `Data` sections, to follow the repository's wiki doctrine.
