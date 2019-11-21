package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := new(IntSlick)
	for i:=0;i<100;i++ {
		*array = append(*array, rand.Intn(1000))
	}
	array.quickSort(0,array.Len()-1)
	fmt.Println("array:", array)
}
