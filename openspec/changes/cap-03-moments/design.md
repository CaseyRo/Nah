# Moments design

## Context

M1 sends text moments as a version 1 envelope behind a seal byte of 0, and reads every envelope as hostile input: bounded sizes, and a fallback sentence for anything it does not know (`apps/mobile/lib/moment.dart`). The server orders moments by when it received them and knows nothing about what they contain. Photo, voice, the queue and media uploads are M3.

## Goals / Non-Goals

**Goals**

- Three types that each feel finished.
- A place and a line that belong to the person who added them, and to nobody else.
- Deletion that is honest about itself.

**Non-Goals**

- Video, music, location as a type of its own, and status (ADR-0006).
- Reactions and comments (cap-04, cap-09).
- The archive and resurfacing (CDI-1857, M6).
- Building anything. This round is specification only.

## Decisions

### A short line on photo and voice, a place on anything

Ruled 2026-09-15, recorded as notes on ADR-0006 and ADR-0008. A photo or a voice clip often needs a few words to land. A text moment already is words, and a line above it would be the headline ADR-0006 rules out. A place fits any moment. Both are optional, and both live inside the sealed envelope, so the server learns neither.

### Where a place comes from, and what Nah? keeps

Ruled 2026-09-15. When the person asks to add their location, the app asks the phone for it once. It passes the location to the operating system's place lookup, which returns a name the person can edit before posting. The person can always type a place instead, and typing is the way in when location access is refused or the lookup fails.

Nah? keeps only the words: no coordinates, no history, and nothing on any server that is not sealed inside a moment. A place is metadata for the person who added it. If the archive uses places later, it uses them for that person, never for the project.

Rejected: a map pin, because tiles tell a tile server where people are looking; and a bundled list of towns, which is private but too coarse for what people want to say.

### Deleting leaves a marker

Ruled 2026-09-15. The person who posted a moment can delete it. The server discards its content and media and keeps only who posted it and when. Every feed and page shows "This moment was deleted" where it stood, so nobody who saw it wonders whether it vanished or whether they were removed. Editing stays out (ADR-0008).

### One person's page is open to their circle

Ruled 2026-09-15. Tapping a name opens that person's moments, and "you" opens your own. The page is visible to the person and their circle and never outside it (`product-principles`). It shows a name and moments, and nothing that counts.

### What cannot be shown keeps its place

A moment of a type the app does not know shows the fallback sentence its author's app wrote (ADR-0016). A malformed moment says it could not be read. Neither disappears, because a gap in a chronological feed misstates what happened.

### Media is prepared, and sealed, on the phone

A photo is resized and re-encoded on the phone, and its thumbnail is made there too, about 550 KB for the pair (CDI-1843). A voice clip is recorded for up to a minute and encoded to Opus on the phone, about 150 KB (CDI-1845). Both are sealed before they leave (ADR-0012), which keeps the server free of media tooling.

## Risks / Trade-offs

- **The operating system's place lookup is a request outside Nah?.** Apple or Google see a coordinate when a person asks for their location. cap-01's no-reporting requirement names it as an exception, and typing is always offered instead.
- **A deleted moment's marker keeps metadata.** The server keeps that someone posted at a given time, which it already knew (ADR-0012's stated exception).
- **A person page is a surface people can tend.** It is limited to the circle and shows no counts. If people start curating it, ADR-0017's warning about performing applies.
- **Voice is the most expensive composer** (ADR-0006). CDI-1845 names the hard parts: recording without jank on iOS, and a call interrupting a recording.
- **The media upload path is unmeasured** (CDI-1824).

## Open Questions

- The short line's length limit.
- How far back a person page goes, and how older moments load without becoming an infinite scroll.
- Whether a deleted moment's media is cleared from phones that cached it, beyond their next load.
