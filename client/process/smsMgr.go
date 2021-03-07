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
