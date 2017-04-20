package digraph

import (
	"testing"
	"reflect"
)

func TestDirectedCycle(t *testing.T) {
	directedCycle := &DirectedCycle{}
	directedCycle.Run(testDigraph1)

	if directedCycle.HasCycle() == false {
		t.Errorf("HasCycle failed to detect cycle")
	}

	expected := []int{3,4,5,3}

	if !reflect.DeepEqual(directedCycle.Cycle(), expected) {
		t.Errorf("Cycle failed: expected: %v got: %v", expected, directedCycle.Cycle())
	}

}
