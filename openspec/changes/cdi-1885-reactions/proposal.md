# Reactions

Linear: [CDI-1885](https://linear.app/cdit/issue/CDI-1885)

## Why

ADR-0008 left reactions out, and the retention research names that as the riskiest rule in the product: in a small circle, a moment met with silence is how the person who posts stops, and everyone else drifts out behind them (CDI-1856, `docs/research/what-makes-people-come-back.md`). The January version of this change specified five illustrated reactions shown to everyone, view receipts, and a "who viewed" list with a count. ADR-0018, accepted 2026-09-15, keeps the five reactions and removes everything that turns them into an audience or a score.

## What Changes

- **reactions**:
  - five illustrated reactions: Smile, Wink, Sad, Wow and Love
  - one reaction per person per moment, which they can change or take back
  - only the poster sees who reacted, by name and face, and nobody sees a number
  - the kind of reaction is sealed, so the server knows only that someone reacted
  - reactions are deleted with their moment
- **Decision recorded**: ADR-0018, amending ADR-0008.
- **`DESIGN.md`**: the reaction button section, removed on 2026-09-15 while reactions were out, returns with its January values.
- **Removed from the January version**: reactions visible to everyone, view receipts and the setting for them, the "who viewed" list and its count, "+N more", and the double-tap shortcut.

## Capabilities

### New Capabilities

- `reactions`: reacting to a moment, and what the poster and everyone else can see of it.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- `apps/mobile`: the reaction button and picker on every card, and the poster's view of who reacted.
- `apps/server`: storing a sealed reaction, returning a moment's reactions only to its poster, and deleting them with the moment.
- Design: five illustrations, legible at 24px.
- Linear: CDI-1856 is decided by ADR-0018 and can be closed.
- `docs/decisions`: ADR-0018 is new and ADR-0008 is amended, so the codebase wiki needs recompiling.
