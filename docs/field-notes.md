# Field Notes

The useful part of this repository is the small rule set around dry-run spread and operator cost.

The domain cases cover `dry-run spread`, `rename risk`, `operator cost`, and `idempotence`. They sit beside the smaller starter fixture so the project has both a compact scoring check and a domain-flavored review check.

The widest spread is between `rename risk` and `operator cost`, so those are the first two cases I would preserve during a refactor.

The point is not to make the repository bigger. The point is to make the important judgment testable.
