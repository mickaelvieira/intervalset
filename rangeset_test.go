package intervalset

import (
	"fmt"
	"testing"
)

func genExpectedRangeSet(p []Interval[int]) *IntervalSet[int] {
	if p == nil {
		return EmptySet[int]()
	}
	return &IntervalSet[int]{intervals: p}
}

func TestRangeSet_AddValuesAndMaintainOrder(t *testing.T) {
	/*----------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10 |
	| (+) |                 |---|                  |
	| (+) |                                 |---|  |
	| (+) |     |---|                              |
	| (+) |                         |---|          |
	------------------------------------------------
	|  R  |     |---|       |---|   |---|   |---|  |
	----------------------------------------------*/
	set := EmptySet[int]().
		Add(
			NewRange[int](5, 6),
			NewRange[int](9, 10),
			NewRange[int](2, 3),
			NewRange[int](7, 8),
		)

	expected := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
		NewRange[int](5, 6),
		NewRange[int](7, 8),
		NewRange[int](9, 10),
	})

	if !set.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, set.AsSlice())
	}
}

func TestRangeSet_AddAndAdjoinValues(t *testing.T) {
	/*----------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10 |
	| (+) |                 |---|                  |
	| (+) |                                 |---|  |
	| (+) |     |---|                              |
	| (+) |                     |---|              |
	| (+) |                             |---|      |
	| (+) | |---|                                  |
	------------------------------------------------
	|  R  | |-------|       |-------|   |-------|  |
	----------------------------------------------*/
	set := EmptySet[int]().
		Add(
			NewRange[int](5, 6),
			NewRange[int](9, 10),
			NewRange[int](2, 3),
			NewRange[int](6, 7),
			NewRange[int](8, 9),
			NewRange[int](1, 2),
		)

	expected := genExpectedRangeSet([]Interval[int]{
		NewRange[int](1, 3),
		NewRange[int](5, 7),
		NewRange[int](8, 10),
	})

	if !set.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, set.AsSlice())
	}
}

func ExampleIntervalSet_Add_range() {
	s := EmptySet[int]().
		Add(
			NewRange[int](5, 6),
			NewRange[int](9, 10),
			NewRange[int](2, 3),
			NewRange[int](6, 7),
			NewRange[int](8, 9),
			NewRange[int](1, 2),
		)

	for _, p := range s.AsSlice() {
		fmt.Printf("%d - %d\n", p.Min(), p.Max())
	}

	// Output:
	// 1 - 3
	// 5 - 7
	// 8 - 10
}

func TestRangeSet_AddAndMergeOverlappingPeriods(t *testing.T) {
	/*----------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10 |
	| (+) |             |-------|                  |
	| (+) |                             |---|      |
	| (+) |     |---|                              |
	| (+) |                 |-------|              |
	| (+) |                                 |---|  |
	------------------------------------------------
	|  R  |     |---|   |-----------|   |--------| |
	----------------------------------------------*/
	set := EmptySet[int]().
		Add(
			NewRange[int](4, 6),
			NewRange[int](8, 9),
			NewRange[int](2, 3),
			NewRange[int](5, 7),
			NewRange[int](9, 10),
		)

	expected := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
		NewRange[int](4, 7),
		NewRange[int](8, 10),
	})

	if !set.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, set.AsSlice())
	}
}

func TestRangeSet_AddAndMergeMultipleOverlappingValues(t *testing.T) {
	/*-------------------------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
	| (+) |             |---|                                     |
	| (+) |     |---|                                             |
	| (+) |                     |---|                             |
	| (+) |                             |---|                     |
	| (+) |                                          |---|        |
	| (+) |             |-----------------------|                 |
	| (+) |                                               |---|   |
	---------------------------------------------------------------
	|  R  |     |---|   |-----------------------|    |--------|   |
	-------------------------------------------------------------*/
	set := EmptySet[int]().
		Add(
			NewRange[int](4, 5),
			NewRange[int](2, 3),
			NewRange[int](6, 7),
			NewRange[int](8, 9),
			NewRange[int](11, 12),
			NewRange[int](4, 10),
			NewRange[int](12, 13),
		)

	expected := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
		NewRange[int](4, 10),
		NewRange[int](11, 13),
	})

	if !set.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, set.AsSlice())
	}
}

func TestRangeSet_Sub(t *testing.T) {
	/*----------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10 |
	| (+) |                 |---|                  |
	| (+) |     |---|                              |
	| (+) |                                 |---|  |
	| (+) |                         |---|          |
	------------------------------------------------
	|  R  |     |---|       |---|   |---|   |---|  |
	------------------------------------------------
	| (-) |             |---------------|          |
	------------------------------------------------
	|  R  |     |---|                       |---|  |
	------------------------------------------------
	| (-) |                                 |---|  |
	------------------------------------------------
	|  R  |     |---|                              |
	------------------------------------------------
	| (-) | |-----------|                          |
	------------------------------------------------
	|  R  |                  ∅                     |
	----------------------------------------------*/
	set := EmptySet[int]().
		Add(
			NewRange[int](5, 6),
			NewRange[int](2, 3),
			NewRange[int](9, 10),
			NewRange[int](7, 8),
		)

	expected1 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
		NewRange[int](5, 6),
		NewRange[int](7, 8),
		NewRange[int](9, 10),
	})

	if !set.Equal(expected1) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected1, set.AsSlice())
	}

	set.Sub(NewRange[int](4, 8))

	expected2 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
		NewRange[int](9, 10),
	})

	if !set.Equal(expected2) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected2, set.AsSlice())
	}

	set.Sub(NewRange[int](9, 10))

	expected3 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
	})

	if !set.Equal(expected3) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected3, set.AsSlice())
	}

	// try to subtract a period that is no longer present
	set.Sub(NewRange[int](9, 10))

	if !set.Equal(expected3) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected3, set.AsSlice())
	}

	set.Sub(NewRange[int](1, 4))

	empty := genExpectedRangeSet(nil)

	if !set.Equal(empty) {
		t.Errorf("both sets should be equal, expected %v, got %v", empty, set.AsSlice())
	}

	// try to subtract a period from an empty set
	set.Sub(NewRange[int](1, 10))

	if !set.Equal(empty) {
		t.Errorf("both sets should be equal, expected %v, got %v", empty, set.AsSlice())
	}
}

func ExampleIntervalSet_Sub_range() {
	s := EmptySet[int]().
		Add(
			NewRange[int](1, 10),
		).
		Sub(
			NewRange[int](2, 4),
			NewRange[int](6, 8),
		)

	for _, p := range s.AsSlice() {
		fmt.Printf("%d - %d\n", p.Min(), p.Max())
	}

	// Output:
	// 1 - 2
	// 4 - 6
	// 8 - 10
}

func TestRangeSet_IsEmpty(t *testing.T) {
	set := EmptySet[int]()

	if !set.IsEmpty() {
		t.Errorf("the set should be empty")
	}

	set.Add(NewRange[int](5, 6))

	if set.IsEmpty() {
		t.Errorf("the set should not be empty")
	}

	set.Sub(NewRange[int](5, 6))

	if !set.IsEmpty() {
		t.Errorf("the set should be empty")
	}
}

func TestRangeSet_RemoveAndTruncateOverlappingPeriods(t *testing.T) {
	/*-------------------------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
	| (+) |             |-----------|                             |
	| (+) |                                 |---|                 |
	| (+) |                                          |---|        |
	| (+) |     |---|                                             |
	| (+) |                         |---|                         |
	---------------------------------------------------------------
	|  R  |     |---|   |---------------|   |---|    |---|        |
	---------------------------------------------------------------
	| (-) |                     |-------|                         |
	---------------------------------------------------------------
	|  R  |     |---|   |-------|           |---|    |---|        |
	---------------------------------------------------------------
	| (-) |                 |------------------------|            |
	---------------------------------------------------------------
	|  R  |     |---|   |---|                        |---|        |
	---------------------------------------------------------------
	| (-) |     |----------------------------------------|        |
	---------------------------------------------------------------
	|  R  |                         ∅                             |
	-------------------------------------------------------------*/
	set := EmptySet[int]().
		Add(
			NewRange[int](4, 7),
			NewRange[int](9, 10),
			NewRange[int](11, 12),
			NewRange[int](2, 3),
			NewRange[int](7, 8),
		)

	expected1 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
		NewRange[int](4, 8),
		NewRange[int](9, 10),
		NewRange[int](11, 12),
	})

	if !set.Equal(expected1) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected1, set.AsSlice())
	}

	set.Sub(NewRange[int](6, 8))

	expected2 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
		NewRange[int](4, 6),
		NewRange[int](9, 10),
		NewRange[int](11, 12),
	})

	if !set.Equal(expected2) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected2, set.AsSlice())
	}

	set.Sub(NewRange[int](5, 11))

	expected3 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](2, 3),
		NewRange[int](4, 5),
		NewRange[int](11, 12),
	})

	if !set.Equal(expected3) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected3, set.AsSlice())
	}

	set.Sub(NewRange[int](2, 12))

	expected4 := genExpectedRangeSet(nil)
	if !set.Equal(expected4) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected4, set.AsSlice())
	}
}

func TestRangeSet_Overlaps(t *testing.T) {
	var table = []struct {
		e *IntervalSet[int]
		p Range[int]
		s *IntervalSet[int]
	}{
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		| (+) |             |-----------|                             |
		| (+) |                                 |---|                 |
		| (+) |                                          |---|        |
		| (+) |     |---|                                             |
		| (+) |                         |---|                         |
		---------------------------------------------------------------
		|  R  |     |---|   |---------------|   |---|    |---|        |
		---------------------------------------------------------------
		|  U  |                         |----------------|            |
		---------------------------------------------------------------
		|  R  |                         |---|   |---|                 |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](7, 8),
				NewRange[int](9, 10),
			}),
			NewRange[int](7, 11),
			EmptySet[int]().Add(
				NewRange[int](4, 7),
				NewRange[int](9, 10),
				NewRange[int](11, 12),
				NewRange[int](2, 3),
				NewRange[int](7, 8),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		| (+) |     |---|                                             |
		---------------------------------------------------------------
		|  R  |     |---|                                             |
		---------------------------------------------------------------
		|  U  |     |---|                                             |
		---------------------------------------------------------------
		|  R  |     |---|                                             |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](2, 3),
			}),
			NewRange[int](2, 3),
			EmptySet[int]().Add(NewRange[int](2, 3)),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		| (+) |             |-----------|                             |
		| (+) |     |---|                                             |
		| (+) |                         |---|                         |
		---------------------------------------------------------------
		|  R  |     |---|   |---------------|                         |
		---------------------------------------------------------------
		|  U  |                         |---|                         |
		---------------------------------------------------------------
		|  R  |                         |---|                         |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](7, 8),
			}),
			NewRange[int](7, 8),
			EmptySet[int]().Add(
				NewRange[int](4, 7),
				NewRange[int](2, 3),
				NewRange[int](7, 8),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		| (+) |         |-------|                                     |
		| (+) |                         |---|                         |
		---------------------------------------------------------------
		|  R  |         |-------|       |---|                         |
		---------------------------------------------------------------
		|  U  |     |-------|                                         |
		---------------------------------------------------------------
		|  R  |         |---|                                         |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](3, 4),
			}),
			NewRange[int](2, 4),
			EmptySet[int]().Add(
				NewRange[int](3, 5),
				NewRange[int](7, 8),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		| (+) |                                 |---|                 |
		| (+) |                                          |---|        |
		| (+) |                         |---|                         |
		---------------------------------------------------------------
		|  R  |                         |---|   |---|    |---|        |
		---------------------------------------------------------------
		|  U  |     |---------------|                                 |
		---------------------------------------------------------------
		|  R  |                        ∅                              |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet(nil),
			NewRange[int](2, 6),
			EmptySet[int]().Add(
				NewRange[int](9, 10),
				NewRange[int](11, 12),
				NewRange[int](7, 8),
			),
		}, /*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		| (+) |             |---|                                     |
		| (+) |                     |---|                             |
		| (+) |     |---|                                             |
		---------------------------------------------------------------
		|  R  |     |---|   |---|   |---|                             |
		---------------------------------------------------------------
		|  R  |                             |-------|                 |
		---------------------------------------------------------------
		|  R  |                         ∅                             |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet(nil),
			NewRange[int](8, 10),
			EmptySet[int]().Add(
				NewRange[int](4, 5),
				NewRange[int](6, 7),
				NewRange[int](2, 3),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  R  |                                                       |
		---------------------------------------------------------------
		|  U  |                             |-------|                 |
		---------------------------------------------------------------
		|  R  |                         ∅                             |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet(nil),
			NewRange[int](8, 10),
			EmptySet[int](),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.s.Overlaps(tc.p)
			if !got.Equal(tc.e) {
				t.Errorf("expected both sets to be equal got %+v, expected %+v", got, tc.e)
			}
		})
	}
}

func ExampleIntervalSet_Overlaps_range() {
	s := EmptySet[int]().
		Add(
			NewRange[int](4, 6),
			NewRange[int](8, 10),
		)

	o := s.Overlaps(NewRange[int](5, 9))

	for _, p := range o.AsSlice() {
		fmt.Printf("%d - %d\n", p.Min(), p.Max())
	}

	// Output:
	// 5 - 6
	// 8 - 9
}

func TestRangeSet_Complement(t *testing.T) {
	var table = []struct {
		e *IntervalSet[int]
		r Range[int]
		s *IntervalSet[int]
	}{
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |---|       |-------|       |-------|         |----|  |
		---------------------------------------------------------------
		|  U  |     |------------------------------------|            |
		---------------------------------------------------------------
		|  R  |     |-------|       |-------|       |----|            |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](2, 4),
				NewRange[int](6, 8),
				NewRange[int](10, 11),
			}),
			NewRange[int](2, 11),
			EmptySet[int]().Add(
				NewRange[int](1, 2),
				NewRange[int](4, 6),
				NewRange[int](8, 10),
				NewRange[int](12, 13),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  |                                                       |
		---------------------------------------------------------------
		|  U  |     |------------------------------------|            |
		---------------------------------------------------------------
		|  R  |     |------------------------------------|            |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](2, 11),
			}),
			NewRange[int](2, 11),
			EmptySet[int](),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  |     |---|                                |----|       |
		---------------------------------------------------------------
		|  U  |             |-------------------|                     |
		---------------------------------------------------------------
		|  R  |             |-------------------|                     |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](4, 9),
			}),
			NewRange[int](4, 9),
			EmptySet[int]().Add(
				NewRange[int](2, 3),
				NewRange[int](11, 12),
			),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.s.Complement(tc.r)
			if !got.Equal(tc.e) {
				t.Errorf("expected both sets to be equal got %+v, expected %+v", got, tc.e)
			}
		})
	}
}

func ExampleIntervalSet_Complement_range() {
	s := EmptySet[int]().
		Add(
			NewRange[int](1, 2),
			NewRange[int](4, 5),
			NewRange[int](7, 8),
		)

	c := s.Complement(NewRange[int](1, 10))

	for _, p := range c.AsSlice() {
		fmt.Printf("%d - %d\n", p.Min(), p.Max())
	}

	// Output:
	// 2 - 4
	// 5 - 7
	// 8 - 10
}

func TestRangeSet_IsSubset(t *testing.T) {
	var table = []struct {
		e  bool
		s1 *IntervalSet[int]
		s2 *IntervalSet[int]
	}{
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  |                         ∅                             |
		|  S  |     |-----------|                    |----|           |
		-------------------------------------------------------------*/
		{
			false,
			EmptySet[int](),
			EmptySet[int]().Add(
				NewRange[int](2, 5),
				NewRange[int](10, 11),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |-------|       |-------|       |--------------|      |
		|  S  |                         ∅                             |
		-------------------------------------------------------------*/
		{
			true,
			EmptySet[int]().Add(
				NewRange[int](1, 3),
				NewRange[int](5, 7),
				NewRange[int](9, 12),
			),
			EmptySet[int](),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |-------|       |-------|       |--------------|      |
		|  S  |     |-----------|                    |----|           |
		-------------------------------------------------------------*/
		{
			false,
			EmptySet[int]().Add(
				NewRange[int](1, 3),
				NewRange[int](5, 7),
				NewRange[int](9, 12),
			),
			EmptySet[int]().Add(
				NewRange[int](2, 5),
				NewRange[int](10, 11),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |-------|       |-------|       |--------------|      |
		|  S  |     |---|                            |----|           |
		-------------------------------------------------------------*/
		{
			true,
			EmptySet[int]().Add(
				NewRange[int](1, 3),
				NewRange[int](5, 7),
				NewRange[int](9, 12),
			),
			EmptySet[int]().Add(
				NewRange[int](2, 3),
				NewRange[int](10, 11),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |-------|       |-------|       |--------------|      |
		|  S  | |-------|       |-------|       |--------------|      |
		-------------------------------------------------------------*/
		{
			true,
			EmptySet[int]().Add(
				NewRange[int](1, 3),
				NewRange[int](5, 7),
				NewRange[int](9, 12),
			),
			EmptySet[int]().Add(
				NewRange[int](1, 3),
				NewRange[int](5, 7),
				NewRange[int](9, 12),
			),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.s1.IsSubset(tc.s2)
			if got != tc.e {
				if tc.e {
					t.Errorf("expected s2 to be a subset of s1: s1 %+v, s2 %+v", tc.s1, tc.s2)
				} else {
					t.Errorf("expected s2 to not be a subset of s1: s1 %+v, s2 %+v", tc.s1, tc.s2)
				}
			}
		})
	}
}

func TestRangeSet_Difference(t *testing.T) {
	var table = []struct {
		e  *IntervalSet[int]
		s1 *IntervalSet[int]
		s2 *IntervalSet[int]
	}{
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |-------|           |---|       |--------------|      |
		|  S  |     |----------|                     |----|           |
		---------------------------------------------------------------
		|  R  | |---|               |---|       |----|    |----|      |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](1, 2),
				NewRange[int](6, 7),
				NewRange[int](9, 10),
				NewRange[int](11, 12),
			}),
			EmptySet[int]().Add(
				NewRange[int](1, 3),
				NewRange[int](6, 7),
				NewRange[int](9, 12),
			),
			EmptySet[int]().Add(
				NewRange[int](2, 5),
				NewRange[int](10, 11),
			),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  |     |-------|                       |----|            |
		|  S  |     |----------|                    |----|            |
		---------------------------------------------------------------
		|  R  |                         ∅                             |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet(nil),
			EmptySet[int]().Add(
				NewRange[int](2, 4),
				NewRange[int](10, 11),
			),
			EmptySet[int]().Add(
				NewRange[int](2, 5),
				NewRange[int](10, 11),
			),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.s1.Difference(tc.s2)
			if !got.Equal(tc.e) {
				t.Errorf("expected both sets to be equal got %+v, expected %+v", got, tc.e)
			}
		})
	}
}

func ExampleIntervalSet_Difference_range() {
	a := EmptySet[int]().
		Add(
			NewRange[int](1, 2),
			NewRange[int](4, 5),
			NewRange[int](7, 8),
		)

	b := EmptySet[int]().
		Add(
			NewRange[int](1, 6),
		)

	// ab = a – b
	ab := a.Difference(b)

	for _, p := range ab.AsSlice() {
		fmt.Printf("%d - %d\n", p.Min(), p.Max())
	}

	// Output:
	// 7 - 8
}

func TestUnion_range(t *testing.T) {
	/*-------------------------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
	---------------------------------------------------------------
	|  S  | |---|                                         |----|  |
	|  S  |     |----------|                    |----|            |
	|  S  |             |-------|       |-------|         |----|  |
	---------------------------------------------------------------
	|  R  | |-------------------|       |------------|    |----|  |
	-------------------------------------------------------------*/
	set1 := EmptySet[int]()
	set1.Add(
		NewRange[int](1, 2),
		NewRange[int](12, 13),
	)

	set2 := EmptySet[int]()
	set2.Add(
		NewRange[int](2, 5),
		NewRange[int](10, 11),
	)

	set3 := EmptySet[int]()
	set3.Add(
		NewRange[int](4, 6),
		NewRange[int](8, 10),
		NewRange[int](12, 13),
	)

	union := Union(set1, set2, set3)

	expected := genExpectedRangeSet([]Interval[int]{
		NewRange[int](1, 6),
		NewRange[int](8, 11),
		NewRange[int](12, 13),
	})

	if !union.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, union)
	}
}

func ExampleUnion_range() {
	a := EmptySet[int]().
		Add(
			NewRange[int](1, 2),
			NewRange[int](4, 5),
		)

	b := EmptySet[int]().
		Add(
			NewRange[int](2, 3),
			NewRange[int](6, 7),
		)

	// ab = a ∪ b
	ab := Union(a, b)

	for _, p := range ab.AsSlice() {
		fmt.Printf("%d - %d\n", p.Min(), p.Max())
	}

	// Output:
	// 1 - 3
	// 4 - 5
	// 6 - 7
}

func TestIntersection_range(t *testing.T) {
	var table = []struct {
		e *IntervalSet[int]
		s []*IntervalSet[int]
	}{
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |-----------|                       |---------------| |
		|  S  |     |-----------|               |-------------|       |
		|  S  | |---|           |---|   |-------|                     |
		|  S  |     |---------------------------|   |----------|      |
		---------------------------------------------------------------
		|  R  |                         ∅                             |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet(nil),
			[]*IntervalSet[int]{
				EmptySet[int]().
					Add(
						NewRange[int](1, 4),
						NewRange[int](10, 13),
					),
				EmptySet[int]().
					Add(
						NewRange[int](2, 5),
						NewRange[int](9, 12),
					),
				EmptySet[int]().
					Add(
						NewRange[int](1, 2),
						NewRange[int](5, 6),
						NewRange[int](7, 9),
					),
				EmptySet[int]().
					Add(
						NewRange[int](2, 9),
						NewRange[int](10, 12),
					),
			},
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |-----------|                       |---------------| |
		|  S  |     |-----------|               |-------------|       |
		|  S  | |-------|           |---|   |------------|            |
		|  S  |     |---------------------------|   |----------|      |
		---------------------------------------------------------------
		|  R  |     |---|                           |----|            |
		-------------------------------------------------------------*/
		{
			genExpectedRangeSet([]Interval[int]{
				NewRange[int](2, 3),
				NewRange[int](10, 11),
			}),
			[]*IntervalSet[int]{
				EmptySet[int]().
					Add(
						NewRange[int](1, 4),
						NewRange[int](10, 13),
					),
				EmptySet[int]().
					Add(
						NewRange[int](2, 5),
						NewRange[int](9, 12),
					),
				EmptySet[int]().
					Add(
						NewRange[int](1, 3),
						NewRange[int](6, 7),
						NewRange[int](8, 11),
					),
				EmptySet[int]().
					Add(
						NewRange[int](2, 9),
						NewRange[int](10, 12),
					),
			},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := Intersection(tc.s...)
			if !got.Equal(tc.e) {
				t.Errorf("expected s1 to be equal to s2: s1 %+v, s2 %+v", got, tc.e)
			}
		})
	}
}

func ExampleIntersection_range() {
	a := EmptySet[int]().
		Add(
			NewRange[int](1, 2),
			NewRange[int](4, 5),
			NewRange[int](7, 8),
		)

	b := EmptySet[int]().
		Add(
			NewRange[int](1, 6),
		)

	// ab = a ∩ b
	ab := Intersection(a, b)

	for _, p := range ab.AsSlice() {
		fmt.Printf("%d - %d\n", p.Min(), p.Max())
	}

	// Output:
	// 1 - 2
	// 4 - 5
}

func TestRangeSet_Equal(t *testing.T) {
	table := []struct {
		e  bool
		s1 *IntervalSet[int]
		s2 *IntervalSet[int]
	}{
		{
			false,
			EmptySet[int]().Add(NewRange[int](1, 3)),
			EmptySet[int](),
		},
		{
			false,
			EmptySet[int]().Add(NewRange[int](1, 3)),
			EmptySet[int]().Add(NewRange[int](1, 4)),
		},
		{
			true,
			EmptySet[int]().Add(NewRange[int](1, 3)),
			EmptySet[int]().Add(NewRange[int](1, 3)),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.s1.Equal(tc.s2)
			if got != tc.e {
				if tc.e {
					t.Errorf("expected s1 to be equal to s2, s1 %+v, s2 %+v", tc.s1, tc.s2)
				} else {
					t.Errorf("expected s1 to not be equal to s2, s1 %+v, s2 %+v", tc.s1, tc.s2)
				}
			}
		})
	}
}

func TestRangeSet_Iter(t *testing.T) {
	s1 := EmptySet[int]().Add(
		NewRange[int](1, 2),
		NewRange[int](3, 4),
		NewRange[int](5, 6),
		NewRange[int](7, 8),
		NewRange[int](9, 10),
		NewRange[int](11, 12),
	)

	s2 := EmptySet[int]()

	s1.Iter(func(i Interval[int]) bool {
		s2.Add(i)
		return true
	})

	if !s1.Equal(s2) {
		t.Errorf("expected both sets to be equal, s1 %+v, s2 %+v", s1, s2)
	}

	s3 := EmptySet[int]()
	e3 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](1, 2),
		NewRange[int](3, 4),
		NewRange[int](5, 6),
	})

	c := 0
	s1.Iter(func(i Interval[int]) bool {
		s3.Add(i)
		c++

		return c < 3
	})

	if !s3.Equal(e3) {
		t.Errorf("expected both sets to be equal, s3 %+v, e3 %+v", s3, e3)
	}
}

func TestRangeSet_IterBetween(t *testing.T) {
	s1 := EmptySet[int]().Add(
		NewRange[int](1, 4),
		NewRange[int](5, 6),
		NewRange[int](7, 10),
		NewRange[int](11, 12),
	)

	p2 := NewRange[int](3, 9)

	s2 := EmptySet[int]()
	e2 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](3, 4),
		NewRange[int](5, 6),
		NewRange[int](7, 9),
	})

	s1.IterBetween(p2, func(i Interval[int]) bool {
		s2.Add(i)
		return true
	})

	if !s2.Equal(e2) {
		t.Errorf("expected both sets to be equal, got %+v, want %+v", s2, e2)
	}

	s3 := EmptySet[int]()
	e3 := genExpectedRangeSet([]Interval[int]{
		NewRange[int](3, 4),
		NewRange[int](5, 6),
	})

	c := 0
	s1.IterBetween(p2, func(i Interval[int]) bool {
		s3.Add(i)
		c++

		return c < 2
	})

	if !s3.Equal(e3) {
		t.Errorf("expected both sets to be equal, s3 %+v, e3 %+v", s3, e3)
	}
}
