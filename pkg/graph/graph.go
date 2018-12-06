package graph

import "interview/pkg/bag"

// Graph datatype
type Graph struct {
	V int
	E int
	adj []bag.Bag
}

// Init initiates/clears a graph data type
func (g *Graph) Init (V int) *Graph {
	g.V = V
	g.E = 0
	g.adj = make([]bag.Bag, V)
	return g
}

// New creates a graph
func New(V int) *Graph {
	return new(Graph).Init(V)
}

// AddEdge adds a undirected edge to the graph
func (g *Graph) AddEdge(v int, w int) {
	g.E++
	g.adj[v].Add(w)
	g.adj[w].Add(v)
}

// Iterate iterates through the graph
func (g *Graph) Iterate() [][]int{
	items := make([][]int, g.V)
	for i := 0; i < g.V; i++ {
		items[i] = make([]int, g.adj[i].Size())		
	}

	for i := 0; i < g.V; i++ {
			items[i] = g.adj[i].Iterate()
	}
	return items
}