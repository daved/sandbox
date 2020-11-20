package main

import "testing"

var t bool

func BenchmarkAtomicSet(b *testing.B) {
	a := atomicBool{}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			a.setValue(true)
		}
	})
}

func BenchmarkAtomicGet(b *testing.B) {
	a := atomicBool{}
	a.setValue(true)

	var v bool

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v = a.value()
		}
	})

	t = v
}

func BenchmarkMutex(b *testing.B) {
	s := safeBool{}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.setValue(true)
		}
	})
}

func BenchmarkMutexGet(b *testing.B) {
	s := safeBool{}
	s.setValue(true)

	var v bool

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v = s.value()
		}
	})

	t = v
}
