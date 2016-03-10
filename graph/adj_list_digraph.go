package graph

import "github.com/tornyak/goalg/bag"

// AdjListDigraph implements Digraph interface
// using adjacency list graph mode
type AdjListDigraph struct {
	v   int
	e   int
	adj []bag.Bag
}

// NewAdjListDigraph creates an empty AdjListDigraph containing v vertices and returns pointer to it
// v - number of vertices in the Graph
func NewAdjListDigraph(v int) Digraph {
	g := &AdjListDigraph{
		v:   v,
		adj: []bag.Bag{},
	}
	for i := 0; i < v; i++ {
		g.adj[i] = bag.NewLinkedBag()
	}
	return g
}

func (g *AdjListDigraph) GetNumVertices() int { return g.v }

func (g *AdjListDigraph) GetNumEdges() int { return g.e }

func (g *AdjListDigraph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.e++
}

func (g *AdjListDigraph) Adj(v int) Iterable {
	return g.adj[v]
}

func (g *AdjListDigraph) Reverse() Digraph {
	revGraph := NewAdjListDigraph(g.v)
	for _, v := range g.adj {
		it := v.GetIterator()
		for it.HasNext() {
			w := it.Next().(int)
			revGraph.AddEdge(w, v)
		}
	}
	return revGraph
}
