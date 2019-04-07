package Next_Greader_Node_In_LinketList

/*
Example 1:

Input: [2,1,5]
Output: [5,5,0]
Example 2:

Input: [2,7,4,3,5]
Output: [7,0,5,5,0]
Example 3:

Input: [1,7,5,1,9,2,5,1]
Output: [7,9,9,9,0,5,0,0]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// My solution
func nextLargerNodes(head *ListNode) []int {
	tmp := head
	// get all elment from list -> ret
	ret := []int{}
	for tmp != nil {
		ret = append(ret, tmp.Val)
		tmp = tmp.Next
	}
	var max int
	//Input: [1,7,5,1,9,2,5,1]
	for i := 0; i < len(ret)-1; i++ {
		for j := i + 1; j < len(ret); j++ {
			if ret[j] <= ret[i] {
				if j == len(ret)-1 {
					ret[i] = -1
				}
				continue
			}
			ret[i] = ret[j]
			max = ret[j]
			break
		}
		if ret[i] != -1 {
			ret[i] = max
		} else {
			ret[i] = 0
		}
	}
	ret[len(ret)-1] = 0
	return ret
}

// Answer from others
// Using stack

func nextLargerNodes2(head *ListNode) []int {
	return nil
}
