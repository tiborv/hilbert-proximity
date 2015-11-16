package tree

import (
	"bytes"
	"log"

	"github.com/tiborv/hilbert-gis/geo"
)

//NewTree creates a rootnode
func NewTree(maxPointsEachNode int) *Node {
	maxPoints = maxPointsEachNode
	pointsInserted = 0

	return &Node{orient: orientU, children: make([]*Node, 4), splitted: false}
}

//GetQuadrant returns the sector for a given node
func (n Node) GetQuadrant() byte {
	return n.quadrant
}

//InsertPoint Inserts a point into the first availbe node from root and down
func (n *Node) InsertPoint(point geo.Point) *Node {
	if n == nil {
		log.Fatal("InsertPoint n==nil!")
	}
	return n.insert(point, 0)
}

//ContainsPoint returns true if a point is in the node
func (n Node) ContainsPoint(p geo.Point) bool {
	for _, v := range n.points {
		if p.Equals(v) {
			return true
		}
	}
	return false
}

//Find Searches the tree for the node that contains a point
func (n *Node) Find(point geo.Point) *Node {
	return n.find(point, 0)
}

//GetHash returns the hash for the node (this method is only used for testing)
func (n *Node) getHash() string {
	if n.parent == nil {
		return "ROOT"
	}
	var buffer bytes.Buffer
	for n.parent != nil {
		buffer.WriteByte(n.GetQuadrant())
		n = n.parent
	}
	return buffer.String()

}
func (n *Node) rigthLeaf() *Node {
	for n.splitted {
		n = n.children[3]
	}
	return n
}

func (n *Node) leftLeaf() *Node {
	for n.splitted {
		n = n.children[0]
	}
	return n
}
