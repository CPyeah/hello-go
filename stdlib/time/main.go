package main

import (
	"fmt"
	"time"
)

func main() {
	durationOperation()
	locationOperation()
}

func locationOperation() {
	var l, err = time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	fmt.Println(l)

	l, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
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
