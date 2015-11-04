package geo

import (
	"log"
	"strconv"
)

//Point Object
type Point struct {
	x        string
	y        string
	xInt     int64
	yInt     int64
	bitDepth int
}

//NewPoint Point constructor
func NewPoint(x string, y string) Point {
	p := Point{x: x, y: y}
	if len(x) != len(y) {
		log.Fatal("Point must have same cordinate lenght!")
	}
	p.bitDepth = len(x)
	if xint, err := strconv.ParseInt(x, 2, 64); err != nil {
		log.Fatal(err)
	} else {
		p.xInt = xint
	}
	if yint, err := strconv.ParseInt(y, 2, 64); err != nil {
		log.Fatal(err)
	} else {
		p.yInt = yint
	}

	return p
}

//GetX returns the x coordinate of a point
func (p Point) GetX() string {
	return p.x
}

//GetY returns the y coordinate of a point
func (p Point) GetY() string {
	return p.y
}

//Equals checks if two points are equal (the same position in space)
func (p Point) Equals(p2 Point) bool {
	return p.x == p2.x && p.y == p2.y
}

//GetConcatAt returns concat of x and y values at a position i
func (p Point) GetConcatAt(i int) (concat string, end bool) {
	if i >= len(p.x) {
		log.Fatal("GetConcat out of range!")
	}
	if i == len(p.x)-1 {
		end = true
	}
	concat = string(p.x[i]) + string(p.y[i])
	return
}

//WithinArea checks if point is withing area defined by min and max point (Rectangle)
func (p Point) WithinArea(min Point, max Point) bool {
	return p.bitDepth == min.bitDepth && p.bitDepth == max.bitDepth && min.xInt <= p.xInt && min.yInt <= p.yInt && max.xInt >= p.xInt && max.yInt >= p.yInt
}
