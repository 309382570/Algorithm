package Add_Two_Numbers

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// related : leetcode 445
// 注意组合链表时的头节点，先记录下来
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := ListNode{}
	pH := &head
	var add_One int = 0

	for l1 != nil || l2 != nil {
		add_One = add_One / 10
		num := 0
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}

		pH.Next = &ListNode{Val: (num + add_One) % 10}
		add_One = num + add_One
		pH = pH.Next

	}

	if (add_One / 10) == 1 {
		pH.Next = &ListNode{Val: 1}
	}

	return head.Next
}
