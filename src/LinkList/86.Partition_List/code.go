package Partition_List

/*
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// 从头遍历链表
	// 构造两个链表分别加入 Val >= x 或者 Val < x 的节点
	// 两个链表做合并

	var bigger, smaller *ListNode = &ListNode{}, &ListNode{}

	h1 := smaller
	h2 := bigger

	for head != nil {

		if head.Val < x {
			h1.Next = head
			h1 = h1.Next
		} else {
			h2.Next = head
			h2 = h2.Next
		}
		head = head.Next
	}
	h2.Next = nil
	h1.Next = bigger.Next

	return smaller.Next

}
