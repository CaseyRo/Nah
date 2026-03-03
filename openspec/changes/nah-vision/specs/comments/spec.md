# Comments Spec

Threaded replies on moments from friends.

## ADDED Requirements

### Requirement: Comments on moments

Users SHALL be able to post text comments on any moment shared by their friends.

#### Scenario: Adding a comment

- **WHEN** user taps the comment count or comment icon on a moment card
- **THEN** a bottom sheet opens showing existing comments and a text input
- **AND** user can type and submit a comment

#### Scenario: Comment attribution

- **WHEN** a comment is posted
- **THEN** it displays the commenter's avatar, display name, timestamp, and comment text

### Requirement: Threaded replies (single level)

Users SHALL be able to reply to a specific comment. Replies are single-level only — no nested reply chains in v1.

#### Scenario: Replying to a comment

- **WHEN** user taps "Reply" on an existing comment
- **THEN** the text input pre-fills with @mention and reply is visually grouped under the parent comment

#### Scenario: No deep nesting

- **WHEN** user replies to a reply
- **THEN** the reply is attached to the original parent comment (flat threading, not recursive)

### Requirement: Comment count on card

The moment card footer SHALL display the total comment count. Tapping it opens the comments sheet.

#### Scenario: Comment count display

- **WHEN** a moment has comments
- **THEN** the card footer shows "N comments" (e.g., "12 comments")

#### Scenario: No comments yet

- **WHEN** a moment has zero comments
- **THEN** the card footer shows a comment icon with no count (or "Comment")

### Requirement: Privacy model — friends-only visibility

Comments SHALL only be visible to friends who can see the original moment. If a moment is inner-circle-only, only inner circle friends see its comments.

#### Scenario: Friend visibility

- **WHEN** user A comments on user B's moment
- **THEN** only friends who can see user B's moment can see the comment

#### Scenario: Inner circle restriction

- **WHEN** a moment is shared to inner circle only
- **THEN** comments are only visible to inner circle members

### Requirement: Comment notifications

Moment authors and previous commenters SHALL receive notifications for new comments.

#### Scenario: Author notification

- **WHEN** a friend comments on user's moment
- **THEN** user receives a push notification (if enabled) and in-app badge

#### Scenario: Thread participant notification

- **WHEN** a reply is added to a comment thread the user participated in
- **THEN** the user receives a notification

### Requirement: Comment deletion

Users SHALL be able to delete their own comments. Moment authors SHALL be able to delete any comment on their moments.

#### Scenario: Deleting own comment

- **WHEN** user long-presses their own comment
- **THEN** a "Delete" option appears with confirmation

#### Scenario: Author moderation

- **WHEN** moment author long-presses any comment on their moment
- **THEN** a "Delete" option appears with confirmation

### Requirement: Comment character limit

Comments SHALL have a reasonable character limit to keep conversations focused.

#### Scenario: Character limit enforcement

- **WHEN** user types a comment exceeding 500 characters
- **THEN** a character counter shows remaining characters and input is capped
