package comb

// Compose {n} allows you to take compositions of n
type Compose int

// Positive are positive compositions of an integer
func (cm Compose) Positive(k int) (iter Iter) {
	iter = newIntChan()
	go func() {
		for c := range Combinations(int(cm-1), k-1) {
			c2 := []int{-1}
			c2 = append(c2, c...)
			c2 = append(c2, int(cm-1))
			diffs := make([]int, len(c2)-1)
			for i := 0; i < len(diffs); i++ {
				diffs[i] = c2[i+1] - c2[i]
			}
			iter <- diffs
		}
		close(iter)
	}()
	return
}

var a = `    for comp in positive(n+k, k):
        yield tuple(x-1 for x in comp)

`

// NonNegative non-negative compositions
func (cm Compose) NonNegative(k int) (iter Iter) {
	iter = newIntChan()
	go func() {
		for pos := range Compose(int(cm) + k).Positive(k) {
			for i := 0; i < len(pos); i++ {
				pos[i]--
			}
		}
		close(iter)
	}()
	return
}
