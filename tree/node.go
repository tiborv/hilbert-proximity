package tree

import "github.com/tiborv/hilbert-gis/geo"

//Node struct
type Node struct {
	mapper   func(value string, nodes []*Node) *Node
	sector   string
	points   []geo.Point
	parent   *Node
	children []*Node
}

func newNode(parent *Node, mapper func(value string, nodes []*Node) *Node, sector string) Node {
	return Node{
		sector:   sector,
		parent:   parent,
		mapper:   mapper,
		children: make([]*Node, 4),
	}
}

//NewTree creates a rootnode
func NewTree() Node {
	return Node{
		mapper:   map2,
		children: make([]*Node, 4),
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
	b := n.mapper(value, n.children)
	return b
}

//Split a node and re-distribute points.
func (n *Node) split() *Node {
	child := newNode(n, map1, "00")
	n.children[0] = &child

	child2 := newNode(n, map2, "01")
	n.children[1] = &child2

	child3 := newNode(n, map3, "10")
	n.children[2] = &child3

	child4 := newNode(n, map4, "11")
	n.children[3] = &child4

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
