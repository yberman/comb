package comb

// ToSliceOfSlice takes a channel of slices and makes it a slice of slices
func ToSliceOfSlice(ch chan []int) [][]int {
	r := make([][]int, 0)
	for s := range ch {
		r = append(r, s)
	}
	return r
}
