# Your name, and how you show yourself

Linear: [CDI-1895](https://linear.app/cdit/issue/CDI-1895)

## Why

Nobody in Nah? has a name. The specs, `DESIGN.md` and the site all say "Maya deleted this moment" and "Maya's circle is full", but nothing asks a person what they are called, the server keeps no name, and the M1 feed shows another person's moment with only its time. Two phones in CDI-1835 could tell "You" from a timestamp and nothing else.

On 2026-09-23 Casey ruled that a person gives the name everyone recognises them by as the first step after joining, and that they are also asked their gender, their pronouns and their sexual orientation, each with "rather not tell"; the pronouns are suggested from the gender answer and are the person's own choice. These are how people can show themselves to their circle, and later they can shape a person's page (CDI-1898, in the backlog).

## What Changes

- **who-you-are**:
  - a name, asked first after joining, which the person can change from their own page
  - gender, pronouns (suggested from gender) and sexual orientation, asked in the first run once profiles are sealed, each with "rather not tell"
  - all of it held in one sealed profile, shown only to the person's circle, and only after a connection completes
  - the app's own sentences use a person's name, never a pronoun
- **Elsewhere**:
  - CDI-1883's first run starts with the name, gains the three questions and a slow beat that greets the person by name, and keeps its avatar photo in the same profile. An invite link, the page for a phone without Nah? and a refusal before connecting name nobody.
  - CDI-1886's export includes the profile, leaving deletes it, and the privacy statement names what it holds.
  - ADR-0005 and ADR-0012 carry dated notes.

## Capabilities

### New Capabilities

- `who-you-are`: a person's name, gender, pronouns and orientation, and who can see them.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- `apps/server`: stores one opaque profile blob per person, and returns the profiles of the authors on a feed page in the same pass that reads their moments.
- `apps/mobile`: the name step, the three questions, the profile envelope, names on every moment, and editing the profile from your own page.
- Linear: CDI-1896 (the name, M1, blocks CDI-1835), CDI-1897 (gender, pronouns and orientation, blocked by CDI-1863), CDI-1898 (personalisation, backlog).
- `docs/decisions`: notes on ADR-0005 and ADR-0012, so the codebase wiki needs recompiling.
