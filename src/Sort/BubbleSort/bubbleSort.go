package bubbleSort

// 冒泡排序
/*
1、时间复杂度 : O(n^2)
2、是稳定的算法，它满足稳定算法的定
3、思路：相邻比较，逐步减少每一次的遍历长度。最大追加到最后
*/

// func bubbleSort(nums []int) {
// 	for j := len(nums) - 1; j > 0; j-- {
// 		for i := 0; i < j; i++ {
// 			if nums[i] > nums[i+1] {
// 				nums[i], nums[i+1] = nums[i+1], nums[i]
// 			}
// 		}
// 	}
// }

func bubbleSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	for i := 0; i < n-1; i++ {
		isChange := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			isChange = true
		}
		if !isChange {
			break
		}
	}
}

/*
for(int i=0;i<arr.length-1;i++){//外层循环控制排序趟数
　　　　　　for(int j=0;j<arr.length-1-i;j++){//内层循环控制每一趟排序多少次
　　　　　　　　if(arr[j]>arr[j+1]){
　　　　　　　　　　int temp=arr[j];
　　　　　　　　　　arr[j]=arr[j+1];
　　　　　　　　　　arr[j+1]=temp;
　　　　　　　　}
　　　　　　}
　　　　}
*/
