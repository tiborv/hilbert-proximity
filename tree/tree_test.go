package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiborv/hilbert-gis/geo"
)

func TestNode(t *testing.T) {
	tree := NewTree()
	assert.Equal(t, tree.GetSector(), "ROOT")

	point := geo.NewPoint("1", "1")
	point2 := geo.NewPoint("11", "10")
	point3 := geo.NewPoint("110", "100")

	tree.InsertPoint(point)
	tree.InsertPoint(point2)
	tree.InsertPoint(point3)

	node := tree.Find(point)
	assert.Equal(t, node.GetHash(), "10")

	node2 := tree.Find(point2)
	assert.Equal(t, node2.GetHash(), "1110")

	node3 := tree.Find(point3)
	assert.Equal(t, node3.GetHash(), "101110")

	assert.True(t, node3.ContainsPoint(point3))

}
