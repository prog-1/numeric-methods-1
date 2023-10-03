package main

import (
	"log"
	"math"
)

func SolveBisect(a, b float64, eps float64, f func(float64) float64) float64 {
	if f(a)*f(b) >= 0 {
		log.Fatal("invalid interval")
	}
	var t float64
	for math.Abs(a-b) > eps {
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
	return t
}
