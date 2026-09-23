# Your circle

Linear: [CDI-1883](https://linear.app/cdit/issue/CDI-1883)

## Why

Who is in a person's circle is the whole product (ADR-0017). The January version of this change specified friend requests, an inner circle for selective sharing, an "X/150" ring and suggestions to review inactive friends. Every one of them contradicts a decision made since. M1 built the skeleton of the real model: invitations that register and connect, and a limit of 150 counted on both sides, with a pasted code standing in for the touch and the link. This restatement specifies joining and connecting as M2 will build them, and ending a connection, which nothing does yet.

## What Changes

- **joining**: an invitation registers a person and connects them to whoever invited them, in one step. The first person on a server uses an operator's invitation. An invitation survives installing the app. The first run starts with the person's name (CDI-1895), follows ADR-0005's ritual and shows only real moments.
- **connecting**:
  - touching phones, with a scanned code where NFC is not available
  - a link for people far away, with its secret after the `#`
  - every invitation works once, never expires, and can be withdrawn by whoever made it
  - refusals are said in a sentence that says whose circle is full, theirs or yours, with no number and no name
  - everyone is on one server until M7; the directory moves there (ruled 2026-09-23)
- **your-circle**: "your circle" is the product's name for the single network of ADR-0017. This covers seeing the people in it, and ending a connection without the other person being told.
- **Removed from the January version**: friend requests and accepting them, the inner circle, the "X/150" counter and ring, inactive-friend suggestions, and profile links as a way to connect.
- **Moved here from CDI-1882**: onboarding and the full-circle message.

## Capabilities

### New Capabilities

- `joining`: arriving in Nah? by invitation, and the first run.
- `connecting`: the touch, the link, the life of an invitation, refusals, and one server until M7.
- `your-circle`: the people a person is connected to, and ending a connection.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- `apps/server`: a person's invitations become single-use and revocable, and still never expire. A route ends a connection. The refusal for a full circle says whose it is. The directory waits for M7.
- `apps/mobile`: the touch and the scanned code, opening links, resuming after install, the first-run ritual, the circle screen, and ending a connection.
- Linear: CDI-1839 asks for short-lived links and CDI-1842 for expiry. Both predate this ruling and need rewording. CDI-1842 also leaves the full-circle sentence open, and the ruling is to say whose circle it is. CDI-1838, the directory, moved to M7 on 2026-09-23.
- Words: "your circle" replaces "network" and "your people" in the app, `DESIGN.md` and the specs, and ADR-0017 gains a note saying so (ruled 2026-09-15).
