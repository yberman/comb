package comb

func Binomial(n, k int) int {
	if k == 0 {
		return 1
	}
	return n * Binomial(n-1, k-1) / k
}

func Combinations(n, k int) chan Set {
	c := make(chan Set)
	go func() {
		if k == 0 {
			c <- Set{}
			close(c)
			return
		}
		if k == n {
			s := Set{}
			for i := 0; i < n; i++ {
				s.Add(i)
			}
			c <- s
			close(c)
			return
		}
		for s := range Combinations(n-1, k) {
			c <- s
		}
		for s := range Combinations(n-1, k-1) {
			s.Add(n)
			c <- s
		}
		close(c)
	}()
	return c
}
