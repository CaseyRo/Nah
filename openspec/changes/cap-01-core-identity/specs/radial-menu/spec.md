# Radial Menu Spec

The signature Path-inspired "+" content creation menu with bloom animation.

## ADDED Requirements

### Requirement: Floating action button position

The radial menu trigger SHALL be a 56px circular button fixed to the bottom-left corner of the viewport with 24px inset.

#### Scenario: Button visibility on timeline

- **WHEN** user views the home timeline
- **THEN** the red "+" FAB is visible in the bottom-left corner

#### Scenario: Button visibility during scroll

- **WHEN** user scrolls the timeline
- **THEN** the FAB remains fixed in position (does not scroll away)

### Requirement: Seven menu items in radial fan

The menu SHALL contain exactly 7 items arranged in a radial fan pattern: Photo, Video, Location, Music, Text, With (tag friends), Sleep/Wake.

#### Scenario: Menu fully expanded

- **WHEN** user opens the radial menu
- **THEN** 7 icons are visible in a fan arrangement around the trigger button

#### Scenario: Icon identification

- **WHEN** menu is open
- **THEN** each icon is visually distinct and recognizable (camera, video, pin, note, quote, silhouettes, moon)

### Requirement: Spring animation on open

The menu SHALL animate open with spring physics — items fan out with staggered timing and slight rotation.

#### Scenario: Opening animation

- **WHEN** user taps the "+" button
- **THEN** the "+" rotates 45° to become "×"
- **AND** menu items fan out with 30ms staggered delay
- **AND** each item scales from 0 with spring easing
- **AND** total animation completes in ~400ms

#### Scenario: Animation smoothness

- **WHEN** menu opens
- **THEN** animation runs at 60fps with no jank

### Requirement: Dismissal on close

The menu SHALL close when user taps the "×" button, taps outside the menu, or selects an item.

#### Scenario: Dismiss via X button

- **WHEN** user taps the "×" button
- **THEN** menu items collapse back with fade
- **AND** "×" rotates back to "+"

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
