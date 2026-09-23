## Purpose

How a person leaves Nah? for good: from the app or the site, hidden from everyone at once, with 30 days to change their mind, and then deleted, as the app stores and the GDPR require.

## ADDED Requirements

### Requirement: Leaving hides everything at once

A person SHALL be able to leave Nah? from inside the app, and SHALL find a page on Nah?'s site that explains how to leave from the app, including recovering a lost phone through people in their circle first (CDI-1865). Anyone who cannot SHALL be able to write to Nah?'s published contact, and the operator SHALL confirm the request with someone in their circle before the 30 days start (ruled 2026-09-15). Leaving SHALL at once hide from everyone every moment the person posted, every reaction they gave, and their place in every circle. Nobody in their former circle SHALL be told, and nothing SHALL mark where their moments stood (ruled 2026-09-15).

#### Scenario: Leaving

- **WHEN** Sam leaves Nah?
- **THEN** from their next load, nobody sees any moment, reaction or trace of Sam

#### Scenario: Without their phone

- **WHEN** Sam wants to leave but has lost his phone
- **THEN** the site tells him to recover through his circle and leave from the app, or to write to Nah?'s contact if he cannot

#### Scenario: Nobody is told

- **WHEN** Sam leaves Nah?
- **THEN** nobody in his former circle receives a notification, and no marker stands where his moments were

### Requirement: Thirty days to come back, then deletion

When a person leaves, the app SHALL tell them the date, 30 days later, on which everything will be deleted, and that returning before then restores it. On that date the server SHALL delete the person's file with their profile, moments, media, reactions, connections and invitations, and remove their entry from the directory once there is one (M7). Until that date, their places in other people's circles SHALL stay held, so that returning restores every connection. Backups SHALL be discarded within 30 days of being made, so that nothing of theirs survives more than 30 days past deletion (ruled 2026-09-15; Apple's account deletion rules, Google Play's account deletion policy, GDPR articles 12 and 17).

#### Scenario: Told the date

- **WHEN** Sam confirms that he is leaving
- **THEN** the app shows the date his data will be deleted, and says that coming back before then restores it

#### Scenario: Coming back

- **WHEN** Sam opens Nah? on a phone that holds his key before that date, and chooses to return
- **THEN** his moments, reactions and connections are visible again

#### Scenario: The date passes

- **WHEN** 30 days pass without Sam returning
- **THEN** Sam's file, profile, moments, media, reactions, connections and invitations are deleted from the server

#### Scenario: A held place

- **WHEN** Maya was connected to Sam when he left
- **THEN** Sam's place in her circle stays taken until he returns or his data is deleted

#### Scenario: Backups

- **WHEN** 30 days have passed since Sam's data was deleted
- **THEN** no backup holds anything of his

#### Scenario: Opening the app afterwards

- **WHEN** Sam opens the app after his data was deleted
- **THEN** he is told that his account was deleted, and can join again only with a new invitation
