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
	prTree = NewTree(5, false)
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

func TestPoints(t *testing.T) {
	assert.Equal(t, prTree.GetSector(), "ROOT")

	for z := prTree.right; z != nil; z = z.right {
		z.checkPoint()

		assert.False(t, z.alreadySeen(), "Should see every node only once")
	}
	assert.Empty(t, testpoints, "All created testpoints should have been found")
}
