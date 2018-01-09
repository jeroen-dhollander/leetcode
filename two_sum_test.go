package twoSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		givenNumbers      []int
		target            int
		expectedPositions []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{2, 7, 11, 15},
			13,
			[]int{0, 2},
		},
		{
			[]int{7, 2, 11, 15},
			13,
			[]int{1, 2},
		},
	}

	for _, test := range tests {
		result := twoSum(test.givenNumbers, test.target)
		assert.Equal(t, test.expectedPositions, result, "Target %v, Numbers %v", test.target, test.givenNumbers)
	}
}

func TestTwoSumDoesNotReturnSamePositionTwice(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
}
