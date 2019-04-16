package RemoveDuplicatesFromSortedListII

/*
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

Example 1:

Input: 1->2->3->3->4->4->5
Output: 1->2->5
Example 2:

Input: 1->1->1->2->3
Output: 2->3
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//

func deleteDuplicates(head *ListNode) *ListNode {
	ret := &ListNode{Next: head}
	pre, cur := ret, head

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return ret.Next

}
