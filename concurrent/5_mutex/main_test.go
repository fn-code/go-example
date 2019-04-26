package main

import (
	"testing"
)

func BenchmarkMutexParalell10(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			runningMutex()
		}
	})
}

func BenchmarkMutex10(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		runningMutex()
	}
}

func BenchmarkNotMutex10(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		runningNotMutex()
	}
}
