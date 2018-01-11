package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		str    string
		length int
	}{
		{"a", 1},
		{"aa", 1},
		{"aba", 2},
		{"aab", 2},
		{"abcad", 4},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.str)
		assert.Equal(t, test.length, result, "Longest substring in '%v' has length %v but got %v", test.str, test.length, result)
	}
}
