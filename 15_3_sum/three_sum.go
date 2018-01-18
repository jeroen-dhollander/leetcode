package three_sum

import (
	"sort"
)

func threeSum(numbers []int) [][]int {
	results := [][]int{}

	sort.Ints(numbers)

	for a := 0; a < len(numbers)-2; a++ {

		if a > 0 && numbers[a-1] == numbers[a] {
			continue // skip duplicates
		}

		b, c := a+1, len(numbers)-1

		for b < c {
			sum := numbers[a] + numbers[b] + numbers[c]
			if sum == 0 {
				results = append(results, []int{numbers[a], numbers[b], numbers[c]})

				for b++; b < c && numbers[b-1] == numbers[b]; b++ {
					// advance b but skip duplicates
				}
			} else if sum < 0 {
				b++
			} else {
				c--
			}
		}
	}
	return results
}
