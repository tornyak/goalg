package util

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSliceUtil(t *testing.T) {
	log.Println("<TEST> TestSliceUtil")
	log.Println("testReverse")
	testReverse(t)
}

func testReverse(t *testing.T) {
	var input []interface{}
	var expected []interface{}
	Reverse(input)
	assert.Equal(t, expected, input, "Reverse failed for empty slice")
	input = append(input, 5)
	expected = append(expected, 5)
	assert.Equal(t, expected, input, "Reverse failed for slice with one element")
	input = append(input, "Stockholm")
	expected = append(expected, "Stockholm")
	assert.Equal(t, expected, input, "Reverse failed for slice with two elements")
	input = append(input, 3.14)
	expected = append(expected, 3.14)
	assert.Equal(t, expected, input, "Reverse failed for slice with three elements")

}
