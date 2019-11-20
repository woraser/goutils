package http

import (
	"testing"
)

func Benchmark(b *testing.B) {
	b.ResetTimer()
	size := 20
	ch := make(chan int, size)
	for i := 0; i < b.N; i++ {
		go Fortest(ch)
	}
	for i := 0; i < b.N; i++ {
		<-ch
	}
}
