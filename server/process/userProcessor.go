package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
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

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXIST {
			loginResMes.Code = message.UserNotExist
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PASSWORD {
			loginResMes.Code = message.UserOrPasswordInvalid
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = message.ServerError
			loginResMes.Error = err.Error()
		}
	} else {
		loginResMes.Code = message.SuccessCode
		fmt.Printf("user[%v, %v] login successfully", user.UserId, user.UserName)
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
