# Reactions design

## Context

Nothing in the app or the server handles reactions today. The January design had a heart button on every card that opened a picker of five custom illustrations, a double-tap shortcut, reactions shown to everyone with avatars, and view receipts on by default. ADR-0004 forbids counts, ADR-0008 forbids visible read tracking, and cap-02 keeps every circle private, so most of that could not stand. CDI-1856 asked for the question to be decided in a decision record first, and ADR-0018 is that record.

## Goals / Non-Goals

**Goals**

- A poster hears back.
- Nothing about a reaction can be counted, compared or performed for.

**Non-Goals**

- Comments (cap-09) and messaging (cap-05).
- Notifications about reactions (see Open Questions).
- Building anything. This round is specification only.

## Decisions

### Five illustrated reactions

Ruled 2026-09-15 (ADR-0018). A small set can say how a moment landed, not only that it arrived: a sad moment and a happy one deserve different answers. They are Nah?'s own illustrations rather than system emoji, so they look the same on every phone.

### Only the poster sees them

Ruled 2026-09-15 (ADR-0018). Reactions shown to the whole circle would give people something to compare and perform for. Names on a moment would also tell some readers who else is in the poster's circle, which cap-02 says nobody can see.

### One reaction per person, and it can change

A person gives a moment at most one reaction, and can change it or take it back. Five reactions are a choice of how a moment landed, and people change their minds.

### The kind of reaction is sealed

The server stores that a person reacted to a moment and when, and the reaction itself travels sealed, as a moment's type does (`docs/wiki/concepts/opaque-moments.md`). Until CDI-1863 gives everyone a content key, reactions travel unsealed behind a seal byte of 0, exactly like moments.

### No view receipts

ADR-0008 rules out visible read tracking, and a list of who saw a moment is the quiet version of a count. Nah? records nothing about who looked.

### Values in `DESIGN.md`

The reaction button and picker values come back into `DESIGN.md` as they were before 2026-09-15. The spec links there for sizes and motion.

## Risks / Trade-offs

- **Five illustrations have to be drawn**, and have to stay legible at 24px.
- **The server learns that someone reacted to someone's moment.** It already knows the two are connected (ADR-0012's stated exception).
- **Silence still happens.** Reactions make it rarer. They cannot end it.
- **Nobody sees that a moment was well received** except its poster, so there is no crowd to join. That removes a pull to react, on purpose.

## Open Questions

- Whether reactions notify the poster at all, or appear in the daily digest (ADR-0007).
- Whether a poster can react to their own moment.
- What happens to a reaction when the two people's connection ends.
