package calculator

func SumFromAToB(a, b int) int {
	//if a > b {
	//	return 0
	//}
	//return a + SumFromAToB(a+1, b)
	return ProcessFromAToB(a, b, 0, func(x, y int) int {
		return x + y
	})
}

func MultiplayFromAToB(a, b int) int {
	//if a > b {
	//	return 1
	//}
	//
	//return a * MultiplayFromAToB(a+1, b)
	return ProcessFromAToB(a, b, 1, func(x, y int) int {
		return x * y
	})
}

func ProcessFromAToB(a, b, initValue int, fn func(int, int) int) int {
	if a > b {
		return initValue
	}
	return fn(a, ProcessFromAToB(a+1, b, initValue, fn))
}
