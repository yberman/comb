package comb

type Set map[int]struct{}

func (s Set) In(n int) bool {
	_, ok := s[n]
	return ok
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) Add(n int) {
	s[n] = struct{}{}
}

func (s Set) Del(n int) bool {
	r := s.In(n)
	delete(s, n)
	return r
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) I() chan int {
	c := make(chan int)
	go func() {
		for n, _ := range s {
			c <- n
		}
		close(c)
	}()
	return c
}

func (s1 Set) And(s2 Set) Set {
	s := Set{}
	if s1.Size() > s2.Size() {
		s1, s2 = s2, s1
	}
	for n := range s1.I() {
		if s2.In(n) {
			s.Add(n)
		}
	}
	return s
}

func (s1 Set) Or(s2 Set) Set {
	s := Set{}
	for n := range s1.I() {
		s.Add(n)
	}
	for n := range s2.I() {
		s.Add(n)
	}
	return s
}

func (s1 Set) Not(s2 Set) Set {
	s := Set{}
	for n := range s1.I() {
		if !s2.In(n) {
			s.Add(n)
		}
	}
	return s
}
