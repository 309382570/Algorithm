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
func sortList(head *ListNode) *ListNode {

	tmp := head
	stack := []int{}

	// [2,]
	for tmp != nil {
		stack = append(stack, tmp.Val)
		tmp = tmp.Next
	}

	for i := 0; i < len(stack)-2; i++ {
		for j := 0; j < len(stack)-1-i; j++ {
			if stack[j] > stack[j+1] {
				stack[j], stack[j+1] = stack[j+1], stack[j]
			}

		}
	}

	count := 0
	tmp2 := head
	for tmp2 != nil {
		tmp2.Val = stack[count]
		tmp2 = tmp2.Next
		count++
	}

	return head
}
