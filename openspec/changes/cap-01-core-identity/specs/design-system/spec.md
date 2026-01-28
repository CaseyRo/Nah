# Design System Spec

Tailwind configuration, CSS custom properties, and shadcn-svelte theme customization.

## ADDED Requirements

### Requirement: Color tokens defined as CSS custom properties

The design system SHALL define all colors as CSS custom properties, allowing runtime theming and dark mode support.

#### Scenario: Primary color usage

- **WHEN** a component uses the primary action color
- **THEN** it references `--color-primary` (#EE3423 pomegranate red)

#### Scenario: Dark mode colors

- **WHEN** the system is in dark mode
- **THEN** surface colors swap to dark variants via CSS custom properties

### Requirement: Typography scale with system fonts

The design system SHALL use system font stack (-apple-system, BlinkMacSystemFont, etc.) with a defined size scale from xs (12px) to 2xl (24px).

#### Scenario: Body text rendering

- **WHEN** body text is displayed
- **THEN** it uses the system font at --text-base (16px) with --font-normal weight

#### Scenario: Heading rendering

- **WHEN** a section heading is displayed
- **THEN** it uses --text-xl or larger with --font-medium weight

### Requirement: Spacing scale based on 4px unit

The design system SHALL define spacing tokens in 4px increments from space-1 (4px) to space-12 (48px).

#### Scenario: Component padding

- **WHEN** a card component has padding
- **THEN** it uses spacing tokens (e.g., --space-4 for 16px)

#### Scenario: Layout gaps

- **WHEN** elements are spaced in a layout
- **THEN** gaps use spacing tokens for consistency

### Requirement: Border radius tokens

The design system SHALL define border radius tokens: sm (6px), md (8px), lg (12px), and full (pill shape).

#### Scenario: Button corners

- **WHEN** a button is rendered
- **THEN** it uses --radius-sm or --radius-full depending on style

#### Scenario: Card corners

- **WHEN** a card is rendered
- **THEN** it uses --radius-md for subtle rounding

### Requirement: Animation timing tokens

The design system SHALL define animation duration and easing tokens, including spring curves for bouncy effects.

#### Scenario: Fast micro-interaction

- **WHEN** a hover state animates
- **THEN** it uses --duration-fast (150ms) with standard easing

#### Scenario: Bouncy entrance animation

- **WHEN** a modal or menu enters
- **THEN** it uses --duration-spring (500ms) with --ease-spring (overshoot curve)

### Requirement: Tailwind config exports design tokens

The Tailwind configuration SHALL extend the default theme with all custom tokens (colors, spacing, animation).

#### Scenario: Using tokens in Tailwind classes

- **WHEN** a developer writes `bg-primary` or `p-4`
- **THEN** it maps to the defined design tokens

#### Scenario: Custom animation classes

- **WHEN** a developer needs spring animation
- **THEN** they can use `animate-spring-in` class defined in Tailwind config
