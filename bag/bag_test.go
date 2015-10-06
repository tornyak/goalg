package bag

import (
	"testing"
	"time"
)

func TestBag(t *testing.T) {
	testCreateEmpty(t)
	testAddOne(t)
	testAddMultiple(t)
	testAddDifferentTypes(t)
	testIterateEmpty(t)
	testIterateOne(t)
}

func testCreateEmpty(t *testing.T) {
	b := New()
	if !b.IsEmpty() {
		t.Error("testCreateEmpty: bag not empty")
	}
}

func testAddOne(t *testing.T) {
	b := New()

	b.Add("Hello")
	if b.Size() != 1 {
		t.Errorf("TestBagSize expected 1 received %d", b.Size())
	}
}

func testAddMultiple(t *testing.T) {
	b := New()
	b.Add("New York")
	b.Add("Stockholm")
	b.Add("London")
	b.Add("Paris")
	if b.Size() != 4 {
		t.Errorf("testAddMultiple expected 4 received %d", b.Size())
	}
	if b.IsEmpty() {
		t.Error("testAddMultiple: bag is empty")
	}
}

func testAddDifferentTypes(t *testing.T) {
	b := New()
	b.Add("New York")
	b.Add(1.23)
	b.Add(time.Now())
	b.Add([]int{1, 2, 3})
	if b.Size() != 4 {
		t.Errorf("testAddMultiple expected 4 received %d", b.Size())
	}
	if b.IsEmpty() {
		t.Error("testAddMultiple: bag is empty")
	}
}

func testIterateEmpty(t *testing.T) {
	b := New()
	it := b.GetIterator()
	for it.Next() {
		t.Error("Iterator should not have next element")
	}
}

func testIterateOne(t *testing.T) {
	b := New()
	b.Add(5)
	it := b.GetIterator()
	for it.Next() {
		v := it.Value().(int)
		if v != 5 {
			t.Errorf("Expected value was 5 received %d", v)
		}
	}
}

func testIterateMultiple(t *testing.T) {
	t.Fail()
}

func testIterateDifferentTypes(t *testing.T) {
	t.Fail()
}
