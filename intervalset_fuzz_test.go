package intervalset

import (
	"fmt"
	"math/rand"
	"testing"
)

// The tests below check IntervalSet against a brute-force model: a bitmask in
// which bit i is set when the value i is covered by the set. Intervals are
// treated as half-open [Min, Max), which is how the set merges and punches them.
//
// They exercise Range only: add & sub are generic over T, so Period runs the
// exact same code path.
//
// Intervals are always generated with a lower bound below or equal to their
// upper bound, as Add does not validate its input yet and an inverted interval
// breaks the set's ordering.

const fuzzDomain = 24

func maskOfRange(l, u int) uint32 {
	var m uint32
	for i := l; i < u; i++ {
		m |= 1 << uint(i)
	}
	return m
}

func maskOfSet(t *testing.T, s *IntervalSet[int]) uint32 {
	t.Helper()

	var m uint32
	for _, v := range s.AsSlice() {
		m |= maskOfRange(v.Min(), v.Max())
	}
	return m
}

// assertSetHolds reports whether the set covers the expected values and whether
// its intervals are still ordered, non-empty and neither overlapping nor adjacent.
func assertSetHolds(t *testing.T, s *IntervalSet[int], expected uint32, ops string) {
	t.Helper()

	if got := maskOfSet(t, s); got != expected {
		t.Fatalf("set should cover the expected values after %s, expected %024b, got %024b (%v)", ops, expected, got, s.AsSlice())
	}

	intervals := s.AsSlice()
	for i, v := range intervals {
		if v.Min() >= v.Max() {
			t.Fatalf("set should not contain an empty interval after %s, got %v at index %d", ops, intervals, i)
		}
		if i > 0 && v.Min() <= intervals[i-1].Max() {
			t.Fatalf("set should not contain overlapping or adjacent intervals after %s, got %v at index %d", ops, intervals, i)
		}
	}
}

func TestRangeSet_AddAndSubAgainstModel(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for i := 0; i < 20000; i++ {
		set := EmptySet[int]()
		var expected uint32
		ops := ""

		for o := 0; o < 6; o++ {
			l := r.Intn(fuzzDomain)
			u := min(l+r.Intn(4), fuzzDomain) // may be empty when l == u

			if r.Intn(2) == 0 {
				ops += fmt.Sprintf(" Add(%d,%d)", l, u)
				set.Add(NewRange[int](l, u))
				expected |= maskOfRange(l, u)
			} else {
				ops += fmt.Sprintf(" Sub(%d,%d)", l, u)
				set.Sub(NewRange[int](l, u))
				expected &^= maskOfRange(l, u)
			}

			assertSetHolds(t, set, expected, ops)
		}
	}
}

func TestRangeSet_OperationsAgainstModel(t *testing.T) {
	r := rand.New(rand.NewSource(7))

	genSet := func() (*IntervalSet[int], uint32) {
		set := EmptySet[int]()
		var m uint32

		for o := 0; o < 4; o++ {
			l := r.Intn(fuzzDomain)
			u := min(l+r.Intn(4), fuzzDomain)
			set.Add(NewRange[int](l, u))
			m |= maskOfRange(l, u)
		}
		return set, m
	}

	for i := 0; i < 5000; i++ {
		p, mp := genSet()
		q, mq := genSet()
		s, ms := genSet()

		assertSetHolds(t, Union(p, q), mp|mq, "Union")
		assertSetHolds(t, Intersection(p), mp, "Intersection of one set")
		assertSetHolds(t, Intersection(p, q), mp&mq, "Intersection of two sets")
		assertSetHolds(t, Intersection(p, q, s), mp&mq&ms, "Intersection of three sets")
		assertSetHolds(t, p.Difference(q), mp&^mq, "Difference")

		l := r.Intn(fuzzDomain)
		u := min(l+r.Intn(8), fuzzDomain)
		assertSetHolds(t, p.Complement(NewRange[int](l, u)), maskOfRange(l, u)&^mp, "Complement")
	}
}

// FuzzRangeSet_AddAndSub explores longer sequences of operations than the seeded
// tests above. Each operation is read from a pair of bytes: the first holds the
// operation and the lower bound, the second the distance to the upper bound.
//
// Run it with: go test -fuzz FuzzRangeSet_AddAndSub
func FuzzRangeSet_AddAndSub(f *testing.F) {
	f.Add([]byte{0x05, 0x02, 0x07, 0x03, 0x85, 0x04})
	f.Add([]byte{0x00, 0x00, 0x00, 0x00})
	f.Add([]byte{0x09, 0x01, 0x89, 0x01, 0x09, 0x02})

	f.Fuzz(func(t *testing.T, ops []byte) {
		if len(ops) > 64 {
			ops = ops[:64]
		}

		set := EmptySet[int]()
		var expected uint32

		for i := 0; i+1 < len(ops); i += 2 {
			l := int(ops[i]&0x7f) % fuzzDomain
			u := min(l+int(ops[i+1])%5, fuzzDomain)

			if ops[i]&0x80 == 0 {
				set.Add(NewRange[int](l, u))
				expected |= maskOfRange(l, u)
			} else {
				set.Sub(NewRange[int](l, u))
				expected &^= maskOfRange(l, u)
			}

			assertSetHolds(t, set, expected, "a fuzzed sequence of operations")
		}
	})
}
