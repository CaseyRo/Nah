# Reactions Spec

Expressive emoji reactions, read receipts, "who viewed" transparency.

## ADDED Requirements

### Requirement: Expressive emoji reactions

Users SHALL be able to react to moments with expressive emoji reactions (smile, laugh, love, wow, sad) rather than a simple "like" button.

#### Scenario: Adding a reaction

- **WHEN** user taps/holds on a moment
- **THEN** a selection of emoji reactions appears for them to choose

#### Scenario: Viewing reactions on a moment

- **WHEN** a moment has reactions
- **THEN** reactions are displayed with the friend's avatar who reacted

#### Scenario: Multiple reaction types

- **WHEN** different friends react to the same moment
- **THEN** all reaction types are displayed (not just a count)

### Requirement: Reaction display with friend context

Reactions SHALL show which friend reacted, creating a warm, personal feeling rather than anonymous counts.

#### Scenario: Viewing who reacted

- **WHEN** user views reactions on their moment
- **THEN** they see each friend's face/avatar next to their reaction

### Requirement: Read receipts for moments

Users SHALL be able to see which friends have viewed their moments.

#### Scenario: Viewing read receipts

- **WHEN** user views their own moment
- **THEN** they can see a list of friends who have viewed it

#### Scenario: Read receipt privacy

- **WHEN** a user views a friend's moment
- **THEN** the friend can see that they viewed it (no anonymous viewing)

### Requirement: "Who viewed" transparency

The system SHALL be transparent about views — there is no way to view content anonymously within your friend circle.

#### Scenario: Viewing content creates receipt

- **WHEN** a moment appears on screen for sufficient time
- **THEN** a view is registered for that user

#### Scenario: View counts available to author

- **WHEN** moment author checks their moment
- **THEN** they can see total view count and list of viewers

### Requirement: No engagement metrics gaming

Reactions and views SHALL NOT be used to rank content or create competitive dynamics. They exist for connection, not comparison.

#### Scenario: No reaction counts in feed

- **WHEN** moments appear in timeline
- **THEN** reaction counts are subtle/secondary, not prominently displayed

#### Scenario: No "most popular" features

- **WHEN** user looks for trending or popular content
- **THEN** no such features exist
