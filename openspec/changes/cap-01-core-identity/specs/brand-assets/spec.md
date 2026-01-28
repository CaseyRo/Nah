# Brand Assets Spec

Logo files, app icons, splash screens, and favicon.

## ADDED Requirements

### Requirement: Logo in multiple formats

The brand SHALL have a logo available in SVG (scalable), PNG (raster at multiple sizes), and optionally as icon font or sprite.

#### Scenario: Logo for web display

- **WHEN** logo is needed for web header or about page
- **THEN** SVG version is used for crisp rendering at any size

#### Scenario: Logo for social sharing

- **WHEN** logo is needed for Open Graph image
- **THEN** PNG at appropriate resolution is available

### Requirement: App icons at required sizes

The brand SHALL provide app icons at minimum: 192x192, 512x512 (PWA manifest), 180x180 (Apple touch icon), 32x32 and 16x16 (favicon).

#### Scenario: PWA icon installation

- **WHEN** user installs PWA
- **THEN** 512x512 icon is used for home screen (high DPI)

#### Scenario: Browser favicon

- **WHEN** user views Nah in browser tab
- **THEN** favicon (32x32 or 16x16) displays in tab

#### Scenario: iOS home screen

- **WHEN** user adds to home screen on iOS
- **THEN** 180x180 Apple touch icon is used

### Requirement: Icon design matches brand identity

App icons SHALL use the pomegranate red (#EE3423) background with white or contrasting logo mark.

#### Scenario: Icon recognition

- **WHEN** user sees Nah icon among other apps
- **THEN** it is immediately recognizable by the red color and "nah" wordmark or symbol

#### Scenario: Icon corners

- **WHEN** icon is displayed on iOS
- **THEN** it works with iOS's automatic corner masking (no important content in corners)

### Requirement: Open Graph images for social sharing

The brand SHALL provide Open Graph images (og:image) for link previews on social platforms.

#### Scenario: Sharing Nah link on social media

- **WHEN** someone shares a Nah URL on Twitter/Facebook/LinkedIn
- **THEN** a branded preview image appears (1200x630 recommended)

#### Scenario: Default vs custom OG images

- **WHEN** sharing the homepage
- **THEN** default branded OG image is used
- **WHEN** sharing a user profile (future)
- **THEN** custom OG image with user info could be generated

### Requirement: Favicon with multiple formats

The brand SHALL provide favicon in ICO format (for legacy browsers) and PNG format (modern browsers).

#### Scenario: Favicon loading

- **WHEN** browser loads Nah
- **THEN** appropriate favicon format is served based on browser support

### Requirement: Splash screen assets for mobile

The brand SHALL provide splash screen images for iOS and Android PWA launch.

#### Scenario: iOS splash screen assets

- **WHEN** Nah is added to iOS home screen
- **THEN** appropriate splash images are available for different device sizes

### Requirement: Brand color consistency

All brand assets SHALL use the defined primary color (#EE3423) consistently.

#### Scenario: Color accuracy

- **WHEN** brand assets are viewed across different contexts
- **THEN** the pomegranate red is consistent (accounting for color space differences)

### Requirement: Accessibility contrast

Logo and icons SHALL maintain sufficient contrast for visibility against common backgrounds.

#### Scenario: Icon on dark background

- **WHEN** icon is shown on dark OS theme
- **THEN** it remains visible and recognizable

#### Scenario: Icon on light background

- **WHEN** icon is shown on light OS theme
- **THEN** it remains visible and recognizable
