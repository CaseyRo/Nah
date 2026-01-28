---
title: "Why Design Craft Matters"
date: 2025-01-28 14:00:00 +0100
categories: [Design, Philosophy]
tags: [design-system, animation, ux, path]
---

We spent a lot of time this week staring at [Brian Lovin's Path dissection](https://brianlovin.com/app-dissection/path-ios). If you haven't seen it, go look. He recorded every animation, every modal, every micro-interaction.

What struck us wasn't any single feature. It was the *consistency*. Every modal bounced the same way. Every menu had the same timing. The app felt like one person touched every pixel.

## The thing about Path

Path wasn't successful in the end. It got acquired, shut down, became a footnote. But people who used it *remember* it. There's a reason we're building Nah — people still talk about how Path *felt*.

That feeling came from craft. From someone deciding that the radial menu shouldn't just pop open, it should bloom. Each icon should rotate slightly as it expands. The whole thing should take exactly 400ms with a spring curve that overshoots just a little.

[TechWalls' review from 2013](https://www.techwalls.com/path-app-iphone-review/) called it "one of the most beautiful apps on iPhone." [Digital Trends](https://www.digitaltrends.com/social-media/path-2-0-relaunches-as-your-personal-social-network/) said Path 2.0 was "stunning." These weren't feature reviews. They were reactions to craft.

## What we decided

We spent time on [designpieces.com](https://www.designpieces.com/palette/path-color-palette-hex-and-rgb/) looking at Path's exact color values. Pomegranate red: `#EE3423`. Not pure red (too aggressive), not salmon (too soft). Warm, energetic, alive.

We're using it. Not because we're copying Path, but because they found the right answer and we'd be foolish to pretend we'd find a better one.

Same logic for the animations. Spring physics with overshoot. 400ms for entrances. 250ms for exits. Staggered timing on the radial menu. These aren't arbitrary numbers — they're what makes something feel "bouncy" versus "snappy" versus "sluggish."

## The cost

This stuff takes longer. A modal that just appears is one line of code. A modal that scales from 95% with spring easing, overshoots, then settles — that needs an animation library, careful timing, testing on low-end devices.

We're paying that cost. Because every time someone opens Nah, we want their brain to register: *someone cared about this*.

That's the whole point.

---

**Sources:**

- [Brian Lovin - Path App Dissection](https://brianlovin.com/app-dissection/path-ios)
- [TechWalls - Path App Review](https://www.techwalls.com/path-app-iphone-review/)
- [Digital Trends - Path 2.0 Review](https://www.digitaltrends.com/social-media/path-2-0-relaunches-as-your-personal-social-network/)
- [Design Pieces - Path Color Palette](https://www.designpieces.com/palette/path-color-palette-hex-and-rgb/)
