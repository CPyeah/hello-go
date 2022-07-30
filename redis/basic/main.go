package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

var Ctx context.Context
var Rdb *redis.Client

func main() {
	defer func(Rdb *redis.Client) {
		err := Rdb.Close()
		if err != nil {
			panic(err)
		}
	}(Rdb)

	stringOperation()

	listOperation()

}

func listOperation() {
	var key = "go-List"
	count, err := Rdb.RPush(Ctx, key, "Tom", "Jerry").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, "has", count, "elements.")

	count = Rdb.LLen(Ctx, key).Val()
	fmt.Println(key, "has", count, "elements.")

	items, _ := Rdb.RPopCount(Ctx, key, 1).Result()
	fmt.Println(items)

	t := Rdb.LRem(Ctx, key, 1, "Tom").Val()
	fmt.Println("removed", t, "items.")
}

func stringOperation() {
	// have
	result, err := Rdb.Get(Ctx, "hello").Result()
	if err == redis.Nil {
		fmt.Println("key hello not exist.")
	} else {
		fmt.Println("key hello exist, value is", result)
	}

	// save
	Rdb.Set(Ctx, "hello", "go redis world", 0)

	// save with expiration time
	Rdb.Set(Ctx, "hello B", "go redis world", time.Minute*5)

	// get
	value := Rdb.Get(Ctx, "hello").String()
	fmt.Println(value)

	// delete
	Rdb.Del(Ctx, "hello", "hello B")

	// have
	result, err = Rdb.Get(Ctx, "hello").Result()
	if err == redis.Nil {
		fmt.Println("key hello not exist.")
	} else {
		fmt.Println("key hello exist, value is", result)
	}
}

func init() {
	Ctx = context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})

	pong, err := Rdb.Ping(Ctx).Result()
	if err == nil {
		fmt.Println("connect success", pong)
	} else {
		fmt.Println("connect fail")
		panic(err)
	}
}
