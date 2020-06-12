package cache

import (
	"conf"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var Cache *redis.Client

var SyncEventID chan int

func Init() {
	makeConnect()
	SyncEventID = make(chan int, 1000)
}

func makeConnect() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     conf.App.RedisHost + ":" + conf.App.RedisPort,
		Password: conf.App.RedisPwd,
		DB:       conf.App.RedisDB,
	})

	result, err := Cache.Ping().Result()
	if err != nil {
		log.Fatal(fmt.Errorf("redis NewClient Ping failed: %v result:%s", err, result))
		return
	}
}
