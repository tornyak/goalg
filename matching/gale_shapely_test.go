package matching_test

import (
	"testing"
	"github.com/tornyak/goalg/matching"
)

func TestGaleShapely_Match(t *testing.T) {
	menWishes := [][]int{
		{1,0,2,3},
		{3,0,1,2},
		{0,2,1,3},
		{1,2,0,3},
	}

	womenWishes := [][]int{
		{0,2,1,3},
		{2,3,0,1},
		{3,1,2,0},
		{2,1,0,3},
	}

	expectedMatches := []int{0,3,2,1}

	gs, _ := matching.InitGaleShapely(menWishes, womenWishes)
	result := gs.Match()

	for i := range result {
		if expectedMatches[i] != result[i] {
			t.Errorf("i=%v expected=%v got=%v", i, expectedMatches[i], result[i])
		}
	}
	t.Log(result)
}

