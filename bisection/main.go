package main

import (
	"fmt"
	"log"
	"math"
)

// SolveBisect uses the bisection approximation method to find the root
// of the given `f` with the given precision `eps`.
func SolveBisect(a, b float64, eps float64, f func(float64) float64) float64 {
	if f(a)*f(b) >= 0 {
		log.Fatal("invalid interval")
	}
	t := (a + b) / 2 // the midpoint of a and b
	var count int
	for ; f(t) != 0 && math.Abs(a-b) > eps; t, count = (a+b)/2, count+1 {
		if f(t)*f(a) < 0 {
			b = t
		} else {
			a = t
		}
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*3)
	return t
}

func main() {
	fmt.Println(SolveBisect(-3, 4, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
	fmt.Println(SolveBisect(1, 5, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
	fmt.Println(SolveBisect(2, 20, 1e-3, func(x float64) float64 { return x*x*x - 2*x - 5 }))
}
