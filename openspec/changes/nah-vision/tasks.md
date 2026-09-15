# Nah? vision tasks

Specification only this round. The principles are checked here against what exists; the work for each capability lives in its own change.

## 1. Principles that already hold in M1

- [x] 1.1 The server enforces the limit of 150 (CDI-1830)
- [x] 1.2 Joining is by invitation only; the first person on a server uses `server invite`
- [x] 1.3 Moments are stored as opaque blobs behind a seal byte (CDI-1862)
- [x] 1.4 The feed is newest first, one page, unranked (CDI-1833)
- [x] 1.5 AGPL-3.0-or-later, with the attribution term in `NOTICE`

## 2. Principles waiting on a milestone

- [ ] 2.1 Seal moments on the device with the poster's content key (CDI-1863, M2)
- [ ] 2.2 Connect by touching two phones (CDI-1840, M2), with the link as the exception (CDI-1839)
- [ ] 2.3 A daily digest that stays silent on quiet days (CDI-1850, M4)
- [ ] 2.4 Publish the no-AI and no-sharing guarantee (CDI-1859)
- [ ] 2.5 Keep the app free of analytics, trackers and AI SDKs, checked in CI (CDI-1860)

## 3. Capability changes, restated one at a time

- [x] 3.1 `cap-01-core-identity`
- [x] 3.2 `cap-02-your-circle`
- [x] 3.3 `cap-03-moments`
- [ ] 3.4 `cap-04-reactions`
- [ ] 3.5 `cap-05-messaging`
- [ ] 3.6 `cap-06-ambient-presence`
- [ ] 3.7 `cap-07-user-ownership`
- [ ] 3.8 `cap-08-community-funding`
- [ ] 3.9 `cap-09-comments`
