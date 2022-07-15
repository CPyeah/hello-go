package main

import (
	"fmt"
	"hello-go/grammar/note"
)

var globalVariable = 100

const globalConstant = 101

func main() {
	variable()
	constant()
	fmt.Println(note.GlobalVersion)
}

func variable() {
	var v1 int
	v1 = 1
	var v2 int = 2
	var v3 = 3
	v4 := 4
	fmt.Println(v1, v2, v3, v4)

	var (
		v5 = 5
		v6 = 6
	)
	fmt.Println(v5, v6)
	fmt.Println(globalVariable)
}

func constant() {
	const c1 = 1 // 1
	const (
		c2 = 2    //2
		c3 = iota //1 当前行数（从0开始）
		c4 = iota //2
		c5        //3 默认值为上一行
		c6        //4
		c7 = 7    //7
		c8        //7 默认值为上一行
		c9        //7
	)
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9)
	fmt.Println(globalConstant)
}
