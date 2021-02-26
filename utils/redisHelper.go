package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ctx = context.Background()
)

func RedisPingpong() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func GetRedisClient() *redis.Client {
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

func GetRedisClientPool() *redis.Client {
	redisOptions := &redis.Options{
		Addr:         "localhost:6379",
		DB:           0,
		PoolSize:     12,
		MinIdleConns: 2,
		IdleTimeout:  time.Second,
	}
	redis.NewClient()
}
