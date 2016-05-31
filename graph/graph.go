package graph

// Graph is interface that implementation of
// undirected graph should support
type Graph interface {
	GetNumVertices() int
	GetNumEdges() int
	AddEdge(v, w int)   // add edge v-w to the Graph
	Adj(v int) Iterable // Vertices adjacent to v
}

// Digraph is interface that implementation of
// directed graph should support
type Digraph interface {
	GetNumVertices() int
	GetNumEdges() int
	AddEdge(v, w int)   // add edge v->w to the Digraph
	Adj(v int) Iterable // Vertices adjacent to v
	Reverse()
}

// Iterable defines interface that collection
// iterator has to implement
type Iterable interface {
	HasNext() bool
	Next() interface{}
}
