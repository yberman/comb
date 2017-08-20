package comb

type Graph map[int]Set

func (g Graph) empty() bool {
	return len(g) == 0
}

func (g Graph) AddVertex(v int) {
	_, ok := g[v]
	if !ok {
		g[v] = Set{}
	}
}

func (g Graph) Add(v, w int) {
	g.AddVertex(v)
	g.AddVertex(w)
	g[v].Add(w)
	g[w].Add(v)
}

type StringGraph struct {
	intGraph Graph
	box      StringBox
}

func (g *StringGraph) empty() bool {
	return g.intGraph.empty()
}

func (g *StringGraph) init() {
	if g.empty() {
		g.intGraph = Graph{}
		g.box = StringBox{}
	}
}

func (g *StringGraph) AddVertex(s string) {
	g.init()
	g.intGraph.AddVertex(g.box.To(s))
}

func (g *StringGraph) Add(s, t string) {
	g.init()
	g.intGraph.Add(g.box.To(s), g.box.To(s))
}
