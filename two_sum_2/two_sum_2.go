// https://leetcode.com/problems/two-sum/description
package twoSum2

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
}
