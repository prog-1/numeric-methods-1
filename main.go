package main

import (
	"fmt"
	"math"
)

// ic - iteration count | fc - function count

// ### BISECTION METHOD ###
func bisectionMethod(f func(float64, *int) float64, a, b, eps float64) (c float64, ic, fc int) {

	for ; math.Abs(a-b) > eps; ic++ {

		if f(c, &fc) == 0 {
			return c, ic, fc
		}

		c = (a + b) / 2
		if f(c, &fc)*f(a, &fc) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return a, ic, fc
}

// ### NEWTON'S METHOD ###
func newtonMethod(f func(float64, *int) float64, x0, e float64) (float64, int, int) {
	var ic, fc int

	dx := e * 10e-2
	for ; math.Abs(f(x0, &fc)) > e; ic++ {

		y0 := f(x0, &fc)
		x1 := x0 + dx
		y1 := f(x1, &fc)

		dy := y1 - y0
		k := dy / dx

		//next
		x1 = x0 - y0/k
		dx = x0 - x1
		x0 = x1
	}
	return x0, ic, fc
}

// ### SECANT METHOD ###
func secantMethod(f func(float64, *int) float64, x0, x1, e float64) (x2 float64, ic, fc int) {
	root := func(x0, x1 float64) float64 {
		return x0 - f(x0, &fc)*(x1-x0)/(f(x1, &fc)-f(x0, &fc))
	}

	for px0 := x0; math.Abs(f(x2, &fc)) > e; px0, x0, x1, ic = x0, x1, x2, ic+1 {
		x2 = root(x0, x1)

		if x2 < math.Min(x0, x1) || x2 > math.Max(x0, x1) { //if x2 not e [x0,x1]
			x2 = px0 //taking previous x0 as x2
		}
	}

	return x2, ic, fc
}

// ### COMPARRISON ###
func main() {
	type eq struct {
		equation func(float64, *int) float64
		a, b, e  float64
	}
	equations := []eq{
		{equation1, -1, 3, 0.005},
		{equation2, -5, 5, 0.005},
		{equation3, 0, 10, 0.005},
		{equation4, -2, 2, 0.005},
		{equation5, -3, 1, 0.005},
	}

	for i, eq := range equations {
		fmt.Println()
		fmt.Println("Equation", i+1)

		gotbm, icbm, fcbm := bisectionMethod(eq.equation, eq.a, eq.b, eq.e)
		fmt.Printf("Bisection Mentod: Root: %v | Iteration count: %v | Function count: %v \n", gotbm, icbm, fcbm)

		gotnm, icnm, fcnm := newtonMethod(eq.equation, eq.a, eq.e)
		fmt.Printf("Bisection Mentod: Root: %v | Iteration count: %v | Function count: %v \n", gotnm, icnm, fcnm)

		gotsm, icsm, fcsm := secantMethod(eq.equation, eq.a, eq.b, eq.e)
		fmt.Printf("Bisection Mentod: Root: %v | Iteration count: %v | Function count: %v \n", gotsm, icsm, fcsm)

	}
}

// ### EQUATIONS ###

func equation1(x float64, fc *int) float64 {
	*fc++
	return (x * x * x) - 2*x - 5 //root: 2.095
}

func equation2(x float64, fc *int) float64 {
	*fc++
	return 3*(x*x*x*x) - 4*(x*x*x) - 12*(x*x) - 5 //roots: -1.593 || 2.824
}

func equation3(x float64, fc *int) float64 {
	*fc++
	return x - 5*math.Sin(x) - 3.5 //roots: 3.201 || 8.201
}

func equation4(x float64, fc *int) float64 {
	*fc++
	return (1+x)*math.Pow(10, 3*x) - 5 //root: 0.204
}

func equation5(x float64, fc *int) float64 {
	*fc++
	return (1+x)*math.Pow(10, -2*x) - (-8) //root: -1.061
}
