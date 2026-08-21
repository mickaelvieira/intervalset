// Package time compares the time values the periods are made of.
package time

import stdtime "time"

// Min returns the earliest of the given times,
// or the zero time when called without any value.
func Min(values ...stdtime.Time) stdtime.Time {
	var m stdtime.Time
	for i, v := range values {
		if i == 0 || v.Before(m) {
			m = v
		}
	}
	return m
}

// Max returns the latest of the given times,
// or the zero time when called without any value.
func Max(values ...stdtime.Time) stdtime.Time {
	var m stdtime.Time
	for i, v := range values {
		if i == 0 || v.After(m) {
			m = v
		}
	}
	return m
}
