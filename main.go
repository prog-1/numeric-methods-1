package main

import (
	"math"
)

func SolveBisect(f func(x float64) float64, a, b, eps float64) float64 {
	var c float64

	for math.Abs(a-b) > eps {

		if f(c) == 0 {
			return c
		}

		c = (a + b) / 2
		if f(c)*f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return a
}
