# Recovery and key rotation

Linear: [CDI-1865](https://linear.app/cdit/issue/CDI-1865)

## Why

ADR-0012's first release had one content key per person, handed over inside the invite link, never rotated, and lost for good with the phone. CDI-1865 named the three problems that decide whether families can live with that: a new phone, recovery, and rotation when someone is removed. Two things raised the stakes on 2026-09-15. A screenshot of an invite would hold a person's key. And recovery by vouching invites impersonation, above all by someone already inside a person's circle. This change settles how keys travel, when they change, and how a person gets back in, while keeping the property CDI-1865 insists on: a person confirms who can read, never the server.

## What Changes

- **content-keys**: an invite carries a one-time secret, never a content key. The key, with its history, follows a completed connection, sealed to the other phone and checked against the secret. A person's key rotates whenever one of their connections ends or they block someone, sealed to the device keys their phone already knows.
- **one-device**: a person uses Nah? on exactly one phone, deliberately. A new phone takes over only through recovery, even while the old one works.
- **recovery**:
  - two different people from the person's circle vouch in person, by touch or scanned code, never by link
  - 48 hours follow, in which the old phone and the whole circle are told who vouched, and anyone can stop it
  - then the identity moves, the old key stops working, and the person's own history comes back from a voucher's phone
  - someone with fewer than two connections cannot recover, and joins again
- **Elsewhere**: `cdi-1883-your-circle` no longer puts a key in the invite. ADR-0012 and ADR-0015 carry dated notes. CDI-1863 is reworded to match.

## Capabilities

### New Capabilities

- `content-keys`: how a content key reaches a connection, and when it rotates.
- `one-device`: one phone per person, and what that rules out.
- `recovery`: moving to a new phone through two people who know you.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- `apps/mobile`: sealing keys to device keys, keeping key history, rotation, the vouching touch, the 48-hour wait with its notices, and refusing a second device.
- `apps/server`: storing sealed keys it cannot open, a recovery request with its wait and its veto, replacing a person's device key when recovery completes, and refusing a second device key.
- Depends on content keys existing at all (CDI-1863, M2), and on push (M4) for the recovery notices.
- `docs/decisions`: notes on ADR-0012 and ADR-0015, so the codebase wiki needs recompiling.
