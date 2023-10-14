package main

import (
	"math"
	"testing"
)

func TestSolveNewton(t *testing.T) {
	for _, tc := range []struct {
		x, eps float64
		f      func(float64) float64
		want   float64
	}{
		{x: 1, eps: 1e-3, f: func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, want: 2.095},
		{x: 3, eps: 1e-3, f: func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, want: 2.824},
		{x: 2, eps: 1e-5, f: func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, want: 3.20136},
		{x: -1, eps: 1e-2, f: func(x float64) float64 { return (1+x)*math.Pow(10, 3*x) - 5 }, want: 0.206},
		{x: -2, eps: 1e-2, f: func(x float64) float64 { return (1+x)*math.Pow(10, (-2)*x) + 8 }, want: -1.06},
	} {
		got := SolveNewton(tc.x, tc.eps, tc.f)
		if math.Abs(got-tc.want) > tc.eps {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}

	}
}
