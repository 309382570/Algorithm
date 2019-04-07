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
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	tmp := head
	var length int
	for tmp != nil {
		length++
		tmp = tmp.Next
	}

	return HelperFunc(head, 0, length-1)

}

func HelperFunc(root *ListNode, start, end int) *TreeNode {

	return nil
}

/*
private ListNode node;
public TreeNode sortedListToBST(ListNode head) {
	if(head == null){
		return null;
	}

	int size = 0;
	ListNode runner = head;
	node = head;

	while(runner != null){
		runner = runner.next;
		size ++;
	}

	return inorderHelper(0, size - 1);
}

[-10,-3,0,5,9]
public TreeNode inorderHelper(int start, int end){
	if(start > end){
		return null;
	}

	int mid = start + (end - start) / 2;
	TreeNode left = inorderHelper(start, mid - 1);

	TreeNode treenode = new TreeNode(node.val);
	treenode.left = left;
	node = node.next;

	TreeNode right = inorderHelper(mid + 1, end);
	treenode.right = right;

	return treenode;
}
*/
