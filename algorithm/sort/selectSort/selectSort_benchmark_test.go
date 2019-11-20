package selectSort

import (
	"math/rand"
	"testing"
)

func BenchmarkSelectSort(b *testing.B) {
	b.ResetTimer()
	array := new(IntSlick)
	for i:=0;i<1000;i++ {
		*array = append(*array, rand.Intn(1000))
	}
	for i:=0;i<b.N;i++ {
		array.SelectSortFunc()
	}
}