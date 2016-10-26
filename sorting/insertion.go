package sorting

import "sort"


// Insertion sort
func Insertion(a sort.Interface) {
	for i := 2; i < a.Len(); i++ {
		// insert a[j] into sorted slice a[0:j]
		for j := i; j > 0 && a.Less(j, j-1); j-- {
			a.Swap(j, j-1)
		}
	}
}
