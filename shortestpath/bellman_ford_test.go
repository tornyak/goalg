package shortestpath

import (
	"testing"
	"os"
	"github.com/tornyak/goalg/digraph"
	"log"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	//myTeardownFunction()
	os.Exit(retCode)
}

var testGraph1 digraph.EdgeWeightedDigraph

func setUp() {
	testGraph1 = digraph.NewEdgeWeightedDigraph(8)
	testGraph1.AddEdge(0,2,0.26).AddEdge(0,4,0.38)
	testGraph1.AddEdge(1,3,0.29)
	testGraph1.AddEdge(2,7,0.34)
	testGraph1.AddEdge(3,6,0.52)
	testGraph1.AddEdge(4,7,0.37).AddEdge(4,5,0.35)
	testGraph1.AddEdge(5,1,0.32).AddEdge(5,7,0.28).AddEdge(5,4,0.35)
	testGraph1.AddEdge(6,4,0.93).AddEdge(6,0,0.58).AddEdge(6,2,0.40)
	testGraph1.AddEdge(7, 3, 0.39).AddEdge(7,5,0.28)
}


func TestBellmanFord_Run(t *testing.T) {
	bf := &BellmanFord{}
	bf.Run(0, testGraph1)
	log.Printf("distanceTo: %v", bf.distanceTo)
	log.Printf("predecessorTo: %v", bf.predecessorTo)
}
