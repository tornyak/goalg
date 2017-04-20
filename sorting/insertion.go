package sorting

import (
	"sort"
	"fmt"
)


// Insertion sort
func Insertion(a sort.Interface) {
	for i := 1; i < a.Len(); i++ {
		for j := i; j > 0 && a.Less(j, j-1); j-- {
			fmt.Printf("Swap i=%d j=%d a=%v\n", i, j , a)
			a.Swap(j, j-1)
		}
	}
}
