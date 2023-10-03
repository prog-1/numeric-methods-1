package main

import (
	"math"
)

func SolveBisect(a, b, eps float64, f func(float64) float64) float64 {
	for math.Abs(a-b) >= eps {
		t := (a + b) / 2
		as, ts := f(a), f(t)
		if ts == 0 {
			return t
		}
		if as*ts > 0 {
			a = t
		} else {
			b = t
		}
	}
	return (a + b) / 2
}
