package main

import (
	"math"
	"testing"
)

func TestSolveBisect(t *testing.T) {
	for _, tc := range []struct {
		name      string
		a, b, eps float64
		f         func(float64) float64
		want      float64
	}{
		{"x^3-2x=5 [1.75, 2.5] 1e-10", 1.75, 2.5, 1e-10, func(x float64) float64 { return x*x*x - 2*x - 5 }, 2.094551481542326591482386540579302963857306105628239180},
		{"3x^4-4x^3-12x^2=5 [2, 3] 1e-10", 2, 3, 1e-10, func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 }, 2.82385308322479064470680340671677772776501477},
		{"x-5sin(x)=3.5 [3, 4] 1e-10", 3, 4, 1e-10, func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, 3.201356853718969429519726837549393477051},
	} {

		if got := SolveBisect(tc.a, tc.b, tc.eps, tc.f); math.Abs(got-tc.want) > tc.eps {
			t.Errorf("%v failed: got %v, want %v", tc.name, got, tc.want)
		}

	}
}
