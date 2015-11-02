package util

// Reverse does in place reversing elements of a slice
func Reverse(a []interface{}) {
	length := len(a)
	for i := 0; i < length/2; i++ {
		a[i], a[length-i-1] = a[length-i-1], a[i]
	}
}
