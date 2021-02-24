package main

import (
	util "chatroom/common"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

const (
	Network = "tcp"
	Address = "127.0.0.1:8081"
)

func main() {
	fmt.Println("服务端在8081端口监听...")
	listener, err := net.Listen(Network, Address)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("读取连接出错...")
			return
		}
		fmt.Printf("接受客户端连接[%v]...", conn.RemoteAddr())
		go process(conn)
	}
}

func serverProcessLogin(conn net.Conn, msg *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(msg.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal loginMes error=", err)
		return err
	}

	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = message.SuccessCode
	} else {
		loginResMes.Code = message.UserOrPasswordInvalid
		loginResMes.Error = "用户名或密码错误"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) is error=", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.LoginResMesType
	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal(resMsg) error=", err)
	}
	util.WritePkg(conn, data)
	return
}

func serverProcessRegister(conn net.Conn, msg *message.Message) (err error) {
	return nil
}

func serverProcessMes(conn net.Conn, msg *message.Message) (err error) {
	fmt.Println("服务端开始处理消息")
	switch msg.Type {
	case message.LoginMesType:
		err = serverProcessLogin(conn, msg)
	case message.RegisterMesType:
		err = serverProcessRegister(conn, msg)
	default:
		fmt.Println("消息类型不正确，请确认...")
	}
	return
}

func process(conn net.Conn) {
	fmt.Println("处理连接开始...")
	defer conn.Close()
	for {
		msg, err := util.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return
			} else {
				fmt.Println("readPkg error=", err)
				return
			}
		}
		fmt.Println("msg=", msg)
		err = serverProcessMes(conn, &msg)
		if err != nil {
			return
		}
	}
}
