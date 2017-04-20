package graph

import (
	"testing"
	"reflect"
)

func TestDFS(t *testing.T) {
	g := NewAdjListGraphEdges([]Edge{
		{0,1},{0,2},{0,5},{2,1},{2,3},{2,4},{3,4},{3,5},
	})

	marked, edgeTo := DFS(g, 0)

	var markedTests = []struct {
		vertex        int // input
		expected bool // expected result
	}{
		{0, true},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
	}

	for _, mt := range markedTests {
		if marked[mt.vertex] != mt.expected {
			t.Errorf("marked[%d] expected: %v, was: %v", mt.vertex, marked[mt.vertex], mt.expected)
		}
	}

	var pathTests = []struct {
		vertex   int // input
		expected []int // expected result
	}{
		{0, []int{0}},
		{1, []int{0,2,1}},
		{2, []int{0,2}},
		{3, []int{0,2,3}},
		{4, []int{0,2,3,4}},
		{5, []int{0,2,3,5}},
	}

	for _, pt := range pathTests {
		if  reflect.DeepEqual(PathTo(0, pt.vertex, marked, edgeTo), pt.expected) {
			t.Errorf("marked[%d] expected: %v, was: %v")
		}
	}
}
