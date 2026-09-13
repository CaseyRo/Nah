---
title: "Concept audit: Nah? ten months in"
permalink: /research/concept-audit/
layout: page
date: 2026-09-11
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

*Where the idea stands, who else is in the room, whether one person can build it, and what the server looks like if any family should be able to run its own.*

## The short version

- **The idea is sharp, but the product doesn't exist yet.** Ten months after the first prompt there are eight decision records, a design system and a blog, but no app and no server. The last commit was on 19 May.

- **The one-click server wish changes the product, not just the hosting.** Path gave every person their own circle on one shared server. "Every family runs its own server" means every group owns a home. That is a different social model, and it is the better one for Nah?.

- **If homes win, Mastodon goes.** Mastodon is built for public, federated, one-way following, and it needs five processes and 2–4 GB of memory per instance. A home needs one small binary and one database file.

- **"One click" for a normal family means a hosted home they pay a little for.** Self-hosting stays open as the way out. This is the Home Assistant and Nabu Casa pattern, and it answers the funding question the vision never settled.

- **The corner Nah? wants is empty, for a reason.** Retro and Yope raised $33M in 2026 on "no algorithm, no ads", but both are closed. The one open, self-hostable family app, FUTO Circles, was archived in February 2025. The survivors in that corner all have a paid hosted plan, which is what hosted homes would be.

- **It is achievable if the MVP stays the size ADR-0008 set.** Build one home for your own circle first. Hosting other families is a second project, and it should start only after six weeks of real use.

## Where Nah? stands

Nah? began as a prompt on 19 November 2025. Here is everything that has happened since, oldest first, the way Nah?'s own feed would show it.

- **19 Nov 2025** — **The prompt.** Path again, web-first, on Mastodon's backend. "Sign up should be easy and effortless, creating private groups too."

- **28 Jan 2026** — **The vision.** Proposal, technical design and eight capability specs: a Mastodon fork plus a SvelteKit PWA. The build-in-public blog goes live.

- **3 Mar 2026** — **Tightening.** Onboarding, invites, error states, a notification table and a full design system are written into the specs.

- **25–27 Apr 2026** — **The decision session.** No counts anywhere. Invitation is connection. Text, voice and photo. Ritual onboarding. A daily digest. An MVP for five to ten friends. Flutter instead of a PWA. A single MVP spec is drafted in chat and never committed.

- **19 May 2026** — **The Flutter pivot lands.** README, vision and the first capability spec move to Flutter. Eight decision records, PRODUCT.md and DESIGN.md are committed. This is the last commit.

- **19 Jul 2026** — **The stack ticket closes.** The Linear issue "Native vs Flutter vs PWA" is cancelled. No Nah? work is tracked in Linear or Things after this.

- **11 Sep 2026** — **Today.** No app folder, no server folder, no MVP spec.

### What's genuinely strong

- **A point of view in one sentence.** "No counts anywhere" is structural, visible and nearly free to build. Most products never find their sentence.

- **Decisions that can be proven wrong.** Every ADR names what would falsify it. That is rare discipline, and it made this audit easy.

- **A real design system.** Tokens, type, motion rules and anti-references are written down, down to oklch values.

### What's broken

- **The specs still describe January.** The decisions moved in April. The MVP spec meant to reconcile them lives only in a chat transcript.

- **Anyone reading the repo gets the old product.** That includes a future you and any AI assistant you hand it to.

- **The public Flutter post tells an old story.** It describes two AIs disagreeing. Your April notes already corrected that to one session pivoting mid-flow.

### Where the specs and the decisions disagree

| Topic | The specs say | The decisions say |
|---|---|---|
| Friend limit | A "X/150 friends" counter and a constellation ring that fills up | No counts or meters. The cap is a sentence: "your circle is full" (ADR-0004) |
| Notifications | Push for requests and comments, a count badge on a bell | One daily digest, no badges, opt in to more (ADR-0007) |
| Onboarding | Email and password, profile photo, a three-card explainer, "Welcome home" | Four beats: a slow sentence, one real question, a haptic, the quiet space (ADR-0005) |
| Reactions | A heart button on every card, double-tap for Love | No reactions in the MVP (ADR-0008) |
| Connecting | Send a request, the other person accepts or declines | Invitation is connection, in one motion (ADR-0003) |

### One factual error to fix

PRODUCT.md, ADR-0001 and ADR-0007 all say EU iPhone users lose push notifications under the Digital Markets Act. Apple announced that in February 2024 for home-screen web apps, then reversed it on 1 March 2024. Native apps were never affected, so the Flutter app gets push in the EU like everywhere else. ADR-0001 still stands on Bluetooth and widgets. The "EU users get a quieter app" caveat can simply go.

### And one pattern to name

Every return to Nah? has produced documents and a re-decision: SvelteKit in January, a PWA matrix in April, Flutter a few messages later, and now the backend is open again. This audit could easily become the next round of that. So it ends in a build order, not in more options.

## What we're actually trying to do

The intent hasn't wavered across ten months of documents.

> A private home for your closest people, to share life without performing for the internet.

*PRODUCT.md*

The job is narrow: share a real moment with the few people who should see it, and see theirs. Success is years of quiet use, not engagement. Your notes add a layer the specs never picked up. The name means Not Alone Here. The idea seeds in SiYuan are about support: people carrying things together that used to need a big network. And the very first prompt asked for private groups, not only a personal circle.

Put together, Nah? is less "Path again" and more a home for the handful of people who carry each other. That matters for the choice below.

### The fork the one-click wish exposes

You asked for something new: any family or clique should be able to run its own server in one click. That sounds like a deployment detail. It isn't, because it decides who a server belongs to, and that decides who sees what.

<figure>

<div class="fig">

<svg viewBox="0 0 880 372" role="img" aria-label="Left: one Nah? server where each person has their own overlapping circle. Right: three small home servers with fixed members, and your app merging the two homes you belong to into one feed.">

<defs>

<marker id="arrF" viewBox="0 0 10 10" refX="9" refY="5" markerWidth="8" markerHeight="8" orient="auto-start-reverse">
<path d="M0,0 L10,5 L0,10 z" class="nd-arrowhead"/>
</path>
</marker>

</defs>

<text x="10" y="22" class="nd-h">
Nah? as a network
</text>

<text x="10" y="44" class="nd-s">
One server. A circle per person.
</text>

<rect x="10" y="60" width="400" height="272" rx="14" class="nd-box"/>
</rect>

<ellipse cx="165" cy="200" rx="115" ry="55" class="nd-ring you"/>
</ellipse>

<ellipse cx="255" cy="202" rx="130" ry="85" class="nd-ring"/>
</ellipse>

<line x1="160" y1="200" x2="80" y2="200" class="nd-line"/>
</line>

<line x1="160" y1="200" x2="250" y2="200" class="nd-line"/>
</line>

<line x1="250" y1="200" x2="340" y2="160" class="nd-line"/>
</line>

<line x1="250" y1="200" x2="340" y2="245" class="nd-line"/>
</line>

<circle cx="80" cy="200" r="17" class="nd-node"/>
</circle>
<text x="80" y="204" class="nd-in">
Sam
</text>

<circle cx="160" cy="200" r="17" class="nd-node you"/>
</circle>
<text x="160" y="204" class="nd-in you">
You
</text>

<circle cx="250" cy="200" r="17" class="nd-node"/>
</circle>
<text x="250" y="204" class="nd-in">
Wife
</text>

<circle cx="340" cy="160" r="17" class="nd-node"/>
</circle>
<text x="340" y="164" class="nd-in">
Maya
</text>

<circle cx="340" cy="245" r="17" class="nd-node"/>
</circle>
<text x="340" y="249" class="nd-in">
Jo
</text>

<text x="56" y="284" class="nd-acc">
your circle
</text>

<text x="300" y="108" class="nd-s" text-anchor="middle">
Wife's circle
</text>

<text x="26" y="318" class="nd-s">
One Nah? server, run by the project
</text>

<line x1="440" y1="8" x2="440" y2="364" class="nd-faint"/>
</line>

<text x="470" y="22" class="nd-h">
Nah? as homes
</text>

<text x="470" y="44" class="nd-s">
A small server per group. Your app merges yours.
</text>

<rect x="470" y="60" width="128" height="164" rx="14" class="nd-box"/>
</rect>

<text x="482" y="82" class="nd-t">
Family
</text>

<path d="M506,116 L562,116 L534,160 Z" class="nd-faint"/>
</path>

<circle cx="506" cy="116" r="15" class="nd-node you"/>
</circle>
<text x="506" y="120" class="nd-in you">
You
</text>

<circle cx="562" cy="116" r="15" class="nd-node"/>
</circle>
<text x="562" y="120" class="nd-in">
Wife
</text>

<circle cx="534" cy="160" r="15" class="nd-node"/>
</circle>
<text x="534" y="164" class="nd-in">
Mum
</text>

<text x="482" y="208" class="nd-s">
on Nah? Cloud
</text>

<rect x="606" y="60" width="128" height="164" rx="14" class="nd-box"/>
</rect>

<text x="618" y="82" class="nd-t">
Uni friends
</text>

<path d="M642,116 L698,116 L670,160 Z" class="nd-faint"/>
</path>

<circle cx="642" cy="116" r="15" class="nd-node you"/>
</circle>
<text x="642" y="120" class="nd-in you">
You
</text>

<circle cx="698" cy="116" r="15" class="nd-node"/>
</circle>
<text x="698" y="120" class="nd-in">
Sam
</text>

<circle cx="670" cy="160" r="15" class="nd-node"/>
</circle>
<text x="670" y="164" class="nd-in">
Ari
</text>

<text x="618" y="208" class="nd-s">
on Sam's NAS
</text>

<rect x="742" y="60" width="128" height="164" rx="14" class="nd-soft"/>
</rect>

<text x="754" y="82" class="nd-t">
Wife's friends
</text>

<path d="M778,116 L834,116 L806,160 Z" class="nd-faint"/>
</path>

<circle cx="778" cy="116" r="15" class="nd-node"/>
</circle>
<text x="778" y="120" class="nd-in">
Wife
</text>

<circle cx="834" cy="116" r="15" class="nd-node"/>
</circle>
<text x="834" y="120" class="nd-in">
Maya
</text>

<circle cx="806" cy="160" r="15" class="nd-node"/>
</circle>
<text x="806" y="164" class="nd-in">
Jo
</text>

<text x="754" y="208" class="nd-s">
on Nah? Cloud
</text>

<text x="806" y="246" class="nd-s" text-anchor="middle">
you're not in this one
</text>

<line x1="534" y1="224" x2="598" y2="280" class="nd-line" marker-end="url(#arrF)"/>
</line>

<line x1="670" y1="224" x2="660" y2="280" class="nd-line" marker-end="url(#arrF)"/>
</line>

<text x="552" y="258" class="nd-s" text-anchor="end">
member
</text>

<text x="672" y="258" class="nd-s">
member
</text>

<rect x="548" y="282" width="170" height="64" rx="14" class="nd-key"/>
</rect>

<text x="633" y="308" class="nd-t" text-anchor="middle">
Your app
</text>

<text x="633" y="328" class="nd-s" text-anchor="middle">
one feed from two homes
</text>

</svg>

</div>

<figcaption>
On the left, Maya sees Wife's moments but not yours, and the server has to work that out for every moment. On the right, a home is the audience. Everyone in it sees everything in it, and your app stitches your homes into one feed.
</figcaption>

</figure>

|   | Nah? as a network | Nah? as homes |
|---|---|---|
| Who owns a server | The project | The group, or Nah? Cloud on its behalf |
| Who sees a moment | Your personal circle, which overlaps other people's | Everyone in the home you posted to |
| The 150 limit | Per person, enforced across one big graph | Per home, a simple member cap |
| Self-hosting | Pointless without federation, because circles cross servers | Natural: one server per group |
| Closest to | Path | The family group chat, done properly |
| Funding | Donations | Hosting fees, plus donations |

### Why homes recommended

- **Privacy by architecture becomes literal.** Your family's moments sit on your family's server. "Who can see this?" has a one-word answer: the home.

- **It's simpler to build.** There's no personal graph, no cross-circle visibility rule and no inner-circle feature. For a smaller audience, start a smaller home.

- **It matches how people already organise.** Families and cliques already live in named groups. One motivated person brings the whole group, which softens the cold start. Yope reports that a fifth of its active users invited older relatives.

- **It can pay for itself.** Hosting a home is a service people understand paying for.

### What homes cost

- **It's less like Path.** Path's magic was your own circle. The merged feed is your personal view, but every post now goes to a home. People with one home never see that choice.

- **Two decision records need replacing.** ADR-0002 (Mastodon) and ADR-0003 (invitation is connection) change. Invitation becomes "join this home". The other six survive intact.

- **One account per home.** The app hides this by reusing your name and photo everywhere.

## The competition

The market wants private sharing, and in 2026 money is flowing to it. But the winners are closed, venture-funded apps with a hook and a money lever. Nobody owns the combination Nah? is reaching for: a private space the group owns, open source and self-hostable. The one serious attempt was archived in February 2025.

<figure>

<div class="fig">

<svg viewBox="0 0 880 420" role="img" aria-label="A two-by-two map. Closed, company-run apps for friends and family hold the money and users. Public apps sit below. The group-owned, private corner has photo libraries and intranets but no warm social home; FUTO Circles tried it and was archived in 2025. Nah? homes aims at that corner.">

<text x="100" y="104" class="nd-t" text-anchor="end">
For your
</text>

<text x="100" y="122" class="nd-t" text-anchor="end">
people
</text>

<text x="100" y="288" class="nd-t" text-anchor="end">
For an
</text>

<text x="100" y="306" class="nd-t" text-anchor="end">
audience
</text>

<text x="302" y="408" class="nd-t" text-anchor="middle">
Run by a company
</text>

<text x="677" y="408" class="nd-t" text-anchor="middle">
Run by the group
</text>

<rect x="120" y="20" width="365" height="180" rx="12" class="nd-box"/>
</rect>

<text x="136" y="48" class="nd-t">
Closed, friends-only apps
</text>

<text x="136" y="76" class="nd-s">
Retro · Yope · Locket · Friendlinq
</text>

<text x="136" y="98" class="nd-s">
FamilyAlbum · Tinybeans · Marco Polo
</text>

<text x="136" y="120" class="nd-s">
WhatsApp groups · iCloud Shared Albums
</text>

<text x="136" y="176" class="nd-lab">
Where the money and users are
</text>

<rect x="495" y="20" width="365" height="180" rx="12" class="nd-key"/>
</rect>

<text x="511" y="48" class="nd-t">
Owned by the group, for your people
</text>

<text x="511" y="76" class="nd-s">
Immich, Ente: photo libraries, not homes
</text>

<text x="511" y="98" class="nd-s">
Nextcloud, HumHub: they feel like an intranet
</text>

<text x="511" y="120" class="nd-s">
FUTO Circles: this exact idea, archived 2025
</text>

<rect x="511" y="150" width="118" height="30" rx="15" class="nd-box" style="stroke:var(--accent);stroke-width:2"/>
</rect>

<text x="570" y="170" class="nd-acc" text-anchor="middle">
Nah? homes
</text>

<text x="641" y="170" class="nd-s">
the empty seat
</text>

<rect x="120" y="210" width="365" height="170" rx="12" class="nd-soft"/>
</rect>

<text x="136" y="238" class="nd-t">
Public, company-run
</text>

<text x="136" y="266" class="nd-s">
Instagram · BeReal · Noplace
</text>

<text x="136" y="288" class="nd-s">
Built on reach, counts and ads
</text>

<rect x="495" y="210" width="365" height="170" rx="12" class="nd-soft"/>
</rect>

<text x="511" y="238" class="nd-t">
Public, on small servers
</text>

<text x="511" y="266" class="nd-s">
Mastodon · Pixelfed · GoToSocial
</text>

<text x="511" y="288" class="nd-s">
Small servers, but public by default
</text>

</svg>

</div>

<figcaption>
Money and users sit top left. The top-right corner, a private space the group owns, has tools but no home. The one app built for exactly that corner was archived when its backer judged it too costly to commercialise.
</figcaption>

</figure>

### The closest rivals

| Product | Where it stands in 2026 | What it teaches Nah? |
|---|---|---|
| Retro | $21M Series A announced in August. About 1M users, and the top app in Germany in December 2025. Premium is about $5 a month | The closest live product. A weekly friends-only recap proves "no algorithm, no ads" can win with a subscription |
| Yope | $12.3M in July. About 15M registered, and a fifth of active users invited older relatives | Private micro-groups that pull in parents and grandparents. Its turn toward AI features is the part to reject |
| Locket | 80M downloads, about 9M daily users, profitable | Presence on the home screen through a widget. It sells its way past its own friend cap |
| Friendlinq | Launched 2 September 2026, no numbers yet | Nearly Nah?'s pitch word for word: chronological, invite-only, no ads, "free forever". No stated way to pay for itself |
| FamilyAlbum | 30M users, used by about 65% of Japanese parents | Free unlimited storage. Nah? can't match that and shouldn't try |
| WhatsApp groups | Pre-installed, free and end-to-end encrypted. Ads stay out of chats | The real competitor. Nah? has to beat the family group chat, not Instagram |
| FUTO Circles | Archived February 2025 | Encrypted family social, open and self-hostable: exactly Nah?'s corner. It had a backer and still stopped |
| Immich and Ente | Very active, and both ship Flutter apps | Proof that Flutter plus a self-hostable server works for families, and that a paid hosted plan can fund the open work |

### Signs the market wants this

- **Money is moving.** Retro and Yope raised $33M between them in 2026 with "no algorithm, no ads" in the headline.

- **The incumbents admit it.** Meta's own trial evidence puts friends' content at about 7% of Instagram time. Instagram's head says personal sharing moved into DMs years ago.

- **The scale is real.** FamilyAlbum has 30M users and Locket about 9M a day.

### Signs it doesn't

- **Group chats already absorb the need.** They're free, installed and encrypted.

- **The winners are narrow hooks.** A widget, a weekly recap, a camera, music. Nobody is winning as a general private network.

- **Path-shaped apps keep dying.** Path in 2018, Cocoon in 2023, Circles in 2025. Cohost, a community-funded network, closed in 2024 with 2,630 paying members and a deficit.

### Why Path actually died

Three things killed it. Facebook's network effect came first, and Path's own cap made it worse: "most of my friends weren't on Path". Then came a 2012 address-book upload scandal that ended in an FTC settlement over children's data. Finally it was sold to Kakao in 2015, mostly for its Indonesian users. The cap it is remembered for went from 50 to 150 to 500 as growth stalled. Homes answer the first problem better than a personal circle does, because one person brings the whole group.

### What everyone does that Nah? refuses

Every commercial player owns the server and runs two things Nah? rejects. One is a hook to bring you back: a daily push, a widget, recaps, AI. The other is a money lever: ads, or paywalls on storage, history or friend count. Refusing both is Nah?'s only durable difference from well-funded rivals, and group ownership makes the refusal believable rather than a promise.

It is also why the principled projects keep dying. In the open-source corner, every survivor has a backer or a paid hosted product. Nah? needs one of those from the start, and that's the hosted home.

What families lack isn't another feed. It's everything group chats get wrong: photos get buried, there's no archive, notifications are noisy, and nothing works for grandparents. A chronological home with a quiet digest answers the first three. Tinybeans' email digest for grandparents who never install an app is worth considering later.

## Is it achievable?

It's three questions, and they have different answers.

| Question | Answer | Why |
|---|---|---|
| Nah? for your own circle | Yes | The MVP is small: accounts, invites, three moment types, a chronological feed and a digest. On a simple backend it's about three months of focused work, or around six months of evenings next to client work. TestFlight covers five to ten friends with only a light beta review. |
| Any family, one click | In stages | It's a second product: provisioning, payment, backups, support and terms. Start it only once your circle has used Nah? for six weeks and still wants it. |
| A lasting project | Unproven | Money is flowing to private sharing in 2026, but to closed apps with hooks, while open projects keep hitting funding cliffs. Families already have free defaults in WhatsApp and iCloud Shared Albums, so Nah? has to be clearly nicer, not just more principled. |

### Rough effort for the MVP

The server figures come from the research pass. The app figure is my own estimate. All are focused weeks, not calendar weeks.

| Piece | Weeks |
|---|---|
| Home server: homes, invites, three moment types, feed, digest job, export | 3–4 |
| Push relay for Apple and Google | about 1 |
| Invite domain, home directory, demo home | about 1 |
| Flutter app: ritual onboarding, feed, three composers, invites, offline queue | 6–8 |
| **Total** | **11–14** |

Hosting other families adds about two weeks for hosted homes and one to two for app-store packaging. A Nah? relay for home boxes adds two to three more.

### The risks, biggest first

- **Re-deciding instead of building.** The history above is the evidence.

- **Solo capacity.** Nah? competes with paid client work for the same evenings. Voice is the costliest composer, so build it last inside the MVP.

- **The quiet-app trap.** No counts, a daily digest and no reactions make a calm app, and calm apps can die quietly. ADR-0004 and ADR-0007 already name this, and the six-week test is where you'll find out.

- **Store rules.** Apple's guideline 1.2 requires filtering, reporting, blocking and a published contact for any app with user content, private groups included. Apple holds the app's publisher responsible even when the content sits on someone else's server. The MVP list has none of this. "Remove from home", "report to the home's admin", a project abuse contact and the directory's cut-off switch would cover it. In-app account deletion is required too.

- **Children.** Family homes will hold photos of children, and some members will be teenagers. Path's FTC settlement was about children's data. Set a minimum member age before the first family joins. In Germany, under-16s need parental consent under GDPR.

- **Hosting other people's family photos.** Nah? Cloud would make CDIT a hosting provider for EU users. Closed homes likely fall outside the Digital Services Act's "online platform" duties, but notice-and-action, a privacy policy and GDPR terms still apply. Run a compliance pass before the first paying family.

## Server architecture, in depth

Start from what a home server has to do, because it's far less than Mastodon does.

### The whole job

- Create accounts and sign people in.

- Turn an invite link into membership of a home.

- Store moments: a line of text, one photo, or up to 60 seconds of voice.

- Serve the feed.

- Wake phones once a day when the home had new moments.

- Export everything.

```sql
`-- The entire feed algorithm for a home
SELECT * FROM moments
WHERE home_id = ?
ORDER BY created_at DESC
LIMIT 30;`
```

A home has at most 150 members, so there's no fan-out, no timeline cache and no ranking. There's also no search, no counts and no discovery, so the server never has to understand what a moment says. That pays off later, in the encryption section.

### Backends compared

Five realistic options, judged on one question: could a family run one without ever calling you?

| Backend | What runs | Memory | Fits Nah?'s model | One click per family |
|---|---|---|---|---|
| Mastodon 4.6 | Web, streaming, Sidekiq, PostgreSQL and Redis: five containers, plus TLS and email | 2–4 GB | No. One-way follows, counts in every API response, and a permanent fork to maintain | No. Managed hosts exist because it needs an operator |
| GoToSocial 0.22 | One Go binary with SQLite | 250–350 MB | Partly. Still follow-based, still beta, and its allowlist mode is "experimental" | Nearly, but the product fights the model |
| PocketBase 0.39, as a Go framework | One Go binary with SQLite, files on disk or S3 | under 100 MB | Yes. You write a small social layer, and there's an official Dart SDK | Yes |
| Serverpod 3 | A Dart server plus PostgreSQL | Postgres-sized | Yes, and it shares Dart models with the app | No. PostgreSQL is required |
| Supabase, self-hosted | Ten or more containers | 1.5–4 GB idle | Yes | No |

### The recommendation

**Build "Nah? Home" as one Go binary on PocketBase, used as a framework rather than as an app.** PocketBase brings sign-in, file storage, live updates, scheduled jobs, backups to S3 and an admin screen in one file, with an official Dart SDK for the Flutter side. You add the social layer, which for this MVP is small: homes, invites, moments, the digest job and export. Voice and photos are encoded on the phone, so the server never runs FFmpeg.

Two honest caveats. PocketBase is still 0.x, with one maintainer and breaking changes between releases, so pin the version and vendor it. If it ever stalls, what's underneath is plain Go and SQLite, which you can keep running without it. GoToSocial is the runner-up if keeping the Mastodon API matters. It saves writing a server, but every Nah?-specific rule would fight a federation server, which is how Hometown ended up a year behind.

One lesson to take from Mastodon: an instance's domain can never change, so its identity is welded to a hostname. Give every home an ID backed by its own key from day one, so it can prove it's the same home after it moves.

### How the pieces fit

<figure>

<div class="fig">

<svg viewBox="0 0 900 470" role="img" aria-label="Moments travel directly between the Nah? app and each home server. Homes send sealed wake-ups to a push relay run by the Nah? project, which forwards them through Apple and Google push to the phone. Invite links on the Nah? domain open the app, and a small directory tells it where the home lives.">

<defs>

<marker id="arrA" viewBox="0 0 10 10" refX="9" refY="5" markerWidth="8" markerHeight="8" orient="auto-start-reverse">
<path d="M0,0 L10,5 L0,10 z" class="nd-arrowhead"/>
</path>
</marker>

</defs>

<rect x="300" y="44" width="280" height="340" rx="16" class="nd-region"/>
</rect>

<text x="316" y="68" class="nd-lab">
The group's homes
</text>

<rect x="650" y="44" width="220" height="340" rx="16" class="nd-region"/>
</rect>

<text x="666" y="68" class="nd-lab">
Run by the Nah? project
</text>

<rect x="20" y="60" width="200" height="64" rx="12" class="nd-soft"/>
</rect>

<text x="120" y="88" class="nd-t" text-anchor="middle">
Apple / Google push
</text>

<text x="120" y="108" class="nd-s" text-anchor="middle">
APNs and FCM
</text>

<rect x="20" y="170" width="200" height="110" rx="14" class="nd-box"/>
</rect>

<text x="120" y="204" class="nd-t" text-anchor="middle">
Nah? app
</text>

<text x="120" y="226" class="nd-s" text-anchor="middle">
one session per home
</text>

<text x="120" y="246" class="nd-s" text-anchor="middle">
offline cache and queue
</text>

<rect x="315" y="84" width="250" height="96" rx="12" class="nd-box"/>
</rect>

<text x="440" y="114" class="nd-t" text-anchor="middle">
Home A, on Nah? Cloud
</text>

<text x="440" y="136" class="nd-s" text-anchor="middle">
one binary, one SQLite file
</text>

<text x="440" y="156" class="nd-s" text-anchor="middle">
media on disk or S3
</text>

<rect x="315" y="264" width="250" height="96" rx="12" class="nd-box"/>
</rect>

<text x="440" y="294" class="nd-t" text-anchor="middle">
Home B, on a cousin's NAS
</text>

<text x="440" y="316" class="nd-s" text-anchor="middle">
the same binary
</text>

<text x="440" y="336" class="nd-s" text-anchor="middle">
reached through a tunnel
</text>

<rect x="665" y="84" width="190" height="80" rx="12" class="nd-key"/>
</rect>

<text x="760" y="116" class="nd-t" text-anchor="middle">
Push relay
</text>

<text x="760" y="138" class="nd-s" text-anchor="middle">
forwards sealed wake-ups
</text>

<rect x="665" y="264" width="190" height="80" rx="12" class="nd-key"/>
</rect>

<text x="760" y="296" class="nd-t" text-anchor="middle">
nah.app/i
</text>

<text x="760" y="318" class="nd-s" text-anchor="middle">
invites and home directory
</text>

<line x1="222" y1="206" x2="313" y2="146" class="nd-line" marker-start="url(#arrA)" marker-end="url(#arrA)"/>
</line>

<line x1="222" y1="248" x2="313" y2="300" class="nd-line" marker-start="url(#arrA)" marker-end="url(#arrA)"/>
</line>

<text x="262" y="222" class="nd-s" text-anchor="middle">
moments
</text>

<text x="262" y="238" class="nd-s" text-anchor="middle">
& media
</text>

<line x1="565" y1="124" x2="663" y2="124" class="nd-line" marker-end="url(#arrA)"/>
</line>

<text x="615" y="116" class="nd-s" text-anchor="middle">
wake-up
</text>

<polyline points="565,304 615,304 615,146 663,146" class="nd-line" marker-end="url(#arrA)"/>
</polyline>

<polyline points="855,112 884,112 884,26 120,26 120,58" class="nd-line" marker-end="url(#arrA)"/>
</polyline>

<text x="440" y="19" class="nd-s" text-anchor="middle">
push, signed with the publisher's key
</text>

<line x1="120" y1="124" x2="120" y2="168" class="nd-line" marker-end="url(#arrA)"/>
</line>

<text x="130" y="151" class="nd-s">
notification
</text>

<polyline points="760,344 760,420 120,420 120,282" class="nd-line" marker-end="url(#arrA)"/>
</polyline>

<text x="440" y="444" class="nd-s" text-anchor="middle">
the invite link opens the app, and the directory says where the home lives
</text>

</svg>

</div>

<figcaption>
Moments and media only ever travel between phones and their home. The project runs two small things every home needs: a push relay, and the invite domain with its directory. Neither one ever sees a moment.
</figcaption>

</figure>

### What the project must run, whatever else happens

- **A push relay.** Apple and Google only accept pushes signed with the app publisher's keys, and a self-hosted home can't hold those. Homes send a sealed wake-up to a small relay that forwards it. Mastodon's apps, Ice Cubes, and Bitwarden's apps talking to self-hosted Vaultwarden all work this way. With the payload encrypted to the phone under the Web Push standard, the relay only learns that some home woke some phone.

- **The link domain.** iOS only opens an app from links on domains listed in the app's entitlements, so the invite domain has to be yours. Put the home's ID in the link and the invite token after the `#`, which browsers never send to a server.

- **A home directory.** A small table mapping each home's ID to its current address. It lets a home move hosts without anyone needing a new invite. It also gives you a switch to cut off a home that breaks the rules, which Apple expects from the app's publisher. It knows which homes exist, not who is in them.

- **A demo home for App Review.** Reviewers need a working account on a working home.

Nothing else has to be central: no global accounts, no analytics, no copy of anyone's moments.

### The one-click ladder

Only the top rung is one click for an ordinary family. The rungs below it exist so nobody is locked in, and that is what makes the top rung trustworthy.

| Rung | Who it's for | What they do | What it costs them |
|---|---|---|---|
| Nah? Cloud | Any family | Tap "Start a home" in the app, pay, share the invite link | A few euros a month, set by you |
| PikaPods | The relative who has heard of self-hosting | Pick Nah? Home from the catalogue. TLS, updates and backups are handled | About $1–2 a month. PikaPods shares 20% of revenue with you |
| Home-server app stores | Owners of an Umbrel, CasaOS, YunoHost or TrueNAS box | Install from the store, then scan a QR code in the app | Hardware they already own |
| Container or binary | Tinkerers | Run one container or one binary on a NAS or VPS | Their own time |

Getting listed is mostly pull requests: Umbrel, CasaOS, YunoHost and TrueNAS each take a manifest in their app repositories. Coolify wants 1,000 GitHub stars first, and Hetzner's app catalogue isn't taking new apps.

Ente, the open-source encrypted photo service with a Flutter app, is the useful warning. Its self-hosting is well built, yet the app hides the "use my own server" switch behind tapping the logo seven times. Hosted is the product and self-hosting is the guarantee. Plan for the same split.

**Payment is part of the click.** Selling hosting inside the iOS app means Apple's in-app purchase, with a 15% cut under the Small Business Program. It is also the most one-click payment there is.

### Home boxes behind a router

A box in someone's living room has no fixed address, no domain and no certificate. The options, from least to most work for you:

- **Tailscale Funnel, built into the binary.** Tailscale's Go library can publish the home on a public `ts.net` address with a certificate. The family still needs a free Tailscale account, and Funnel's bandwidth cap isn't published.

- **Cloudflare Tunnel.** Solid, but it needs a domain on a Cloudflare account. The free quick tunnels are for testing only and drop the live-update connection PocketBase uses.

- **A Nah? relay, later.** Copy Home Assistant's remote access. The home dials out to a relay the project runs, and the relay routes traffic by hostname without decrypting it, so the certificate's key stays at home. A thousand homes would cost tens of euros a month in servers. Plan for Let's Encrypt's limit of 50 certificates per domain per week before you issue one per home.

### Why not peer-to-peer

Every serverless stack hits the same wall. iOS suspends background connections, so a photo posted at 9pm reaches nobody until its poster reopens the app. You end up needing an always-on peer, which is a home server by another name. None of the candidates has a Flutter-ready, stable SDK a solo developer could ship on in months. Iroh reached 1.0 in June 2026 and is worth a second look later as a way to reach home boxes without a relay, but it has no Dart bindings yet.

### What a home costs to run

These are estimates. They assume photos are re-encoded on the phone to about 550 KB including a thumbnail, voice clips are 32 kbps Opus at about 150 KB each, originals aren't kept, and storage costs Backblaze's $6.95 per terabyte per month.

| Home | Moments a day | Media a year | Storage a month | Traffic a month |
|---|---|---|---|---|
| Busy clique, 150 people | 20, half photos | about 2.5 GB | about $0.02 | 16–20 GB |
| Typical family, 12–20 people | 3 | about 0.35 GB | under $0.01 | 1–2 GB |

The smallest Hetzner server costs €5.49 a month, includes 20 TB of traffic, and can hold dozens of hosted homes. A Raspberry Pi at home uses about €1 of electricity a month. Storage and compute round to zero per home. What costs money is your time: support, backups and upgrades. Price for that, not for the disk.

### What it could earn

Nabu Casa sells Home Assistant Cloud for $6.50 a month. It is profitable without outside investors and funds the Open Home Foundation. It also sits on more than two million active Home Assistant installs, and its subscriber numbers aren't public. The model works, but it monetises a huge free base, and Nah? starts from zero.

| Paying homes at €3 a month | Before Apple's cut | After 15% |
|---|---|---|
| 100 | €300 a month | €255 a month |
| 1,000 | €3,000 a month | €2,550 a month |

A hundred homes covers the infrastructure many times over and pays for none of your time. A thousand is a small but real income. Treat hosting as what keeps Nah? alive, not as a business case, until the circle test says people want it.

### Backup and moving out

A home is one SQLite file plus a folder of media. Litestream streams the database to object storage continuously, and the media folder syncs alongside it. "Move my home" is a zip of both, downloaded from the app and restored anywhere the home binary runs. Leaving in one tap is the promise that makes paying for Nah? Cloud feel safe.

### Encryption: later, but cheap

With Nah? Cloud you could read the families' moments, exactly as the Mastodon plan admitted for its single server. Homes give a clean way out later. Because the server never searches, ranks or counts, it never needs to read a moment. Moments and media can become encrypted blobs without changing the server: the home's key travels after the `#` in the invite link, and phones make their own thumbnails. The hard part is removing someone, which means handing a new key to everyone who stays. That isn't MVP work. Just don't add server features that need to read moments, and the door stays open.

## What to decide, in order

1. **Pick the model: network or homes.** Everything below assumes homes. If you choose network, keep Mastodon and drop the one-click goal.

2. **Write two decision records.** One replaces ADR-0002 with the home server. One replaces ADR-0003 with "invitation is joining a home".

3. **Clean the specs in one pass.** Commit the MVP spec against homes, mark the January capability scaffolds as superseded, and remove the EU push claim.

4. **Build the walking skeleton.** Run one home server on your existing Komodo fleet behind a Cloudflare tunnel. Build a Flutter app that joins a home by invite link, posts text and shows the feed. Then add photo, then voice, then the digest.

5. **Spend six weeks with your circle.** Judge it by the "what would prove this wrong" lines already written in the ADRs.

6. **Only then build for other families.** Harden the push relay, then Nah? Cloud, then the one-click listings.

## Sources

### Audit and policy

- [TechCrunch: Apple reverses EU web-app removal](https://techcrunch.com/2024/03/01/apple-reverses-decision-about-blocking-web-apps-on-iphones-in-the-eu/)

- [The Register: the same reversal](https://www.theregister.com/2024/03/02/apple_reverses_pwa_decision/)

- [Digital Services Act, recital 14](https://dsa-info.eu/index.php/recital-14/)

- [Apple App Review Guidelines](https://developer.apple.com/app-store/review/guidelines/)

### Servers

- [Mastodon 4.6.0 release](https://github.com/mastodon/mastodon/releases/tag/v4.6.0) and [configuration docs](https://docs.joinmastodon.org/admin/config/)

- [Hometown maintenance issue](https://github.com/hometown-fork/hometown/issues/1358)

- [GoToSocial federation modes](https://docs.gotosocial.org/en/latest/admin/federation_modes/)

- [PocketBase as a Go framework](https://pocketbase.io/docs/go-overview) and its [Dart SDK](https://github.com/pocketbase/dart-sdk)

- [Serverpod SQLite support issue](https://github.com/serverpod/serverpod/issues/4785)

- [Litestream 0.5](https://fly.io/blog/litestream-v050-is-here/)

- [Ente self-hosting, post-install](https://ente.com/help/self-hosting/installation/post-install/)

- [Iroh 1.0](https://www.iroh.computer/blog/v1)

### Push, hosting and connectivity

- [Mastodon push docs](https://docs.joinmastodon.org/methods/push/) and [webpush-apn-relay](https://github.com/mastodon/webpush-apn-relay)

- [Vaultwarden push through Bitwarden's relay](https://github.com/dani-garcia/vaultwarden/wiki/Enabling-Mobile-Client-push-notification)

- [RFC 8291, Web Push encryption](https://www.rfc-editor.org/rfc/rfc8291)

- [PikaPods](https://www.pikapods.com/), [Umbrel apps](https://github.com/getumbrel/umbrel-apps), [Coolify service rules](https://coolify.io/docs/get-started/contribute/service)

- [Tailscale Funnel](https://tailscale.com/docs/features/tailscale-funnel) and [tsnet](https://pkg.go.dev/tailscale.com/tsnet)

- [Cloudflare quick tunnels](https://developers.cloudflare.com/cloudflare-one/networks/connectors/cloudflare-tunnel/do-more-with-tunnels/trycloudflare/)

- [Nabu Casa remote access security](https://support.nabucasa.com/hc/en-us/articles/26508882007581-Remote-access-Security-aspects), [Home Assistant Cloud](https://www.home-assistant.io/cloud/), [State of the Open Home 2025](https://www.home-assistant.io/blog/2025/04/16/state-of-the-open-home-recap/)

- [Let's Encrypt rate limits](https://letsencrypt.org/docs/rate-limits/)

- [Hetzner pricing](https://costgoat.com/pricing/hetzner) and [Backblaze B2 pricing](https://www.backblaze.com/cloud-storage/pricing)

### Competition

- [Retro raises $21M](https://techcrunch.com/2026/08/28/friend-focused-photo-sharing-app-retro-snags-21m/) and [Retro in December 2025](https://techcrunch.com/2025/12/12/retro-a-photo-sharing-app-for-friends-lets-you-time-travel-through-your-camera-roll/)

- [Yope raises $12.3M](https://techcrunch.com/2026/07/22/yope-raises-12-3m-to-build-a-private-social-network-without-algorithms-or-ads/)

- [Locket's scale](https://finance.yahoo.com/news/photo-sharing-app-locket-banking-130000795.html)

- [Friendlinq launch](https://www.pr.com/press-release/978121)

- [FamilyAlbum at 30M users](https://mixi.co.jp/en/news/2026/0507/52325/) and [Tinybeans FY25 report](https://announcements.asx.com.au/asxpdf/20250821/pdf/06n4m9jjp1zr87.pdf)

- [FUTO Circles, archived](https://github.com/circles-project/circles-android) and [Cocoon's goodbye](https://www.iphoneblog.de/2023/05/12/cocoon-is-saying-goodbye/)

- [Path shuts down](https://www.engadget.com/2018-09-17-path-private-social-network-stickers-dead.html) and [Path's FTC settlement](https://www.ftc.gov/news-events/news/press-releases/2013/02/path-social-networking-app-settles-ftc-charges-it-deceived-consumers-improperly-collected-personal)

- [Cohost shuts down](https://cohost.org/staff/post/7611443-cohost-to-shut-down)

- [Ente pricing](https://ente.com/pricing/) and [Immich 2.0](https://immich.app/blog/v2.0.0-release)

- [FTC post-trial findings on Meta](https://www.ftc.gov/system/files/ftc_gov/pdf/Plaintiff%20Federal%20Trade%20Commission%E2%80%99s%20Post-Trial%20Findings%20of%20Fact..pdf) and [Om Malik on Mosseri's memo](https://om.co/2026/01/01/what-is-instagrams-adam-mosseri-really-saying-in-his-year-end-memo/)

- [WhatsApp ads in Updates](https://www.campaignasia.com/article/whatsapp-officially-introduces-ads-in-updates-tab/h9u80y1qaupbn1aczlc6sni1lu)

*Prepared from the Nah? repository, the April decision notes in SiYuan, Linear, and two research passes on 11 September 2026. Estimates are marked as estimates.*
