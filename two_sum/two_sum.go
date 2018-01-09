// https://leetcode.com/problems/two-sum/description
package twoSum

func twoSum(nums []int, target int) []int {
	valueAtPosition := make(map[int]int, len(nums))

	for position, value := range nums {
		if firstPosition, ok := valueAtPosition[value]; ok {
			return []int{firstPosition, position}
		} else {
			valueAtPosition[target-value] = position
		}
	}

	panic("No solution found")
}
