package digraph

import (
	"bytes"
	"strconv"
)

type digraph struct{
	numVertices int
	numEdges int
	adj [][]int
}

// NewAdjListDigraph creates an empty AdjListDigraph containing v vertices and returns pointer to it
// v - number of vertices in the Digraph
func NewAdjListDigraph(v int) Digraph {
	g := &digraph{
		numVertices:   v,
		adj: make([][]int, v),
	}
	return g
}

func (g *digraph) NumVertices() int {
	return g.numVertices
}

func (g *digraph) NumEdges() int {
	return g.numEdges
}

func (g *digraph)AddEdge(v int, w ...int) Digraph {
	for _, ww := range w {
		g.adj[v] = append(g.adj[v], ww)
		g.numEdges++
	}
	return g
}

func (g *digraph)Adj(v int) []int {
	return g.adj[v]
}

func (g *digraph)Reverse() Digraph {
	r := NewAdjListDigraph(g.numVertices)
	for v := 0; v < g.numVertices; v++ {
		for _, w := range g.adj[v] {
			r.AddEdge(w, v)
		}
	}
	return r
}

func (g *digraph)String() string {
	var buffer bytes.Buffer
	for v := 0; v < g.numVertices; v++ {
		buffer.WriteString(strconv.Itoa(v))
		buffer.WriteString(":")
		for _, w := range g.Adj(v) {
			buffer.WriteString(strconv.Itoa(w))
			buffer.WriteString(" ")
		}

		buffer.WriteString("\n")
	}
	return buffer.String()
}

