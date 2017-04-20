package digraph

import (
	"testing"
	"os"
	"reflect"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	//myTeardownFunction()
	os.Exit(retCode)
}

var testDigraph1 Digraph

func setUp() {
	testDigraph1 = NewAdjListDigraph(13)
	testDigraph1.AddEdge(0, 1, 5)
	testDigraph1.AddEdge(2, 0, 3)
	testDigraph1.AddEdge(3, 5, 2)
	testDigraph1.AddEdge(4, 3, 2)
	testDigraph1.AddEdge(5, 4)
	testDigraph1.AddEdge(6, 9,4,8,0)
	testDigraph1.AddEdge(7, 6,9)
	testDigraph1.AddEdge(8, 6)
	testDigraph1.AddEdge(9, 11, 10)
	testDigraph1.AddEdge(10, 12)
	testDigraph1.AddEdge(11, 4, 12)
	testDigraph1.AddEdge(12, 9)
}

func TestAdjListDigraph_Adj(t *testing.T) {
	var expected = [][]int{
		{1,5},
		nil,
		{0,3},
		{5,2},
		{3,2},
		{4},
		{9,4,8,0},
		{6,9},
		{6},
		{11,10},
		{12},
		{4, 12},
		{9},
	}

	t.Log(testDigraph1)

	for v, exp := range expected {
		if !reflect.DeepEqual(testDigraph1.Adj(v), exp) {
			t.Errorf("Wrong Adj slice returned: %v: expected %v", testDigraph1.Adj(v), exp)
		}
	}
}

func TestAdjListDigraph_Reverse(t *testing.T) {

	var expected = [][]int{
		{2,6},
		{0},
		{3,4},
		{2,4},
		{5,6,11},
		{0,3},
		{7,8},
		nil,
		{6},
		{6,7,12},
		{9},
		{9},
		{10,11},
	}

	reversed := testDigraph1.Reverse()
	t.Log(reversed)

	for v, exp := range expected {
		if !reflect.DeepEqual(reversed.Adj(v), exp) {
			t.Errorf("Wrong Adj slice returned: %v: expected %v", reversed.Adj(v), exp)
		}
	}
}
