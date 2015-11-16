package tree

import (
	"fmt"
	"log"

	"github.com/tiborv/hilbert-gis/geo"
)

//Node struct
type Node struct {
	quadrant byte         // Hilbert quadrant
	orient   rune         // The Hilbert orientation of the current node (Used for confirimation within test)
	mapper   map[byte]int // Function giving the Hilbert mapped next children node
	points   []geo.Point  // Points within a node
	parent   *Node        // Pointer to parrent node
	children []*Node      // Pointers to children nodes
	next     *Node        // A pointer to the next node
	adj      *Node        // A pointer to the adjecent node (Next child of same parent)
	splitted bool         // Boolean indicating if this node is splitted (has children)
	level    int          //
	zpoint   geo.Point
}

var maxPoints, pointsInserted int

const orientU, orientD, orientL, orientR = 'u', 'd', 'l', 'r'

//Creates a new node
func newNode(parent *Node, orient rune, quadrant byte, zvalue byte) *Node {

	newNode := Node{
		quadrant: quadrant,
		parent:   parent,
		children: make([]*Node, 4),
		orient:   orient,
		splitted: false,
	}
	if parent == nil { //If root
		newNode.level = 0
		newNode.zpoint = geo.NewPoint("", "")

	} else {
		newNode.level = parent.level + 1
		newNode.zpoint = parent.zpoint.Append(zvalue)
	}
	switch orient {
	case orientU:
		newNode.mapper = mapU
	case orientD:
		newNode.mapper = mapD
	case orientL:
		newNode.mapper = mapL
	case orientR:
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
func (n Node) descend(b byte) *Node {
	if !n.splitted {
		log.Fatal("No children to explore!")
	}
	return n.children[n.mapper[b]]
}

//Split a node and re-distribute points.

func (n *Node) split() *Node {
	if n.splitted {
		return n
	}
	switch n.orient {
	case orientU:
		n.children[0] = newNode(n, orientR, 0, 0)
		n.children[1] = newNode(n, orientU, 1, 1)
		n.children[2] = newNode(n, orientU, 2, 3)
		n.children[3] = newNode(n, orientL, 3, 2)
	case orientD:
		n.children[0] = newNode(n, orientL, 0, 1)
		n.children[1] = newNode(n, orientD, 1, 0)
		n.children[2] = newNode(n, orientD, 2, 2)
		n.children[3] = newNode(n, orientR, 3, 3)
	case orientL:
		n.children[0] = newNode(n, orientD, 0, 3)
		n.children[1] = newNode(n, orientL, 1, 1)
		n.children[2] = newNode(n, orientL, 2, 0)
		n.children[3] = newNode(n, orientU, 3, 1)
	case orientR:
		n.children[0] = newNode(n, orientU, 0, 1)
		n.children[1] = newNode(n, orientR, 1, 3)
		n.children[2] = newNode(n, orientR, 2, 2)
		n.children[3] = newNode(n, orientD, 3, 0)
	}

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
		n.points = []geo.Point{}
		for _, p := range tempPoints {
			if p.GetBitDepth() > n.level { // Only re-insert points that can be added deeper in the tree
				n.insert(p, n.level) //Insert point into the tree
				pointsInserted--     //Dont count points that need to be re-inserted
			} else { // Keep points that cant go deeper into the tree, in the same node
				n.points = append(n.points, p)
			}
		}
		tempPoints = nil
	}
	return n
}

//Inserts a point into the first availbe children node
func (n *Node) insert(point geo.Point, pos int) *Node {
	concat, last := point.GetConcatAt(pos)

	n = n.split().descend(concat) //.split() only splits if not splitted
	if n.addPoint(point) {
		return n
	}
	if last {
		fmt.Println(n.parent.children[n.mapper[1]])
		fmt.Println(n.parent.children[n.mapper[2]])

		fmt.Println(n.points, n.treePos(), n.level)
		log.Fatal("Couldnt insert, not enough point", point)
	}

	return n.insert(point, pos+1)
}

func (n *Node) find(point geo.Point, pos int) *Node {
	concat, last := point.GetConcatAt(pos)
	if n.descend(concat).ContainsPoint(point) {
		return n.descend(concat)
	}

	if last {
		return nil //Not found
	}

	return n.descend(concat).find(point, pos+1)

}

func (n *Node) getRoot() *Node {
	for n.parent != nil {
		n = n.parent
	}
	return n
}

func (n Node) treePos() []int {
	var arr []int
	for n.parent != nil {
		for i, c := range n.parent.children {
			if c.zpoint.Equals(n.zpoint) {
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
