package slice

import (
	"fmt"
	"math/rand"
	"time"
)

// Reverse does in place reversing elements of a slice
func Reverse(a []interface{}) {
	length := len(a)
	for i := 0; i < length/2; i++ {
		opp := length - i -1
		a[i], a[opp] = a[opp], a[i]
	}
}

// Shuffle implements in place shuffling algorithm
// known as Fisher–Yates  or Knuth shuffle
func Shuffle(a []interface{}) {
	rand.Seed(time.Now().UnixNano())
	n := len(a)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// Cut removes range from i-th to j-th elements from slice
// Usuallly i should be lower than j but if that is not case they
// are swapped
func Cut(a []interface{}, i, j int) ([]interface{}, error) {
	lenA := len(a)
	if i < 0 || i >= lenA {
		return a, fmt.Errorf("Index out of range: %v, len: %v", i, lenA)
	}
	if j < 0 || j >= lenA {
		return a, fmt.Errorf("Index out of range: %v, len: %v", i, lenA)
	}

	if i > j {
		i, j = j, i
	}

	copy(a[i:], a[j:])
	for k, n := len(a) - j + i, len(a); k < n; k++ {
		a[k] = nil // or the zero value of T
	}
	a = a[:len(a) - j + i]
	return a, nil
}

// Delete i-th element from slice
func Delete(a []interface{}, i int) ([]interface{}, error) {
	if i < 0 || i >= len(a) {
		return a, fmt.Errorf("Index out of range: %v, len: %v", i, len(a))
	}

	copy(a[i:], a[i+1:])
	a[len(a)-1] = nil // or the zero value of T
	a = a[:len(a)-1]
	return a, nil
}

func DeleteInt(a []int, i int) ([]int, error) {
	if i < 0 || i >= len(a) {
		return a, fmt.Errorf("Index out of range: %v, len: %v", i, len(a))
	}

	copy(a[i:], a[i+1:])
	a[len(a)-1] = 0 // or the zero value of T
	a = a[:len(a)-1]
	return a, nil
}

// CopyInt returns a copy of slice a
func CopyInt(a []int) []int {
	aCopy := make([]int, len(a))
	copy(aCopy, a)
	return aCopy
}


func EqualInt(a, b []int) bool {

	if len(a) != len (b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) == 0 {
		return true
	}


	if &a[0] == &b[0] {
		return true
	}

	b = b[:len(a)] // To use BCE (Boundary Check Elimination)
	// TODO if it does not work use range
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}