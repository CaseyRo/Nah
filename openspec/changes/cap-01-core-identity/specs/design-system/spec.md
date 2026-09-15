## Purpose

How Nah? looks and moves, stated as behaviour that can be checked. Every value (colour, type scale, spacing, radius, shadow, motion timing) lives in `DESIGN.md` at the repository root, and this spec never repeats one.

## ADDED Requirements

### Requirement: DESIGN.md is the one source of design values

The app SHALL take every colour, type size and weight, spacing, radius, shadow and motion timing from `DESIGN.md`. Screens SHALL read these values from the app's theme, and no value SHALL be defined a second time in a spec or in an individual screen.

#### Scenario: A value changes

- **WHEN** a value in `DESIGN.md` changes and the theme is updated to match
- **THEN** every screen that uses it changes with no further edit

#### Scenario: Primary action colour

- **WHEN** a primary action is drawn in the light theme
- **THEN** it uses the primary colour `DESIGN.md` defines for the light theme

### Requirement: Light and dark follow the system

The app SHALL use `DESIGN.md`'s light palette when the system appearance is light and its dark palette when the system appearance is dark, and SHALL switch while running.

#### Scenario: System set to dark

- **WHEN** the system appearance is dark
- **THEN** surfaces and text use the dark palette
- **AND** surfaces that carry a shadow in light mode carry a 1px border instead, as `DESIGN.md` specifies

#### Scenario: Appearance changes while open

- **WHEN** the system appearance changes while Nah? is in the foreground
- **THEN** the app switches palette without restarting and without losing what is on screen

### Requirement: Display font for brand voice, platform font for everything else

The app SHALL set the logo, screen titles and onboarding headings in the display font `DESIGN.md` names, and SHALL set body text, labels, buttons, timestamps and form fields in the platform's default sans.

#### Scenario: A button

- **WHEN** a button is drawn
- **THEN** its label uses the platform sans, never the display font

#### Scenario: A text moment

- **WHEN** a text moment is shown in the feed
- **THEN** its text uses the platform sans

### Requirement: Design assets ship inside the app

The display font and every icon and illustration SHALL be bundled with the app. The app SHALL NOT fetch fonts, icons or other design assets from any network location.

#### Scenario: First launch with no network

- **WHEN** Nah? is opened for the first time with no connectivity
- **THEN** titles render in the display font

#### Scenario: Traffic during launch

- **WHEN** the app's network traffic is captured during launch
- **THEN** no request goes to a font, icon or asset host

### Requirement: Spring motion, and reduced motion honoured

Entrances, presses and dismissals SHALL animate with the spring timings in `DESIGN.md` and SHALL NOT use bounce or elastic curves. When the system's reduce-motion setting is on, each of these animations SHALL become an instant change or a short opacity fade.

#### Scenario: Reduce motion off

- **WHEN** reduce motion is off and a bottom sheet opens
- **THEN** it enters with the spring entrance timing from `DESIGN.md`, without bouncing

#### Scenario: Reduce motion on

- **WHEN** reduce motion is on and a bottom sheet opens
- **THEN** it appears without sliding, springing or overshooting

### Requirement: Readable and operable by everyone

In both palettes, all text SHALL meet WCAG 2.1 AA contrast against whatever it is drawn on. No text SHALL be set on the brand primary colour: a primary action SHALL be a primary-colour circle carrying a glyph with at least 3:1 contrast, with its label beside the circle. Every touch target SHALL be at least 44 by 44 points. Every interactive element SHALL be announced by VoiceOver and TalkBack with a name and a role, and text SHALL scale with the system text size.

#### Scenario: A primary action

- **WHEN** a primary action such as joining or sending is drawn, in either palette
- **THEN** its primary-colour fill carries only a glyph at 3:1 or more, and its label sits beside the fill at 4.5:1 or more

#### Scenario: Screen reader

- **WHEN** VoiceOver or TalkBack moves across the feed screen
- **THEN** every control is announced with a name and a role, in visual order

#### Scenario: Largest text size

- **WHEN** the system text size is set to its largest accessibility size
- **THEN** moment text grows, and no control is clipped or out of reach
