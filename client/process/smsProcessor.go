package process

import (
	"chatroom/common/message"
	"chatroom/utils"
	"encoding/json"
	"fmt"
)

type SmsProcessor struct {
}

func (this *SmsProcessor) SendSmsGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.UserId = CurrentUser.UserId
	smsMes.Content = content
	smsMes.Status = message.UserOnline

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendSmsGroupMes json.Marshal err=", err)
		return err
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendSmsGroupMes json.Marshal err=", err)
		return err
	}

	transfer := &utils.Transfer{
		Conn: CurrentUser.Conn,
	}

	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("SendSmsGroupMes json.Marshal err=", err)
		return err
	}

	return nil
}
