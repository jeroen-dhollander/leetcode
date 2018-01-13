package reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input, expected int
	}{
		{1, 1},
		{12, 21},
		{123, 321},
		{-123, -321},
		{1534236469, 0}, // result would be bigger than max-int32
	}

	for _, test := range tests {
		actual := reverse(test.input)
		assert.Equal(t, test.expected, actual, "Reverse of %v is %v but got %v", test.input, test.expected, actual)
	}

}
