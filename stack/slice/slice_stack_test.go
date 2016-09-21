package stack_test

import (
	"testing"
	"github.com/tornyak/goalg/stack
	"github.com/tornyak/goalg/bag/slice"

	"sort"
)

func TestIsEmpty(t *testing.T) {
	var isEmptyTests = []struct {
		stack      stack.Stack
		expected bool // expected result
	}{
		{slice.NewStack(), true},
		{stack.NewStack().Push("Hello"), false},
		{stack.NewStack().Push("New York", "Stockholm", "London", "Paris"), false},
		{stack.NewStack().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), false},
	}

	for _, test := range isEmptyTests {
		if test.stack.IsEmpty() != test.expected {
			t.Errorf("IsEmpty() for Bag: %v, expected: %v, was: %v", test.stack, test.expected, test.stack.IsEmpty())
		}
	}

	sort.IntsAreSorted()
}

//func TestArrayStackPopOne(t *testing.T) {
//	s := NewArrayStack()
//	expected := "Stockholm"
//	s.Push(expected)
//	assert.False(t, s.IsEmpty(), "Stack empty")
//	assert.Equal(t, s.Size(), 1, "Wrong Stack size")
//
//	city := s.Pop().(string)
//	assert.Equal(t, expected, city, "Stack.Pop returned wrong element")
//}
//
//func TestArrayStackPopMultiple(t *testing.T) {
//	s := NewArrayStack()
//	cities := []string{"New York", "Stockholm", "Paris", "London"}
//	for _, city := range cities {
//		s.Push(city)
//	}
//
//	assert.Equal(t, s.Size(), 4, "Wrong Stack size")
//	for i := len(cities) - 1; !s.IsEmpty(); i-- {
//		city := s.Pop().(string)
//		assert.Equal(t, cities[i], city, "Stack.Pop returned wrong element")
//	}
//
//	newCity := "Tokyo"
//	s.Push(newCity)
//	assert.Equal(t, 1, s.Size(), "Wrong Stack size")
//	assert.Equal(t, newCity, s.Pop().(string), "Stack.Pop returned wrong element")
//	assert.Equal(t, 0, s.Size(), "Wrong Stack size")
//	assert.True(t, s.IsEmpty(), "Stack not empty")
//}
//
//func TestArrayStackPopEmpty(t *testing.T) {
//	s := NewArrayStack()
//	e := s.Pop()
//	assert.Nil(t, e)
//}
//
//func TestArrayStackResize(t *testing.T) {
//	s := NewArrayStack().(*ArrayStack)
//	assert.Equal(t, 0, cap(s.a), "Stack created with wrong capacity")
//	cities := []string{"New York", "Stockholm", "Paris", "London", "Tokyo"}
//	for _, city := range cities {
//		s.Push(city)
//	}
//	assert.Equal(t, 8, cap(s.a), "Stack + resized with wrong capacity")
//	for s.Size() > 1 {
//		s.Pop()
//	}
//	assert.Equal(t, 8, cap(s.a), "Stack - resized with wrong capacity")
//}
//
