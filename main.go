package main

import (
	"fmt"
	"math"
)

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

func SolveSecantImproved(a, b float64, eps float64, f func(float64) float64) float64 {
	// The point of improvement is that in case intersection of our secant
	// with OX axis is beyond the range from x0 to x1, then in further calculations
	// we can use the nearest argument out of the initial two.
	fa, fb := f(a), f(b)
	var c, fc float64
	for {
		c = (a*fb - b*fa) / (fb - fa)
		// How to check whether c is within the range from a to b?
		// Option 1:
		if a > b {
			if c > a {
				c = a
			} else if c < b {
				c = b
			}
		} else {
			if c < a {
				c = a
			} else if c > b {
				c = b
			}
		}
		/////////////////////////////
		fc = f(c)
		if math.Abs(fc) <= eps {
			return c
		}
		a, b, fa, fb = b, c, fb, fc
	}
}

func main() {
	// fmt.Println(SolveSecantImproved(-1.570796315078, 1, 1e-3, func(x float64) float64 {
	// 	return 2*math.Sin(x) + 1.5
	// })) // -0.84806
	fmt.Println(SolveSecantImproved(1.089, 5.148, 1e-3, func(x float64) float64 {
		return (x-3)*math.Sin(x) - 0.00501
	})) //3.071
}
