package queue

import (
	"testing"

	"github.com/tornyak/goalg/util"
)

func TestQueue(t *testing.T) {
	testCreateEmpty(t)
	testQueueDequeueOne(t)
	testQueueDequeueMultiple(t)
	//testAddDifferentTypes(t)
	//testIterateEmpty(t)
	//testIterateOne(t)
	//testIterateMultiple(t)
}

func testCreateEmpty(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Error(util.GetFuncName(), ":Queue not empty")
	}
	if q.Size() != 0 {
		t.Errorf("%q :size expected 0 received %d", util.GetFuncName(), q.Size())
	}
}

func testQueueDequeueOne(t *testing.T) {
	q := New()
	expected := "Stockholm"
	q.Enqueue(expected)
	if q.IsEmpty() {
		t.Error(util.GetFuncName(), ":Queue empty")
	}

	if q.Size() != 1 {
		t.Errorf("%q :size expected 1 received %d", util.GetFuncName(), q.Size())
	}

	city := q.Dequeue().(string)
	if city != expected {
		t.Errorf("%q :Dequeue expected %q received %q", util.GetFuncName(), expected, city)
	}

}

func testQueueDequeueMultiple(t *testing.T) {
	q := New()
	cities := []string{"New York", "Stockholm", "Paris", "London"}
	for _, city := range cities {
		q.Enqueue(city)
	}

	if q.Size() != len(cities) {
		t.Errorf("%q :size expected %d received %d", util.GetFuncName(), len(cities), q.Size())
	}

	for i := 0; !q.IsEmpty(); i++ {
		city := q.Dequeue().(string)
		if city != cities[i] {
			t.Errorf("%q :Dequeue expected %q received %q", util.GetFuncName(), cities[i], city)
		}
	}
	newCity := "Tokyo"
	q.Enqueue(newCity)
	if q.Size() != 1 {
		t.Errorf("%q :size expected 1 received %d", util.GetFuncName(), q.Size())
	}
	city := q.Dequeue().(string)
	if city != newCity {
		t.Errorf("%q :Dequeue expected %q received %q", util.GetFuncName(), newCity, city)
	}
}
