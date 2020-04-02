package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day10/chat_mysql/proto"
	"net"
)

//读取信息
func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var buf [8192]byte
	n, err := conn.Read(buf[0:4]) //通信协议
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	//fmt.Println("read package:", buf[0:4])

	var packLen uint32
	packLen = binary.BigEndian.Uint32(buf[0:4])
	//fmt.Printf("receive len:%d\n", packLen)

	n, err = conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	//获取信息 解压
	fmt.Printf("receive data:%s\n", string(buf[0:packLen]))
	fmt.Printf("receive len2:%d\n", packLen)
	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed err", err)
	}
	return
}
