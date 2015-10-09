package bag

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tornyak/goalg/util"
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
	assert.True(t, b.IsEmpty(), "bag not empty")
	assert.Equal(t, 0, b.Size(), util.GetFuncName()+" :wrong size")
}

func testAddOne(t *testing.T) {
	b := New()
	b.Add("Hello")
	assert.Equal(t, 1, b.Size(), util.GetFuncName()+" :wrong size")
}

func testAddMultiple(t *testing.T) {
	b := New()
	b.Add("New York", "Stockholm", "London", "Paris")
	assert.Equal(t, 4, b.Size(), util.GetFuncName()+" :wrong size")
	assert.False(t, b.IsEmpty(), util.GetFuncName()+" :bag empty")
}

func testAddDifferentTypes(t *testing.T) {
	b := New()
	b.Add("New York", 1.23, time.Now(), []int{1, 2, 3})
	assert.Equal(t, 4, b.Size(), util.GetFuncName()+" :wrong size")
	assert.False(t, b.IsEmpty(), util.GetFuncName()+" :bag empty")
}

func testIterateEmpty(t *testing.T) {
	b := New()
	it := b.GetIterator()
	for it.Next() {
		t.Errorf("%q Iterator should not have next element", util.GetFuncName())
	}
}

func testIterateOne(t *testing.T) {
	b := New()
	b.Add(5)
	it := b.GetIterator()
	for it.Next() {
		v := it.Value().(int)
		if v != 5 {
			t.Errorf("%q expected value was 5 received %d", util.GetFuncName(), v)
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

	for it := b.GetIterator(); it.Next(); {
		city := it.Value().(string)
		cities[city]--
	}

	// Test that all cities in map were found
	for city, count := range cities {
		if count != 0 {
			t.Errorf("%q count for city %q is %d instead of 0", util.GetFuncName(), city, count)
		}
	}
}

func testEmptyIteratorGetValue(t *testing.T) {
	b := New()
	it := b.GetIterator()

	if v := it.Value(); v != nil {
		t.Errorf("%q empty iterator did not return nil but %v", util.GetFuncName(), v)
	}
}
