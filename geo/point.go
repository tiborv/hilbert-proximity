package geo

//Point Object
type Point struct {
	x string
	y string
}

//NewPoint Point constructor
func NewPoint(x string, y string) Point {
	p := Point{x: x, y: y}
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
	if i+1 == len(p.x) {
		end = true
	}
	concat = string(p.x[i]) + string(p.y[i])
	return
}
