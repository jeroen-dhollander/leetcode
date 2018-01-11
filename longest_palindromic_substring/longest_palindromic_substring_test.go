package longest_palindromic_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrom(t *testing.T) {
	tests := []struct {
		str       string
		substring string
	}{
		{"aba", "aba"},
		{"cabad", "aba"},
		{"aa", "aa"},
		{"baad", "aa"},
	}

	for _, test := range tests {
		substring := longestPalindrome(test.str)
		assert.Equal(t, test.substring, substring, "Longest substring in '%v' is '%v' but got '%v'", test.str, test.substring, substring)
	}
}
