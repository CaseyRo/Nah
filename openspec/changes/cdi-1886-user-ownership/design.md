# User ownership design

## Context

Each person is one SQLite file on a server (ADR-0011, amended), which makes export, deletion and moving a matter of file operations rather than queries. The server cannot read a sealed moment (ADR-0012), so anything that needs a moment's content, an export or a report, has to be assembled on a phone. Nah? keeps no email address or phone number for anyone. Nothing of this capability exists yet.

## Goals / Non-Goals

**Goals**

- Nothing a person posted outlives their wish to leave.
- Protection from someone in your circle that works without anyone reading anything.
- A statement of what the server sees that a person can check against the product.

**Non-Goals**

- The personal archive and resurfacing (`cdi-1857-archive`, CDI-1857).
- Second devices and recovery (CDI-1865).
- Moving to another server, which the directory already allows (CDI-1883).
- Building anything. This round is specification only.

## Decisions

### Leaving hides at once, and deletes after 30 days

Ruled 2026-09-15, after checking what is required. Apple requires deletion to start inside the app, and to cover the account, personal data and content shared with others. Offering only to deactivate is not enough, but a deletion that takes time is acceptable when the person is told how long, and given confirmation ([Apple](https://developer.apple.com/support/offering-account-deletion-in-your-app/)). Google Play requires an in-app path and a web link, and allows retaining data only for reasons like security or legal compliance, disclosed ([Google Play](https://support.google.com/googleplay/android-developer/answer/13327111)). The GDPR requires erasure without undue delay, acted on within one month ([article 12](https://gdpr-info.eu/art-12-gdpr/), [article 17](https://gdpr-info.eu/art-17-gdpr/)), and none of its exceptions covers a network wanting to keep someone's moments.

A 30-day wait fits these rules if everything is hidden at once, the person is told the date, and deletion then happens by itself. This is a reading of the rules, not legal advice, and CDI-1867's compliance pass should confirm it.

Rejected: keeping a leaver's moments for the circle, which the rules above do not allow; deactivation on its own, which Apple calls insufficient; and deleting at once, which leaves no way back for someone who left in a bad moment.

### Leaving is quiet

Ruled 2026-09-15. A person who leaves disappears: no marker where their moments stood, and nobody told. Deleting a single moment still leaves a marker (CDI-1884), because that is a correction people may have seen happen. Leaving is a person who wants to be forgotten.

### Ending a connection takes what it shared

Ruled 2026-09-15, amending CDI-1883. When two people's connection ends, each person's moments, earlier ones included, leave the other's feed and page. Blocking ends a connection the same way, and also refuses to let it be made again.

### Your own moments only, unsealed on your phone

An export holds what the person posted and nothing anyone else did (`product-principles`, CDI-1857). The server cannot unseal anything, so the phone builds the export and hands it to the person.

### A report is the one way a moment leaves its circle

Nobody can inspect a sealed moment, so a report carries the reporter's own copy, sent only when they confirm (CDI-1867). `product-principles` names this as the one exception to moments staying in their circle.

### The statement names who runs every outside lookup

The platforms' push services, the operating system's place lookup, and the song-link lookup, which is Odesli, owned by Linktree (CDI-1884). Until CDI-1863 seals moments, the statement also says the server could read them.

### Places are held for the 30 days

Ruled 2026-09-15. A person who left keeps their places in other people's circles until their data is deleted, so returning restores every connection intact. The cost is that someone's circle can stay full for up to a month because of a place they cannot see.

### Backups age out within 30 days

Ruled 2026-09-15. Backups older than 30 days are discarded, so nothing of a deleted person survives more than 30 days past deletion, or 60 days after they left. The privacy statement says so.

### Exports leave reactions out

Ruled 2026-09-15. A reaction is someone else's act, so an export holds only what the person posted.

### A blocked person is told plainly

Ruled 2026-09-15. When a blocked person tries to connect again, they are told that they have been blocked, so they know where they stand.

### Leaving from the web means recovering first

Ruled 2026-09-15. Google Play wants a web path to deletion, but the key lives on the phone. The site explains leaving from the app, and someone who lost their phone recovers through people in their circle first (CDI-1865). Anyone who cannot writes to Nah?'s published contact, and the operator confirms with someone in their circle before the 30 days start.

## Risks / Trade-offs

- **Nah? has no way to reach a person who deleted the app.** The deletion date is shown when they leave, and opening the app afterwards says the account was deleted. Apple asks for a confirmation when deletion completes, and this is the nearest Nah? can offer without collecting contact details.
- **Backups outlive the file by up to 30 days.** Restoring a backup taken before a deletion would bring that person back, and the restore drill has to account for it (CDI-1826).
- **Telling a blocked person plainly can provoke them.** Reporting and the operator's cut-off exist for that case.
- **A report cannot be checked against the original,** because the operator holds no key. A forged report is possible, so cutting someone off needs judgement, not automation.
- **The site's leaving page cannot prove who is asking** when the only key is on a lost phone.

## Open Questions

None left open: the last were settled on 2026-09-15.
