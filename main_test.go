package main

import (
	"math"
	"testing"
)

func TestSolveBisect(t *testing.T) {
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
		{"x^3-x*2=5", Input{5.0, -5.0, epsilon, func(x float64) float64 {
			return x*x*x - 2*x - 5
		}}, 2.094551},
		{"3x^4-4x^3-12x^2=5", Input{2, -4, epsilon, func(x float64) float64 {
			return 3*math.Pow(x, 4) - 4*math.Pow(x, 3) - 12*x*x - 5
		}}, -1.592088},
		{"x-5sin(x)=3.5", Input{0, 6, epsilon, func(x float64) float64 {
			return x - 5*math.Sin(x) - 3.5
		}}, 3.2013568},
		{"(1+x)e^{3*x}=5", Input{-1, 1, epsilon, func(x float64) float64 {
			return (1+x)*math.Pow(10, 3*x) - 5
		}}, 0.2058877},
		{"(1+x)e^{-2*x}=-8", Input{-1.1, 1, epsilon, func(x float64) float64 {
			return (1+x)*math.Pow(10, -2*x) + 8
		}}, -1.060536},
	} {
		if got := SolveBisect(tc.input.a, tc.input.b, tc.input.eps, tc.input.f); math.Abs(got-tc.want) > tc.input.eps {
			t.Errorf("%s failed: got = %f, want = %f", tc.name, got, tc.want)
		}
	}
}

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
