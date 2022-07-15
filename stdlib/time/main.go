package main

import (
	"fmt"
	"time"
)

func main() {
	durationOperation()
}

func durationOperation() {
	var d = time.Second * 1000
	fmt.Println(d)
	time.Sleep(time.Second)

	var d1, err = time.ParseDuration("5m40s")
	if err != nil {
		panic(err)
	}
	fmt.Println(d1)

	var t1, err2 = time.Parse("2006年1月2日", "2019年10月13日")
	if err2 != nil {
		panic(err)
	}
	fmt.Println(t1)

	fmt.Println(time.Since(t1).Hours())
}
