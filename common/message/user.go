package message

import "net"

type User struct {
	UserId     int    `json:"userId"`
	ReceiverId int    `json:receiverId`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	Status     int    `json:"status"`
}

type CurrUser struct {
	User
	Conn net.Conn
}
