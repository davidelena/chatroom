package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcessor struct {
	Conn net.Conn
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
	smsTransfer.Code = message.SuccessCode

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
		return
	}
}

func (this *SmsProcessor) SendSmsPrivateMes(mes *message.Message) {
	var smsMes message.SmsMes
	var smsTransfer message.SmsTransferMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendSmsPrivateMes err=", err)
		return
	}

	// find the online users
	var upConn net.Conn
	_, ok := UserManager.OnlineUsers[smsMes.ReceiverId]

	if !ok {
		fmt.Printf("Current Reciver [%d] is not online\n", smsMes.ReceiverId)
		smsTransfer.User = smsMes.User
		smsTransfer.ReceiverId = smsMes.ReceiverId //get receiverId
		smsTransfer.Content = smsMes.Content
		smsTransfer.Code = message.UserNotOnline
		smsTransfer.Error = model.ERROR_USER_NOT_ONLINE.Error()
		// if online not exist need reply this message to the sender
		upConn = this.Conn
	} else {
		smsTransfer.User = smsMes.User
		smsTransfer.ReceiverId = smsMes.ReceiverId //get receiverId
		smsTransfer.Content = smsMes.Content
		smsTransfer.Code = message.SuccessCode
		// if online need reply this message to the receiver
		upConn = UserManager.OnlineUsers[smsMes.ReceiverId].Conn
	}

	data, err := json.Marshal(smsTransfer)
	if err != nil {
		fmt.Println("SendSmsPrivateMes err=", err)
		return
	}

	var mesResp message.Message
	mesResp.Data = string(data)
	mesResp.Type = message.SmsTransferPrivateMesType

	respData, err := json.Marshal(mesResp)
	if err != nil {
		fmt.Println("SendSmsGroupMes Result err=", err)
		return
	}

	this.SendMsgToOnlineUsers(respData, upConn)
}
