package graph

type Edge struct {
	v, w int
}

// Graph is interface that implementation of
// undirected graph should support
type Graph interface {
	GetNumVertices() int
	GetNumEdges() int
	AddEdge(v, w int)  Graph  // add edge v-w to the Graph
	Adj(v int) []int // Vertices adjacent to v
}

// Degree computes the degree of vertex v in graph g
func Degree(g Graph, v int) int {
	return len(g.Adj(v))
}

// MaxDegree computes maximum degree in graph g
// returns tuple: max degree vertex, max degree
func MaxDegree(g Graph) (int, int) {
	maxV, maxD := -1, -1

	for v := 0; v < g.GetNumVertices(); v++ {
		if dgr := Degree(g, v); dgr > maxD {
			maxD, maxV = dgr, v
		}
	}
	return maxV, maxD
}

// AverageDegree computes average degree of the Graph g
func AverageDegree(g Graph) float64 {
	return 2.0 * float64(g.GetNumEdges()) / float64(g.GetNumVertices())
}

// NumOfSelfLoops counts self-loops in graph g
func NumOfSelfLoops(g Graph) int {
	cnt := 0
	for v := 0; v < g.GetNumVertices(); v++ {
		for _, w := range g.Adj(v) {
			if v == w {
				cnt++;
			}
		}
	}
	return cnt/2
}
