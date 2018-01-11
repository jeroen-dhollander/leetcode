package longest_substring_without_repeating_characters

import ()

func lengthOfLongestSubstring(s string) int {
	positions := make(map[rune]int)

	longest := -1
	substring_start := 0

	for position, character := range s {

		previous_position := positions[character]

		if previous_position > substring_start {
			length := position - substring_start
			longest = max(longest, length)
			// The unique substring starts again just after the previous time we encountered 'character'
			substring_start = previous_position
		}

		positions[character] = position + 1
	}

	longest = max(longest, len(s)-substring_start)
	return longest
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}
