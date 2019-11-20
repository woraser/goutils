package sort

func (s *IntSlick) InsertionSortFunc() {
	len := s.Len()
	for i:=0;i<len; i++ {
		for j :=i; j>0 && s.Less(j, j-1); j-- {
			s.Swap(j, j-1)
		}
	}
}
