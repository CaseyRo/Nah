# Moments

## Why

A moment is the only thing anyone posts in Nah?. The January version of this change listed photo and video, text, music, location, sleep and wake, tagging who you were with, and a journal on every profile, all as followers-only posts on Mastodon. Since then:

- ADR-0006 cut the types to text, voice and photo.
- ADR-0008 made moments immutable, with no drafts.
- ADR-0012 sealed them on the device.
- ADR-0016 fixed the envelope they travel in.

M1 posts and reads text moments. This restatement specifies all three types, what a moment may carry, how it is deleted, and how the feed and a person's page show it, so M3 builds photo and voice once.

## What Changes

- **moments**:
  - three types: text, voice and photo
  - an optional short line on a photo or voice moment, and an optional place on any moment; the place comes from the phone's own place lookup or is typed by hand, and Nah? keeps no location
  - media is resized, encoded and sealed on the phone
  - no drafts and no editing; the person who posted a moment can delete it, and it is replaced by a plain marker
  - a moment composed without signal waits visibly, and is never lost
- **feed**: a card per type, the timeline clock, a sentence in place of any moment the app cannot show, a marker where a moment was deleted, and an end to the page.
- **person-page**: a page of one person's moments, which that person and everyone in their circle can open, and nobody else.
- **Decisions amended (ruled 2026-09-15)**: ADR-0006 and ADR-0008 each carry a dated note covering the short line, the place and deletion.
- **Elsewhere**: cap-01's no-reporting rule names the phone's place lookup as an exception, and nah-vision gains the principle that Nah? keeps no record of where anyone is.
- **Removed from the January version**: video, music, location as a moment type, sleep and wake, tagging who you were with, followers-only visibility, and map previews.

## Capabilities

### New Capabilities

- `moments`: what a moment is, what it may carry, and its life from composing to deleting.
- `feed`: how the feed shows moments.
- `person-page`: one person's moments on a page of their own.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- `apps/mobile`: voice and photo composers, the short line and the place, the queue, the three cards, the clock, the deleted marker, and the person page.
- `apps/server`: an upload path for media blobs (CDI-1824 still has to measure it), a route that deletes a moment and leaves a marker, and a feed for one person.
- Linear: CDI-1843 to CDI-1847 cover most of it. Deleting a moment, the short line, the place and the person page have no issue yet.
- `docs/decisions`: notes on ADR-0006 and ADR-0008, so the codebase wiki needs recompiling.
