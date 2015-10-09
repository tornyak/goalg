// Package bag implements a Bag collection where removing items is not
// supported. Its purpose is to provide clients with the ability to collect
// items and then to iterate through the collected items (the client can also
// test if a bag is empty and find its number of items).
// The order of iteration is unspecified and should be immaterial to the client.
package bag

// Element is an element of the Bag
type Element struct {
	next  *Element
	Value interface{}
}

// A Bag structure
type Bag struct {
	first *Element
	size  int
}

// Iterator data structure
type Iterator struct {
	current *Element
}

// Next moves iterator to the next element in collection
// if it exists and returns true, othewise it just returns false
func (it *Iterator) Next() bool {
	if it.current != nil {
		it.current = it.current.next
	}
	if it.current == nil {
		return false
	}
	return true
}

// Value returns the value of the current element or nil
// if current element is nil
func (it *Iterator) Value() interface{} {
	if it.current != nil {
		return it.current.Value
	}
	return nil
}

// New creates a new empty bag and returns a pointer to it
func New() *Bag { return new(Bag) }

// Add items to the bag
func (b *Bag) Add(v ...interface{}) {
	for _, item := range v {
		b.add(item)
	}
}

// Add an item to the bag
func (b *Bag) add(v interface{}) {
	e := &Element{b.first, v}
	b.first = e
	b.size++
}

// IsEmpty returns true if the Bag is empty, othewise false
func (b *Bag) IsEmpty() bool {
	return b.size == 0
}

// Size returns number of items in the Bag
func (b *Bag) Size() int {
	return b.size
}

func (b *Bag) GetIterator() *Iterator {
	it := new(Iterator)
	it.current = new(Element)
	it.current.next = b.first
	return it
}
