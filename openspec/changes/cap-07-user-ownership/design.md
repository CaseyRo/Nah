# User ownership design

## Context

Each person is one SQLite file on a server (ADR-0011, amended), which makes export, deletion and moving a matter of file operations rather than queries. The server cannot read a sealed moment (ADR-0012), so anything that needs a moment's content, an export or a report, has to be assembled on a phone. Nah? keeps no email address or phone number for anyone. Nothing of this capability exists yet.

## Goals / Non-Goals

**Goals**

- Nothing a person posted outlives their wish to leave.
- Protection from someone in your circle that works without anyone reading anything.
- A statement of what the server sees that a person can check against the product.

**Non-Goals**

- The personal archive and resurfacing (`cap-10-archive`, CDI-1857).
- Second devices and recovery (CDI-1865).
- Moving to another server, which the directory already allows (cap-02).
- Building anything. This round is specification only.

## Decisions

### Leaving hides at once, and deletes after 30 days

Ruled 2026-09-15, after checking what is required. Apple requires deletion to start inside the app, and to cover the account, personal data and content shared with others. Offering only to deactivate is not enough, but a deletion that takes time is acceptable when the person is told how long, and given confirmation ([Apple](https://developer.apple.com/support/offering-account-deletion-in-your-app/)). Google Play requires an in-app path and a web link, and allows retaining data only for reasons like security or legal compliance, disclosed ([Google Play](https://support.google.com/googleplay/android-developer/answer/13327111)). The GDPR requires erasure without undue delay, acted on within one month ([article 12](https://gdpr-info.eu/art-12-gdpr/), [article 17](https://gdpr-info.eu/art-17-gdpr/)), and none of its exceptions covers a network wanting to keep someone's moments.

A 30-day wait fits these rules if everything is hidden at once, the person is told the date, and deletion then happens by itself. This is a reading of the rules, not legal advice, and CDI-1867's compliance pass should confirm it.

Rejected: keeping a leaver's moments for the circle, which the rules above do not allow; deactivation on its own, which Apple calls insufficient; and deleting at once, which leaves no way back for someone who left in a bad moment.

### Leaving is quiet

Ruled 2026-09-15. A person who leaves disappears: no marker where their moments stood, and nobody told. Deleting a single moment still leaves a marker (cap-03), because that is a correction people may have seen happen. Leaving is a person who wants to be forgotten.

### Ending a connection takes what it shared

Ruled 2026-09-15, amending cap-02. When two people's connection ends, each person's moments, earlier ones included, leave the other's feed and page. Blocking ends a connection the same way, and also refuses to let it be made again.

### Your own moments only, unsealed on your phone

An export holds what the person posted and nothing anyone else did (`product-principles`, CDI-1857). The server cannot unseal anything, so the phone builds the export and hands it to the person.

### A report is the one way a moment leaves its circle

Nobody can inspect a sealed moment, so a report carries the reporter's own copy, sent only when they confirm (CDI-1867). `product-principles` names this as the one exception to moments staying in their circle.

### The statement names who runs every outside lookup

The platforms' push services, the operating system's place lookup, and the song-link lookup, which is Odesli, owned by Linktree (cap-03). Until CDI-1863 seals moments, the statement also says the server could read them.

## Risks / Trade-offs

- **Nah? has no way to reach a person who deleted the app.** The deletion date is shown when they leave, and opening the app afterwards says the account was deleted. Apple asks for a confirmation when deletion completes, and this is the nearest Nah? can offer without collecting contact details.
- **Backups outlive the file.** A restored backup must not bring a deleted person back, and backup retention has to fit what the statement promises (CDI-1826, the restore drill).
- **A report cannot be checked against the original,** because the operator holds no key. A forged report is possible, so cutting someone off needs judgement, not automation.
- **The site's leaving page cannot prove who is asking** when the only key is on a lost phone.

## Open Questions

- Whether a hidden person keeps their places in other people's circles during the 30 days, and what happens on return if one of those circles has filled up.
- How the leaving page on the site confirms who is asking.
- How long backups may keep a deleted person.
- Whether an export includes the reactions a person received.
- What a blocked person is told when a connection is refused.
