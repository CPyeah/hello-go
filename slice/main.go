package main

import "fmt"

func main() {
	createSlice()
	lenAndCap()
	appendSlice()
	copySlice()
	stringAndByte()
	paramSlice("ls", "-a", "-h")
}

func paramSlice(command string, params ...string) {
	fmt.Println("command is", command)
	fmt.Println("params is", params)
}

func stringAndByte() {
	var s = "Hello 世界"
	var b = []byte(s)
	var subString = s[0:5]
	fmt.Println(s, b, subString)

	for i, v := range s {
		fmt.Println(i, string(v))
	}
}

func copySlice() {
	var s1 = []int{1, 2, 3, 4, 5}
	var s2 = []int{6, 7}
	fmt.Println(s1, s2)
	copy(s1, s2)
	fmt.Println(s1, s2)
}

func appendSlice() {
	var s1 = []int{1, 2, 3}
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1, 4)
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1, []int{5, 6, 7}...)
	fmt.Println(s1, len(s1), cap(s1))
}

func lenAndCap() {
	var slice = make([]int, 8, 10)
	fmt.Println(slice, len(slice), cap(slice))
}

func createSlice() {
	// from array
	array := [5]int{1, 2, 3, 4, 5}
	var s1 = array[0 : len(array)-1]
	fmt.Println(s1)

	// from slice
	var s2 = s1[1:3]
	fmt.Println(s2)

	// make
	s3 := make([]int, 0)
	fmt.Println(s3)
	s3 = append(s3, 3)
	fmt.Println(s3)

	// sugar
	s4 := []int{4}
	fmt.Println(s4)
}
