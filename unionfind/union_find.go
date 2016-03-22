package unionfind

import "fmt"

// UnionFind represents interface for union-find algorithm
type UnionFind interface {
	Union(int, int)
	Find(int) (int, error)
	Connected(int, int) bool
	Count() int
}

// unionFindCommon implements common data structures shared by all union-find implementations
type unionFindData struct {
	id    []int
	count int
}

// Creates a new UF instance, n is number of components
func newUnionFindData(n int) *unionFindData {
	uf := &unionFindData{
		id:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return uf
}

func NewQuickFind(n int) *QuickFind {
	qf := &QuickFind{
		unionFindData: *newUnionFindData(n),
	}
	return qf
}

// Count returns number of components
func (uf *unionFindData) Count() int {
	return uf.count
}

// QuickFind - complexity: constructor N, union N, find 1
type QuickFind struct {
	unionFindData
}

// Union assigns ID of component containing q to all elements of component containing p
func (qf *QuickFind) Union(p, q int) {
	pID, pErr := qf.Find(p)
	if pErr != nil {
		return
	}
	qID, qErr := qf.Find(q)
	if qErr != nil {
		return
	}
	if pID == qID {
		return
	}
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pID {
			qf.id[i] = qID
		}
	}
	qf.count--
}

// Connected returns true if components p and q are connected, otherwise false
func (qf *QuickFind) Connected(p, q int) bool {
	pFind, pErr := qf.Find(p)
	if pErr != nil {
		return false
	}
	qFind, qErr := qf.Find(q)
	if qErr != nil {
		return false
	}
	return qFind == pFind
}

// Find returns ID of component containing p. Just indexing slice
func (qf *QuickFind) Find(p int) (int, error) {
	if p >= len(qf.id) {
		return 0, fmt.Errorf("Find(p) failed, p out of bounds: p=%d qf.count=%d", p, qf.count)
	}
	return qf.id[p], nil
}

// QuickUnion - complexity: constructor N, union tree height, find tree height
type QuickUnion struct {
	unionFindData
}

// WeightedQuickUnion - complexity: constructor N, union lgN, find lgN
type WeightedQuickUnion struct {
	unionFindData
}

// WeightedQuickUnionCompress - complexity: constructor N, union >>1, find >>1
type WeightedQuickUnionCompress struct {
	unionFindData
}
