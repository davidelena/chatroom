package process

import (
	"chatroom/common/message"
	"chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

const (
	Network = "tcp"
	Address = "127.0.0.1:8081"
)

type UserProcessor struct {
}

func (this *UserProcessor) Login(userId int, userPwd string) (err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	fmt.Printf("客户端发送消息成功, 发送长度:%v, 发送内容:%v\n", len(data), string(data))

	// Read the response data from connection
	msg, err := tf.ReadPkg()
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
		//隐藏启动goroutine保持和服务端的通讯，如果服务端有数据推送给客户端需要保持联系
		go serverProcessMes(conn)
		//显示菜单
		ShowMenu()
	} else if loginResMes.Code == 401 {
		fmt.Println("登录失败，用户名或密码不正确")
	} else if loginResMes.Code == 402 {
		fmt.Println("用户不存在")
	} else {
		fmt.Println("系统异常")
	}
	return nil
}

func serverProcessMes(conn net.Conn) {
	transfer := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("服务器端readPkg出错err=", err)
			return
		}
		fmt.Printf("mes=%v", mes)
	}
}
