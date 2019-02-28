package ConstruckBinaryTreeII

/*
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

// TreeNode data
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
注意递归时传递的 前序 和 中序 切片
*/
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	if len(preorder) == 1 {
		return root
	}
	var idx int
	for i, v := range inorder {
		if v == root.Val {
			idx = i
		}
	}

	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
