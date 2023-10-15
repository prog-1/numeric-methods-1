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

// SolveSecant uses the secant approximation method to find the root
// of the given `f` with the given precision `eps`.
func SolveSecant(a, b float64, eps float64, f func(float64) float64) float64 {
	fa, fb := f(a), f(b)
	var c, fc float64
	for {
		c = (a*fb - b*fa) / (fb - fa)
		fc = f(c)
		if math.Abs(fc) <= eps {
			return c
		}
		a, b, fa, fb = b, c, fb, fc
	}
}

func main() {
	fmt.Println(SolveSecant(-1.5, 1, 1e-3, func(x float64) float64 {
		return 2*math.Sin(x) + 1.5
	})) // -0.84806
}
