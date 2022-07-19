package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	mutex()
	condition()
	once()
	syncMap()
	atomicOperation()
}

func atomicOperation() {
	var count int64 = 0
	var wg sync.WaitGroup
	wg.Add(2000)
	for i := 0; i < 1000; i++ {
		go func() {
			//count++
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		go func() {
			//count--
			atomic.AddInt64(&count, -1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("result count is", count)
}

func syncMap() {
	var syncMap sync.Map
	syncMap.Store("k1", "v1")
	syncMap.Store("k2", "v2")
	syncMap.Store("k3", "v3")
	syncMap.Store("k4", "v4")

	fmt.Println(syncMap.Load("k1"))
	fmt.Println(syncMap.LoadAndDelete("k4"))
	fmt.Println(syncMap.LoadOrStore("k5", "v5"))
	fmt.Println(syncMap.LoadOrStore("k5", "v6"))

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, "=", value)
		return true
	})
}

func once() {
	var o sync.Once
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			o.Do(func() {
				fmt.Println("execute once.")
			})
			wg.Done()
		}()
	}
	wg.Wait()
}

func condition() {
	var mutex sync.Mutex
	var cond = sync.NewCond(&mutex)
	var latch sync.WaitGroup
	latch.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			cond.L.Lock()
			cond.Wait()
			fmt.Println("routine", i, "be wake up.")
			cond.L.Unlock()
			latch.Done()
		}(i)
	}
	for i := 0; i < 5; i++ {
		cond.Signal()
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("wake all routine")
	cond.Broadcast()
	latch.Wait()
}

func mutex() {
	var count = 0
	var mutex sync.Mutex
	var latch sync.WaitGroup
	for i := 2; i < 100001; i++ {

		latch.Add(1)
		go func() {
			if isPrimeNum(i) {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
			latch.Done()
		}()

		latch.Wait()
	}

	fmt.Println("count is", count)
}

// 是否是素数
func isPrimeNum(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 && i != num {
			return false
		}
	}
	return true
}
