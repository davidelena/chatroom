package process

import (
	"chatroom/common/message"
	"chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	Conn net.Conn
}

func (this *UserProcessor) ServerProcessLogin(msg *message.Message) (err error) {
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
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	transfer.WritePkg(data)
	return
}

func (this *UserProcessor) ServerProcessRegister(msg *message.Message) (err error) {
	return nil
}
