package util

import (
	"math/rand"
	"time"
)

// Reverse does in place reversing elements of a slice
func Reverse(a []interface{}) {
	length := len(a)
	for i := 0; i < length/2; i++ {
		a[i], a[length-i-1] = a[length-i-1], a[i]
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
