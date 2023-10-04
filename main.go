package main

func SolveBisect(a, b float64, eps float64, f func(float64) float64) float64 {
	if a > b {
		a, b = b, a
	}
	for b-a > 1e-4 {
		if f((a+b)/2)*f(a) > 0 {
			a = (a + b) / 2
		} else {
			b = (a + b) / 2
		}
	}
	return (a + b) / 2
}
