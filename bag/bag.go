// Package bag implements a Bag collection where removing items is not
// supported. Its purpose is to provide clients with the ability to collect
// items and then to iterate through the collected items (the client can also
// test if a bag is empty and find its number of items).
// The order of iteration is unspecified and should be immaterial to the client.
package bag

// Bag interface that different Bag
// implementations must implement
type Bag interface {
	Add(...interface{}) Bag
	IsEmpty() bool
	Size() int
	ForEach(func(interface{})) Bag
}
