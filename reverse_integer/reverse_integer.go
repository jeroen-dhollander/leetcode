package reverse_integer

import (
	"math"
)

func reverse(input int) int {
	sign := signBit(input)

	result := 0
	for remainder := abs(input); remainder > 0; remainder = remainder / 10 {
		result = result*10 + remainder%10
	}

	if result > math.MaxInt32 {
		return 0
	}
	return sign * result
}

func signBit(input int) int {
	if input < 0 {
		return -1
	}
	return 1
}

func abs(input int) int {
	if input < 0 {
		return 0 - input
	}
	return input
}
