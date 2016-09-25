package stack

import "fmt"


// Stack is a slice based Stack implementation
// Implementation uses slice's inherent resizing support
type stack struct {
	a []interface{}
}

// Slice creates and returns a new empty slice backed stack
func Slice() Stack {
	return &stack{
		a: make([]interface{}, 0),
	}
}

// IsEmpty returns true if the Stack is empty, otherwise false
func (s *stack) IsEmpty() bool {
	return len(s.a) == 0
}

// Push adds item e to the top of the Stack
func (s *stack) Push(e ...interface{}) Stack {
	s.a = append(s.a, e...)
	return s
}

// Pop the item from the top of the stack.
// Returns (nil, false) if stack is empty
// otherwise (*topElement, true)
func (s *stack) Pop() interface{} {
	if len(s.a) == 0 {
		return nil
	}
	e := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return e
}

// Size returns number of items in the Stack
func (s *stack) Size() int {
	return len(s.a)
}

// String returns formatted string representing stack and its elements
func (s *stack) String() string  {
	return fmt.Sprintf("%v", s.a)
}