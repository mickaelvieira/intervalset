package time

import stdtime "time"

func Min(values ...stdtime.Time) stdtime.Time {
	var m stdtime.Time
	for _, v := range values {
		if m.IsZero() || v.Before(m) {
			m = v
		}
	}
	return m
}

func Max(values ...stdtime.Time) stdtime.Time {
	var m stdtime.Time
	for _, v := range values {
		if m.IsZero() || v.After(m) {
			m = v
		}
	}
	return m
}
