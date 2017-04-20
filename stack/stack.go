// Package stack implements a LIFO (Last In First Out) Stack
// collection.
package stack

// Stack interface that different Stack
// implementations must implement
type Stack interface {
	Push(...interface{}) Stack
	Pop() interface{}
	IsEmpty() bool
	Size() int
}



