---
title: "ADR-0015 — No passwords, a key on the device"
permalink: /decisions/0015-no-passwords-a-key-on-the-device/
layout: page
status: Accepted
date: 2026-09-13
---

## Status

Accepted (2026-09-13). Follows [ADR-0011](0011-plain-go-and-a-database-per-circle.html), which removed PocketBase and with it the sign-in that came free. Sits alongside [ADR-0012](0012-encrypted-on-device.html), which already puts a key on the device.

Revised the same day, before anything was built on it. The first version kept sessions in memory and accepted that a deploy would log everyone out. That was the wrong default for a product that intends to deploy constantly, and the revision is recorded in the decision below rather than left in the history.

## Context

Dropping PocketBase took authentication with it, so sign-in is now ours to write. It is the first thing a person does and the thing they will be doing at eleven at night when it does not work, so it is a support decision as much as a security one.

The product has an unusual constraint that narrows the field before taste enters. Under ADR-0012 the device already holds a circle's content key, because that is how it reads a moment. Whatever sign-in we choose, a person who has lost every device has already lost the content. An authentication scheme with a recovery path stronger than the encryption would only be pretending.

There is also nothing to sign in *to* in the usual sense. There is no account, no profile and no cross-circle identity on the server. A person is a member of a circle, and membership is a row in that circle's own database file.

## Decision

**The app generates an Ed25519 keypair on first run. The public key is the identity, and signing a server-issued challenge is the sign-in.**

- Joining a circle registers the public key against that circle's membership. There is no separate registration step and no account.
- To get a session, the app asks the circle for a challenge, signs it, and exchanges the signature for a bearer token. `crypto/ed25519` is in the standard library, so this is verification in one line rather than a dependency.
- **A challenge, not a signed timestamp.** The server issues 32 random bytes, holds them briefly, and accepts each exactly once. Signing a client-chosen timestamp would be smaller, but a captured signature would then be replayable for as long as the window allowed, and the window is the only thing standing between an intercepted request and a session.
- **A session token carries its own claims and a MAC over them**, so it outlives the process that issued it. The token is an expiry, a public key and a circle id, with an HMAC-SHA256 over all three. Nothing about a session is stored anywhere.
- **The key that signs those tokens lives in the data directory**, generated on first run, beside the circle files rather than in the environment. A deploy cannot forget a file it never touches, a restored volume restores the sessions with it, and two instances sharing a volume during a rolling deploy agree without being told anything.
- **The identity key is not the content key.** The content key belongs to the circle and arrives in the invite. The identity key belongs to the device and never leaves it. They are generated and stored separately even though both live in the same secure enclave, because they have different lifetimes: the content key rotates when someone is removed, the identity key does not.

The first version of this decision kept sessions in a map and accepted that every deploy would log everyone out. The intent is to ship constantly and for nobody to be able to tell, which makes a per-deploy re-authentication for every member exactly the wrong default: it is invisible only for as long as the client's silent-retry path is perfect, and it turns one client bug into a login screen in front of somebody's family. Stateless tokens cost about forty lines and remove the failure rather than managing it.

## Considered alternatives

- **Email and password.** The familiar answer, and the reason it is rejected is not that passwords are unfashionable. It needs password hashing, an SMTP dependency we otherwise do not have, a reset flow, and a breach story. Password reset is the single largest support surface in anything aimed at families, and we would be adding it to a product that still cannot recover a lost circle key. The weakest credential would then be the one guarding an account that owns nothing.
- **Email magic links.** No password to leak, but it puts the credential in whoever still has access to a grandparent's mailbox, adds the same SMTP dependency, and makes sign-in depend on mail delivery — a category of failure the project would then be answering support mail about. It also fails the guardrail that joining never involves typing an address.
- **Passkeys / WebAuthn.** The right answer for a product with a web client, and genuinely better recovery through the platform's own sync. Rejected because it makes the platform an intermediary in an app whose whole argument is that there is no intermediary, and because the recovery it buys is recovery of the *account*, not of the content, which is the part people will actually lose.
- **A signed timestamp instead of a challenge.** Smaller, and what request-signing schemes generally do. Rejected above: replay within the window, for the sake of about twenty-five lines.
- **Sessions in a table in each circle's file.** Survives a restart too, and can revoke one session before it expires. Rejected because it buys revocation this product has no use for — a member is removed, not a session — while adding a table, a migration, a write on every sign-in, and rows nobody ever cleans up.
- **Sessions in memory, as this decision first said.** Rejected on the deploy story above.
- **Signing every request rather than holding a session.** Removes session state entirely and is tempting for that reason. Rejected because it puts a signature on the hot path of a feed poll, pushes replay protection into every endpoint instead of one, and makes the client harder to write on the only platform we have.

## Consequences

Positive:

- No password to leak, reset, hash or be breached, and no SMTP anywhere in the product.
- Sign-in and decryption fail and recover together. A person either has a working device or does not, and there is one story to tell rather than two that can disagree.
- Social recovery in M5 then covers identity and content at once: an existing member approves a new phone, and it gets both.
- Nothing is added to the dependency budget. Ed25519, HMAC and `crypto/rand` are standard library.
- A deploy is invisible. Sessions survive a restart, and survive two instances running at once, so nothing in the authentication design stands between this and continuous deployment.

Negative / acknowledged cost:

- **Lose every device and you are out** until someone vouches for you, which is the same trade ADR-0012 already makes for content. It is one trade, not two, but it is still the trade.
- **We own the crypto**, though only the easy part: generate, sign, verify. No key agreement, no ratchet.
- **No web client is possible** without solving key storage in a browser. There is no web client and none is planned, but this closes the door rather than leaving it ajar.
- **One session cannot be revoked before it expires.** Nothing is stored, so there is nothing to delete. What revokes access is removing the membership, which the server checks on every request — but until someone does that, a thirty-day token is thirty days of a stolen phone staying signed in.
- **The signing key is now a thing that can be lost.** Lose it and everyone signs in again, which is a recoverable annoyance. Leak it and anyone can mint a session for any public key, which is not. It sits at `0600` in the data directory, so it is as protected as the circle files beside it — and those are ciphertext, while it is not.
- A stolen unlocked phone is a valid member until that membership is removed. True of every scheme here, but worth writing down rather than discovering.

## What would prove this wrong

- Thirty days turns out to be far too long for a token nobody can revoke. The first stolen or sold phone will say. Shortening it is cheap; what it would really mean is that revocation belongs somewhere after all.
- Someone gets a new phone before M5 exists and there is no way to move them without the developer touching the database.
- A second device turns out to be the common case rather than the exception, which would make the "one device, one key" assumption underneath this the wrong shape rather than an early simplification.
- The in-memory challenge store becomes a correctness problem — a rolling deploy routes a device's challenge to one instance and its signature to another, and sign-in starts failing intermittently under exactly the deployment model this revision was written for. That would mean challenges have to be shared, or derived rather than stored.
