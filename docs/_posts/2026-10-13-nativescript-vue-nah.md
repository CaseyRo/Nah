---
title: "NativeScript-Vue (or: a good tool for someone else)"
date: 2026-10-13 12:00:00 +0200
categories: [Decisions]
tags: [flutter, nativescript, vue, decision-log]
draft: true
published: false
---

This morning I came across [a post about NativeScript-Vue](https://www.opensourceprojects.dev/post/nativescript-vue) and my first thought was: could this help Nah? talk to the phone? I asked Claude to look into it with me. The short answer is no, and I want to write down why, because the reasons say more about Nah? than about NativeScript.

## What the new version is

NativeScript lets you build iOS and Android apps in JavaScript, and draws real native buttons and lists instead of a web page dressed up as an app. NativeScript-Vue puts Vue on top of that. The news is that Vue 3 support is now generally available, with better TypeScript types, a plugin system, and Vue Devtools working on a phone. Vue 2 moves to its own branch, and there is an upgrade guide for people coming from it.

Its real party trick is older than this release: from JavaScript you can call the phone's own APIs directly, the iOS and Android ones, without writing Swift or Kotlin. If you already think in Vue, that is a lovely way to build a mobile app.

## Why it isn't for Nah?

"Native integration" for Nah? means a few concrete things. Touching phones to connect, with NFC where both phones allow it and a code to scan where they don't. A key that lives on your device instead of a password. Later maybe Bluetooth proximity and a home screen widget. Flutter reaches all of those today, through plugins, and where a plugin falls short I can write a small piece of Swift or Kotlin and call it from Dart. I accepted that cost back in [the Flutter decision]({{ '/posts/the-flutter-decision/' | relative_url }}).

NativeScript-Vue can't sit next to Flutter. It is a whole way of building the app, so using it would mean rewriting the client. What I would get for that is writing a few platform files in JavaScript instead of Swift. What I would give up is a Dart codebase that works, and months I would rather spend on getting two real phones talking to the server.

And the post itself is a release note. It is honest about what changed, but it has no numbers and no stories from apps in the wild, so there is nothing in it that says I'd end up somewhere better.

## What would change my mind

If Nah? ever needs something from the phone that Flutter can't reach, not even with a bit of Swift, I'll look again. The same goes if one of the things that would prove the Flutter decision wrong actually happens, like friends giving up on the install because the app is too big. Either one would start with a new decision record, not with a framework swap.

Until then, I'm noticing the pull of a shiny new tool and letting it pass. Nah? needs finishing more than it needs a new foundation.
