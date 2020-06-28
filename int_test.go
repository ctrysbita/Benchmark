package main

import (
	"math/rand"
	"testing"
)

func BenchmarkIntAdd32(b *testing.B) {
	num := rand.Int31n(10)
	b.ResetTimer()
	for i := int32(0); i < int32(b.N); i++ {
		num += i
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkIntMul32(b *testing.B) {
	num := rand.Int31n(10)
	b.ResetTimer()
	for i := int32(0); i < int32(b.N); i++ {
		num *= i
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkIntAdd64(b *testing.B) {
	num := rand.Int63n(10)
	b.ResetTimer()
	for i := int64(0); i < int64(b.N); i++ {
		num += i
	}
	b.StopTimer()
	keep(num)
}

func BenchmarkIntMul64(b *testing.B) {
	num := rand.Int63n(10)
	b.ResetTimer()
	for i := int64(0); i < int64(b.N); i++ {
		num *= i
	}
	b.StopTimer()
	keep(num)
}
