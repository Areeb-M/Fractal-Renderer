package main

import "math/cmplx"

// Defines Fractal Functions
//
// Arguments
// complex128: a point in the complex set
// float: the divergence threshold for the fractal
// int: max iterations to run
//
// Returns:
// int: how many iterations before the threshold was crossed
type FractalFunction func(complex128, float64, int) int

// Mandelbrot Set
// f(z) = z^2 + c
func MandelbrotSet(c complex128, threshold float64, maxIterations int) int {
	z := complex128(0)

	for i := 0; i < maxIterations; i++ {
		z = z*z + c

		if cmplx.Abs(z) > threshold {
			return i
		}
	}

	return maxIterations
}
