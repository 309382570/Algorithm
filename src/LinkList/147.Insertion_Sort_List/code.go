package Insertion_Sort_List

/*
Sort a linked list using insertion sort.

A graphical example of insertion sort. The partial sorted list (black) initially contains only the first element in the list.
With each iteration one element (red) is removed from the input data and inserted in-place into the sorted list

Algorithm of Insertion Sort:

Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
It repeats until no input elements remain.

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

// 用递归
func insertionSortList(head *ListNode) *ListNode {

	tmp := head

	for tmp != nil {
		tmp = exchange(tmp, tmp.Next)
		tmp = tmp.Next
	}
	return head

}

func exchange(first, second *ListNode) *ListNode {

	if second == nil {
		return first
	}

	second = exchange(second, second.Next)

	if first.Val > second.Val {
		tmp := first.Val
		first.Val = second.Val
		second.Val = tmp
	}

	return first

}

// Leetcode

func test(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := head
	arr := make([]*ListNode, 0)
	for tmp != nil {
		arr = append(arr, tmp)
		tmp = tmp.Next
	}

	n := len(arr)

	for i := 1; i < n; i++ {
		for j := i - 1; j > 0; j-- {
			if arr[j].Val > arr[j+1].Val {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				break
			}
		}
	}

	var pre, cur *ListNode
	cur = arr[0]
	pre = cur

	for i := 1; i < n; i++ {
		cur.Next = arr[i]
		cur = cur.Next
	}

	cur.Next = nil

	return pre
}
