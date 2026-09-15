## Purpose

The installed app: what it is called, what it opens on, what surrounds the feed, how it stays current, how it asks for notifications, and what it never talks to.

## ADDED Requirements

### Requirement: A native app called Nah?

Nah? SHALL ship as a native app for iOS and Android, installed publicly from the App Store and Google Play, with TestFlight and Play internal testing used before public release. The name under its icon SHALL be "Nah?".

#### Scenario: On the home screen

- **WHEN** Nah? is installed on an iPhone or an Android phone
- **THEN** its icon appears labelled "Nah?"

#### Scenario: First store submission

- **WHEN** the app is submitted to either store for the first time
- **THEN** it carries its final bundle identifier and application id, because neither store allows changing them afterwards

### Requirement: Opens on the feed, or asks for an invitation

On launch the app SHALL show the feed if this device has joined, signing in with the device's key without asking the person anything (ADR-0015). If the device has not joined, the app SHALL ask for an invitation, because Nah? is by invitation only. The app SHALL NOT show a sign-up form, an email field or a password field.

#### Scenario: A device that has joined

- **WHEN** the app launches on a device that has joined
- **THEN** the feed appears with no sign-in step

#### Scenario: A device that has not joined

- **WHEN** the app launches on a device that has never joined
- **THEN** it asks for an invitation and offers no other way in

### Requirement: The feed is the whole screen

The app SHALL have no tab bar and no side menu. The feed SHALL fill the screen, the person's circle and their own page SHALL open from the feed's header, and the timeline clock and the + SHALL be the only other chrome over the feed (ADR-0008). Nothing in the app and nothing on its icon SHALL carry a badge (ADR-0007).

#### Scenario: Seeing your circle

- **WHEN** the person opens their circle from the feed header and then closes that screen
- **THEN** they are back on the feed where they left it

#### Scenario: New moments have arrived

- **WHEN** moments have arrived since the person last looked
- **THEN** no dot, badge or number appears in the header, on the app icon or anywhere else

### Requirement: Staying current without a connection held open

The app SHALL load the newest page of the feed on launch, when it returns to the foreground, and when the person pulls down to refresh. It SHALL show that one page and never load a second (CDI-1833). It SHALL NOT hold a streaming connection, and SHALL NOT poll for new moments while in the background.

#### Scenario: Back from the background

- **WHEN** the app returns to the foreground
- **THEN** it loads the newest page and shows the moments posted meanwhile

#### Scenario: The end of the page

- **WHEN** the person scrolls to the end of the page
- **THEN** the feed ends there and nothing more loads

#### Scenario: In the background

- **WHEN** the app is in the background
- **THEN** it does not ask the server for new moments

### Requirement: The last feed is there offline

The app SHALL keep the last feed it loaded on the device and show it at launch, before any network response arrives and when none arrives. When the server cannot be reached, the app SHALL say so in a sentence and keep showing what it has.

#### Scenario: No connection

- **WHEN** the app opens with no connectivity, after it has loaded a feed before
- **THEN** the last loaded feed is shown, with a sentence saying it cannot reach their circle right now

#### Scenario: Cold start

- **WHEN** the app launches from a terminated state on an iPhone 13 or Pixel 6 class device
- **THEN** the stored feed is visible and scrollable within 2 seconds

### Requirement: Launch screen and system bars match the brand and the theme

The app SHALL show a native launch screen with the Nah? logo on the brand background until its first frame is drawn, on both platforms. The status bar and navigation bar content SHALL suit the current light or dark theme.

#### Scenario: Launching

- **WHEN** the app is launched on either platform
- **THEN** the logo on the brand background shows until the first frame, with no blank flash in between

#### Scenario: Dark theme

- **WHEN** the theme is dark
- **THEN** the status bar content is light

### Requirement: Notification permission comes with a reason

The app SHALL show the system notification permission prompt only after saying, in one sentence in the app, what the person will receive. Notifications SHALL arrive through Nah?'s push relay (ADR-0010), and tapping one SHALL open the feed. What a notification says and when it is sent is decided by ADR-0007, not here.

#### Scenario: Asking

- **WHEN** the app needs notification permission
- **THEN** it first says in one sentence what the person will receive, and only then shows the system prompt

#### Scenario: Declined

- **WHEN** the person declines notification permission
- **THEN** everything else in the app works as before, and the app does not ask again unless the person asks for notifications

#### Scenario: Tapping a notification

- **WHEN** the person taps a Nah? notification
- **THEN** the app opens on the feed

### Requirement: Nothing reports on the person

The app SHALL contain no analytics, crash-reporting, advertising, attribution or AI software, directly or through anything it depends on (CDI-1860). Every request the app makes itself SHALL go to the person's Nah? server or to another address Nah? runs, such as the directory or the push relay. The operating system's own push service is outside this rule.

#### Scenario: Adding a dependency

- **WHEN** a package is added to the app
- **THEN** neither it nor anything it pulls in collects analytics, reports crashes, serves advertising, attributes installs or runs AI

#### Scenario: A whole session's traffic

- **WHEN** the app's traffic is captured through joining, reading the feed and posting
- **THEN** every request goes to an address Nah? runs
