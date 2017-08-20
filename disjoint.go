package comb

type Disjoint struct{}

func (d *Disjoint) Add(n int) {
}

func (d *Disjoint) root(n int) int {
	return 0
}

func (d *Disjoint) Join(n1, n2 int) {
}

func (d *Disjoint) Test(n1, n2 int) bool {
	return false
}

func (d *Disjoint) Sets() []Set {
	return []Set{}
}
