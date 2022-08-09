package calculator

// Packageebene Vaariable
// PI:= 3.14159265359 funktioniert nicht

var PI float64 = 3.14159265359

func Add(left, right int) int {

	var sum int = left + right

	// Möglichkeiten für Variablen
	// var sum int
	// var sum = left + right
	// sum := left + right

	return sum
}

func Divide(left, right int) (quotient int, remainder int) {
	quotient = left / right
	remainder = left % right

	return
}
