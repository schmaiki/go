package main

import (
	calculator "Go/Calculator"
	"Go/composite"
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)
	go func(q chan int) { //Go routine wird async aufgerufen
		time.Sleep(time.Second * 1)
		fmt.Println("Eine Sec vorbei")
		q <- 23
	}(queue)

	fmt.Println("24 + 45 =", calculator.Add(24, 45))
	fmt.Println(calculator.Divide(17, 3))

	valueFromQueue := <-queue

	fmt.Println(calculator.Sum(1, valueFromQueue))
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

	result, err := calculator.SquerRoot(-5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
