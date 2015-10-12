package tree

import (
	"log"

	"github.com/tiborv/hilbert-gis/geo"
)

//Node struct
type Node struct {
	sector   string
	points   []geo.Point
	parent   *Node
	children []*Node
	orient   string
	mapper   func(value string, nodes []*Node) *Node
}

func (n *Node) newNode(orient string, sector string, mapper func(value string, nodes []*Node) *Node) Node {
	return Node{
		sector:   sector,
		parent:   n,
		children: make([]*Node, 4),
		orient:   orient,
		mapper:   mapper,
	}

}

//NewTree creates a rootnode
func NewTree() Node {
	return Node{
		children: make([]*Node, 4),
		orient:   "U",
		mapper:   mapU,
	}
}

//GetSector returns the sector for a geiven node
func (n Node) GetSector() string {
	if n.parent == nil {
		return "ROOT"
	}
	return n.sector
}

func (n *Node) addPoint(p geo.Point, maxPoints int) bool {
	if n == nil {
		log.Fatal("addPoint n=nil!")
	}
	if len(n.points) < maxPoints {
		n.points = append(n.points, p)
		return true
	}
	return false

}
func (n Node) next(value string) *Node {
	if n.children[0] == nil {
		n.split()
	}
	if n.children[0] == nil {
		log.Fatal("next children is nil after spilt!")
	}
	b := n.mapper(value, n.children)
	if b == nil {
		log.Fatal("next b=nil!")
	}
	return b
}

//Split a node and re-distribute points.
func (n *Node) split() *Node {

	switch n.orient {
	case "U":
		c1 := n.newNode("R", "00", mapR)
		n.children[0] = &c1
		c2 := n.newNode("U", "01", mapU)
		n.children[1] = &c2
		c3 := n.newNode("U", "10", mapU)
		n.children[2] = &c3
		c4 := n.newNode("L", "11", mapL)
		n.children[3] = &c4
	case "D":
		c1 := n.newNode("D", "00", mapD)
		n.children[0] = &c1
		c2 := n.newNode("R", "01", mapR)
		n.children[1] = &c2
		c3 := n.newNode("L", "10", mapL)
		n.children[2] = &c3
		c4 := n.newNode("D", "11", mapD)
		n.children[3] = &c4
	case "L":
		c1 := n.newNode("L", "00", mapL)
		n.children[0] = &c1
		c2 := n.newNode("L", "01", mapL)
		n.children[1] = &c2
		c3 := n.newNode("D", "10", mapD)
		n.children[2] = &c3
		c4 := n.newNode("U", "11", mapU)
		n.children[3] = &c4
	case "R":
		c1 := n.newNode("U", "00", mapU)
		n.children[0] = &c1
		c2 := n.newNode("D", "01", mapD)
		n.children[1] = &c2
		c3 := n.newNode("R", "10", mapR)
		n.children[2] = &c3
		c4 := n.newNode("R", "11", mapR)
		n.children[3] = &c4
	}
	//Re-insert points from the root-node
	if len(n.points) > 0 {
		root := n.getRoot()
		points := n.points
		for _, p := range points {
			root.InsertPoint(p) //Insert point into the tree
		}
		n.points = []geo.Point{}
	}
	return n
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
