package sorting

import (
	"fmt"
	"github.com/tornyak/goalg/util/slice"
	"sort"
)


// Merge sort slice a. If a is not slice it will panic
// less returns true if element a[i] < a[j]
func MergeInts(a []int) {

	if len(a) == 0 {
		return
	}

	mergeSortInt(a, slice.CopyInt(a))
	fmt.Println(a)
}

// MergeIntsBottomUp sorts slice a using bottom up merge sort
func MergeIntsBottomUp(a []int){
	if len(a) == 0 {
		return
	}

	mergeSortIntBottom(a, slice.CopyInt(a))
	fmt.Println(a)
}

func mergeSortIntBottom(a, aCopy []int)  {
	for size := 2; size < len(a); size = size*2 {

	}
}

func mergeSortInt(a, aCopy []int) {
	fmt.Printf("mergeSortInt %v %v\n", a, aCopy)

	len := len(a)

	if len < 12 {
		fmt.Printf("Insertion sort: %v\n", a)
		Insertion(sort.IntSlice(a))
		return
	}

	middle := len/2

	// Swap a and aCopy to avoid copying
	mergeSortInt(aCopy[:middle], a[:middle])
	mergeSortInt(aCopy[middle:], a[middle:])

	// is already sorted?
	if a[middle-1] < a[middle] {
		return
	}

	Merge(a, aCopy[:middle], aCopy[middle:])
}


func Merge(a, b, c []int) {

	var i, j, k int
	for i, j, k = 0, 0, 0; j < len(b) && k < len(c); i++ {
		if b[j] < c[k] {
			a[i] = b[j]
			j++
		} else {
			a[i] = c[k]
			k++
		}
	}

	copy(a[i:], b[j:])
	copy(a[i:], c[k:])
}

