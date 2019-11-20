package sort


func (s *IntSlick) ShellSortFunc()  {
	len := s.Len()
	//
	h := 1

	for {
		if h > (len/3) {
			break
		}
		h = 3 * h + 1
	}
	for {
		if h < 1 {
			break
		}

		for i:=h; i< len; i++ {
			for j:=i;j>=h && s.Less(j,j-h);j-=h {
				s.Swap(j,j-h)
			}
		}
		h = h/3
	}

}