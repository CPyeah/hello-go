package main

import (
	"fmt"
	"hello-go/grammar/note"
)

func main() {
	fmt.Println("hello world!") // in current function
	sayHello()                  // in current file
	speakHello()                // in current package
	note.SayHello()             // in other package

	fmt.Println()
}

func sayHello() {
	fmt.Println("hello!")
}
