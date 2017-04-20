package sorting_test

import (
	"testing"
	"github.com/tornyak/goalg/sorting"
	"sort"
)


func TestMergeSortIntSlice(t *testing.T) {
	a := ints[:]
	sorting.MergeInts(a)
	if !sort.IntsAreSorted(a){
		t.Errorf("sorted: %v\n", ints)
		t.Errorf("   got: %v\n", a)
	}
}

func TestMergeInts(t *testing.T) {
	a := make([]int, 12)
	b := []int{-3,0,2,4,4,98}
	c := []int{-7,-3,1,2,5,103}


	sorting.Merge(a, b, c)

	if !sort.IntsAreSorted(a){
		t.Errorf("sorted: %v\n", ints)
		t.Errorf("   got: %v\n", a)
	}
}