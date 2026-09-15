## Purpose

One person's moments on a page of their own, open to that person and the people in their circle, and to nobody else.

## ADDED Requirements

### Requirement: A page of one person's moments

Tapping a person's name SHALL open a page of their moments, newest first, and "you" SHALL open the person's own page (ruled 2026-09-15). The top of the page SHALL show the person's avatar and name, with no cover photo, followed by their first moment, the answer they gave on arriving (CDI-1883). Below it SHALL come their moments as the feed shows them, including deleted markers, one page at a time: older moments SHALL load only when the reader asks for them, never by scrolling alone. The page SHALL show no count of anything (ADR-0004, ruled 2026-09-15).

#### Scenario: Opening someone's page

- **WHEN** a person taps Maya's name on one of her moments
- **THEN** a page opens with Maya's avatar and name, her first moment, and her other moments newest first, and no number of moments, connections or anything else

#### Scenario: Your own page

- **WHEN** a person opens "you" from the feed header
- **THEN** they see their own moments, newest first

#### Scenario: Older moments

- **WHEN** a reader reaches the end of the moments shown on a person's page
- **THEN** they are offered to show earlier moments, and nothing more loads until they take it

### Requirement: Only the circle can open it

A person's page SHALL be available only to that person and the people in their circle. When a connection ends, each person's page SHALL stop being available to the other.

#### Scenario: Outside the circle

- **WHEN** a request for Maya's page comes from anyone who is neither Maya nor in her circle
- **THEN** the server refuses it, and returns nothing of Maya's

#### Scenario: After a connection ends

- **WHEN** Sam ends his connection with Maya
- **THEN** neither of them can open the other's page any longer
