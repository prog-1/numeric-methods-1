package main

import "math"

func SolveSecant(x0, x1, eps float64, f func(float64) float64) float64 {
	x2 := math.MaxFloat64
	for math.Abs(f(x2)) > eps {
		x2 = (x0*f(x1) - x1*f(x0)) / (f(x1) - f(x0))
		x0, x1 = x1, x2
	}
	return x2
}
