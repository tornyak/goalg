package stack

// ArrayStack is stack implementation where data is stored in slice
// Since slices inherently support resizing implementation is simple
type ArrayStack struct {
	size int
	a    []interface{}
}

// ArrayStackIterator is iterator for ArrayStack
type ArrayStackIterator struct {
	current int
	a       []interface{}
}

// resize creates a new slice of size max and
// copies all data from a into it and replaces a with the new slice
func (s *ArrayStack) resize(max int) {
	newA := make([]interface{}, max)
	copy(newA, s.a)
	s.a = newA
}

// NewArrayStack creates and returns new
// ArrayStack data structure of capacity 1
func NewArrayStack() Stack {
	s := &ArrayStack{
		a: make([]interface{}, 1),
	}
	return s
}

// IsEmpty returns true if the Stack is empty, othewise false
func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}

// Push adds item e to the top of the Stack
func (s *ArrayStack) Push(e interface{}) {
	if s.size == cap(s.a) {
		s.resize(s.size * 2)
	}
	s.a[s.size] = e
	s.size++
}

// Pop the item from the top of the stack. Return nil
// if stack is empty
func (s *ArrayStack) Pop() interface{} {
	if s.Size() == 0 {
		return nil
	}
	s.size--
	e := s.a[s.size]
	s.a[s.size] = nil
	if s.size > 0 && s.size == cap(s.a)/4 {
		s.resize(cap(s.a) / 2)
	}
	return e
}

// Size returns number of items in the Stack
func (s *ArrayStack) Size() int {
	return s.size
}

// GetIterator returns a Stack iterator
func (s *ArrayStack) GetIterator() Iterable {
	it := &ArrayStackIterator{
		current: s.Size(),
		a:       s.a,
	}
	return it
}

// HasNext moves iterator to the next element in collection
// if it exists and returns true, othewise it just returns false
func (it *ArrayStackIterator) HasNext() bool {
	return it.current > 0
}

// Next returns the value of the next element. If there is no
// next element nil is returned
func (it *ArrayStackIterator) Next() interface{} {
	if it.current > 0 {
		it.current--
		return it.a[it.current]
	}
	return nil
}
