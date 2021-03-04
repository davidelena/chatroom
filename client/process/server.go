package process

import (
	"fmt"
	"os"
)

func ShowMenu(userId int) {
	for {
		fmt.Printf("---------------恭喜用户%d登录成功---------------\n", userId)
		fmt.Println("---------------1. 显示在线用户列表---------------")
		fmt.Println("---------------2. 发送消息---------------")
		fmt.Println("---------------3. 信息列表---------------")
		fmt.Println("---------------请选择（1-4）:---------------")
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			//fmt.Println("显示在线用户列表")
			outputOnlineUsers()
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("显示信息列表")
		case 4:
			fmt.Println("你选择退出系统...")
			os.Exit(0)
		default:
			fmt.Println("你输入的选项不正确...")
		}
	}
}
