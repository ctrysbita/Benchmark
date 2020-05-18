package main

import (
	"math/rand"
	"testing"
)

func BenchmarkFloatAdd32(b *testing.B) {
	num := rand.Float32()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num += 0.1
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkFloatMul32(b *testing.B) {
	num := rand.Float32()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num *= 0.6
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkFloatAdd64(b *testing.B) {
	num := rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num += 0.1
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkFloatMul64(b *testing.B) {
	num := rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num *= 0.6
	}
	b.StopTimer()
	keep(num)
}
