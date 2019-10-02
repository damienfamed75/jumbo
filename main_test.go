package main

import (
	"testing"
)

func BenchmarkCalpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calpha()
	}
}

func BenchmarkGoAlphabet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alphabet()
	}
}
