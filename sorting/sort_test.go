package sorting_test

/*
 * Data and tests are from golang sort package
 */

import (
	"math"
	"testing"
	"sort"
	"github.com/tornyak/goalg/sorting"
	"math/rand"
	"strconv"
	"github.com/tornyak/goalg/util"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}


// sorting algorithms
var algorithms = []func(sort.Interface){
	sorting.Bubble,
	sorting.Insertion,
	sorting.Shell,
	sorting.Selection,
}

func TestSortIntSlice(t *testing.T) {
	data := ints

	for _, alg := range algorithms {
		a := sort.IntSlice(data[:])
		alg(a)
		if !sort.IsSorted(a) {
			t.Errorf("%v\n", util.GetFuncName(alg))
			t.Errorf("sorted: %v\n", ints)
			t.Errorf("   got: %v\n", data)
		}
	}
}

func TestSortFloat64Slice(t *testing.T) {
	data := float64s

	for _, alg := range algorithms {
		a := sort.Float64Slice(data[:])
		alg(a)
		if !sort.IsSorted(a) {
			t.Errorf("%v\n", util.GetFuncName(alg))
			t.Errorf("sorted: %v\n", float64s)
			t.Errorf("   got: %v\n", data)
		}
	}
}

func TestSortStringSlice(t *testing.T) {
	data := strings

	for _, alg := range algorithms {
		a := sort.StringSlice(data[:])
		alg(a)
		if !sort.IsSorted(a) {
			t.Errorf("%v\n", util.GetFuncName(alg))
			t.Errorf("sorted: %v\n", strings)
			t.Errorf("   got: %v\n", data)
		}
	}
}

func TestInts(t *testing.T) {
	data := ints

	for _, alg := range algorithms {
		sorting.Ints(data[:], alg)
		if !sort.IntsAreSorted(data[:]) {
			t.Errorf("%v\n", util.GetFuncName(alg))
			t.Errorf("sorted: %v\n", ints)
			t.Errorf("   got: %v\n", data)
		}
	}
}

func TestFloat64s(t *testing.T) {
	data := float64s

	for _, alg := range algorithms {
		sorting.Float64s(data[:], alg)
		if !sort.Float64sAreSorted(data[:]) {
			t.Errorf("%v\n", util.GetFuncName(alg))
			t.Errorf("sorted: %v\n", float64s)
			t.Errorf("   got: %v\n", data)
		}
	}
}

func TestStrings(t *testing.T) {
	data := strings

	for _, alg := range algorithms {
		sorting.Strings(data[:], alg)
		if !sort.StringsAreSorted(data[:]) {
			t.Errorf("%v\n", util.GetFuncName(alg))
			t.Errorf("sorted: %v\n", strings)
			t.Errorf("   got: %v\n", data)
		}
	}
}

func TestSortLarge_Random(t *testing.T) {
	n := 10000
	if testing.Short() {
		n /= 100
	}
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}
	if sort.IntsAreSorted(data) {
		t.Fatalf("terrible rand.rand")
	}

	for _, alg := range algorithms {
		sorting.Ints(data, alg)
		if !sort.IntsAreSorted(data) {
			t.Errorf("%v failed to sort 10K ints\n", util.GetFuncName(alg))
		}
	}
}

func benchInt(b *testing.B, size int, algorithm func(sort.Interface)) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, size)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		sorting.Ints(data, algorithm)
		b.StopTimer()
	}
}

func benchString(b *testing.B, size int, algorithm func(sort.Interface)) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, size)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		sorting.Strings(data, algorithm)
		b.StopTimer()
	}
}

func BenchmarkBubbleSortString1K(b *testing.B) { benchString(b, 1<<10, sorting.Bubble) }
func BenchmarkInsertionSortString1K(b *testing.B) { benchString(b, 1<<10, sorting.Insertion) }
func BenchmarkSelectionSortString1K(b *testing.B) { benchString(b, 1<<10, sorting.Selection) }
func BenchmarkShellSortString1K(b *testing.B) { benchString(b, 1<<10, sorting.Shell) }

func BenchmarkBubbleSortInt1K(b *testing.B) { benchInt(b, 1<<10, sorting.Bubble) }
func BenchmarkInsertionSortInt1K(b *testing.B) { benchInt(b, 1<<10, sorting.Insertion) }
func BenchmarkSelectionSortInt1K(b *testing.B) { benchInt(b, 1<<10, sorting.Selection) }
func BenchmarkShellSortInt1K(b *testing.B) { benchInt(b, 1<<10, sorting.Shell) }

func BenchmarkBubbleSortInt64K(b *testing.B) { benchInt(b, 1<<16, sorting.Bubble) }
func BenchmarkInsertionSortInt64K(b *testing.B) { benchInt(b, 1<<16, sorting.Insertion) }
func BenchmarkSelectionSortInt64K(b *testing.B) { benchInt(b, 1<<16, sorting.Selection) }
func BenchmarkShellSortInt64K(b *testing.B) { benchInt(b, 1<<16, sorting.Shell) }


