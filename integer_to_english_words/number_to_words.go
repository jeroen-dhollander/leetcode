// https://leetcode.com/problems/integer-to-english-words/description/
package integer_to_english_words

import (
	"strings"
)

var (
	belowTwenty  = [...]string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	BelowHundred = [...]string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	units        = [...]string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(number int) string {

	remainder := number
	unit := 0
	parts := []string{}

	if number == 0 {
		return "Zero"
	}

	for ; remainder > 0; remainder /= 1000 {
		number := remainder % 1000
		parts = append(formatNumberWithUnits(number, units[unit]), parts...)

		unit++
	}
	return strings.Join(parts, " ")

}

/* This formats numbers up to 999, and then adds the units string */
func formatNumberWithUnits(number int, unit string) []string {
	if number >= 1000 {
		panic("Too high")
	}

	if number == 0 {
		return []string{}
	}

	result := append(formatHundreds(number), formatTens(number)...)
	if len(unit) > 0 {
		result = append(result, unit)
	}
	return result
}

func formatHundreds(number int) []string {
	if number >= 100 {
		return []string{belowTwenty[number/100], "Hundred"}
	}
	return []string{}
}

/*
Formats a number up to 99
*/
func formatTens(number int) []string {
	number = number % 100

	switch {
	case number == 0:
		return []string{}
	case number < 20:
		return []string{belowTwenty[number]}
	case number%10 == 0:
		return []string{BelowHundred[number/10]}
	default:
		return []string{BelowHundred[number/10], belowTwenty[number%10]}
	}
}
