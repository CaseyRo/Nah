# Radial Menu Spec

The signature Path-inspired "+" content creation menu with bloom animation.

## ADDED Requirements

### Requirement: Floating action button position

The radial menu trigger SHALL be a 56px circular button fixed to the bottom-right corner of the viewport with 24px inset, positioned above the bottom tab bar. A user setting SHALL allow moving the FAB to bottom-left for left-hand accessibility.

#### Scenario: Button visibility on timeline

- **WHEN** user views the home timeline
- **THEN** the red "+" FAB is visible in the bottom-right corner, above the tab bar

#### Scenario: Button visibility during scroll

- **WHEN** user scrolls the timeline
- **THEN** the FAB remains fixed in position (does not scroll away)

#### Scenario: Left-hand preference

- **WHEN** user enables "Left-hand mode" in Settings → Accessibility
- **THEN** the FAB moves to bottom-left corner

### Requirement: Five menu items in radial fan

The menu SHALL contain exactly 5 items arranged in a radial fan pattern: Photo, Text, Music, Location, Status.

#### Scenario: Menu fully expanded

- **WHEN** user opens the radial menu
- **THEN** 5 icons are visible in a fan arrangement around the trigger button

#### Scenario: Icon identification

- **WHEN** menu is open
- **THEN** each icon is visually distinct and recognizable (camera, quote/note, music note, pin, moon/sun)

#### Scenario: Removed items rationale

- "With" (friend tagging) is a composer toolbar option available on any moment type, not a standalone content type
- "Video" is combined with Photo — the camera icon opens a media picker for both photo and video
- Sleep/Wake is combined into "Status" — a single entry point with sleep/wake picker inside the composer

### Requirement: Spring animation on open

The menu SHALL animate open with spring physics — items fan out with staggered timing and slight rotation.

#### Scenario: Opening animation

- **WHEN** user taps the "+" button
- **THEN** the "+" rotates 45deg to become "x"
- **AND** menu items fan out with 30ms staggered delay
- **AND** each item scales from 0 with spring easing
- **AND** total animation completes in ~400ms

#### Scenario: Animation smoothness

- **WHEN** menu opens
- **THEN** animation runs at 60fps with no jank

### Requirement: Backdrop overlay

When the radial menu is open, a frosted glass backdrop overlay SHALL dim the content behind it. This is one of exactly 2 approved glassmorphism contexts in the app (the other being profile header compression).

#### Scenario: Backdrop appearance

- **WHEN** radial menu opens
- **THEN** a semi-transparent backdrop with blur appears behind the menu items
- **AND** CSS: `backdrop-filter: blur(12px) saturate(150%)`
- **AND** Fallback: semi-transparent dark overlay for unsupported browsers

### Requirement: Dismissal on close

The menu SHALL close when user taps the "x" button, taps outside the menu, or selects an item.

#### Scenario: Dismiss via X button

- **WHEN** user taps the "x" button
- **THEN** menu items collapse back with fade
- **AND** "x" rotates back to "+"

#### Scenario: Dismiss via outside tap

- **WHEN** user taps anywhere outside the menu
- **THEN** menu closes with same animation as X tap

#### Scenario: Dismiss via item selection

- **WHEN** user taps a menu item
- **THEN** menu closes and corresponding composer opens

### Requirement: Touch target size

Each menu item SHALL have a touch target of at least 44x44px for accessibility.

#### Scenario: Tapping menu item

- **WHEN** user attempts to tap a menu item
- **THEN** the touch target is large enough to tap without precision

### Requirement: Visual feedback on item hover/press

Menu items SHALL provide visual feedback when pressed (scale down slightly, brightness change).

#### Scenario: Pressing menu item

- **WHEN** user presses down on a menu item
- **THEN** item scales to 90% and dims slightly
- **WHEN** user releases
- **THEN** item returns to normal and action triggers

### Requirement: Accessibility

The radial menu SHALL be fully accessible via keyboard and screen readers.

#### Scenario: Keyboard navigation

- **WHEN** FAB receives keyboard focus and user presses Enter/Space
- **THEN** menu opens and focus moves to first menu item
- **AND** Arrow keys navigate between items
- **AND** Escape closes the menu and returns focus to FAB

#### Scenario: ARIA roles

- **WHEN** menu is rendered
- **THEN** FAB has `role="button"` and `aria-expanded` state
- **AND** menu container has `role="menu"`
- **AND** each item has `role="menuitem"` with descriptive label

#### Scenario: Focus order

- **WHEN** menu is open
- **THEN** focus is trapped within the menu (Tab cycles through items)
- **AND** focus order matches visual order (Photo → Text → Music → Location → Status)
