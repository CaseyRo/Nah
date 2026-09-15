## Purpose

How a person answers a moment without writing anything: five illustrated reactions, seen only by the person who posted the moment, by name and never by number.

## ADDED Requirements

### Requirement: Five illustrated reactions

Each moment card SHALL offer a reaction button that opens five reactions, Smile, Wink, Sad, Wow and Love, drawn as Nah?'s own illustrations rather than system emoji (ADR-0018). A person SHALL be able to give a moment one reaction, change it, and take it back. A person SHALL NOT be able to react to their own moment (ruled 2026-09-15).

#### Scenario: Reacting

- **WHEN** Sam taps the reaction button on Maya's moment and picks Wow
- **THEN** Sam's Wow is recorded on Maya's moment, and Sam's card shows the Wow he gave

#### Scenario: Changing it

- **WHEN** Sam later picks Love on the same moment
- **THEN** Love replaces his Wow, and Maya's moment still has one reaction from Sam

#### Scenario: Taking it back

- **WHEN** Sam takes his reaction back
- **THEN** Maya's moment no longer has a reaction from Sam

#### Scenario: Your own moment

- **WHEN** Maya looks at a moment she posted
- **THEN** it offers her no reaction button

#### Scenario: The same on every phone

- **WHEN** a reaction is shown on an iPhone and on an Android phone
- **THEN** the illustration is the same on both

### Requirement: Only the poster sees who reacted

The poster SHALL see each reaction on their moment with the reacting person's name and face and the reaction they chose. Nobody else SHALL see any reaction except their own (ADR-0018), and no count of reactions SHALL appear anywhere (ADR-0004).

#### Scenario: The poster

- **WHEN** Maya opens her moment that Sam and Ana reacted to
- **THEN** she sees Sam's face and name with his reaction, Ana's with hers, and no number

#### Scenario: Someone else in the circle

- **WHEN** Ana looks at Maya's moment
- **THEN** she sees her own reaction, and nothing of Sam's

#### Scenario: Many reactions

- **WHEN** forty people have reacted to a moment
- **THEN** its poster sees forty faces and names, with no number and no "and others"

### Requirement: Reactions reach the poster in the daily digest

A reaction SHALL NOT send a notification of its own. On a day when people reacted to a person's moments, that person's daily digest (ADR-0007) SHALL say who reacted, by name and never by number (ruled 2026-09-15). When more than three people reacted, it SHALL name three and say that more of the person's circle reacted, without a number.

#### Scenario: A day with reactions

- **WHEN** Maya and Sam react to one of Ana's moments today
- **THEN** no notification is sent as they react, and Ana's digest names Maya and Sam without a number

#### Scenario: Many reactions in a day

- **WHEN** twelve people react to Ana's moment today
- **THEN** Ana's digest names three of them and says that more of her circle reacted, with no number

### Requirement: The server cannot tell one reaction from another

A reaction SHALL be sealed on the reacting phone, so that the server knows only that a person reacted to a moment and when (ADR-0012). The server SHALL return a moment's reactions only to its poster, and SHALL delete them when the moment is deleted, or when the connection between the two people ends (CDI-1883).

#### Scenario: Storing a reaction

- **WHEN** the server stores Sam's reaction to Maya's moment
- **THEN** nothing on the server can tell whether it was Sad or Love

#### Scenario: Someone else asks

- **WHEN** anyone other than Maya asks the server for the reactions on Maya's moment
- **THEN** the server refuses

#### Scenario: A connection ends

- **WHEN** Sam and Maya's connection ends
- **THEN** Sam's reactions on Maya's moments, and Maya's on Sam's, are deleted

#### Scenario: A deleted moment

- **WHEN** Maya deletes a moment that has reactions
- **THEN** its reactions are deleted with it, and its marker shows none

### Requirement: No record of who saw a moment

Nah? SHALL NOT record or show who saw a moment, or whether anyone did (ADR-0008).

#### Scenario: A moment nobody reacted to

- **WHEN** Maya looks at a moment of hers that has no reactions
- **THEN** nothing tells her who saw it, or whether anyone did
