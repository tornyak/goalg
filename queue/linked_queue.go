package queue

// Element is an element of the Queue
type element struct {
	next  *element
	value interface{}
}

// A LinkedQueue structure
type LinkedQueue struct {
	first *element
	last  *element
	size  int
}

// LinkedQueueIterator data structure
type LinkedQueueIterator struct {
	current *element
}

// NewLinkedQueue creates an empty Queue and returns pointer to it
func NewLinkedQueue() *LinkedQueue { return new(LinkedQueue) }

// Enqueue adds item v to the end of the Queue
func (q *LinkedQueue) Enqueue(v interface{}) {
	oldLast := q.last
	e := new(element)
	e.value = v
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
func (q *LinkedQueue) Dequeue() interface{} {
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
func (q *LinkedQueue) IsEmpty() bool {
	return q.size == 0
}

// Size returns number of items in the Queue
func (q *LinkedQueue) Size() int {
	return q.size
}

// GetIterator returns a Queue iterator
func (q *LinkedQueue) GetIterator() Iterable {
	it := new(LinkedQueueIterator)
	it.current = q.first
	return it
}

// HasNext moves iterator to the next element in collection
// if it exists and returns true, othewise it just returns false
func (it *LinkedQueueIterator) HasNext() bool {
	return it.current != nil
}

// Next returns the value of the current element or nil
// if current element is nil
func (it *LinkedQueueIterator) Next() interface{} {
	if it.current != nil {
		curr := it.current
		it.current = it.current.next
		return curr.value
	}
	return nil
}
