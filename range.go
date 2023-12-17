package intervalset

import "golang.org/x/exp/constraints"

// Number represents a value in a range.
type Number interface {
	constraints.Float | constraints.Integer
}

// NewRange returns a new range between lower and upper values.
func NewRange[T Number](l, u T) Range[T] {
	return Range[T]{lower: l, upper: u}
}

// Range represents a range between two numbers.
type Range[T Number] struct {
	lower T
	upper T
}

// Min returns the range's minimum value.
func (p Range[T]) Min() T {
	return p.lower
}

// Max returns the range's maximum value.
func (p Range[T]) Max() T {
	return p.upper
}

// IsZero reports whether both lower & upper values are zero values.
func (p Range[T]) IsZero() bool {
	return p.lower == 0 && p.upper == 0
}

// IsValid reports whether the range's lower value is actually lower than or equal to the upper value.
func (p Range[T]) IsValid() bool {
	return p.lower <= p.upper
}

// IsEmpty reports whether the range's values are equal.
func (p Range[T]) IsEmpty() bool {
	return p.lower == p.upper
}

// Equal reports whether p is equal to q.
// Two ranges are equal when their lower & upper values are equal.
func (p Range[T]) Equal(q Interval[T]) bool {
	return p.lower == q.Min() && p.upper == q.Max()
}

// Before reports whether p upper value is lower than q lower value.
func (p Range[T]) Before(q Interval[T]) bool {
	return p.upper < q.Min()
}

// After reports whether p lower value is greater than q upper value.
func (p Range[T]) After(q Interval[T]) bool {
	return p.lower > q.Max()
}

// Overlaps reports whether p overlaps q.
func (p Range[T]) Overlaps(q Interval[T]) bool {
	return !p.Before(q) && !p.After(q)
}

// Contains reports whether p contains q.
func (p Range[T]) Contains(q Interval[T]) bool {
	return q.Min() >= p.lower && q.Max() <= p.upper
}

// Intersect returns a new range representing the intersection of both ranges.
// The new range is either a valid and non-empty range (its lower value being strictly
// greater than its upper value) or a zero value range.
func (p Range[T]) Intersect(q Interval[T]) Interval[T] {
	if !p.Overlaps(q) {
		return Range[T]{}
	}

	l := max(p.lower, q.Min())
	u := min(p.upper, q.Max())

	if l >= u {
		return Range[T]{}
	}

	return Range[T]{
		lower: l,
		upper: u,
	}
}

// Encompass returns a new range encompassing both ranges.
// Both ranges must overlap otherwise it returns a zero value range.
func (p Range[T]) Encompass(q Interval[T]) Interval[T] {
	if !p.Overlaps(q) {
		return Range[T]{}
	}

	l := min(p.lower, q.Min())
	u := max(p.upper, q.Max())

	if l >= u {
		return Range[T]{}
	}

	return Range[T]{
		lower: l,
		upper: u,
	}
}

// Punch cuts q out of p and returns the remaining ranges.
func (p Range[T]) Punch(q Interval[T]) (Interval[T], Interval[T]) {
	i := p.Intersect(q)
	if i.IsZero() {
		if p.Before(q) {
			return p, Range[T]{}
		}
		return Range[T]{}, p
	}

	l := Range[T]{}
	r := Range[T]{}

	if p.lower != i.Min() {
		l = Range[T]{lower: p.lower, upper: i.Min()}
	}

	if i.Max() != p.upper {
		r = Range[T]{lower: i.Max(), upper: p.upper}
	}

	return l, r
}
