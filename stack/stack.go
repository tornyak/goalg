// Package stack implements a LIFO (Last In First Out) Stack
// collection.
package stack

// Element is an element of the Stack
type element struct {
	next  *element
	value interface{}
}

// A Stack structure
type Stack struct {
	first *element
	size  int
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

// NewStack creates an empty Stack and returns pointer to it
func NewStack() *Stack { return new(Stack) }

// Push adds item v to the top of the Stack
func (s *Stack) Push(v interface{}) {
	e := new(element)
	e.value = v
	e.next = s.first
	s.first = e
	s.size++
}

// Pop the item from the top of the stack. Return nil
// if stack is empty
func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	top := s.first.value
	s.first = s.first.next
	s.size--
	return top
}

// IsEmpty returns true if the Stack is empty, othewise false
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Size returns number of items in the Stack
func (s *Stack) Size() int {
	return s.size
}

// GetIterator returns a Stack iterator
func (s *Stack) GetIterator() *Iterator {
	it := new(Iterator)
	it.current = new(element)
	it.current.next = s.first
	return it
}
