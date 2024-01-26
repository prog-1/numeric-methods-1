package main

import (
	"math"
	"testing"
)

func TestSolveSecant(t *testing.T) {
	const epsilon float64 = 1e-3
	type Input struct {
		a, b, eps float64
		f         func(float64) float64
	}
	for _, tc := range []struct {
		name  string
		input Input
		want  float64
	}{
		// Test case conditions:
		// 1. One root within range
		// 2. No extremums within range
		{"x-5sin(x)-3.5=0", Input{2, 4, epsilon, func(x float64) float64 {
			return x - 5*math.Sin(x) - 3.5
		}}, 3.2013568},
		{"2sin(x)+1.5", Input{-1.5, 1, epsilon, func(x float64) float64 {
			return 2*math.Sin(x) + 1.5
		}}, -0.84806},
	} {
		if got := SolveSecant(tc.input.a, tc.input.b, tc.input.eps, tc.input.f); math.Abs(got-tc.want) > tc.input.eps {
			t.Errorf("SolveSecant failed %s: got = %f, want = %f", tc.name, got, tc.want)
		}
	}
}
