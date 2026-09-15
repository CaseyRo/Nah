## Purpose

The + that starts a moment. It blooms into the three moment types and hands straight over to the chosen composer, with nothing in between.

## ADDED Requirements

### Requirement: The + stays within reach on the feed

The feed SHALL show a round + button fixed in a bottom corner, which stays in place while the feed scrolls. It SHALL sit bottom-right by default, and a left-hand setting SHALL move it to bottom-left.

#### Scenario: Scrolling

- **WHEN** the person scrolls the feed
- **THEN** the + stays where it is

#### Scenario: Left-hand setting

- **WHEN** the person turns on the left-hand setting
- **THEN** the + moves to the bottom-left corner and the menu blooms towards the right

### Requirement: Three moment types, and no audience step

Opening the + SHALL show exactly three items, text, voice and photo (ADR-0006), each with an icon and a label. Choosing one SHALL open that type's composer directly. No step SHALL ask who the moment is for, because every moment goes to the person's whole network (ADR-0017).

#### Scenario: Opening the menu

- **WHEN** the person taps the +
- **THEN** text, voice and photo appear, and nothing else

#### Scenario: Choosing an item

- **WHEN** the person taps voice
- **THEN** the menu closes and the voice composer opens
- **AND** no audience, visibility or recipient choice appears on the way

### Requirement: The bloom

When the menu opens, the + SHALL turn into a close mark, the items SHALL fan out from it one after another with the spring timings in `DESIGN.md`, and the feed behind SHALL dim. With reduce motion on, the items SHALL appear together, with no stagger, rotation or overshoot.

#### Scenario: Opening with motion

- **WHEN** reduce motion is off and the person taps the +
- **THEN** the + rotates into a close mark, the items fan out in turn, and the feed dims

#### Scenario: Opening with reduce motion

- **WHEN** reduce motion is on and the person taps the +
- **THEN** the items appear at once and the feed dims without animating

### Requirement: Closing without starting a moment

The menu SHALL close, without starting a moment, when the person taps the close mark, taps outside the items, or uses the system back gesture or button.

#### Scenario: Tapping outside

- **WHEN** the menu is open and the person taps the dimmed feed
- **THEN** the menu closes and the feed is where it was

#### Scenario: Back on Android

- **WHEN** the menu is open and the person presses back on Android
- **THEN** the menu closes and the app stays open

### Requirement: An accessible menu

The + and every item SHALL have a touch target of at least 44 by 44 points. A screen reader SHALL announce the + as a button with its open or closed state, and each item by its moment type. When the menu opens, focus SHALL move to the first item and stay within the menu. When the menu closes without a choice, focus SHALL return to the +.

#### Scenario: Opening with VoiceOver

- **WHEN** VoiceOver focus is on the + and the person activates it
- **THEN** the + is announced as expanded and focus moves to text

#### Scenario: Closing returns focus

- **WHEN** a screen-reader user closes the menu without choosing
- **THEN** focus returns to the +
