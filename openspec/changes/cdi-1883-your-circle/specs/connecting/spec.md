## Purpose

How two people who are both in Nah? join each other's circles: by touching phones when they are together, by a link when they are not, what an invitation can and cannot do, and and why nobody types an address.

## ADDED Requirements

### Requirement: Touching phones is the ordinary way

Two people SHALL be able to connect by bringing their phones together, using NFC where both phones allow it, and a code shown on one screen and scanned by the other where they do not (CDI-1840). The exchange SHALL connect both of them, and there SHALL be no accept step afterwards. Each person's content key follows the connection, sealed to the other's phone, as CDI-1865 specifies.

#### Scenario: Two phones with NFC

- **WHEN** two people hold their phones together with Nah? open
- **THEN** both are connected, and each phone confirms it by naming the other person

#### Scenario: A phone without NFC

- **WHEN** one of the two phones cannot use NFC
- **THEN** one phone shows a code, the other scans it, and both are connected just the same

#### Scenario: The keys

- **WHEN** a connection completes by touch or by code
- **THEN** each phone receives the other person's content key sealed to it, and no server can read it

### Requirement: A link for people who are not in the same room

A person SHALL be able to make an invite link and send it through any app. The link SHALL carry the person in its path and a one-time secret after the `#`, and SHALL NOT carry a content key (CDI-1836, CDI-1865, ruled 2026-09-15), and SHALL open Nah? directly on both platforms when the app is installed (CDI-1837). Opening it SHALL say that someone invited the person, without naming them, and ask for one confirmation, after which the two are connected and each sees the other's name (CDI-1839, CDI-1895, ruled 2026-09-23).

#### Scenario: Opening a link

- **WHEN** someone who has Nah? taps an invite link in a chat app
- **THEN** Nah? opens, says someone invited them, and connects the two after one confirmation, after which each sees the other's name

#### Scenario: What a server receives

- **WHEN** an invite link is opened, in the app or in a browser
- **THEN** no server receives the part of the link after the `#`

#### Scenario: A screenshot of a used link

- **WHEN** someone photographs an invite link after it has been used
- **THEN** the photo neither connects anyone nor unlocks any moment

### Requirement: An invitation works once, never expires, and can be withdrawn

Every invitation, whether it travels as a link, a touch, a scanned code or an operator's command, SHALL work exactly once. No invitation SHALL expire. Whoever made an invitation SHALL be able to see the ones they made that are still unused, and revoke any of them (ruled 2026-09-15).

#### Scenario: Opened a second time

- **WHEN** an invitation that has already been used is opened again
- **THEN** it is refused, with a sentence saying it has already been used

#### Scenario: Months later

- **WHEN** an unused invitation is opened months after it was made
- **THEN** it still works

#### Scenario: Withdrawn

- **WHEN** the person who made an unused invitation revokes it, and someone opens it afterwards
- **THEN** it is refused, with a sentence saying it was withdrawn

### Requirement: A refused connection says why, in words

When a connection cannot be made, the app SHALL say why in one sentence written for people, and neither person's circle SHALL change. When a circle is full, the sentence SHALL say whether it is the other person's circle or the person's own, and SHALL NOT give a number (ADR-0004, ruled 2026-09-15). It SHALL name nobody, because the two are not connected yet (CDI-1895, ruled 2026-09-23). The person whose circle is full SHALL be told in words that someone could not connect because their circle is full, again with no number and no name.

#### Scenario: Their circle is full

- **WHEN** Sam opens Maya's invitation and Maya's circle already holds 150 people
- **THEN** Sam is told that the circle of the person who invited Sam is full, with no number, and neither circle changes

#### Scenario: Your own circle is full

- **WHEN** a person whose circle holds 150 people opens someone's invitation
- **THEN** they are told that their own circle is full, with no number

#### Scenario: The owner of the full circle

- **WHEN** Sam is refused because Maya's circle is full
- **THEN** Maya is told in words that someone could not connect because her circle is full, with no number and no name

### Requirement: One server, and nobody types an address

Until Nah? hosts servers for other people (M7), everyone SHALL be on one server, whose address the app carries as a build setting, so that nobody ever types or chooses an address (ruled 2026-09-23). Connecting SHALL need both people on that server. A directory that maps a person to the server they are on, keeps no record of lookups, and lets a person move server without connecting again is M7 work (CDI-1838) and is not specified here.

#### Scenario: Joining

- **WHEN** a person joins or connects
- **THEN** they are never asked for a server, a domain or an address
