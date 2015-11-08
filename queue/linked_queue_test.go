package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedQueueCreateEmpty(t *testing.T) {
	q := NewLinkedQueue()
	assert.True(t, q.IsEmpty(), "Stack not empty")
	assert.Equal(t, q.Size(), 0, "Wrong queue size")
}

func TestLinkedQueueEnqueueDequeueOne(t *testing.T) {
	q := NewLinkedQueue()
	expected := "Stockholm"
	q.Enqueue(expected)
	assert.False(t, q.IsEmpty(), "Queue empty")
	assert.Equal(t, q.Size(), 1, "Wrong Stack size")

	e := q.Dequeue().(string)
	assert.Equal(t, expected, e, "Dequeue returned wrong element")
}

func TestLinkedStackQueueDequeueMultiple(t *testing.T) {
	q := NewLinkedQueue()
	cities := []string{"New York", "Stockholm", "Paris", "London"}
	for _, city := range cities {
		q.Enqueue(city)
	}
	assert.Equal(t, len(cities), q.Size(), "Wrong Size")

	for i := 0; !q.IsEmpty(); i++ {
		city := q.Dequeue().(string)
		assert.Equal(t, cities[i], city, "Dequeue returned wrong element")
	}

	newCity := "Tokyo"
	q.Enqueue(newCity)
	assert.Equal(t, 1, q.Size(), "Wrong Size")

	city := q.Dequeue().(string)
	assert.Equal(t, newCity, city, "Dequeue returned wrong element")
}

func TestLinkedQueueDequeueEmpty(t *testing.T) {
	q := NewLinkedQueue()
	e := q.Dequeue()
	assert.Nil(t, e)
}

func TestLinkedStackIterateEmpty(t *testing.T) {
	q := NewLinkedQueue()
	it := q.GetIterator()
	assert.Nil(t, it.Next(), "Iterator Next() returned wrong value")
}

func TestLinkedStackIterateOne(t *testing.T) {
	q := NewLinkedQueue()
	e := "New York"
	q.Enqueue(e)
	it := q.GetIterator()
	assert.Equal(t, e, it.Next(), "Iterator Next() returned wrong value")
}

func TestLinkedStackIterateMultiple(t *testing.T) {
	q := NewLinkedQueue()
	cities := []string{"New York", "Stockholm", "Paris", "London"}
	for _, city := range cities {
		q.Enqueue(city)
	}
	it := q.GetIterator()
	i := 0
	for it.HasNext() {
		assert.Equal(t, cities[i], it.Next(), "Iterator Next() returned wrong value")
		i++
	}
}

func TestLinkedStackIterateNextNil(t *testing.T) {
	q := NewLinkedQueue()
	cities := []string{"New York", "Stockholm", "Paris"}
	for _, city := range cities {
		q.Enqueue(city)
	}
	it := q.GetIterator()
	for it.HasNext() {
		it.Next()
	}
	assert.Nil(t, it.Next(), "Iterator Next() returned wrong value")
}
