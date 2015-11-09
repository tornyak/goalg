package bag

import (
	"github.com/stretchr/testify/assert"
	"github.com/tornyak/goalg/util"
	"testing"
	"time"
)

func TestLinkedRandomBagCreateEmpty(t *testing.T) {
	b := NewLinkedRandomBag()
	assert.True(t, b.IsEmpty(), "Bag not empty")
	assert.Equal(t, b.Size(), 0, "Wrong Bag size")
}

func TestLinkedRandomBagAddOne(t *testing.T) {
	b := NewLinkedRandomBag()
	b.Add("Hello")
	assert.Equal(t, 1, b.Size(), "Wrong Bag size")
}

func TestLinkedRandomBagAddMultiple(t *testing.T) {
	b := NewLinkedRandomBag()
	b.Add("NewBag York")
	b.Add("Stockholm")
	b.Add("London")
	b.Add("Paris")
	assert.Equal(t, 4, b.Size(), "Wrong Bag size")
	assert.False(t, b.IsEmpty(), "Bag empty")
}

func TestLinkedRandomBagAddDifferentTypes(t *testing.T) {
	b := NewLinkedRandomBag()
	b.Add("NewBag York")
	b.Add(1.23)
	b.Add(time.Now())
	b.Add([]int{1, 2, 3})
	assert.Equal(t, 4, b.Size(), "Wrong Bag size")
	assert.False(t, b.IsEmpty(), "Bag empty")
}

func TestLinkedRandomBagIterateEmpty(t *testing.T) {
	b := NewLinkedRandomBag()
	it := b.GetIterator()
	assert.False(t, it.HasNext(), "Iterator should not have next element")
}

func TestLinkedRandomBagIterateOne(t *testing.T) {
	b := NewLinkedRandomBag()
	b.Add(5)
	it := b.GetIterator()
	for it.HasNext() {
		v := it.Next().(int)
		assert.Equal(t, 5, v, "Iterator returned wrong value")
	}
}

func TestLinkedRandomBagIterateMultiple(t *testing.T) {
	b := NewLinkedRandomBag().(*LinkedRandomBag)
	cities := []string{
		"NewBag York",
		"Stockholm",
		"London",
		"Paris"}

	for _, city := range cities {
		b.Add(city)
	}

	var it1Cities []interface{}
	var it2Cities []interface{}

	// 1st iterator will not shuffle
	it := b.getIterator(func([]interface{}) {})
	for it.HasNext() {
		city := it.Next().(string)
		it1Cities = append(it1Cities, city)
	}

	// 2nd iterator shold return elements in reverse order
	it = b.getIterator(util.Reverse)
	for it.HasNext() {
		city := it.Next().(string)
		it2Cities = append(it2Cities, city)
	}
	util.Reverse(it1Cities)

	assert.Equal(t, it1Cities, it2Cities, "Shuffling does not work for LinkedRandomBag")
}

func TestLinkedRandomBagIterateNextNil(t *testing.T) {
	b := NewLinkedRandomBag()
	b.Add(5)
	it := b.GetIterator()
	for it.HasNext() {
		v := it.Next().(int)
		assert.Equal(t, 5, v, "Iterator returned wrong value")
	}
	assert.Nil(t, it.Next(), "Next should have returned nil")
}
