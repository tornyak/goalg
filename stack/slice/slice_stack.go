package stack

import "github.com/tornyak/goalg/stack"

// Stack is a slice based Stack implementation
// Implementation uses slice's inherent resizing support
type Stack struct {
	a []interface{}
}

// NewArrayStack creates and returns new
// ArrayStack data structure of capacity 1
func NewStack() Stack {
	return &Stack{
		a: make([]interface{}, 0),
	}
}

// IsEmpty returns true if the Stack is empty, otherwise false
func (s *Stack) IsEmpty() bool {
	return len(s.a) == 0
}

// Push adds item e to the top of the Stack
func (s *Stack) Push(e ...interface{}) stack.Stack {
	s.a = append(s.a, e...)
	return s
}

// Pop the item from the top of the stack.
// Returns (nil, false) if stack is empty
// otherwise (*topElement, true)
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.a) == 0 {
		return nil, false
	}
	e := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return e, true
}

// Size returns number of items in the Stack
func (s *Stack) Size() int {
	return len(s.a)
}
