package main

import (
	"math"
	"testing"
)

func TestMethods(t *testing.T) {
	for _, tc := range []struct {
		name     string
		equation func(float64, *int) float64
		a, b, e  float64
		want     float64
		want2    float64
	}{
		{"1", equation1, -1, 3, 0.005, 2.095, 0},
		{"2", equation2, -5, 5, 0.005, -1.593, 2.824},
		{"3", equation3, 0, 10, 0.005, 3.201, 8.201},
		{"4", equation4, -2, 2, 0.005, 0.204, 0},
		{"5", equation5, -3, 1, 0.005, -1.061, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Equation", tc.name)
			gotbm, _, _ := bisectionMethod(tc.equation, tc.a, tc.b, tc.e)
			if math.Abs(gotbm-tc.want) > tc.e && math.Abs(gotbm-tc.want2) > tc.e {
				t.Errorf("Bisection method: got = %v, want = %v or %v with e: %v", gotbm, tc.want, tc.want2, tc.e)
			}

			gotnm, _, _ := newtonMethod(tc.equation, tc.a, tc.e)
			if math.Abs(gotnm-tc.want) > tc.e && math.Abs(gotnm-tc.want2) > tc.e {
				t.Errorf("Bisection method: got = %v, want = %v or %v with e: %v", gotnm, tc.want, tc.want2, tc.e)
			}

			gotsm, _, _ := secantMethod(tc.equation, tc.a, tc.b, tc.e)
			if math.Abs(gotsm-tc.want) > tc.e && math.Abs(gotsm-tc.want2) > tc.e {
				t.Errorf("Secant method: got = %v, want = %v or %v with e: %v", gotsm, tc.want, tc.want2, tc.e)
			}
		})
	}
}

// ⬇️ Uncomment for tests ⬇️

func main2() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test Methods", F: TestMethods},
		},
		/* benchmarks */ nil /* examples */, nil)
}
