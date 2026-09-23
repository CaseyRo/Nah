## Purpose

How the feed shows moments: each on the two-lane ground `DESIGN.md` describes, a moment opened on its own in focus, the timeline clock while scrolling, an honest sentence for anything that cannot be shown, and a plain end.

## ADDED Requirements

### Requirement: Each moment in the two lanes

Each moment SHALL be shown in the two-lane feed of `DESIGN.md`: its time in the time lane, and the poster's name (CDI-1895), its body and its foot in the content column, with no box around it. A text moment SHALL be shown as words; a long one SHALL show its title, if it has one, and its opening lines with a way to read on. A photo SHALL be shown in the feed treatment `DESIGN.md` gives it. A voice moment SHALL show its waveform and length, and SHALL NOT play on the feed. A music moment SHALL show its sealed artwork, title and artist. A short line and a place SHALL appear as words when the moment has them, and no map SHALL be drawn (ruled 2026-09-15 and 2026-09-16).

#### Scenario: A photo

- **WHEN** a photo moment is in the feed
- **THEN** it appears under the poster's name, with its time in the lane, and no box around it

#### Scenario: A voice moment on the feed

- **WHEN** a voice moment is in the feed and nobody taps it
- **THEN** it shows its waveform and length, and nothing plays

#### Scenario: A long text

- **WHEN** a text moment runs longer than its opening lines
- **THEN** the feed shows its title, if any, and its opening lines, with a way to read on

#### Scenario: A line and a place

- **WHEN** a voice moment carries a short line and a place
- **THEN** both appear under it as words

### Requirement: A moment on its own, in focus

Tapping a photo, voice or music moment, or reading on in a long text, SHALL open that moment in focus: on the bare ground, with nothing else from the feed on the screen, as `DESIGN.md` describes (ruled 2026-09-16, CDI-1900). A photo SHALL be shown alone. A voice moment SHALL start playing on entering focus and stop on leaving it. A long text SHALL be shown in full, with its title. A short text SHALL NOT open in focus, because the feed already shows all of it. Leaving focus SHALL return the feed exactly where it was.

#### Scenario: Opening a photo

- **WHEN** a person taps a photo in the feed
- **THEN** the photo opens on its own on the bare ground, with no name, line or control beside it

#### Scenario: Playing a voice moment

- **WHEN** a person taps a voice moment
- **THEN** it opens in focus and plays, and stops when they leave focus

#### Scenario: Reading a long text

- **WHEN** a person reads on in a long text
- **THEN** its title and whole text open in focus, and leaving returns them to the same place in the feed

### Requirement: A song plays a sample, then opens where the reader listens

Tapping a music moment SHALL open it in focus, where a short sample of the song plays, fetched from one of the song's sealed links only after the tap. Focus SHALL offer one action that opens the song in the reader's own music service, which the reader chooses once and the app remembers. Before the tap, the app SHALL fetch, preview or contact nothing at any music service (ADR-0016, ruled 2026-09-16).

#### Scenario: Opening a song

- **WHEN** a reader taps a music moment for the first time
- **THEN** a sample plays in focus, and the one action asks which service they use, opens the song there, and remembers the choice

#### Scenario: Scrolling past a song

- **WHEN** a music moment is shown in the feed and nobody taps it
- **THEN** no request goes to any music service, and its artwork comes from the moment itself

### Requirement: The timeline clock

While the feed scrolls, a clock SHALL ride the time lane at the middle of the feed's visible area, its hands reading the time at that line between the moment above and the moment below, with the date beneath it, and SHALL fade once scrolling stops, as `DESIGN.md` describes (ruled 2026-09-16). With reduce motion on, there SHALL be no clock face: the lane label of the moment nearest the middle SHALL be emphasised, with the date under it, and nothing SHALL move.

#### Scenario: Scrolling back through a week

- **WHEN** the person scrolls from today's moments back to last Tuesday's
- **THEN** the clock's hands turn through the time between the moments passing it, and its date follows

#### Scenario: Scrolling stops

- **WHEN** the person stops scrolling
- **THEN** the clock fades after about two seconds

#### Scenario: Reduce motion

- **WHEN** reduce motion is on and the feed scrolls
- **THEN** no clock face appears, and the time and date of the moment nearest the middle stand out in the lane

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
