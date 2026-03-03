# Friend Circles Spec

Dunbar-limited connections (150 max), mutual friendship model, inner circle support.

## ADDED Requirements

### Requirement: Hard limit of 150 friends

The system SHALL enforce a maximum of 150 friends per user. This limit SHALL be enforced at the database level and cannot be bypassed.

#### Scenario: User at friend limit

- **WHEN** a user with 150 friends attempts to add a 151st friend
- **THEN** the system rejects the request with a friendly dialog: "You've reached 150 friends — Nah's limit for meaningful connections. To add someone new, you might consider reviewing your friend list."

#### Scenario: Friend count display

- **WHEN** a user views their profile or friends list
- **THEN** the system displays "X/150 friends" as positive framing

### Requirement: Visual friend cap indicator on profile

The 150-friend limit SHALL be visible on the user's profile as an intentional, beautiful design element — not just a number.

#### Scenario: Profile friend indicator

- **WHEN** user views their own or a friend's profile
- **THEN** a visual element (constellation ring, progress arc, or radial indicator) shows how many of the 150 slots are filled
- **AND** the element feels intentional and designed, not like a progress bar or limitation warning

#### Scenario: Near-capacity indicator

- **WHEN** user has 140+ friends
- **THEN** the visual indicator subtly communicates nearness to the cap (e.g., ring nearly complete, constellation nearly full)

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

### Requirement: Invite flow

Users SHALL be able to invite new people to Nah via shareable invite links.

#### Scenario: Invite location in UI

- **WHEN** user navigates to the Friends tab
- **THEN** an "Invite a Friend" button is prominently displayed at the top of the friend list
- **AND** an "Invite" option is available on the empty state screen for new users

#### Scenario: Generating an invite link

- **WHEN** user taps "Invite a Friend"
- **THEN** a shareable link is generated that can be sent via text, email, or any external channel
- **AND** the system share sheet opens for easy distribution

#### Scenario: Invite link landing — new user

- **WHEN** a non-user opens an invite link
- **THEN** they see a branded landing page with: inviter's name/avatar, Nah tagline, "Join Nah" CTA
- **AND** after account creation, the friend request from the inviter is auto-sent

#### Scenario: Invite link landing — existing user

- **WHEN** an existing Nah user opens an invite link
- **THEN** they are deep-linked to the inviter's profile with an "Accept Friend Request" prompt

#### Scenario: Pending request inbox

- **WHEN** user has pending friend requests
- **THEN** the Friends tab shows a "Pending Requests" section at the top with count badge
- **AND** each request shows the requester's avatar, name, and accept/decline buttons

#### Scenario: Invite accepted notification

- **WHEN** an invited person accepts the friend request
- **THEN** the inviter receives a push notification: "[Name] joined Nah!" (for new users) or "[Name] accepted your friend request" (for existing users)

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
