package knapsack_test

import (
	"testing"
	"github.com/tornyak/goalg/knapsack"
)

func TestUnbounded(t *testing.T) {
	var tests = []struct{
		items []knapsack.Item
		cap int
		expected int
	}{
		{items: []knapsack.Item{}, cap: 0, expected: 0},
		{items: []knapsack.Item{{1, 1}}, cap: 0, expected: 0},
		{items: []knapsack.Item{{1, 1}}, cap: 1, expected: 1},
		{items: []knapsack.Item{{1, 1},{2, 3} }, cap: 4, expected: 6},
		{items: []knapsack.Item{{3, 4},{4, 5},{7, 10},{8, 11},{9, 13} }, cap: 17, expected: 24},
		{items: []knapsack.Item{{12, 24},{7, 13},{11, 23},{8, 15},{9, 16}}, cap: 26, expected: 51},
		{items: []knapsack.Item{{23, 92},{31, 57},{29, 49},{44, 68},{53, 60},{38, 43},{63, 67},{85, 84},{89, 87},{82, 72}}, cap: 165, expected: 644},
	}

	for _, test := range tests {
		ks := knapsack.Unbounded(test.cap, test.items)
		if ks.ItemValue != test.expected {
			t.Errorf("Wrong max found for cap %d, items: %v: expected %v, was %v",
				test.cap, test.items, test.expected, ks.ItemValue)
		}
	}
}
