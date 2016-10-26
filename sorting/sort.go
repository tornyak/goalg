package sorting

import "sort"

// Ints sorts a slice of ints in increasing order.
func Ints(a []int, sortFunc func(sort.Interface)) {
	sortFunc(sort.IntSlice(a))
}

// Float64s sorts a slice of float64s in increasing order.
func Float64s(a []float64, sortFunc func(sort.Interface)) {
	sortFunc(sort.Float64Slice(a))
}

// Strings sorts a slice of strings in increasing order.
func Strings(a []string, sortFunc func(sort.Interface)) {
	sortFunc(sort.StringSlice(a))
}