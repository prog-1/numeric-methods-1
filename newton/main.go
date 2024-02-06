package main

import (
	"fmt"
	"math"
)

func SolveNewton(x0, eps float64, f func(float64) float64) float64 {
	var count int
	for dx := eps; math.Abs(f(x0)) > eps; count++ {
		fx0 := f(x0)
		dy := f(x0+dx) - fx0
		k := dy / dx // it's also the derivative of the function
		x1 := x0 - fx0/k
		dx = x0 - x1
		x0 = x1
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*3+1) // because the program calculates f(x0) 1 additional time before exiting the loop
	return x0
}

func main() {
	fmt.Println(SolveNewton(1, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
	fmt.Println(SolveNewton(5, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
	fmt.Println(SolveNewton(-3, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
}
