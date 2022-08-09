package main

import (
	calculator "Go/Calculator"
	"fmt"
)

func main() {
	fmt.Println("24 + 45 =", calculator.Add(24, 45))
	fmt.Println(calculator.Divide(17, 3))
}
