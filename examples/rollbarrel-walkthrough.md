# Rollbarrel Walkthrough

The fixture is intentionally compact, so the review starts with the cases that pull farthest apart.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | dry-run spread | 175 | ship |
| stress | rename risk | 201 | ship |
| edge | operator cost | 154 | ship |
| recovery | idempotence | 179 | ship |
| stale | dry-run spread | 182 | ship |

Start with `stress` and `edge`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

If `edge` becomes less cautious without a clear reason, I would inspect the drag input first.
