package intervalset

import (
	"time"

	limit "github.com/mickaelvieira/intervalset/internal/time"
)

// NewPeriod returns a new period with the given start and end dates.
func NewPeriod(s, e time.Time) Period[time.Time] {
	return Period[time.Time]{start: s, end: e}
}

// Period represents a portion of time.
type Period[T time.Time] struct {
	start T
	end   T
}

// Min returns the period's minimum value.
func (p Period[T]) Min() T {
	return p.start
}

// Max returns the period's maximum value.
func (p Period[T]) Max() T {
	return p.end
}

// IsZero reports whether both start & end dates are zero values.
func (p Period[T]) IsZero() bool {
	return time.Time(p.start).IsZero() && time.Time(p.end).IsZero()
}

// IsValid reports whether the period's start date is lower than or equal to its end date.
// A period with a start date greater than its end date would be indeed invalid.
func (p Period[T]) IsValid() bool {
	return !time.Time(p.start).After(time.Time(p.end))
}

// IsEmpty reports whether the period's start date is equal to its end date,
// meaning the period's duration would be equal to zero and therefore be empty.
func (p Period[T]) IsEmpty() bool {
	return time.Time(p.start).Equal(time.Time(p.end))
}

// Equal reports whether p is equal to q.
// Two periods are equal when their start & end dates are equal.
func (p Period[T]) Equal(q Interval[T]) bool {
	return time.Time(p.start).Equal(time.Time(q.Min())) &&
		time.Time(p.end).Equal(time.Time(q.Max()))
}

// Before reports whether p ends
// before the beginning of q.
func (p Period[T]) Before(q Interval[T]) bool {
	return time.Time(p.end).Before(time.Time(q.Min()))
}

// After reports whether p starts
// after the end of q.
func (p Period[T]) After(q Interval[T]) bool {
	return time.Time(p.start).After(time.Time(q.Max()))
}

// Overlaps reports whether p overlaps q.
func (p Period[T]) Overlaps(q Interval[T]) bool {
	return !p.Before(q) && !p.After(q)
}

// Contains reports whether p contains q.
func (p Period[T]) Contains(q Interval[T]) bool {
	s := time.Time(q.Min())
	e := time.Time(q.Max())

	return (s.After(time.Time(p.start)) || s.Equal(time.Time(p.start))) &&
		(e.Before(time.Time(p.end)) || e.Equal(time.Time(p.end)))
}

// Intersect returns a new period representing the intersection of both periods.
// The new period is either a valid and non-empty period (its start date being strictly
// greater than its end date) or a zero value period.
func (p Period[T]) Intersect(q Interval[T]) Interval[T] {
	if !p.Overlaps(q) {
		return Period[T]{}
	}

	s := limit.Max(time.Time(p.start), time.Time(q.Min()))
	e := limit.Min(time.Time(p.end), time.Time(q.Max()))

	if !s.Before(e) {
		return Period[T]{}
	}

	return Period[T]{
		start: T(s),
		end:   T(e),
	}
}

// Encompass returns a new period encompassing both periods.
// Both periods must overlap otherwise it returns a zero value period.
func (p Period[T]) Encompass(q Interval[T]) Interval[T] {
	if !p.Overlaps(q) {
		return Period[T]{}
	}

	s := limit.Min(time.Time(p.start), time.Time(q.Min()))
	e := limit.Max(time.Time(p.end), time.Time(q.Max()))

	if !s.Before(e) {
		return Period[T]{}
	}

	return Period[T]{
		start: T(s),
		end:   T(e),
	}
}

// Punch cuts q out of p and returns the remaining periods.
func (p Period[T]) Punch(q Interval[T]) (Interval[T], Interval[T]) {
	i := p.Intersect(q)
	if i.IsZero() {
		if p.Before(q) {
			return p, Period[T]{}
		}
		return Period[T]{}, p
	}

	l := Period[T]{}
	r := Period[T]{}

	if !time.Time(p.start).Equal(time.Time(i.Min())) {
		l = Period[T]{start: p.start, end: i.Min()}
	}

	if !time.Time(i.Max()).Equal(time.Time(p.end)) {
		r = Period[T]{start: i.Max(), end: p.end}
	}

	return l, r
}
