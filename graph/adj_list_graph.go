package graph

import (
	"fmt"
	"strconv"
)

// make sure that AdjListGraph implements Graph interface
//var _ AdjListGraph = (*Graph)(nil)

// AdjListGraph implements Graph interface
// using adjacency list graph representation
type AdjListGraph struct {
	v   int
	e   int
	adj [][]int
}

// NewAdjListGraph creates an empty AdjListGraph containing v vertices and returns pointer to it
// v - number of vertices in the Graph
func NewAdjListGraph(v int) Graph {
	g := &AdjListGraph{
		v:   v,
		adj: make([][]int, v),
	}
	for i := 0; i < v; i++ {
		g.adj[i] = make([]int,0)
	}
	return g
}

// NewAdjListGraph creates an empty AdjListGraph containing v vertices and returns pointer to it
// v - number of vertices in the Graph
func NewAdjListGraphEdges(edges []Edge) Graph {
	g := NewAdjListGraph(len(edges))

	for _, e := range edges {
		g.AddEdge(e.v, e.w)
	}

	return g
}

func (g *AdjListGraph) GetNumVertices() int { return g.v }

func (g *AdjListGraph) GetNumEdges() int { return g.e }

func (g *AdjListGraph) AddEdge(v, w int) Graph {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
	return g
}

func (g *AdjListGraph) Adj(v int) []int {
	return g.adj[v]
}

func (g *AdjListGraph)String() string {
	s := fmt.Sprintf("%v vertices, %v edges\n", g.GetNumVertices(), g.GetNumEdges())
	for v := 0; v < g.GetNumVertices(); v++ {
		s += ": "
		for _, w := range g.Adj(v) {
			s += strconv.Itoa(w) + " "
		}
		s += "\n"
	}
	return s
}
