package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())
	fmt.Println(runtime.GOMAXPROCS(8))

	go routineFunction()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
}

func routineFunction() {
	fmt.Println("routineFunction")
	runtime.Goexit()
	fmt.Println("never be printed")
}
