package main

import (
	"math/rand"
	"testing"
)

func BenchmarkIntAdd32(b *testing.B) {
	num := rand.Int31n(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num += 1
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkIntMul32(b *testing.B) {
	num := rand.Int31n(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num *= int32(i)
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkIntAdd64(b *testing.B) {
	num := rand.Int63n(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num += 1
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkIntMul64(b *testing.B) {
	num := rand.Int63n(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		num *= int64(i)
	}
	b.StopTimer()
	keep(num)
}
