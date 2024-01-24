package main

import (
	"fmt"
)

func SolveNewton(x0, eps float64, f func(float64) float64) float64 {
	var count int
	for ; f(x0) != 0; count++ {
		//	der := (f(x+dx) - f(x)) / dx
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*3)
	return x0
}

func main() {

}
