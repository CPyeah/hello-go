package main

import (
	"fmt"
	"hello-go/grammar/note"
	"strconv"
)

func main() {
	countAndSay(4)
	fmt.Println("hello world!") // in current function
	sayHello()                  // in current file
	//speakHello()                // in current package
	note.SayHello() // in other package

	fmt.Println()
}

func sayHello() {
	fmt.Println("hello!")
}

func countAndSay(n int) string {
	num := "1"
	for i := 2; i <= n; i++ {
		num = say(num)
	}
	return num
}
func say(seq string) string {
	count := 0
	target := seq[0]
	res := ""
	for i := 0; i < len(seq); i++ {
		if seq[i] == target {
			count++
		} else {
			res += strconv.Itoa(count)
			res += string(target)
			target = seq[i]
			count = 1
		}
	}
	// last one target
	res += strconv.Itoa(count)
	res += string(target)
	return res
}
