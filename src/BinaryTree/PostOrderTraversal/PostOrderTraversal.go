package PostOrderTraversal

// 左 右 根

// TreeNode data
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PostOrderTraversal in Recrusion
func PostOrderTraversal(root *TreeNode) []int {
	var s []int
	if root != nil {
		s = append(s, PostOrderTraversal(root.Left)...)
		s = append(s, PostOrderTraversal(root.Right)...)
		s = append(s, root.Val)
	}
	return s
}

// PostOrderTraversal1 in Iteration
func PostOrderTraversal1(root *TreeNode) []int {
	// 思路 后序是左右中。我们获取到中右左，然后反转slice即可
	var ret []int
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			ret = append(ret, curr.Val)
			curr = curr.Right
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = curr.Left
	}

	return reverse(ret)
}

// reverse the slice element
func reverse(nums []int) []int {
	for index := 0; index < len(nums)/2; index++ {
		j := len(nums) - index - 1
		nums[index], nums[j] = nums[j], nums[index]
	}
	return nums
}

// func reverse(numbers []int) []int {
// for i := 0; i < len(numbers)/2; i++ {
// j := len(numbers) - i - 1
// numbers[i], numbers[j] = numbers[j], numbers[i]
// }
// return numbers
// }
