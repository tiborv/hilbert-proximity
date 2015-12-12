package tree

import (
	"log"

	"github.com/tiborv/go-bitarray"
	"github.com/tiborv/hilbert-proximity/const"
	"github.com/tiborv/hilbert-proximity/geo"
)

//Node struct
type Node struct {
	quadrant string         //Hilbert Quadrant
	points   []geo.Point    //Point bucket
	parent   *Node          //Parent node pointer
	children []*Node        //Array of pointers to children nodes
	orient   rune           //Hilbert orientation
	mapper   map[string]int //Hilber mappinng function
	next     *Node          //Next node in hilbert-order
	adj      *Node          //Next sibling-node (adjacent in the tree)
	splitted bool           //Has children
	level    int            //Tree level
	hash     string         //Hilbert hash
	zx       ba.BitArray    //Z-order x coordinate
	zy       ba.BitArray    //Z-order y coordinate
}

var maxPoints, pointsInserted int

func newNode(parent *Node, orient rune, quadrant string, zvalue string) *Node {
	newNode := Node{
		quadrant: quadrant,
		parent:   parent,
		children: make([]*Node, 4),
		orient:   orient,
		splitted: false,
	}
	if parent == nil { //root
		newNode.level = 0
		newNode.hash = ""
		newNode.zx = ba.NewBitArray("")
		newNode.zy = ba.NewBitArray("")

	} else {
		newNode.level = parent.level + 1
		newNode.hash = parent.hash + quadrant //Set the Hilbert-hash of this node
		newNode.zx = ba.NewBitArray(parent.zx, zvalue[:1])
		newNode.zy = ba.NewBitArray(parent.zy, zvalue[1:])
	}

	switch orient { //Set the Hilbert-mapper
	case c.U:
		newNode.mapper = mapU
	case c.D:
		newNode.mapper = mapD
	case c.L:
		newNode.mapper = mapL
	case c.R:
		newNode.mapper = mapR
	}
	return &newNode

}

func (n *Node) addPoint(p geo.Point) bool {
	if n == nil {
		log.Fatal("addPoint n=nil!")
	}
	if len(n.points) < maxPoints {
		n.points = append(n.points, p)
		pointsInserted++
		return true
	}
	return false

}

func (n Node) descend(value string) *Node {
	if !n.splitted {
		log.Fatal("No children to explore!")
	}
	return n.children[n.mapper[value]]
}

//Split a node and re-distribute points.
func (n *Node) split() *Node {
	if n.splitted {
		return n
	}
	switch n.orient {
	case c.U:
		n.children[0] = newNode(n, c.R, "00", "00")
		n.children[1] = newNode(n, c.U, "01", "01")
		n.children[2] = newNode(n, c.U, "10", "11")
		n.children[3] = newNode(n, c.L, "11", "10")
	case c.D:
		n.children[0] = newNode(n, c.L, "00", "01")
		n.children[1] = newNode(n, c.D, "01", "00")
		n.children[2] = newNode(n, c.D, "10", "10")
		n.children[3] = newNode(n, c.R, "11", "11")
	case c.L:
		n.children[0] = newNode(n, c.D, "00", "11")
		n.children[1] = newNode(n, c.L, "01", "01")
		n.children[2] = newNode(n, c.L, "10", "00")
		n.children[3] = newNode(n, c.U, "11", "10")
	case c.R:
		n.children[0] = newNode(n, c.U, "00", "01")
		n.children[1] = newNode(n, c.R, "01", "11")
		n.children[2] = newNode(n, c.R, "10", "10")
		n.children[3] = newNode(n, c.D, "11", "00")
	}
	//Fix next Hilbert node for last child
	for iter := n; iter.parent != nil; iter = iter.parent {
		if iter.adj != nil {
			n.children[3].next = iter.adj
			break
		}
		if iter.parent == nil {
			n.children[3].next = nil
			break
		}
	}

	n.children[0].next = n.children[1]
	n.children[0].adj = n.children[1]
	n.children[1].next = n.children[2]
	n.children[1].adj = n.children[2]
	n.children[2].next = n.children[3]
	n.children[2].adj = n.children[3]
	n.next = n.children[0]
	n.splitted = true

	//	Check if points needs to be reinserted
	if maxPoints == len(n.points) {
		tempPoints := n.points
		n.points = []geo.Point{}       //Remove all points from the node
		for _, p := range tempPoints { //Insert all points
			if p.GetBitDepth() > n.level { //Check if a point can go deeper
				n.Insert(p)      //Insert
				pointsInserted-- //Dont re-count points that need to be re-inserted
			} else { // Keep points that cant go deeper in the same node
				n.points = append(n.points, p)
			}
		}
		tempPoints = nil
	}
	return n
}

func (n *Node) getRoot() *Node {
	for n.parent != nil {
		n = n.parent
	}
	return n
}

func (n *Node) getParent() *Node {
	return n.parent
}

func (n Node) treePos() []int {
	var arr []int
	for n.parent != nil {
		for i, c := range n.parent.children {
			if c.hash == n.hash {
				arr = append(arr, i)
				break
			}
		}
		n = *n.parent
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
