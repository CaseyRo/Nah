## Purpose

The guarantees Nah? makes to the people in it. Every other capability keeps them, and breaking one takes a decision record first.

## ADDED Requirements

### Requirement: One circle of up to 150 per person

Each person SHALL have one circle of at most 150 people, enforced by the server. A person's circle is the single network ADR-0017 describes: drawn around that person alone, and never a group or a room that others share or enter. There SHALL be no groups or rooms, and nothing SHALL let a person choose part of their circle as the audience for a moment.

#### Scenario: A circle that is full

- **WHEN** two people try to connect and either of them already has 150 people in their circle
- **THEN** the server refuses the connection, and neither circle changes

#### Scenario: Posting

- **WHEN** a person posts a moment
- **THEN** it goes to their whole circle, and no step offers to narrow who sees it

### Requirement: Connection is mutual, and physical first

A connection SHALL exist for both people or for neither; there SHALL be no one-way follow. The ordinary way to connect SHALL be two phones brought together in person (CDI-1840), and an invite link SHALL remain for people who are not in the same room (CDI-1839, ADR-0017).

#### Scenario: Two people connect

- **WHEN** two people complete a connection
- **THEN** each is in the other's circle, and neither can see the other's moments without the other being able to see theirs

#### Scenario: Someone far away

- **WHEN** a person wants to connect with someone who is not with them
- **THEN** they can send an invite link, and the connection completes without the two meeting

### Requirement: By invitation, and nothing public

A person SHALL join Nah? only with an invitation: from someone already in it, or, for the first person on a server, from that server's operator. There SHALL be no sign-up without one, no search for people, no suggested people, no public profile, no public moment and no federation with other networks (ADR-0008, ADR-0010). Joining SHALL never ask a person for a domain name, an address or a server setting.

#### Scenario: No invitation

- **WHEN** someone installs the app without an invitation
- **THEN** there is no way in

#### Scenario: Looking for people

- **WHEN** a person looks for someone to connect with inside the app
- **THEN** there is no search for people and no suggestion of people

#### Scenario: Joining

- **WHEN** a person joins with an invitation
- **THEN** they are never asked to type or choose a server address

### Requirement: A chronological, unranked feed

The feed SHALL show moments from the person and their circle newest first, ordered by time alone, with no ranking, no inserted content and no infinite scroll (ADR-0008, ADR-0017).

#### Scenario: Order

- **WHEN** the feed is shown
- **THEN** its moments appear in the order they were posted, and nothing appears that the person or their circle did not post

#### Scenario: Attention changes nothing

- **WHEN** one moment has been looked at far more than another
- **THEN** neither moves in the feed

### Requirement: No numbers about people

No screen, notification or message SHALL show a count of people or of what people did: no count of moments, connections, views, reads or reactions, no unread count, no streak, and no meter or progress toward the limit of 150 (ADR-0004). Where Nah? needs to say that someone did something, it SHALL name them.

#### Scenario: Close to the limit

- **WHEN** a person has 149 people in their circle
- **THEN** nothing anywhere shows 149, 150 or a proportion of either, and a full circle is said in words

#### Scenario: A notification

- **WHEN** a notification tells a person that their circle posted
- **THEN** it gives no number of people or of moments

### Requirement: The server keeps moments it cannot read

A moment's content SHALL reach the server only as an opaque blob whose first byte names how it is sealed, and the server SHALL never parse, index or transform what follows (ADR-0012, ADR-0016). Moments SHALL be sealed on the posting device with the poster's content key, so the project holds no keys. Until CDI-1863 ships, moments travel with seal byte 0, unsealed.

#### Scenario: Storing a moment

- **WHEN** the server stores a moment
- **THEN** it keeps the bytes it received unchanged, and knows only who posted it and when

#### Scenario: After CDI-1863

- **WHEN** a moment is posted from an app that holds content keys
- **THEN** its seal byte is not 0, and nothing on the server can turn it back into text or media

### Requirement: Moments stay in the circle they were posted to

Nothing in Nah? SHALL move a moment outside the circle of the person who posted it: no share, forward or repost, and no public link to a moment (CDI-1859).

#### Scenario: Passing a moment on

- **WHEN** a person views a moment someone else posted
- **THEN** Nah? offers no way to pass it on, including to people in the viewer's own circle

### Requirement: No chat

Nah? SHALL have no private messages or chats between people, one-to-one or in groups. People already have a place for that, and Nah? fills the gap next to it (ADR-0017, ruled 2026-09-15).

#### Scenario: Wanting to say something to one person

- **WHEN** a person wants to say something about a moment to just one person
- **THEN** Nah? offers no way to message them

### Requirement: No AI in the product

Nah? SHALL contain no AI features: no recaps, summaries, suggestions or generated content. No model, on the device or anywhere else, SHALL receive a person's moment (CDI-1859). Using AI tools to write Nah?'s code is separate, and allowed.

#### Scenario: A quiet week

- **WHEN** a person's circle has posted little this week
- **THEN** nothing generates a recap, a prompt or a suggestion to fill the gap

#### Scenario: Preparing a photo

- **WHEN** a photo moment is prepared on the device
- **THEN** it is resized and encoded, and Nah? passes it to no model for tagging, captioning or analysis

### Requirement: No record of where anyone is

Nah? SHALL ask a phone for its location only when the person chooses to add a place to a moment, and SHALL keep no coordinates and no location history anywhere. A place SHALL exist only as the words a person attached to a moment, sealed inside it, for that person and their circle and never for the project (ruled 2026-09-15).

#### Scenario: Adding a place

- **WHEN** a person adds a place to a moment from their phone's location
- **THEN** the moment carries the place as words, and neither the app nor any Nah? server keeps the coordinates

#### Scenario: Not adding a place

- **WHEN** a person posts a moment without a place
- **THEN** Nah? has not asked the phone where it is

### Requirement: Nothing is sold, and nothing advertises

Nah? SHALL show no advertising, and SHALL NOT sell, rent or share data about the people in it.

#### Scenario: Paying for Nah?

- **WHEN** Nah? needs money to run
- **THEN** none of it comes from advertising or from data about the people in it

### Requirement: Notifications default to the most respectful option

A person who has changed nothing SHALL be notified about new moments at most once a day, in a digest, and not at all on a day their circle posted nothing. Anything more frequent SHALL be something the person turned on (ADR-0007).

#### Scenario: A day with moments

- **WHEN** a person has not changed their notification settings and their circle posted today
- **THEN** they receive one digest about it that day, and no other notification about moments

#### Scenario: A quiet day

- **WHEN** nobody in a person's circle posted today
- **THEN** that person receives no notification about moments

### Requirement: Open code that stays traceable

Nah?'s code SHALL stay open source under AGPL-3.0-or-later, and the section 7(b) term in `NOTICE` SHALL require every modified version to keep its attribution to Nah?.

#### Scenario: A fork

- **WHEN** someone distributes a modified version of Nah?
- **THEN** its source is available under AGPL-3.0-or-later, and it carries the attribution `NOTICE` requires
