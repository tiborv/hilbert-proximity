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

func init() {
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
	fmt.Println("Points inserted:", pointsInserted)

}

var seen []Node

func (n Node) alreadySeen() bool {
	for _, p := range seen {
		if n.GetHash() == p.GetHash() {

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
	assert.Equal(t, testTree.GetSector(), "ROOT")

	for z := testTree.next; z != nil; z = z.next {
		z.checkPoint()

		assert.False(t, z.alreadySeen(), "Should see every node only once")
	}
	assert.Empty(t, testpoints, "All created testpoints should have been found")
}

func TestPoint7(t *testing.T) {

	assert.Equal(t, testTree.GetSector(), "ROOT")

	point1 := geo.NewPoint("1", "1")
	point2 := geo.NewPoint("11", "10")
	point3 := geo.NewPoint("110", "100")

	node1 := testTree.Find(point1)
	assert.Equal(t, node1.hash, "10", "node1 hash")

	node2 := testTree.Find(point2)
	assert.Equal(t, node2.hash, "1011", "node2 hash")

	node3 := testTree.Find(point3)
	assert.Equal(t, node3.hash, "101110", "node3 hash")

	assert.Equal(t, node1.parent.orient, "U", "node1 orient")
	assert.Equal(t, node2.parent.orient, "U", "node2 orient")
	assert.Equal(t, node3.parent.orient, "L", "node3 orient")

	assert.True(t, node3.ContainsPoint(point3))

}

func TestPoint8(t *testing.T) {

	assert.Equal(t, testTree.GetSector(), "ROOT")

	point1 := geo.NewPoint("1", "0")
	point2 := geo.NewPoint("11", "01")

	point3 := geo.NewPoint("110", "010")
	point4 := geo.NewPoint("1101", "0100")
	point5 := geo.NewPoint("1100", "0111")

	node1 := testTree.Find(point1)
	assert.Equal(t, node1.hash, "11", "node1 hash")

	node2 := testTree.Find(point2)
	assert.Equal(t, node2.hash, "1100", "node2 hash")

	node3 := testTree.Find(point3)
	assert.Equal(t, node3.hash, "110010", "node3 hash")

	node4 := testTree.Find(point4)
	assert.Equal(t, node4.hash, "11001001", "node4 hash")

	node5 := testTree.Find(point5)
	assert.Equal(t, node5.hash, "11001111", "node5 hash")

	assert.Equal(t, node1.parent.orient, "U", "node1 orient")
	assert.Equal(t, node2.parent.orient, "L", "node2 orient")
	assert.Equal(t, node3.parent.orient, "D", "node3 orient")
	assert.Equal(t, node4.parent.orient, "D", "node4 orient")
	assert.Equal(t, node5.parent.orient, "R", "node5 orient")

	assert.True(t, node4.ContainsPoint(point4))

}

func TestPoint3(t *testing.T) {

	point1 := geo.NewPoint("1", "0")
	point2 := geo.NewPoint("10", "00")
	point3 := geo.NewPoint("100", "000")
	point4 := geo.NewPoint("1000", "0001")
	point5 := geo.NewPoint("1010", "0001")
	point6 := geo.NewPoint("110", "001")

	node1 := testTree.Find(point1)
	assert.Equal(t, node1.hash, "11", "node1 hash")

	node2 := testTree.Find(point2)
	assert.Equal(t, node2.hash, "1110", "node2 hash")

	node3 := testTree.Find(point3)
	assert.Equal(t, node3.hash, "111010", "node3 hash")

	node4 := testTree.Find(point4)
	assert.Equal(t, node4.hash, "11101001", "node4 hash")

	node5 := testTree.Find(point5)
	assert.Equal(t, node5.hash, "11101101", "node5 hash")

	node6 := testTree.Find(point6)
	assert.Equal(t, node6.hash, "111101", "node6 hash")

	assert.Equal(t, node1.parent.orient, "U", "node1 orient")
	assert.Equal(t, node2.parent.orient, "L", "node2 orient")
	assert.Equal(t, node3.parent.orient, "L", "node3 orient")
	assert.Equal(t, node4.parent.orient, "L", "node4 orient")
	assert.Equal(t, node5.parent.orient, "U", "node5 orient")
	assert.Equal(t, node6.parent.orient, "U", "node6 orient")

	assert.True(t, node4.ContainsPoint(point4))

}
