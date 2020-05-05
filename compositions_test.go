package comb

import (
	"testing"
)

func TestCompositionss(t *testing.T) {
	sos := ToSliceOfSlice(Compose(4).Positive(2))
	if !equalSliceOfSlice(sos, [][]int{{1, 3}, {2, 2}, {3, 1}}) {
		t.Error("Positive compositions fails")
	}
}

func TestNonNegCompositionss(t *testing.T) {
	sos := ToSliceOfSlice(Compose(4).NonNegative(2))
	if !equalSliceOfSlice(sos, [][]int{{0, 4}, {1, 3}, {2, 2}, {3, 1}, {4, 0}}) {
		t.Error("Non-negative compositions fails")
	}
}
