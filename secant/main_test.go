package main

import (
	"math"
	"testing"
)

func TestSolveSecant(t *testing.T) {
	eps := 1e-3
	for _, tc := range []struct {
		name string
		a, b float64
		f    func(float64) float64
		want float64
	}{
		{"1", -3, 4, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"2", 1, 5, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"3", 2, 20, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"4", 2, 4, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, 2.824},
		{"5", 3, 7, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, 2.824},
		{"6", -2, -1, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, -1.592},
		{"7", 1.5, 4.5, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, 3.201},
		{"8", -2, 1, func(x float64) float64 { return (1+x)*math.Pow(math.E, 3*x) - 5 }, 0.42},
		{"9", -3, 0, func(x float64) float64 { return (1+x)*math.Pow(math.E, (-2)*x) - (-8) }, -1.445},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := SolveSecant(tc.a, tc.b, eps, tc.f)
			if math.Abs(got-tc.want) > eps {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
