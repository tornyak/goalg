package bag

import (
	"fmt"
	"strings"
)

// element is an element of the Bag
type element struct {
	next  *element
	Value interface{}
}

// linkedListBag structure defines
// Bag implemented as a linked list
type linkedListBag struct {
	first *element
	size  int
}

// NewLinkedBag creates an empty LinkedBag and returns pointer to it
func Linked() Bag { return new(linkedListBag) }

// Add an item to the bag
func (b *linkedListBag) Add(items ...interface{}) Bag {
	for _, it := range items {
		e := &element{b.first, it}
		b.first = e
		b.size++
	}
	return b
}

// IsEmpty returns true if the Bag is empty, othewise false
func (b *linkedListBag) IsEmpty() bool {
	return b.size == 0
}

// Size returns number of items in the Bag
func (b *linkedListBag) Size() int {
	return b.size
}

// Size returns number of items in the Bag
func (b *linkedListBag) ForEach(f func(interface{})) Bag {
	for e := b.first; e != nil; e = e.next {
		f(e.Value)
	}
	return b
}

// String returns formatted string representing LinkedBag and its elements
func (b *linkedListBag) String() string  {
	ret := "["
	b.ForEach(func(e interface{}) {
		ret+= fmt.Sprintf("%v ", e)
	})
	return strings.TrimSpace(ret) + "]"
}
