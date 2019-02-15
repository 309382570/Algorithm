package SymmetricTree

// 判断是否是对称树
/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3

*/

// TreeNode data
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	isSymmetric in Recursion
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return ish(root.Left, root.Right)
}

func ish(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}
	if l.Val != r.Val {
		return false
	}

	return ish(l.Left, r.Right) && ish(l.Right, r.Left)
}

// isSymmetric in iteraion
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode
	if root.Left == nil {
		if root.Right != nil {
			return false
		}
		// 不能用下面这个， 加了 nil之后，slice里面长度会加 1 ！！！
		// stack = append(stack, root.Left, root.Right)
		if root.Right == nil {
			return true
		}

	} else if root.Right == nil {
		return false
	}

	stack = append(stack, root.Right, root.Left)

	for len(stack) > 0 {

		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		right := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// if left != nil || right !=nil {
		// 	return left == right
		// }

		// if left != nil {
		// 	if right == nil {
		// 		return false
		// 	}
		// 	if left.Val != right.Val {
		// 		return false
		// 	}
		// } else if right != nil {
		// 	return false
		// }

		if left.Left != nil {
			if right.Right == nil {
				return false
			}
			stack = append(stack, right.Right, left.Left)

		} else if right.Right != nil {
			return false
		}

		if left.Right != nil {
			if right.Left == nil {
				return false
			}
			stack = append(stack, left.Right, right.Left)
		} else if right.Left == nil {
			return false
		}

	}

	return true
}

/*
public boolean isSymmetric(TreeNode root) {
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        if(root == null) return true;
        q.add(root.left);
        q.add(root.right);
        while(q.size() > 1){
            TreeNode left = q.poll(),
                     right = q.poll();
            if(left== null&& right == null) continue;
            if(left == null ^ right == null) return false;
            if(left.val != right.val) return false;
            q.add(left.left);
            q.add(right.right);
            q.add(left.right);
            q.add(right.left);
        }
        return true;
	}
*/
