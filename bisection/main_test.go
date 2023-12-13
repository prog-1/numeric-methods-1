package main

import (
	"math"
	"testing"
)

func TestSolveBisect(t *testing.T) {
	for _, tc := range []struct {
		name string
		a, b float64
		f    func(float64) float64
		want float64
	}{
		{"1", 1, 5, func(x float64) float64 { return x*x*x - 2*x - 5 }, 2.0946},
		{"2", 1, 5, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 }, 2.8239},
		{"3", -5, -1, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 }, -1.592},
		{"4", 1, 5, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, 3.2014},
		{"5", -1, 5, func(x float64) float64 { return (1+x)*math.Pow(2.718, 3*x) - 5 }, 0.4197},
		{"6", -5, 1, func(x float64) float64 { return (1+x)*math.Pow(2.718, (-2)*x) - (-8) }, -1.4448},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := SolveBisect(tc.a, tc.b, 1e-4, tc.f); math.Abs(got-tc.want) > 1e-4 {
				t.Errorf("testcase %v  got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}
