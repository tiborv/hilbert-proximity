package tree

import (
	"log"
	"strconv"

	"github.com/tiborv/hilbert-gis/geo"
)

const zero = int64(0)

func (n *Node) rangeQuery(min geo.Point, max geo.Point) []geo.Point {
	if n == nil {
		log.Fatal("Query: N is nil!")
	}
	nodeQueue := []*Node{n}
	foundPoints := []geo.Point{}
	curr := nodeQueue[0]
	for len(nodeQueue) > 0 {
		foundPoints = append(foundPoints, curr.points...)

		if curr.splitted {
			for _, p := range curr.children {
				if p.isWithinRange(min, max) {
					nodeQueue = append(nodeQueue, p)
				}
			}
		}

		curr = nodeQueue[len(nodeQueue)-1]
		nodeQueue = nodeQueue[:len(nodeQueue)-1]

	}
	nodeQueue = nil
	return foundPoints
}

func compareValues(val1 string, val2 string) int64 {
	var v1, v2 int64
	if len(val1) == len(val2) {
		v1, _ = strconv.ParseInt(val1, 2, 64)
		v2, _ = strconv.ParseInt(val2, 2, 64)
	} else if len(val1) > len(val2) {
		v1, _ = strconv.ParseInt(val1[:len(val2)], 2, 64)
		v2, _ = strconv.ParseInt(val2, 2, 64)
	} else {
		v1, _ = strconv.ParseInt(val1, 2, 64)
		v2, _ = strconv.ParseInt(val2[:len(val1)], 2, 64)
	}

	return v1 - v2

}

func (n Node) isWithinRange(min geo.Point, max geo.Point) bool {
	return compareValues(n.zx, min.GetX()) >= zero &&
		compareValues(n.zy, min.GetY()) >= zero &&
		compareValues(n.zx, max.GetX()) <= zero &&
		compareValues(n.zy, max.GetY()) <= zero
}
