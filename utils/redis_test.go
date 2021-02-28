package utils

import (
	"chatroom/server/model"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisKVCRUD(t *testing.T) {
	ctx := context.Background()
	redisCli := GerRedisClient()
	key := "redis_name"
	statusCmd := redisCli.Set(ctx, key, "davidelena", time.Minute)
	fmt.Println(statusCmd)

	value := redisCli.Get(ctx, key)
	fmt.Println(value)

	fmt.Println("after delete operation")
	statusCmdDel := redisCli.Del(ctx, key)
	fmt.Println(statusCmdDel)

	value2 := redisCli.Get(ctx, key)
	fmt.Println(value2)
}

func TestRedisHashCRUD(t *testing.T) {
	ctx := context.Background()
	redisCli := GerRedisClient()
	defer redisCli.Close()
	key := "redis_name"
	statusCmd := redisCli.HSet(ctx, key, "davidelena", time.Minute)
	fmt.Println(statusCmd)

	value := redisCli.Get(ctx, key)
	fmt.Println(value)

	fmt.Println("after delete operation")
	statusCmdDel := redisCli.Del(ctx, key)
	fmt.Println(statusCmdDel)

	value2 := redisCli.Get(ctx, key)
	fmt.Println(value2)
}

func TestLoginSuccess(t *testing.T) {
	ctx := context.Background()
	redisCli := GerRedisClient()
	defer redisCli.Close()
	userDao := model.NewUserDao(ctx, redisCli)

	user1, err := userDao.Login(100, "123456")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user1)
}

func TestLoginUserNotExist(t *testing.T) {
	ctx := context.Background()
	redisCli := GerRedisClient()
	defer redisCli.Close()
	userDao := model.NewUserDao(ctx, redisCli)

	user1, err := userDao.Login(200, "123456")
	t.Log(user1)
	if err != nil {
		assert.EqualError(t, err, model.ERROR_USER_NOT_EXIST.Error())
		return
	}
}

func TestLoginUserPwdIncorrect(t *testing.T) {
	ctx := context.Background()
	redisCli := GerRedisClient()
	defer redisCli.Close()
	userDao := model.NewUserDao(ctx, redisCli)

	user1, err := userDao.Login(100, "12345")
	if err != nil {
		assert.EqualError(t, err, model.ERROR_USER_PASSWORD.Error())
		t.Log(user1)
		return
	}
}
