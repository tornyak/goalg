// Package queue implements a FIFO (First In First Out) Queue
// collection.
package queue

// Element is an element of the Queue
type element struct {
	next  *element
	value interface{}
}

// A Queue structure
type Queue struct {
	first *element
	last  *element
	size  uint
}

// Iterator data structure
type Iterator struct {
	current *element
}

// Next moves iterator to the next element in collection
// if it exists and returns true, othewise it just returns false
func (it *Iterator) Next() bool {
	if it.current != nil {
		it.current = it.current.next
	}
	if it.current == nil {
		return false
	}
	return true
}

// Value returns the value of the current element or nil
// if current element is nil
func (it *Iterator) Value() interface{} {
	if it.current != nil {
		return it.current.value
	}
	return nil
}

// New creates an empty Queue and returns pointer to it
func New() *Queue { return new(Queue).init() }

func (q *Queue) init() *Queue {
	q.first = nil
	q.last = nil
	q.size = 0
	return q
}

// Enqueue adds item v to the end of the Queue
func (q *Queue) Enqueue(v interface{}) {
	oldLast := q.last
	e := new(element)
	e.value = v
	e.next = q.last
	q.last = e
	// if empty
	if q.size == 0 {
		q.first = e
	} else {
		oldLast.next = e
	}
	q.size++
}

// Dequeue returns the first item from the queue
func (q *Queue) Dequeue() interface{} {
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
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Size returns number of items in the Queue
func (q *Queue) Size() uint {
	return q.size
}

// GetIterator returns a Queue iterator
func (q *Queue) GetIterator() *Iterator {
	it := new(Iterator)
	it.current = new(element)
	it.current.next = q.first
	return it
}
