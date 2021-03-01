package model

import (
	"chatroom/common/message"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	Ctx    context.Context
	Client *redis.Client
}

func NewUserDao(ctx context.Context, client *redis.Client) *UserDao {
	MyUserDao = &UserDao{
		Ctx:    ctx,
		Client: client,
	}
	return MyUserDao
}

func (this *UserDao) GetUserById(id int) (user *message.User, err error) {
	cmd := this.Client.HGet(this.Ctx, "users", strconv.Itoa(id))
	if cmd.Err() != nil && cmd.Err() == redis.Nil {
		err = ERROR_USER_NOT_EXIST
		return
	}
	res, err := cmd.Result()
	if err != nil {
		err = ERROR_INTERNAL_SERVER_ERROR
		return
	}

	jsonErr := json.Unmarshal([]byte(res), &user)
	if jsonErr != nil {
		err = ERROR_INTERNAL_SERVER_ERROR
		return
	}
	return
}

func (this *UserDao) AddUser(user *message.User) (err error) {
	_, err = this.GetUserById(user.UserId)
	if err == ERROR_USER_NOT_EXIST {
		jsonRes, _ := json.Marshal(user)
		cmd := this.Client.HSet(this.Ctx, "users", user.UserId, string(jsonRes))
		intRes, err := cmd.Result()
		if err != nil {
			err = ERROR_INTERNAL_SERVER_ERROR
			return err
		}
		fmt.Printf("cmd:%v\n", intRes)
	} else {
		err = ERROR_USER_EXISTS
		return err
	}

	return nil
}

func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	user, err = this.GetUserById(userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PASSWORD
		return
	}
	return user, nil
}

func (this *UserDao) Register(user *message.User) (err error) {
	return this.AddUser(user)
}
