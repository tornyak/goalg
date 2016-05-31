package queue

// SliceQueue represents Queue based on slices
type SliceQueue struct {
	a []interface{}
}

// SliceQueueIterator data structure
type SliceQueueIterator struct {
	current int
	a       []interface{}
}

// NewSliceQueue creates an empty Queue and returns pointer to it
func NewSliceQueue() Queue { return &SliceQueue{a: make([]interface{}, 0)} }

// Enqueue adds item v to the end of the Queue
func (q *SliceQueue) Enqueue(v interface{}) {
	q.a = append(q.a, v)
}

// Dequeue returns the first item from the queue
func (q *SliceQueue) Dequeue() interface{} {
	if len(q.a) == 0 {
		return nil
	}
	retValue := q.a[0]
	q.a = q.a[1:]
	return retValue
}

// IsEmpty returns true if the Queue is empty, otherwise false
func (q *SliceQueue) IsEmpty() bool {
	return len(q.a) == 0
}

// Size returns number of items in the Queue
func (q *SliceQueue) Size() int {
	return len(q.a)
}

// GetIterator returns a Queue iterator
func (q *SliceQueue) GetIterator() Iterable {
	sliceCopy := make([]interface{}, len(q.a))
	copy(sliceCopy, q.a)
	return &SliceQueueIterator{
		current: 0,
		a:       sliceCopy,
	}
}

// HasNext moves iterator to the next element in collection
// if it exists and returns true, othewise it just returns false
func (it *SliceQueueIterator) HasNext() bool {
	return it.current < len(it.a)
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
