## Purpose

What a person tells their circle about themselves: the name everyone recognises them by, and, if they choose, their gender and sexual orientation. Only their circle sees any of it.

## ADDED Requirements

### Requirement: A name, asked first

The first step after joining SHALL ask the person for the name the people close to them recognise them by, before anything else in the first run (ruled 2026-09-23). A name SHALL be required, need not be unique, and SHALL be limited to 50 characters with no counter (ADR-0004). The person SHALL be able to change it from their own page.

#### Scenario: Joining

- **WHEN** a person has joined with an invitation
- **THEN** the first thing the app asks is their name, and the first run continues only once they have given one

#### Scenario: Changing it

- **WHEN** a person changes their name on their own page
- **THEN** their circle sees the new name on their moments from the next load

### Requirement: Every moment names who posted it

The feed, a person's page and every sentence the app says about a person SHALL name them by the name they gave. The app's own sentences SHALL use that name and never a pronoun (ruled 2026-09-23).

#### Scenario: Reading the feed

- **WHEN** a person reads a moment someone in their circle posted
- **THEN** it shows that person's name

#### Scenario: A deleted moment

- **WHEN** Maya deletes a moment
- **THEN** her circle reads "Maya deleted this moment.", with no pronoun

### Requirement: Gender and orientation, each with "rather not tell"

Once profiles are sealed (CDI-1863), the first run SHALL ask the person's gender and their sexual orientation, after the name (ruled 2026-09-23). Each question SHALL offer "rather not tell", which SHALL NOT be preselected and SHALL be as prominent and as easy to choose as every other answer. The orientation question SHALL say, in one sentence, who will see the answer and that it can be changed or removed at any time. The person SHALL be able to change or withdraw either answer from their own page. Neither question SHALL be asked while profiles travel unsealed.

#### Scenario: Rather not tell

- **WHEN** a person chooses "rather not tell" for either question
- **THEN** the first run continues exactly as it would with any other answer, and nothing about that question appears on their page

#### Scenario: Before profiles are sealed

- **WHEN** the app in use still sends profiles unsealed
- **THEN** the first run asks for the name only

#### Scenario: Withdrawing an answer

- **WHEN** a person sets their orientation to "rather not tell" on their own page
- **THEN** their circle no longer sees it, and the earlier answer is not kept

### Requirement: Only the circle sees it, and only after connecting

A person's name, gender, orientation and avatar SHALL travel together as one profile, sealed on their phone like a moment (ADR-0012, ADR-0016), and SHALL be shown only to that person and the people in their circle. Gender and orientation SHALL appear only on the person's page, never on the feed, and SHALL be used for nothing else: no search, filter, suggestion, ranking or count. An invitation, the page for a phone without Nah?, and a refusal before a connection completes SHALL name nobody (ruled 2026-09-23).

#### Scenario: Opening an invite link

- **WHEN** someone opens an invite link
- **THEN** they are told that someone invited them, and see no name until the connection completes

#### Scenario: After connecting

- **WHEN** two people complete a connection
- **THEN** each sees the other's name, and whatever the other chose to show on their page

#### Scenario: What the server holds

- **WHEN** profiles are sealed and the server stores a person's profile
- **THEN** it holds bytes it cannot read, and knows only whose profile it is and when it changed

### Requirement: The profile leaves with the person

A person's export SHALL include their own profile. Leaving Nah? SHALL hide it at once and delete it with everything else after 30 days (CDI-1886). Ending a connection SHALL stop each person seeing the other's profile.

#### Scenario: Exporting

- **WHEN** a person exports their moments
- **THEN** the export also holds their name, and the gender and orientation they gave, if any

#### Scenario: After a connection ends

- **WHEN** Sam ends his connection with Maya
- **THEN** neither can see the other's profile any longer
