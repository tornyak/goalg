package sorting

import "sort"

// Selection sort
func Selection(data sort.Interface) {
	dataLen := data.Len()
	for i := 0; i < dataLen; i++ {
		minIndex := i
		for j := i + 1; j < dataLen; j++ {
			if data.Less(j, minIndex) {
				minIndex = j
			}
		}
		data.Swap(i, minIndex)
	}
}
