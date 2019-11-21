package sort

func (s *IntSlick) InPlaceMergeSort(l, r int)  {
	if l >= r {
		return
	}

	mid := (r + l) / 2
	s.InPlaceMergeSort(l, mid)
	s.InPlaceMergeSort(mid+1, r)

	s.Merge(l, mid, r)
}

func (arr *IntSlick) Merge(l int, mid int, r int) {
	// create a new array
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = (*arr)[i]
	}

	// offset  of two parts
	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		// left is use up. use right
		if left > mid {
			(*arr)[i] = temp[right-l]
			right++
			// right is use up. use left
		} else if right > r {
			(*arr)[i] = temp[left-l]
			left++
			// compare left and right, choose the smaller
		} else if temp[left - l] > temp[right - l] {
			(*arr)[i] = temp[right - l]
			right++
		} else {
			(*arr)[i] = temp[left - l]
			left++
		}
	}
}