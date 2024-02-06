package main

import (
	"math"
	"testing"
)

func TestSolveNewton(t *testing.T) {
	eps := 1e-3
	for _, tc := range []struct {
		name string
		x0   float64
		f    func(float64) float64
		want float64
	}{
		{"1", 0, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"2", 1, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"3", 4, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"4", -3, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"5", 1, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"6", 5, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"7", 2, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"8", 20, func(x float64) float64 { return math.Pow(x, 3) - 2*x - 5 }, 2.095},
		{"9", 2, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, 2.824},
		{"10", 5, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, 2.824},
		{"11", 2, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, 2.824},
		{"12", 7, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, 2.824},
		{"13", 1, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, -1.592},
		{"14", -5, func(x float64) float64 { return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5 }, -1.592},
		{"15", 1.5, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, 3.201},
		{"16", 4.5, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, 3.201},
		{"17", 2, func(x float64) float64 { return (1+x)*math.Pow(math.E, 3*x) - 5 }, 0.42},
		{"18", -1, func(x float64) float64 { return (1+x)*math.Pow(math.E, 3*x) - 5 }, 0.42},
		{"19", -1, func(x float64) float64 { return (1+x)*math.Pow(math.E, (-2)*x) - (-8) }, -1.445},
		{"20", -3, func(x float64) float64 { return (1+x)*math.Pow(math.E, (-2)*x) - (-8) }, -1.445},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := SolveNewton(tc.x0, eps, tc.f)
			if math.Abs(got-tc.want) > eps {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
