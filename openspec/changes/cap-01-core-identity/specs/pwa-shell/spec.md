# PWA Shell Spec

SvelteKit PWA configuration with manifest, service worker, and offline support.

## ADDED Requirements

### Requirement: Web app manifest with app identity

The PWA SHALL include a valid manifest.json with app name, icons, theme color, and display mode.

#### Scenario: Manifest contents

- **WHEN** browser requests manifest.json
- **THEN** it returns valid JSON with name "Nah", short_name "Nah", theme_color "#EE3423", display "standalone"

#### Scenario: Icon definitions

- **WHEN** manifest is parsed
- **THEN** it includes icons at 192x192 and 512x512 sizes minimum

### Requirement: Add to home screen prompt

The PWA SHALL be installable on mobile devices via browser "Add to Home Screen" functionality.

#### Scenario: iOS installation

- **WHEN** user accesses Nah on Safari iOS
- **THEN** they can add to home screen via share menu
- **AND** app launches in standalone mode without browser chrome

#### Scenario: Android installation

- **WHEN** user accesses Nah on Chrome Android
- **THEN** browser shows install prompt (after engagement threshold)
- **AND** installed app appears in app drawer

### Requirement: Service worker for offline caching

The PWA SHALL register a service worker that caches static assets and API responses for offline access.

#### Scenario: First visit caching

- **WHEN** user first visits Nah
- **THEN** service worker caches app shell (HTML, CSS, JS, icons)

#### Scenario: Offline access

- **WHEN** user opens Nah without network connectivity
- **THEN** cached timeline content is displayed
- **AND** app shell loads normally

#### Scenario: Background sync for posts

- **WHEN** user creates a post while offline
- **THEN** post is queued locally
- **WHEN** connectivity returns
- **THEN** queued posts sync automatically

### Requirement: Splash screen configuration

The PWA SHALL display appropriate splash screens during app launch on iOS and Android.

#### Scenario: iOS splash screen

- **WHEN** user launches Nah from iOS home screen
- **THEN** splash screen shows with app icon and background color (#EE3423)

#### Scenario: Android splash screen

- **WHEN** user launches Nah from Android
- **THEN** splash screen uses manifest theme_color and icon

### Requirement: Status bar theming

The PWA SHALL configure mobile status bar color to match app theme.

#### Scenario: Status bar color

- **WHEN** app is in foreground
- **THEN** status bar uses theme color (#EE3423) or transparent overlay

### Requirement: Push notification support (opt-in)

The PWA SHALL support web push notifications when user grants permission.

#### Scenario: Permission request

- **WHEN** appropriate trigger occurs (e.g., first friend request received)
- **THEN** app requests notification permission with clear explanation

#### Scenario: Notification delivery

- **WHEN** push notification is received and permission granted
- **THEN** notification appears in system tray even if app is closed

### Requirement: Cache invalidation strategy

The service worker SHALL use appropriate caching strategies (network-first for API, cache-first for assets).

#### Scenario: API request with network

- **WHEN** user requests timeline with network available
- **THEN** fresh data is fetched and cache is updated

#### Scenario: Asset request

- **WHEN** browser requests static asset (JS, CSS, image)
- **THEN** cached version is returned immediately (cache-first)
