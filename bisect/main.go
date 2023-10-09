package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(SolveBisect(1, 5, 0.001, func(x float64) float64 { return x*x*x - 2*x - 5 }))
	SolveBisect(1, 5, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 })
	SolveBisect(1, 5, 1e-3, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 })
	SolveBisect(-1, -5, 1e-3, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 })
	SolveBisect(1.5, 4.5, 1e-3, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 })
	SolveBisect(-1, 1.5, 1e-3, func(x float64) float64 { return (1+x)*math.Pow(2.718, 3*x) - 5 })
	SolveBisect(0, -3, 1e-3, func(x float64) float64 { return (1+x)*math.Pow(2.718, (-2)*x) - (-8) })
}

func SolveBisect(a, b float64, eps float64, f func(float64) float64) float64 {
	var half float64
	var count int
	for math.Abs(a-b) > eps {
		count++
		half = (a + b) / 2
		if f(half) == 0 {
			return half
		}
		if f(half)*f(a) > 0 {
			a = half
		} else {
			b = half
		}
	}
	fmt.Printf("Solve newton cycle ran %v times, and f(x) was count %v times\n", count, count*3)
	return half
}
