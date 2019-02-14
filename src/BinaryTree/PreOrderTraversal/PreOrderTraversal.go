package PreOrderTraversal

//  根, 左, 右

// TreeNode data
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrderTraversal1 in recrusion
func PreOrderTraversal1(root *TreeNode) []int {
	ret := []int{}
	if root != nil {
		ret = append(ret, root.Val)
		ret = append(ret, PreOrderTraversal1(root.Left)...)
		ret = append(ret, PreOrderTraversal1(root.Right)...)
	}
	return ret
}

// PreOrderTraversal in iteration
func PreOrderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	result := []int{}

	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, p.Val)

		// 这里的顺序，必须是先右后左，弹栈的顺序才会是先从左边开始
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return result
}
