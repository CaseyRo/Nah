## Purpose

How the feed shows moments: a card for each type, the timeline clock while scrolling, an honest sentence for anything that cannot be shown, and a plain end.

## ADDED Requirements

### Requirement: A card for each type

Each moment SHALL be shown as a card that names who posted it and when, with the treatment `DESIGN.md` gives its type. A text moment SHALL be shown as padded words. A photo SHALL be shown edge to edge at its own aspect ratio. A voice moment SHALL play in place with a tap and show its waveform. A short line and a place SHALL appear on the card as words when the moment has them, and no map SHALL be drawn.

#### Scenario: A photo

- **WHEN** a photo moment is in the feed
- **THEN** it spans the card's width without being cropped, under the poster's name and the time

#### Scenario: A voice moment

- **WHEN** the person taps a voice moment
- **THEN** it plays in place, without leaving the feed

#### Scenario: A line and a place

- **WHEN** a voice moment carries a short line and a place
- **THEN** both appear on its card as words

### Requirement: The timeline clock

While the feed scrolls, a clock SHALL appear at the left edge showing the time and date of the moment nearest the middle of the screen, and SHALL fade once scrolling stops, as `DESIGN.md` describes. With reduce motion on, it SHALL show the time as static digits.

#### Scenario: Scrolling back through a week

- **WHEN** the person scrolls from today's moments back to last Tuesday's
- **THEN** the clock's hands and date follow the moments passing the middle of the screen

#### Scenario: Scrolling stops

- **WHEN** the person stops scrolling
- **THEN** the clock fades after about two seconds

#### Scenario: Reduce motion

- **WHEN** reduce motion is on and the feed scrolls
- **THEN** the clock shows the time as digits, with no moving hands

### Requirement: What cannot be shown keeps its place

Where a moment cannot be shown as itself, the feed SHALL keep it in its place in time and say so. A type the app does not know SHALL show the fallback sentence its author's app wrote (ADR-0016). A malformed moment SHALL say that it could not be read. A deleted moment SHALL say that it was deleted, with the poster's name and the time.

#### Scenario: A newer type

- **WHEN** a moment of a type this app does not know arrives
- **THEN** its author's fallback sentence appears where the moment stands, and nothing around it changes

#### Scenario: A malformed moment

- **WHEN** a moment's envelope is malformed or larger than the app accepts
- **THEN** the feed says that the moment could not be read, and the app keeps running

#### Scenario: A deleted moment

- **WHEN** a moment has been deleted by its poster
- **THEN** the feed shows, with the poster's name and the time, that the moment was deleted

### Requirement: The feed ends plainly

After the last moment on its page, the feed SHALL end with a short sentence saying the person is caught up.

#### Scenario: The last moment

- **WHEN** the person scrolls past the last moment on the page
- **THEN** a short sentence says they are caught up, and nothing follows it
