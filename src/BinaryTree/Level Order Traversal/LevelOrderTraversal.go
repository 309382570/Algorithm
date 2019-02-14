package LevelOrderTraversal

// TreeNode data
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		lenQ := len(q)
		var tmp []int
		for i := 0; i < lenQ; i++ {
			n := q[i]
			tmp = append(tmp, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		q = q[lenQ:]
		ret = append(ret, tmp)
	}

	return ret
}
