package composite

import "math"

type HasDistance interface {
	DistanceFromZero() float64
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) HasDistance {
	return Point{x, y}
}

func (p Point) DistanceFromZero() float64 { //Receiver Parameter
	return math.Sqrt(
		math.Pow(float64(p.X), 2) +
			math.Pow(float64(p.Y), 2))
}
