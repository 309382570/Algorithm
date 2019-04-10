package SortList

/*
Sort a linked list in O(n log n) time using constant space complexity.
Example 1:
Input: 4->2->1->3
Output: 1->2->3->4
Example 2:
Input: -1->5->3->4->0
Output: -1->0->3->4->5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// [2,1]
// 使用递归
func sortList(head *ListNode) *ListNode {

	for tmp != nil {
		tmp = Recrusion(tmp, tmp.Next)
		tmp = tmp.Next
	}
	return Recrusion(head, head.Next)
}

func Recrusion(a, b *ListNode) *ListNode {
	if b == nil {
		return a
	}

	b = Recrusion(b, b.Next)

	if a.Val > b.Val {
		tmp := b.Val
		b.Val = a.Val
		a.Val = tmp
	}

	return a

}

//Leetcode
