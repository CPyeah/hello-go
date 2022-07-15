package main

import (
	"math/rand"
	"time"
)

func main() {
	println(rand.Int31n(11))
	println(rand.Float32())

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		println(rand.Int31n(11))
	}
}
