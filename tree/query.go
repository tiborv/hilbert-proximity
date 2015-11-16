package tree

import (
	"log"

	"github.com/tiborv/hilbert-gis/geo"
)

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
			for _, c := range curr.children {
				if c.zpoint.IsWithinRange(min, max) {
					nodeQueue = append(nodeQueue, c)
				}
			}
		}

		curr = nodeQueue[len(nodeQueue)-1]
		nodeQueue = nodeQueue[:len(nodeQueue)-1]

	}
	nodeQueue = nil
	return foundPoints
}
