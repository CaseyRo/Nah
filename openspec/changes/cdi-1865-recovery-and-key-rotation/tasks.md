# Recovery and key rotation tasks

Specification only this round. Nothing below is started.

## 1. Content keys

- [ ] 1.1 Invitations carry a one-time secret and no key (with CDI-1836 and CDI-1863)
- [ ] 1.2 After a connection completes, seal the content key and its history to the other device key, checked against the secret for links
- [ ] 1.3 Rotate on every ending and block, sealed to the device keys the phone already holds

## 2. One device

- [ ] 2.1 Refuse a second device key for a person
- [ ] 2.2 Stop the old device key the moment recovery completes

## 3. Recovery

- [ ] 3.1 Vouching by touch or scanned code from two different people in the circle; never by link
- [ ] 3.2 The 48-hour wait, with notices to the old phone and the circle, and a veto for any of them (needs push, M4)
- [ ] 3.3 Move the identity; restore the person's own history from a voucher's phone
- [ ] 3.4 Re-seal every other connection's key to the new phone, trusting the vouchers' signatures, never the server
- [ ] 3.5 Refuse recovery to anyone with fewer than two connections, and point them to joining again
