package search

import "fmt"

func BinarySearch(array []int, target int)  {
	left, right := 0, len(array) -1
	index :=0
	for {
		mid := (right-left)/2
		midVal := array[mid]
		if midVal > target{
			right = mid
		} else if midVal < target {
			left = mid
		} else {
			index = mid
			break
		}
	}
	fmt.Println("index value:",index)
}
