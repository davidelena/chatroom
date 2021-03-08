package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	var transferMes message.SmsTransferMes
	err := json.Unmarshal([]byte(mes.Data), &transferMes)
	if err != nil {
		fmt.Println("outputGroupMes json.Unmarshal err=", err)
		return
	}

	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", transferMes.UserId, transferMes.Content)
	fmt.Println(info)
	fmt.Println()
}

func outputPrivateMes(mes *message.Message) {
	var transferMes message.SmsTransferMes
	err := json.Unmarshal([]byte(mes.Data), &transferMes)
	if err != nil {
		fmt.Println("outputGroupMes json.Unmarshal err=", err)
		return
	}

	if transferMes.Code != message.SuccessCode {
		info := fmt.Sprintf("用户id:\t%d发送的消息未能被接受，错误原因:%s", transferMes.UserId, transferMes.Error)
		fmt.Println(info)
		fmt.Println()
	} else {
		info := fmt.Sprintf("用户id:\t%d对你说:\t%s", transferMes.UserId, transferMes.Content)
		fmt.Println(info)
		fmt.Println()
	}

}
