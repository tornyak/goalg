package digraph

import "fmt"

// Digraph is interface that implementation of
// directed graph should support
type Digraph interface {
	NumVertices() int
	NumEdges() int
	AddEdge(v int, w ...int) Digraph  // add edges v->w1, v->w2 ... to the Digraph
	Adj(v int) []int    // Vertices adjacent to v
	Reverse() Digraph
}

// EdgeWeightedDigraph is interface that implementation of
// edge weighted directed graph should support
type EdgeWeightedDigraph interface {
	NumVertices() int
	NumEdges() int
	AddEdge(v,w int, weight float64) EdgeWeightedDigraph
	Adj(v int) []*Edge    // Vertices adjacent to v
	Reverse() EdgeWeightedDigraph
}

type Edge struct {
	From, To int
	Weight float64
}

func NewEdge(v, w int, weight float64) *Edge {
	return &Edge{From: v, To: w, Weight: weight}
}

func (e *Edge)String() string {
	return fmt.Sprintf("%v->%v %.2f", e.From, e.To, e.Weight)
}


