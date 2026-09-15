---
title: Rules of engagement
icon: fas fa-handshake
order: 2
---

How people come into Nah?, what happens when a connection ends, how someone gets back in on a new phone, and how they leave for good.

These rules were settled on 15 September 2026. Most of them are not built yet: today's test app connects two people by pasting an invitation. The detail sits in the [specs](https://github.com/CaseyRo/Nah/tree/main/openspec/changes), and the reasons in the [decisions]({{ '/decisions/' | relative_url }}).

## Why by invitation, and why in person

Nah? has no sign-up page. Nobody gets in without someone who is already here, and the ordinary way in is to be in the same room as that person and hold your phones together. A link exists for the people who are far away, as the exception.

Between Germans, the question in the name means: I know you, I care about you, and if anything is up, tell me. That is who a circle is for.

**Your circle is people you know.** Each person has one circle of up to 150 people, the number of stable relationships one person can hold. Everything you post goes to all of them, with no audience to pick. That only works if everyone in it is someone you would show, and then there is nobody to perform for.

**It cannot be grown.** A network you have to build in person cannot be grown by a growth loop, and no later pressure can quietly turn it into one, ours included. There is no search for people, no suggested people, no public profile and no public moment. "Viral? Nah." is built into how people arrive, not written on a wall.

**It is how these relationships start.** The people who belong in someone's 150 are overwhelmingly people they have stood next to.

**It is the safest way to connect.** When two phones touch, they exchange what they need face to face. A link can end up in the wrong group chat.

**The person who invites tells the story.** Nah? does not pitch itself to newcomers. The person inviting explains why, in their own words, in the place they already talk ([ADR-0005]({{ '/decisions/0005-ritual-onboarding/' | relative_url }})).

The full argument is in [ADR-0017]({{ '/decisions/0017-one-network-of-a-hundred-and-fifty/' | relative_url }}).

### The best invitation is an evening together

Bring the people you want in your circle to one table: a dinner, a birthday, a Sunday coffee. Anyone new installs Nah? there and then, and the phones go round. Every touch connects two people, both ways, at once, with nothing to approve afterwards.

An evening is not a room. Nobody joins "the dinner", and there are no groups in Nah?. Each touch is one decision between two people, so everyone at the table connects only with the people they are close to.

## Inviting someone

- **In person**, hold your phones together. Where a phone cannot do that, one shows a code and the other scans it. Both of you are connected, and neither has anything to accept.
- **Far away**, send a link through any app you like. It opens Nah?, shows who sent it, and connects you after one confirmation. Someone without the app lands on a plain page that names who invited them and sends them to their store, and after installing, the app carries on with the same invitation.
- **Every invitation works once.** Opened a second time, it is refused, with a sentence saying so.
- **No invitation expires.** The people who most need a link are often the slowest to install, and a dead link may be the last thing that ever happens between you.
- **You can withdraw an invitation.** You can always see the ones you made that nobody has used, and withdraw any of them, such as the link that leaked into a group chat.
- **A link holds nothing worth stealing.** It carries a one-time secret, and the part that matters never reaches any server. A screenshot of a used link connects nobody and unlocks nothing.
- **A full circle is said in words.** When the other person's circle is full, you are told it is theirs, never with a number. They are told too, in words, so they can make room if they want to. Filling the 150 is not a goal, and nobody is ever shown how many people they have.
- **Nobody can see who is in your circle**, how many there are, or who you have in common with anyone.

When someone new arrives, Nah? shows one sentence on a calm screen, then asks one question: *What do you want to share with the ones closest to you right now?* The answer sits under their name on their page and reaches their circle as their first moment. They can add a photo or skip it, and they land on a feed that shows only moments people really posted.

The first person on a new server is invited by whoever runs that server.

## When a connection ends

### Ending a connection

Either of you can end a connection at any time, and the other person is not told. Your moments leave their feed and page, and theirs leave yours, earlier ones included. A place opens in both circles.

Your phone then makes a new key for what you post next, so they cannot read anything you post afterwards. What they already saw, they saw. You can connect again later in the ordinary way.

### Blocking

Blocking ends the connection in the same way, again without telling the other person. It also refuses every later attempt to connect, by touch, code or link, until you unblock them. If they try, they are told plainly that you blocked them, so they know where they stand.

### Reporting

Nobody can read a moment except the people in its circle: not the server, and not whoever runs it. So a report carries your own copy of the moment, with who posted it and when, and it is sent to Nah?'s abuse contact only when you confirm. You are offered to block the poster as well, and they are never told who reported them.

Nah? can cut a person off, so that nobody can reach them through Nah?, without reading anything they posted. A report cannot be checked against the original, so cutting someone off takes a person's judgement, never an automatic rule.

### Deleting a moment

You can delete any moment you posted. A plain marker saying the moment was deleted stays where it was, because people may have seen it and a correction should be visible. Moments cannot be edited.

## Leaving Nah?

- **You leave from inside the app**, and everything is hidden from everyone at once: your moments, your reactions, and your place in every circle.
- **Leaving is quiet.** Nobody is told, and nothing marks where your moments stood.
- **You get 30 days to change your mind.** The app tells you the date, and coming back before then restores everything, your connections included. Until then your places in other people's circles stay held for you.
- **Then it is deleted.** On that date the server deletes your moments, media, reactions, connections, invitations and your entry in the directory. Backups are discarded within 30 days of being made, so nothing of yours survives more than 30 days after deletion.
- **Coming back after that** means joining again with a new invitation.
- **Your moments are yours to take.** You can export what you posted, unsealed on your own phone. Other people's reactions stay out of it, because they are someone else's.
- **Lost your phone and want to leave?** Get back in through your circle first, as below, and leave from the app. If you cannot, write to Nah?'s contact, and someone from your circle is asked to confirm it is you before the 30 days start.

Why this way: the app stores and European law require real deletion, started from inside the app, without undue delay. Hiding everything at once and deleting after a stated wait meets that, and it leaves a way back for someone who left in a bad moment. Keeping a leaver's moments for their circle would not be allowed, and would not be right.

## A new phone

You use Nah? on one phone, deliberately. There is no second device, no password and no recovery phrase. Your circle is how you get back in.

1. **Two people vouch, in person.** Two different people from your circle each touch your new phone, or scan its code, and confirm it is you. Nobody can vouch by link.
2. **Your circle gets 48 hours.** Everyone in your circle, and your old phone if it still has Nah?, is told that you are moving to a new phone and who vouched for you, by name. Any of them can stop it. Nothing reaches the new phone during the wait.
3. **Then you are back.** Your old phone stops working at once. Your own history comes back from one of the people who vouched, whose phone already holds your key. Everyone else's phone opens their moments to your new one the next time it comes online, because two people from your circle vouched for it, never because a server says so.

- **The same path for everyone.** Buying a new phone while the old one still works also goes through recovery, so there is no quiet way to add a reader.
- **A stopped recovery waits 7 days**, then needs two people again and tells your circle again, so nobody can wear a circle down by trying over and over.
- **Fewer than two connections means starting again.** Someone with one connection or none cannot recover. They join again with a new invitation, and their history is gone.

### Why it is this strict

The dangerous impostor is not a stranger. It is someone already in your circle, a controlling partner for instance, who holds a phone and says it is yours. One person vouching would let them in. Two people in person, with your whole circle told who they were and able to say no, is what stops it.

The cost is real: moving to a new phone takes two people and two days, even with the old phone in your hand. That is accepted.

## Related

- [ADR-0017: one network of a hundred and fifty]({{ '/decisions/0017-one-network-of-a-hundred-and-fifty/' | relative_url }}), why one circle and why in person
- [ADR-0005: ritual onboarding]({{ '/decisions/0005-ritual-onboarding/' | relative_url }}), how someone arrives
- [ADR-0004: no counts anywhere]({{ '/decisions/0004-no-counts-anywhere/' | relative_url }}), why a full circle has no number
- [ADR-0012: encrypted on the device]({{ '/decisions/0012-encrypted-on-device/' | relative_url }}) and [ADR-0015: a key on the device]({{ '/decisions/0015-no-passwords-a-key-on-the-device/' | relative_url }}), why nobody else can read a moment
- The specs: [your circle](https://github.com/CaseyRo/Nah/tree/main/openspec/changes/cdi-1883-your-circle), [leaving, blocking and reporting](https://github.com/CaseyRo/Nah/tree/main/openspec/changes/cdi-1886-user-ownership), [recovery and keys](https://github.com/CaseyRo/Nah/tree/main/openspec/changes/cdi-1865-recovery-and-key-rotation)
