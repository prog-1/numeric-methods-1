package main

import (
	"fmt"
	"math"
)

func SolveSecant(x0, x1, eps float64, f func(float64) float64) float64 {
	var x2, fx0, fx1 float64
	var count int
	for math.Abs(x1-x0) >= eps {
		fx0, fx1 = f(x0), f(x1)
		x2 = x0 + ((fx0 * (x1 - x0)) / (fx0 - fx1))
		x0, x1 = x1, x2
		count++
	}
	fmt.Println(count)
	return x1
}
