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
