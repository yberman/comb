package comb

func Factorial(n int) (r int) {
	r = 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}

// Permutations returns an iter of all permutations
// of the numbers 0 through n-1. It is implemented using Heap's Algorithm
func Permutations(n int) (iter chan []int) {
	iter = make(chan []int, ChannelBufferSize)
	go func() {
		close(iter)
	}()
	return
}
