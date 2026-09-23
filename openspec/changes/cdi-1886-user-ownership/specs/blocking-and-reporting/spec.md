## Purpose

What a person can do when someone in Nah? is a problem: block them, and report what they posted, in a network whose server cannot read anything. Apple's guideline 1.2 requires both, and a published contact, for private networks too (CDI-1867).

## ADDED Requirements

### Requirement: Blocking

A person SHALL be able to block anyone in their circle. Blocking SHALL end the connection as ending one does (CDI-1883), without notifying the blocked person, and SHALL refuse every later attempt to connect the two, by touch, code or invitation, until the person who blocked chooses to unblock. When the blocked person tries, the refusal SHALL tell them plainly that they have been blocked (ruled 2026-09-15).

#### Scenario: Blocking

- **WHEN** Maya blocks Sam
- **THEN** they are no longer connected, Sam receives no notification, and neither sees any of the other's moments

#### Scenario: Trying to connect again

- **WHEN** Sam later opens an invitation Maya made, or touches phones with her
- **THEN** the connection is refused, and Sam is told plainly that Maya has blocked him

#### Scenario: Unblocking

- **WHEN** Maya unblocks Sam
- **THEN** the two can connect again in the ordinary way, and are not connected until they do

### Requirement: Reporting a moment

A person SHALL be able to report a moment posted by someone in their circle. Because no server can read moments, the report SHALL carry the reporter's own unsealed copy of the moment, with who posted it and when, and SHALL be sent to Nah?'s abuse contact only when the reporter confirms (CDI-1867). Reporting SHALL offer to block the poster as well. The poster SHALL NOT be told who reported them.

#### Scenario: Reporting

- **WHEN** Ana reports one of Sam's moments and confirms
- **THEN** Nah?'s abuse contact receives that moment as Ana sees it, with Sam as its poster and its time, and Ana is offered to block Sam

#### Scenario: Not confirming

- **WHEN** Ana starts a report and does not confirm it
- **THEN** nothing is sent

#### Scenario: The reported person

- **WHEN** a report about one of Sam's moments is sent
- **THEN** nothing tells Sam who sent it

### Requirement: A published contact, and cutting someone off

Nah? SHALL publish an abuse contact, in the app and on its site, that anyone can reach without joining. Nah?'s operator SHALL be able to cut a person off from the server and from push, and from the directory once there is one (M7, CDI-1838), so that nobody can reach them through Nah?, without the operator reading anything they posted (CDI-1867).

#### Scenario: Finding the contact

- **WHEN** someone looks for how to report abuse
- **THEN** the contact is in the app and on the site, and reaching it needs no account

#### Scenario: Cutting someone off

- **WHEN** the operator cuts a person off
- **THEN** the server no longer serves them or anything of theirs and no push reaches them, and the operator has read none of their moments
