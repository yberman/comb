package comb

import "testing"

func TestTo(t *testing.T) {
	b := StringBox{}
	n1 := b.To("one")
	n2 := b.To("two")
	if n1 != 1 || n2 != 2 || n1 != b.To("one") {
		t.Fail()
	}
}

func TestFrom(t *testing.T) {
	b := StringBox{}
	s := "one"
	n := b.To(s)
	if s2, _ := b.From(n); s2 != s {
		t.Fail()
	}
}
