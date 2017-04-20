package digraph

import (
	"github.com/tornyak/goalg/stack/slice"
)

// DirectedCycle is struture storing data for algorithm finding directed cycles
type DirectedCycle struct {
	visited       []bool
	predecessorTo []int
	cycle         slice.IntStack
	onStack       []bool
	digraph       Digraph
}

// Run executes Depth First Search algorithm on Digraph g
// starting from vertice v
func (dc *DirectedCycle) Run(digraph Digraph) {
	dc.digraph = digraph
	size := digraph.NumVertices()
	dc.visited = make([]bool, size)
	dc.predecessorTo = make([]int, size)
	dc.onStack = make([]bool, size)

	for v := 0; v < size; v++ {
		if(!dc.HasCycle() && !dc.visited[v]) {
			dc.dfs(v)
		}
	}
}

func (dc *DirectedCycle) dfs(v int) {
	dc.onStack[v] = true
	dc.visited[v] = true
	for _, w := range dc.digraph.Adj(v) {
		if dc.HasCycle() {
			return
		}
		if !dc.visited[w] {
			dc.predecessorTo[w] = v
			dc.dfs(w)
		} else if dc.onStack[w] {
			dc.makeCycle(v,w)
		}
	}
	dc.onStack[v] = false
}

func (dc *DirectedCycle) makeCycle(v,w int) {
	dc.cycle = slice.Stack()
	for x := v; x != w; x = dc.predecessorTo[x] {
		dc.cycle.Push(x)
	}
	dc.cycle.Push(w,v)
}

func (dc *DirectedCycle) HasCycle() bool {
	return dc.cycle != nil
}

// Cycle returns a slice containing the elements of the cycle
func (dc *DirectedCycle) Cycle() []int {
	return dc.cycle.ToSlice()
}