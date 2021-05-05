package main

import "testing"

func BenchmarkIsShutdown(b *testing.B) {
	b.Run("chan shutdown", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isShutdown()
		}
	})

	b.Run("bool shutdown", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			boolShutdown()
		}
	})
}
