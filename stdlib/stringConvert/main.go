package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmtConvert()

	stringConvert()

	stringOperation()
}

func stringOperation() {
	println(strings.Contains("hello world", " "))
	println(strings.Count("hello world", " "))
	println(strings.Index("hello world", " "))

	println(strings.Replace("hello world", " ", "_", 1))
	println(strings.ReplaceAll("hello world", " ", "_"))

	println(strings.Repeat("ha", 2))

	println(strings.EqualFold("ABC", "abc"))

	fmt.Println(strings.Fields("a b \b c\nd\t"))
	fmt.Println(strings.Split("abc", ""))

	println(strings.TrimSpace("   asdf  \n  "))
}

func stringConvert() {
	var i int64 = 666
	var s = strconv.FormatInt(i, 2)
	fmt.Println(s)

	var ui, err = strconv.ParseUint(s, 2, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(ui)
}

func fmtConvert() {
	var i = 666
	var s = "cp.org"
	var r = fmt.Sprintf("%d@%v", i, s)
	fmt.Println(r)

	var (
		i1 int
		s1 string
	)
	var n, err = fmt.Sscanf(r, "%d@%v", &i1, &s1)
	if err == nil {
		fmt.Println(n, i1, s1)
	}
}
