package calculator

func Sum(min, max int) int {
	sum := 0

	for i := min; i < max; i++ {
		sum += i
	}
	return sum
}

//Diese Scheife entspricht einer While Schleife

func SumUntil(max int) int {

	sum := 1

	for sum < max {
		sum += sum
	}

	return sum
}
