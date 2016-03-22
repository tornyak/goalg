package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	var input []interface{}
	var expected []interface{}
	Reverse(input)
	assert.Equal(t, expected, input, "Reverse failed for empty slice")
	input = append(input, 5)
	Reverse(input)
	expected = append(expected, 5)
	assert.Equal(t, expected, input, "Reverse failed for slice with one element")

	input = make([]interface{}, 2)
	input[0] = 5
	input[1] = "Stockholm"
	Reverse(input)

	expected = make([]interface{}, 2)
	expected[0] = "Stockholm"
	expected[1] = 5

	assert.Equal(t, expected, input, "Reverse failed for slice with two elements")

	input = append(input, 3.14)
	expected = append(expected, 3.14)
	assert.Equal(t, expected, input, "Reverse failed for slice with three elements")

}

func TestRemove(t *testing.T)  {
	tests := []struct{
		A []interface{}
		I int
		E []interface{}
	}{
		{A: []interface{}{}, I: 2, E: []interface{}{}},
		{A: []interface{}{1}, I: -1, E: []interface{}{1}},
		{A: []interface{}{1}, I: 0, E: []interface{}{}},
		{A: []interface{}{1,2,3}, I: 0, E: []interface{}{2,3}},
		{A: []interface{}{1,2,3}, I: 1, E: []interface{}{1,3}},
		{A: []interface{}{1,2,3}, I: 2, E: []interface{}{1,2}},
	}

	for id, mTest := range tests {
		res, _ := Remove([]interface{}(mTest.A), mTest.I)
		assert.Equal(t, mTest.E,res, "TestRemove, id: %v Remove Fail", id)
	}
}
