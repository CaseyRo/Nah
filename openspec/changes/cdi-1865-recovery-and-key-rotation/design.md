# Recovery and key rotation design

## Context

A person's content key seals their moments on their phone (ADR-0012), and their identity is a device key the server knows them by (ADR-0015). ADR-0012 planned to hand the content key over inside the invite link and never rotate it, and left losing the phone as losing the history. CDI-1865 asks for a second device, social recovery and rotation on removal, and warns that any design letting the server distribute keys lets a dishonest server add a reader.

Connecting already gives each person the other's device key and content key. That is what makes recovery through the circle possible: every person in someone's circle holds that someone's content key.

## Goals / Non-Goals

**Goals**

- A leaked or photographed invitation unlocks nothing.
- Removing someone from a circle keeps them out of what comes after.
- A lost phone does not have to mean a lost history.
- No step lets the server decide who can read.

**Non-Goals**

- A second device. Ruled out on purpose.
- A recovery phrase or recovery file.
- Building anything. This round is specification only.

## Decisions

### The key never travels in the invite

Ruled 2026-09-15. An invite carries a one-time secret after the `#`. Once the connection completes, each phone seals its person's content key, and every earlier key, to the other person's device key. For a link, it first checks that device key against the secret, so a server in the middle cannot swap in its own. For a touch, the device keys were exchanged in person. A screenshot of a used invite holds nothing, and a newcomer can still read the other person's earlier moments, as the first run (CDI-1883) and a person's page (CDI-1884) need.

Rejected: keeping the key in the invite and rotating it on every new connection, at most daily. That shortens how long a leaked key is useful, but does not remove it.

### Rotation on every ending and every block

Ruled 2026-09-15. When one of a person's connections ends, or they block someone, their phone makes a new content key for future moments. It seals that key to the device keys it received from each remaining connection, never to keys the server supplies. Earlier moments keep their keys. What someone could already read stays readable to them, which is honest: anything shown to a person once has been shown.

### One device, deliberately

Ruled 2026-09-15. A person uses Nah? on one phone. There is no second device, and a new phone takes over only through recovery, even when the old one still works, so there is one path for everyone and no quiet way to add a reader.

### Two people vouch in person, and the circle gets 48 hours

Ruled 2026-09-15, after asking how impersonation is stopped. The dangerous impostor is not a stranger but someone already in the person's circle, such as a controlling partner, who could otherwise vouch for a phone they hold. So:

- two different people from the circle must each touch the new phone, and a link can never vouch
- 48 hours follow, in which the old phone, if it still has Nah?, and everyone in the circle are told who vouched, and any of them can stop it
- no key moves until the wait ends

A stopped recovery can start again after 7 days, with two vouchers and the notice again, so it cannot be used to wear a circle down (ruled 2026-09-15).

When it ends, the identity moves to the new phone and the old device key stops working. The person's own content key and its history come back from a voucher's phone, which already holds it. Each other person's phone re-seals their key to the new phone when it next comes online, trusting it because two people from the circle signed for it.

### Fewer than two connections means starting again

Ruled 2026-09-15. Someone with fewer than two connections cannot recover. They join again with a new invitation, and their history is gone. That is the price of never letting one person, or the server, vouch alone.

## Risks / Trade-offs

- **Moving to a new phone is slow for everyone,** even with the old phone in hand: two people in person and two days.
- **Two vouchers can still be coerced together.** The 48-hour notice to the whole circle is what catches that.
- **The notice tells the circle that someone changed phones.** Accepted as the cost of the veto.
- **A person with one connection loses everything with their phone.** The privacy statement's lost-phone sentence (CDI-1886) has to say so.
- **Every rotation seals a key to up to 149 device keys,** and key history grows with each rotation.
- **Recovery depends on push** (M4) to reach the old phone and the circle in time.

## Open Questions

- The wording of the recovery notice.
