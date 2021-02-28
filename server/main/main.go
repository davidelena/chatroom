package main

import (
	"chatroom/server/model"
	process "chatroom/server/process"
	"chatroom/utils"
	"context"
	"fmt"
	"io"
	"net"
)

const (
	Network = "tcp"
	Address = "127.0.0.1:8081"
)

func initConfig(ctx context.Context) {
	utils.InitRedisClient()
	model.MyUserDao = model.NewUserDao(ctx, utils.RedisCli)
}

func main() {
	ctx := context.Background()
	fmt.Println("服务端在8081端口监听...")
	listener, err := net.Listen(Network, Address)
	if err != nil {
		fmt.Println(err)
		return
	}
	initConfig(ctx)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("读取连接出错...")
			return
		}
		fmt.Printf("接受客户端连接[%v]...", conn.RemoteAddr())
		go routeProcess(conn)
	}
}

func routeProcess(conn net.Conn) {
	fmt.Println("处理连接开始...")
	defer conn.Close()
	processor := &process.Processor{Conn: conn}
	err := processor.RouteProcess()
	if err != io.EOF && err != nil {
		fmt.Println("服务端和客户端通讯错误err=", err)
		return
	}
}
