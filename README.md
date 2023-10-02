# Bisection Method Algorithm

Follow the below procedure to get the solution for the continuous function:

For any continuous function $f(x)$,

-   Find two points, say $a$ and $b$ such that $a < b$ and $f(a)* f(b)$ < 0
-   Find the midpoint of $a$ and $b$, say $t$
-   $t$ is the root of the given function if $f(t) = 0$; else follow the next step
-   Divide the interval $[a, b]$ – If $f(t)*f(a) <0$, there exist a root between $t$ and $a$  
    – else if $f(t) *f (b) < 0$, there exist a root between $t$ and $b$
-   Repeat above three steps until $|a-b| < \epsilon$.

The bisection method is an approximation method to find the roots of the given equation by repeatedly dividing the interval. This method will divide the interval until the resulting interval is found, which is extremely small.

## Exercises

- $x^3-2x=5$
- $3x^4-4x^3-12x^2=5$
- $x-5sin(x)=3.5$
- $(1+x)e^{ax}=b$ with some $a$ and $b$ e.g. $3; 5$ and $-2; -8$
