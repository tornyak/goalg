package digraph// Digraph is interface that implementation of
// directed graph should support
type Digraph interface {
	GetNumVertices() int
	GetNumEdges() int
	AddEdge(v, w int)   // add edge v->w to the Digraph
	Adj(v int) []int    // Vertices adjacent to v
	Reverse()
	String() string
}
