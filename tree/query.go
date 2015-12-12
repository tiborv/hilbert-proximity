package tree

import (
	"log"

	"github.com/tiborv/hilbert-proximity/geo"
)

//RangeQuery returns all the Points withing a range
func (n *Node) RangeQuery(min geo.Point, max geo.Point) []geo.Point {
	if n == nil {
		log.Fatal("Query: N is nil!")
	}
	nodeQueue := []*Node{n} //Search queue
	foundPoints := []geo.Point{}
	curr := nodeQueue[0]
	for len(nodeQueue) > 0 {
		//Assume all Points from found nodes are within the query
		foundPoints = append(foundPoints, curr.points...)

		if curr.splitted {
			for _, c := range curr.children {
				if c.isWithinRange(min, max) { //Find all children within query
					nodeQueue = append(nodeQueue, c) //Append to serach que
				}
			}
		}
		curr = nodeQueue[len(nodeQueue)-1]
		nodeQueue = nodeQueue[:len(nodeQueue)-1]
	}
	nodeQueue = nil //gc
	return foundPoints
}

const zero = int64(0)

func (n Node) isWithinRange(min geo.Point, max geo.Point) bool {
	return n.zx.CompareTo(min.GetX()) >= zero &&
		n.zy.CompareTo(min.GetY()) >= zero &&
		n.zx.CompareTo(max.GetX()) <= zero &&
		n.zy.CompareTo(max.GetY()) <= zero
}
