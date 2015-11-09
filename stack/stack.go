// Package stack implements a LIFO (Last In First Out) Stack
// collection.
package stack

// Stack interface that different Stack
// implementations must implement
type Stack interface {
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
	GetIterator() Iterable
}

// Iterable defines interface that collection
// iterator has to implement
type Iterable interface {
	HasNext() bool
	Next() interface{}
}
