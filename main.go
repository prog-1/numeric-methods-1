package main

import (
	"fmt"
	"math"
)

// SolveBisect uses the bisection approximation method to find the root
// of the given `f` with the given precision `eps`.
func SolveBisect(a, b float64, eps float64, f func(float64) float64) float64 {
	fa, fb := f(a), f(b)
	if fa*fb >= 0 {
		panic("Invalid range")
	}
	t := (a + b) / 2
	ft := f(t)
	if ft == 0 || math.Abs(a-b) <= eps {
		return t
	}
	// Shortening the one with the same sign
	if fa*ft > 0 { // same sign
		return SolveBisect(t, b, eps, f)
	} else {
		return SolveBisect(a, t, eps, f)
	}
}

func main() {
	fmt.Println(SolveBisect(0, 6, 1e-3, func(x float64) float64 {
		return x - 5*math.Sin(x) - 3.5
	}))
}
