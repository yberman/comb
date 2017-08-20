package comb

type StringBox struct {
	to   map[string]int
	from map[int]string
}

func (b *StringBox) init() {
	if b.empty() {
		b.to = map[string]int{}
		b.from = map[int]string{}
	}
}

func (b *StringBox) empty() bool {
	if b.to == nil && b.from == nil {
		return true
	}
	return false
}


func (b *StringBox) To(s string) int {
	b.init()
	_, ok := b.to[s]
	if !ok {
		b.to[s] = len(b.to) + 1
		b.from[b.to[s]] = s
	}
	return b.to[s]
}

func (b *StringBox) From(n int) (string, bool) {
	if b.empty() {
		return "", false
	}
	s, ok := b.from[n]
	return s, ok
}

