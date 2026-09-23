## Purpose

How a person arrives in Nah?: only with an invitation, which gives them a place on a server and puts them in their inviter's circle in the same step, followed by the short first-run ritual of ADR-0005.

## ADDED Requirements

### Requirement: Joining takes an invitation

A person SHALL join only with an invitation. An invitation from someone already in Nah? SHALL register the new person and connect the two of them in one step, with nothing for either to accept afterwards. The first person on a server SHALL join with an invitation that the server's operator creates. A refused invitation SHALL leave nothing behind on the server.

#### Scenario: Invited by someone

- **WHEN** a person joins with an invitation Maya made
- **THEN** they have a place on a server, and each is in the other's circle, with no request to approve

#### Scenario: The first person on a server

- **WHEN** someone joins with an invitation the operator created
- **THEN** they are the server's first person, and their circle is empty

#### Scenario: A refused invitation

- **WHEN** someone tries to join with an invitation that cannot be used
- **THEN** the app says why in one sentence, and no trace of them remains on the server

### Requirement: An invitation survives installing the app

When an invite link is opened on a phone without Nah?, the person SHALL land on a plain page that says someone invited them, names nobody, and sends them to their platform's store (ruled 2026-09-23, CDI-1895). After installing, the app SHALL continue joining with that same invitation without the link being opened again. The invitation's secret SHALL NOT reach any server on the way (CDI-1836, CDI-1841).

#### Scenario: No app yet

- **WHEN** someone without Nah? taps an invite link from Maya
- **THEN** a page says someone invited them, without naming Maya, and offers the store for their phone

#### Scenario: After installing

- **WHEN** they install Nah? from that page and open it for the first time
- **THEN** the app continues joining with Maya's invitation, and does not ask for the link again

#### Scenario: What the host sees

- **WHEN** the page is requested
- **THEN** the server receives the link's path, and not the part after the `#`

### Requirement: The first run is a ritual

After joining, the app SHALL first ask the person's name, and once profiles are sealed their gender and orientation, as CDI-1895 specifies. It SHALL then show a single sentence alone on a calm screen, fading in over a few seconds, with no way to skip it. It SHALL then ask one real question, "What do you want to share with the ones closest to you right now?" (ADR-0005, wording ruled 2026-09-15), and the answer SHALL appear under the person's name on their page and SHALL reach their circle as their first moment (ruled 2026-09-15). After the question, it SHALL offer a photo for the person's avatar, kept in their sealed profile (CDI-1895); the person SHALL be able to skip it and keep their initials, and to add or change the photo later from their own page (ruled 2026-09-15). The person SHALL arrive with a single soft haptic tap and nothing celebratory, and SHALL land on their feed. That feed SHALL show only moments that people in their circle really posted.

#### Scenario: The slow beat

- **WHEN** the person has given their name
- **THEN** one sentence fades in on a calm screen, and nothing on it lets the person skip ahead

#### Scenario: The answer

- **WHEN** a new person answers the first-run question
- **THEN** the answer appears under their name on their page, and reaches their circle as their first moment

#### Scenario: Skipping the photo

- **WHEN** a new person skips the avatar photo in the first run
- **THEN** they arrive with their initials as their avatar, and can add a photo later from their own page

#### Scenario: Arriving

- **WHEN** the person reaches their feed for the first time
- **THEN** the phone gives one soft haptic tap, and no confetti, badge or welcome banner appears

#### Scenario: An inviter who has posted

- **WHEN** the person's inviter has posted moments
- **THEN** the inviter's most recent moments appear in the new person's feed

#### Scenario: An inviter who has not posted

- **WHEN** nobody in the new person's circle has posted anything
- **THEN** the feed names the inviter by the name they gave and says it is quiet, and shows no example or placeholder moment
