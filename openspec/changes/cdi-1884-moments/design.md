# Moments design

## Context

M1 sends text moments as a version 1 envelope behind a seal byte of 0, and reads every envelope as hostile input: bounded sizes, and a fallback sentence for anything it does not know (`apps/mobile/lib/moment.dart`). The server orders moments by when it received them and knows nothing about what they contain. Photo, voice, music, the queue and media uploads are all still to come.

## Goals / Non-Goals

**Goals**

- Four types that each feel finished.
- A place, a line and a song that belong to the person who added them, and to nobody else.
- Deletion that is honest about itself.

**Non-Goals**

- Video, location as a type of its own, and status (ADR-0006).
- Live presence of any kind, including what someone is playing right now. cap-06 is removed.
- Reactions and comments (CDI-1885, CDI-1888).
- The archive and resurfacing (CDI-1857, M6).
- Building anything. This round is specification only.

## Decisions

### Music, shared from any music app and looked up on the phone

Ruled 2026-09-15. Music is the part of ambient presence worth keeping, as a deliberate share rather than a live status, so it is a moment. A song arrives through the music app's own share action or a pasted link, which works for every service on both platforms without an SDK or a special permission.

The phone sends that link to Odesli (song.link), `GET https://api.song.link/v1-alpha.1/links`, which returns the title, artist, artwork and the song's links on other services. The moment keeps the title, artist and per-service links, sealed. It does not keep the song.link page, which would send every reader through a third party. The artwork is fetched once by the posting phone and sealed as an attachment, so no reader contacts a music service until they tap. ADR-0016 carries a note for the links.

Rejected: the music services' own SDKs, which are flaky and different for every service; reading what is playing on the phone, which needs notification access on Android and sees only Apple Music on iOS; and routing the lookup through a Nah? server, which would let the server learn what a moment is about.

### A text can be a full post, with a title

Ruled 2026-09-23 by Casey with Emma, with Hyves coming back. People were reflective on Path, and a line is not enough room for that. A text moment runs from one line to 10,000 characters, with an optional title of up to 100, both without counters. It stays one type, text, so the + keeps four items and ADR-0006 gets a note rather than a fifth type. On the feed a long text shows its title and opening lines, and reading on opens it in focus.

10,000 characters keep the envelope inside the 64 KiB an app already accepts, so an older app shows the text or its fallback rather than calling the moment unreadable. The title is a new field, which older apps ignore (ADR-0016). No photos inside a text: a photo is its own moment.

### Nothing plays on the feed

Ruled 2026-09-16 in `DESIGN.md`. Where other feeds lean into the next thing, Nah? spends time on one: a photo, voice or music moment opens in focus on the bare ground, and voice and music play only there. A song plays a short sample from one of its links, fetched after the tap, and one action opens it in the reader's own service, chosen once and remembered. This replaces the earlier "plays in place" and "choose a service on every tap" (CDI-1900).

### A short line on photo, voice and music, a place on anything

Ruled 2026-09-15, recorded as notes on ADR-0006 and ADR-0008. A photo, a voice clip or a song often needs a few words to land, and 140 characters is room for them, with no counter. A text moment already is words, and a line above it would be the headline ADR-0006 rules out. A place fits any moment. Both are optional, and both live inside the sealed envelope, so the server learns neither.

### Where a place comes from, and what Nah? keeps

Ruled 2026-09-15. When the person asks to add their location, the app asks the phone for it once. It passes the location to the operating system's place lookup, which returns a name the person can edit before posting. The person can always type a place instead, and typing is the way in when location access is refused or the lookup fails.

Nah? keeps only the words: no coordinates, no history, and nothing on any server that is not sealed inside a moment. A place is metadata for the person who added it. If the archive uses places later, it uses them for that person, never for the project.

Rejected: a map pin, because tiles tell a tile server where people are looking; and a bundled list of towns, which is private but too coarse for what people want to say.

### Deleting leaves a marker

Ruled 2026-09-15. The person who posted a moment can delete it. The server discards its content and media and keeps only who posted it and when. Every feed and page shows "This moment was deleted" where it stood, so nobody who saw it wonders whether it vanished or whether they were removed. Editing stays out (ADR-0008). Phones that already hold the moment are told at once, through the push relay (M4), to drop it and its media, and a phone that is off or out of reach drops them at its next load. Background pushes are best effort on both platforms, which is why the next load stays the backstop.

### One person's page is open to their circle

Ruled 2026-09-15. Tapping a name opens that person's moments, and "you" opens your own. The page is visible to the person and their circle and never outside it (`product-principles`). Its top is an avatar and a name with no cover photo: Path had one, and it would be one more thing to curate. Under the name come the gender, pronouns and orientation the person chose to show, and their first moment, their answer on arriving, quoted. Then their moments in the same two lanes as the feed, with no cards (ruled 2026-09-23), one page at a time, with older ones loading only when asked; the first moment appears again at its date. Nothing on it counts.

### What cannot be shown keeps its place

A moment of a type the app does not know shows the fallback sentence its author's app wrote (ADR-0016). A malformed moment says it could not be read. Neither disappears, because a gap in a chronological feed misstates what happened.

### Media is prepared, and sealed, on the phone

A photo is resized and re-encoded on the phone, and its thumbnail is made there too, about 550 KB for the pair (CDI-1843). A voice clip is recorded for up to a minute and encoded to Opus on the phone, about 150 KB (CDI-1845). A song's artwork is fetched and sealed on the phone. All of it is sealed before it leaves (ADR-0012), which keeps the server free of media tooling.

## Risks / Trade-offs

- **The song-link lookup belongs to Linktree,** which bought Songlink/Odesli in 2021. It learns which song a phone asked about, from that phone's address, though never who the person is or who the moment is for. CDI-1882 names it as an exception, alongside the operating system's place lookup.
- **The lookup's API is still `v1-alpha.1`.** It allows 10 requests a minute per caller without a key, and 60 with one. Calls come from phones, so each phone has its own allowance; a key shipped inside the app would be shared by everyone, and anyone could extract it.
- **Fetching a song's artwork contacts whoever hosts it,** once, from the posting phone. **A song's sample contacts that music service** from the reader's phone, only after the reader taps; CDI-1882 and the privacy statement name it.
- **If the lookup disappears,** a person can still post a song by typing its title and artist, with the link they shared.
- **The operating system's place lookup is a request outside Nah?.** Apple or Google see a coordinate when a person asks for their location, and typing is always offered instead.
- **A deleted moment's marker keeps metadata.** The server keeps that someone posted at a given time, which it already knew (ADR-0012's stated exception).
- **A person page is a surface people can tend.** It is limited to the circle and shows no counts. If people start curating it, ADR-0017's warning about performing applies.
- **Voice is the most expensive composer** (ADR-0006). CDI-1845 names the hard parts: recording without jank on iOS, and a call interrupting a recording.
- **The media upload path is unmeasured** (CDI-1824).

## Open Questions

None left open: the last were settled on 2026-09-15.
