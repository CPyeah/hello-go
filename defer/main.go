package main

import "fmt"

var count = 0

func main() {
	deferDemo()
	deferRecover()
	fmt.Println("main function end.")
}

func deferDemo() {
	defer f(1)
	defer f(2)
	defer f(f(4))
	f(5)
}

func f(i int) int {
	count++
	fmt.Printf("param is %v, this is the %v time invoke \n", i, count)
	return i
}

func deferRecover() {
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println(error)
		}
	}()
	i := 0
	fmt.Println(1 / i)
}
