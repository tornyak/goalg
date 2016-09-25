package bag_test

import (
	"fmt"
	"github.com/tornyak/goalg/bag"
	"testing"
)

func TestBagIsEmpty(t *testing.T) {
	var isEmptyTests = []struct {
		bag      bag.Bag
		expected bool // expected result
	}{
		{bag.Slice(), true},
		{bag.Slice().Add("Hello"), false},
		{bag.Slice().Add("New York", "Stockholm", "London", "Paris"), false},
		{bag.Slice().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), false},
	}

	for _, test := range isEmptyTests {
		if test.bag.IsEmpty() != test.expected {
			t.Errorf("IsEmpty() for Bag: %v, expected: %v, was: %v", test.bag, test.expected, test.bag.IsEmpty())
		}
	}
}

func TestBagSize(t *testing.T) {
	var sizeTests = []struct {
		bag      bag.Bag
		expected int
	}{
		{bag.Slice(), 0},
		{bag.Slice().Add("Hello"), 1},
		{bag.Slice().Add("New York", "Stockholm", "London", "Paris"), 4},
		{bag.Slice().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), 4},
	}

	for _, test := range sizeTests {
		if test.bag.Size() != test.expected {
			t.Errorf("Size() for Bag: %v, expected: %v, was: %v", test.bag, test.expected, test.bag.Size())
		}
	}
}

func TestForEachEmpty(t *testing.T) {
	b := bag.Slice()
	b.ForEach(func(interface{}) {
		t.Error("Empty bag ForEach called function")
	})
}

func TestForEachOne(t *testing.T) {
	b := bag.Slice().Add(5)
	i := 0
	b.ForEach(func(e interface{}) {
		i++
		if e != 5 {
			t.Errorf("ForEach got wrong element: %v", e)
		}
	})
	if i != 1 {

	}
}

func TestForEachMultiple(t *testing.T) {
	b := bag.Slice()
	cities := map[string]int{
		"NewBag York": 1,
		"Stockholm":   1,
		"London":      1,
		"Paris":       1}

	for city := range cities {
		b.Add(city)
	}

	b.ForEach(func(e interface{}) {
		cities[e.(string)]--
	})

	// Test that all cities in map were found
	for city, count := range cities {
		if count != 0 {
			t.Errorf("Wrong count %d for city: %s", count, city)
		}
	}
}

func TestString(t *testing.T) {
	var sizeTests = []struct {
		bag      bag.Bag
		expected string
	}{
		{bag.Slice(), "[]"},
		{bag.Slice().Add("Hello"), "[Hello]"},
		{bag.Slice().Add("New York", "Stockholm", "London", "Paris"), "[New York Stockholm London Paris]"},
		{bag.Slice().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), "[New York 1.23 1234567890 [1 2 3]]"},
	}

	for _, test := range sizeTests {
		if fmt.Sprint(test.bag) != test.expected {
			t.Errorf("String() for Bag: %v, expected: %v", test.bag, test.expected)
		}
	}
}
