package ReverseLinkListII

/*
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/

// 找到要反转的链表，反转后记得将头尾连接

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if m == n {
		return head
	}

	// Header := new(ListNode)
	// Header.Next = head
	// cur := Header
	// var pre *ListNode

	pHead := new(ListNode)
	pHead.Next = head
	cur := pHead
	var pre *ListNode

	// pre 为 m-1 ， cur为 m
	for i := 0; i < m; i++ {
		pre = cur
		cur = cur.Next
	}

	pre2 := pre
	cur2 := cur
	var p *ListNode

	for j := m; j <= n; j++ {
		p = cur2.Next
		cur2.Next = pre2
		pre2 = cur2
		cur2 = p
	}

	pre.Next = pre2
	cur.Next = cur2

	return pHead.Next
}
