package sort


func (a IntSlick) quickSort(lo, hi int) {
	if lo >= hi {
		return
	}
	p := a.partition(lo, hi)
	a.quickSort(lo, p-1)
	a.quickSort(p+1, hi)

}

func (a IntSlick) partition(lo, hi int) int {
	pivot := getMiddleValue(a[lo],a[hi/2],a[hi])
	i := lo -1
	for j := lo; j < hi ; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1] ,a[hi] = a[hi], a[i+1]
	return i + 1
}




func getMiddleValue(a, b, c int) int{
	if a < b && b < c {
		return b
	}
	if a < c && c < b {
		return c
	}
	return b
}