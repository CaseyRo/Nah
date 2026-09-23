# User ownership

Linear: [CDI-1886](https://linear.app/cdit/issue/CDI-1886)

## Why

Owning your part of Nah? comes down to a few practical things:

- taking your moments with you
- leaving for good
- protecting yourself from someone
- knowing what the server can see

The January version promised export, deletion "within a reasonable timeframe", no tracking, and a privacy policy admitting that admins could read the database. ADR-0012 removed that admission. Both app stores now require in-app account deletion. Apple's guideline 1.2 requires reporting and blocking even in a private network (CDI-1867). This restatement specifies all of it.

## What Changes

- **export**: a person's own moments and media, unsealed on their phone, as plain files. Nothing anyone else posted.
- **leaving**: from the app, or through a page on the site. Everything is hidden at once, and deleted after 30 days unless the person comes back. Nobody is told, and no marker is left.
- **blocking-and-reporting**:
  - blocking ends a connection and refuses any reconnection until it is lifted
  - a report carries the reporter's own copy of a moment, sent only when they confirm, to a published abuse contact
  - the operator can cut a person off from the server and from push
- **privacy-statement**: a plain statement of what a server can and cannot see, naming every outside service, and saying so while moments are still unsealed. It also carries the sentence about losing every phone that holds your key.
- **Elsewhere**:
  - `product-principles` names a report as the one way a moment leaves its circle
  - in CDI-1883, ending a connection now takes both people's moments with it, and a block stops the two connecting again
- **New**: `cdi-1857-archive`, a scaffold for the personal archive and resurfacing (CDI-1857, M6).
- **Removed from the January version**: "admins can technically access the database", messages and reactions received in the export, and deletion "within a reasonable timeframe".

## Capabilities

### New Capabilities

- `export`: taking your own moments out of Nah?.
- `leaving`: leaving Nah?, the 30 days, and deletion.
- `blocking-and-reporting`: blocking a person, reporting a moment, the abuse contact, and cutting someone off.
- `privacy-statement`: what Nah? says about what it can see, and the cost of holding your own key.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- `apps/server`: a leaving state with a scheduled deletion, blocks, receiving reports, and the operator's cut-off; the directory entry's removal waits for M7.
- `apps/mobile`: export, leaving and returning, blocking, reporting, and the statement.
- The site: the privacy statement, the abuse contact, and the page for leaving.
- Linear: CDI-1867 covers reporting and blocking. Export, leaving and the statement have no issue yet. CDI-1867 also asks for a compliance pass before anyone outside Casey's own circle joins.
