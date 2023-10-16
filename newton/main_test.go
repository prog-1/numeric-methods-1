package main

import (
	"math"
	"testing"
)

func TestTangent(t *testing.T) {
	const eps = 1e-3
	for _, tc := range []struct {
		name string
		x    float64
		f    func(float64) float64
		want float64
	}{
		{"case-1", 1, func(x float64) float64 { return x*x*x - 2*x - 5 }, 2.095},
		{"case-2", 11, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 }, 2.824},
		{"case-3", -1, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 }, -1.592},
		{"case-4", 1.5, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, 3.201},
		{"case-5", -1, func(x float64) float64 { return (1+x)*math.Pow(2.718, 3*x) - 5 }, 0.42},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := SolveNewton(tc.x, 1e-3, tc.f)
			if math.Abs(got-tc.want) > eps {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
