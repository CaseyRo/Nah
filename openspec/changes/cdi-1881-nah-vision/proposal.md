# Nah? vision

Linear: [CDI-1881](https://linear.app/cdit/issue/CDI-1881)

## Why

Nah? is a private social network for the people one person is actually close to: up to 150 of them, joined by invitation, reading one chronological feed that ends. The January version of this change described a Path successor built on a Mastodon fork, with friend requests, an inner circle, reactions, comments, view receipts and a donation model. Most of that has since been decided differently (ADR-0004 to ADR-0017), and its nine specs duplicated capabilities that CDI-1882 to CDI-1888 own. What holds across every capability is a short list of principles. This change now specifies only those, so each capability can be checked against them.

**Viral? Nah. Vital.**

## What Changes

- **product-principles**: the rules every capability keeps:
  - one circle of up to 150 per person, drawn around that person alone, with no groups and no audience picker
  - connection is mutual, and physical first
  - joining is by invitation, and nothing is public or discoverable
  - the feed is chronological and unranked
  - no numbers about people
  - the server keeps moments it cannot read
  - moments never leave the circle they were posted to
  - no chat, one-to-one or in groups
  - no AI in the product
  - nothing is sold and nothing advertises
  - notifications default to the most respectful option
  - the code stays open under AGPL-3.0-or-later, and every fork keeps the attribution
- **Moved out**: the January specs for friend circles, moments, reactions, comments, messaging, ambient presence, user ownership and community funding. Each now sits in its capability's folder as `from-nah-vision.md`, source material to restate. Comments get a change of their own, `cdi-1888-comments`. The January core-identity spec is removed, because CDI-1882 restated it on 2026-09-15.
- **Removed**, each superseded by a decision record listed in design.md:
  - the Mastodon fork and ActivityPub
  - friend requests, the inner circle and followers-only visibility
  - email and password accounts
  - Sentry and Prometheus
  - S3 with a CDN
  - the `nah_ui` and `nah_api` packages

## Capabilities

### New Capabilities

- `product-principles`: the guarantees Nah? makes to the people in it, which every other capability has to keep.

### Modified Capabilities

None: `openspec/specs/` holds no specs yet.

## Impact

- Every capability change, CDI-1882 to CDI-1888, is checked against `product-principles` when it is restated.
- The architecture this change once described lives in the decision records and the codebase wiki; design.md points at them.
- No code changes. As far as the M1 server and app exist, the principles already hold in them. The exceptions are named in the spec with their milestone: sealing moments waits on CDI-1863, the touch on CDI-1840, and the digest on CDI-1850.
