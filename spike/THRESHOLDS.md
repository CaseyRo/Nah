# Spike thresholds

Written **before** any measurement, per CDI-1818. The spike exists to confirm or
refute these numbers, not to explore until something interesting turns up.

Anything measured that is not on this list is interesting rather than decisive,
and does not get a vote.

## The workload we are sizing for

One server hosts many circles. A circle has at most 150 members. From the
research:

| Shape | Members | Moments a day | Media a year |
|---|---|---|---|
| Family circle | 12–20 | ~3 | ~0.35 GB |
| Busy clique | 150 | ~20, half photos | ~2.5 GB |

The seeded fixture for every test is **50 circles with one year of history**,
mixed between those two shapes, which is a plausible small hosted server.

Two properties of this product make the load unusual, and both are deliberate:

- **Moments are opaque blobs.** The server never parses, indexes, ranks or
  counts content, so there is no query more complex than "the newest thirty rows
  in this circle".
- **Digests are the only synchronised load.** Everything else arrives when a
  human opens an app. Digests fire at times people chose, and people choose
  round numbers.

## Thresholds

| # | Measure | Threshold | Why this number |
|---|---|---|---|
| 1 | Feed query, p95, 50 concurrent readers | **< 50 ms** | It is one indexed query against local storage. Anything slower means the shape is wrong, not the machine. |
| 2 | Post a moment, p95 | **< 100 ms** | Includes a write and the response. Media upload is measured separately. |
| 3 | Idle memory, 50 circles seeded | **< 150 MB** | Has to leave room on a 1 GB box for everything else a person runs. |
| 4 | Memory under load | **< 400 MB** | Same reason. A server that balloons under load cannot be hosted densely. |
| 5 | Cold start to first response | **< 2 s** | Restarts must be boring, including after an update on someone's NAS. |
| 6 | Deployable artifact size | **< 100 MB** | It has to be handed to a person as one thing. |
| 7 | Restore one circle from backup | **< 60 s** | Every promise about leaving, moving or paying for hosting rests on this. |
| 8 | Digest storm, 5,000 circles in one minute | **completes < 60 s, memory within #4** | The only load that arrives all at once. |
| 9 | Circles per 2 GB box, extrapolated | **≥ 200** | Decides whether hosting can ever cost less than it charges. |
| 10 | Decrypt a screenful of 30 moments, low-end phone | **< 100 ms** | This is what "switching circles has no pause" means in practice (ADR-0013). |
| 11 | Encrypt a photo plus thumbnail, low-end phone | **< 300 ms** | Slower than this and composing feels broken. |

Thresholds 10 and 11 need a physical Android device and are expected to run
after the server work, in CDI-1871. They are listed here so the numbers are
fixed in advance like all the others.

## What we are deliberately not measuring

- Requests per second at scale. A circle is 150 people. If throughput is ever
  the binding constraint, something has gone wrong with the product, not the
  server.
- Query planner behaviour on complex joins. There are no complex joins.
- Horizontal scaling. A circle lives on one server, by design.

## The tie-break, agreed in advance

If two candidates both pass every threshold, take **Go**. "Close" means both
inside the thresholds, not within a rounding error of each other. A candidate
that fails any threshold is out regardless of how pleasant it was to write.

Time box: if this has not produced a decision within two weeks of work, take Go
and move on. The documented failure mode of this project is research
substituting for code.
