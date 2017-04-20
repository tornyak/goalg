package digraph

// DirectedDFS is struture storing data of DFS algorithm
// for directed graphs
type DirectedDFS struct {
	marked[] bool
	digraph Digraph
}

// Run executes Depth First Search algorithm on Digraph g
// starting from vertice v
func (d *DirectedDFS) Run(digraph Digraph, v int) {
	d.marked = make([]bool, digraph.NumVertices())
	d.digraph = digraph
	d.dfs(v)
}

func (d *DirectedDFS) dfs(v int) {
	d.marked[v] = true
	for _, w := range d.digraph.Adj(v) {
		if !d.marked[w] {
			d.dfs(w)
		}
	}
}

// Marked returns true if node v was marked during Run
func (dfs *DirectedDFS) Marked(v int) bool {
	return dfs.marked[v]
}
