## Purpose

The marks that say Nah?: the logo, the app icons, and the site's favicon and link previews. Colours come from `DESIGN.md`.

## ADDED Requirements

### Requirement: The name is written Nah?

Every asset that spells the product's name SHALL write it "Nah?", with the question mark.

#### Scenario: A wordmark

- **WHEN** the logo, an icon or a preview image includes the name
- **THEN** it reads "Nah?"

### Requirement: One logo master

The logo SHALL exist as an SVG master, and every raster version (icons, launch screen, favicon, preview image) SHALL be exported from it.

#### Scenario: A new size is needed

- **WHEN** an asset needs the logo at a size that has not been exported
- **THEN** it is exported from the SVG master, not redrawn or scaled up from a smaller raster

### Requirement: App icons that survive every mask

The app SHALL have an iOS icon set made from a 1024 by 1024 master and an Android adaptive icon with separate foreground and background layers. Both SHALL put a white mark on the brand primary colour from `DESIGN.md`, and neither SHALL place any part of the mark where a platform mask can cut it off.

#### Scenario: iOS corners

- **WHEN** iOS rounds the icon's corners
- **THEN** no part of the mark is cut off

#### Scenario: Android launcher shapes

- **WHEN** a launcher masks the adaptive icon as a circle, a squircle or a rounded square
- **THEN** the whole mark stays visible

#### Scenario: Light and dark home screens

- **WHEN** the icon is shown on a light or a dark home screen
- **THEN** it stays recognisable

### Requirement: The site's favicon and default link preview

The published site SHALL serve a favicon and one default link-preview image of 1200 by 630 pixels.

#### Scenario: A browser tab

- **WHEN** the site is opened in a browser
- **THEN** the tab shows the Nah? favicon

#### Scenario: Sharing the site

- **WHEN** a link to the site is pasted into a chat app
- **THEN** the preview shows the default image

### Requirement: A link preview never shows a person

No link preview SHALL show a person's name, photo or moments. That includes the preview of an invite link, because Nah? has no public profiles (ADR-0008) and an invitation is private.

#### Scenario: Pasting an invite link

- **WHEN** an invite link is pasted into WhatsApp, Signal or iMessage
- **THEN** the preview shows the default Nah? image and nothing about who sent it
