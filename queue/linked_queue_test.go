package queue_test

import (
	"testing"
	"fmt"
	"github.com/tornyak/goalg/queue"
)

func TestLinkedQueueIsEmpty(t *testing.T) {
	var tests = []struct {
		queue      queue.Queue
		expected bool // expected result
	}{
		{queue.Linked(), true},
		{queue.Linked().Enqueue("Hello"), false},
		{queue.Linked().Enqueue("New York", "Stockholm", "London", "Paris"), false},
		{queue.Linked().Enqueue("New York", 1.23, 1234567890, []int{1, 2, 3}), false},
	}

	for _, test := range tests {
		if test.queue.IsEmpty() != test.expected {
			t.Errorf("IsEmpty() for Queue: %v, expected: %v, was: %v", test.queue, test.expected, test.queue.IsEmpty())
		}
	}
}

func TestLinkedQueueSize(t *testing.T) {
	var tests = []struct {
		queue      queue.Queue
		expected int
	}{
		{queue.Linked(), 0},
		{queue.Linked().Enqueue("Hello"), 1},
		{queue.Linked().Enqueue("New York", "Stockholm", "London", "Paris"), 4},
		{queue.Linked().Enqueue("New York", 1.23, 1234567890, []int{1, 2, 3}), 4},
	}

	for _, test := range tests {
		if test.queue.Size() != test.expected {
			t.Errorf("Size() for Queue: %v, expected: %v, was: %v", test.queue, test.expected, test.queue.Size())
		}
	}
}

func TestLinkedQueueDequeue(t *testing.T) {
	var tests = []struct {
		queue      queue.Queue
		expected interface{}
	}{
		{queue.Linked(), nil},
		{queue.Linked().Enqueue("Hello"), "Hello"},
		{queue.Linked().Enqueue("New York", "Stockholm", "London", "Paris"), "New York"},
		{queue.Linked().Enqueue(1234567890, "New York", 1.23, []int{1, 2, 3}), 1234567890},
	}

	for _, test := range tests {
		deq := test.queue.Dequeue()
		if deq != test.expected{
			t.Errorf("Dequeue() for Queue: %v, expected: %v was: %v",
				test.queue, test.expected, deq)
		}
	}
}

func TestLinkedQueueDequeueMultiple(t *testing.T) {
	q := queue.Linked()
	cities := []string{"New York", "Stockholm", "Paris", "London"}
	q.Enqueue("New York", "Stockholm", "Paris", "London")

	for i := 0; !q.IsEmpty(); i++ {
		deq := q.Dequeue()
		if deq != cities[i]{
			t.Errorf("Dequeue() for Queue: %v, expected: %v was: %v",
				q, cities[i], deq)
		}
	}

	if q.IsEmpty() != true {
		t.Errorf("IsEmpty() for Queue: %v, expected: %v, was: %v", q, true, q.IsEmpty())
	}

	newCity := "Tokyo"
	q.Enqueue(newCity)

	if q.IsEmpty() != false {
		t.Errorf("IsEmpty() for Queue: %v, expected: %v, was: %v", q, false, q.IsEmpty())
	}

	if q.Size() != 1 {
		t.Errorf("Size() for Queue: %v, expected: %v, was: %v", q, 1, q.Size())
	}

	pop := q.Dequeue()
	if pop != newCity {
		t.Errorf("Dequeue() for Queue: %v, expected: %v was: %v",
			q, newCity, pop)
	}

	if q.IsEmpty() != true {
		t.Errorf("IsEmpty() for Queue: %v, expected: %v, was: %v", q, true, q.IsEmpty())
	}

	if q.Size() != 0 {
		t.Errorf("Size() for Queue: %v, expected: %v, was: %v", q, 0, q.Size())
	}
}

func TestLinkedQueueString(t *testing.T) {
	var stringTest = []struct {
		queue      queue.Queue
		expected string
	}{
		{queue.Linked(), "[]"},
		{queue.Linked().Enqueue("Hello"), "[Hello]"},
		{queue.Linked().Enqueue("New York", "Stockholm", "London", "Paris"), "[New York Stockholm London Paris]"},
		{queue.Linked().Enqueue("New York", 1.23, 1234567890, []int{1, 2, 3}), "[New York 1.23 1234567890 [1 2 3]]"},
	}

	for _, test := range stringTest {
		if fmt.Sprint(test.queue) != test.expected {
			t.Errorf("String() for Queue: %v, expected: %v", test.queue, test.expected)
		}
	}
}

