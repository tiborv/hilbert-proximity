package tree

import (
	"bytes"
	"fmt"
	"log"

	"github.com/tiborv/hilbert-proximity/const"
	"github.com/tiborv/hilbert-proximity/geo"
)

//NewTree creates a rootnode
func NewTree(maxPointsEachNode int) *Node {
	maxPoints = maxPointsEachNode
	pointsInserted = 0
	return newNode(nil, c.U, "", "")
}

//GetQuadrant returns the sector for a given node
func (n Node) GetQuadrant() string {
	if n.parent == nil {
		return ""
	}
	return n.quadrant
}

//ContainsPoint returns true if a point is in the given node
func (n Node) ContainsPoint(p geo.Point) bool {
	for _, v := range n.points {
		if p.Equals(v) {
			return true
		}
	}
	return false
}

//Insert inserts a point into the first availbe descendand of a node
func (n *Node) Insert(point geo.Point) *Node {
	concat, last := point.GetMortonAt(n.level)
	n = n.split().descend(concat) //.split() only splits if node not splitted
	if n.addPoint(point) {
		return n
	}
	if last {
		fmt.Println(n.points, n.treePos(), n.level)
		log.Fatal("Couldnt insert, not enough point", point)
	}

	return n.Insert(point)
}

//Find Searches the tree for a given node that contains a point
func (n *Node) Find(point geo.Point) *Node {
	concat, last := point.GetMortonAt(n.level)
	if n.descend(concat).ContainsPoint(point) {
		return n.descend(concat)
	}
	if last { //Point not found
		return nil
	}
	return n.descend(concat).Find(point)
}

//GetHash returns the hash for a given node
func (n *Node) GetHash() string {
	if n.parent == nil {
		return "ROOT"
	}
	var buffer bytes.Buffer
	for n.parent != nil {
		buffer.WriteString(n.GetQuadrant())
		n = n.getParent()
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
