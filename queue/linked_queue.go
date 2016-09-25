package queue

import (
	"fmt"
	"strings"
)

// element represents an element of the Queue
type element struct {
	next  *element
	value interface{}
}

// A linkedQueue is underlying struct for Linked List backed Queue
type linkedQueue struct {
	first *element
	last  *element
	size  int
}

// Linked creates an empty Queue and returns pointer to it
func Linked() Queue { return new(linkedQueue) }

// Enqueue adds an item or list of items v to the end of the Queue
func (q *linkedQueue) Enqueue(items ...interface{}) Queue {
	for _, it := range items {
		oldLast := q.last
		e := &element{
			value: it,
		}
		q.last = e
		// if empty
		if q.size == 0 {
			q.first = e
		} else {
			oldLast.next = e
		}
		q.size++
	}
	return q
}

// Dequeue returns the first item from the queue
func (q *linkedQueue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	}
	e := q.first
	q.first = e.next
	if q.size == 1 {
		q.last = nil
	}
	q.size--
	return e.value
}

// IsEmpty returns true if the Queue is empty, othewise false
func (q *linkedQueue) IsEmpty() bool {
	return q.size == 0
}

// Size returns number of items in the Queue
func (q *linkedQueue) Size() int {
	return q.size
}

// String returns formatted string representing linkedQueue and its elements
func (b *linkedQueue) String() string  {
	ret := "["
	for e := b.first; e != nil; e = e.next {
		ret+= fmt.Sprintf("%v ", e.value)
	}
	return strings.TrimSpace(ret) + "]"
}