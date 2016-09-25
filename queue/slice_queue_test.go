package queue_test

import (
	"testing"
	"github.com/tornyak/goalg/queue"
	"fmt"
)


func TestSliceQueueIsEmpty(t *testing.T) {
	var isEmptyTests = []struct {
		queue      queue.Queue
		expected bool // expected result
	}{
		{queue.Slice(), true},
		{queue.Slice().Enqueue("Hello"), false},
		{queue.Slice().Enqueue("New York", "Stockholm", "London", "Paris"), false},
		{queue.Slice().Enqueue("New York", 1.23, 1234567890, []int{1, 2, 3}), false},
	}

	for _, test := range isEmptyTests {
		if test.queue.IsEmpty() != test.expected {
			t.Errorf("IsEmpty() for Queue: %v, expected: %v, was: %v", test.queue, test.expected, test.queue.IsEmpty())
		}
	}
}

func TestSliceQueueSize(t *testing.T) {
	var sizeTests = []struct {
		queue      queue.Queue
		expected int
	}{
		{queue.Slice(), 0},
		{queue.Slice().Enqueue("Hello"), 1},
		{queue.Slice().Enqueue("New York", "Stockholm", "London", "Paris"), 4},
		{queue.Slice().Enqueue("New York", 1.23, 1234567890, []int{1, 2, 3}), 4},
	}

	for _, test := range sizeTests {
		if test.queue.Size() != test.expected {
			t.Errorf("Size() for Queue: %v, expected: %v, was: %v", test.queue, test.expected, test.queue.Size())
		}
	}
}

func TestSliceQueueDequeue(t *testing.T) {
	var tests = []struct {
		queue      queue.Queue
		expected interface{}
	}{
		{queue.Slice(), nil},
		{queue.Slice().Enqueue("Hello"), "Hello"},
		{queue.Slice().Enqueue("New York", "Stockholm", "London", "Paris"), "New York"},
		{queue.Slice().Enqueue(1234567890, "New York", 1.23, []int{1, 2, 3}), 1234567890},
	}

	for _, test := range tests {
		deq := test.queue.Dequeue()
		if deq != test.expected{
			t.Errorf("Dequeue() for Queue: %v, expected: %v was: %v",
				test.queue, test.expected, deq)
		}
	}
}

func TestSliceQueueDequeueMultiple(t *testing.T) {
	q := queue.Slice()
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

	deq := q.Dequeue()
	if deq != newCity {
		t.Errorf("Dequeue() for Queue: %v, expected: %v was: %v",
			q, newCity, deq)
	}

	if q.IsEmpty() != true {
		t.Errorf("IsEmpty() for Queue: %v, expected: %v, was: %v", q, true, q.IsEmpty())
	}

	if q.Size() != 0 {
		t.Errorf("Size() for Queue: %v, expected: %v, was: %v", q, 0, q.Size())
	}
}

func TestSliceQueueString(t *testing.T) {
	var stringTests = []struct {
		queue      queue.Queue
		expected string
	}{
		{queue.Slice(), "[]"},
		{queue.Slice().Enqueue("Hello"), "[Hello]"},
		{queue.Slice().Enqueue("New York", "Stockholm", "London", "Paris"), "[New York Stockholm London Paris]"},
		{queue.Slice().Enqueue("New York", 1.23, 1234567890, []int{1, 2, 3}), "[New York 1.23 1234567890 [1 2 3]]"},
	}

	for _, test := range stringTests {
		if fmt.Sprint(test.queue) != test.expected {
			t.Errorf("String() for Queue: %v, expected: %v", test.queue, test.expected)
		}
	}
}
