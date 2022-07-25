package main

import "fmt"

func main() {
	newUser()
}

func newUser() {
	var user = NewUser("tom", "111111")
	fmt.Println(*user)
	println(user.getName())
	user.setPassword("123456")
	fmt.Println(*user)
}
