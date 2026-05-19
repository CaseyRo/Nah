# App Shell Spec

Flutter app entry point: launch, navigation scaffold, offline cache, lifecycle, platform integration. Replaces the earlier `pwa-shell` capability as part of the Flutter pivot ([decision blog post](../../../../docs/_posts/2026-05-19-the-flutter-decision.md)).

## ADDED Requirements

### Requirement: App identity in native bundles

Nah SHALL ship as a native iOS app and Android app with consistent identity (display name, bundle ID, icon, theme color) configured in platform manifests.

#### Scenario: iOS bundle configuration

- **WHEN** the iOS app is built
- **THEN** `Info.plist` declares `CFBundleDisplayName = "Nah"`, `CFBundleIdentifier` per environment, and `UIStatusBarStyle` consistent with the active theme

#### Scenario: Android manifest configuration

- **WHEN** the Android app is built
- **THEN** `AndroidManifest.xml` declares `android:label="Nah"`, `applicationId` per environment, and a theme using pomegranate red (#EE3423) as the primary color

### Requirement: App install via official stores

Nah SHALL be installable via the App Store (iOS) and Google Play (Android). Sideloading and TestFlight/internal testing tracks are acceptable during alpha but the public install path SHALL be the official stores.

#### Scenario: iOS installation

- **WHEN** a user installs Nah from the App Store
- **THEN** the app appears on the home screen with the configured icon and launches as a native iOS app

#### Scenario: Android installation

- **WHEN** a user installs Nah from Google Play
- **THEN** the app appears in the app drawer with the configured adaptive icon and launches as a native Android app

### Requirement: Cold-start performance

The app SHALL reach interactive state within 2 seconds on mid-tier devices (iPhone 13, Pixel 6) under normal cold-start conditions.

#### Scenario: Cold start measurement

- **WHEN** the app is launched from a fully terminated state
- **THEN** time-to-first-frame is < 1 second
- **AND** time-to-interactive (feed visible, scrollable) is < 2 seconds

### Requirement: Offline cache initialization

The app SHALL initialize a local persistence layer (Hive or Drift) on first launch and use it for offline timeline reading and queued moments.

#### Scenario: First launch cache bootstrap

- **WHEN** the app launches for the first time
- **THEN** the local persistence layer is initialized and ready before the feed renders

#### Scenario: Offline feed rendering

- **WHEN** a user opens the app without network connectivity
- **THEN** the previously cached timeline is displayed
- **AND** the app shell renders normally (navigation, profile, settings all accessible)

#### Scenario: Queued moments on offline post

- **WHEN** a user creates a moment while offline
- **THEN** the moment is persisted locally with a pending indicator
- **WHEN** connectivity returns
- **THEN** queued moments sync automatically and the pending indicator clears

### Requirement: Splash screen

The app SHALL display a native splash screen during launch on both platforms.

#### Scenario: iOS launch

- **WHEN** the user launches Nah on iOS
- **THEN** the iOS LaunchScreen displays the Nah logo on a pomegranate red (#EE3423) background until the first Flutter frame renders

#### Scenario: Android launch

- **WHEN** the user launches Nah on Android
- **THEN** the Android 12+ splash API (or pre-12 splash drawable) displays the Nah logo on a pomegranate red (#EE3423) background until the first Flutter frame renders

### Requirement: Status bar theming

The app SHALL configure mobile status bar color and content style to match the active theme.

#### Scenario: Light theme status bar

- **WHEN** the app is in light mode and the user is on a screen with a white surface
- **THEN** the status bar uses dark content style with a transparent or matching surface background

#### Scenario: Dark theme status bar

- **WHEN** the app is in dark mode
- **THEN** the status bar uses light content style with a transparent or matching dark surface background

### Requirement: Push notification support (opt-in)

The app SHALL support push notifications on both platforms when the user grants permission.

#### Scenario: Permission request

- **WHEN** an appropriate trigger occurs (e.g., first friend request received)
- **THEN** the app requests notification permission with a clear, friendly explanation rationale shown before the system prompt

#### Scenario: Notification delivery

- **WHEN** a push notification is received and permission has been granted
- **THEN** the notification appears in the system notification center even when the app is in the background or terminated

#### Scenario: Deep-link from notification

- **WHEN** a user taps a notification
- **THEN** the app opens directly to the relevant content (e.g., moment, friend profile, comment thread) per the deep-link destination defined in nah-vision

### Requirement: App lifecycle awareness

The app SHALL respond appropriately to lifecycle events (foreground, background, terminated).

#### Scenario: Returning from background

- **WHEN** the app returns to the foreground after being backgrounded
- **THEN** the feed checks for new content via the streaming connection
- **AND** any pending queued moments attempt to sync

#### Scenario: Backgrounding

- **WHEN** the app is backgrounded
- **THEN** in-flight network requests are allowed to complete or are persisted for retry
- **AND** the streaming connection is gracefully closed (and re-established on resume)

### Requirement: Cache strategy

The app SHALL apply a defined cache strategy per resource type.

#### Scenario: Static assets bundled with the app

- **WHEN** the app needs fonts, illustrations, or icon assets
- **THEN** these are loaded from the app bundle (no network round-trip)

#### Scenario: Timeline data

- **WHEN** the app needs timeline data
- **THEN** local cache is read first for instant render
- **AND** a background fetch updates the cache with fresh data
- **AND** the UI updates when fresh data arrives

#### Scenario: Media (photos, videos)

- **WHEN** the app loads media for the timeline
- **THEN** media is cached on disk with a size cap (e.g., 200MB) and LRU eviction
