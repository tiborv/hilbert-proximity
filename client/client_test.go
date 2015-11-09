package client

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/tiborv/hilbert-gis/geo"
	"github.com/tiborv/hilbert-gis/tree"
)

var prTree tree.Node

func inti() {
	prTree = tree.NewTree(1, false)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			x := fmt.Sprintf("%01s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%01s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			prTree.InsertPoint(g)
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			x := fmt.Sprintf("%02s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%02s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			prTree.InsertPoint(g)
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			x := fmt.Sprintf("%03s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%03s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			prTree.InsertPoint(g)
		}
	}
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			x := fmt.Sprintf("%04s", strconv.FormatInt(int64(i), 2))
			y := fmt.Sprintf("%04s", strconv.FormatInt(int64(j), 2))
			g := geo.NewPoint(x, y)
			prTree.InsertPoint(g)
		}
	}
}

func Test_Client(t *testing.T) {
	fmt.Println()

}
