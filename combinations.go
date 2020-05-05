package comb

const ChannelBufferSize = 8

// Binomial is the number of k combinations over n elements.
func Binomial(n, k int) int {
	if k == 0 {
		return 1
	}
	return n * Binomial(n-1, k-1) / k
}

func duplicateSlice(s []int) (sCopy []int) {
	sCopy = make([]int, len(s))
	copy(sCopy, s)
	return
}

// Combinations produces iter of all the k combinations over the integers
// 0 up to n. Based on Python's combinations in
// https://docs.python.org/3/library/itertools.html
func Combinations(n, k int) (iter chan []int) {
	iter = make(chan []int, ChannelBufferSize)
	go func() {
		result := make([]int, k)
		for i := 0; i < k; i++ {
			result[i] = i
		}
		iter <- duplicateSlice(result)
		for {
			var i int
			var f bool
			for i = k - 1; i >= 0; i-- {
				if result[i] != i+n-k {
					f = true
					break
				}
			}
			if !f {
				close(iter)
				return
			}
			result[i] += 1
			for j := i + 1; j < k; j++ {
				result[j] = result[j-1] + 1
			}
			iter <- duplicateSlice(result)
		}
	}()
	return
}