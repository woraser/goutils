package sort

//complexity:O(N^2)

func (s *IntSlick) SelectSortFunc() {
	len := s.Len()
	for i:=0; i < len; i++ {
		minIndex := i
		for j:=i+1; j < len; j++ {
			if s.Less(j, minIndex) {
				minIndex = j
			}
		}
		s.Swap(i, minIndex)
	}

}

