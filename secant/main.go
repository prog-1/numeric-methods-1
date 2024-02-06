package main

import (
	"fmt"
	"math"
)

func SolveSecant(x0, x1, eps float64, f func(float64) float64) float64 {
	var count int
	for ; math.Abs(x1-x0) > eps; count++ {
		fx0, fx1 := f(x0), f(x1)
		x2 := (x0*fx1 - x1*fx0) / (fx1 - fx0)
		x0, x1 = x1, x2
		if math.Abs(f(x2)) < eps {
			break
		}
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*3)
	return x1
}

func main() {
	fmt.Println(SolveSecant(-3, 4, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
	fmt.Println(SolveSecant(1, 5, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
	fmt.Println(SolveSecant(2, 20, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
}
