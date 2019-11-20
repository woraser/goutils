package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestShellSortFunc(t *testing.T)  {
	array := new(IntSlick)
	for i:=0;i<1000;i++ {
		*array = append(*array, rand.Intn(10000))
	}

	array.ShellSortFunc()
	fmt.Println("array:",array)
}
