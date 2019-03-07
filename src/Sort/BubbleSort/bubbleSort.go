package bubbleSort

// 冒泡排序
/*
1、时间复杂度 : O(n^2)
2、是稳定的算法，它满足稳定算法的定
3、思路：相邻比较，逐步减少每一次的遍历长度。最大追加到最后
*/

func bubbleSort(nums []int) {
	for j := len(nums) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}
