package queue

import "fmt"

// sliceQueue represents Queue based on slices
type sliceQueue struct {
	a []interface{}
}

// Slice creates an empty Queue and returns pointer to it
func Slice() Queue { return &sliceQueue{a: make([]interface{}, 0)} }

// Enqueue adds item v to the end of the Queue
func (q *sliceQueue) Enqueue(v ...interface{}) Queue {
	q.a = append(q.a, v...)
	return q
}

// Dequeue returns the first item from the queue
func (q *sliceQueue) Dequeue() interface{} {
	if len(q.a) == 0 {
		return nil
	}
	retValue := q.a[0]
	q.a = q.a[1:]
	return retValue
}

// IsEmpty returns true if the Queue is empty, otherwise false
func (q *sliceQueue) IsEmpty() bool {
	return len(q.a) == 0
}

// Size returns number of items in the Queue
func (q *sliceQueue) Size() int {
	return len(q.a)
}

// String returns formatted string representing sliceQueue and its elements
func (s *sliceQueue) String() string  {
	return fmt.Sprintf("%v", s.a)
}