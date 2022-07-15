package main

import "fmt"

func main() {
	type cpType string
	var myVar cpType = "cp"
	fmt.Printf("myVar 's value is %v, it's type is %T \n", myVar, myVar)

	var s = "yeah"
	myVar = cpType(s)
	fmt.Printf("myVar 's value is %v, it's type is %T \n", myVar, myVar)

	type myString = string
	var ms myString = "hello"
	fmt.Printf("myVar 's value is %v, it's type is %T \n", ms, ms)
}
