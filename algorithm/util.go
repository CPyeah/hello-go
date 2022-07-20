package algorithm

func Swap(array []int, i int, j int) {
	var temp = array[i]
	array[i] = array[j]
	array[j] = temp
}
