# Architecture spike

Throwaway measurement code for M0. None of this ships.

It exists to answer two questions with numbers rather than argument: which
runtime the circle server should be written in, and whether each circle gets its
own SQLite database or shares one.

- `THRESHOLDS.md` — the pass/fail numbers, written before any measurement
- `harness/` — seed and benchmark scripts, shared by every candidate so the
  comparison is a comparison
- `candidates/` — one minimal server per language, each implementing exactly two
  endpoints: post a moment, and list a circle's feed
- `RESULTS.md` — what was measured, filled in as it happens
- `candidates/go-fanin` and `harness.py --degree N` — added for CDI-1879 after
  ADR-0017: one file per person, and a feed read fan-in across everyone they are
  connected to. Its outcome is the 2026-09-14 amendment to ADR-0011.

The outcome is ADR-0011. When that is merged, this directory has done its job.
