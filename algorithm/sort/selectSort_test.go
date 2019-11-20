package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSelectSort(t *testing.T) {
	array := new(IntSlick)
	for i:=0;i<100;i++ {
		*array = append(*array, rand.Intn(1000))
	}
	array.SelectSortFunc()
	fmt.Println("array:", array)
}