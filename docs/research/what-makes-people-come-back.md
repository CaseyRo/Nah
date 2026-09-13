---
title: "What makes people come back"
permalink: /research/what-makes-people-come-back/
layout: page
date: 2026-09-12
---

<style>
.fig{border:1px solid rgba(128,128,128,.28);border-radius:12px;padding:1rem;overflow-x:auto;margin:1.5rem 0}
.fig svg{display:block;width:100%;min-width:640px;height:auto;color:inherit}
.nd-box,.nd-soft{fill:none;stroke:currentColor;stroke-width:1.25;opacity:.55}
.nd-key{fill:rgba(238,52,35,.10);stroke:#EE3423;stroke-width:1.5;opacity:1}
.nd-region{fill:none;stroke:currentColor;stroke-width:1;stroke-dasharray:4 5;opacity:.4}
.nd-line{fill:none;stroke:currentColor;stroke-width:1.4;opacity:.75}
.nd-faint{fill:none;stroke:currentColor;stroke-width:1.25;opacity:.35}
.nd-ring{fill:none;stroke:currentColor;stroke-width:1.5;stroke-dasharray:6 5;opacity:.45}
.nd-ring.you{stroke:#EE3423;opacity:.9}
.nd-node{fill:none;stroke:currentColor;stroke-width:1.4;opacity:.75}
.nd-node.you{fill:rgba(238,52,35,.12);stroke:#EE3423;stroke-width:2;opacity:1}
.nd-h{fill:currentColor;font:800 16px inherit}
.nd-t{fill:currentColor;font-weight:700;font-size:13.5px}
.nd-s{fill:currentColor;opacity:.7;font-size:12px}
.nd-lab{fill:currentColor;opacity:.6;font-weight:700;font-size:11px;letter-spacing:.08em;text-transform:uppercase}
.nd-acc{fill:#EE3423;font-weight:700;font-size:12px}
.nd-in{fill:currentColor;font-weight:700;font-size:11px;text-anchor:middle}
.nd-in.you{fill:#EE3423}
.nd-arrowhead{fill:currentColor;opacity:.75}
.nc-grid{stroke:currentColor;opacity:.18}
.nc-axis{stroke:currentColor;opacity:.35}
.nc-line{fill:none;stroke:#EE3423;stroke-width:2;stroke-linejoin:round;stroke-linecap:round}
.nc-dot{fill:#EE3423}
.nc-tick{fill:currentColor;opacity:.65;font-size:11.5px}
.nc-val{fill:currentColor;font-weight:700;font-size:12px}
</style>

*How the private-social apps are really doing, which return mechanics actually work, and what both mean for an app built on purpose to be calm.*

## The short version

- **Nobody in this category publishes retention.** Not Retro, Yope, Locket, Marco Polo, FamilyAlbum, Lapse or Dispo. The closest thing to an audited number is Tinybeans, where about 5.7% of monthly users pay, and the company still loses money.

- **BeReal is the closest thing to Nah?'s idea that ever reached scale, and it fell from 20 million daily users to under 6 million in five months.** The peer-reviewed post-mortem blames performative drift and feature creep, not the one-post-a-day rule. People left because there was nothing to do, not because they were only allowed one moment.

- **The apps that hold people don't hold them by being calm.** They hold them because leaving costs something: an archive, a network, a vault. Calm technology has a design literature and no retention evidence.

- **One product did grow for years on a chronological, algorithm-free friend feed.** Letterboxd went from 1.8 million members to over 30 million on a subscription. Its engine is a catalogue and a logging habit, though, not friends.

- **For Nah?: plan a weekly product with a daily surface.** Add a widget and an archive, ship a reaction that carries no number, keep no counts, and charge for the room.

## How they're actually doing

Every figure below is labelled by where it came from. A company claim is what the company says about itself, usually to advertisers or investors. An estimate comes from an outside tracker such as Sensor Tower, Apptopia, Appfigures or Similarweb, and those are modelled, not measured.

| App | Scale | Engagement | What it signals |
|---|---|---|---|
| BeReal | 40M monthly users, company claim. An outside tracker put it at 16M in 2023 | 55% post inside the two-minute window, company claim. Peak churn of 20.7% in 2022 against under 10% for rivals, estimated | Sold for €500M, now ad-funded, profitability targeted for 2027 |
| Retro | About 1M users and 7M downloads, company claim | About 45.7% engage daily and over half of daily visitors post, company claim | 585 App Store ratings against a claimed million users. Raised $21M |
| Yope | 15M registered, company claim. 2.2M monthly users in March 2025 | Over half open it five or more days a week, company claim. 40% day-7 retention, reported in early 2025 | Markets itself on no algorithms and no ads, and ships streaks and widgets |
| Locket | 80M+ installs and about 9M daily users, company claim | None published | Added a weekly prompt and celebrity accounts in 2025 to keep growing |
| Airbuds | 5M monthly and 1.5M daily users, company claim | Roughly 30% daily-to-monthly ratio, which is genuinely good | Median 12 to 15 followers per user. No subscription at all as of late 2025 |
| Marco Polo | 2.2M App Store ratings, the largest in this set | None published | Profitable on $9.99 a month, and not growing |
| Tinybeans | About 900K monthly users, 51K paying | 93% subscriber retention, 7% churn, company reported | Revenue fell 11% in 2025, then rose 35% in 2026 on the back of an acquisition, reaching its first positive operating cash flow after fourteen years |
| FamilyAlbum | 30M users, company claim | None published | Grew ad-free for a decade, then added banner ads to the free tier in late 2025 |
| Lapse | 118K ratings | None published | Fell from 29th to 99th in its category and retreated to being a camera app |
| Noplace | 3.9K ratings since a number-one debut in July 2024 | None published | 3.9 stars. A chart-topping launch that left almost nothing behind |

Two numbers in that table are worth staring at. Retro has a claimed million users and 585 ratings, which is an unusual ratio for an app people love. And Locket, with nine million daily users, earns roughly half a dollar per daily user per year. That is the most sobering figure in this whole investigation for anyone planning to fund a private feed with subscriptions.

## The BeReal experiment

BeReal matters more than any other app here, because it is the only product that ever tested Nah?'s core bet at scale: one deliberate moment a day, no scrolling, no follower counts, no algorithm. It worked, spectacularly, and then it stopped working.

<figure>

<div class="fig">

<svg viewBox="0 0 720 320" role="img" aria-label="A line chart of BeReal's estimated daily active users. It rises from 10 million in August 2022 to a peak of 20 million in October 2022, then falls to 10.4 million by February 2023 and under 6 million by March 2023.">

<line x1="60" y1="45" x2="700" y2="45" class="nc-grid"/>
</line>

<line x1="60" y1="95" x2="700" y2="95" class="nc-grid"/>
</line>

<line x1="60" y1="145" x2="700" y2="145" class="nc-grid"/>
</line>

<line x1="60" y1="195" x2="700" y2="195" class="nc-grid"/>
</line>

<line x1="60" y1="245" x2="700" y2="245" class="nc-axis"/>
</line>

<text x="50" y="49" class="nc-tick" text-anchor="end">
20M
</text>

<text x="50" y="99" class="nc-tick" text-anchor="end">
15M
</text>

<text x="50" y="149" class="nc-tick" text-anchor="end">
10M
</text>

<text x="50" y="199" class="nc-tick" text-anchor="end">
5M
</text>

<text x="50" y="249" class="nc-tick" text-anchor="end">
0
</text>

<polyline points="60,145 240,45 600,141 690,185" class="nc-line"/>
</polyline>

<g>
<title>
August 2022: about 10 million daily users
</title>
<circle cx="60" cy="145" r="5" class="nc-dot"/>
</circle>
</g>

<g>
<title>
October 2022: about 20 million daily users, the peak
</title>
<circle cx="240" cy="45" r="5" class="nc-dot"/>
</circle>
</g>

<g>
<title>
February 2023: about 10.4 million daily users
</title>
<circle cx="600" cy="141" r="5" class="nc-dot"/>
</circle>
</g>

<g>
<title>
March 2023: under 6 million daily users
</title>
<circle cx="690" cy="185" r="5" class="nc-dot"/>
</circle>
</g>

<text x="240" y="32" class="nc-val" text-anchor="middle">
20M at the peak
</text>

<text x="686" y="207" class="nc-val" text-anchor="end">
under 6M
</text>

<text x="56" y="268" class="nc-tick">
Aug 2022
</text>

<text x="240" y="268" class="nc-tick" text-anchor="middle">
Oct 2022
</text>

<text x="600" y="268" class="nc-tick" text-anchor="middle">
Feb 2023
</text>

<text x="694" y="268" class="nc-tick" text-anchor="end">
Mar 2023
</text>

<text x="60" y="296" class="nc-tick">
Estimated daily active users, Apptopia and Sensor Tower
</text>

</svg>

</div>

<figcaption>
BeReal doubled and then lost 70% of its daily users inside eight months. The company later disputed outside trackers and claimed 25 million daily users in September 2023, a figure no independent source supports.
</figcaption>

</figure>

The useful part is why. A peer-reviewed study of 333 users, published in New Media and Society in November 2025, traced a path from enthusiasm to scepticism to disinterest. What people valued was social connection and a private memory archive. What drove them out was growing cynicism about authenticity, posting that became performative anyway, and feature creep. Trade coverage at the time of the collapse is blunter: users called it "pointless" and the authenticity "monotonous".

Read carefully, that is not an indictment of the one-a-day rule. People liked the constraint. What decayed was the reason to open the app once the novelty of seeing friends unposed wore off. There was nothing else in there.

The sequel matters too. Under Voodoo the app added public discovery, friends-of-friends, several posts a day, brand accounts and ads. That produced €30 million of revenue and an audience that stabilised rather than recovered. Every one of those additions is something Nah? has decided not to do, and they did not bring the users back.

## What brings people back

Each of these is a real mechanic in a shipping product, with a cue and a reward. The last column is the one that matters for Nah?, because a mechanic that breaks the rules isn't available.

| Mechanic | Who uses it | Does it work? | Fits Nah?'s rules |
|---|---|---|---|
| Ambient presence | Locket, Yope: a widget on the home or lock screen | Structurally the strongest. Returning costs no attention, so it cannot fatigue. No independent decay data exists | Perfect fit and Nah? currently has no ambient surface at all |
| A digest that arrives | Tinybeans, newsletters, Nah?'s own plan | Paid newsletters run 50 to 70% open rates. No numbers published for family apps | Already the plan. One a day is the ceiling of what people tolerate |
| Reciprocity in a small group | Marco Polo organically, BeReal by force | Well supported in the research on online communities: good contributions create a felt obligation to answer | Fits, as long as the obligation is to a person. "Mum posted" is not a dark pattern |
| A rhythm tied to a time | Retro weekly, Duolingo daily | The one habit intervention with trial evidence is anchoring a behaviour to an existing routine | Fits, and argues for a fixed digest time rather than BeReal's random one |
| Archive value | Retro's Rewind, Timehop, Day One | Timehop reached millions of daily users on nostalgia alone. BeReal's own users named the archive as what they valued | Fits completely, and it's the mechanic that makes year two better than year one |
| An audience with nowhere else to go | Tinybeans and FamilyAlbum: grandparents | Structural rather than measured. A fifth of Yope's active users invited an older relative | Fits, and may be the real moat |
| Streaks and counts | Snapchat, Duolingo, Yope | They work. Over nine million Duolingo users hold a year-long streak | Excluded by rule, and the Snapchat research shows why |
| An algorithmic feed | Instagram, TikTok | The only mechanic that works when nobody you know posted anything | Excluded, and impossible at 150 people anyway |

Snapchat's streaks are worth one extra note, because they are the purest version of engineered return. The research on them is consistent and unflattering: teenagers send blank photos of the floor purely to keep a streak alive, and breaking one is experienced as personal rejection. That is what a retention mechanic looks like when it has fully detached from the relationship it was meant to serve.

## What the research actually supports

- **Habits form slower than product folklore says.** The standard study found a median of 66 days to reach automaticity, a range from 18 to 254 days, and fewer than half of participants getting there at all.

- **Anchoring beats reminding.** The strongest evidence for building a routine is attaching it to something already in the person's day, rather than notifying them at a clever moment.

- **The habit-loop model popular in product circles is a heuristic, not a validated theory.** Its empirical core is variable reward from operant conditioning, which is exactly the part critics object to.

- **One notification a day is the ceiling of tolerance.** Around two thirds of people say they want one push a day or fewer, and nearly half opt out after two to five in a week. Vendor claims that daily notifications multiply retention are selection effects, not causation.

- **The "healthy social media" literature is weaker than it looks.** A meta-analysis of 141 studies and about 145,000 participants found mostly negligible effects, and concluded the return on fifteen years of research has been poor. One finding does help Nah?: passive use looks harmful in public feeds but not inside social groups. Lurking in a room full of people you know is not the same act as lurking in a stranger's feed.

- **Calm technology has no retention evidence.** It is a design philosophy, and a good one. It tells you what to build, not whether anyone comes back.

## The small-group risk

This is the part that should worry Nah? most, and it has a formal literature.

In small private communities, the famous 90-9-1 split of lurkers to contributors softens to something closer to 70-20-10. It never inverts. Most people in a room are reading, not posting, and that is normal rather than a failure.

The danger is what happens when a poster stops. The model built from Friendster's collapse treats a member as leaving when the benefit from their still-active friends no longer covers the effort of showing up. What survives is the subset where everyone still has at least a few active friends, and departures cascade from the edges inward. Hungary's iWiW died exactly this way: peripheral members first, then everyone.

In a 150-person room with one dominant poster, most members have exactly one reason to come back. Lose that person and the collapse isn't gradual. That is the single most likely way a Nah? home dies, and it is a product problem, not a technical one.

One more number, for honesty: nobody has published a minimum posting frequency for keeping a private group alive. Anyone who quotes you one is guessing.

## Does anyone pay?

Yes, but thinly, and for a specific framing.

| Product | Price | What it earns | Result |
|---|---|---|---|
| Tinybeans | About $75 a year | 51K payers out of about 860K monthly users, so 5.9% convert | Reached its first full year of positive operating cash flow in 2026, fourteen years in, after abandoning ads and buying a competitor. Its own paid base was flat for two straight years |
| Ente | About $65 a year | 20,703 paying out of 460,461 registered, so 4.5% convert | Open source, encrypted, self-hostable, and funded by the hosted plan |
| Locket | $3.99 a month or $36 a year | About $400K a month from 9M daily users, which is roughly half a dollar per daily user per year | Profitable, but also runs ads |
| Marco Polo | $9.99 a month | Around $800K a month | Profitable, ad-free, flat |
| FamilyAlbum | $5.99 a month | Not broken out | Added ads to the free tier in late 2025 |
| Retro | About $5 a month | In-app spend up sharply, but only a small subset subscribe | Venture-funded instead |

The pattern across all of them: **people pay for storage and for an archive. They do not pay for a feed.** Tinybeans sells a keepsake, Locket sells a widget, Marco Polo sells a way to talk to family. Nobody has sold a private timeline on its own.

Two findings make paying itself worth doing. Subscription apps retain about 14% of users at day 30 against 5.4% for ad-supported ones. And annual plans retain dramatically better than monthly ones, around 36% against 6.7% at the one-year mark. Paying for something is itself a reason to keep using it.

### The number to plan with

Two unrelated businesses, in adjacent categories, with nothing in common but a private-memories product, land on the same figure. Tinybeans converts 5.9% at about $75 a year. Ente converts 4.5% at about $65 a year. **Roughly one in twenty pays, at roughly $65 to $75 a year.** That is the strongest quantitative signal in this investigation, and it is the assumption to plan with rather than "surely a family would pay five euros a month".

One caveat that cuts in Nah?'s favour, and one that cuts against. In favour: both of those are freemium funnels, where most people never intended to pay. If a Nah? home is paid from the start, there is no funnel to convert, only a decision to make. Against: nobody in this set has tested whether a family will pay for a private space *before* they have felt it work, which is exactly what a paid-from-day-one home asks them to do.

### What the comparison class has already taught families

This is the finding that most directly challenges a per-group price. Across iCloud, Google One, Apple One, 1Password and Ente, the family offer is always the same sentence: *share with up to five more people at no extra cost*. Extra seats are a device for making the one payer's subscription harder to cancel, not a reason to charge more. Marco Polo's family tier is the only real exception in the set.

So a household comparing "a few euros a month for our room" is comparing it against a market that has trained them to expect the other people to be free. Basecamp is the one clean counter-model, a flat fee per organisation with no per-person pricing at all, but that is a business tool bought with a business budget.

### "Pay us so we don't need ads" is said everywhere and proven almost nowhere

Cocoon said it and shut down. Marco Polo said it, launched a genuinely additive paid tier, and then moved previously-free features behind the paywall anyway. Tinybeans *was* the ad business and spent four years escaping it, cutting advertising revenue from about $9M to about $1.3M while subscriptions failed to fill the gap.

The two places it works cleanly are instructive because neither has to convert anyone. Immich sells an optional lifetime key with an explicit promise that no feature will ever be paywalled, and it has a patron behind it. Home Assistant's cloud subscription has held the same price for four and a half years, under a foundation, with hardware revenue alongside. Donations alone do not work: Plausible's self-hosters contribute about $300 a month in total, and Mastodon converts well under 1% of its users into supporters.

## Why they die

These come from shutdown notices, founder interviews and company filings rather than from commentary.

### Not enough people, at any conversion rate

Cohost published its books when it closed, which almost nobody does. Expenses of $41,605 a month against $16,307 of income. 16,846 monthly users, 3,046 of them paying, 3.18% monthly churn, $6.39 per subscriber. That is an **18% conversion rate**, three to four times what anyone else in this report manages, and it was still fatal.

> We've managed to build a social media platform that many of our users love, but we just don't have enough users.

*Cohost's shutdown post, September 2024. Its first line names lack of funding and burnout, and notes that none of the team were being paid.*

For an invite-only product the base is capped on purpose, so no conversion rate rescues it. Nah? has one real advantage Cohost did not: a home costs cents a month to run rather than thousands. Its equivalent failure isn't the server bill, it's the founder's time.

### The maintenance outlives the founder's energy

Three teams, three completely different funding situations, one sentence between them. Cocoon closed because it had become "unsustainable to maintain a high quality experience". Cohost named burnout in its opening line. And FUTO Circles, which had a patron, a liberal licence and no revenue pressure at all, ended like this:

> Making this project into a successful commercial product would require a level of effort far beyond what was feasible for our small team.

*The Circles repository, archived February 2025*

### The group chat is the competitor, and the founders knew it

Cocoon's founders said both of the important things out loud in 2019, four years before they shut down. On why an ad model was impossible: "15 to 20 people are never going to create enough content to aggregate enough attention to build an ads business." On what they were actually up against: "the biggest competitor to Cocoon is what goes down in the small groups you have in iMessage or any of your other chat apps."

They were right on both counts, charged a subscription because of it, and died anyway. Being right about the market is not the same as beating it.

### The paid tier eats its own goodwill, and ads come back

Marco Polo launched its subscription in 2020 as a purely additive tier and was praised for it. The one-star revolt came weeks later, when features that had been free moved behind the paywall. Decide once what is free forever, write it down publicly, and never move anything back across that line.

On ads: Tinybeans spent four years escaping an advertising business, cutting ad revenue from about $9M to about $1.3M. FamilyAlbum grew ad-free for a decade, then added banner ads to its free tier in late 2025. BeReal was bought and monetised. The drift runs one way, and the only defence is a cost base so small you never need the money.

### Donations do not work. Hosting does

This evidence is unusually clean. Plausible receives about $300 a month in total from all of its self-hosted users. Mastodon says under 1% of its users donate regularly, and in one recent year roughly $1.6M of its $2.4M came from a single donor. Its founder was paid €5,000 a month, less than his own infrastructure engineer, and stepped down in late 2025 citing his health. core-js, which runs on a large share of the web, earned its maintainer under two dollars an hour.

Meanwhile the open products that sell hosting are genuinely healthy. Ghost runs at about $11M a year from around 30,000 paying customers, under a foundation that cannot be sold. Plausible went from $64 of monthly revenue in 2019 to a million a year with ten people. Home Assistant's $65-a-year cloud subscription funds more than fifty full-time staff. Ente is bootstrapped with twelve. Immich took a patron and sells an optional key with a promise that no feature will ever sit behind it.

### An empty room teaches people the product is dead

Peach reached the top ten and lost its retention inside a week. Google's own post-mortem of Google+ gave the number that says everything: 90% of sessions lasted under five seconds. Once a room reads as empty, people stop checking, and you cannot un-teach that.

Nextdoor is the only product here with a hard mechanic against it. A new neighbourhood **expires if it doesn't reach ten verified members within 21 days**. It refuses to let a dead room exist. That's a deletion mechanic rather than an invitation mechanic, and it's the most directly stealable idea in this investigation.

### Numbers in this field that turn out to be folklore

Worth knowing, because they circulate as fact: that 1 to 4% of self-hosters convert to paid, Discord's often-quoted 2.9% Nitro conversion, Path's $40M to $50M sale price, and any revenue figure attributed to Home Assistant's cloud business. None has a traceable source, and Nabu Casa has never published a subscriber count at all.

## What this means for Nah?

### Four things to add, none of which break a rule

1. **A widget.** The highest-value, lowest-conflict move available. The latest moment from your home, on the home screen. Returning costs nothing, it cannot fatigue, and it still shows something on a week when nobody posted.

2. **A fixed digest time.** Seven in the evening, every evening, not a clever moment chosen by a model. BeReal's randomness was its fear-of-missing-out engine. A predictable time is an anchor, which is the one habit mechanism with trial evidence behind it.

3. **The archive, from day one.** "A year ago today" costs almost nothing to build, breaks no rule, and is what people in this category actually pay for. BeReal's own users named the memory archive as the thing they valued.

4. **Named reciprocity in the digest.** "Anna posted" is not a count and not a streak. Obligation to a person is the humane version of the same force.

5. **An expiry for empty rooms.** Copy Nextdoor, which deletes a new neighbourhood that hasn't reached ten members in three weeks. A home that never gets going should wind itself up rather than sit there teaching everyone the product is dead.

### One rule to reconsider

**No reactions in the first version is the expensive rule.** Posting into silence is the documented death spiral for a small network: nobody responds, the poster stops, and the cascade starts. A reaction is not a count. Ship an acknowledgement that the poster can see, with no number attached anywhere, and the no-counts headline survives intact.

### What to keep without hesitation

No counts is cheap, structural and still the sentence. No algorithm costs nothing at 150 people, because there is no feed to rank. One digest a day is exactly at the ceiling of what people tolerate, which is the right place to sit.

### The reframe

"A maybe daily check-in" is not a daily product. Retro, the best-performing peer, is weekly and gets about 45% of people engaging daily. Yope beats that with streaks and widgets. BeReal got true daily behaviour and lost 70% of it in eight months. **Plan Nah? as a weekly product with a daily surface.** That isn't a downgrade, it's what the healthiest app in the category already does.

### The contradiction worth resolving

The research pulled in two directions here, and the tension is the most useful thing in it.

One side says a small network needs several posters. The model built from Friendster's collapse shows members leaving once their own active friends have gone, cascading from the edges inward. A room with one poster gives everybody else exactly one reason to come back, and takes it away all at once.

The other side says the opposite. Every family product that survived a decade is **built so that one uploader is enough**. Tinybeans, FamilyAlbum and Retro all work when a single person posts and everyone else receives. The products that required reciprocal posting, Path and Cocoon and Peach and Cohost, are all dead.

They are answering different questions. The first describes what happens to a network of peers. The second describes what the survivors built instead. The resolution for Nah? is to stop treating a reader as a failed contributor: **design so that one poster is enough, and treat more as upside**. The grandmother who reads everything and posts nothing isn't a retention problem. She's the product working, and she's the reason the poster keeps posting.

That reframing changes what the first version should contain. The receiving experience, the digest and the widget and the archive, stops being secondary to the composer.

### What to watch in the six-week test

Signs it's working:

- **At least two people post without being prompted in weeks three to six.** Posting in the first fortnight is politeness, not signal.

- **Someone opens it on a day you didn't post.** - **Someone asks to add a person you didn't invite.** - **Someone complains about a missing feature** rather than praising the idea.

- **Someone re-shares something out of Nah? into a group chat.** The strongest signal available, because it means Nah? has become the source of record.

Signs it isn't:

- **Every post has your name on it.** - **The most engaged behaviour is reacting rather than posting.** - **People reply in WhatsApp about things they saw in Nah?** - **Week-four opens are below week-one opens.** That's the curve Peach rode to death in four days.

### The odds, stated plainly

Of roughly fourteen private-social attempts documented in this research, three are alive after five years, and all three are funded teams rather than individuals. None is open source. The researcher's own estimate, marked clearly as an estimate rather than a statistic, puts a solo-developer open-source private social app reaching sustained use by even a handful of groups at two years somewhere around 10 to 20%, and reaching enough revenue to justify the time under 5%.

The reason isn't demand. Demand is visible on every page of this report. It's that maintenance cost stays roughly constant however few users you have, while revenue scales with a base that is capped on purpose.

The most likely causes of death, in order: your own energy somewhere around month nine to eighteen; the room going quiet without anyone telling you, because they like you; the app-store release treadmill eating the time meant for the product; nobody paying, not from refusal but from the absence of a moment where paying is the obvious next step; and media sync bugs consuming the maintenance budget.

## What we couldn't verify

- No retention figures are published for Retro, Yope, Locket, Marco Polo, FamilyAlbum, Lapse, Dispo or Noplace. Every retention claim in this report about them is either a company statement or absent.

- BeReal's own daily-user claims after 2023 contradict every outside tracker, and no credible 2026 figure exists. Numbers circulating this year trace back to a corrupted restatement of an April 2022 figure.

- Retro's engagement figures come from its founders around a funding round.

- Marco Polo publishes nothing at all. Its revenue figures are third-party estimates.

- A large share of app-statistics websites in 2026 are machine-generated and mutually contradictory. Those were excluded where identified, which is also why some obvious-looking numbers are missing here.

## Sources

### Scale, money and the BeReal collapse

- [Sensor Tower on BeReal's churn](https://sensortower.com/blog/bereal-data-snapshot) and [TechCrunch on its daily-open rate](https://techcrunch.com/2022/10/10/bereal-tops-53m-installs-but-only-9-open-the-app-daily-estimates-claim/)

- [Glossy, February 2023](https://www.glossy.co/beauty/is-bereal-in-its-flop-era/) and [BeReal disputing the trackers](https://techcrunch.com/2023/09/29/bereal-pushes-back-at-report-that-its-losing-steam-says-it-now-has-25m-daily-users/)

- [The rise and fall of BeReal, New Media and Society, 2025](https://researchportal.northumbria.ac.uk/en/publications/the-rise-and-fall-of-bereal-values-of-and-motivations-for-disenga/)

- [Voodoo's 2025 results](https://voodoo.io/news/2025-a-year-of-growth-and-transformation-for-voodoo)

- [Retro's Series A](https://techcrunch.com/2026/08/28/friend-focused-photo-sharing-app-retro-snags-21m/) and [its engagement claims](https://techcrunch.com/2025/12/12/retro-a-photo-sharing-app-for-friends-lets-you-time-travel-through-your-camera-roll/)

- [Yope's raise](https://techcrunch.com/2026/07/22/yope-raises-12-3m-to-build-a-private-social-network-without-algorithms-or-ads/), [its early retention](https://techcrunch.com/2025/02/24/yope-is-sparking-genz-and-vc-interest-with-an-instagram-like-app-for-private-groups) and [its 2025 user numbers](https://www.businessinsider.com/yope-app-popular-gen-z-2025-3)

- [Airbuds on its daily and monthly users](https://www.businessinsider.com/airbuds-widget-app-alexis-ohanian-social-music-teens-gen-z-2025-11)

- [Locket adds celebrity accounts](https://techcrunch.com/2025/08/06/photo-sharing-app-locket-is-banking-on-a-new-celebrity-focused-feature-to-fuel-its-growth)

- [Tinybeans FY25 results](https://ecommercenews.com.au/story/tinybeans-reports-growth-in-subscriptions-lower-cash-outflows)

- [Snap's Q2 2026 numbers](https://www.cnbc.com/2026/08/03/snap-q2-earnings-report-2026.html) and [app retention benchmarks](https://www.businessofapps.com/data/app-retention-rates/)

### Habits, notifications and group dynamics

- [How long habits take to form](https://pmc.ncbi.nlm.nih.gov/articles/PMC3505409/) and [the anchoring trial](https://mhealth.jmir.org/2021/12/e32794)

- [Duolingo on streaks](https://www.lennysnewsletter.com/p/behind-the-product-duolingo-streaks) and [the Busuu comparison](https://sensortower.com/blog/duolingo-streak-feature-app-engagement-growth)

- [Research on Snapchat streaks](https://www.sciencedirect.com/science/article/pii/S2772503023000476)

- [The active versus passive use meta-analysis](https://academic.oup.com/jcmc/article/29/1/zmad055/7595758)

- [Push notification tolerance](https://www.mobiloud.com/blog/push-notification-statistics)

- [Participation inequality](https://www.nngroup.com/articles/participation-inequality/), [the Friendster collapse model](https://arxiv.org/abs/1302.6109) and [the iWiW study](https://www.nature.com/articles/s41598-017-17135-1)

- [Subscription retention benchmarks](https://www.revenuecat.com/state-of-subscription-apps)

- [Timehop's nostalgia engine](https://medium.com/@sm_app_intel/did-facebooks-on-this-day-have-an-impact-on-timehop-6c9439fb441c)

### Shutdowns

- [Cocoon's goodbye](http://web.archive.org/web/20230529110637/https://ourcocoon.com/news/cocoon-is-saying-goodbye) and [its pricing page](http://web.archive.org/web/20230616154351/https://ourcocoon.com/membership)

- [Cohost's shutdown post and financials](https://cohost.org/staff/post/7611443-cohost-to-shut-down)

- [FUTO Circles, archived](https://github.com/circles-project/circles-android) and [Home Assistant Cloud](https://www.home-assistant.io/cloud/)

*Compiled on 12 September 2026 from three research passes, alongside the concept audit of 11 September. Figures are labelled as company claims or outside estimates where the distinction matters.*
