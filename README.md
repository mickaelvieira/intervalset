# IntervalSet [![Go Reference](https://pkg.go.dev/badge/github.com/mickaelvieira/intervalset.svg)](https://pkg.go.dev/github.com/mickaelvieira/intervalset)

`intervalset` is a Go package that provides [operations](https://pkg.go.dev/github.com/mickaelvieira/intervalset#pkg-examples) on set of intervals. It supports two types of intervals:
- `Range`: a range is an interval of numbers, either floats or integers;
- `Period`: a period is an interval of [Time](https://pkg.go.dev/time#Time).

Intervals are half-open: an interval covers its lower limit but stops short of its upper one, so `[0,5)` and `[5,10)` share no value and a set coalesces them into `[0,10)`. The [package documentation](https://pkg.go.dev/github.com/mickaelvieira/intervalset) covers what follows from it.

## Limitations

- A range cannot cover the maximum value of a bounded integer type. Covering `255` in a `Range[uint8]` would take an upper limit of `256`, which overflows, and the same goes for `math.MaxInt64` and friends. Lower limits are unaffected, and periods are not concerned as `time.Time` has no such bound.
- Ranges of integers read with an off-by-one: the values `1` to `5` are `NewRange(1, 6)`.
- Empty intervals, whose limits are equal, and invalid ones, whose limits are out of order, are ignored rather than reported. Adding either to a set leaves the set unchanged.
