# Golden Cases

The golden files are not a benchmark. They are a stable review surface for automation.

The main golden fixture is `fixtures/golden/scoreboard.csv`. The matrix fixture is `fixtures/golden/lane-matrix.csv`. Together they cover `dry-run spread`, `rename risk`, `operator cost`, and `idempotence` with different score ranges.

The purpose is to make intentional rule changes show up in review.
