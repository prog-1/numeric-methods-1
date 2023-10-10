package main

import (
	"math"
)

func SolveNewton(x, eps float64, f func(float64) float64) float64 {
	dx := eps
	for math.Abs(f(x)) > eps {
		k := (f(x+dx) - f(x)) / dx
		x1 := x - f(x)/k
		dx = x - x1
		x = x1
	}
	return x
}
