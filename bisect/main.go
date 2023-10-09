package main

import (
	"fmt"
	"log"
	"math"
)

func SolveBisect(a, b float64, eps float64, f func(float64) float64) float64 {
	if f(a)*f(b) >= 0 {
		log.Fatal("invalid interval")
	}
	var t float64
	var count int
	for math.Abs(a-b) > eps {
		count++
		t = (a + b) / 2
		if f(t) == 0 {
			return t
		}
		if f(t)*f(a) > 0 {
			a = t
		} else {
			b = t
		}
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*3)
	return t
}

func main() {
	SolveBisect(1, 5, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 })
	SolveBisect(1, 5, 1e-3, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 })
	SolveBisect(-1, -5, 1e-3, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 })
	SolveBisect(1.5, 4.5, 1e-3, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 })
	SolveBisect(-1, 1.5, 1e-3, func(x float64) float64 { return (1+x)*math.Pow(2.718, 3*x) - 5 })
	SolveBisect(0, -3, 1e-3, func(x float64) float64 { return (1+x)*math.Pow(2.718, (-2)*x) - (-8) })
}
