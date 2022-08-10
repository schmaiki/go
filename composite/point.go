package composite

import "math"

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) DistanceFromZero() float64 { //Receiver Parameter
	return math.Sqrt(
		math.Pow(float64(p.X), 2) +
			math.Pow(float64(p.Y), 2))
}
