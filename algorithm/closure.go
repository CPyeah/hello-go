package algorithm

import "fmt"

func closure() {
	var f = closureFunc()
	for i := 0; i < 10; i++ {
		f()
	}
	fmt.Println("reset count")
	f = closureFunc()
	for i := 0; i < 10; i++ {
		f()
	}
}

func closureFunc() func() {
	var count = 0
	return func() {
		count++
		fmt.Printf("this is %v time invoke function\n", count)
	}
}
