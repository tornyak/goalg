package slice

import (
	"fmt"
	"math"
)

type IntStack interface {
	Push(...int) IntStack
	Pop() int
	IsEmpty() bool
	Size() int
	ToSlice() []int
}

// intStack is an integer slice based Stack implementation
type intStack struct {
	a []int
}

// IntSlice creates and returns a new empty slice backed stack
func Stack() IntStack {
	return &intStack{}
}

// IsEmpty returns true if the Stack is empty, otherwise false
func (s *intStack) IsEmpty() bool {
	return len(s.a) == 0
}

// Push adds item e to the top of the Stack
func (s *intStack) Push(e ...int) IntStack {
	s.a = append(s.a, e...)
	return s
}

// Pop the item from the top of the stack.
// Returns math.MinInt64 if stack is empty
// otherwise element from the top of the stack
func (s *intStack) Pop() int {

	if len(s.a) == 0 {
		return math.MinInt64
	}
	e := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return e
}

// Size returns number of items in the Stack
func (s *intStack) Size() int {
	return len(s.a)
}

// ToSlice returns a copy of the underlying IntStack slice
func (s *intStack) ToSlice() []int  {
	var retCopy []int
	return append(retCopy, s.a...)
}

// String returns formatted string representing stack and its elements
func (s *intStack) String() string  {
	return fmt.Sprintf("%v", s.a)
}