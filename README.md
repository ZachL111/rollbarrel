# rollbarrel

`rollbarrel` is a Go project in automation. Its focus is to plan backup retention and dry-run deletion reports.

## Use Case

The point is to make a small domain rule concrete enough that a reader can change it and immediately see what broke.

## Rollbarrel Review Notes

For a quick review, compare `rename risk` with `operator cost` before reading the middle cases.

## Highlights

- `fixtures/domain_review.csv` adds cases for dry-run spread and rename risk.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/rollbarrel-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `rename risk` and `operator cost`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Code Layout

The repository has two validation layers: the original compact policy fixture and the domain review fixture. They are separate so one can change without hiding failures in the other.

The Go implementation avoids hidden state so fixture changes are easy to reason about.

## Run The Check

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Regression Path

The same command runs the local verification path. The highest-scoring domain case is `stress` at 201, which lands in `ship`. The most cautious case is `edge` at 154, which lands in `ship`.

## Future Work

The repository is intentionally scoped to local checks. I would expand it by adding adversarial fixtures before adding features.
