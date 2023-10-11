package main

import (
	"fmt"
	"math"
)

func SolveNewton(x, eps float64, f func(float64) float64) float64 {
	dx := eps
	var count int
	for math.Abs(f(x)) > eps {
		count++
		k := (f(x+dx) - f(x)) / dx
		x1 := x - f(x)/k
		dx = x - x1
		x = x1
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*4+1)
	return x
}

func main() {
	SolveNewton(1, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 })
	SolveNewton(2, 1e-3, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 })
	SolveNewton(-1, 1e-3, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 })
	SolveNewton(-2, 1e-3, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 })
	SolveNewton(-1, 1e-3, func(x float64) float64 { return (1+x)*math.Pow(2.718, 3*x) - 5 })
}
