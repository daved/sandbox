package main

import "testing"

var t bool

func BenchmarkAtomicSet(b *testing.B) {
	a := atomicBool{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.setValue(true)
	}
}

func BenchmarkAtomicGet(b *testing.B) {
	a := atomicBool{}
	a.setValue(true)

	var v bool

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = a.value()
	}

	t = v
}

func BenchmarkMutex(b *testing.B) {
	s := safeBool{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.setValue(true)
	}
}

func BenchmarkMutexGet(b *testing.B) {
	s := safeBool{}
	s.setValue(true)

	var v bool

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = s.value()
	}

	t = v
}
