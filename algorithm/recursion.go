package algorithm

var cache []int

// fibonacci 1 1 2 3 5 8 13....
func fibonacci(num int) int {
	if cache == nil || len(cache) < num+1 {
		cache = make([]int, num+1)
	}
	if num <= 2 {
		return 1
	}
	if cache[num] == 0 {
		cache[num] = fibonacci(num-1) + fibonacci(num-2)
	}
	return cache[num]
}
