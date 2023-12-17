package intervalset

import (
	"fmt"
	"testing"
)

func TestRange_IsEmpty(t *testing.T) {
	t1 := NewRange[float32](10, 10)

	if !t1.IsEmpty() {
		t.Errorf("interval should be empty %+v", t1)
	}

	t2 := NewRange[float32](10, 10.1)
	if t2.IsEmpty() {
		t.Errorf("interval should not be empty %+v", t2)
	}
}

func TestRange_IsValid(t *testing.T) {
	t1 := NewRange[float32](10, 10)
	if !t1.IsValid() {
		t.Errorf("interval should be valid %+v", t1)
	}

	t2 := NewRange[float32](10, 10.1)
	if !t2.IsValid() {
		t.Errorf("interval should be valid %+v", t2)
	}

	t3 := NewRange[float32](10, 9)
	if t3.IsValid() {
		t.Errorf("interval should not be valid %+v", t3)
	}
}

func TestRange_Equal(t *testing.T) {
	var table = []struct {
		e  bool
		i1 Range[int32]
		i2 Range[int32]
	}{
		{
			true,
			NewRange[int32](1, 2),
			NewRange[int32](1, 2),
		},
		{
			false,
			NewRange[int32](1, 2),
			NewRange[int32](1, 3),
		},
		{
			false,
			NewRange[int32](1, 2),
			NewRange[int32](2, 2),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			if tc.i1.Equal(tc.i2) != tc.e {
				if tc.e {
					t.Errorf("expected i1 to be equal to i2: i1 %+v, i2 %+v", tc.i1, tc.i2)
				} else {
					t.Errorf("expected i1 to not be equal to i2: i1 %+v, i2 %+v", tc.i1, tc.i2)
				}
			}
		})
	}
}

func TestRange_Before(t *testing.T) {
	var table = []struct {
		e  bool
		i1 Range[int32]
		i2 Range[int32]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |-------|              |
		| (2) |         |---|                          |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](5, 7),
			NewRange[int32](3, 4),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |         |-------|                      |
		| (2) |             |-------|                  |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](3, 5),
			NewRange[int32](4, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](2, 4),
			NewRange[int32](5, 6),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			if tc.i1.Before(tc.i2) != tc.e {
				if tc.e {
					t.Errorf("expected i1 to end before the beginning of i2: i1 %+v, i2 %+v", tc.i1, tc.i2)
				} else {
					t.Errorf("expected i1 to not end before the beginning of i2: i1 %+v, i2 %+v", tc.i1, tc.i2)
				}
			}
		})
	}
}

func TestRange_After(t *testing.T) {
	var table = []struct {
		e  bool
		i1 Range[int32]
		i2 Range[int32]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](2, 4),
			NewRange[int32](5, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |-------|                  |
		| (2) |         |-------|                      |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](4, 6),
			NewRange[int32](3, 5),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |-------|              |
		| (2) |     |-------|                          |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](5, 7),
			NewRange[int32](2, 4),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			if tc.i1.After(tc.i2) != tc.e {
				if tc.e {
					t.Errorf("expected i1 to start after the end of i2: i1 %+v, i2 %+v", tc.i1, tc.i2)
				} else {
					t.Errorf("expected i1 to not start after the end of i2: i1 %+v, i2 %+v", tc.i1, tc.i2)
				}
			}
		})
	}
}

func TestRange_Overlap(t *testing.T) {
	var table = []struct {
		e  bool
		i1 Range[int32]
		i2 Range[int32]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](2, 4),
			NewRange[int32](5, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |         |-------|                      |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](3, 5),
			NewRange[int32](5, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |-------|                  |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](4, 6),
			NewRange[int32](5, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |-------|              |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](5, 7),
			NewRange[int32](5, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                     |-------|          |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](6, 8),
			NewRange[int32](5, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                         |-------|      |
		| (2) |                 |---|                  |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](7, 9),
			NewRange[int32](5, 6),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.i1.Overlaps(tc.i2)
			if got != tc.e {
				if tc.e {
					t.Errorf("expected i1 to overlap i2: i1 %+v, i2 %+v, %t", tc.i1, tc.i2, got)
				} else {
					t.Errorf("expected i1 to not overlap i2: i1 %+v, i2 %+v, %t", tc.i1, tc.i2, got)
				}
			}
		})
	}
}

func TestRange_Contains(t *testing.T) {
	var table = []struct {
		e  bool
		i1 Range[int32]
		i2 Range[int32]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (2) | |-------|                              |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](4, 8),
			NewRange[int32](1, 3),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |     |-------|                          |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](4, 8),
			NewRange[int32](2, 4),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |         |-------|                      |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](4, 8),
			NewRange[int32](3, 5),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |             |-------|                  |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](4, 8),
			NewRange[int32](4, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |                 |-------|              |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](4, 8),
			NewRange[int32](5, 7),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |                     |-------|          |
		----------------------------------------------*/
		{
			true,
			NewRange[int32](4, 8),
			NewRange[int32](6, 8),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |                         |-------|      |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](4, 8),
			NewRange[int32](7, 9),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |             |---------------|          |
		| (1) |                             |-------|  |
		----------------------------------------------*/
		{
			false,
			NewRange[int32](4, 8),
			NewRange[int32](8, 10),
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.i1.Contains(tc.i2)
			if got != tc.e {
				if tc.e {
					t.Errorf("expected i2 to be a subset of i1: i1 %+v, i2 %+v", tc.i1, tc.i2)
				} else {
					t.Errorf("expected i2 to not be a subset of i1: i1 %+v, i2 %+v", tc.i1, tc.i2)
				}
			}
		})
	}
}

func TestRange_Intersect(t *testing.T) {
	var table = []struct {
		i1 Range[int32]
		i2 Range[int32]
		e  Range[int32]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |-------|              |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewRange[int32](2, 4),
			NewRange[int32](5, 7),
			Range[int32]{},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |---------------|                  |
		| (2) |             |---|                      |
		------------------------------------------------
		|  R  |             |---|                      |
		----------------------------------------------*/
		{
			NewRange[int32](2, 6),
			NewRange[int32](4, 5),
			NewRange[int32](4, 5),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |         |-------|                      |
		| (2) |     |-------|                          |
		------------------------------------------------
		|  R  |         |---|                          |
		----------------------------------------------*/
		{
			NewRange[int32](3, 5),
			NewRange[int32](2, 4),
			NewRange[int32](3, 4),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |         |-------|                      |
		------------------------------------------------
		|  R  |         |---|                          |
		----------------------------------------------*/
		{
			NewRange[int32](2, 4),
			NewRange[int32](3, 5),
			NewRange[int32](3, 4),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |             |---|                      |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewRange[int32](2, 4),
			NewRange[int32](4, 5),
			Range[int32]{},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.i1.Intersect(tc.i2)
			if !got.Equal(tc.e) {
				t.Errorf("both intervals should be equal, expected %v, got %v", tc.e, got)
			}
		})
	}
}

func ExampleRange_Intersect() {
	a := NewRange[int](2, 7)
	b := NewRange[int](3, 5)

	r := a.Intersect(b)

	fmt.Printf("%d - %d\n", r.Min(), r.Max())

	// Output:
	// 3 - 5
}

func TestRange_Punch(t *testing.T) {
	var table = []struct {
		i1 Range[int32]
		i2 Range[int32]
		e1 Range[int32]
		e2 Range[int32]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------------------|              |
		| (2) |         |-------|                      |
		------------------------------------------------
		|  R  |     |---|       |-------|              |
		----------------------------------------------*/
		{
			NewRange[int32](2, 7),
			NewRange[int32](3, 5),
			NewRange[int32](2, 3),
			NewRange[int32](5, 7),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |                 |-------|              |
		------------------------------------------------
		|  R  |     |-------|                          |
		----------------------------------------------*/
		{
			NewRange[int32](2, 4),
			NewRange[int32](5, 7),
			NewRange[int32](2, 4),
			Range[int32]{},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |-------|              |
		| (2) |     |-------|                          |
		------------------------------------------------
		|  R  |                 |-------|              |
		----------------------------------------------*/
		{
			NewRange[int32](5, 7),
			NewRange[int32](2, 4),
			Range[int32]{},
			NewRange[int32](5, 7),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------------------|              |
		| (2) |     |-------------------|              |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewRange[int32](2, 7),
			NewRange[int32](2, 7),
			Range[int32]{},
			Range[int32]{},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got1, got2 := tc.i1.Punch(tc.i2)

			if !tc.e1.Equal(got1) {
				t.Errorf("both intervals should be equal, expected %v, got %v", tc.e1, got1)
			}

			if !tc.e2.Equal(got2) {
				t.Errorf("both intervals should be equal, expected %v, got %v", tc.e2, got2)
			}
		})
	}
}

func ExampleRange_Punch() {
	a := NewRange(2, 7)
	b := NewRange(3, 5)

	l, r := a.Punch(b)

	fmt.Printf("%d - %d\n", l.Min(), l.Max())
	fmt.Printf("%d - %d\n", r.Min(), r.Max())

	// Output:
	// 2 - 3
	// 5 - 7
}

func TestRange_Encompass(t *testing.T) {
	var table = []struct {
		i1 Range[int32]
		i2 Range[int32]
		e  Range[int32]
	}{
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) | |-------|                              |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewRange[int32](1, 3),
			NewRange[int32](4, 6),
			Range[int32]{},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |     |-------|                          |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |     |---------------|                  |
		----------------------------------------------*/
		{
			NewRange[int32](2, 4),
			NewRange[int32](4, 6),
			NewRange[int32](2, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |         |-------|                      |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |         |-----------|                  |
		----------------------------------------------*/
		{
			NewRange[int32](3, 5),
			NewRange[int32](4, 6),
			NewRange[int32](3, 6),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                     |-------|          |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |             |---------------|          |
		----------------------------------------------*/
		{
			NewRange[int32](6, 8),
			NewRange[int32](4, 6),
			NewRange[int32](4, 8),
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                         |-------|      |
		| (2) |             |-------|                  |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewRange[int32](7, 9),
			NewRange[int32](4, 6),
			Range[int32]{},
		},
		/*----------------------------------------------
		|  T  | 1   2   3   4   5   6   7   8   9   10 |
		| (1) |                 |                      |
		| (2) |                 |                      |
		------------------------------------------------
		|  R  |                 ∅                      |
		----------------------------------------------*/
		{
			NewRange[int32](5, 5),
			NewRange[int32](5, 5),
			Range[int32]{},
		},
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			got := tc.i1.Encompass(tc.i2)

			if !tc.e.Equal(got) {
				t.Errorf("both intervals should be equal, expected %v, got %v", tc.e, got)
			}
		})
	}
}

func ExampleRange_Encompass() {
	a := NewRange[int](3, 5)
	b := NewRange[int](4, 6)

	ab := a.Encompass(b)

	fmt.Printf("%d - %d\n", ab.Min(), ab.Max())

	// Output:
	// 3 - 6
}
