## Purpose

How a person's content key reaches the people they connect with, and when it changes, so that the project never holds a key and a leaked invitation unlocks nothing.

## ADDED Requirements

### Requirement: An invitation carries a secret, never a key

An invitation SHALL carry a one-time secret and SHALL NOT carry a content key. After a connection completes, each phone SHALL seal its person's content key, together with every earlier key, to the other person's device key. For a connection made by link, the phone SHALL first check that device key against the invitation's secret; for a touch or scanned code, the device keys SHALL come from the in-person exchange (ruled 2026-09-15).

#### Scenario: A photographed invitation

- **WHEN** someone photographs an invite link or code after it has been used
- **THEN** the photo holds no key and unlocks no moment

#### Scenario: Earlier moments on arrival

- **WHEN** Sam connects with Maya
- **THEN** Sam's phone receives Maya's current key and every earlier one, and can read her earlier moments

#### Scenario: A server in the middle

- **WHEN** a server replaces the newcomer's device key with one of its own during a link connection
- **THEN** the key does not check out against the invitation's secret, and no content key is sealed to it

### Requirement: A key rotates on every ending and every block

When one of a person's connections ends, or they block someone, that person's phone SHALL make a new content key for their future moments. It SHALL seal the new key to the device keys it received from each remaining connection, and never to a device key the server supplies. Earlier moments SHALL stay sealed with the keys they had (ruled 2026-09-15).

#### Scenario: Ending a connection

- **WHEN** Maya ends her connection with Sam
- **THEN** every moment Maya posts afterwards is sealed with a key Sam's phone never receives

#### Scenario: What was readable stays readable

- **WHEN** Maya's key has rotated
- **THEN** the people still in her circle can read her earlier moments and her new ones

#### Scenario: A key the server offers

- **WHEN** the server lists a device key that no phone received from a connection
- **THEN** no rotated key is sealed to it
