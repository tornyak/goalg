package sort
import "sort"


// Insertion sort
func InsertionSort(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// ShellSort
func ShellSort(data sort.Interface) {

	N := data.Len()
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}

	for h > 0 {
		// Insertion sort
		for i := h; i < N; i++ {
			for j := i; j >= h && data.Less(j, j-h); j-=h {
				data.Swap(j, j-h)
			}
		}
		h = h/3
	}
}
