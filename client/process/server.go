package process

import (
	"fmt"
	"os"
)

// client -> senderId+receiverId+content
// server 接受信息，根据userId->找到对应连接，发送对应内容
// client -> 接收到信息

func ShowMenu(userId int) {
	for {
		fmt.Printf("---------------恭喜用户%d登录成功---------------\n", userId)
		fmt.Println("---------------1. 显示在线用户列表---------------")
		fmt.Println("---------------2. 发送消息---------------")
		fmt.Println("---------------3. 发送私聊消息---------------")
		fmt.Println("---------------请选择（1-4）:---------------")
		var key int
		var content, privateContent string
		var receiverId int
		smsProcessor := &SmsProcessor{}
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			outputOnlineUsers()
		case 2:
			fmt.Println("请输入你想法送的内容")
			fmt.Scanf("%s\n", &content)
			smsProcessor.SendSmsGroupMes(content)
		case 3:
			fmt.Println("发送私聊消息")
			fmt.Println("请输入你要发送的用户id和内容:")
			fmt.Scanf("%s\n", &content)
			fmt.Sscanf(content, "%d:%s", &receiverId, &privateContent)
			smsProcessor.SendSmsPrivateMes(receiverId, privateContent)
		case 4:
			fmt.Println("你选择退出系统...")
			os.Exit(0)
		default:
			fmt.Println("你输入的选项不正确...")
		}
	}
}
