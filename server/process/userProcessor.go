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
	UserId int
	Conn   net.Conn
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
		// login userId set to the this.UserId
		this.UserId = loginMes.UserId
		UserManager.AddOrUpdateOnlineUser(this)
		// notify other online users current login user is online
		this.NotifyOtherOnlineUsers(loginMes.UserId)
		for id, _ := range UserManager.OnlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Printf("user[%v, %v] login successfully\n", user.UserId, user.UserName)
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

func (this *UserProcessor) NotifyOtherOnlineUsers(userId int) {
	userMap := UserManager.OnlineUsers
	for uid, up := range userMap {
		if uid == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcessor) NotifyMeOnline(userId int) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	msg.Data = string(data)

	msgData, msgErr := json.Marshal(msg)
	if msgErr != nil {
		fmt.Println("json.Marshal err=", msgErr)
		return
	}

	transfer := &utils.Transfer{
		Conn: this.Conn,
	}

	transferErr := transfer.WritePkg(msgData)
	if transferErr != nil {
		fmt.Println("transfer.WritePkg err=", transferErr)
		return
	}

}

func (this *UserProcessor) ServerProcessRegister(msg *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(msg.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal loginMes error=", err)
		return err
	}

	var registerResMes message.LoginResMes

	err = model.MyUserDao.Register(registerMes.UserVO)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = message.UserRegisterExisted
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = message.ServerError
			registerResMes.Error = err.Error()
		}
	} else {
		registerResMes.Code = message.SuccessCode
		fmt.Printf("user[%v, %v] login successfully", registerMes.UserVO.UserId, registerMes.UserVO.UserName)
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) is error=", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.RegisterResMesType
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
