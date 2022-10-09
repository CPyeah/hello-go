package main

import (
	"fmt"
	"math/rand"
)

func main() {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	ifDemo()
	switchDemo()
	forDemo()
}

func ifDemo() {
	var s = rand.Int31n(100)
	fmt.Println(s)
	if s > 60 {
		fmt.Println("pass")
	} else {
		fmt.Println("fall")
	}
}

func switchDemo() {
	var s int32 = rand.Int31n(100)
	fmt.Println(s)
	switch {
	case s < 10:
		fmt.Println("太差了")
	case s < 60:
		fmt.Println("不及格")
	case s < 80:
		fallthrough
	case s < 100:
		fmt.Println("good")

	}
}

func forDemo() {
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Print(i)
		i++
	}

	fmt.Println()

	i = 0
	for i < 10 {
		fmt.Print(i)
		i++
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

}
func trap(height []int) int {
	stack := []int{}
	total := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// pop
			bottomIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			// peek
			leftIndex := stack[len(stack)-1]

			v := (i - leftIndex - 1) * (min(height[i], height[leftIndex]) - height[bottomIndex])
			total += v
		}
		stack = append(stack, i)
	}
	return total
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
