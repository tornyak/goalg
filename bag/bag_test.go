package bag

import (
	"github.com/tornyak/goalg/util"
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
	testIterateMultiple(t)
}

func testCreateEmpty(t *testing.T) {
	b := New()
	if !b.IsEmpty() {
		t.Error(util.GetFuncName(), ":bag not empty")
	}
	if b.Size() != 0 {
		t.Errorf("%q size expected 0 received %d", util.GetFuncName(), b.Size())
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
	b := New()
	cities := map[string]int{
		"New York":  1,
		"Stockholm": 1,
		"London":    1,
		"Paris":     1}

	for city := range cities {
		b.Add(city)
	}

	it := b.GetIterator()
	for it.Next() {
		city := it.Value().(string)
		cities[city]--
	}

	// Test that all cities in map were found
	for city, count := range cities {
		if count != 0 {
			t.Errorf("testIterateMultiple: count for city %q is %d instead of 0", city, count)
		}
	}
}
