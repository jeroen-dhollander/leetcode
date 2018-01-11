// https://leetcode.com/problems/add-two-numbers/description/
package addTwoNumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		first  []int
		second []int
		sum    []int
	}{
		{
			[]int{1}, []int{1},
			[]int{2},
		},
		{
			[]int{1, 1}, []int{2, 2},
			[]int{3, 3},
		},
		{
			[]int{1, 1, 1}, []int{2, 2, 2},
			[]int{3, 3, 3},
		},
		{
			[]int{1, 1}, []int{2},
			[]int{3, 1},
		},
		{
			[]int{5, 0}, []int{5, 0},
			[]int{0, 1},
		},
		{
			[]int{5}, []int{5},
			[]int{0, 1},
		},
		{
			[]int{1}, []int{9, 9, 9},
			[]int{0, 0, 0, 1},
		},
	}

	for _, test := range tests {
		first := toListNode(test.first)
		second := toListNode(test.second)
		sum := fromListNode(addTwoNumbers(first, second))
		assert.Equal(t, test.sum, sum, "%v + %v = %v, but got %v", test.first, test.second, test.sum, sum)
	}
}

func toListNode(numbers []int) *ListNode {
	if len(numbers) == 0 {
		return nil
	}

	return &ListNode{numbers[0], toListNode(numbers[1:])}
}

func fromListNode(node *ListNode) []int {
	var result []int

	for n := node; n != nil; n = n.Next {
		result = append(result, n.Val)
	}
	return result
}
