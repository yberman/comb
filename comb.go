package comb

func Binomial(n, k int) int {
    return 0
}

func Combinations(n, k int) chan []int {
    c := make(chan []int)
    go func() {
        close(c)
    }()
    return c
}
