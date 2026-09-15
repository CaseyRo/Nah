---
title: About
icon: fas fa-info-circle
order: 1
---

## Nah?

**A private home for your closest people.**

> *"Share with your circle, not the world."*

Nah? ("Not Alone Here") is a private social network for the people you are actually close to. Each person has one circle of up to 150 people, connected in person first and only by invitation, with one chronological feed and nobody to perform for. It is inspired by [Path](https://en.wikipedia.org/wiki/Path_(social_network)), open source, and designed never to compromise on intimacy.

How people join, leave, and get back in on a new phone is in the [Rules of engagement]({{ '/rules-of-engagement/' | relative_url }}).

## Principles

- **Small by design.** Your circle holds up to 150 people, Dunbar's number, said in words and never shown as a count.
- **By invitation, in person first.** Nobody joins without someone already here, and the ordinary way to connect is to hold two phones together.
- **Private by default.** Moments are sealed on your phone, and the server keeps what it cannot read. No public profiles, no public moments, no search for people.
- **Real friends, real moments.** Every connection is mutual. Text, voice, photo and music, with reactions only the poster sees, and no numbers about anyone.
- **No algorithmic theater.** One chronological feed that ends, and at most one daily digest unless you ask for more.
- **Nothing else in the way.** No chat, no AI, no ads, and no data for sale.
- **Open source.** AGPL, community-funded, and yours to verify.
- **Viral? Nah. Vital.**

## Tech Stack

| Layer | Technology |
|-------|------------|
| Mobile app | Flutter / Dart (iOS and Android) |
| State | Bloc / Cubit |
| Backend | Nah? Home: one static Go binary, no framework |
| Database | SQLite, one file per person |
| Sign-in | An Ed25519 key made on the phone, no passwords |
| Encryption | Moments sealed on the phone, so no server can read them (planned) |

The app moved from a SvelteKit PWA to Flutter in May 2026. In September 2026 the Mastodon fork gave way to a small server of our own, after [measuring the candidates]({{ '/posts/measuring-instead-of-arguing/' | relative_url }}) rather than arguing about them. Both moves are documented in [Decisions]({{ '/decisions/' | relative_url }}) and [Research]({{ '/research/' | relative_url }}).

## Status

🚧 **Pre-alpha.** A first working version runs on one machine: two people join by invitation, connect, post text moments and read each other's feed. Next is the same on two real phones against a deployed server. The rules of engagement were settled in September 2026, and most of them are not built yet.

## Get Involved

- **GitHub**: [CaseyRo/Nah](https://github.com/CaseyRo/Nah)
- **Specs**: [OpenSpec changes](https://github.com/CaseyRo/Nah/tree/main/openspec/changes)

Nah? is not ready for contributions yet. Watch the repository to follow along.

## Support

Nah? is community-funded, with donations only: no ads, no data sales, no VC. Supporting it changes nothing anyone can see, and everything in the app is free for everyone. Donations through Open Collective are coming soon, with public finances and a public roadmap.

---

*"Nah? is a private home for your closest people, to share life without performing for the internet."*
