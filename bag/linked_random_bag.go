package bag

import "github.com/tornyak/goalg/util"

// LinkedRandomBag is a Bag where iterator provides
// elements in random order. All permutations are
// equaly probable
type LinkedRandomBag struct {
	LinkedBag
}

// NewLinkedRandomBag creates an empty LinkedBag and returns pointer to it
func NewLinkedRandomBag() Bag { return new(LinkedRandomBag) }

// GetIterator returns a Bag iterator
func (b *LinkedRandomBag) getIterator(shuffle func([]interface{})) Iterable {
	randomIt := new(LinkedRandomBagIterator)
	randomIt.elements = make([]interface{}, 0)
	it := b.LinkedBag.GetIterator()
	for it.HasNext() {
		randomIt.elements = append(randomIt.elements, it.Next())
	}
	shuffle(randomIt.elements)
	return randomIt
}

// GetIterator returns a Bag iterator using default shuffling
// function based on Fisher-Yates (Knuth) shuffle
func (b *LinkedRandomBag) GetIterator() Iterable {
	return b.getIterator(util.Shuffle)
}

// LinkedRandomBagIterator data structure
type LinkedRandomBagIterator struct {
	// random shuffle of bag elements
	elements []interface{}
	current  int
}

// HasNext returns true if there is next element of
// the iterator, othewise it just returns false
func (it *LinkedRandomBagIterator) HasNext() bool {
	return it.current < len(it.elements)
}

// Next returns the value of the next element. If there is no
// next element error is set
func (it *LinkedRandomBagIterator) Next() interface{} {
	if it.current < len(it.elements) {
		curr := it.current
		it.current++
		return it.elements[curr]
	}
	return nil
}
