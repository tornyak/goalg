package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedStackCreateEmpty(t *testing.T) {
	s := NewLinkedStack()
	assert.True(t, s.IsEmpty(), "Stack not empty")
	assert.Equal(t, s.Size(), 0, "Wrong Stack size")
}

func TestLinkedStackPopOne(t *testing.T) {
	s := NewLinkedStack()
	expected := "Stockholm"
	s.Push(expected)
	assert.False(t, s.IsEmpty(), "Stack empty")
	assert.Equal(t, s.Size(), 1, "Wrong Stack size")

	city := s.Pop().(string)
	assert.Equal(t, expected, city, "Stack.Pop returned wrong element")
}

func TestLinkedStackPopMultiple(t *testing.T) {
	s := NewLinkedStack()
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

func TestLinkedStackPopEmpty(t *testing.T) {
	s := NewLinkedStack()
	e := s.Pop()
	assert.Nil(t, e)
}

func TestLinkedStackIterateEmpty(t *testing.T) {
	s := NewLinkedStack()
	it := s.GetIterator()
	assert.Nil(t, it.Next(), "Iterator Next() returned wrong value")
}

func TestLinkedStackIterateOne(t *testing.T) {
	s := NewLinkedStack()
	e := "New York"
	s.Push(e)
	it := s.GetIterator()
	assert.Equal(t, e, it.Next(), "Iterator Next() returned wrong value")
}

func TestLinkedStackIterateMultiple(t *testing.T) {
	s := NewLinkedStack()
	cities := []string{"New York", "Stockholm", "Paris", "London"}
	for _, city := range cities {
		s.Push(city)
	}
	it := s.GetIterator()
	i := len(cities) - 1
	for it.HasNext() {
		assert.Equal(t, cities[i], it.Next(), "Iterator Next() returned wrong value")
		i--
	}
}

func TestLinkedStackIterateNextNil(t *testing.T) {
	s := NewLinkedStack()
	cities := []string{"New York", "Stockholm", "Paris"}
	for _, city := range cities {
		s.Push(city)
	}
	it := s.GetIterator()
	for it.HasNext() {
		it.Next()
	}
	assert.Nil(t, it.Next(), "Iterator Next() returned wrong value")
}
