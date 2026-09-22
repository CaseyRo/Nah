---
title: "Hyves is back (or: the company that shut it down is running it)"
date: 2026-09-22 20:00:00 +0200
categories: [Reflections]
tags: [hyves, business-model, privacy, advertising]
published: true
---

Hyves came back this morning, on what Mediahuis calls its 22nd birthday. It has krabbels, WieWatWaar, glitter graphics and the dancing banana. It has a chronological feed, no recommendation algorithm and no endless scroll. It is hosted in Europe, and Mediahuis, which owns it, [says](https://www.mediahuis.nl/nl/nieuws/hyves-is-terug-nederlands-sociaal-netwerk-maakt-op-22-september-comeback) no user data is shared with third parties. Just before ten the counter on hyves.nl read 10,210 sign-ups. The cap is [10,000 new accounts on day one and 5,000 a day after that](https://www.dutchnews.nl/2026/09/dutch-social-media-site-hyves-makes-a-comeback/), and the target is 1.5 million by the end of the year.

## Why I'm writing this

For ten months my brain has been cooking on a social network of my own. It is called Nah?, short for "Not Alone Here", and it is the opposite of a growth product: one circle of up to 150 people per person, joined by invitation and connected in person first, one chronological feed that ends, no public profiles, no counts, no chat, no AI, no ads. Moments are sealed on your phone before they leave it, so the server can't read them. The code is open source and the project is funded by donations. Path, the 2012 app that got intimacy right and then ran out of money, is the ancestor. The build-in-public log is at [caseyro.github.io/Nah](https://caseyro.github.io/Nah/).

So when a €1.26 billion publisher relaunches the network I grew up on, with a feature list that reads like my own decision log, I look at it the way you look at a competitor's product teardown: what did they promise, what did they ship, and who pays for it. This post is that teardown, written from the Nah? side of the table.

## The good news

A publisher with €1.26 billion in turnover last year has decided there is money in a social network that isn't Meta's. DutchNews reports that Mediahuis hopes to tap into nostalgia and distrust of American social media companies. It had Newcom survey just over 2,000 Dutch people in April, and almost 60 percent said they would probably or certainly try a new Hyves.

Read the feature list in the press release and it could be Nah?'s decision log: people you know, a chronological feed, no algorithm, no endless scroll. I've spent ten months arguing that people want this, and now a company with a sales department is betting on it too. Every evening someone spends on Hyves instead of Instagram is an evening Meta doesn't sell.

## What Hyves was

I was on Hyves, like most of the Netherlands: [10.3 million accounts in May 2010](https://en.wikipedia.org/wiki/Hyves), in a country of 16.6 million. It was Dutch, and it did a handful of things. You decorated your profile until it hurt to look at, you left krabbels on your friends' pages, and you joined the hyve for your school or your town. I met a ton of great people there.

<!-- Casey: one specific person or evening from Hyves goes here. -->

The limits were what made it work. It was local, it didn't try to be everything, and most of the people on it were people you could actually meet.

## What happened to it

In November 2010 Telegraaf Media Groep bought Hyves for [€43.7 million](https://www.agconnect.nl/artikel/hyves-kostte-telegraaf-44-miljoen-euro). In March 2013 it [wrote off €36.5 million](https://www.techzine.nl/nieuws/applications/72632/telegraaf-schrijft-365-miljoen-euro-af-op-hyves/) of that. By October 2013 [only 30 percent of visits](https://www.emerce.nl/nieuws/hyves-stopt-sociaal-netwerk) were still about the social network. On 2 December 2013 the social network closed and Hyves Games carried on. Facebook did most of the damage, and the new owner ran what was left as a games portal.

Mediahuis took over TMG in 2017, and the Hyves brand [never left the group](https://www.adformatie.nl/media/media/hyves-is-terug-inclusief-krabbels-en-dansende-bananen). The company that relaunched Hyves today is the company that bought it in 2010 and switched it off in 2013.

## What the paperwork says

I read what hyves.nl loads, what its cookie banner asks, what its App Store listing declares and which privacy statement it links to.

- **The publisher** is Mediahuis Nederland B.V. in Amsterdam, KvK 34283848, according to the [house rules](https://www.mediahuis.nl/nl/gebruiksvoorwaarden-hyves). The iPhone app is published by DT Media B.V., KvK 34231355, the Mediahuis subsidiary that also publishes Dumpert.
- **The owner** is Mediahuis NV in Antwerp, which also owns De Telegraaf, NRC, De Standaard and the Irish Independent. Its shareholders are held mainly by the Leysen, Baert and Van Puijenbroek families, and the Belgian investment company Ackermans & van Haaren [holds 14.42 percent](https://www.avh.be/en/participations/growth-capital/mediahuis). In 2025 the group turned over €1,258 million and made €157 million net.
- **The launch was paid for.** Albert Heijn, Bol, Corendon, Chiquita and hollandsnieuwe are the launching partners, and Adformatie reports that dentsu Benelux helped make the launch possible. The App Store listing says the app contains ads. The house rules have a section on how advertising works: the launch partners paid for fixed ads that are not targeted, shown "tijdelijk", for now. The same section promises that minors won't see ads based on a profile of their behaviour. It makes no such promise about adults.
- **The page loads ad tech before the app.** The source of hyves.nl pulls in a "Piano DMP", which a code comment in the page calls a "Mediahuis Ad-Tech-built component", and an ad tag from Opt Out Advertising that waits for an `AdView` to create an ad slot. Piano Software is a Philadelphia company whose DMP, formerly Cxense, sorts visitors into audience segments. In my browser, Piano's script came down from cdn.cxense.com before I had answered the cookie banner.
- **The cookie banner** asks consent for two ad-tech vendors, Opt Out Advertising and Piano, for purposes that include "create profiles for personalised advertising" and "use profiles to select personalised advertising".
- **The App Store privacy label**, filled in by DT Media B.V., lists identifiers and usage data under "Data used to track you". Apple defines tracking as linking data from an app with data from other companies' apps and websites for targeted advertising, or sharing it with data brokers.
- **The privacy statement** isn't Hyves's own. The link goes to [Mediahuis Nederland's group-wide statement](https://www.mediahuis.nl/nl/privacyverklaring), last updated on 21 September, the day before launch, and it now covers Hyves. It allows Mediahuis to share your data with every company in the group, to build profiles by combining your account with your browsing, to send your IP address to "media agencies, advertisers, ad networks and/or data partners" for automated ad sales, to hand those parties the profiles, to match your email address against an advertiser's list, and to enrich your record through data partners such as CDDN. The profiling needs your consent, and the banner is where you give it.

Some of this is better than Instagram. The house rules commit to a feed that shows your friends' posts newest first. What you post stays yours, and the licence you give Hyves only covers running Hyves, which is narrower than the one in Mediahuis's general terms. Opt Out Advertising sells contextual ads without cookies, and the launch ads are untargeted. But "no user data is shared with third parties" holds only until you press accept on a banner that names two ad-tech companies.

## Where Nah? isn't going

I expected to write that Hyves is being lined up for another sale. I found no sign of one, and it doesn't need one. The sale happened in 2010, the buyer kept the brand, and the relaunch is wired into the buyer's ad business from the first page load. A DMP, an ad server, five paying launch partners, a tracking declaration and a group privacy statement were all in place on day one.

That is where Nah? isn't going. Nah? makes the same promises about the feed as Hyves: friends only, newest first, no algorithm, an end to the page. The difference is structural. There is no place in Nah? to put an ad, and no owner who needs one:

- **No ads, no data sale, no venture capital.** Nah? is funded by donations from the people who use it. Nobody upstream needs a segment to sell.
- **Nothing to profile.** Moments are sealed on your phone with keys the project never holds. The server stores what it cannot read. A DMP pointed at Nah? would find ciphertext and a list of who knows whom.
- **Nothing to buy.** The code is AGPL. If Nah? were ever sold, the buyer would get a server full of moments it can't open, a connection graph, and a codebase anyone can fork and run without them. The thing that made Hyves worth €43.7 million in 2010, an audience to sell to, does not exist in Nah? by design.

I'll make a Hyves account this week, because some of the people I met there might come back too. I'll decline everything on the banner.
