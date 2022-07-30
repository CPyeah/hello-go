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

	hashOperation()

	setOperation()

	sortedSetOperation()

}

func sortedSetOperation() {
	var zsetKey = "go-ZSet"
	Rdb.ZAdd(Ctx, zsetKey, redis.Z{Member: "Tom", Score: 7}, redis.Z{Member: "Jerry", Score: 3})
	count := Rdb.ZCard(Ctx, zsetKey).Val()
	members := Rdb.ZRandMember(Ctx, zsetKey, int(count)).Val()
	fmt.Println("members is", members)

	s := Rdb.ZScore(Ctx, zsetKey, "Tom").Val()
	fmt.Println("Tom's age is", s)
}

func setOperation() {
	var setKey = "go-Set"
	Rdb.SAdd(Ctx, setKey, "Tom", "Jerry", "Tom", "Tommy")
	Rdb.SRem(Ctx, setKey, "Tommy")
	count := Rdb.SCard(Ctx, setKey).Val()
	fmt.Println(setKey, "has", count, "elements.")
	members := Rdb.SMembers(Ctx, setKey).Val()
	fmt.Println(members)
}

func hashOperation() {
	var hashKey = "go-Hash"
	Rdb.HSet(Ctx, hashKey, "key1", "value1", "key2", "value2")
	var count = Rdb.HLen(Ctx, hashKey).Val()
	fmt.Println(hashKey, "has", count, "key-value paris.")

	value2 := Rdb.HGet(Ctx, hashKey, "key2").Val()
	fmt.Println("key2 's value is", value2)

	Rdb.HDel(Ctx, hashKey, "key1")

	values := Rdb.HGetAll(Ctx, hashKey).Val()
	fmt.Println(values)
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
