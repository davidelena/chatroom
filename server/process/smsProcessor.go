package process

import (
	"chatroom/common/message"
	"chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcessor struct {
}

func (this *SmsProcessor) SendSmsGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	var smsTransfer message.SmsTransferMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendSmsGroupMes err=", err)
		return
	}

	smsTransfer.User = smsMes.User
	smsTransfer.Content = smsMes.Content

	data, err := json.Marshal(smsTransfer)
	if err != nil {
		fmt.Println("SendSmsGroupMes err=", err)
		return
	}

	var mesResp message.Message
	mesResp.Data = string(data)
	mesResp.Type = message.SmsTransferMesType

	respData, err := json.Marshal(mesResp)
	if err != nil {
		fmt.Println("SendSmsGroupMes Result err=", err)
		return
	}

	for id, up := range UserManager.OnlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMsgToOnlineUsers(respData, up.Conn)
	}
}

func (this *SmsProcessor) SendMsgToOnlineUsers(data []byte, conn net.Conn) {
	tf := &utils.Transfer{Conn: conn}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败, err=", err)
	}
}
