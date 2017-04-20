package shortestpath

import (
	"github.com/tornyak/goalg/digraph"
	"math"
)

// Implementation of Bellman-Ford algorithm
// for finding single source shortest paths in directed graphs
// O(V E)
type BellmanFord struct {
	source int
	distanceTo []float64
	predecessorTo []int
	graph digraph.EdgeWeightedDigraph
}

func (bf *BellmanFord)Run(s int, g digraph.EdgeWeightedDigraph) {

	// for hops = 0..N-1
	// for all nodes update cost in a way that
	// cost = min current cost or for all edges cost of predecessor + cost of edge
	bf.initialize(s, g)
	for hop := 0; hop < g.NumVertices(); hop++ {
		changed := false
		for v := 0; v < g.NumVertices(); v++ {
			for _, e := range g.Adj(v) {
				if bf.distanceTo[e.To] > (bf.distanceTo[v] + e.Weight) {
					bf.distanceTo[e.To] = bf.distanceTo[v] + e.Weight
					bf.predecessorTo[e.To] = v
					changed = true
				}
			}
		}
		if !changed {
			return
		}
	}
}

func (bf *BellmanFord)initialize(s int, g digraph.EdgeWeightedDigraph) {
	bf.distanceTo = make([]float64, g.NumVertices())
	bf.predecessorTo = make([]int, g.NumVertices())

	for i := 0; i < g.NumVertices(); i++ {
		bf.distanceTo[i] = math.MaxFloat64
		bf.predecessorTo[i] = -1
	}
	bf.distanceTo[s] = 0
	bf.predecessorTo[s] = 0
}
