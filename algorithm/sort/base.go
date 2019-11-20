package sort

type IntSlick []int

func (i IntSlick) Len() int {
	return len(i)
}

func (s IntSlick) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}

func (s IntSlick) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
