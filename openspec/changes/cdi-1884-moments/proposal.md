# Moments

Linear: [CDI-1884](https://linear.app/cdit/issue/CDI-1884)

## Why

A moment is the only thing anyone posts in Nah?. The January version of this change listed photo and video, text, music, location, sleep and wake, tagging who you were with, and a journal on every profile, all as followers-only posts on Mastodon. Since then:

- ADR-0006 cut the types to text, voice and photo.
- ADR-0008 made moments immutable, with no drafts.
- ADR-0012 sealed them on the device.
- ADR-0016 fixed the envelope they travel in.

On 2026-09-15 music came back as a fourth type. On 2026-09-16 the feed design moved playback into focus, and on 2026-09-23 a text became able to carry a full written post with a title. M1 posts and reads text moments. This restatement specifies all four types, what a moment may carry, how it is deleted, and how the feed and a person's page show it, so M3 builds each type once.

## What Changes

- **moments**:
  - four types: text, voice, photo and music
  - a text can be a full written post of up to 10,000 characters, with an optional title (ruled 2026-09-23, CDI-1899)
  - a music moment is a song shared from any music app, or pasted as a link, and looked up on the phone for its title, artist, artwork and links on other music services
  - an optional short line on a photo, voice or music moment, and an optional place on any moment; the place comes from the phone's own place lookup or is typed by hand, and Nah? keeps no location
  - media and artwork are prepared and sealed on the phone
  - no drafts and no editing; the person who posted a moment can delete it, and it is replaced by a plain marker
  - a moment composed without signal waits visibly, and is never lost
- **feed**: each moment on the two-lane ground of `DESIGN.md` with no box around it; photo, voice, music and long text open on their own in focus, where voice and music play (ruled 2026-09-16, CDI-1900); the timeline clock; a sentence in place of any moment the app cannot show; a marker where a moment was deleted; and an end to the page. A song plays a sample in focus and opens in the reader's own service, and nothing is fetched from music services before that tap.
- **person-page**: a page of one person's moments, which that person and everyone in their circle can open, and nobody else.
- **Decisions amended (ruled 2026-09-15)**: ADR-0006, ADR-0008 and ADR-0016 each carry a dated note covering music, the short line, the place, deletion and song links.
- **Elsewhere**: CDI-1882's + offers four items, and its no-reporting rule names the place lookup and the song-link lookup as exceptions. nah-vision gains the principle that Nah? keeps no record of where anyone is. cap-06 is removed, because music was the part of it worth keeping.
- **Removed from the January version**: video, location as a moment type, sleep and wake, tagging who you were with, followers-only visibility, and map previews.

## Capabilities

### New Capabilities

- `moments`: what a moment is, what it may carry, and its life from composing to deleting.
- `feed`: how the feed shows moments.
- `person-page`: one person's moments on a page of their own.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- `apps/mobile`: voice, photo and music composers, receiving a song from other apps' share actions, the song-link lookup, the short line and the place, the queue, the four cards, the clock, the deleted marker, and the person page.
- `apps/server`: an upload path for media blobs (CDI-1824 still has to measure it), a route that deletes a moment and leaves a marker, and a feed for one person. The server never sees a song lookup.
- An outside service: the song-link lookup, called from phones (see design.md).
- Linear: CDI-1843 to CDI-1847 cover photo, voice and the queue, CDI-1899 the long text and CDI-1900 focus. Music, deleting a moment, the short line, the place and the person page have no issue yet.
- `docs/decisions`: notes on ADR-0006, ADR-0008 and ADR-0016, and a second note on ADR-0006 for the long text, so the codebase wiki needs recompiling.
