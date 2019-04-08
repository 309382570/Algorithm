package InsertSort

func insertSort(nums []int) {

	if len(nums) < 1 {
		return
	}

	l := len(nums)
	for i := 1; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			// not use i here
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				break
			}
		}
	}

}
