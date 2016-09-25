// Package queue implements a FIFO (First In First Out)
// Queue collection.
package queue

// Queue specifies interface for
// Iterable FIFO queue
type Queue interface {
	Enqueue(...interface{}) Queue
	Dequeue() interface{}
	IsEmpty() bool
	Size() int
}
