# Funding design

## Context

Nothing about funding exists yet. The January version offered donations alongside supporter perks, digital goods, feature voting and crowdfunded features. The audit on 2026-09-11 recommended a paid hosted service instead, on the Nabu Casa model. `product-principles` already rules out advertising and selling data, and ADR-0004 rules out counts.

## Goals / Non-Goals

**Goals**

- Pay for Nah? without changing anything for the people in it.
- Make the money and the plans visible to anyone.

**Non-Goals**

- Paid hosting, subscriptions or tiers of any kind.
- Building anything. This round is specification only.

## Decisions

### Donations only

Ruled 2026-09-15. Nah? is free to use in full, and paid for by voluntary donations made on its site through a donation platform. The app takes no payments.

Rejected: a yearly fee for a hosted service with free self-hosting, the direction the September audit recommended; and leaving the model undecided until M7.

### Open Collective, with its ledger as the public finances

Ruled 2026-09-15. Donations go through Open Collective, whose ledger shows every donation and expense publicly by default. The finances are therefore public continuously, rather than in periodic reports.

### Supporting changes nothing anyone can see

Ruled 2026-09-15. Inside a circle, a supporter badge is status, and status is what people perform for (ADR-0017). A supporter gets at most the donation platform's own thank-you. Because donations happen outside the app, Nah? never learns which person in it gave, and never tries to.

### Public money, public roadmap, nothing for sale

Ruled 2026-09-15. What Nah? receives and what it pays for are published, and so is the roadmap, and anyone can suggest something. Features are not voted on, because a vote is a count (ADR-0004), and not crowdfunded, because pledges make the roadmap a market.

### No money that needs growth

Carried over from January unchanged. Investment that needs Nah? to grow or be sold works directly against a network that cannot grow past 150 per person.

### One plain line in settings

Ruled 2026-09-15. The app's settings carry one plain sentence about supporting Nah?, with a link to the site, and nothing else in the app mentions it. Apple and Google have rules about apps that point people to payments outside the store, which differ between storefronts and have been changing, so the line ships only where the rules in force allow it.

### Suggestions wait until after the MVP

Ruled 2026-09-15. Most places for public suggestions count votes, and Nah? needs a setup that does not. Until after the MVP the roadmap is public, and there is no suggestion channel.

## Risks / Trade-offs

- **Donations may not cover hosting** once Nah? hosts other people's circles (M7).
- **A donation platform sees donors' payment details.** Donors are told which platform it is, and nothing from it flows into Nah?.
- **Most places for public suggestions count votes,** so the roadmap needs a home that does not.

## Open Questions

None left open: the last were settled on 2026-09-15.
