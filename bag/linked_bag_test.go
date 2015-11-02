package bag

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateEmpty(t *testing.T) {
	b := NewLinkedBag()
	assert.True(t, b.IsEmpty(), "Bag not empty")
	assert.Equal(t, b.Size(), 0, "Wrong Bag size")
}

func TestAddOne(t *testing.T) {
	b := NewLinkedBag()
	b.Add("Hello")
	assert.Equal(t, 1, b.Size(), "Wrong Bag size")
}

func TestAddMultiple(t *testing.T) {
	b := NewLinkedBag()
	b.Add("NewBag York")
	b.Add("Stockholm")
	b.Add("London")
	b.Add("Paris")
	assert.Equal(t, 4, b.Size(), "Wrong Bag size")
	assert.False(t, b.IsEmpty(), "Bag empty")
}

func TestAddDifferentTypes(t *testing.T) {
	b := NewLinkedBag()
	b.Add("NewBag York")
	b.Add(1.23)
	b.Add(time.Now())
	b.Add([]int{1, 2, 3})
	assert.Equal(t, 4, b.Size(), "Wrong Bag size")
	assert.False(t, b.IsEmpty(), "Bag empty")
}

func TestIterateEmpty(t *testing.T) {
	b := NewLinkedBag()
	it := b.GetIterator()
	assert.False(t, it.HasNext(), "Iterator should not have next element")
}

func TestIterateOne(t *testing.T) {
	b := NewLinkedBag()
	b.Add(5)
	it := b.GetIterator()
	for it.HasNext() {
		v := it.Next().(int)
		assert.Equal(t, 5, v, "Iterator returned wrong value")
	}
}

func TestIterateMultiple(t *testing.T) {
	b := NewLinkedBag()
	cities := map[string]int{
		"NewBag York": 1,
		"Stockholm":   1,
		"London":      1,
		"Paris":       1}

	for city := range cities {
		b.Add(city)
	}

	it := b.GetIterator()
	for it.HasNext() {
		city := it.Next().(string)
		cities[city]--
	}

	// Test that all cities in map were found
	for city, count := range cities {
		assert.Equal(t, 0, count, "Wrong count for city", city)
	}
}

func TestLinkedBagIterateNextNil(t *testing.T) {
	b := NewLinkedBag()
	b.Add(5)
	it := b.GetIterator()
	for it.HasNext() {
		v := it.Next().(int)
		assert.Equal(t, 5, v, "Iterator returned wrong value")
	}
	assert.Nil(t, it.Next(), "Next should have returned nil")
}

func TestLinkedBagIterateRemove(t *testing.T) {
	b := NewLinkedBag()
	b.Add(5)
	it := b.GetIterator()
	for it.HasNext() {
		it.Next()
		it.Remove()
	}
	it = b.GetIterator()
	for it.HasNext() {
		v := it.Next().(int)
		assert.Equal(t, 5, v, "Iterator returned wrong value")
	}
}
