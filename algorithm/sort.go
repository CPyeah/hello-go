package algorithm

func bubbleSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j, j+1)
			}
		}
	}
	return nums
}
