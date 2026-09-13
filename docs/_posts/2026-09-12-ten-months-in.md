---
title: "Ten Months In (or: the part where there's still no code)"
date: 2026-09-12 10:00:00 +0200
categories: [Reflections, Decisions]
tags: [audit, mastodon, business-model, architecture, ai]
---

Ten months ago I wrote a prompt about an app I missed. Today this repo has thirteen commits, fourteen counting this post, eight decision records, a design system with oklch values, a blog you're reading, and not one line of app code.

That's the honest starting point for this post. I asked for [a full audit of Nah?]({{ '/research/concept-audit/' | relative_url }}) this week, and the first thing it told me was the thing I already knew and hadn't written down.

## What ten months actually produced

The good part is real. Nah? has something most side projects never get: a sentence. **No counts anywhere.** No likes, no views, no followers, no read receipts, no streaks. It's structural, you notice it in the first thirty seconds, and it costs almost nothing to build because it's mostly a list of things we don't ship.

There are eight decision records, and every one of them ends with what would prove it wrong. There's a design system down to the motion curves. There's a vision that hasn't wobbled in ten months.

The bad part is also real:

- The decisions moved in April. The specs didn't. The single MVP spec that was supposed to reconcile them was drafted in a chat and never committed, so anyone opening this repo today reads the January product.
- The specs still show a "X/150 friends" counter and a constellation ring that fills up, which directly contradicts the no-counts decision. They still describe a heart button, reactions, and an email-and-password onboarding that the April session replaced.
- Three client stack decisions in eight months. SvelteKit, then a PWA capability matrix, then Flutter. Now the backend is open again.
- One factual error I want to correct publicly: three of our documents claim EU iPhone users lose push notifications under the Digital Markets Act. That's wrong. Apple announced it for home-screen web apps in February 2024 and reversed it on 1 March 2024, and native apps were never affected at all. The Flutter decision still stands on Bluetooth and widgets, but that particular argument was never true.

The pattern underneath all of it: every time I come back to Nah?, I produce documents and a new decision instead of code. I'm aware this post is also a document. More on that at the end.

## What changed outside the repo

In March I wrote that I think of AI as an exoskeleton. Ten months later that's not a metaphor any more, it's a budget line.

A private social network used to mean a backend team, an iOS developer, an Android developer, a designer and an ops person. That's the reason "Path but again" stayed a bar conversation for a decade. It isn't that any more. One person with an exoskeleton can plausibly ship a server, a Flutter app for both platforms, and the hosting around it. Not easily. Plausibly.

But I want to be careful with that thought, because it cuts both ways. AI made building cheap. It didn't make deciding cheap, and it made producing convincing documents almost free. Ten months of evidence in this repo says my bottleneck was never the typing.

## The Mastodon question

We forked Mastodon for two reasons. The practical one was that it had already solved accounts, timelines, media and OAuth. The other one mattered more to me: I wanted to build on top of great open source software and give something back, instead of assembling yet another proprietary stack on rented infrastructure.

That second stance stays. It's not negotiable for me. But I've stopped believing Mastodon is how it gets honoured, for three reasons.

**It's too technical to be good for anyone.** That was true when we picked it and I talked myself past it. A Mastodon instance is five processes, a Postgres, a Redis, a job runner and a streaming server, and it wants two to four gigabytes of memory before anyone has posted a single photo. Managed Mastodon hosts exist precisely because a human has to babysit it. Nobody's dad is running that for the family.

**The model fights us.** Mastodon is public, federated and follow-based. Every API response carries the counts we spent a decision record removing. Mutual friendship isn't a concept it has. We'd be patching against all of that forever, on every release. Hometown is the cautionary tale here: a minimal Mastodon fork, maintained by one person, that spent over a year unable to catch up with upstream.

**Building on open source doesn't mean building on Mastodon.** SQLite, Litestream, PocketBase, Go, Flutter and Dart are all open source, all maintained, and all things a single binary can be built from. The give-back stays: Nah? is AGPL, the server would be open, and a family that stops trusting me can keep running it.

## Whose server is it, anyway?

This is the part I didn't see coming. I asked a simple-sounding question this week: could any family or friend group run their own Nah? server in one click?

It sounds like a hosting detail. It isn't. It decides who a server belongs to, and that decides who sees what.

Path's model, and ours until now, is one big server where every person has their own circle. My wife's friend is in her circle but not in mine, and the server works that out for every single moment. Self-hosting makes no sense in that world, because your circle spans other people's servers.

The other model is one small server per group. A family. A friend group. Everyone in it sees everything in it, you can belong to several, and the app stitches them into one feed. Suddenly "who can see this?" has a one-word answer, self-hosting is natural, the 150 cap becomes the size of the room rather than a number on your profile, and the thing you'd pay for is obvious.

I'm leaning hard towards the second one. It's a real change, it would supersede two decision records, and it deserves its own post once I've committed to it.

## Anti-scroll is a product decision. It's also the biggest risk

Nah? is deliberately anti-scroll and anti-sticky. No infinite feed, no counts, no streaks, one digest a day at most. A maybe-daily check-in with people you love, not an avalanche of slop.

So I had [the alternatives properly investigated]({{ '/research/what-makes-people-come-back/' | relative_url }}) this week, and the evidence is not kind to the romantic version of that idea. It is, however, unusually clear about what to do instead.

- BeReal is the closest thing to "one deliberate check-in a day" that ever reached scale. Daily users went from roughly fifteen million at the end of 2022 to about six million by the spring of 2023. It had the strongest cue anyone has ever built for this, and it decayed fastest.
- Yope sells itself on no algorithms and no ads, which is almost our sentence, and ships streaks and widgets anyway. Anti-algorithm does not apparently mean anti-habit-mechanic.
- Retro is the best-performing peer in this category, and it's weekly, not daily. Around 45% of its users engage daily, and more than half of the people who show up actually post rather than lurk.
- The products that hold people for years, Signal and Day One and Strava and the good newsletters, don't retain because they're calm. They retain because leaving costs something: an archive, a network, a vault.

So here's a correction to my own thinking. "A maybe daily check-in" isn't a daily product. It's a weekly product with a daily surface, and that surface should be a widget rather than a notification, because glancing at a widget costs nobody anything and can't wear out.

Two things follow that weren't in the plan:

1. **The archive is a feature, not a by-product.** "A year ago today" is the mechanic that makes the second year better than the first, and it breaks none of our rules.
2. **No reactions in the first version is probably wrong.** Posting into silence is how small networks die. Nobody responds, the one person who posts stops posting, and everyone else drifts out behind them. There's a formal model of this built from the Friendster collapse, and it says the network thins from the edges until it caves. A reaction is not a count. It's the minimum evidence that a moment was received, and it can ship without a single number attached.

No counts stays, without hesitation. It's cheap, it's structural, and it's still the sentence.

The honest version of the risk is this: the rules aren't what will kill Nah?. An empty feed is.

There's a split in the evidence worth sitting with, though. The formal model, built from the collapse of Friendster, says a small network caves in when its one poster stops. But every family product that survived a decade, Tinybeans and FamilyAlbum and Marco Polo, is built so that one person posting is enough, because the people receiving are treated as a real audience rather than as failed contributors. The ones that needed everybody to post, Path and Cocoon and Peach and Cohost, are all dead.

That reframes what I'm building. The lurker isn't a failure state. Someone's mother reading everything and posting nothing is the product working, and she's the reason the person posting keeps posting. It also means the digest, the widget and the archive matter more than the composer does, which is not how I've been thinking about it.

One more idea worth stealing, from Nextdoor of all places: a new neighbourhood there expires if it doesn't reach ten members within three weeks. It refuses to let a dead room exist. An empty room teaches everyone in it that the product is dead, and you can't un-teach that.

## The business model I actually believe in

Here's the thing I keep coming back to. The apps closest to Nah? raised thirty-three million dollars between them this year on the pitch of no algorithm and no ads. I don't begrudge them a cent. But watch what happens next, because it's happening already: one adds AI recaps, one adds celebrity accounts, one adds ads after being acquired.

That isn't villainy. It's arithmetic. If you raise money for a feature, product management starts answering to the raise. Every roadmap conversation acquires a second question that isn't "is this better for the people using it".

"Invite your hundred and fifty people, pay a few euros a month for the room they live in" is a worse pitch deck and a much better business. The customer is the family. The bill is small, honest and directly tied to the thing being provided. Storage for a busy group costs cents a month, so the price is paying for my time, the backups and the updates, and I can say that out loud.

I had the numbers checked, and they are a useful cold shower. Tinybeans converts about six percent of its monthly users into subscribers and still loses money. Locket has nine million people opening it every day and earns something like fifty cents per daily user per year. Nobody, anywhere, has got rich selling a private feed.

But look at what people in this category do pay for: storage, keepsakes, an archive. Tinybeans sells a scrapbook, not a timeline. And the one product that has genuinely grown for years on a chronological, algorithm-free friend feed is Letterboxd, from under two million members to over thirty million on a subscription, where the real engine is a catalogue you keep rather than a feed you watch.

So the thing worth selling isn't the timeline. It's the room, and everything ever said in it, kept safe, kept yours, and removable in one archive the day you want out. That happens to be exactly what a per-group subscription pays for.

It's not free of problems, and I'd rather name them now:

- Asking friends to pay for a social app is a genuine barrier, and "free" is what everyone else charges.
- The comparison class has trained families to expect the opposite of per-group pricing. iCloud, Google One and 1Password all sell the same sentence: add five more people at no extra cost. Extra seats are how they stop the one payer cancelling, not a reason to charge more. A room priced above what one person pays for storage is arguing with everything a family already subscribes to.
- The two businesses in this space that publish anything, Tinybeans and Ente, both convert about one user in twenty, at sixty-five to seventy-five a year. That's the planning assumption, not the optimistic one.
- One person per group can pay, which helps a lot, but it also makes that person the account holder for everyone else's memories.
- Slow money means slow growth, and a private network has no viral loop by design. That's the trade I'm choosing.
- Charging means responsibilities: hosting other people's family photos in Germany comes with GDPR, a real privacy policy, and a duty not to disappear.

The escape hatch is what makes it fair. If the hosted version ever stops being worth it, a family should be able to take their room somewhere else, or run it themselves, and have that be a genuinely supported path rather than a theoretical one.

## What happens next

The audit ended with a build order, and I'm taking it:

1. Commit to the model: one shared network, or one server per group.
2. Write the two decision records that supersede the Mastodon and connection decisions.
3. Clean the specs in one pass, so the repo stops describing a product we abandoned in April.
4. Build the walking skeleton: one server, one app, join by invite link, post text, see the feed.
5. Use it with five to ten friends for six weeks, and judge it by the falsification lines already written into the decision records.
6. Only then build hosting for other families.

Steps one to three are documents, which is exactly what I said my problem was. So I'm giving them a deadline instead of a plan, and the measure of whether this post was worth writing is simple.

The next post has a screenshot in it, or it doesn't count.

Both research documents live in the repo now, under [Research]({{ '/research/' | relative_url }}), figures, sources and all.

nah!
