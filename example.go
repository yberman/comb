package main

import (
	"fmt"

	"github.com/yberman/comb"
)

func main() {
	fmt.Println(comb.Binomial(4, 2))
	for c := range comb.combinations(4, 2) {
		fmt.Printf("%+v\n", c)
	}
}
