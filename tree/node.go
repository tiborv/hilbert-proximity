package tree

import (
	"fmt"
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
	next     *Node
	adj      *Node
	splitted bool
	level    int
	hash     string
	zx       string
	zy       string
}

var maxPoints, pointsInserted int

func newNode(parent *Node, orient string, sector string, zvalue string) *Node {
	newNode := Node{
		sector:   sector,
		parent:   parent,
		children: make([]*Node, 4),
		orient:   orient,
		splitted: false,
	}
	if parent == nil { //root
		newNode.level = 0
		newNode.hash = ""
		newNode.zx = ""
		newNode.zy = ""

	} else {
		newNode.level = parent.level + 1
		newNode.hash = parent.hash + sector
		newNode.zx += parent.zx + zvalue[:1]
		newNode.zy += parent.zy + zvalue[1:]

	}
	switch orient {
	case "U":
		newNode.mapper = mapU
	case "D":
		newNode.mapper = mapD
	case "L":
		newNode.mapper = mapL
	case "R":
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
func (n Node) desend(value string) *Node {
	if !n.splitted {
		log.Fatal("No children to explore!")
	}
	return n.mapper(value, n.children)
}

//Split a node and re-distribute points.
func (n *Node) split() *Node {
	if n.splitted {
		return n
	}
	switch n.orient {
	case "U":
		n.children[0] = newNode(n, "R", "00", "00")
		n.children[1] = newNode(n, "U", "01", "01")
		n.children[2] = newNode(n, "U", "10", "11")
		n.children[3] = newNode(n, "L", "11", "10")
	case "D":
		n.children[0] = newNode(n, "L", "00", "01")
		n.children[1] = newNode(n, "D", "01", "00")
		n.children[2] = newNode(n, "D", "10", "10")
		n.children[3] = newNode(n, "R", "11", "11")
	case "L":
		n.children[0] = newNode(n, "D", "00", "11")
		n.children[1] = newNode(n, "L", "01", "01")
		n.children[2] = newNode(n, "L", "10", "00")
		n.children[3] = newNode(n, "U", "11", "10")
	case "R":
		n.children[0] = newNode(n, "U", "00", "01")
		n.children[1] = newNode(n, "R", "01", "11")
		n.children[2] = newNode(n, "R", "10", "10")
		n.children[3] = newNode(n, "D", "11", "00")
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
	n = n.split().desend(concat) //.split() only splits if not splitted if not splitted
	if n.addPoint(point) {
		return n
	}
	if last {
		fmt.Println(n.points, n.treePos(), n.level)
		log.Fatal("Couldnt insert, not enough point", point)
	}

	return n.insert(point, pos+1)
}

func (n *Node) find(point geo.Point, pos int) *Node {
	concat, last := point.GetConcatAt(pos)
	if n.desend(concat).ContainsPoint(point) {
		return n.desend(concat)
	}

	if last {
		return nil
	}

	return n.desend(concat).find(point, pos+1)

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
