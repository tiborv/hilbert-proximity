package tree

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiborv/hilbert-gis/geo"
)

var prTree Node
var testpoints []geo.Point

func init() {
	prTree = NewTree(1, false)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			x := fmt.Sprintf("%01s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%01s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			prTree.InsertPoint(g)
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			x := fmt.Sprintf("%02s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%02s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			prTree.InsertPoint(g)
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			x := fmt.Sprintf("%03s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%03s", strconv.FormatInt(int64(j), 2))

			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			prTree.InsertPoint(g)
		}
	}

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			x := fmt.Sprintf("%04s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%04s", strconv.FormatInt(int64(j), 2))

			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			prTree.InsertPoint(g)
		}
	}
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			x := fmt.Sprintf("%05s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%05s", strconv.FormatInt(int64(j), 2))

			g := geo.NewPoint(x, y)
			testpoints = append(testpoints, g)
			prTree.InsertPoint(g)
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
	assert.Equal(t, prTree.GetSector(), "ROOT")

	for z := prTree.right; z != nil; z = z.right {
		z.checkPoint()
		assert.False(t, z.alreadySeen(), "Should see every node only once")
	}
	assert.Empty(t, testpoints, "All created testpoints should have been found")
}

func TestPoint1(t *testing.T) {

	assert.Equal(t, prTree.GetSector(), "ROOT")

	point1 := geo.NewPoint("1", "1")
	point2 := geo.NewPoint("11", "10")
	point3 := geo.NewPoint("110", "100")

	node1 := prTree.Find(point1)
	assert.Equal(t, node1.GetHash(), "10", "node1 hash")

	node2 := prTree.Find(point2)
	assert.Equal(t, node2.GetHash(), "1110", "node2 hash")

	node3 := prTree.Find(point3)
	assert.Equal(t, node3.GetHash(), "101110", "node3 hash")

	assert.Equal(t, node1.parent.orient, "U", "node1 orient")
	assert.Equal(t, node2.parent.orient, "U", "node2 orient")
	assert.Equal(t, node3.parent.orient, "L", "node3 orient")

	assert.True(t, node3.ContainsPoint(point3))

}

func TestPoint2(t *testing.T) {

	assert.Equal(t, prTree.GetSector(), "ROOT")

	point1 := geo.NewPoint("1", "0")
	point2 := geo.NewPoint("11", "01")

	point3 := geo.NewPoint("110", "010")
	point4 := geo.NewPoint("1101", "0100")
	point5 := geo.NewPoint("1100", "0111")

	node1 := prTree.Find(point1)
	assert.Equal(t, node1.GetHash(), "11", "node1 hash")

	node2 := prTree.Find(point2)
	assert.Equal(t, node2.GetHash(), "0011", "node2 hash")

	node3 := prTree.Find(point3)
	assert.Equal(t, node3.GetHash(), "100011", "node3 hash")

	node4 := prTree.Find(point4)
	assert.Equal(t, node4.GetHash(), "01100011", "node4 hash")

	node5 := prTree.Find(point5)
	assert.Equal(t, node5.GetHash(), "11110011", "node5 hash")

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

	node1 := prTree.Find(point1)
	assert.Equal(t, node1.GetHash(), "11", "node1 hash")

	node2 := prTree.Find(point2)
	assert.Equal(t, node2.GetHash(), "1011", "node2 hash")

	node3 := prTree.Find(point3)
	assert.Equal(t, node3.GetHash(), "101011", "node3 hash")

	node4 := prTree.Find(point4)
	assert.Equal(t, node4.GetHash(), "01101011", "node4 hash")

	node5 := prTree.Find(point5)
	assert.Equal(t, node5.GetHash(), "01111011", "node5 hash")

	node6 := prTree.Find(point6)
	assert.Equal(t, node6.GetHash(), "011111", "node6 hash")

	assert.Equal(t, node1.parent.orient, "U", "node1 orient")
	assert.Equal(t, node2.parent.orient, "L", "node2 orient")
	assert.Equal(t, node3.parent.orient, "L", "node3 orient")
	assert.Equal(t, node4.parent.orient, "L", "node4 orient")
	assert.Equal(t, node5.parent.orient, "U", "node5 orient")
	assert.Equal(t, node6.parent.orient, "U", "node6 orient")

	assert.True(t, node4.ContainsPoint(point4))

}
