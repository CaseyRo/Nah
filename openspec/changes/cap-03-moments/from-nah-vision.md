# Moments Spec

Timeline of life moments — photos, text, music, location, sleep/wake status.

## ADDED Requirements

### Requirement: Chronological home timeline

The home timeline SHALL display moments from friends in reverse chronological order. There SHALL be no algorithmic ranking or engagement optimization.

#### Scenario: Viewing home timeline

- **WHEN** user opens the app
- **THEN** they see their friends' moments in chronological order (newest first)

#### Scenario: No algorithmic interference

- **WHEN** moments are displayed
- **THEN** order is strictly by timestamp, not by predicted engagement

### Requirement: Multiple moment types

Users SHALL be able to create different types of moments: photo/video, text, music, location, sleep/wake status, and "with" tags.

#### Scenario: Creating a photo moment

- **WHEN** user selects photo from radial menu
- **THEN** they can capture or select a photo and add optional caption

#### Scenario: Creating a music moment

- **WHEN** user selects music from radial menu
- **THEN** they can share what they're listening to (manual entry or integration)

#### Scenario: Creating a location moment

- **WHEN** user selects location from radial menu
- **THEN** they can share their current location with configurable precision

#### Scenario: Creating a sleep/wake moment

- **WHEN** user selects sleep or wake from menu
- **THEN** a status moment is posted indicating they went to sleep or woke up

#### Scenario: Tagging friends ("with")

- **WHEN** user creates any moment
- **THEN** they can optionally tag friends they're with

### Requirement: All moments default to followers-only

Every moment SHALL default to "followers-only" visibility. Users cannot post publicly.

#### Scenario: Creating any moment

- **WHEN** user creates a moment
- **THEN** visibility is automatically set to followers-only (friends)

#### Scenario: No public posting option

- **WHEN** user looks for visibility options
- **THEN** there is no "public" option available

### Requirement: Rich moment display

Moments SHALL be displayed with rich formatting appropriate to their type — music moments show album art, location moments show maps, etc.

#### Scenario: Viewing a music moment

- **WHEN** a music moment appears in timeline
- **THEN** it displays song name, artist, and album art if available

#### Scenario: Viewing a location moment

- **WHEN** a location moment appears in timeline
- **THEN** it displays location name and optional map preview

### Requirement: Media optimization

Photos and videos SHALL be compressed/resized on the client before upload to optimize storage and bandwidth.

#### Scenario: Uploading a photo

- **WHEN** user attaches a high-resolution photo
- **THEN** the client compresses it to reasonable quality before upload

### Requirement: Timeline as personal journal

The user's own profile SHALL function as a personal journal/timeline of their moments, viewable by themselves and their friends.

#### Scenario: Viewing own profile

- **WHEN** user views their own profile
- **THEN** they see all their moments as a chronological journal

#### Scenario: Friend viewing profile

- **WHEN** a friend views another friend's profile
- **THEN** they see that friend's moments (respecting inner circle settings)
