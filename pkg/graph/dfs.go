package graph

// DFS seraches a graph in deph fist search manner 
type DFS struct {
	graph *Graph
	marked []bool
	count int
	src int
}

// Init initializes/clears the vistedItems and items array
func (d *DFS) Init() *DFS {
	d.marked = make([]bool, d.graph.V)
	d.search()
	return d
}

// NewDFS returns a new DFS initialized with state and algorithm
func NewDFS(g *Graph, s int) *DFS {
	d := &DFS{graph:g, src: s}
	return d.Init()
}

func (d *DFS) search() {
	src := d.src
	graph := d.graph
	d.marked[src] = true
	d.count++

	itemsInBag := graph.adj[src].Iterate()
	for i := 0; i < graph.adj[src].Size(); i++ {
		if d.marked[itemsInBag[i]] {
			continue
		}

		d.marked[itemsInBag[i]] = true
		d.count++
		d.dfs(itemsInBag[i])
	}
}

func (d *DFS)dfs(itemInBag int) {
	graph := d.graph
	if graph.adj[itemInBag].Size() == 0 {
		return
	}

	itemsInBag := graph.adj[itemInBag].Iterate()
	for i := 0; i < graph.adj[itemInBag].Size(); i++ {
		if d.marked[itemsInBag[i]] {
			continue
		}

		d.marked[itemsInBag[i]] = true
		d.count++
		d.dfs(itemsInBag[i])
	}
}

// Marked check if a vertex is marked or not
func (d *DFS) Marked(v int) bool {
	return d.marked[v]
}

// Count returns the number of attached vertices
func (d *DFS) Count() int {
	return d.count
}