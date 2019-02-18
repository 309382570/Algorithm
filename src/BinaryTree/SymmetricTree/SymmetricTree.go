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

//   1
//  / \
// 2   2
//  \   \
//   3    3

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return root.Left == root.Right
	}

	stack := []*TreeNode{root.Left, root.Right}

	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		l := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if r == nil && l == nil {
			continue
		}

		if r != nil {
			if l == nil {
				return false
			}
			if r.Val != l.Val {
				return false
			}

			if l.Left != nil {
				if r.Right == nil {
					return false
				}
			} else if r.Right != nil {
				return false
			}

			stack = append(stack, l.Left, r.Right)

			if l.Right != nil {
				if r.Left == nil {
					return false
				}
			} else if r.Left != nil {
				return false
			}
			stack = append(stack, l.Right, r.Left)

		} else if l == nil {
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
