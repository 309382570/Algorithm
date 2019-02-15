package MaximumDepth

// TreeNode data
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth : 考虑如果知道子树的结果，能否得出节点的结果。 bottom to top
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return getL(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func getL(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maxDepth1 : 考虑从 节点 带去一个值给 子树，时刻更新目标值。 Top to bottom 类似于前序遍历
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getL(maxDepth1(root.Left), maxDepth1(root.Right))
}
