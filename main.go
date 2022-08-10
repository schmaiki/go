package main

import (
	calculator "Go/Calculator"
	"Go/composite"
	"fmt"
)

func main() {
	fmt.Println("24 + 45 =", calculator.Add(24, 45))
	fmt.Println(calculator.Divide(17, 3))
	fmt.Println(calculator.Sum(1, 10))
	fmt.Println(calculator.SumUntil(10))
	fmt.Println(calculator.IsSquareNumber(24))
	fmt.Println(calculator.RunOperation("add", 12, 13))
	fmt.Println(calculator.RunOperation("sub", 12, 13))
	fmt.Println(calculator.MultiplayFromAToB(1, 10))
	fmt.Println(composite.Add(10, 23))

	point := composite.Point{
		X: 3,
		Y: 2,
	}

	fmt.Println(point.X, point.Y)

	point1 := composite.NewPoint(3, 8)
	fmt.Println(point1.(composite.Point).X, point1.(composite.Point).Y)
	//point1 := composite.NewPoint(3, 8)
	//fmt.Println(point1.X, point1.Y)

	//composite.DemoCollections()
	fmt.Println(point.DistanceFromZero())
}
