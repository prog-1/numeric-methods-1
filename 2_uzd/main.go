package main

import (
	"math"
)

func SolveNewton(x, eps float64, f func(float64) float64) float64 {
	for math.Abs(f(x)) > eps {
		k := (f(x+eps) - f(x)) / (eps)
		x = x - f(x)/k
	}
	return x
}
