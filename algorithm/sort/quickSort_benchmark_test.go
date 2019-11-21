package sort

import (
	"math/rand"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	b.ResetTimer()
	array := new(IntSlick)
	for i:=0;i<10000;i++ {
		*array = append(*array, rand.Intn(10000))
	}
	for i:=0;i<b.N;i++ {
		array.quickSort(0, array.Len()-1)
	}
}