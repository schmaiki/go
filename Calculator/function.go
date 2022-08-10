package calculator

import (
	"errors"
	"math"
)

func Abs(value int) int {
	if value >= 0 {
		return value
	}
	return -value
}

func IsSquareNumber(value int) bool {
	if sqrt := math.Sqrt(float64(value)); sqrt != float64(int(sqrt)) {
		return false
	}

	return true
}

func RunOperation(operation string, left, right int) int {
	var result int

	switch operation {
	case "add":
		result = left + right
	case "sub":
		result = left - right
	default:
		/// .....
	}

	return result
	/*
		Alternative Schreibweise
		switch {
		case operation == "add":
			result = left + right
		case operation == "sub":
			result = left - right
		default:
			/// .....
		}
		return result*/
}

func SquerRoot(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("value must not be nagative")
	}
	return math.Sqrt(value), nil
}
