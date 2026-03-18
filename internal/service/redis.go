package service

import (
	"log"

	"github.com/go-redis/redis"
)

var rdbClient *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Println("Redis not available, continuing without cache")
		return
	}

	rdbClient = rdb
}
