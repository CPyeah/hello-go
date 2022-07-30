package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

var Ctx context.Context
var Rdb *redis.Client

func main() {
	var value, err = Rdb.Get(Ctx, "hello").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("hello", value)
}

func init() {
	Ctx = context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})
}
