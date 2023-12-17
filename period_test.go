package intervalset

import (
	"fmt"
	"testing"
	"time"
)

func TestPeriod_IsEmpty(t *testing.T) {
	t1 := Period[time.Time]{
		start: time.Date(2023, time.December, 10, 12, 50, 0, 0, time.UTC),
		end:   time.Date(2023, time.December, 10, 12, 50, 0, 0, time.UTC),
	}

	if !t1.IsEmpty() {
		t.Errorf("period should be empty %+v", t1)
	}

	t2 := Period[time.Time]{
		start: time.Date(2023, time.December, 10, 12, 50, 0, 0, time.UTC),
		end:   time.Date(2023, time.December, 10, 12, 50, 1, 0, time.UTC),
	}

	if t2.IsEmpty() {
		t.Errorf("period should not be empty %+v", t2)
	}
}

func TestPeriod_IsValid(t *testing.T) {
	t1 := Period[time.Time]{
		start: time.Date(2023, time.December, 10, 12, 50, 0, 0, time.UTC),
		end:   time.Date(2023, time.December, 10, 12, 50, 0, 0, time.UTC),
	}

	if !t1.IsValid() {
		t.Errorf("period should be valid %+v", t1)
	}

	t2 := Period[time.Time]{
		start: time.Date(2023, time.December, 10, 12, 50, 0, 0, time.UTC),
		end:   time.Date(2023, time.December, 10, 12, 50, 1, 0, time.UTC),
	}

	if !t2.IsValid() {
		t.Errorf("period should be valid %+v", t2)
	}

	t3 := Period[time.Time]{
		start: time.Date(2023, time.December, 10, 12, 50, 1, 0, time.UTC),
		end:   time.Date(2023, time.December, 10, 12, 50, 0, 0, time.UTC),
	}

	if t3.IsValid() {
		t.Errorf("period should not be valid %+v", t3)
	}
}

func TestPeriod_Equal(t *testing.T) {
	var table = []struct {
		e bool
		i []time.Time
	}{
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t1 := Period[time.Time]{
				start: tc.i[0],
				end:   tc.i[1],
			}
			t2 := Period[time.Time]{
				start: tc.i[2],
				end:   tc.i[3],
			}

			if t1.Equal(t2) != tc.e {
				if tc.e {
					t.Errorf("expected t1 to be equal to t2: t1 %+v, t2 %+v", t1, t2)
				} else {
					t.Errorf("expected t1 to not be equal to t2: t1 %+v, t2 %+v", t1, t2)
				}
			}
		})
	}
}

func TestPeriod_Before(t *testing.T) {
	var table = []struct {
		e bool
		i []time.Time
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |-------|              |
		| (2) |         |---|                          |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |         |-------|                      |
		| (2) |             |-------|                  |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t1 := Period[time.Time]{
				start: tc.i[0],
				end:   tc.i[1],
			}
			t2 := Period[time.Time]{
				start: tc.i[2],
				end:   tc.i[3],
			}

			if t1.Before(t2) != tc.e {
				if tc.e {
					t.Errorf("expected t1 to end before the beginning of t2: t1 %+v, t2 %+v", t1, t2)
				} else {
					t.Errorf("expected t1 to not end before the beginning of t2: t1 %+v, t2 %+v", t1, t2)
				}
			}
		})
	}
}

func TestPeriod_After(t *testing.T) {
	var table = []struct {
		e bool
		i []time.Time
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |-------|                  |
		| (2) |         |-------|                      |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |-------|              |
		| (2) |     |-------|                          |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t1 := Period[time.Time]{
				start: tc.i[0],
				end:   tc.i[1],
			}
			t2 := Period[time.Time]{
				start: tc.i[2],
				end:   tc.i[3],
			}

			if t1.After(t2) != tc.e {
				if tc.e {
					t.Errorf("expected t1 to start after the end of t2: t1 %+v, t2 %+v", t1, t2)
				} else {
					t.Errorf("expected t1 to not start after the end of t2: t1 %+v, t2 %+v", t1, t2)
				}
			}
		})
	}
}

func TestPeriod_Overlap(t *testing.T) {
	var table = []struct {
		e bool
		i []time.Time
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |         |-------|                      |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |-------|                  |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |-------|              |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                     |-------|          |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                         |-------|      |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t1 := Period[time.Time]{
				start: tc.i[0],
				end:   tc.i[1],
			}
			t2 := Period[time.Time]{
				start: tc.i[2],
				end:   tc.i[3],
			}

			got := t1.Overlaps(t2)
			if got != tc.e {
				if tc.e {
					t.Errorf("expected t1 to overlap t2: t1 %+v, t2 %+v, %t", t1, t2, got)
				} else {
					t.Errorf("expected t1 to not overlap t2: t1 %+v, t2 %+v, %t", t1, t2, got)
				}
			}
		})
	}
}

func TestPeriod_Contains(t *testing.T) {
	var table = []struct {
		e bool
		i []time.Time
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (2) | |-------|                              |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |     |-------|                          |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |         |-------|                      |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |             |-------|                  |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |                 |-------|              |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |                     |-------|          |
		----------------------------------------------*/
		{
			true,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |                         |-------|      |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |                             |-------|  |
		----------------------------------------------*/
		{
			false,
			[]time.Time{
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t1 := Period[time.Time]{
				start: tc.i[0],
				end:   tc.i[1],
			}
			t2 := Period[time.Time]{
				start: tc.i[2],
				end:   tc.i[3],
			}
			got := t1.Contains(t2)
			if got != tc.e {
				if tc.e {
					t.Errorf("expected t2 to be a subset of t1: t1 %+v, t2 %+v", t1, t2)
				} else {
					t.Errorf("expected t2 to not be a subset of t1: t1 %+v, t2 %+v", t1, t2)
				}
			}
		})
	}
}

func TestPeriod_Intersect(t *testing.T) {
	var table = []struct {
		p1 Period[time.Time]
		p2 Period[time.Time]
		e  Period[time.Time]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |-------|              |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |---------------|                  |
		| (2) |             |---|                      |
		------------------------------------------------
		|  R  |             |---|                      |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{
				start: time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |         |-------|                      |
		| (2) |     |-------|                          |
		------------------------------------------------
		|  R  |         |---|                          |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{
				start: time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |         |-------|                      |
		------------------------------------------------
		|  R  |         |---|                          |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{
				start: time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |             |---|                      |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.p1.Intersect(tc.p2)

			if !tc.e.Equal(got) {
				t.Errorf("both periods should be equal, expected %v, got %v", tc.e, got)
			}
		})
	}
}

func ExamplePeriod_Intersect() {
	a := NewPeriod(
		time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
	)

	b := NewPeriod(
		time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
	)

	r := a.Intersect(b)

	fmt.Printf("%s - %s\n", r.Min(), r.Max())

	// Output:
	// 2023-12-03 00:00:00 +0000 UTC - 2023-12-05 00:00:00 +0000 UTC
}

func TestPeriod_Punch(t *testing.T) {
	var table = []struct {
		p1 Period[time.Time]
		p2 Period[time.Time]
		e1 Period[time.Time]
		e2 Period[time.Time]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------------------|              |
		| (2) |         |-------|                      |
		------------------------------------------------
		|  R  |     |---|       |-------|              |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{
				start: time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			},
			Period[time.Time]{
				start: time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |-------|              |
		------------------------------------------------
		|  R  |     |-------|                          |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{
				start: time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			},
			Period[time.Time]{},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |-------|              |
		| (2) |     |-------|                          |
		------------------------------------------------
		|  R  |                 |-------|              |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{},
			Period[time.Time]{
				start: time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------------------|              |
		| (2) |     |-------------------|              |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{

			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{},
			Period[time.Time]{},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got1, got2 := tc.p1.Punch(tc.p2)

			if !tc.e1.Equal(got1) {
				t.Errorf("both periods should be equal, expected %v, got %v", tc.e1, got1)
			}

			if !tc.e2.Equal(got2) {
				t.Errorf("both periods should be equal, expected %v, got %v", tc.e2, got2)
			}
		})
	}
}

func ExamplePeriod_Punch() {
	a := NewPeriod(
		time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
	)

	b := NewPeriod(
		time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
	)

	l, r := a.Punch(b)

	fmt.Printf("%s - %s\n", l.Min(), l.Max())
	fmt.Printf("%s - %s\n", r.Min(), r.Max())

	// Output:
	// 2023-12-02 00:00:00 +0000 UTC - 2023-12-03 00:00:00 +0000 UTC
	// 2023-12-05 00:00:00 +0000 UTC - 2023-12-07 00:00:00 +0000 UTC
}

func TestPeriod_Encompass(t *testing.T) {
	var table = []struct {
		p1 Period[time.Time]
		p2 Period[time.Time]
		e  Period[time.Time]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) | |-------|                              |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |     |---------------|                  |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |         |-------|                      |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |         |-----------|                  |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                     |-------|          |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |             |---------------|          |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                         |-------|      |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 8, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |                      |
		| (2) |                 |                      |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			NewPeriod(
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
			),
			Period[time.Time]{},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.p1.Encompass(tc.p2)

			if !tc.e.Equal(got) {
				t.Errorf("both periods should be equal, expected %v, got %v", tc.e, got)
			}
		})
	}
}

func ExamplePeriod_Encompass() {
	a := NewPeriod(
		time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 5, 0, 0, 0, 0, time.UTC),
	)

	b := NewPeriod(
		time.Date(2023, time.December, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.December, 6, 0, 0, 0, 0, time.UTC),
	)

	ab := a.Encompass(b)

	fmt.Printf("%s - %s\n", ab.Min(), ab.Max())

	// Output:
	// 2023-12-03 00:00:00 +0000 UTC - 2023-12-06 00:00:00 +0000 UTC
}
