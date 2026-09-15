## Purpose

A person's own moments, taken out of Nah? whenever they like, as plain files they can open without Nah?.

## ADDED Requirements

### Requirement: Export your own moments

A person SHALL be able to export, at any time and from the app, every moment they posted that still exists: its words, time, short line, place and song, and its media as posted. The export SHALL be one readable data file with the media files beside it. It SHALL contain nothing that anyone else posted (`product-principles`, CDI-1857).

#### Scenario: Exporting

- **WHEN** a person exports
- **THEN** they receive their moments' words, times, lines, places and songs in one readable file, with every photo, voice clip and song artwork beside it

#### Scenario: Someone else's moments

- **WHEN** the person's feed holds moments from people in their circle
- **THEN** none of those moments is in the export

#### Scenario: Opened without Nah?

- **WHEN** the export is opened on a computer that has never had Nah?
- **THEN** the moments can be read, and the media opened, with ordinary software

### Requirement: The export is unsealed only on the person's phone

Because no server can read a sealed moment, the export SHALL be assembled and unsealed on the person's phone. It SHALL be saved to the phone for the person to put wherever they choose, and no Nah? server SHALL ever hold an unsealed copy (ADR-0012).

#### Scenario: Where the export goes

- **WHEN** an export is made
- **THEN** it is saved on the phone for the person to keep or send on, and no server received it unsealed
