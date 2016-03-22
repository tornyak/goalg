package unionfind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickUnionFind(t *testing.T) {
	uf := NewQuickFind(10)
	assert.Equal(t, uf.Count(), 10, "Count returned wrong number of components")
	uf.Union(1, 3)
	assert.Equal(t, uf.Count(), 9, "Count returned wrong number of components")
	c, err := uf.Find(1)
	assert.Equal(t, c, 3, "Find returned wrong component value")
	uf.Union(3, 3)
	assert.Equal(t, uf.Count(), 9, "Count returned wrong number of components")
	c, err = uf.Find(1)
	assert.Equal(t, c, 3, "Find returned wrong component value")
	uf.Union(1, 5)
	uf.Union(3, 5)
	uf.Union(3, 7)
	uf.Union(2, 4)
	// Same element
	assert.True(t, uf.Connected(1, 1), "Wrong value from Connected")
	assert.True(t, uf.Connected(1, 7), "Wrong value from Connected")
	assert.False(t, uf.Connected(1, 2), "Wrong value from Connected")
	assert.False(t, uf.Connected(1, 33), "Connected out of range should return false")
	assert.False(t, uf.Connected(33, 1), "Connected out of range should return false")
	// Get out of range
	uf.Union(33, 1)
	uf.Union(1, 33)
	c, err = uf.Find(33)
	assert.NotNil(t, err, "Find returned nil error")
}
