# Friend Circles Spec

Dunbar-limited connections (150 max), mutual friendship model, inner circle support.

## ADDED Requirements

### Requirement: Hard limit of 150 friends
The system SHALL enforce a maximum of 150 friends per user. This limit SHALL be enforced at the database level and cannot be bypassed.

#### Scenario: User at friend limit
- **WHEN** a user with 150 friends attempts to add a 151st friend
- **THEN** the system rejects the request with a clear explanation of the limit

#### Scenario: Friend count display
- **WHEN** a user views their profile or friends list
- **THEN** the system displays "X/150 friends" as positive framing

### Requirement: Mutual friendship model
All connections SHALL be mutual — both parties must accept for a friendship to be established. There are no one-way follows.

#### Scenario: Sending a friend request
- **WHEN** user A sends a friend request to user B
- **THEN** user B receives a notification and must explicitly accept or decline

#### Scenario: Accepting a friend request
- **WHEN** user B accepts a friend request from user A
- **THEN** both users become friends simultaneously and can see each other's content

#### Scenario: Declining a friend request
- **WHEN** user B declines a friend request from user A
- **THEN** no connection is established and user A is not notified of the decline

### Requirement: Inner circle for selective sharing
Users SHALL be able to designate a subset of friends as their "inner circle" for more selective sharing of certain moments.

#### Scenario: Marking a friend as inner circle
- **WHEN** user designates a friend as part of their inner circle
- **THEN** that friend can see posts marked "inner circle only"

#### Scenario: Posting to inner circle only
- **WHEN** user creates a moment and selects "inner circle only"
- **THEN** only inner circle friends can see that moment

### Requirement: Friend curation tools
The system SHALL provide tools to help users curate their network, including highlighting inactive friendships.

#### Scenario: Inactive friend suggestion
- **WHEN** a user hasn't interacted with a friend for an extended period
- **THEN** the system may suggest reviewing that friendship (without pressuring removal)

### Requirement: No friend discoverability
Users SHALL NOT be able to discover other users through search, suggestions, or public listings. New friends can only be added via invitation or direct sharing of profile links.

#### Scenario: No user search
- **WHEN** a user looks for ways to find new people
- **THEN** no search, explore, or suggestion features exist

#### Scenario: Adding friends by invitation
- **WHEN** a user wants to invite someone new
- **THEN** they share a direct invitation link or invite via external channel (text, email)
