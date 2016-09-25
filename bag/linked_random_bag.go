package bag

import (
	"github.com/tornyak/goalg/util/slice"
	"strings"
	"fmt"
)

// randomLinkedListBag is a Bag where iterator provides
// elements in random order. All permutations are
// equaly probable
type randomLinkedListBag struct {
	linkedListBag
}

// RandomList creates an empty Bag backed with linked list whose elements
// are traversed in random order. Returns a pointer to created Bag
func RandomLinked() Bag { return new(randomLinkedListBag) }

// Add an item to the bag
func (b *randomLinkedListBag) Add(items ...interface{}) Bag {
	b.linkedListBag.Add(items...)
	return b
}

// ForEach applies function f to each element of the randomLinkedListBag
// Bag elements are traversed in random order
func (b *randomLinkedListBag) ForEach(f func(interface{})) Bag  {
	var tmp []interface{}
	for e := b.first; e != nil; e = e.next {
		tmp = append(tmp, e.Value)
	}
	slice.Shuffle(tmp)
	for _, e := range tmp {
		f(e)
	}
	return b
}

// String returns formated string representing LinkedBag and its elements
func (b *randomLinkedListBag) String() string  {
	ret := "["
	b.ForEach(func(e interface{}) {
		ret+= fmt.Sprintf("%v ", e)
	})
	return strings.TrimSpace(ret) + "]"
}
