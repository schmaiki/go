package main

import (
	calculator "Go/Calculator"
	"fmt"
)

func main() {
	fmt.Println("24 + 45 =", calculator.Add(24, 45))
	fmt.Println(calculator.Divide(17, 3))

	fmt.Println(calculator.Sum(1, 10))
	fmt.Println(calculator.SumUntil(10))
}
