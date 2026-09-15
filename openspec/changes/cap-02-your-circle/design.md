# Your circle design

## Context

M1 connects two people through one server route, `POST /v1/people/{person}/connections`, which redeems an invitation. Registering with a person's invitation connects the two in the same step. A touch and a link carry the same two things, who made the invitation and its secret, so the server cannot tell them apart. Today a person's invitation never expires, is not single-use and cannot be revoked; an operator's is single-use. Nothing ends a connection. Nobody's name is stored on the server: names travel inside the ciphertext.

## Goals / Non-Goals

**Goals**

- Connecting that needs no typing, no addresses and no approvals.
- Invitations that cannot quietly work twice.
- A way out of a connection.

**Non-Goals**

- Key rotation after a connection ends, second devices and social recovery (CDI-1865, M5).
- Reporting and blocking (CDI-1867).
- Leaving Nah? altogether (cap-07).
- Building anything. This round is specification only.

## Decisions

### "Your circle", one per person

Ruled 2026-09-15. "Network" reads as rigid and businesslike, and "your people" sounds one-directional. A circle drawn around one person fits a set of people who are each close to that person, and it says nothing about whether they know each other. It is not ADR-0009's circles: nobody shares a circle, nobody enters one, and each person has exactly one. ADR-0017 carries a note saying so.

### Single-use, never expiring, revocable

Ruled 2026-09-15. A person's invitation is now the only way a new person arrives, and it carries a content key, so it must not work twice. An expiry was rejected: the people who most need a link are the slowest to install, and a dead link may be the last thing that happens with them. Revocation covers the link that leaked into a group chat. This supersedes CDI-1839's "short-lived".

### A full circle is named

Ruled 2026-09-15. When the other person's circle is full, the refusal says it is theirs. The rejected alternative, one sentence that hides which side is full, was kinder to the person at 150 and baffling to the person refused. The sentence still carries no number (ADR-0004).

### Ending a connection is silent and on both sides

Either person can end a connection, and the other is not told. Each person's moments, earlier ones included, leave the other's feed and page, and a place opens in both circles. Ruled 2026-09-15: a connection that ends takes what it shared with it.

### The touch and the scanned code are one exchange

NFC where both phones allow it, and a code shown on one screen and scanned by the other where they do not (CDI-1840). Both carry the same invitation a link carries, so one server route serves all three.

### Only real moments on arrival

ADR-0005 allowed seeded moments from the inviter, "real or mocked". Mocked ones are out: `product-principles` says nothing appears in a feed that nobody in it posted. An inviter who has not posted is shown as a quiet feed.

### Resuming after install is the hard part

The invitation's secret lives after the `#`, so an install from the store has to carry it through without any server seeing it (CDI-1841). The spec fixes that outcome and leaves the mechanism to M2.

## Risks / Trade-offs

- **An unused invitation lives until it is revoked.** A link pasted into a group chat and never used can connect whoever finds it, and carries the content key. The mitigation is that a person can always see their unused invitations and withdraw any of them.
- **Naming whose circle is full tells someone that a person is at 150.** Accepted by the ruling.
- **An ended connection keeps the content key it was given** until rotation exists (CDI-1865). Until then, the server refusing to serve new moments is the only thing that keeps them out.
- **Physical first strands distant people if the link is weak** (ADR-0017). Resuming after install is where that is won or lost.

## Open Questions

- The first run's one real question: its wording, and whether the answer goes to the circle, to the person's own page, or nowhere (ADR-0005).
- Whether the maker of an invitation is told when it is refused because their circle is full.
