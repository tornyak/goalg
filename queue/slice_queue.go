package queue

// SliceQueue represents Queue based on slices
type SliceQueue []interface{}

// LinkedQueueIterator data structure
type SliceQueueIterator struct {
	current int
	a SliceQueue
}

// NewSliceQueue creates an empty Queue and returns pointer to it
func NewSliceQueue() Queue { return make(SliceQueue) }

// Enqueue adds item v to the end of the Queue
func (q *SliceQueue) Enqueue(v interface{}) {
	q = append(q, v)
}

// Dequeue returns the first item from the queue
func (q *SliceQueue) Dequeue() interface{} {
	retValue := q[0]
	q = q[1:]
	return retValue
}

// IsEmpty returns true if the Queue is empty, otherwise false
func (q *SliceQueue) IsEmpty() bool {
	return len(q) == 0
}

// Size returns number of items in the Queue
func (q *SliceQueue) Size() int {
	return len(q)
}

// GetIterator returns a Queue iterator
func (q *SliceQueue) GetIterator() Iterable {
	return &SliceQueueIterator{
		current: 0,
		a: q[:],
	}
}

// HasNext moves iterator to the next element in collection
// if it exists and returns true, othewise it just returns false
func (it *SliceQueueIterator) HasNext() bool {
	return it.current < len(it.a) - 1
}

// Next returns the value of the current element or nil
// if current element is nil
func (it *SliceQueueIterator) Next() interface{} {
	if it.current < len(it.a) {
		curr := it.current
		it.current++
		return it.a[curr]
	}
	return nil
}

