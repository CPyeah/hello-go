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

	basicChannel()

	goChannelInvoke()
}

func goChannelInvoke() {
	var c = make(chan int)

	// find and add
	go func(c chan int) {
		for i := 2; i < 100001; i++ {
			func(num int, c chan int) {
				for j := 2; j < num; j++ {
					if num%j == 0 {
						return
					}
				}
				c <- num
			}(i, c)
		}
		close(c)
	}(c)

	fmt.Println(c)
	// get and count
	var count = 0
out:
	for true {
		select {
		case v := <-c:
			count++
			fmt.Println(v)
		default:
			break out
		}
	}
	for v := range c {
		fmt.Println(v)
		count++
	}
	fmt.Println("count is ", count)

	//time.Sleep(time.Duration(2) * time.Second)
}

func basicChannel() {
	var c = make(chan int, 10)
	go consumer1(&c)
	go producer1(&c)
	go producer2(&c)
	go consumer2(&c)

	time.Sleep(time.Duration(2) * time.Second)
}

func producer1(c *chan int) {
	for i := 0; i < 100000; i++ {
		*c <- i
	}
	//close(*c)
}

func producer2(c *chan int) {
	for i := 100000; i < 200000; i++ {
		*c <- i
	}
	//close(*c)
}

func consumer1(c *chan int) {
	for {
		var out, ok = <-*c
		if ok {
			fmt.Println("out 1 ->", out)
		} else {
			break
		}
	}
}

func consumer2(c *chan int) {
	for v := range *c {
		fmt.Println("out 2 ->", v)
	}
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
