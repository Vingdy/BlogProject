package cache

import (
	"github.com/go-redis/redis"
	"log"
	"fmt"
)

var Cache *redis.Client

var SyncEventID chan int

func Init() {
	makeConnect()
	SyncEventID = make(chan int, 1000)
}

func makeConnect() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1" + ":" + "6379",
		Password: "",
		DB:       0,
	})

	result, err := Cache.Ping().Result()
	if err != nil {
		log.Fatal(fmt.Errorf("redis NewClient Ping failed: %v result:%s", err, result))
		return
	}
}

