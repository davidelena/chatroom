package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte // This is the buffer which will be used in the transfer
}

/**
Read Message data from tcp connection
*/
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端发送的数据...")
	// 只有在conn在没有关闭情况下才会阻塞，如果任意一方关闭则直接不阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	fmt.Println("pkgLen=", pkgLen)
	//Read data about 0~pkgLen from connection into buf bytes array
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	fmt.Println("readN=", n)
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal error=", err)
		return
	}
	return
}

/**
Write Message data into tcp connection
*/
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write header fail", err)
		return
	}

	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write body fail", err)
		return
	}
	return nil
}
