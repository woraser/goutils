package selectSort


type IntSlick []int

// implement interface for sort
func (i IntSlick) len() int {
	return len(i)
}

func (s IntSlick) less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}

func (s IntSlick) swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *IntSlick) SelectSortFunc() {
	len := s.len()
	for i:=0; i < len; i++ {
		minIndex := i
		for j:=i+1; j < len; j++ {
			if s.less(j, minIndex) {
				minIndex = j
			}
		}
		s.swap(i, minIndex)
	}

}

