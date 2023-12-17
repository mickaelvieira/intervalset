package intervalset

import "sort"

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

	// Overlaps reports whether the interval overlaps the given interval.
	Overlaps(Interval[T]) bool

	// Intersect returns a new interval representing the intersection of both intervals.
	Intersect(Interval[T]) Interval[T]

	// Encompass returns a new interval encompassing two overlapping intervals.
	Encompass(Interval[T]) Interval[T]

	// Punch cuts the given interval out of the interval and returns the remaining intervals.
	Punch(Interval[T]) (Interval[T], Interval[T])

	// IsZero reports whether both limits are zero values
	IsZero() bool
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

// AsSlice returns the underlying set of intervals as a slice.
func (p *IntervalSet[T]) AsSlice() []Interval[T] {
	return p.intervals
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

	for cur != nil && !cur.IsZero() {
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

	for cur != nil && !cur.IsZero() {
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

// IsSubset reports whether s is a subset of p.
func (p *IntervalSet[T]) IsSubset(s *IntervalSet[T]) bool {
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
func Intersection[T any](sets ...*IntervalSet[T]) *IntervalSet[T] {
	a := EmptySet[T]()

	for i, s := range sets {
		// no need to go through the intervals of the last set
		// as it's already been compared against the other sets.
		if i == len(sets)-1 {
			continue
		}

		for _, p := range s.intervals {
			var o *IntervalSet[T]
			for j, ss := range sets {
				if j == i {
					continue
				}

				if o == nil {
					o = ss.Overlaps(p)
					continue
				}

				ol := make([]*IntervalSet[T], 0)
				for _, p := range o.intervals {
					ol = append(ol, ss.Overlaps(p))
				}
				o = Union(ol...)
			}

			if !o.IsEmpty() {
				a.Add(o.intervals...)
			}
		}
	}

	return a
}
