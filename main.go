package main

import (
	"math"
)

func SolveBisect(a, b float64, eps float64, f func(float64) float64) float64 {
	var half float64
	for math.Abs(a-b) > eps {
		half = (a + b) / 2
		if f(half) == 0 {
			return half
		}
		if f(half)*f(a) > 0 {
			a = half
		} else {
			b = half
		}
	}
	return half
}
