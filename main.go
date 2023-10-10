package main

import (
	"log"
	"math"
)

func SolveBisect(a, b, eps float64, f func(float64) float64) float64 {
	if f(a)*f(b) >= 0 {
		log.Fatal("errrrrror")
	}
	for math.Abs(a-b) > eps {
		if f((a+b)/2) == 0 {
			return (a + b) / 2
		}
		if f((a+b)/2)*f(a) > 0 {
			a = (a + b) / 2
		} else {
			b = (a + b) / 2
		}
	}
	return (a + b) / 2
}
