package geo

import (
	"bytes"
	"log"

	"github.com/smalllixin/bitarray"
)

//Point Object
type Point struct {
	hash *bitarray.BitArray
	x    string
	y    string
}

//NewPoint Point constructor
func NewPoint(x, y string) Point {
	if len(x) != len(y) {
		log.Fatal("Point must have same cordinate lenght!")
	}

	p := Point{x: x, y: y}
	p.genHash()
	return p
}

func (p *Point) genHash() *Point {
	p.hash = bitarray.NewBitArray(uint32(len(p.x)), 2)
	for i, c := range p.x {
		if c == '1' {
			if p.y[i] == '1' {
				p.hash.SetB(uint32(i), byte(3))
			} else {
				p.hash.SetB(uint32(i), byte(2))

			}
		} else {
			if p.y[i] == '1' {
				p.hash.SetB(uint32(i), byte(1))
			} else {
				p.hash.SetB(uint32(i), byte(0))
			}
		}
	}
	return p
}

//GetX returns the x coordinate of a point
func (p Point) Append(b byte) Point {
	c := Point{hash: bitarray.NewBitArray(uint32(len(p.x)+1), 2)}
	c.hash.SetB(uint32(len(p.x)+1), b)

	return c
}

//Equals checks if two points are equal (the same position in space)
func (p Point) Equals(p2 Point) bool {
	return bytes.Equal(p.hash.GetBytes(), p2.hash.GetBytes())
}

//Equals checks if two points are equal (the same position in space)
func (p Point) GetHashString(p2 Point) bool {
	return bytes.Equal(p.hash.GetBytes(), p2.hash.GetBytes())
}

//GetConcatAt returns concat of x and y values at a position is
func (p Point) GetConcatAt(i int) (concat byte, end bool) {
	if i == p.hash.GetAllocLen() {
		end = true
	}
	concat = p.hash.GetB(uint32(i))
	return
}

//GetBitDepth returns bitDepth of a point
func (p Point) GetBitDepth() int {
	return p.hash.GetAllocLen()
}

const zero = int64(0)

func (p Point) IsWithinRange(min Point, max Point) bool {
	return bytes.Compare(p.hash.GetBytes(), min.hash.GetBytes()) >= 0 &&
		bytes.Compare(p.hash.GetBytes(), max.hash.GetBytes()) <= 0
}
