// Package intervalset provides operations on sets of intervals, either ranges
// of numbers or periods of time.
//
// Intervals are half-open: an interval covers its lower limit but stops short
// of its upper one, so [0,5) and [5,10) share no value. It follows that:
//
//   - an interval whose limits are equal is empty and covers nothing. A set
//     ignores empty intervals, as it ignores invalid ones whose limits are out
//     of order;
//   - subtracting an interval splits the remaining ones at its limits without
//     losing or duplicating a single value;
//   - a set coalesces adjacent intervals: adding [0,5) and [5,10) to a set
//     leaves it holding [0,10).
//
// Two predicates tell intervals apart at their limits: Intersects reports
// whether they share a value, while Touches also accepts intervals that merely
// meet at a limit, which is what lets a set coalesce them. Adjacent intervals
// therefore touch without intersecting.
package intervalset

import (
	"slices"
	"sort"
)

// Interval is an interface that represents an interval within the set.
type Interval[T any] interface {
	// Min returns the minimum value of the interval.
	Min() T

	// Max returns the maximum value of the interval.
	Max() T

	// Equal reports whether two intervals are equal.
	Equal(Interval[T]) bool

	// Before reports whether the interval is before the given interval.
	Before(Interval[T]) bool

	// After reports whether the interval is after the given interval.
	After(Interval[T]) bool

	// Contains reports whether the interval contains the given interval.
	Contains(Interval[T]) bool

	// Intersects reports whether the interval shares a value with the given interval.
	Intersects(Interval[T]) bool

	// Touches reports whether the interval shares a value with the given
	// interval or meets it at a limit.
	Touches(Interval[T]) bool

	// Intersect returns a new interval representing the intersection of both intervals.
	Intersect(Interval[T]) Interval[T]

	// Encompass returns a new interval encompassing two overlapping intervals.
	Encompass(Interval[T]) Interval[T]

	// Punch cuts the given interval out of the interval and returns the remaining intervals.
	Punch(Interval[T]) (Interval[T], Interval[T])

	// IsZero reports whether both limits are zero values
	IsZero() bool

	// IsEmpty reports whether the interval's limits are equal,
	// meaning the interval covers nothing.
	IsEmpty() bool

	// IsValid reports whether the interval's lower limit is
	// lower than or equal to its upper limit.
	IsValid() bool
}

// EmptySet returns an empty set.
func EmptySet[T any]() *IntervalSet[T] {
	return &IntervalSet[T]{
		intervals: make([]Interval[T], 0),
	}
}

// IntervalSet is an ordered set of intervals.
type IntervalSet[T any] struct {
	intervals []Interval[T]
}

// AsSlice returns the set of intervals as a slice.
// The slice is a copy of the set's own: appending to it or replacing
// its intervals leaves the set untouched.
func (p *IntervalSet[T]) AsSlice() []Interval[T] {
	return slices.Clone(p.intervals)
}

// IsEmpty reports whether the set is empty.
func (p *IntervalSet[T]) IsEmpty() bool {
	return len(p.intervals) == 0
}

// Equal reports whether the set is equal to another set.
func (p *IntervalSet[T]) Equal(q *IntervalSet[T]) bool {
	if len(p.intervals) != len(q.intervals) {
		return false
	}
	for i, v := range p.intervals {
		if !v.Equal(q.intervals[i]) {
			return false
		}
	}
	return true
}

// Add adds the given interval to the set.
// Intervals will be merged with the intervals present in the set.
func (p *IntervalSet[T]) Add(intervals ...Interval[T]) *IntervalSet[T] {
	for _, q := range intervals {
		p.add(q)
	}
	return p
}

func (p *IntervalSet[T]) add(q Interval[T]) {
	// an invalid interval would break the set's ordering and an empty interval
	// covers nothing, so adding either of them to the set is a no-op.
	// keeping them out of the set also guarantees that the intervals we
	// encompass below always produce a non-zero interval.
	if !q.IsValid() || q.IsEmpty() {
		return
	}

	// the set is empty we can simply append the interval to the set.
	if p.IsEmpty() {
		p.intervals = append(p.intervals, q)
		return
	}

	// find the first interval that ends either during or after the given interval:
	// |   |
	// | T |---------------------------------->
	// |   |   x      i     i+1    i+2    i+3
	// | P | -----  -----  -----  -----  -----
	// | Q |          -------
	// |   |
	i := sort.Search(len(p.intervals), func(i int) bool {
		return !p.intervals[i].Before(q)
	})

	// there are no intervals to the right, all intervals are positioned before q
	// we then can simply append the interval to the set.
	if i == len(p.intervals) {
		p.intervals = append(p.intervals, q)
		return
	}

	stack := make([]Interval[T], 0)
	left, right := p.intervals[0:i], p.intervals[i:]
	cur, right := right[0], right[1:]

	interval := q

	for cur != nil {
		if cur.After(interval) {
			// we can safely append both intervals to the stack
			stack = append(stack, interval)
			stack = append(stack, cur)
			break
		}

		// both intervals must be overlapping
		// we create a new interval encompassing both
		interval = interval.Encompass(cur)
		if interval.IsZero() {
			panic("we should be able to get an encompassing interval")
		}

		// there are no more intervals to the right
		if len(right) == 0 {
			// we can simply append the interval to the stack
			stack = append(stack, interval)
			cur = nil
		} else {
			// otherwise, we keep looking overlapping intervals to the right
			cur, right = right[0], right[1:]
		}
	}

	// append the remaining intervals that might not have been discovered
	if len(right) > 0 {
		stack = append(stack, right...)
	}

	p.intervals = make([]Interval[T], 0)
	p.intervals = append(p.intervals, left...)  // 👈
	p.intervals = append(p.intervals, stack...) // 👉
}

// Sub subtracts the given intervals from the set.
func (p *IntervalSet[T]) Sub(intervals ...Interval[T]) *IntervalSet[T] {
	for _, q := range intervals {
		p.sub(q)
	}
	return p
}

func (p *IntervalSet[T]) sub(q Interval[T]) {
	// the set is empty we do not need to remove the interval
	if p.IsEmpty() {
		return
	}

	// find the first interval that ends either during or after the given interval:
	// |   |
	// | T |---------------------------------->
	// |   |   x      i     i+1    i+2    i+3
	// | P | -----  -----  -----  -----  -----
	// | Q |          -------
	// |   |
	i := sort.Search(len(p.intervals), func(i int) bool {
		return !p.intervals[i].Before(q)
	})

	// there are no intervals to the right, all intervals are positioned before q
	// there is therefore nothing to subtract.
	if i == len(p.intervals) {
		return
	}

	stack := make([]Interval[T], 0)
	left, right := p.intervals[0:i], p.intervals[i:]
	cur, right := right[0], right[1:]

	for cur != nil {
		// the interval is no longer overlapping the subtraction
		// we can stop here the next ones will be after the subtraction too
		if cur.After(q) {
			stack = append(stack, cur)
			break
		}

		l, r := cur.Punch(q)
		if !l.IsZero() {
			stack = append(stack, l)
		}
		if !r.IsZero() {
			stack = append(stack, r)
		}

		// there are no more intervals to the right
		if len(right) == 0 {
			// we can stop looking up overlapping intervals
			cur = nil
		} else {
			// we keep looking up overlapping intervals
			cur, right = right[0], right[1:]
		}
	}

	// append the remaining intervals that might not have been discovered
	if len(right) > 0 {
		stack = append(stack, right...)
	}

	p.intervals = make([]Interval[T], 0)
	p.intervals = append(p.intervals, left...)  // 👈
	p.intervals = append(p.intervals, stack...) // 👉
}

// Overlaps returns a new set containing the intervals overlapping q.
func (p *IntervalSet[T]) Overlaps(q Interval[T]) *IntervalSet[T] {
	l, h := p.rangeOfOverlap(q)

	s := EmptySet[T]()

	for _, v := range p.intervals[l:h] {
		i := v.Intersect(q)
		if !i.IsZero() {
			s.Add(i)
		}
	}

	return s
}

// IsSupersetOf reports whether p covers s, that is whether every interval
// of s is contained in one of p's intervals.
func (p *IntervalSet[T]) IsSupersetOf(s *IntervalSet[T]) bool {
	c := 0

	for _, q := range s.intervals {
		l, h := p.rangeOfOverlap(q)

		for _, v := range p.intervals[l:h] {
			if v.Contains(q) {
				c = c + 1
				break
			}
		}
	}

	return c == len(s.intervals)
}

// Complement returns a new set containing the intervals in q that are not in p.
func (p *IntervalSet[T]) Complement(q Interval[T]) *IntervalSet[T] {
	l, h := p.rangeOfOverlap(q)

	return EmptySet[T]().
		Add(q).
		Sub(p.intervals[l:h]...)
}

// Difference returns a new set containing the intervals in p that are not in q.
func (p *IntervalSet[T]) Difference(q *IntervalSet[T]) *IntervalSet[T] {
	return EmptySet[T]().
		Add(p.intervals...).
		Sub(q.intervals...)
}

// Iter iterates over the set and pass intervals to the anonymous function.
// It stops when the function returns false or when there are no more intervals to consume.
func (p *IntervalSet[T]) Iter(f func(Interval[T]) bool) {
	for _, i := range p.intervals {
		if !f(i) {
			break
		}
	}
}

// IterBetween iterates over the set between the given interval.
// Each interval is truncated to fit within q and pass intervals to the anonymous function.
// It stops when the function returns false or when there are no more intervals to consume.
func (p *IntervalSet[T]) IterBetween(q Interval[T], f func(Interval[T]) bool) {
	l, h := p.rangeOfOverlap(q)

	for _, v := range p.intervals[l:h] {
		i := v.Intersect(q)
		if i.IsZero() {
			continue
		}

		if !f(i) {
			break
		}
	}
}

// rangeOfOverlap returns the range to obtain a slice of intervals overlapping the given interval:
// - the lower limit is the index of the first interval that ends during or after the given interval
// - the higher limit is the index of the first interval that starts after the given interval
//
// |   |
// | T |---------------------------------->
// |   |   l     l+1    l+2    l+3     h
// | P | -----  -----  -----  -----  -----
// | Q |    ---------------------
// |   |
func (p *IntervalSet[T]) rangeOfOverlap(q Interval[T]) (int, int) {
	// an invalid interval overlaps nothing: its limits being out of order,
	// the lower limit would otherwise be searched past the higher one.
	if !q.IsValid() {
		return 0, 0
	}

	l := sort.Search(len(p.intervals), func(i int) bool {
		return !p.intervals[i].Before(q)
	})
	h := sort.Search(len(p.intervals), func(i int) bool {
		return p.intervals[i].After(q)
	})

	return l, h
}

// Union returns a new set that is the union of the sets.
func Union[T any](sets ...*IntervalSet[T]) *IntervalSet[T] {
	s := EmptySet[T]()
	for _, set := range sets {
		s.Add(set.intervals...)
	}
	return s
}

// Intersection returns a new set that is the intersection of the sets.
// It returns an empty set when called without any set.
func Intersection[T any](sets ...*IntervalSet[T]) *IntervalSet[T] {
	if len(sets) == 0 {
		return EmptySet[T]()
	}

	// start from the first set, then narrow it down with each of the others.
	a := EmptySet[T]().Add(sets[0].intervals...)

	for _, s := range sets[1:] {
		n := EmptySet[T]()
		for _, q := range a.intervals {
			n.Add(s.Overlaps(q).intervals...)
		}
		a = n
	}

	return a
}
