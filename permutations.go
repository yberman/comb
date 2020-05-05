package comb

// Factorial returns n!
func Factorial(n int) (r int) {
	r = 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}

// Permutations returns an iter of all permutations
// of the numbers 0 through n-1. It is implemented recursively
func Permutations(n int) (iter chan []int) {
	iter = make(chan []int, channelBufferSize)
	go func() {
		if n == 0 {
			iter <- []int{}
			close(iter)
			return
		}
		for p := range Permutations(n - 1) {
			for i := 0; i <= len(p); i++ {
				// insert n-1 in position i
				pC := append(CopySlice(p), 0)
				copy(pC[i+1:], pC[i:])
				pC[i] = n - 1
				iter <- pC
			}
		}
		close(iter)
	}()
	return
}
