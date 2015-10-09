package tree

import (
	"bytes"

	"github.com/tiborv/hilbert-gis/geo"
)

//InsertPoint Inserts a point into the first availbe node from root and down
func (n *Node) InsertPoint(point geo.Point) *Node {
	return n.insert(point, 0)
}

func (n *Node) insert(point geo.Point, pos int) *Node {

	concat, last := point.GetConcatAt(pos)
	n = n.next(concat)
	if last {
		n.addPoint(point, 1)
		return n
	}
	return n.insert(point, pos+1)
}

//Find Searches the tree for the node that contains a point
func (n *Node) Find(point geo.Point) *Node {
	return n.find(point, 0)
}

func (n *Node) find(point geo.Point, pos int) *Node {
	concat, last := point.GetConcatAt(pos)
	if n.ContainsPoint(point) || last {
		return n.next(concat)
	}
	return n.next(concat).find(point, pos+1)

}

func (n *Node) getRoot() *Node {
	for n.parent != nil {
		n = n.prev()
	}
	return n
}

//GetHash returns the hash for the node
func (n *Node) GetHash() string {
	var buffer bytes.Buffer
	for n.parent != nil {
		buffer.WriteString(n.GetSector())
		n = n.prev()
	}
	return buffer.String()

}

func (n *Node) prev() *Node {
	return n.parent
}
