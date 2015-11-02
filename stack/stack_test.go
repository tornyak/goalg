package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	testCreateEmpty(t)
	testStackPopOne(t)
	testStackPopMultiple(t)
	//testAddDifferentTypes(t)
	//testIterateEmpty(t)
	//testIterateOne(t)
	//testIterateMultiple(t)
}

func testCreateEmpty(t *testing.T) {
	s := NewStack()
	assert.True(t, s.IsEmpty(), "Stack not empty")
	assert.Equal(t, s.Size(), 0, "Wrong Stack size")
}

func testStackPopOne(t *testing.T) {
	s := NewStack()
	expected := "Stockholm"
	s.Push(expected)
	assert.False(t, s.IsEmpty(), "Stack empty")
	assert.Equal(t, s.Size(), 1, "Wrong Stack size")

	city := s.Pop().(string)
	assert.Equal(t, expected, city, "Stack.Pop returned wrong element")
}

func testStackPopMultiple(t *testing.T) {
	s := NewStack()
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
