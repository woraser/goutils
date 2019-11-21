package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInPlaceMergeSort(t *testing.T) {
	array := new(IntSlick)
	for i:=0;i<5;i++ {
		*array = append(*array, rand.Intn(1000))
	}
	array.InPlaceMergeSort(0,  array.Len() - 1)
	fmt.Println("array:", array)
}
