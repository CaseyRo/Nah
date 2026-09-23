# Your name, and how you show yourself: design

## Context

The server keeps a person's device key, connections, invitations and moments, and no name (`apps/server/README.md`). Moments travel as a versioned envelope behind a seal byte, unsealed (seal byte 0) until CDI-1863 gives each person a content key. The first run in CDI-1883 asks one question and offers an avatar photo "sealed like a moment", with nowhere yet for that photo to live. The feed reads at most `limit + 1` files, and adding a query per connection to it undoes the measured win (`spike/RESULTS.md`).

## Goals / Non-Goals

**Goals**

- Everyone in a circle can tell who posted what.
- A person can show who they are, and can always decline to.
- The server stays unable to read any of it once profiles are sealed.

**Non-Goals**

- Personalising a person's page from their answers (CDI-1898, backlog).
- Search, suggestions, filters or anything else that uses gender or orientation. Nah? has none of these (`product-principles`).
- Building anything. This round is specification only.

## Decisions

### The name comes first

Ruled 2026-09-23. It is the first step after joining, before ADR-0005's slow beat, so the calm sentence that follows can greet the person by name. It is required, because a circle has to know who posted. It is the name people recognise the person by, not a legal name or a handle: it need not be unique, is limited to 50 characters with no counter (ADR-0004), and can be changed from the person's own page.

### One profile, sealed like a moment

The name, gender, orientation and avatar photo travel together as one profile: a versioned JSON envelope behind a seal byte, exactly as a moment does (ADR-0016), held as one opaque blob in the person's file. Unknown fields are ignored, so the questions can be added without a new version. It is sealed with the person's content key once CDI-1863 lands. Until then it travels unsealed, like moments.

The server returns the profiles of the authors on a feed page in the same pass that reads their moments, and never reads one per connection. The app keeps the profiles it has seen with the stored feed.

### Gender and orientation, each with "rather not tell"

Ruled 2026-09-23. After the name, the first run asks both, with the answers below (lists ruled the same day); "rather not tell" is never preselected and is never smaller, lower or harder to reach than the others. Either answer can be changed or withdrawn later from the person's own page.

- Gender: woman, man, non-binary, in my own words, rather not tell.
- Orientation: straight, gay, lesbian, bisexual, pansexual, asexual, queer, in my own words, rather not tell.

They appear only on the person's page, to the person and their circle, under their name. "Rather not tell" shows nothing, not a label saying so. Neither answer appears on the feed.

### Orientation waits for sealing, and needs explicit consent

Sexual orientation is special-category data under GDPR Art. 9, which allows processing it on the person's explicit consent (Art. 9(2)(a)). So the question says in one sentence who will see the answer and that it can be changed or removed at any time, and choosing an answer is the consent. It is never asked while profiles travel unsealed, so the project never holds it in a form it can read. Gender is asked at the same step, after CDI-1863, which keeps the first run in one shape (CDI-1897). Until then the first run asks the name only.

### The app never guesses a pronoun

The app's own sentences use a person's name: "Maya deleted this moment", never "she deleted". Gender is how a person shows themselves, not a switch for grammar, and a name works in every language the app will speak.

### Nobody is named before a connection completes

Ruled 2026-09-23. An invite link carries a one-time secret and no content key (CDI-1865), so it cannot carry a sealed name, and putting a name in plain text in the link would show it to everyone the link reaches. So a link, the page for a phone without Nah?, and the confirmation all say that someone invited the person, and name nobody. Once the two are connected, each sees the other's name. In person, the two already know who they are touching phones with.

A refusal that happens before connecting names nobody either: the person is told that the other person's circle is full, or that their own is, and the owner of a full circle is told that someone could not connect.

## Risks / Trade-offs

- **Confirming a link blind.** The person accepts an invitation without seeing the inviter's name. Accepted: links travel through chats the two already share, and what the link says there is the context.
- **A circle learns things people choose to share.** Orientation shown to up to 150 people can be dangerous for some. "Rather not tell" is equal to every other answer, the answer can be withdrawn at any time, and nothing is shown outside the circle.
- **Special-category data on a server, sealed.** The server holds a blob it cannot read. Its metadata says only that a profile changed, and when.
- **A name can be anything.** Someone can call themselves by another person's name. Connecting in person is the check, as it is for everything else.

## Open Questions

None left open: the answer lists were ruled on 2026-09-23.
