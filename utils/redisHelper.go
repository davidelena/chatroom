package utils

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	RedisCli *redis.Client
)

func InitRedisClient() *redis.Client {
	redisOptions := &redis.Options{
		Addr:         "localhost:6379",
		DB:           0,
		PoolSize:     12,
		MinIdleConns: 2,
		IdleTimeout:  time.Second,
	}
	RedisCli = redis.NewClient(redisOptions)
	return RedisCli
}

func GerRedisClient() *redis.Client {
	redisOptions := &redis.Options{
		Addr:         "localhost:6379",
		DB:           0,
		PoolSize:     12,
		MinIdleConns: 2,
		IdleTimeout:  time.Second,
	}
	client := redis.NewClient(redisOptions)
	return client
}
