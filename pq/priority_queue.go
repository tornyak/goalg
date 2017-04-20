package pq


// TODO should it be extracted??, now used in st adn pq

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

type MaxPQ interface {
	Insert(v Key) // Insert a key into the priority queue
	Max() Key // Return the largest key
	DelMax() // Retrurn and remove the largest key
	IsEmpty() bool // Is the PQ empty?
	Size() int // Number of keys in the priority queue
}
