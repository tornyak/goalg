package stack

import (
	"strings"
	"fmt"
)

// Element is an element of the Stack
type element struct {
	next  *element
	value interface{}
}

// linkedStack represents linked list stack structure
type linkedStack struct {
	first *element
	size  int
}

// List creates an empty Stack and returns pointer to it
func Linked() Stack { return new(linkedStack) }

// Push adds items v to the top of the Stack
func (s *linkedStack) Push(v ...interface{}) Stack {
	for _, it := range v {
		e := new(element)
		e.value = it
		e.next = s.first
		s.first = e
		s.size++
	}
	return s
}

// Pop the item from the top of the stack. Return nil
// if stack is empty
func (s *linkedStack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	top := s.first.value
	s.first = s.first.next
	s.size--
	return top
}

// IsEmpty returns true if the Stack is empty, othewise false
func (s *linkedStack) IsEmpty() bool {
	return s.size == 0
}

// Size returns number of items in the Stack
func (s *linkedStack) Size() int {
	return s.size
}

// String returns formatted string representing linkedStack and its elements
func (b *linkedStack) String() string  {
	ret := "["
	for e := b.first; e != nil; e = e.next {
		ret+= fmt.Sprintf("%v ", e.value)
	}
	return strings.TrimSpace(ret) + "]"
}
