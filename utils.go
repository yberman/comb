package comb

// ToSliceOfSlice takes a channel of slices and makes it a slice of slices
func ToSliceOfSlice(ch chan []int) [][]int {
	r := make([][]int, 0)
	for s := range ch {
		r = append(r, s)
	}
	return r
}

// CopySlice copies and integer slice
func CopySlice(s []int) (sCopy []int) {
	sCopy = make([]int, len(s))
	copy(sCopy, s)
	return
}
