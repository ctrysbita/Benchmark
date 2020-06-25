package main

import (
	"testing"

	"github.com/mjibson/go-dsp/fft"
)

// run with: go test -test.bench="."
func BenchmarkFFT(b *testing.B) {
	b.StopTimer()

	N := 1 << 20
	a := make([]complex128, N)
	for i := 0; i < N; i++ {
		a[i] = complex(float64(i)/float64(N), 0)
	}

	fft.EnsureRadix2Factors(N)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		fft.FFT(a)
	}
}
