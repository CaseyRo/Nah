---
version: 1
slug: "apps-mobile-lib-main-dart"
primary_target: "apps/mobile/lib/main.dart"
related_targets: []
---

# Feed (the main timeline)

Scope: the feed screen of the Nah? app, the one a joined phone opens on. Visitor mode: Operate. Deliverable for this round: two 390x844 phone frames in the Nah? Paper file, light and dark, dated 2026-09-15; no Flutter code.

Audience and job: a person on the couch on a Sunday evening, reading what their circle shared, a day or two back. The task is reading, then one deliberate act: react, or start a moment with the +.

Content: moments of the four types (text, photo, voice, music), one of the reader's own with poster-only reactions, and a deleted marker. Synthetic people and words, labelled synthetic in the frame's layer names.

Constraints: DESIGN.md's world unchanged; product-principles (no counts, no ranking, no chat, no badges); app-shell (no tab bar, circle and you from the header, the clock and the + the only chrome); feed spec (a treatment per type, the clock, markers keep their place, a plain end); reactions spec (own moment shows faces and names, no button on it).

Unresolved: the voice card's visual treatment (this round designs it), the dark palette values (this round sets them), the logo (does not exist; the wordmark stands in).

## Direction contract

THESIS: the feed is a day read down a time axis, not a stack of posts. The left lane is time; the right lane is people. It refuses the card stack (every feed's shape) and the chat thread (faces as the axis).

OWN-WORLD: Surface 2 ground with no card boxes; a 64px left lane carrying each moment's time in Caption Text 2 and a 1px Surface 4 rail at its edge; the 48px analog clock riding the lane as the cursor, second hand Pomegranate; content set directly on the ground in the platform sans, names 16/600, text moments 17px; photos bleeding to the right edge; the 56px Pomegranate + bottom-right; Nunito only on the wordmark. With every word removed, the lane, the rail and the clock still say Nah?.

STORY: the reader sees at a glance when each thing happened and who did it, reads down without anything asking for attention, gives one reaction, and leaves. Nothing counts, nothing ranks, nothing suggests.

FIRST VIEWPORT: status bar; header with three overlapping faces left (your circle), "Nah?" centred in Nunito 24/700, your 32px avatar right; hairline. Then, mid-scroll: 21:14 Maya, text with a place, reaction button; 19:02 Sam, photo bleeding right, short line and place, the clock in the lane at this row reading 19:02 with "Sat 13 Sep" below; 17:40 Ana, voice with play and waveform; 16:05 You, music with 64px artwork and two reactions by face and name; 14:30 a deleted marker. The + at bottom-right, 24px inset.

FORM: two lanes, time on the left; fifth on my ranked list of seven; seed key 05819c70.

FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, DESIGN.md, and every shipping raster carrying its provenance.

## Detail boards (2026-09-16)

Shape brief confirmed by Casey on 2026-09-16: eight annotated boards in the Nah? Paper file, page "Feed details — 2026-09-16", documenting the feed's details for the M2 build. Same world and composition as the frames above; no new surface concept.

1. The + (anatomy): disc, arc of four items, hit zones, left-hand mirror, dim and glass backdrop.
2. The bloom (storyboard, 7 frames): press, rotate, fan out with stagger, choose, cancel, Reduce Motion, VoiceOver.
3. Placing content (anatomy): lane, rail, gutters, stack rhythm, photo bleed, foot line, hit zones, safe areas.
4. The clock (anatomy): parts, when it appears and fades, digits under Reduce Motion.
5. Scrolling a day (storyboard): clock in, hands sweeping, stop, fade; pull to refresh as the authored moment.
6. Reacting to someone's moment (storyboard, 6 frames): press, row rises, choose, settle, change or take back, close; no double tap, no long press, no swipe.
7. Opening a moment (storyboard, 4 frames): matched-geometry push to a moment page; edge-swipe back.
8. Playing a voice moment (strip, 3 frames).

Rulings proposed on the boards, for Casey to confirm: arc of 90 degrees with items at 0, 30, 60, 90 and text nearest the thumb; tall photos capped at 4:5; clock hands travel at most one turn across a gap of days; reaction row carries names for VoiceOver only. The moment page is new and needs a requirement in cdi-1884-moments and a Linear issue before build.

Rulings by Casey later on 2026-09-16, after the boards' first review: the clock is fixed at the middle of the lane and runs with you, its hands turning through the time between the moments above and below the middle; the second hand was misaligned and is redrawn through the pivot. The moment page is withdrawn: a photo, a voice moment or a song opens in focus on the bare ground with nothing else on it (a gallery wall), text does not focus; voice and music play only in focus, with the waveform in Pomegranate while it plays, and a song in focus plays a 30 second sample under the one action, open in the listener's own service. Board 07 is now "Focusing on a moment", board 08 "Voice and music in focus". The feed spec's "plays in place" and "chooser of services" need amending in cdi-1884-moments with a Linear issue. A blog post on the + (docs/_posts/2026-09-16-the-plus.md) is drawn from board 01.
