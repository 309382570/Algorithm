package InOrderTraversal

// // 左，根，右

// TreeNode data
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InOrderTraversal in Recursion
func InOrderTraversal(root *TreeNode) []int {
	var ret []int
	if root != nil {

		ret = append(ret, InOrderTraversal(root.Left)...)
		ret = append(ret, root.Val)
		ret = append(ret, InOrderTraversal(root.Right)...)
	}
	return ret
}

// InOrderTraversal2 in Iterative
func InOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := []int{}
	stack := []*TreeNode{root}
	t := root.Left
	for {
		// 必须在判断空stack前，否则右边的叶子会缺失
		for t != nil {
			stack = append(stack, t)
			t = t.Left
		}

		if len(stack) == 0 {
			break
		}

		t = stack[len(stack)-1]
		ret = append(ret, t.Val)
		stack = stack[:len(stack)-1]

		t = t.Right

	}
	return ret
}
