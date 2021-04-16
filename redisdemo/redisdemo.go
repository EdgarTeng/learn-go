package main

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

var (
	redisInstance *redis.Client
)

const (
	RedisTimeout = 10 * time.Millisecond
)

func init() {
	redisInstance = redis.NewClient(&redis.Options{
		DialTimeout:  RedisTimeout,
		ReadTimeout:  RedisTimeout,
		WriteTimeout: RedisTimeout,
		PoolTimeout:  RedisTimeout,
		DB:           1,
	})
}

func main() {
	redisInstance.Set("username", "ken", 1*time.Minute)

	value, err := redisInstance.Get("username").Bytes()

	if err != nil {
		log.Printf("redis error %v", err)
	} else {
		log.Printf("value: %+v", string(value))
	}

}
