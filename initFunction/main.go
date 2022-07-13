package main

import (
	"fmt"
	"hello-go/initFunction/otherPackage"
)

func init() {
	fmt.Println("this main init function")
}

// 被依赖包的全局变量 -> 被依赖包的init函数 -> 当前包的全局变量 -> 当前包的init函数 -> 当前包的函数
func main() {
	otherPackage.SaySomething("something")
}
