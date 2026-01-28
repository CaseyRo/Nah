# User Ownership Spec

Data export, account portability, transparent privacy controls.

## ADDED Requirements

### Requirement: Full data export

Users SHALL be able to export all their data at any time in a standard, portable format.

#### Scenario: Requesting data export

- **WHEN** user requests a data export
- **THEN** the system generates a downloadable archive of all their data

#### Scenario: Export contents

- **WHEN** data export is generated
- **THEN** it includes: moments, photos/videos, messages, friend list, profile info, reactions received

#### Scenario: Export format

- **WHEN** data is exported
- **THEN** it uses standard formats (JSON for data, original files for media)

### Requirement: Account deletion

Users SHALL be able to permanently delete their account and all associated data.

#### Scenario: Requesting account deletion

- **WHEN** user requests account deletion
- **THEN** a confirmation process ensures they understand the permanence

#### Scenario: Deletion execution

- **WHEN** account deletion is confirmed
- **THEN** all user data is permanently removed from the system

#### Scenario: Deletion timeline

- **WHEN** deletion is requested
- **THEN** data is removed within a reasonable timeframe (e.g., 30 days)

### Requirement: Transparent privacy policy

The privacy policy SHALL be written in plain, human-readable language — not legal jargon.

#### Scenario: Reading privacy policy

- **WHEN** user reads the privacy policy
- **THEN** they can understand it without legal expertise

#### Scenario: Admin access disclosure

- **WHEN** privacy policy discusses data access
- **THEN** it honestly states that admins can technically access database (standard for hosted services)

### Requirement: No data sales or advertising

User data SHALL NEVER be sold to third parties or used for advertising purposes.

#### Scenario: Business model commitment

- **WHEN** Nah needs revenue
- **THEN** it relies on community funding, not data monetization

### Requirement: Clear data practices

Users SHALL be able to see exactly what data Nah collects and how it's used.

#### Scenario: Data transparency

- **WHEN** user wants to understand data practices
- **THEN** clear documentation explains what's collected and why

### Requirement: No third-party tracking

The app SHALL NOT include third-party analytics, tracking scripts, or advertising SDKs.

#### Scenario: Third-party code audit

- **WHEN** the codebase is reviewed
- **THEN** no third-party tracking or analytics code is present

#### Scenario: Network requests

- **WHEN** the app makes network requests
- **THEN** they only go to Nah servers, not third-party trackers
