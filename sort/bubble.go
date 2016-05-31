package sort
import "sort"

// BubbleSort implementation
func BubbleSort(data sort.Interface) {

	for i := 0; i < data.Len(); i++ {
		for j := data.Len() - 1; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
