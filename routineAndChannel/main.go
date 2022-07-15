package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = sync.Mutex{}

func main() {
	normalInvoke()

	goroutine()
}

func normalInvoke() {
	fmt.Println("normalInvoke")
	var start = time.Now().UnixMilli()
	var count = 0
	for i := 2; i < 100001; i++ {
		if isPrimeNum(i) {
			count++
		}
	}
	var cost = time.Now().UnixMilli() - start
	fmt.Println("count is", count, "cost time", cost)
}

func goroutine() {
	fmt.Println("goroutine Invoke")
	var start = time.Now().UnixMilli()
	var count = 0
	for i := 2; i < 100001; i++ {
		go primeNum(i, &count)
	}
	var cost = time.Now().UnixMilli() - start
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("count is", count, "cost time", cost)
}

func isPrimeNum(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primeNum(num int, count *int) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return
		}
	}
	lock.Lock()
	*count++
	lock.Unlock()
}
