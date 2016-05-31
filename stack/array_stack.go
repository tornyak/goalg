package stack

// ArrayStack is stack implementation where data is stored in slice
// Since slices inherently support resizing implementation is simple
type ArrayStack struct {
	a    []interface{}
}

// ArrayStackIterator is iterator for ArrayStack
type ArrayStackIterator struct {
	current int
	a       []interface{}
}

// NewArrayStack creates and returns new
// ArrayStack data structure of capacity 1
func NewArrayStack() Stack {
	return &ArrayStack{
		a: make([]interface{}, 0),
	}
}

// IsEmpty returns true if the Stack is empty, othewise false
func (s *ArrayStack) IsEmpty() bool {
	return len(s.a) == 0
}

// Push adds item e to the top of the Stack
func (s *ArrayStack) Push(e interface{}) {
	s.a = append(s.a, e)
}

// Pop the item from the top of the stack. Return nil
// if stack is empty
func (s *ArrayStack) Pop() interface{} {
	if len(s.a) == 0 {
		return nil
	}
	e := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return e
}

// Size returns number of items in the Stack
func (s *ArrayStack) Size() int {
	return len(s.a)
}

// GetIterator returns a Stack iterator
func (s *ArrayStack) GetIterator() Iterable {
	return &ArrayStackIterator{
		current: s.Size(),
		a:       s.a,
	}
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
