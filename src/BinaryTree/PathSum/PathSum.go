package PathSum

/*
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
*/

// TreeNode data
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil && sum == 0 {
		return false
	}
	return handle(root, sum, 0)
}

func handle(root *TreeNode, target, tmp int) bool {
	sum := tmp
	if root == nil {
		return false
	}
	// if root.Left == nil && root.Right == nil {
	// 	sum = sum + root.Val
	// 	if sum != target {
	// 		return false
	// 	}
	// }

	if root.Left == nil && root.Right == nil && tmp+root.Val == target {
		return true
	}

	return handle(root.Left, target, sum+root.Val) || handle(root.Right, target, sum+root.Val)
}

// BestPractise : Answer from discuss
func BestPractise(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
		return true
	}

	return BestPractise(root.Right, sum-root.Val) || BestPractise(root.Left, sum-root.Val)
}
