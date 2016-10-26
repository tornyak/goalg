package st

// Comparable interface with single method Compare
type Comparable interface {
	// Compare compares two values of the same type:
	// A.compare(B) returns:
	//  0 if A == B
	//  1 if A > B
	// -1 if A < B
	// error if A and B are not the same type
	Compare(Comparable) (int, error)
}

type Key Comparable

// Symbol table is generic interface for symbol tables
// Does not allow nil keys or values
// Does not allow duplicate keys, put with existing key and new value will replace the old value
//
type SymbolTable interface {
	Put(key Key, value interface{})		// Puts key-value pair into the table
	Get(Key Key) interface{}			// Get returns value paired with key (null if key is absent)
	Delete(key Key)						// Delete removes key (and its value) from table
	Contains(key Key) bool				// Contains returns true if there is value paired with the key
	IsEmpty() bool 						// is the table empty?
	Size() int 							// Size returns number of key-value pairs in the table
	Keys() []Key						// Keys returns a slice containing all the keys in the table
}

// OrderedSymbolTable takes advantage order among keys to make
// Put and Get operations more efficient
type OrderedSymbolTable interface {
	SymbolTable
	Min() Key 				// Returns smallest key
	Max() Key 				// Return largest key
	Floor(key Key) Key 		// Floor returns largest key less then or equal to key
	Ceiling(key Key)  Key	// Ceiling returns smallest key greater then or equal to key
	Rank(key Key)  int		// Rank returns number of keys less then key
	Select(k int) Key		// Select returns a key of rank k
	DeleteMin()				// DeleteMin deletes smallest key
	DeleteMax()				// DeleteMax deletes largest key
	SizeRange(lo, hi Key) int 	// Size returns number of keys in [lo..hi]
	KeysRange(lo, hi Key) []Key	// Keys returns slice of keys in [lo..hi], in sorted order
}
