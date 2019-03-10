package SwapNodesinPairs

/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.
Example:
Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := head.Next

	tmp := head

	for tmp != nil && tmp.Next != nil {
		t := tmp.Next
		if t.Next != nil {
			t2 := t.Next
			t.Next = tmp
			tmp = t2
		}
		break
	}

	return ret

}

/*
每个操作都一样，交换。
考虑递归
*/

func recurion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := head.Next
	head.Next = recurion(head.Next.Next)
	n.Next = head
	return n
}
