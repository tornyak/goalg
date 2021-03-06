package stack_test

import (
	"testing"
	"github.com/tornyak/goalg/stack"
	"fmt"
)


func TestLinkedStackIsEmpty(t *testing.T) {
	var isEmptyTests = []struct {
		stack      stack.Stack
		expected bool // expected result
	}{
		{stack.Linked(), true},
		{stack.Linked().Push("Hello"), false},
		{stack.Linked().Push("New York", "Stockholm", "London", "Paris"), false},
		{stack.Linked().Push("New York", 1.23, 1234567890, []int{1, 2, 3}), false},
	}

	for _, test := range isEmptyTests {
		if test.stack.IsEmpty() != test.expected {
			t.Errorf("IsEmpty() for Stack: %v, expected: %v, was: %v", test.stack, test.expected, test.stack.IsEmpty())
		}
	}
}

func TestLinkedStackSize(t *testing.T) {
	var sizeTests = []struct {
		stack      stack.Stack
		expected int
	}{
		{stack.Linked(), 0},
		{stack.Linked().Push("Hello"), 1},
		{stack.Linked().Push("New York", "Stockholm", "London", "Paris"), 4},
		{stack.Linked().Push("New York", 1.23, 1234567890, []int{1, 2, 3}), 4},
	}

	for _, test := range sizeTests {
		if test.stack.Size() != test.expected {
			t.Errorf("Size() for Stack: %v, expected: %v, was: %v", test.stack, test.expected, test.stack.Size())
		}
	}
}

func TestLinkedStackPop(t *testing.T) {
	var sizeTests = []struct {
		stack      stack.Stack
		expected interface{}
	}{
		{stack.Linked(), nil},
		{stack.Linked().Push("Hello"), "Hello"},
		{stack.Linked().Push("New York", "Stockholm", "London", "Paris"), "Paris"},
		{stack.Linked().Push("New York", 1.23, []int{1, 2, 3}, 1234567890), 1234567890},
	}

	for _, test := range sizeTests {
		pop := test.stack.Pop()
		if pop != test.expected{
			t.Errorf("Pop() for Stack: %v, expected: %v was: %v",
				test.stack, test.expected, pop)
		}
	}
}

func TestLinkedStackPopMultiple(t *testing.T) {
	s := stack.Linked()
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

func TestLinkedStackString(t *testing.T) {
	var stringTest = []struct {
		stack      stack.Stack
		expected string
	}{
		{stack.Linked(), "[]"},
		{stack.Linked().Push("Hello"), "[Hello]"},
		{stack.Linked().Push("New York", "Stockholm", "London", "Paris"), "[Paris London Stockholm New York]"},
		{stack.Linked().Push("New York", 1.23, 1234567890, []int{1, 2, 3}), "[[1 2 3] 1234567890 1.23 New York]"},
	}

	for _, test := range stringTest {
		if fmt.Sprint(test.stack) != test.expected {
			t.Errorf("String() for Stack: %v, expected: %v", test.stack, test.expected)
		}
	}
}