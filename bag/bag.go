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
	size  uint
}

type Iterator struct {
	current *Element
}

func (it *Iterator) Next() bool {
	if it.current != nil {
		it.current = it.current.next
	}
	if it.current == nil {
		return false
	}
	return true
}

func (it *Iterator) Value() interface{} {
	if it.current != nil {
		return it.current.Value
	} else {
		return nil
	}
}

func New() *Bag { return new(Bag).init() }

func (b *Bag) init() *Bag {
	b.first = nil
	b.size = 0
	return b
}

// Add an item to the bag
func (b *Bag) Add(v interface{}) {
	e := &Element{b.first, v}
	b.first = e
	b.size++
}

// IsEmpty returns true if the Bag is empty, othewise false
func (b *Bag) IsEmpty() bool {
	return b.size == 0
}

// Size returns number of items in the Bag
func (b *Bag) Size() uint {
	return b.size
}

func (b *Bag) GetIterator() *Iterator {
	it := new(Iterator)
	it.current = new(Element)
	it.current.next = b.first
	return it
}
