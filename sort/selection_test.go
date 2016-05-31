package sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestSelectionSortInts(t *testing.T) {
	var testCases = []struct {
		data []int
		exp  []int
	}{
		{
			data: nil,
			exp:  nil,
		},
		{
			data: []int{},
			exp:  []int{},
		},
		{
			data: []int{5},
			exp:  []int{5},
		},
		{
			data: []int{5, 7, 1, 43, 11111, 21, 98, -23, -776, 0, 0, 1},
			exp:  []int{-776, -23, 0, 0, 1, 1, 5, 7, 21, 43, 98, 11111},
		},
	}

	for _, tc := range testCases {
		SelectionSort(sort.IntSlice(tc.data))
		if !reflect.DeepEqual(tc.exp, tc.data) {
			t.Errorf("Exp: %v, Got: %v\n", tc.exp, tc.data)
		}
	}
}
