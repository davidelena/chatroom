package process

import (
	"chatroom/common/message"
	"chatroom/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(msg *message.Message) (err error) {
	fmt.Println("服务端开始处理消息")
	switch msg.Type {
	case message.LoginMesType:
		userProcessor := &UserProcessor{Conn: this.Conn}
		err = userProcessor.ServerProcessLogin(msg)
	case message.RegisterMesType:
		userProcessor := &UserProcessor{Conn: this.Conn}
		err = userProcessor.ServerProcessRegister(msg)
	case message.SmsMesType:
		smsProcessor := &SmsProcessor{}
		smsProcessor.SendSmsGroupMes(msg)
	default:
		fmt.Println("消息类型不正确，请确认...")
	}
	return
}

func (this *Processor) RouteProcess() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return err
			} else {
				fmt.Println("readPkg error=", err)
				return err
			}
		}
		fmt.Println("msg=", msg)
		err = this.serverProcessMes(&msg)
		if err != nil {
			return err
		}
	}
}
