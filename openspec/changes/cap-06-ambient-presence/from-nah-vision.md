# Ambient Presence Spec

Opt-in status sharing — subtle cues that deepen presence without manual updates.

## ADDED Requirements

### Requirement: Opt-in ambient status sharing

Users SHALL be able to opt into sharing ambient status information with friends — what they're listening to, battery level, transit status, etc.

#### Scenario: Enabling ambient presence

- **WHEN** user enables ambient presence in settings
- **THEN** selected status types are shared automatically with friends

#### Scenario: Granular control

- **WHEN** user configures ambient presence
- **THEN** they can enable/disable each status type independently

### Requirement: Music status (Now Playing)

When enabled, the system SHALL share what music the user is currently listening to.

#### Scenario: Sharing music status

- **WHEN** user is listening to music and has music sharing enabled
- **THEN** friends can see "Listening to [song] by [artist]" on their profile

#### Scenario: Music status display

- **WHEN** friend views the user's profile or presence indicator
- **THEN** current music is displayed subtly

### Requirement: Location/Transit status

When enabled, the system MAY share general location context (e.g., "in transit," "at home," "arrived in [city]").

#### Scenario: Transit status

- **WHEN** user is traveling and has transit sharing enabled
- **THEN** friends see a subtle "in transit" indicator

#### Scenario: City arrival

- **WHEN** user arrives in a new city (significant location change)
- **THEN** an automatic "arrived in [city]" moment can be created (with permission)

### Requirement: Battery status (optional)

When enabled, the system MAY share battery level as a subtle presence indicator.

#### Scenario: Low battery indicator

- **WHEN** user's battery is low and battery sharing is enabled
- **THEN** friends see a subtle indicator (helps explain delayed responses)

### Requirement: Status is ambient, not intrusive

Ambient status SHALL be displayed subtly and SHALL NOT create notifications. It's background context, not attention-demanding alerts.

#### Scenario: Status display

- **WHEN** ambient status is available for a friend
- **THEN** it appears as subtle secondary information, not primary content

#### Scenario: No status notifications

- **WHEN** a friend's ambient status changes
- **THEN** no push notification is sent

### Requirement: Privacy-first defaults

All ambient presence features SHALL be opt-in and disabled by default. Users must explicitly enable sharing.

#### Scenario: New user defaults

- **WHEN** a new user joins Nah
- **THEN** all ambient presence features are disabled by default

#### Scenario: Clear privacy controls

- **WHEN** user views ambient presence settings
- **THEN** each option clearly explains what is shared and with whom
