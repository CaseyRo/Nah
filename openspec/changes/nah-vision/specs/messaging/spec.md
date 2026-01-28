# Messaging Spec

Private 1:1 and small group chat with optional ephemeral messages.

## ADDED Requirements

### Requirement: Private direct messaging

Users SHALL be able to send private messages to individual friends.

#### Scenario: Starting a conversation

- **WHEN** user initiates a message to a friend
- **THEN** a private 1:1 conversation thread is created

#### Scenario: Message delivery

- **WHEN** user sends a message
- **THEN** the recipient receives a notification (if enabled)

### Requirement: Small group messaging

Users SHALL be able to create small group conversations with multiple friends.

#### Scenario: Creating a group chat

- **WHEN** user creates a new group message
- **THEN** they can add multiple friends (up to reasonable limit, e.g., 20)

#### Scenario: Group conversation

- **WHEN** any member sends a message to the group
- **THEN** all members can see and respond

### Requirement: Messages are private to participants

Messages SHALL only be visible to conversation participants. They do not appear in timelines or feeds.

#### Scenario: Message privacy

- **WHEN** user A sends a message to user B
- **THEN** only A and B can see that message

### Requirement: Ephemeral messages (v1.5+)

Users SHALL have the option to send ephemeral messages that disappear after being read or after a time period (24 hours).

#### Scenario: Sending ephemeral message

- **WHEN** user enables ephemeral mode for a message
- **THEN** the message disappears after being read or after 24 hours

#### Scenario: Ephemeral indicator

- **WHEN** an ephemeral message is displayed
- **THEN** it is visually distinct to indicate its temporary nature

### Requirement: No message forwarding

Messages SHALL NOT be forwardable to other conversations, preserving the context and privacy of the original conversation.

#### Scenario: Attempting to forward

- **WHEN** user tries to share a message outside the conversation
- **THEN** no forwarding option exists (copy-paste still possible, but not facilitated)

### Requirement: Typing indicators

Users SHALL see when the other person is typing in a conversation.

#### Scenario: Typing indicator

- **WHEN** friend is composing a message
- **THEN** a typing indicator appears for the other participant(s)
