# Reactions Spec

Expressive custom reaction icons, configurable view receipts, "who viewed" transparency.

## ADDED Requirements

### Requirement: Dedicated reaction button on every card

Each moment card SHALL display a dedicated heart/reaction icon button in the card footer. This is the primary entry point for reactions — more discoverable and accessible than long-press.

#### Scenario: Adding a reaction via button

- **WHEN** user taps the heart icon on a moment card
- **THEN** a reaction picker appears showing 5 custom reaction icons (Smile, Wink, Sad, Wow, Love)

#### Scenario: Quick-react shortcut

- **WHEN** user double-taps a moment card
- **THEN** a "Love" reaction is added immediately (shortcut, skips picker)

#### Scenario: Long-press alternative

- **WHEN** user long-presses on a moment card
- **THEN** the same reaction picker appears (alternative to tapping the button)

### Requirement: Custom illustrated reaction icons

Reactions SHALL use custom SVG icons — not system emoji — that match Nah's visual language. Icons are part of the brand-assets deliverable.

**Reaction set:**

- **Smile** — default positive reaction
- **Wink** — strong like, inside joke energy
- **Sad** — sympathy, support
- **Wow** — surprise, amazement
- **Love** — deep appreciation, heart-eyes style

#### Scenario: Viewing reactions on a moment

- **WHEN** a moment has reactions
- **THEN** reactions are displayed as custom icons with the friend's avatar who reacted

#### Scenario: Multiple reaction types

- **WHEN** different friends react to the same moment
- **THEN** all reaction types are displayed as icons (not just a count)

### Requirement: Reaction display with friend context

Reactions SHALL show which friend reacted, creating a warm, personal feeling rather than anonymous counts.

#### Scenario: Viewing who reacted

- **WHEN** user views reactions on their moment
- **THEN** they see each friend's face/avatar next to their custom reaction icon

### Requirement: Configurable view receipts

View receipts SHALL be configurable at the account level. View receipts are enabled by default but users can opt out at any time.

#### Scenario: View receipt toggle

- **WHEN** user navigates to Settings → Privacy
- **THEN** they see a toggle: "Let friends see when you view their moments" (default: ON)

#### Scenario: Opting out of view receipts

- **WHEN** user turns OFF the view receipt toggle
- **THEN** their views are no longer registered or visible to moment authors

#### Scenario: View receipt generation (when enabled)

- **WHEN** a moment appears on screen for 50% visibility for 2+ continuous seconds AND the user has view receipts enabled
- **THEN** a view is registered for that user

#### Scenario: Viewing read receipts as author

- **WHEN** user views their own moment
- **THEN** they can see a list of friends who have viewed it (only those with receipts enabled)

### Requirement: "Who viewed" transparency

For users who opt in, the system SHALL be transparent about views — view receipts are honest indicators, not surveillance tools.

#### Scenario: View counts available to author

- **WHEN** moment author checks their moment
- **THEN** they can see total view count and list of viewers (opted-in only)

#### Scenario: Partial visibility acknowledged

- **WHEN** some friends have view receipts disabled
- **THEN** the view count reflects only opted-in viewers (no "and X others" for opted-out users)

### Requirement: No engagement metrics gaming

Reactions and views SHALL NOT be used to rank content or create competitive dynamics. They exist for connection, not comparison.

#### Scenario: No reaction counts in feed

- **WHEN** moments appear in timeline
- **THEN** reaction counts are subtle/secondary, not prominently displayed

#### Scenario: No "most popular" features

- **WHEN** user looks for trending or popular content
- **THEN** no such features exist
