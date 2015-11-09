package stack

// Element is an element of the Stack
type element struct {
	next  *element
	value interface{}
}

// A LinkedStack structure
type LinkedStack struct {
	first *element
	size  int
}

// LinkedStackIterator data structure
type LinkedStackIterator struct {
	current *element
}

// NewLinkedStack creates an empty Stack and returns pointer to it
func NewLinkedStack() Stack { return new(LinkedStack) }

// Push adds item v to the top of the Stack
func (s *LinkedStack) Push(v interface{}) {
	e := new(element)
	e.value = v
	e.next = s.first
	s.first = e
	s.size++
}

// Pop the item from the top of the stack. Return nil
// if stack is empty
func (s *LinkedStack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	top := s.first.value
	s.first = s.first.next
	s.size--
	return top
}

// IsEmpty returns true if the Stack is empty, othewise false
func (s *LinkedStack) IsEmpty() bool {
	return s.size == 0
}

// Size returns number of items in the Stack
func (s *LinkedStack) Size() int {
	return s.size
}

// GetIterator returns a Stack iterator
func (s *LinkedStack) GetIterator() Iterable {
	it := new(LinkedStackIterator)
	it.current = s.first
	return it
}

// HasNext moves iterator to the next element in collection
// if it exists and returns true, othewise it just returns false
func (it *LinkedStackIterator) HasNext() bool {
	return it.current != nil
}

// Next returns the value of the next element. If there is no
// next element nil is returned
func (it *LinkedStackIterator) Next() interface{} {
	if it.current != nil {
		curr := it.current
		it.current = it.current.next
		return curr.value
	}
	return nil
}
