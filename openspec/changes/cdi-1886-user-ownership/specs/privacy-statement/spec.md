## Purpose

What Nah? can and cannot see, said plainly in the app and on the site, together with the one cost of holding your own key, so nobody has to take a promise on trust.

## ADDED Requirements

### Requirement: A plain statement of what a server sees

Nah? SHALL publish a privacy statement in plain language, in the app and on its site. It SHALL say what a Nah? server can see: who is connected to whom, when each person posted or reacted, and the network addresses phones connect from. It SHALL say what a server cannot see: what a moment says or shows, its place and song, which reaction someone gave, and a person's name, gender, pronouns and orientation (ADR-0012, CDI-1895). It SHALL name every outside service a phone contacts on a person's behalf, and who runs each: the platforms' push services, the operating system's place lookup, the song-link lookup, and the music service a song's sample comes from. It SHALL say that nothing is sold, advertised or given to AI. While moments still travel unsealed (until CDI-1863), it SHALL say so.

#### Scenario: Reading the statement

- **WHEN** a person opens the privacy statement
- **THEN** it lists, in words they need no lawyer for, what a server can see, what it cannot, and every outside service the app contacts

#### Scenario: Before moments are sealed

- **WHEN** the app in use still posts moments unsealed
- **THEN** the statement says plainly that the server could read them

#### Scenario: A new outside service

- **WHEN** a version of the app starts contacting a service that earlier versions did not
- **THEN** the statement names that service before the version ships

### Requirement: The cost of holding your own key is said up front

The app SHALL tell a person, once before their first moment and wherever they look for it afterwards, that losing their phone loses their history unless two people in their circle vouch for a new one, and that with fewer than two connections there is no way back (ADR-0012, CDI-1865).

#### Scenario: Before the first moment

- **WHEN** a person is about to post for the first time
- **THEN** one plain sentence tells them that losing their phone loses their history unless two people in their circle vouch for a new one

#### Scenario: Looking for it later

- **WHEN** a person looks in the app for what happens if they lose their phone
- **THEN** they find the same sentence
