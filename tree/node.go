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
	right    *Node
	splitted bool
	level    int
}

var mx bool
var maxPoints, pointsInserted int

func newNode(parent *Node, orient string, sector string) *Node {
	newNode := Node{
		sector:   sector,
		parent:   parent,
		children: make([]*Node, 4),
		orient:   orient,
		splitted: false,
	}
	if parent == nil {
		newNode.level = 0
	} else {
		newNode.level = parent.level + 1
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
func (n Node) next(value string) *Node {
	if !n.splitted {
		log.Fatal("No children to explore!")
	}
	b := n.mapper(value, n.children)
	if b == nil {
		log.Fatal("next b=nil!")
	}
	return b
}

//Split a node and re-distribute points.
func (n *Node) split() *Node {
	if n.splitted {
		return n
	}
	switch n.orient {
	case "U":
		n.children[0] = newNode(n, "R", "00")
		n.children[1] = newNode(n, "U", "01")
		n.children[2] = newNode(n, "L", "11")
		n.children[3] = newNode(n, "U", "10")
	case "D":
		n.children[0] = newNode(n, "R", "11")
		n.children[1] = newNode(n, "D", "10")
		n.children[2] = newNode(n, "L", "00")
		n.children[3] = newNode(n, "D", "01")
	case "L":
		n.children[0] = newNode(n, "U", "11")
		n.children[1] = newNode(n, "L", "01")
		n.children[2] = newNode(n, "D", "00")
		n.children[3] = newNode(n, "L", "10")
	case "R":
		n.children[0] = newNode(n, "U", "00")
		n.children[1] = newNode(n, "R", "10")
		n.children[2] = newNode(n, "D", "11")
		n.children[3] = newNode(n, "R", "01")
	}
	iter := n
	for iter.right != nil && iter.parent != nil {
		if mx {
			n.children[3].right = iter.right.leftLeaf()
		} else {
			n.children[3].right = iter.right
			break
		}
		iter = iter.parent
	}

	if iter.parent == nil {
		n.children[3].right = nil
	}

	n.children[0].right = n.children[1]
	n.children[1].right = n.children[2]
	n.children[2].right = n.children[3]
	n.right = n.children[0]
	n.splitted = true

	//	Re-insert points from the root-node
	if !mx {
		if len(n.points) > 0 {
			root := n.getRoot()
			tempPoints := n.points
			n.points = []geo.Point{}
			for _, p := range tempPoints {
				root.InsertPoint(p) //Insert point into the tree
				pointsInserted--    //Dont count points that need to be re-inserted
			}
			tempPoints = nil
		}
	}
	return n
}

func (n *Node) insert(point geo.Point, pos int) *Node {
	concat, last := point.GetConcatAt(pos)
	n = n.split().next(concat) //.split() only splits if not splitted if not splitted
	if mx {                    //If MX-tree
		if last { //Only add if on last bit of point
			if !n.addPoint(point) {
				fmt.Println("Couldnt insert point:", point)
			}
			return n
		}
	} else {
		if n.addPoint(point) {
			return n
		}
		if last {
			fmt.Println(n.points, n.treePos(), n.level)
			log.Fatal("Couldnt insert, not enough point", point)
		}
	}

	return n.insert(point, pos+1)
}

func (n *Node) find(point geo.Point, pos int) *Node {
	concat, last := point.GetConcatAt(pos)
	if n.next(concat).ContainsPoint(point) {
		return n.next(concat)
	}

	if last {
		return nil
	}

	return n.next(concat).find(point, pos+1)

}

func (n *Node) getRoot() *Node {
	for n.parent != nil {
		n = n.getParent()
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
			if c.GetHash() == n.GetHash() {
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
