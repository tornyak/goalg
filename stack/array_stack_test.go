package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayStackCreateEmpty(t *testing.T) {
	s := NewArrayStack()
	assert.True(t, s.IsEmpty(), "Stack not empty")
	assert.Equal(t, s.Size(), 0, "Wrong Stack size")
}

func TestArrayStackPopOne(t *testing.T) {
	s := NewArrayStack()
	expected := "Stockholm"
	s.Push(expected)
	assert.False(t, s.IsEmpty(), "Stack empty")
	assert.Equal(t, s.Size(), 1, "Wrong Stack size")

	city := s.Pop().(string)
	assert.Equal(t, expected, city, "Stack.Pop returned wrong element")
}

func TestArrayStackPopMultiple(t *testing.T) {
	s := NewArrayStack()
	cities := []string{"New York", "Stockholm", "Paris", "London"}
	for _, city := range cities {
		s.Push(city)
	}

	assert.Equal(t, s.Size(), 4, "Wrong Stack size")
	for i := len(cities) - 1; !s.IsEmpty(); i-- {
		city := s.Pop().(string)
		assert.Equal(t, cities[i], city, "Stack.Pop returned wrong element")
	}

	newCity := "Tokyo"
	s.Push(newCity)
	assert.Equal(t, 1, s.Size(), "Wrong Stack size")
	assert.Equal(t, newCity, s.Pop().(string), "Stack.Pop returned wrong element")
	assert.Equal(t, 0, s.Size(), "Wrong Stack size")
	assert.True(t, s.IsEmpty(), "Stack not empty")
}

func TestArrayStackPopEmpty(t *testing.T) {
	s := NewArrayStack()
	e := s.Pop()
	assert.Nil(t, e)
}

func TestArrayStackResize(t *testing.T) {
	s := NewArrayStack()
	assert.Equal(t, 1, cap(s.a), "Stack created with wrong capacity")
	cities := []string{"New York", "Stockholm", "Paris", "London", "Tokyo"}
	for _, city := range cities {
		s.Push(city)
	}
	assert.Equal(t, 8, cap(s.a), "Stack + resized with wrong capacity")
	for s.Size() > 1 {
		s.Pop()
	}
	assert.Equal(t, 2, cap(s.a), "Stack - resized with wrong capacity")
}

// func TestArrayStackIterateEmpty(t *testing.T) {
// 	s := NewArrayStack()
// 	it := s.GetIterator()
// 	assert.Nil(t, it.Next(), "Iterator Next() returned wrong value")
// }
//
// func TestArrayStackIterateOne(t *testing.T) {
// 	s := NewArrayStack()
// 	e := "New York"
// 	s.Push(e)
// 	it := s.GetIterator()
// 	assert.Equal(t, e, it.Next(), "Iterator Next() returned wrong value")
// }
//
// func TestArrayStackIterateMultiple(t *testing.T) {
// 	s := NewArrayStack()
// 	cities := []string{"New York", "Stockholm", "Paris", "London"}
// 	for _, city := range cities {
// 		s.Push(city)
// 	}
// 	it := s.GetIterator()
// 	i := len(cities) - 1
// 	for it.HasNext() {
// 		assert.Equal(t, cities[i], it.Next(), "Iterator Next() returned wrong value")
// 		i--
// 	}
// }
//
// func TestArrayStackIterateNextNil(t *testing.T) {
// 	s := NewArrayStack()
// 	cities := []string{"New York", "Stockholm", "Paris"}
// 	for _, city := range cities {
// 		s.Push(city)
// 	}
// 	it := s.GetIterator()
// 	for it.HasNext() {
// 		it.Next()
// 	}
// 	assert.Nil(t, it.Next(), "Iterator Next() returned wrong value")
// }
