// Package stack implements a LIFO (Last In First Out) Stack
// collection.
package stack

// Stack interface that different Stack
// implementations must implement
type Stack interface {
	Push(...interface{}) Stack
	Pop() (interface{}, bool)
	IsEmpty() bool
	Size() int
}
