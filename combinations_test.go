package comb

import (
	"testing"
)

func makeSliceOfSlice(ch chan []int) [][]int {
	r := make([][]int, 0)
	for s := range ch {
		r = append(r, s)
	}
	return r
}

func equalSlice(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, es1 := range s1 {
		if es1 != s2[i] {
			return false
		}
	}
	return true
}

func equalSliceOfSlice(ss1, ss2 [][]int) bool {
	m := map[int]int{}
	if len(ss1) != len(ss2) {
		return false
	}
	for i, s1 := range ss1 {
		s1Matched := false
		for j, s2 := range ss2 {
			if equalSlice(s1, s2) {
				if _, ok := m[j]; !ok {
					s1Matched = true
					m[j] = i
				}
			}
			if s1Matched {
				break
			}
		}
		if !s1Matched {
			return false
		}
	}
	return true
}

func TestHelpers(t *testing.T) {
	if !equalSlice([]int{}, []int{}) {
		t.Errorf("Helper functions failed\n")
	}
	if !equalSlice([]int{3}, []int{3}) {
		t.Errorf("Helper functions failed\n")
	}
	if equalSlice([]int{3}, []int{}) {
		t.Errorf("Helper functions failed\n")
	}
	if !equalSliceOfSlice([][]int{{}}, [][]int{{}}) {
		t.Errorf("Helper functions failed\n")
	}
	if equalSliceOfSlice([][]int{{3}}, [][]int{{}}) {
		t.Errorf("Helper functions failed\n")
	}
}

func TestBinomial(t *testing.T) {
	b := Binomial(20, 10)
	const target = 184756
	if b != target {
		t.Errorf("Binomial(20, 10) = %d != %d\n", b, target)
	}
}

func TestCominations(t *testing.T) {
	if equalSliceOfSlice(makeSliceOfSlice(Combinations(4, 2)), [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {1, 4}}) {
		t.Errorf("Combinations(4, 2) has error\n")
	}
}
