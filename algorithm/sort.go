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

func selectionSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums)-1; i++ {
		var lastIndex = len(nums) - 1 - i
		var maxValue = nums[lastIndex]
		var maxIndex = lastIndex
		for j := 0; j < lastIndex; j++ {
			if nums[j] > maxValue {
				maxValue = nums[j]
				maxIndex = j
			}
		}
		if maxIndex != lastIndex {
			Swap(nums, maxIndex, lastIndex)
		}
	}
	return nums

}

func insertionSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		var currentValue = nums[i]
		var insertIndex = i - 1
		for ; insertIndex >= 0 && currentValue < nums[insertIndex]; insertIndex-- {
			// move backward
			nums[insertIndex+1] = nums[insertIndex]
		}
		// insert element
		if insertIndex != i-1 {
			nums[insertIndex+1] = currentValue
		}
	}
	return nums
}
