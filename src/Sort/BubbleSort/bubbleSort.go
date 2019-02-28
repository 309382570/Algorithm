package bubbleSort

// 冒泡排序
/*
最坏时间复杂度 n^2
最优时间复杂度 n
*/

func bubbleSort(nums []int) {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

}
