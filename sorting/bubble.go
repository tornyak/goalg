package sorting

import "sort"

// BubbleSort implementation
func Bubble(a sort.Interface) {

	for i := 0; i < a.Len(); i++ {
		for j := a.Len() - 1; j > i; j-- {
			if a.Less(j, j-1) {
				a.Swap(j, j-1)
			}
		}
	}
}
