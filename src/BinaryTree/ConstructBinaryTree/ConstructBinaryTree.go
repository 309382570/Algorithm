package ConstructBinaryTree

/*
中序，后续
Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
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

// buildTree
// inorder = [9,3,15,20,7]
// postorder = [9,15,7,20,3]

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]

	var idx int

	for _, v := range inorder {
		if v == val {
			break
		}
		idx++
	}

	l := buildTree(inorder[:idx], postorder[:idx])
	r := buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])

	root := &TreeNode{Val: val, Left: l, Right: r}
	return root
}
