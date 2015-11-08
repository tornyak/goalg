package bag

// element is an element of the Bag
type element struct {
	next  *element
	Value interface{}
}

// LinkedBag structure defines
// Bag implemented as a linked list
type LinkedBag struct {
	first *element
	size  int
}

// LinkedBagIterator data structure
type LinkedBagIterator struct {
	current *element
}

// NewLinkedBag creates an empty LinkedBag and returns pointer to it
func NewLinkedBag() *LinkedBag { return new(LinkedBag) }

// Add an item to the bag
func (b *LinkedBag) Add(v interface{}) {
	e := &element{b.first, v}
	b.first = e
	b.size++
}

// IsEmpty returns true if the Bag is empty, othewise false
func (b *LinkedBag) IsEmpty() bool {
	return b.size == 0
}

// Size returns number of items in the Bag
func (b *LinkedBag) Size() int {
	return b.size
}

// GetIterator returns a Bag iterator
func (b *LinkedBag) GetIterator() Iterable {
	it := new(LinkedBagIterator)
	it.current = b.first
	return it
}

// HasNext returns true if there is next element of
// the iterator, othewise it just returns false
func (it *LinkedBagIterator) HasNext() bool {
	return it.current != nil
}

// Next returns the value of the next element. If there is no
// next element nil is returned
func (it *LinkedBagIterator) Next() interface{} {
	if it.current != nil {
		curr := it.current
		it.current = it.current.next
		return curr.Value
	}
	return nil
}
