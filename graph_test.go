package comb

import "testing"

func TestGraph(t *testing.T) {
	g := StringGraph{}
	g.Add("cabbage", "goat")
	g.Add("wolf", "goat")
}
