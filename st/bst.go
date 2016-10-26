package st

// node of the tree
type node struct {
	key   Comparable
	value interface{}
	left  *node
	right *node
	nodes int // number of nodes in subtree rooted in this node
}

// BST is Binary Search Tree based implementation of the SymbolTable
type BST struct {
	root *node
}

func NewBST() OrderedSymbolTable {
	return new(BST)
}

func (bst *BST) Put(key Key, value interface{}) {
	bst.root = bst.put(bst.root, key, value)
}

func (bst *BST) put(x *node, key Key, value interface{}) *node {
	if x == nil {
		return &node{key: key, value: value, }
	}

	cmp, err := key.Compare(x.key)
	if err != nil {
		return  nil
	}

	switch {
	case cmp < 0:
		x.left = bst.put(x.left, key, value)
	case cmp > 0:
		x.right = bst.put(x.right, key, value)
	default:
		x.value = value
	}
	x.nodes = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *BST) Get(key Key) interface{} {
	return bst.get(bst.root, key)
}

func (bst *BST) get(x *node, key Key) interface{} {
	if x == nil {
		return nil
	}
	cmp, err := key.Compare(x.key)
	if err != nil {
		return nil
	}
	switch {
	case cmp < 0:
		return bst.get(x.left, key)
	case cmp > 0:
		return bst.get(x.right, key)
	default:
		return x.value
	}
}

// Delete is using Hibbard deletion algorithm
func (bst *BST) Delete(key Key) {
	if bst.root != nil {
		bst.root = bst.delete(bst.root, key)
	}
}

func (bst *BST) delete(x *node, key Key) *node {

	cmp, err := key.Compare(x.key)
	if err != nil {
		return x
	}

	switch {
	case cmp < 0:
		x.left = bst.delete(x.left, key)
	case cmp > 0:
		x.right = bst.delete(x.right, key)
	default:
		// found element, need to remove from the tree
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		t := x
		x = bst.min(t.right)
		x.right = bst.deleteMin(t.right)
		x.left = t.left
	}
	x.nodes = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *BST)DeleteMin() {
	if bst.root == nil {
		return
	}
	bst.root = bst.deleteMin(bst.root)
}

func (bst *BST)deleteMin(x *node) *node {
	if x.left == nil { // x is min node
		return x.right
	}
	x.left = bst.deleteMin(x.left)
	x.nodes = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *BST) Contains(key Key) bool {
	return bst.get(bst.root, key) != nil
}

func (bst *BST) IsEmpty() bool {
	return bst.size(bst.root) == 0
}

func (bst *BST) Size() int {
	return bst.size(bst.root)
}

func (bst *BST) size(x *node) int {
	if x == nil {
		return 0
	}
	return x.nodes
}

func (bst *BST) Keys() []Key {
	return nil
}

func (bst *BST) Min() Key {
	if bst.root == nil {
		return nil
	}
	return bst.min(bst.root).key
}

func (bst *BST) min(x *node) *node {
	if x.left == nil {
		return x
	}
	return bst.min(x.left)
}

// Max returns the largest key
func (bst *BST) Max() Key {
	if bst.root == nil {
		return nil
	}
	return bst.max(bst.root)
}

func (bst *BST) max(x *node) *node {
	if x.right == nil {
		return x
	}
	return bst.max(x.right)
}

// Floor returns largest key less then or equal to key
func (bst *BST) Floor(key Key) Key {
	var n *node
	if n := bst.floor(bst.root, key); n == nil {
		return nil
	}
	return n.key

}

func (bst *BST)floor(x *node, key Key) *node {

	if x == nil {
		return nil
	}

	c, err := key.Compare(x.key)
	if err != nil {
		return nil
	}

	switch c {
	case 0:
		return x
	case -1:
		return bst.floor(x.left, key)
	default:
		if t := bst.floor(x.right, key); t != nil {
			return t
		}
		return x

	}
}

// Ceiling returns smallest key greater then or equal to key
func (bst *BST) Ceiling(key Key) Key {
	// TODO
	return nil
}

// Rank returns number of keys less then key
func (bst *BST) Rank(key Key) int {
	// TODO
	return 0
}

// Select returns a key of rank k
func (bst *BST) Select(k int) Key {
	// TODO
	return nil
}

// DeleteMax deletes largest key
func (bst *BST) DeleteMax() {
	// TODO
}

// SizeRange returns number of keys in [lo..hi]
func (bst *BST) SizeRange(lo, hi Key) int {
	// TODO
	return 0
}

// KeysRange returns slice of keys in [lo..hi], in
func (bst *BST) KeysRange(lo, hi Key) []Key {
	// TODO
	return nil
}
