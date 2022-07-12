package main

import "fmt"

func main() {
	result := add(7, 8)
	fmt.Println(result)

	sum, diff := sumAndDiff(9, 4)
	fmt.Println(sum, diff)

	sum, diff = simpleReturn(4, 9)
	fmt.Println(sum, diff)

	anonymousFunc()

	immediateInvoke()
}

func add(a int, b int) int {
	return a + b
}

func sumAndDiff(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func simpleReturn(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func anonymousFunc() {
	var f = func() {
		fmt.Println("I am anonymous function")
	}
	f()
}

func immediateInvoke() {
	func() {
		fmt.Println("Immediate invoke function")
	}()
}
