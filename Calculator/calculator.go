package calculator

func Add(left, right int) int {
	return left + right
}

func Divide(left, right int) (quotient int, remainder int) {
	quotient = left / right
	remainder = left % right

	return
}
