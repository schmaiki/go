package composite

func Add(left, right int) int {
	leftPtr := &left
	rightPtr := &right

	sum := *leftPtr + *rightPtr
	return sum
}
