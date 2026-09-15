## Purpose

A person uses Nah? on exactly one phone, deliberately, so that there is one path onto a new phone and no quiet way to add a reader.

## ADDED Requirements

### Requirement: One phone per person

A person SHALL be signed in on exactly one phone. No second device SHALL be added, and a new phone SHALL take over only through recovery, even while the old phone still works. When recovery completes, the old phone's device key SHALL stop working at once (ADR-0015, ruled 2026-09-15).

#### Scenario: A second phone

- **WHEN** someone signed in on one phone tries to sign in as the same person on another
- **THEN** it is refused, and they are told that a new phone takes over through recovery

#### Scenario: A new phone while the old one works

- **WHEN** a person buys a new phone and their old one still works
- **THEN** moving to it goes through recovery, the same as for a lost phone

#### Scenario: After recovery

- **WHEN** recovery completes on the new phone
- **THEN** the old phone can no longer sign in, post, or receive anything new
