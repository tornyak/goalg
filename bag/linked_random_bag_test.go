package bag_test

import (
	"github.com/tornyak/goalg/bag"
	"testing"
	"fmt"
)

func TestRandomBagIsEmpty(t *testing.T) {
	var isEmptyTests = []struct {
		bag      bag.Bag
		expected bool // expected result
	}{
		{bag.RandomLinked(), true},
		{bag.RandomLinked().Add("Hello"), false},
		{bag.RandomLinked().Add("New York", "Stockholm", "London", "Paris"), false},
		{bag.RandomLinked().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), false},
	}

	for _, test := range isEmptyTests {
		if test.bag.IsEmpty() != test.expected {
			t.Errorf("IsEmpty() for Bag: %v, expected: %v, was: %v", test.bag, test.expected, test.bag.IsEmpty())
		}
	}
}

func TestRandomBagSize(t *testing.T) {
	var sizeTests = []struct {
		bag      bag.Bag
		expected int
	}{
		{bag.RandomLinked(), 0},
		{bag.RandomLinked().Add("Hello"), 1},
		{bag.RandomLinked().Add("New York", "Stockholm", "London", "Paris"), 4},
		{bag.RandomLinked().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), 4},
	}

	for _, test := range sizeTests {
		if test.bag.Size() != test.expected {
			t.Errorf("Size() for Bag: %v, expected: %v, was: %v", test.bag, test.expected, test.bag.Size())
		}
	}
}

func TestRandomBagForEachEmpty(t *testing.T) {
	b := bag.RandomLinked()
	b.ForEach(func(interface{}){
		t.Error("Empty bag ForEach called function")
	})
}

func TestRandomBagForEachOne(t *testing.T) {
	b := bag.RandomLinked().Add(5)
	i := 0
	b.ForEach(func(e interface{}){
		i++
		if e != 5 {
			t.Error("ForEach got wrong element: %v", e)
		}
	})
	if i != 1 {

	}
}

func TestRandomBagForEachMultiple(t *testing.T) {
	b := bag.RandomLinked()
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

func TestRandomBagString(t *testing.T) {
	var sizeTests = []struct {
		bag      bag.Bag
		expected string
	}{
		{bag.RandomLinked(), "[]"},
		{bag.RandomLinked().Add("Hello"), "[Hello]"},
	}

	for _, test := range sizeTests {
		if fmt.Sprint(test.bag) != test.expected {
			t.Errorf("String() for Bag: %v, expected: %v", test.bag, test.expected)
		}
	}
}
