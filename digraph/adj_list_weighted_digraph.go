package digraph

import (
	"bytes"
	"strconv"
)

type edgeWeightedDigraph struct{
	numVertices int
	numEdges int
	adj [][]*Edge
}

// NewAdjListDigraph creates an empty AdjListDigraph containing v vertices and returns pointer to it
// v - number of vertices in the Digraph
func NewEdgeWeightedDigraph(v int) EdgeWeightedDigraph {
	g := &edgeWeightedDigraph{
		numVertices:   v,
		adj: make([][]*Edge, v),
	}
	return g
}

func (g *edgeWeightedDigraph) NumVertices() int {
	return g.numVertices
}

func (g *edgeWeightedDigraph) NumEdges() int {
	return g.numEdges
}

func (g *edgeWeightedDigraph)AddEdge(v, w int, weight float64) EdgeWeightedDigraph {
	return g.addEdge(NewEdge(v, w, weight))
}

func (g *edgeWeightedDigraph)addEdge(edges ...*Edge) EdgeWeightedDigraph {
	for _, e := range edges {
		g.adj[e.From] = append(g.adj[e.From], e)
		g.numEdges++
	}
	return g
}

func (g *edgeWeightedDigraph)Adj(v int) []*Edge {
	return g.adj[v]
}

func (g *edgeWeightedDigraph)Edges() []*Edge {
	var edges []*Edge
	for _, adjEdges := range g.adj {
		for _, e := range adjEdges {
			edges = append(edges, NewEdge(e.From, e.To, e.Weight))
		}
	}
	return edges
}

func (g *edgeWeightedDigraph)Reverse() EdgeWeightedDigraph {
	r := NewEdgeWeightedDigraph(g.numVertices)
	for v := 0; v < g.numVertices; v++ {
		for _, e := range g.adj[v] {
			r.AddEdge(e.To, e.From, e.Weight)
		}
	}
	return r
}

func (g *edgeWeightedDigraph)String() string {
	var buffer bytes.Buffer
	for v := 0; v < g.numVertices; v++ {
		buffer.WriteString(strconv.Itoa(v))
		buffer.WriteString(": ")
		for _, e := range g.Adj(v) {
			buffer.WriteString(e.String())
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
