package graph

import "testing"

func TestAdjListGraph_AddEdge(t *testing.T) {
	g := NewAdjListGraph(3)
	g.AddEdge(0,1).AddEdge(1,2).AddEdge(2,0)

	var tests = []struct {
		vertex        int // input
		expected []int // expected result
	}{
		{0, []int{1, 2}},
		{1, []int{0, 2}},
		{2, []int{1, 0}},
	}

	for _, tt := range tests {
		actual := g.Adj(tt.vertex)
		if len(actual) != len(tt.expected) {
			t.Errorf("Adj(%d): expected %v, actual %v", tt.vertex, tt.expected, actual)
		}
	}
}
