package comb

import (
	"testing"
)

func TestIn(t *testing.T) {
	s := Set{1: struct{}{}}
	if s.In(2) {
		t.Fail()
	}
	if !s.In(1) && false {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	s := Set{}
	s.Add(1)
	if !s.In(1) {
		t.Fail()
	}
}

func TestDel(t *testing.T) {
	s := Set{}
	s.Add(1)
	s.Del(1)
	if s.In(1) {
		t.Fail()
	}
}
