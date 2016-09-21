package bag

import (
	"fmt"
	"reflect"
)

// sliceBag structure defines
// Bag data structure implemented as a slice
type sliceBag struct {
	a []interface{}
}

// Slice creates an empty slice Bag and returns a pointer to it
func Slice() Bag {
	return &sliceBag{
		a: make([]interface{},0),
	}
}

// Add an item to the bag
func (b *sliceBag) Add(items ...interface{}) Bag {
	for _, it := range items {
		b.a = append(b.a, it)
	}
	return b
}

// IsEmpty returns true if the Bag is empty, othewise false
func (b *sliceBag) IsEmpty() bool {
	return len(b.a) == 0
}

// Size returns number of items in the Bag
func (b *sliceBag) Size() int {
	return len(b.a)
}

// ForEach traverses Bag and applies function f to each Bag element
func (b *sliceBag) ForEach(f func(interface{})) Bag {
	for _, e := range b.a {
		f(e)
	}
	return b
}

// String returns formated string representing LinkedBag and its elements
func (b *sliceBag) String() string  {
	t := reflect.TypeOf(b).Elem()
	return fmt.Sprintf("%v: %v", t, b.a)
}
