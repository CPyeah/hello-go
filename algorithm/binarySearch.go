package algorithm

func binarySearch(sorted []int, target int) int {
	var (
		left  = 0
		right = len(sorted) - 1
		mid   int
	)
	for {
		if sorted[left] == target {
			return left
		}
		if sorted[right] == target {
			return right
		}
		if right-left < 2 {
			return -1
		}
		mid = left + (right-left)/2
		if sorted[mid] == target {
			return mid
		}
		if sorted[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
