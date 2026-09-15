## Purpose

Getting back into Nah?, your circle and your history on a new phone, confirmed by people who know you rather than by the server.

## ADDED Requirements

### Requirement: Two people vouch in person

A person SHALL be able to move to a new phone when two different people in their circle each confirm, by touching that phone or scanning its code in person, that it is them. Vouching SHALL NOT be possible by link. A person with fewer than two connections SHALL NOT be able to recover, and SHALL be told to join again with a new invitation (ruled 2026-09-15).

#### Scenario: Two vouchers

- **WHEN** Ana and Maya, both in Sam's circle, each touch Sam's new phone and confirm it is Sam
- **THEN** Sam's recovery starts its waiting period

#### Scenario: Vouching by link

- **WHEN** someone tries to vouch for a new phone through a link
- **THEN** it is refused

#### Scenario: A small circle

- **WHEN** a person with one connection loses their phone and asks to recover
- **THEN** they are told they cannot, and that they can join again with a new invitation

### Requirement: Forty-eight hours in which anyone can stop it

After the second confirmation, recovery SHALL wait 48 hours. During the wait, the person's old phone, if it still has Nah?, and everyone in the person's circle SHALL be told that the person is moving to a new phone and who vouched for it, by name, and any of them SHALL be able to stop the recovery. A stopped recovery SHALL NOT start again for 7 days, and a later attempt SHALL need two vouchers again and SHALL tell the circle again (ruled 2026-09-15). No key SHALL reach the new phone before the wait ends.

#### Scenario: The circle is told

- **WHEN** Sam's recovery starts its wait
- **THEN** everyone in Sam's circle, and his old phone if it still works, is told that Sam is moving to a new phone, confirmed by Ana and Maya

#### Scenario: Stopping it

- **WHEN** anyone who was told stops Sam's recovery during the 48 hours
- **THEN** the recovery ends, and the new phone receives nothing

#### Scenario: Trying again

- **WHEN** Sam's recovery was stopped and he tries again three days later
- **THEN** it is refused until 7 days have passed, and a later attempt needs two vouchers and tells his circle again

### Requirement: What comes back when the wait ends

When the wait ends, the person's identity SHALL move to the new phone, and the old device key SHALL stop working. The person's own content key and its history SHALL come back from a voucher's phone, which already holds it. Each other person in the circle SHALL seal their content key to the new phone when their app next comes online, trusting it because two people from the circle vouched for it, never because the server says so (CDI-1865).

#### Scenario: Your own history

- **WHEN** Sam's recovery completes
- **THEN** his own earlier moments are readable on the new phone

#### Scenario: Your circle's moments

- **WHEN** Ana's app comes online after Sam's recovery completed
- **THEN** it seals Ana's content key to Sam's new phone, and Sam can read her moments again

#### Scenario: A phone nobody vouched for

- **WHEN** the server presents a new device key for Sam that two people from his circle did not vouch for
- **THEN** no phone in the circle seals a key to it
