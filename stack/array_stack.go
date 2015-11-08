package stack

// ArrayStack is stack implementation where data is stored in slice
// Since slices inherently support resizing implementation is simple
type ArrayStack struct {
	size int
	a    []interface{}
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
func NewArrayStack() *ArrayStack {
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
