package intervalset

import (
	"fmt"
	"testing"
	"time"
)

func genExpectedPeriodSet(p []Interval[time.Time]) *IntervalSet[time.Time] {
	if p == nil {
		return EmptySet[time.Time]()
	}
	return &IntervalSet[time.Time]{intervals: p}
}

func TestPeriodSet_AddPeriodsAndMaintainOrder(t *testing.T) {
	/*----------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10 |
	| (+) |                 |---|                  |
	| (+) |                                 |---|  |
	| (+) |     |---|                              |
	| (+) |                         |---|          |
	------------------------------------------------
	|  R  |     |---|       |---|   |---|   |---|  |
	----------------------------------------------*/
	set := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
		)

	expected := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, set.AsSlice())
	}
}

func TestPeriodSet_AddAndAdjoinPeriods(t *testing.T) {
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
	set := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			),
		)

	expected := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, set.AsSlice())
	}
}

func ExampleIntervalSet_Add_period() {
	s := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			),
		)

	for _, p := range s.AsSlice() {
		fmt.Printf("%s - %s\n", p.Min(), p.Max())
	}

	// Output:
	// 2023-12-01 00:00:00 +0000 UTC - 2023-12-03 00:00:00 +0000 UTC
	// 2023-12-05 00:00:00 +0000 UTC - 2023-12-07 00:00:00 +0000 UTC
	// 2023-12-08 00:00:00 +0000 UTC - 2023-12-10 00:00:00 +0000 UTC
}

func TestPeriodSet_AddAndMergeOverlappingPeriods(t *testing.T) {
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
	set := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
		)

	expected := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, set.AsSlice())
	}
}

func TestPeriodSet_AddAndMergeMultipleOverlappingPeriods(t *testing.T) {
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
	set := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 12, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
			),
		)

	expected := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, set.AsSlice())
	}
}

func TestPeriodSet_Sub(t *testing.T) {
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
	set := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
		)

	expected1 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected1) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected1, set.AsSlice())
	}

	set.Sub(NewPeriod(
		time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
	))

	expected2 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected2) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected2, set.AsSlice())
	}

	set.Sub(NewPeriod(
		time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
	))

	expected3 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected3) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected3, set.AsSlice())
	}

	// try to subtract a period that is no longer present
	set.Sub(NewPeriod(
		time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
	))

	if !set.Equal(expected3) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected3, set.AsSlice())
	}

	set.Sub(NewPeriod(
		time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
	))

	empty := genExpectedPeriodSet(nil)

	if !set.Equal(empty) {
		t.Errorf("both sets should be equal, expected %v, got %v", empty, set.AsSlice())
	}

	// try to subtract a period from an empty set
	set.Sub(NewPeriod(
		time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
	))

	if !set.Equal(empty) {
		t.Errorf("both sets should be equal, expected %v, got %v", empty, set.AsSlice())
	}
}

func ExampleIntervalSet_Sub_period() {
	s := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
		).
		Sub(
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
		)

	for _, p := range s.AsSlice() {
		fmt.Printf("%s - %s\n", p.Min(), p.Max())
	}

	// Output:
	// 2023-12-01 00:00:00 +0000 UTC - 2023-12-02 00:00:00 +0000 UTC
	// 2023-12-04 00:00:00 +0000 UTC - 2023-12-06 00:00:00 +0000 UTC
	// 2023-12-08 00:00:00 +0000 UTC - 2023-12-10 00:00:00 +0000 UTC
}

func TestPeriodSet_IsEmpty(t *testing.T) {
	set := EmptySet[time.Time]()

	if !set.IsEmpty() {
		t.Errorf("the set should be empty")
	}

	set.Add(NewPeriod(
		time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
	))

	if set.IsEmpty() {
		t.Errorf("the set should not be empty")
	}

	set.Sub(NewPeriod(
		time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
	))

	if !set.IsEmpty() {
		t.Errorf("the set should be empty")
	}
}

func TestPeriodSet_RemoveAndTruncateOverlappingPeriods(t *testing.T) {
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
	set := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
		)

	expected1 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected1) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected1, set.AsSlice())
	}

	set.Sub(NewPeriod(
		time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
	))

	expected2 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected2) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected2, set.AsSlice())
	}

	set.Sub(NewPeriod(
		time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
	))

	expected3 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
		),
	})

	if !set.Equal(expected3) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected3, set.AsSlice())
	}

	set.Sub(NewPeriod(
		time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
	))

	expected4 := genExpectedPeriodSet(nil)
	if !set.Equal(expected4) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected4, set.AsSlice())
	}
}

func BenchmarkOverlaps(b *testing.B) {
	b.StopTimer()

	set := EmptySet[time.Time]()

	oneDay := time.Second * 60 * 60 * 24

	s1 := time.Date(2024, time.January, 1, 8, 0, 0, 0, time.UTC)
	e1 := time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)

	s2 := time.Date(2024, time.January, 1, 14, 0, 0, 0, time.UTC)
	e2 := time.Date(2024, time.January, 1, 18, 0, 0, 0, time.UTC)

	s3 := time.Date(2024, time.January, 1, 22, 0, 0, 0, time.UTC)
	e3 := time.Date(2024, time.January, 2, 2, 0, 0, 0, time.UTC)

	i := 1

	for i <= 365 {
		set.Add(NewPeriod(s1, e1))
		set.Add(NewPeriod(s2, e2))
		set.Add(NewPeriod(s3, e3))

		s1 = s1.Add(oneDay)
		e1 = e1.Add(oneDay)
		s2 = s2.Add(oneDay)
		e2 = e2.Add(oneDay)
		s3 = s3.Add(oneDay)
		e3 = e3.Add(oneDay)

		i++
	}

	s := time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC)
	e := time.Date(2021, time.June, 2, 0, 0, 0, 0, time.UTC)

	b.StartTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		set.Overlaps(NewPeriod(
			s,
			e,
		))
	}
}

func TestPeriodSet_Overlaps(t *testing.T) {
	var table = []struct {
		e *IntervalSet[time.Time]
		p Period[time.Time]
		s *IntervalSet[time.Time]
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
				),
			}),
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				),
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
			}),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time]().Add(NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			)),
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				),
			}),
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				),
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				),
			}),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				),
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
			genExpectedPeriodSet(nil),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				),
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
			genExpectedPeriodSet(nil),
			NewPeriod(
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
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
			genExpectedPeriodSet(nil),
			NewPeriod(
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time](),
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

func ExampleIntervalSet_Overlaps_period() {
	s := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			),
		)

	o := s.Overlaps(
		NewPeriod(
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
		),
	)

	for _, p := range o.AsSlice() {
		fmt.Printf("%s - %s\n", p.Min(), p.Max())
	}

	// Output:
	// 2023-12-05 00:00:00 +0000 UTC - 2023-12-06 00:00:00 +0000 UTC
	// 2023-12-08 00:00:00 +0000 UTC - 2023-12-09 00:00:00 +0000 UTC
}

func TestPeriodSet_Complement(t *testing.T) {
	var table = []struct {
		e *IntervalSet[time.Time]
		p Period[time.Time]
		s *IntervalSet[time.Time]
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
			}),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
				),
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
			}),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time](),
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
				),
			}),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
			),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.s.Complement(tc.p)
			if !got.Equal(tc.e) {
				t.Errorf("expected both sets to be equal got %+v, expected %+v", got, tc.e)
			}
		})
	}
}

func ExampleIntervalSet_Complement_period() {
	s := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
		)

	c := s.Complement(
		NewPeriod(
			time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
	)

	for _, p := range c.AsSlice() {
		fmt.Printf("%s - %s\n", p.Min(), p.Max())
	}

	// Output:
	// 2023-12-02 00:00:00 +0000 UTC - 2023-12-04 00:00:00 +0000 UTC
	// 2023-12-05 00:00:00 +0000 UTC - 2023-12-07 00:00:00 +0000 UTC
	// 2023-12-08 00:00:00 +0000 UTC - 2023-12-10 00:00:00 +0000 UTC
}

func TestPeriodSet_IsSubset(t *testing.T) {
	var table = []struct {
		e  bool
		s1 *IntervalSet[time.Time]
		s2 *IntervalSet[time.Time]
	}{
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  |                         ∅                             |
		|  S  |     |-----------|                    |----|           |
		-------------------------------------------------------------*/
		{
			false,
			EmptySet[time.Time](),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
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
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time](),
		},
		/*-------------------------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
		---------------------------------------------------------------
		|  S  | |-------|       |-------|       |--------------|      |
		|  S  |     |-----------|                    |----|           |
		-------------------------------------------------------------*/
		{
			false,
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
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
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
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
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
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

func TestPeriodSet_Difference(t *testing.T) {
	var table = []struct {
		e  *IntervalSet[time.Time]
		s1 *IntervalSet[time.Time]
		s2 *IntervalSet[time.Time]
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
			}),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
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
			genExpectedPeriodSet(nil),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
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

func ExampleIntervalSet_Difference_period() {
	a := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
		)

	b := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
		)

	// ab = a – b
	ab := a.Difference(b)

	for _, p := range ab.AsSlice() {
		fmt.Printf("%s - %s\n", p.Min(), p.Max())
	}

	// Output:
	// 2023-12-07 00:00:00 +0000 UTC - 2023-12-08 00:00:00 +0000 UTC
}

func TestUnion(t *testing.T) {
	/*-------------------------------------------------------------
	|  T  | 1   2   3   4   5   6   7   8   9   10   11   12   13 |
	---------------------------------------------------------------
	|  S  | |---|                                         |----|  |
	|  S  |     |----------|                    |----|            |
	|  S  |             |-------|       |-------|         |----|  |
	---------------------------------------------------------------
	|  R  | |-------------------|       |------------|    |----|  |
	-------------------------------------------------------------*/
	set1 := EmptySet[time.Time]()
	set1.Add(
		NewPeriod(
			time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
		),
	)

	set2 := EmptySet[time.Time]()
	set2.Add(
		NewPeriod(
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
		),
	)

	set3 := EmptySet[time.Time]()
	set3.Add(
		NewPeriod(
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
		),
	)

	union := Union(set1, set2, set3)

	expected := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
		),
	})

	if !union.Equal(expected) {
		t.Errorf("both sets should be equal, expected %v, got %v", expected, union)
	}
}

func ExampleUnion_period() {
	a := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
		)

	b := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
		)

	// ab = a ∪ b
	ab := Union(a, b)

	for _, p := range ab.AsSlice() {
		fmt.Printf("%s - %s\n", p.Min(), p.Max())
	}

	// Output:
	// 2023-12-01 00:00:00 +0000 UTC - 2023-12-03 00:00:00 +0000 UTC
	// 2023-12-04 00:00:00 +0000 UTC - 2023-12-05 00:00:00 +0000 UTC
	// 2023-12-06 00:00:00 +0000 UTC - 2023-12-07 00:00:00 +0000 UTC
}

func TestPeriodSet_Intersection(t *testing.T) {
	var table = []struct {
		e *IntervalSet[time.Time]
		s []*IntervalSet[time.Time]
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
			genExpectedPeriodSet(nil),
			[]*IntervalSet[time.Time]{
				EmptySet[time.Time]().
					Add(
						NewPeriod(
							time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
						),
					),
				EmptySet[time.Time]().
					Add(
						NewPeriod(
							time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
						),
					),
				EmptySet[time.Time]().
					Add(
						NewPeriod(
							time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
						),
					),
				EmptySet[time.Time]().
					Add(
						NewPeriod(
							time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
						),
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
			genExpectedPeriodSet([]Interval[time.Time]{
				NewPeriod(
					time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
				NewPeriod(
					time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
				),
			}),
			[]*IntervalSet[time.Time]{
				EmptySet[time.Time]().
					Add(
						NewPeriod(
							time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
						),
					),
				EmptySet[time.Time]().
					Add(
						NewPeriod(
							time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
						),
					),
				EmptySet[time.Time]().
					Add(
						NewPeriod(
							time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
						),
					),
				EmptySet[time.Time]().
					Add(
						NewPeriod(
							time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
						),
						NewPeriod(
							time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
							time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
						),
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

func ExampleIntersection_period() {
	a := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
		)

	b := EmptySet[time.Time]().
		Add(
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
		)

	// ab = a ∩ b
	ab := Intersection(a, b)

	for _, p := range ab.AsSlice() {
		fmt.Printf("%s - %s\n", p.Min(), p.Max())
	}

	// Output:
	// 2023-12-01 00:00:00 +0000 UTC - 2023-12-02 00:00:00 +0000 UTC
	// 2023-12-04 00:00:00 +0000 UTC - 2023-12-05 00:00:00 +0000 UTC
}

func TestPeriodSet_Equal(t *testing.T) {
	table := []struct {
		e  bool
		s1 *IntervalSet[time.Time]
		s2 *IntervalSet[time.Time]
	}{
		{
			false,
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time](),
		},
		{
			false,
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				),
			),
		},
		{
			true,
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
			),
			EmptySet[time.Time]().Add(
				NewPeriod(
					time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				),
			),
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

func TestPeriodSet_Iter(t *testing.T) {
	s1 := EmptySet[time.Time]().Add(
		NewPeriod(
			time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
		),
	)

	s2 := EmptySet[time.Time]()

	s1.Iter(func(i Interval[time.Time]) bool {
		s2.Add(i)
		return true
	})

	if !s1.Equal(s2) {
		t.Errorf("expected both sets to be equal, s1 %+v, s2 %+v", s1, s2)
	}

	s3 := EmptySet[time.Time]()
	e3 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
		),
	})

	c := 0
	s1.Iter(func(i Interval[time.Time]) bool {
		s3.Add(i)
		c++

		return c < 2
	})

	if !s3.Equal(e3) {
		t.Errorf("expected both sets to be equal, s3 %+v, e3 %+v", s3, e3)
	}
}

func TestPeriodSet_IterBetween(t *testing.T) {
	s1 := EmptySet[time.Time]().Add(
		NewPeriod(
			time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 12, 0, 0, 0, 0, time.UTC),
		),
	)

	p2 := NewPeriod(
		time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
	)

	s2 := EmptySet[time.Time]()
	e2 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
		),
	})

	s1.IterBetween(p2, func(i Interval[time.Time]) bool {
		s2.Add(i)
		return true
	})

	if !s2.Equal(e2) {
		t.Errorf("expected both sets to be equal, got %+v, want %+v", s2, e2)
	}

	s3 := EmptySet[time.Time]()
	e3 := genExpectedPeriodSet([]Interval[time.Time]{
		NewPeriod(
			time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
		),
		NewPeriod(
			time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
		),
	})

	c := 0
	s1.IterBetween(p2, func(i Interval[time.Time]) bool {
		s3.Add(i)
		c++

		return c < 2
	})

	if !s3.Equal(e3) {
		t.Errorf("expected both sets to be equal, s3 %+v, e3 %+v", s3, e3)
	}
}
