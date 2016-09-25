package stack_test

import (
	"testing"
	"github.com/tornyak/goalg/stack"
	"fmt"
)


func TestSliceStackIsEmpty(t *testing.T) {
	var isEmptyTests = []struct {
		stack      stack.Stack
		expected bool // expected result
	}{
		{stack.Slice(), true},
		{stack.Slice().Push("Hello"), false},
		{stack.Slice().Push("New York", "Stockholm", "London", "Paris"), false},
		{stack.Slice().Push("New York", 1.23, 1234567890, []int{1, 2, 3}), false},
	}

	for _, test := range isEmptyTests {
		if test.stack.IsEmpty() != test.expected {
			t.Errorf("IsEmpty() for Stack: %v, expected: %v, was: %v", test.stack, test.expected, test.stack.IsEmpty())
		}
	}
}

func TestSliceStackSize(t *testing.T) {
	var sizeTests = []struct {
		stack      stack.Stack
		expected int
	}{
		{stack.Slice(), 0},
		{stack.Slice().Push("Hello"), 1},
		{stack.Slice().Push("New York", "Stockholm", "London", "Paris"), 4},
		{stack.Slice().Push("New York", 1.23, 1234567890, []int{1, 2, 3}), 4},
	}

	for _, test := range sizeTests {
		if test.stack.Size() != test.expected {
			t.Errorf("Size() for Stack: %v, expected: %v, was: %v", test.stack, test.expected, test.stack.Size())
		}
	}
}

func TestSliceStackPop(t *testing.T) {
	var sizeTests = []struct {
		stack      stack.Stack
		expected interface{}
	}{
		{stack.Slice(), nil},
		{stack.Slice().Push("Hello"), "Hello"},
		{stack.Slice().Push("New York", "Stockholm", "London", "Paris"), "Paris"},
		{stack.Slice().Push("New York", 1.23, []int{1, 2, 3}, 1234567890), 1234567890},
	}

	for _, test := range sizeTests {
		pop := test.stack.Pop()
		if pop != test.expected{
			t.Errorf("Pop() for Stack: %v, expected: %v was: %v",
				test.stack, test.expected, pop)
		}
	}
}

func TestSliceStackPopMultiple(t *testing.T) {
	s := stack.Slice()
	cities := []string{"New York", "Stockholm", "Paris", "London"}
	s.Push("New York", "Stockholm", "Paris", "London")

	for i := len(cities) - 1; !s.IsEmpty(); i-- {
		pop := s.Pop()
		if pop != cities[i]{
			t.Errorf("Pop() for Stack: %v, expected: %v was: %v",
				s, cities[i], pop)
		}
	}

	if s.IsEmpty() != true {
		t.Errorf("IsEmpty() for Stack: %v, expected: %v, was: %v", s, true, s.IsEmpty())
	}

	newCity := "Tokyo"
	s.Push(newCity)

	if s.IsEmpty() != false {
		t.Errorf("IsEmpty() for Stack: %v, expected: %v, was: %v", s, false, s.IsEmpty())
	}

	if s.Size() != 1 {
		t.Errorf("Size() for Stack: %v, expected: %v, was: %v", s, 1, s.Size())
	}

	pop := s.Pop()
	if pop != newCity {
		t.Errorf("Pop() for Stack: %v, expected: %v was: %v",
			s, newCity, pop)
	}

	if s.IsEmpty() != true {
		t.Errorf("IsEmpty() for Stack: %v, expected: %v, was: %v", s, true, s.IsEmpty())
	}

	if s.Size() != 0 {
		t.Errorf("Size() for Stack: %v, expected: %v, was: %v", s, 0, s.Size())
	}
}

func TestSliceStackString(t *testing.T) {
	var stringTests = []struct {
		stack      stack.Stack
		expected string
	}{
		{stack.Slice(), "[]"},
		{stack.Slice().Push("Hello"), "[Hello]"},
		{stack.Slice().Push("New York", "Stockholm", "London", "Paris"), "[New York Stockholm London Paris]"},
		{stack.Slice().Push("New York", 1.23, 1234567890, []int{1, 2, 3}), "[New York 1.23 1234567890 [1 2 3]]"},
	}

	for _, test := range stringTests {
		if fmt.Sprint(test.stack) != test.expected {
			t.Errorf("String() for Stack: %v, expected: %v", test.stack, test.expected)
		}
	}
}
