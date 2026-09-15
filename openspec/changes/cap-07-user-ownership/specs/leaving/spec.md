## Purpose

How a person leaves Nah? for good: from the app or the site, hidden from everyone at once, with 30 days to change their mind, and then deleted, as the app stores and the GDPR require.

## ADDED Requirements

### Requirement: Leaving hides everything at once

A person SHALL be able to leave Nah? from inside the app, and SHALL find a page on Nah?'s site that explains how to leave and lets them ask for deletion. Leaving SHALL at once hide from everyone every moment the person posted, every reaction they gave, and their place in every circle. Nobody in their former circle SHALL be told, and nothing SHALL mark where their moments stood (ruled 2026-09-15).

#### Scenario: Leaving

- **WHEN** Sam leaves Nah?
- **THEN** from their next load, nobody sees any moment, reaction or trace of Sam

#### Scenario: Nobody is told

- **WHEN** Sam leaves Nah?
- **THEN** nobody in his former circle receives a notification, and no marker stands where his moments were

### Requirement: Thirty days to come back, then deletion

When a person leaves, the app SHALL tell them the date, 30 days later, on which everything will be deleted, and that returning before then restores it. On that date the server SHALL delete the person's file with their moments, media, reactions, connections and invitations, and remove their entry from the directory. Nothing of theirs SHALL remain in any backup for longer than the privacy statement says (ruled 2026-09-15; Apple's account deletion rules, Google Play's account deletion policy, GDPR articles 12 and 17).

#### Scenario: Told the date

- **WHEN** Sam confirms that he is leaving
- **THEN** the app shows the date his data will be deleted, and says that coming back before then restores it

#### Scenario: Coming back

- **WHEN** Sam opens Nah? on a phone that holds his key before that date, and chooses to return
- **THEN** his moments, reactions and connections are visible again

#### Scenario: The date passes

- **WHEN** 30 days pass without Sam returning
- **THEN** his file, moments, media, reactions, connections, invitations and directory entry are deleted from the server

#### Scenario: Opening the app afterwards

- **WHEN** Sam opens the app after his data was deleted
- **THEN** he is told that his account was deleted, and can join again only with a new invitation
