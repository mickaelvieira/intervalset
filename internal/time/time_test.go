package time

import (
	"testing"
	"time"
)

func TestMinEmpty(t *testing.T) {
	got := Min()
	expected := time.Time{}

	if !got.Equal(expected) {
		t.Errorf("time should be equal, expected %+v got %+v", got, expected)
	}
}

func TestMin(t *testing.T) {
	t1 := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.January, 1, 14, 0, 0, 0, time.UTC)

	got := Min(t1, t2, t3)

	if !got.Equal(t1) {
		t.Errorf("time should be equal, expected %+v got %+v", got, t1)
	}
}

func TestMaxEmpty(t *testing.T) {
	got := Max()
	expected := time.Time{}
	if !got.Equal(expected) {
		t.Errorf("time should be equal, expected %+v got %+v", got, expected)
	}
}

func TestMax(t *testing.T) {
	t1 := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.January, 1, 14, 0, 0, 0, time.UTC)

	got := Max(t1, t2, t3)

	if !got.Equal(t2) {
		t.Errorf("time should be equal, expected %+v got %+v", got, t2)
	}
}

// The zero time is a valid value, not a marker for "no value seen yet".
func TestMinWithZeroTime(t *testing.T) {
	zero := time.Time{}
	t1 := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)

	if got := Min(zero, t1); !got.Equal(zero) {
		t.Errorf("time should be equal, expected %+v got %+v", zero, got)
	}

	if got := Min(t1, zero); !got.Equal(zero) {
		t.Errorf("time should be equal, expected %+v got %+v", zero, got)
	}
}

func TestMaxWithZeroTime(t *testing.T) {
	zero := time.Time{}
	t1 := time.Date(-100, time.January, 1, 0, 0, 0, 0, time.UTC)

	if got := Max(zero, t1); !got.Equal(zero) {
		t.Errorf("time should be equal, expected %+v got %+v", zero, got)
	}

	if got := Max(t1, zero); !got.Equal(zero) {
		t.Errorf("time should be equal, expected %+v got %+v", zero, got)
	}
}
