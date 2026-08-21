package intervalset

import (
	"time"

	limit "github.com/mickaelvieira/intervalset/internal/time"
)

// NewPeriod returns a new period with the given start and end dates.
func NewPeriod(s, e time.Time) Period {
	return Period{start: s, end: e}
}

// Period represents a portion of time.
type Period struct {
	start time.Time
	end   time.Time
}

// Min returns the period's minimum value.
func (p Period) Min() time.Time {
	return p.start
}

// Max returns the period's maximum value.
func (p Period) Max() time.Time {
	return p.end
}

// IsZero reports whether both start & end dates are zero values.
func (p Period) IsZero() bool {
	return p.start.IsZero() && p.end.IsZero()
}

// IsValid reports whether the period's start date is lower than or equal to its end date.
// A period with a start date greater than its end date would be indeed invalid.
func (p Period) IsValid() bool {
	return !p.start.After(p.end)
}

// IsEmpty reports whether the period's start date is equal to its end date,
// meaning the period's duration would be equal to zero and therefore be empty.
func (p Period) IsEmpty() bool {
	return p.start.Equal(p.end)
}

// Equal reports whether p is equal to q.
// Two periods are equal when their start & end dates are equal.
func (p Period) Equal(q Interval[time.Time]) bool {
	return p.start.Equal(q.Min()) && p.end.Equal(q.Max())
}

// Before reports whether p ends
// before the beginning of q.
func (p Period) Before(q Interval[time.Time]) bool {
	return p.end.Before(q.Min())
}

// After reports whether p starts
// after the end of q.
func (p Period) After(q Interval[time.Time]) bool {
	return p.start.After(q.Max())
}

// Intersects reports whether p and q share at least one instant.
// Adjacent periods do not intersect, they only touch.
func (p Period) Intersects(q Interval[time.Time]) bool {
	s := limit.Max(p.start, q.Min())
	e := limit.Min(p.end, q.Max())

	return s.Before(e)
}

// Touches reports whether p and q share an instant or meet at a limit.
// Adjacent periods touch, which is what lets a set coalesce them.
func (p Period) Touches(q Interval[time.Time]) bool {
	return !p.Before(q) && !p.After(q)
}

// Contains reports whether p contains q.
func (p Period) Contains(q Interval[time.Time]) bool {
	s := q.Min()
	e := q.Max()

	return (s.After(p.start) || s.Equal(p.start)) &&
		(e.Before(p.end) || e.Equal(p.end))
}

// Intersect returns a new period representing the intersection of both periods.
// The new period is either a valid and non-empty period (its start date being
// strictly before its end date) or a zero value period.
func (p Period) Intersect(q Interval[time.Time]) Interval[time.Time] {
	s := limit.Max(p.start, q.Min())
	e := limit.Min(p.end, q.Max())

	if !s.Before(e) {
		return Period{}
	}

	return Period{
		start: s,
		end:   e,
	}
}

// Encompass returns a new period encompassing both periods.
// Both periods must touch otherwise it returns a zero value period.
func (p Period) Encompass(q Interval[time.Time]) Interval[time.Time] {
	if !p.Touches(q) {
		return Period{}
	}

	s := limit.Min(p.start, q.Min())
	e := limit.Max(p.end, q.Max())

	if !s.Before(e) {
		return Period{}
	}

	return Period{
		start: s,
		end:   e,
	}
}

// Punch cuts q out of p and returns the remaining periods.
func (p Period) Punch(q Interval[time.Time]) (Interval[time.Time], Interval[time.Time]) {
	i := p.Intersect(q)
	if i.IsZero() {
		if p.Before(q) {
			return p, Period{}
		}
		return Period{}, p
	}

	l := Period{}
	r := Period{}

	if !p.start.Equal(i.Min()) {
		l = Period{start: p.start, end: i.Min()}
	}

	if !i.Max().Equal(p.end) {
		r = Period{start: i.Max(), end: p.end}
	}

	return l, r
}
