package reverseLinkList

/*
Reverse a singly linked list.
Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var next *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = next
		next = head
		head = tmp
	}
	return next
}
