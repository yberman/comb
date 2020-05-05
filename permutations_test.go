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

func TestPermutations(t *testing.T) {
	if len(ToSliceOfSlice(Permutations(5))) != Factorial(5) {
		t.Errorf("permutations has error %d %d\n", Factorial(5), len(ToSliceOfSlice(Permutations(5))))
	}
}
