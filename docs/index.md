---
layout: default
title: Home
---

# Nah

**A private home for your closest people.**

Nah ("Not Alone Here") is a social network for your inner circle — limited to 150 friends, private by default, open source by design.

## Why Nah?

Social media has become performance theater. We broadcast to hundreds of "friends" we barely know, curate ourselves for algorithmic approval, and measure worth in viral moments.

**Viral? Nah. Vital.**

Nah is different:
- **Small by design** — 150 friends max (Dunbar's number)
- **Private by default** — no public timelines, no algorithmic feeds
- **Real friends, real moments** — share life without performing
- **Open source** — community-funded, transparent, yours to verify

## Reflections

We're building in public. These posts document the decisions behind Nah:

{% for post in site.reflections reversed %}
- [{{ post.title }}]({{ post.url | relative_url }}) — {{ post.date | date: "%Y-%m-%d" }}
{% endfor %}

## Get Involved

Nah is open source and community-driven.

- **GitHub**: [Repository](https://github.com/your-org/nah) *(coming soon)*
- **Contribute**: Join the conversation, suggest features, submit code
- **Support**: Community-funded, no ads, no data sales

---

*"Share with your circle, not the world."*
