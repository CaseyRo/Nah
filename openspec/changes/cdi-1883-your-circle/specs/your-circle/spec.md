## Purpose

The people a person is connected to. "Your circle" is the product's name for the single network of up to 150 people that ADR-0017 describes: exactly one per person, drawn around that person, and never a group or a room that others share or enter.

## ADDED Requirements

### Requirement: Seeing your circle

A person SHALL be able to open their circle from the feed's header and see everyone in it by name, in an order that does not rank them, with no count of how many there are (ADR-0004). No person SHALL be able to see who is in anyone else's circle.

#### Scenario: Opening your circle

- **WHEN** a person opens their circle
- **THEN** they see each person in it by name, and no number appears

#### Scenario: Someone else's circle

- **WHEN** a person looks at someone who is in their circle
- **THEN** nothing shows who is in that person's circle, how many there are, or who the two have in common

### Requirement: Ending a connection

Either person SHALL be able to end a connection at any time, and the other person SHALL NOT be notified. Each person's moments, past and future, SHALL leave the other's feed and page, and each SHALL be removed from the other's circle, which frees a place in both (ruled 2026-09-15). Ending a connection SHALL NOT stop the two from connecting again later in the ordinary way, unless one of them has blocked the other (CDI-1886).

#### Scenario: Ending it

- **WHEN** Sam ends his connection with Maya
- **THEN** Maya receives no notification, and from their next load neither sees any of the other's moments, earlier ones included

#### Scenario: A place opens

- **WHEN** a connection ends while one of the two circles is full
- **THEN** that circle can take one new person again

#### Scenario: Connecting again

- **WHEN** two people whose connection ended later touch phones or use an invitation
- **THEN** they are connected again, as they were the first time
