package main

import (
	"testing"

	"github.com/mjibson/go-dsp/fft"
)

// Benckmark for Fast Fourier Transform with 2^10 1D input.
func BenchmarkFFT(b *testing.B) {

	N := 1 << 10
	a := make([]complex128, N)
	for i := 0; i < N; i++ {
		a[i] = complex(float64(i)/float64(N), 0)
	}

	fft.EnsureRadix2Factors(N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fft.FFT(a)
	}
}
