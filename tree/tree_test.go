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

	assert.Equal(t, node.parent.orient, "U")
	assert.Equal(t, node2.parent.orient, "U")
	assert.Equal(t, node3.parent.orient, "L")

	assert.True(t, node3.ContainsPoint(point3))

}

func TestNode2(t *testing.T) {
	tree := NewTree()
	point := geo.NewPoint("0", "0")
	point2 := geo.NewPoint("00", "00")
	point3 := geo.NewPoint("001", "001")

	tree.InsertPoint(point)
	tree.InsertPoint(point2)
	tree.InsertPoint(point3)

	node := tree.Find(point)
	assert.Equal(t, node.GetHash(), "00")

	node2 := tree.Find(point2)
	assert.Equal(t, node2.GetHash(), "0000")

	node3 := tree.Find(point3)
	assert.Equal(t, node3.GetHash(), "100000")

	assert.Equal(t, node.parent.orient, "U")

	assert.Equal(t, node2.parent.orient, "R")

	assert.Equal(t, node3.parent.orient, "U")

	assert.True(t, node3.ContainsPoint(point3))

}
