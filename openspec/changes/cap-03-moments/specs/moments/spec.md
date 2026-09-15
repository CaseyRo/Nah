## Purpose

What a moment is: one of three types, with an optional short line and place, prepared and sealed on the phone, never edited, and deletable by the person who posted it.

## ADDED Requirements

### Requirement: Three types of moment

A moment SHALL be exactly one of text, voice or photo (ADR-0006). A text moment SHALL be words, limited in length with no visible counter (ADR-0004). A voice moment SHALL be one recording of up to sixty seconds. A photo moment SHALL be one photo. There SHALL be no drafts: a moment is either posted or discarded (ADR-0008).

#### Scenario: Recording a voice moment

- **WHEN** a person records a voice moment and keeps talking
- **THEN** recording stops at sixty seconds, and what was recorded can be posted or discarded

#### Scenario: A long text

- **WHEN** a person types past the length a text moment allows
- **THEN** the text stops growing, and no counter or number is shown

#### Scenario: Leaving without posting

- **WHEN** a person leaves a composer without posting
- **THEN** nothing they composed is kept for later

### Requirement: An optional short line on photo and voice

The composer for a photo or voice moment SHALL offer one optional short line of text, which travels inside the moment. A text moment SHALL NOT offer one (ruled 2026-09-15).

#### Scenario: A photo with a line

- **WHEN** a person adds a short line to a photo and posts it
- **THEN** the posted moment carries the photo and that line

#### Scenario: A text moment

- **WHEN** a person composes a text moment
- **THEN** no separate short line is offered

### Requirement: An optional place, and no location kept

Every composer SHALL offer an optional place. When the person chooses to use their location, the app SHALL ask for location permission only then, after saying in one sentence what it is for. It SHALL ask the phone for its location once, and turn it into a place name through the operating system's place lookup, which the person can edit before posting. The person SHALL always be able to type a place instead, and SHALL be offered typing when location access is refused or the lookup fails. The moment SHALL carry the place as words only, sealed inside it, and neither the app nor any Nah? server SHALL keep coordinates (ruled 2026-09-15).

#### Scenario: Using the phone's location

- **WHEN** a person chooses to add their location to a moment
- **THEN** a place name appears that they can edit, and the posted moment carries only that name

#### Scenario: Location refused or unavailable

- **WHEN** location access is refused, or the lookup fails
- **THEN** the person can type the place, and the moment posts as normal

#### Scenario: What is kept

- **WHEN** a moment with a place has been posted
- **THEN** no coordinates exist in the moment, in the app's storage or on any Nah? server

### Requirement: Media is prepared and sealed on the phone

A photo SHALL be resized and re-encoded, and its thumbnail made, on the phone (CDI-1843). A voice recording SHALL be encoded on the phone (CDI-1845). Every moment and every piece of media SHALL be sealed on the phone before it is uploaded (ADR-0012), and SHALL travel in the versioned envelope with a plain-text fallback written by the posting app (ADR-0016).

#### Scenario: Posting a photo

- **WHEN** a person posts a photo taken at the camera's full resolution
- **THEN** what leaves the phone is a sealed, resized photo and a sealed thumbnail

#### Scenario: A call during a recording

- **WHEN** a call arrives while a person is recording a voice moment
- **THEN** recording stops cleanly, and the person can post or discard what was recorded once they return

### Requirement: No editing, and deleting leaves a marker

A posted moment SHALL NOT be edited (ADR-0008). The person who posted a moment SHALL be able to delete it. The server SHALL then discard its content and media and keep only who posted it and when, and wherever the moment appeared SHALL show that it was deleted (ruled 2026-09-15).

#### Scenario: Deleting

- **WHEN** Maya deletes a moment she posted
- **THEN** everyone whose feed or page held it sees, from their next load, that Maya's moment was deleted, at the time it was posted, and nothing of what it held

#### Scenario: Someone else's moment

- **WHEN** a person looks at a moment someone else posted
- **THEN** they have no way to edit or delete it

### Requirement: A moment without signal waits, and is never lost

A moment posted without a connection SHALL be kept on the phone and sent when a connection returns, and SHALL show as waiting until it is sent (CDI-1847). It SHALL survive the app being closed while it waits.

#### Scenario: No signal

- **WHEN** a person posts a voice moment with no signal
- **THEN** it shows in their feed as waiting to send, and goes out by itself once the phone is online

#### Scenario: Closed while waiting

- **WHEN** the app is closed and reopened before a waiting moment is sent
- **THEN** the moment is still waiting, and is sent once the phone is online
