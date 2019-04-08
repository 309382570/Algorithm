/*

Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted linked list: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5

*/

package code

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [-10,-3,0,5]
// My sulotion
func sortedListToBST(head *ListNode) *TreeNode {

	// 1、获取链表总长度，将值顺序存储到slice中
	// 2、获取中间的值，递归构建 BST

	if head == nil {
		return nil
	}

	v := make([]int, 0)
	for head != nil {
		v = append(v, head.Val)
		head = head.Next
	}

	return buildBST(v)
}

func buildBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	// [0,1,2]
	// [0,1,2,3]
	mid := len(nums) / 2
	t := TreeNode{Val: nums[mid]}
	t.Left = buildBST(nums[:mid])
	t.Right = buildBST(nums[mid+1:])
	return &t
}

// Leetcode Sulotion

/*
public class Solution {
public TreeNode sortedListToBST(ListNode head) {
    if(head==null) return null;
    return toBST(head,null);
}
public TreeNode toBST(ListNode head, ListNode tail){
    ListNode slow = head;
    ListNode fast = head;
    if(head==tail) return null;

    while(fast!=tail&&fast.next!=tail){
        fast = fast.next.next;
        slow = slow.next;
    }
    TreeNode thead = new TreeNode(slow.val);
    thead.left = toBST(head,slow);
    thead.right = toBST(slow.next,tail);
    return thead;
}
*/
