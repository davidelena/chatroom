package utils

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRedisKVCRUD(t *testing.T) {
	ctx := context.Background()
	redisClient := GetRedisClient()
	key := "redis_name"
	statusCmd := redisClient.Set(ctx, key, "davidelena", time.Minute)
	fmt.Println(statusCmd)

	value := redisClient.Get(ctx, key)
	fmt.Println(value)

	fmt.Println("after delete operation")
	statusCmdDel := redisClient.Del(ctx, key)
	fmt.Println(statusCmdDel)

	value2 := redisClient.Get(ctx, key)
	fmt.Println(value2)
}

func TestRedisHashCRUD(t *testing.T) {
	ctx := context.Background()
	redisClient := GetRedisClient()
	key := "redis_name"
	statusCmd := redisClient.HSet(ctx, key, "davidelena", time.Minute)
	fmt.Println(statusCmd)

	value := redisClient.Get(ctx, key)
	fmt.Println(value)

	fmt.Println("after delete operation")
	statusCmdDel := redisClient.Del(ctx, key)
	fmt.Println(statusCmdDel)

	value2 := redisClient.Get(ctx, key)
	fmt.Println(value2)
}
