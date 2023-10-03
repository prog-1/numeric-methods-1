package main

import (
	"fmt"
	"math"
)

func SolveBisect(a, b float64, e float64, f func(x float64) float64) float64 {
	var t float64
	for math.Abs(a-b) > e {
		t = (b + a) / 2
		if f(t)*f(a) > 0 {
			a = t
		} else {
			b = t
		}
	}
	return t
}

func main() {
	fmt.Println(SolveBisect(1.75, 2.5, 1e-10,
		func(x float64) float64 {
			return x*x*x - 2*x - 5
		},
	))
}
