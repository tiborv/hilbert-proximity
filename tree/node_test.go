package tree

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiborv/hilbert-gis/geo"
)

var testTree *Node
var testpoints []geo.Point

func TestTreeCreation(t *testing.T) {
	testTree = NewTree(1)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			x := fmt.Sprintf("%01s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%01s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			testTree.InsertPoint(g)
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			x := fmt.Sprintf("%02s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%02s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			testTree.InsertPoint(g)
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			x := fmt.Sprintf("%03s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%03s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			testTree.InsertPoint(g)
		}
	}
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			x := fmt.Sprintf("%04s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%04s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			testTree.InsertPoint(g)
		}
	}
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			x := fmt.Sprintf("%05s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%05s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			testTree.InsertPoint(g)
		}
	}
	fmt.Println("Points inserted:", pointsInserted)

}

var seen []Node

func (n Node) alreadySeen() bool {
	for _, p := range seen {
		if n.getHash() == p.getHash() {

			return true
		}
	}
	seen = append(seen, n)
	return false
}

func removeFromTestPoints(gp geo.Point) bool {
	for i, p := range testpoints {
		if p.Equals(gp) {
			testpoints = append(testpoints[:i], testpoints[i+1:]...)
			return true
		}
	}
	return false
}
func (n Node) checkPoint() {
	for _, p := range n.points {
		removeFromTestPoints(p)
	}
}

func TestPointers(t *testing.T) {
	assert.Equal(t, testTree.GetQuadrant(), "ROOT")

	for z := testTree.next; z != nil; z = z.next {
		z.checkPoint()

		assert.False(t, z.alreadySeen(), "Should see every node only once")
	}
	assert.Empty(t, testpoints, "All created testpoints should have been found")
}

func TestPoint7(t *testing.T) {

	assert.Equal(t, testTree.GetQuadrant(), "ROOT")

	point1 := geo.NewPoint("1", "1")
	point2 := geo.NewPoint("11", "10")
	point3 := geo.NewPoint("110", "100")

	node1 := testTree.Find(point1)
	assert.Equal(t, node1.getHash(), "10", "node1 hash")

	node2 := testTree.Find(point2)
	assert.Equal(t, node2.getHash(), "1011", "node2 hash")

	node3 := testTree.Find(point3)
	assert.Equal(t, node3.getHash(), "101110", "node3 hash")

	assert.Equal(t, node1.parent.orient, orientU, "node1 orient")
	assert.Equal(t, node2.parent.orient, orientU, "node2 orient")
	assert.Equal(t, node3.parent.orient, orientL, "node3 orient")

	assert.True(t, node3.ContainsPoint(point3))

}
