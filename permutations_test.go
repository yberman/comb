package comb

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	b := Factorial(10)
	const target = 3628800
	if b != target {
		t.Errorf("Factorial(10) = %d != %d\n", b, target)
	}
}

// func TestPermutations(t *testing.T) {
// 	if equalSliceOfSlice(makeSliceOfSlice(Permutations(3, 2)), [][]int{{0, 1}, {0, 2}, {1, 2}})
// 		t.Errorf("Combinations(4, 2) has error\n")
// 	}
// }
