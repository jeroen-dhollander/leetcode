package integer_to_english_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		number int
		words  string
	}{
		{0, "Zero"},
		{1, "One"},
		{11, "Eleven"},
		{20, "Twenty"},
		{23, "Twenty Three"},
		{100, "One Hundred"},
		{101, "One Hundred One"},
		{4005, "Four Thousand Five"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	}

	for _, test := range tests {
		result := numberToWords(test.number)

		assert.Equal(t, test.words, result, "%v is '%v' but got '%v'", test.number, test.words, result)
	}

}
