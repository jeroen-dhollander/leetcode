package brick_wall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeastBricks(t *testing.T) {
	tests := []struct {
		bricks      [][]int
		leastBricks int
	}{
		{
			[][]int{
				{1, 2, 2, 1},
				{3, 1, 2},
				{1, 3, 2},
				{2, 4},
				{3, 1, 2},
				{1, 3, 1, 1}},
			2,
		},
	}

	for _, test := range tests {
		result := leastBricks(test.bricks)
		assert.Equal(t, test.leastBricks, result, "%v can be crossed in %v but got %v",
			test.bricks, test.leastBricks, result)
	}

}
