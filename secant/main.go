package main

import (
	"fmt"
	"math"
)

func SolveSecant(a, b, eps float64, f func(float64) float64) float64 {
	var count int
	var x0, x1 float64
	if a < b {
		x0, x1 = a, b
	} else {
		x0, x1 = b, a
	}
	for math.Abs(b-a) > eps {
		count++
		fA, fB := f(a), f(b)
		x := (a*fB - b*fA) / (fB - fA)
		if x < x0 {
			x = x0
		} else if x > x1 {
			x = x1
		}
		a = b
		b = x
		if math.Abs(f(x)) < eps {
			break
		}
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*3)
	return b
}

func main() {
	SolveSecant(1, 5, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 })
	SolveSecant(-1, 5, 1e-3, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 })
	SolveSecant(-1, -5, 1e-3, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 })
	SolveSecant(1.5, 4.5, 1e-3, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 })
	SolveSecant(-1, 1.5, 1e-3, func(x float64) float64 { return (1+x)*math.Pow(2.718, 3*x) - 5 })
	SolveSecant(-1, -3, 1e-3, func(x float64) float64 { return (1+x)*math.Pow(2.718, (-2)*x) - (-8) })
}
