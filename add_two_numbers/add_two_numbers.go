// https://leetcode.com/problems/add-two-numbers/description/
package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy_head := ListNode{}
	tail := &dummy_head

	carry_over := 0

	for ; l1 != nil || l2 != nil ; l1, l2 = next(l1), next(l2) {
		sum := value(l1) + value(l2) + carry_over
		tail.Next = &ListNode{sum % 10, nil}
		carry_over = sum / 10
		tail = tail.Next
	}

	if carry_over != 0 {
		tail.Next = &ListNode{carry_over, nil}
	}

	return dummy_head.Next
}

func next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}
	return nil
}

func value(node *ListNode) int {
	if node != nil {
		return node.Val
	}
	return 0
}
