package knapsack

const (
	undef = -1
)

// N types of items of varying size and value
// M Size of knapsack
// Find combination of items to put in knapsack that would have maximum value

// Item represents an item with its size and value
type Item struct {
	Size int
	Value int
}

// item with quantity
type itemQuantity struct {
	Item
	Quantity int
}

// Knapsack represents a knapsack filled with items
type Knapsack struct {
	Size int
	ItemValue int
	Items []itemQuantity
}

// Unbounded recursively solves unbounded knapsack problem where there
// is no limit on available number of items.
// cap capacity of the knapsack
// items slice of available Items
// returns pointer to Knapsack containing selected items and their value
func Unbounded(cap int, items []Item) *Knapsack {

	qItems := make([]itemQuantity, len(items))
	for i, it := range items {
		qItems[i] = itemQuantity{it,0}
	}

	ks := &Knapsack{
		Size: cap,
		Items: qItems,
	}

	if cap == 0 || len(items) == 0 {
		return ks
	}

	maxKnown := make([]int, cap+1)			// slice containing maximal value for capacity
	itemKnownIndex := make([]int, cap+1)	// slice containing index of item that maximizes value for capacity
	for i := 0; i < len(maxKnown); i++ {
		maxKnown[i] = undef
		itemKnownIndex[i] = undef
	}

	ks.ItemValue = unboundedMemo(cap, items, itemKnownIndex, maxKnown)

	// find items that go in the knapsack
	itemIdx := itemKnownIndex[cap]
	for value := cap; value > 0;  {
		ks.Items[itemIdx].Quantity++
		value-=ks.Items[itemIdx].Size
		if value >= 0 {
			itemIdx = itemKnownIndex[value]
		}
	}

	return ks
}

func unboundedMemo(cap int, items []Item, itemKnownIndex, maxKnown  []int) int {
	var max, maxIndex int

	// if solution is already known just return the result
	if maxKnown[cap] != undef {
		return maxKnown[cap]
	}
	for i, it := range items {
		// is there enough space in knapsack to take item it?
		if space := cap - it.Size; space >= 0 {
			// Does taking item it increases the maximum value or not?
			if t := unboundedMemo(space, items, itemKnownIndex, maxKnown) + it.Value; t > max {
				max = t
				maxIndex = i
			}
		}
	}
	maxKnown[cap] = max
	itemKnownIndex[cap] = maxIndex
	return max
}

