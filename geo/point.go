package geo

import (
	"log"

	"github.com/tiborv/go-bitarray"
)

//Point Object
type Point struct {
	x        ba.BitArray //Morton envoding of X coordinate
	y        ba.BitArray //Morton encoding of Y cordinate
	morton   []string    //Morton encoding of the Point
	bitDepth int         //Number of bits describing this point
}

var stringToByte = map[string]byte{}

//NewPoint Point constructor
func NewPoint(x string, y string) Point {
	newPoint := Point{
		x:      ba.NewBitArray(x),
		y:      ba.NewBitArray(y),
		morton: make([]string, len(x)),
	}
	if len(x) != len(y) {
		log.Fatal("Point must have same cordinate lenght!")
	}
	newPoint.bitDepth = len(x)
	for i := 0; i < len(x); i++ {
		newPoint.morton[i] = string(x[i]) + string(y[i])
	}
	return newPoint
}

//GetX returns the x coordinate of a point
func (p Point) GetX() ba.BitArray {
	return p.x
}

//GetY returns the y coordinate of a point
func (p Point) GetY() ba.BitArray {
	return p.y
}

//Equals checks if two points are equal (the same position in space)
func (p Point) Equals(p2 Point) bool {
	return p.x.Equals(p2.x) && p.y.Equals(p2.y)
}

//GetMortonAt returns concat of x and y values at a position i
func (p Point) GetMortonAt(i int) (concat string, end bool) {
	if i >= p.bitDepth {
		log.Fatal("GetMortonAt out of range!")
	}
	if i == p.bitDepth-1 {
		end = true
	}
	concat = p.morton[i]
	return
}

//GetBitDepth returns bitDepth of a point
func (p Point) GetBitDepth() int {
	return p.bitDepth
}
