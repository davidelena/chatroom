package process

import (
	"chatroom/common/message"
	"fmt"
)

var (
	OnlineUsers = make(map[int]*message.User, 10)
	CurrentUser message.CurrUser
)

func outputOnlineUsers() {
	fmt.Println("当前在线用户列表")
	for uid, _ := range OnlineUsers {
		fmt.Printf("用户id[%d]\n", uid)
	}
}

func updateUserStatus(mes *message.NotifyUserStatusMes) {
	user, ok := OnlineUsers[mes.UserId]
	if !ok {
		user = &message.User{
			UserId: mes.UserId,
		}
	}
	user.Status = mes.Status
	OnlineUsers[mes.UserId] = user
	outputOnlineUsers()
}
