package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SolveNewton(5, 0.003, func(x float64) float64 { return x*x*x - 2*x - 5 }))
}

func SolveNewton(x, eps float64, f func(float64) float64) float64 {
	dx := eps
	var count int
	for math.Abs(f(x)) > eps {
		count++
		kx := (f(x+dx) - f(x)) / dx
		x1 := x - (f(x) / kx)
		dx = x - x1
		x = x1

	}
	fmt.Printf("Solve newton cycle ran %v times, and f(x) was count %v times\n", count, count*4+1)
	return x
}
