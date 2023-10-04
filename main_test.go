package main

import (
	"math"
	"testing"
)

func TestSolveBisect(t *testing.T) {
	for _, tc := range []struct {
		name    string
		formula func(float64) float64
		a, b, e float64
		want    float64
	}{
		{"1", func(x float64) float64 { return (x * x * x) - 2*x - 5 }, -100, 100, 0.005, 2.093505859375},
		{"2", func(x float64) float64 { return 3*(x*x*x*x) - 4*(x*x*x) - 12*(x*x) - 5 }, -100, 100, 0.005, -1.593017578125},
		{"3", func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, -100, 100, 0.005, 3.2012939453125},
		{"4", func(x float64) float64 { return (1+x)*math.Pow(10, 3*x) - 5 }, -100, 100, 0.005, 0.2044677734375},
		{"5", func(x float64) float64 { return (1+x)*math.Pow(10, -2*x) - (-8) }, -100, 100, 0.005, -1.06201171875},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := SolveBisect(tc.formula, tc.a, tc.b, tc.e)
			if math.Abs(got-tc.want) > tc.e {
				t.Errorf("got = %v, want = %v with e: %v", got, tc.want, tc.e)
			}
		})
	}
}

func main() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test Bisection", F: TestSolveBisect},
		},
		/* benchmarks */ nil /* examples */, nil)
}
