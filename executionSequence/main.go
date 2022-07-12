package main

import (
	"fmt"
	"math/rand"
)

func main() {
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
