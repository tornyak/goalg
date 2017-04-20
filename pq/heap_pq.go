package pq

type heapPQ struct {
	heap []Key // Heap ordered complete binary tree, pq[0] not used
	n int
}

// NewHeap returns en empty heap based MaxPQ
func NewHeap(size int) MaxPQ {
	return &heapPQ{
		heap: make([]Key, size+1),
	}
}

// Insert a key into the priority queue
func (pq *heapPQ) Insert(v Key) {
	pq.n++
	pq.heap[pq.n] = v
	pq.swim(pq.n)
}

// Max returns the largest key
func (pq *heapPQ) Max() Key {
	return pq.heap[1]
}

// DelMax returns and removes the largest key
func (pq *heapPQ) DelMax() Key {
	max := pq[1]
	pq.exchange(1, pq.n)
	pq.n--
	pq.sink(1)
	return max
}

// IsEmpty returns true if the priority queue is empty, otherwise false
func (pq *heapPQ) IsEmpty() bool {
	return pq.n == 0
}

// Size returns the number of keys in the priority queue
func (pq *heapPQ) Size() int {
	return pq.n
}

func (pq *heapPQ) exchange(i,j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func (pq *heapPQ) less(i,j int) bool {
	return pq.heap[i].Compare(pq.heap[j]) < 1
}

func (pq *heapPQ) sink(k int) {
	for 2*k <= pq.n {
		j := 2*k
		if j < pq.n && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.exchange(k, j)
		k = j
	}
}

func (pq *heapPQ) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exchange(k/2, k)
		k = k/2
	}
}
