package MergeSort

//  归并排序的针对2个数组，都必须是有序的
func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	k := n / 2
	h1 := mergeSort(nums[:k])
	h2 := mergeSort(nums[k:])

	return merge(h1, h2)

}

func merge(a, b []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			tmp = append(tmp, a[i])
			i++
		} else {
			tmp = append(tmp, b[j])
			j++
		}
	}

	tmp = append(tmp, a[i:]...)
	tmp = append(tmp, b[i:]...)

	return tmp

}
