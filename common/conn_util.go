package util

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

/**
Read Message data from tcp connection
*/
func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	// 只有在conn在没有关闭情况下才会阻塞，如果任意一方关闭则直接不阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	fmt.Println("pkgLen=", pkgLen)
	//Read data about 0~pkgLen from connection into buf bytes array
	n, err := conn.Read(buf[:pkgLen])
	fmt.Println("readN=", n)
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal error=", err)
		return
	}
	return
}

/**
Write Message data into tcp connection
 */
func WritePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write header fail", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write body fail", err)
		return
	}
	return nil
}
