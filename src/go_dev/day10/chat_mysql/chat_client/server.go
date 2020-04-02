package main

import (
	"encoding/json"
	"fmt"
	"go_dev/day10/chat_mysql/proto"
	"net"
	"os"
)

//接收消息
func processServerMessage(conn net.Conn) {
	for {
		msg, err := readPackage(conn)
		fmt.Println("与服务器交互数据中。。。")
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(0)
		}
		var userStatus proto.UserStatusNotify
		err = json.Unmarshal([]byte(msg.Data), &userStatus) //反序列化
		if err != nil {
			fmt.Println("unmarshal failed,err", err)
			return
		}
		fmt.Println("消息类型：", msg.Cmd)
		switch msg.Cmd {
		case proto.UserStatusNotifyRes: //用户上线
			fmt.Println("数据执行中！")
			updateUserStatus(userStatus) //更新用户状态
		case proto.UserRecvMessageCmd: //接收其他用户的消息
			recvMessageFromServer(msg) //进行消息的处理

		}
	}
}

//接收其他用户的消息处理
func recvMessageFromServer(msg proto.Message) {
	var recvMsg proto.UserRecvMessageReq
	err := json.Unmarshal([]byte(msg.Data), &recvMsg)
	if err != nil {
		fmt.Println("unmarshal failed,err:", err)
		return
	}
	msgChan <- recvMsg //放入管道
	listUnReadMsg()    //消息显示
}
