package main

import "fmt"

func main() {
	array()

	towDArray()
}

func array() {
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(array[1], len(array))

	var array1 = [...]int{4,
		5,
		6}
	fmt.Println(array1[0])

	for i, v := range array1 {
		fmt.Println(i, "的值是", v)
	}
}

func towDArray() {
	var points = [3][2]int{{1, 1}, {2, 2}, {3, 3}}
	for _, point := range points {
		fmt.Println(point)
	}
}
