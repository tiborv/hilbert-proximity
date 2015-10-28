package tree

import (
	"bytes"
	"log"

	"github.com/tiborv/hilbert-gis/geo"
)

//NewTree creates a rootnode
func NewTree(maxPointsEachNode int, matrix bool) Node {
	maxPoints = maxPointsEachNode
	mx = matrix
	pointsInserted = 0
	return *newNode(nil, "U", "")
}

//GetSector returns the sector for a given node
func (n Node) GetSector() string {
	if n.parent == nil {
		return "ROOT"
	}
	return n.sector
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

//GetHash returns the hash for the node
func (n *Node) GetHash() string {
	if n.parent == nil {
		return "ROOT"
	}
	var buffer bytes.Buffer
	for n.parent != nil {
		buffer.WriteString(n.GetSector())
		n = n.getParent()
	}
	return buffer.String()

}

func (n *Node) leftLeaf() *Node {
	for n.splitted {
		n = n.children[0]
	}
	return n
}
