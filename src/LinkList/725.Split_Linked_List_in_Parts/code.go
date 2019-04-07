package code

/*
Given a (singly) linked list with head node root, write a function to split the linked list into k consecutive linked list "parts".

The length of each part should be as equal as possible: no two parts should have a size differing by more than 1. This may lead to some parts being null.

The parts should be in order of occurrence in the input list, and parts occurring earlier should always have a size greater than or equal parts occurring later.

Return a List of ListNode's representing the linked list parts that are formed.

Examples 1->2->3->4, k = 5 // 5 equal parts [ [1], [2], [3], [4], null ]
Example 1:
Input:
root = [1, 2, 3], k = 5
Output: [[1],[2],[3],[],[]]
Explanation:
The input and each element of the output are ListNodes, not arrays.
For example, the input root has root.val = 1, root.next.val = 2, \root.next.next.val = 3, and root.next.next.next = null.
The first element output[0] has output[0].val = 1, output[0].next = null.
The last element output[4] is null, but it's string representation as a ListNode is [].
Example 2:
Input:
root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
Output: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
Explanation:
The input has been split into consecutive parts with size difference at most 1, and earlier parts are a larger size than the later parts.
Note:

The length of root will be in the range [0, 1000].
Each value of a node in the input will be an integer in the range [0, 999].
k will be an integer in the range [1, 50].
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	// 1、必须要遍历获得整个链表长度
	tmp := root
	l := 0
	for tmp != nil {
		l++
		tmp = tmp.Next
	}

	ret := make([]*ListNode, 0)

	// 2、比较 链表长度 和 要分成k组 之间的关系，
	// 每组中的个数: n = l / k

	n := l / k

	// 如果 总长度，小于还分开的组数。单个为一组，不足的加空
	// n 只可能 > 0 或 = 0，之前这里出错了
	if n == 0 {
		for root != nil {
			tmp := root.Next
			root.Next = nil
			ret = append(ret, root)
			root = tmp
		}
		for i := 1; i <= k-l; i++ {
			ret = append(ret, nil)
		}
		return ret
	}

	// 按k组平均分完之后，还剩下多少个 r = l % k
	r := l % k
	// 判断  k 和 r 之间的关系
	// r 不可能大于 n

	var head, end *ListNode = root, root

	for k > 0 {
		count := 1
		for count < n {
			end = end.Next
			count++
		}
		// 前 r 的组，个数需要加1
		if r > 0 {
			end = end.Next
			r--
		}

		tmp := end.Next
		end.Next = nil
		ret = append(ret, head)
		head = tmp
		end = head
		k--
	}

	return ret
}
