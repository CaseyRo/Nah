---
title: "How Privacy Actually Works in Nah"
date: 2025-01-28 11:00:00 +0100
categories: [Architecture, Privacy]
tags: [privacy, mastodon, security]
---

One of the first questions we had to answer was deceptively simple: if everyone's data is on the same server, how do we keep it private between friends?

## The Scenario

Let's say:
- **X** joins Nah (first user)
- **Y** joins Nah and adds X as a friend
- **Y** invites **Z**, and they become friends
- But **X and Z are not friends**

What should happen?

```
X ←——friends——→ Y ←——friends——→ Z

X sees: Y's posts
Y sees: X's posts + Z's posts
Z sees: Y's posts
X cannot see: Z's posts
```

X and Z are on the same server. Their data is in the same database. But they shouldn't be able to see each other's posts because they're not connected.

## The Naive Concern

Our initial worry: "If all content is stored on one server and everyone can access it, doesn't that defeat the purpose of privacy?"

It's a reasonable concern. But it misunderstands how social network privacy works.

## How Mastodon Already Handles This

Here's the key insight: **the database stores everything, but the API only returns what you're authorized to see.**

Mastodon posts have visibility levels:
- **Public**: Anyone can see
- **Unlisted**: Anyone with a link can see
- **Followers-only**: Only your followers can see ← **this is what we use**
- **Direct**: Only mentioned users (like a DM)

When X requests their home timeline, Mastodon checks: "Who does X follow?" Then it returns posts only from those people. X doesn't follow Z, so Z's posts are never returned — even though they exist in the same database.

This is standard access control. It's how every social network works, from Facebook to Instagram to Twitter's protected accounts.

## What We Disable

The danger in vanilla Mastodon is the **Local Timeline** — a feed showing all public posts from everyone on an instance. It's the "town square" feature that lets you discover strangers.

For Nah, we:

1. **Never expose the local timeline** in our UI
2. **Never expose the public timeline** either
3. **Default all posts to followers-only** at the server level
4. **Only show the home timeline** (posts from your friends)

Users literally cannot discover Z exists unless Y introduces them. There's no search, no explore page, no algorithmic suggestions. If you're not friends, you don't exist to each other.

## Privacy by Architecture, Not Policy

This is an important distinction. We're not saying "please don't look at other people's posts" (policy). We're saying "you physically cannot access posts from people you don't follow" (architecture).

The API enforces it. The UI enforces it. Even if someone tried to hack around it, the server wouldn't return the data.

## What About Admins?

One honest caveat: server admins can technically see the database. This is true of literally every hosted service — Instagram, Gmail, Slack, all of them. The people running the servers can access the data.

For Nah, we handle this with transparency:
- **Simple privacy policy** (human-readable, not legal jargon)
- **Clear statement**: admins can technically access data, standard for any hosted service
- **Commitment**: no data sales, no ads, no third-party sharing
- **User ownership**: export your data anytime, delete anytime

If end-to-end encryption for posts becomes important (so even admins can't read them), that's a v2 feature. For now, we're honest about the standard trust model of hosted services.

## The Mental Model

Think of Nah like an apartment building:
- Everyone lives in the same building (same server)
- Each apartment has a lock (followers-only visibility)
- You can only enter apartments where you have a key (mutual friendship)
- The building manager has a master key (admin access)
- But there's no directory of residents (no local timeline)

You know your neighbors exist because you've met them. You can't wander the halls looking through windows.

## Why This Matters

Nah's promise is "share with your circle, not the world." That only works if the architecture enforces it. You shouldn't have to trust that people will respect your privacy — the system should make violating it impossible.

By defaulting to followers-only and hiding discovery features, we make private sharing the only kind of sharing. It's not a setting you might forget to enable. It's how the app works.

---

*This is part of an ongoing series documenting the decisions behind Nah. We believe in building trust through transparency — even about the things we can't perfectly solve (like admin access).*
