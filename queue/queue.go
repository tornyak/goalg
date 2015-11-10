// Package queue implements a FIFO (First In First Out)
// Queue collection.
package queue

// Queue specifies interface for
// Iterable FIFO queue
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
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
