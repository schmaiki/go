package calculator

func SumFromAToB(a, b int) int {
	if a > b {
		return 0
	}
	return a + SumFromAToB(a+1, b)
}

func MultiplayFromAToB(a, b int) int {
	if a > b {
		return 1
	}

	return a * MultiplayFromAToB(a+1, b)
}
