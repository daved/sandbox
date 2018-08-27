package main

import "testing"

func BenchmarkMulticomp(b *testing.B) {
	b.Run("Hit", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			multicomp("M")
		}
	})

	b.Run("Miss", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			multicomp("X")
		}
	})
}

func BenchmarkSinglecomp(b *testing.B) {
	b.Run("Hit", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			singlecomp("M")
		}
	})

	b.Run("Miss", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			singlecomp("X")
		}
	})
}
