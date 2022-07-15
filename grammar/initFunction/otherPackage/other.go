package otherPackage

import "fmt"

func init() {
	fmt.Println("this is init function")
}

var A = getFunction()

func getFunction() func() {
	fmt.Println("globe variable")
	return func() {

	}
}

func init() {
	fmt.Println("this is another init function")
}

func SaySomething(something string) {
	fmt.Println(something)
}
