package three_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input   []int
		results [][]int
	}{
		{
			[]int{-1, 0, 1},
			[][]int{
				{-1, 0, 1},
			},
		},
		{
			[]int{-4, -1, -1, 0, 1, 2},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			[]int{0, 0, 0},
			[][]int{
				{0, 0, 0},
			},
		},
		{
			[]int{0, 0, 0, 0},
			[][]int{
				{0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		result := threeSum(test.input)
		assert.Equal(t, test.results, result, "For %v expected %v but got %v", test.input, test.results, result)
	}
}
