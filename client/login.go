package main

import (
	util "chatroom/common"
	message "chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

func login(userId int, userPwd string) (err error) {
	fmt.Printf("userId: %d, userPwd: %s\n", userId, userPwd)

	// Dial to special address
	conn, err := net.Dial(Network, Address)
	if err != nil {
		fmt.Printf("连接服务器[%v]无响应...\n", Address)
		return
	}
	defer conn.Close()

	// Create Message Data
	var mes message.Message
	mes.Type = message.LoginMesType
	loginMes := message.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}
	// LoginMes serialization
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal loginMes error", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal mes error", err)
		return
	}

	// Write the data length first
	util.WritePkg(conn, data)
	fmt.Printf("客户端发送消息成功, 发送长度:%v, 发送内容:%v\n", len(data), string(data))

	// Read the response data from connection
	msg, err := util.ReadPkg(conn)
	if err != nil {
		fmt.Println("客户端获取响应结果出错")
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(msg.Data), &loginResMes)
	if err != nil {
		fmt.Println("获取登录结果出错, error=", err)
		return
	}

	if loginResMes.Code == 200 {
		fmt.Println("登录成功")

	} else if loginResMes.Code == 401 {
		fmt.Println("登录失败，用户名或密码不正确")
	} else {
		fmt.Println("系统异常")
	}
	return nil
}
