// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package median_of_sorted_arrays

import (
	"math"
)

/*
To find the median of M numbers, you need to find the N = M/2 th number
(and potentially also the M/2 + 1 th number).

To find the Nth number in 2 sorted arrays,
we first take the N/2 th number in both arrays,
and compare them with each other.

It is impossible for the lowest one to be the N th number
(because in the 'worst case' the numbers are alternated stored in each array,
in which case the number we end up with is the N-1 th number).

Thus, we can strip this part from that array,
and repeat the algorithm (but this time searching for the N/4 th number).

Repeat until we need to find the first number.
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total_length := len(nums1) + len(nums2)
	N := (total_length + 1) / 2

	for N > 1 {
		i := N / 2
		nums1, nums2, _ = trimNLowest(nums1, nums2, i)
		N -= i
	}

	if isOdd(total_length) {
		_, _, median := trimNLowest(nums1, nums2, 1)
		return float64(median[0])
	} else {
		nums1, nums2, median_1 := trimNLowest(nums1, nums2, 1)
		_, _, median_2 := trimNLowest(nums1, nums2, 1)
		return float64(median_1[0]+median_2[0]) / 2.0
	}
}

func trimNLowest(nums1 []int, nums2 []int, N int) (new_nums1, new_nums2, trimmed_numbers []int) {
	value_1 := get_or_default(nums1, N-1, math.MaxInt32)
	value_2 := get_or_default(nums2, N-1, math.MaxInt32)

	if value_1 < value_2 {
		new_nums1 = nums1[N:]
		new_nums2 = nums2
		trimmed_numbers = nums1[:N]
	} else {
		new_nums1 = nums1
		new_nums2 = nums2[N:]
		trimmed_numbers = nums2[:N]
	}
	return
}

func isOdd(number int) bool {
	return number%2 != 0
}

func get_or_default(numbers []int, index int, default_value int) int {
	if index < len(numbers) {
		return numbers[index]
	} else {
		return default_value
	}
}
