package bag_test

import (
	"testing"
	"github.com/tornyak/goalg/bag"
	"fmt"
)

func TestLinkedBagIsEmpty(t *testing.T) {
	var isEmptyTests = []struct {
		bag      bag.Bag
		expected bool // expected result
	}{
		{bag.List(), true},
		{bag.List().Add("Hello"), false},
		{bag.List().Add("New York", "Stockholm", "London", "Paris"), false},
		{bag.List().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), false},
	}

	for _, test := range isEmptyTests {
		if test.bag.IsEmpty() != test.expected {
			t.Errorf("IsEmpty() for Bag: %v, expected: %v, was: %v", test.bag, test.expected, test.bag.IsEmpty())
		}
	}
}


func TestLinkedBagSize(t *testing.T) {
	var sizeTests = []struct {
		bag      bag.Bag
		expected int
	}{
		{bag.List(), 0},
		{bag.List().Add("Hello"), 1},
		{bag.List().Add("New York", "Stockholm", "London", "Paris"), 4},
		{bag.List().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), 4},
	}

	for _, test := range sizeTests {
		if test.bag.Size() != test.expected {
			t.Errorf("Size() for Bag: %v, expected: %v, was: %v", test.bag, test.expected, test.bag.Size())
		}
	}
}

func TestLinkedBagForEachEmpty(t *testing.T) {
	b := bag.List()
	b.ForEach(func(interface{}){
		t.Error("Empty bag ForEach called function")
	})
}

func TestLinkedBagForEachOne(t *testing.T) {
	b := bag.List().Add(5)
	i := 0
	b.ForEach(func(e interface{}){
		i++
		if e != 5 {
			t.Errorf("ForEach got wrong element: %v", e)
		}
	})
	if i != 1 {

	}
}

func TestLinkedBagForEachMultiple(t *testing.T) {
	b := bag.List()
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

func TestLinkedBagString(t *testing.T) {
	var sizeTests = []struct {
		bag      bag.Bag
		expected string
	}{
		{bag.List(), "bag.linkedListBag: []"},
		{bag.List().Add("Hello"), "bag.linkedListBag: [Hello]"},
		{bag.List().Add("New York", "Stockholm", "London", "Paris"), "bag.linkedListBag: [Paris London Stockholm New York]"},
		{bag.List().Add("New York", 1.23, 1234567890, []int{1, 2, 3}), "bag.linkedListBag: [[1 2 3] 1234567890 1.23 New York]"},
	}

	for _, test := range sizeTests {
		if fmt.Sprint(test.bag) != test.expected {
			t.Errorf("String() for Bag: %v, expected: %v", test.bag, test.expected)
		}
	}
}
