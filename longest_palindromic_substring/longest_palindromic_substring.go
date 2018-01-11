package longest_palindromic_substring

import ()

func longestPalindrome(s string) string {
	result := ""

	for position, _ := range s {
		// look for longest odd palindrome around center, like 'aba'
		start, length := expandAround(s, position, position)
		if length > len(result) {
			result = s[start : start+length]
		}

		// look for longest even palindrome around center, like 'abba'
		start, length = expandAround(s, position, position+1)
		if length > len(result) {
			result = s[start : start+length]
		}
	}

	return result
}

func expandAround(s string, left, right int) (start, length int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	start = left + 1
	length = right - start
	return
}
