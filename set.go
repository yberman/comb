package comb

type Set map[int]struct{}

func (s Set) In(n int) bool {
	_, ok := s[n]
	return ok
}

func (s Set) Add(n int) {
	s[n] = struct{}{}
}

func (s Set) Del(n int) bool {
	r := s.In(n)
	delete(s, n)
	return r
}

func (s1 Set) And(s2 Set) Set {
	return Set{}
}

func (s1 Set) Or(s2 Set) Set {
	return Set{}
}

func (s1 Set) Not(s2 Set) Set {
	return Set{}
}
