package AddTwoNumbersII

/*
You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

Example:

Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	node := ListNode{}
	sum := 0
	for len(stack1) > 0 || len(stack2) > 0 {
		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]

		}

		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack1[:len(stack2)-1]
		}

		node.Val = sum % 10
		head := ListNode{}
		head.Next = &node
		node = head
		sum = sum / 10
	}
	node.Val = sum
	if node.Val == 0 {
		return node.Next
	}
	return &node

}
