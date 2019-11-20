package sort

import (
	"math/rand"
	"testing"
)

func BenchmarkInsertionSort(b *testing.B) {
	b.ResetTimer()
	array := new(IntSlick)
	for i:=0;i<1000;i++ {
		*array = append(*array, rand.Intn(10000))
	}
	for i:=0;i<b.N;i++ {
		array.InsertionSortFunc()
	}
}