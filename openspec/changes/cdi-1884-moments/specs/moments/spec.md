## Purpose

What a moment is: one of four types, with an optional short line and place, prepared and sealed on the phone, never edited, and deletable by the person who posted it.

## ADDED Requirements

### Requirement: Four types of moment

A moment SHALL be exactly one of text, voice, photo or music (ADR-0006, amended 2026-09-15). A text moment SHALL be words, limited in length with no visible counter (ADR-0004). A voice moment SHALL be one recording of up to sixty seconds. A photo moment SHALL be one photo. A music moment SHALL be one song. There SHALL be no drafts: a moment is either posted or discarded (ADR-0008).

#### Scenario: Recording a voice moment

- **WHEN** a person records a voice moment and keeps talking
- **THEN** recording stops at sixty seconds, and what was recorded can be posted or discarded

#### Scenario: A long text

- **WHEN** a person types past the length a text moment allows
- **THEN** the text stops growing, and no counter or number is shown

#### Scenario: Leaving without posting

- **WHEN** a person leaves a composer without posting
- **THEN** nothing they composed is kept for later

### Requirement: A song comes from the music apps people already use

A person SHALL be able to start a music moment from any music app's own share action, or by pasting a link to a song. The phone SHALL look the song up through a song-link lookup service, which returns its title, artist, artwork and its links on other music services. The lookup SHALL run from the phone and never through a Nah? server, and SHALL send the service nothing but the song's link. The artwork SHALL be fetched once on the posting phone and sealed with the moment, and the title, artist and links SHALL travel sealed inside it (ADR-0012, ADR-0016). When the lookup cannot be reached, the person SHALL be able to type the title and artist and post with the link they shared (ruled 2026-09-15).

#### Scenario: Sharing from a music app

- **WHEN** a person taps Share on a song in their music app and chooses Nah?
- **THEN** the music composer opens with the song's title, artist and artwork filled in

#### Scenario: Pasting a link

- **WHEN** a person pastes a link to a song into the music composer
- **THEN** the song's title, artist and artwork are filled in the same way

#### Scenario: The lookup is unavailable

- **WHEN** the song-link lookup cannot be reached
- **THEN** the person can type the title and artist, and post the song with the link they shared

#### Scenario: What the lookup learns

- **WHEN** a song is looked up
- **THEN** the service receives the song's link and no name, identifier or circle of the person, and no Nah? server receives anything about the song

### Requirement: An optional short line on photo, voice and music

The composer for a photo, voice or music moment SHALL offer one optional short line of at most 140 characters, with no counter, which travels inside the moment. A text moment SHALL NOT offer one (ruled 2026-09-15).

#### Scenario: A photo with a line

- **WHEN** a person adds a short line to a photo and posts it
- **THEN** the posted moment carries the photo and that line

#### Scenario: A text moment

- **WHEN** a person composes a text moment
- **THEN** no separate short line is offered

#### Scenario: A long line

- **WHEN** a person types past 140 characters in the short line
- **THEN** the line stops growing, and no counter or number is shown

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

A photo SHALL be resized and re-encoded, and its thumbnail made, on the phone (CDI-1843). A voice recording SHALL be encoded on the phone (CDI-1845). Every moment and every piece of media, including a song's artwork, SHALL be sealed on the phone before it is uploaded (ADR-0012), and SHALL travel in the versioned envelope with a plain-text fallback written by the posting app (ADR-0016).

#### Scenario: Posting a photo

- **WHEN** a person posts a photo taken at the camera's full resolution
- **THEN** what leaves the phone is a sealed, resized photo and a sealed thumbnail

#### Scenario: A call during a recording

- **WHEN** a call arrives while a person is recording a voice moment
- **THEN** recording stops cleanly, and the person can post or discard what was recorded once they return

### Requirement: No editing, and deleting leaves a marker

A posted moment SHALL NOT be edited (ADR-0008). The person who posted a moment SHALL be able to delete it. The server SHALL then discard its content and media and keep only who posted it and when, and wherever the moment appeared SHALL show that it was deleted (ruled 2026-09-15). Phones that already hold the moment SHALL be told at once to drop it and its media, and a phone that cannot be reached SHALL drop them at its next load.

#### Scenario: Deleting

- **WHEN** Maya deletes a moment she posted
- **THEN** everyone whose feed or page held it sees, from their next load, that Maya's moment was deleted, at the time it was posted, and nothing of what it held

#### Scenario: A phone that already showed it

- **WHEN** Maya deletes a moment that Sam's phone had already loaded
- **THEN** Sam's phone is told at once to drop the moment and its media, or drops them at its next load if it could not be reached

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
