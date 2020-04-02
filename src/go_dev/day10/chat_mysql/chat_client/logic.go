package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_dev/day10/chat_mysql/proto"
	"net"
	"os"
)

//聊天消息处理 发送
func sendTextMseeage(conns net.Conn, text string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserSendMessageCmd //发消息的结构体

	var sendReq proto.UserSendMessageReq //发消息的协议
	sendReq.Data = text
	sendReq.UserId = userId

	data, err := json.Marshal(sendReq)
	if err != nil {
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))

	fmt.Println("packlen1:", packLen)
	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conns.Write(buf[:]) //写入消息
	if err != nil || n != 4 {
		fmt.Println("write data failed")
		return
	}
	fmt.Println("packlen2:", buf)
	_, err = conns.Write([]byte(data))
	if err != nil {
		return
	}
	return
}

//聊天对话
func enterTalk() {
	//var destUserId int
	var msg string
	fmt.Println("请输入：\n")
	fmt.Scanf("%s\n", &msg)
	sendTextMseeage(conns, msg) //发送消息
	fmt.Println("++++++++++++++++++")

}

//他人消息显示
func listUnReadMsg() {
	select {
	case msg := <-msgChan:
		fmt.Println(msg.UserId, ":", msg.Data)
	default:
		return
	}
}

//功能选项
func enterMenu() {
	fmt.Println("1,list online user")
	fmt.Println("2,talk")
	fmt.Println("3,list message")
	fmt.Println("4,exit")
	var sel int
	fmt.Scanf("%d\n", &sel)
	switch sel {
	case 1:
		outputUserOnline()
	case 2:
		enterTalk()
	case 3:
		listUnReadMsg()
	case 4:
		os.Exit(0)
	}
}
func logic() {
	enterMenu()
}
