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
	if got != expected {
		t.Errorf("time should be equal, expected %+v got %+v", got, expected)
	}
}

func TestMax(t *testing.T) {
	t1 := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.January, 1, 14, 0, 0, 0, time.UTC)

	got := Max(t1, t2, t3)

	if got != t2 {
		t.Errorf("time should be equal, expected %+v got %+v", got, t2)
	}
}
