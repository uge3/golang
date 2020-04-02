package main

import (
	"fmt"
	"go_dev/day10/chat_mysql/proto"
	"net"
)

var userId int
var passwd string
var conns net.Conn
var err error

var msgChan chan proto.UserRecvMessageReq //接收消息的管道
//管理初始化
func init() {
	msgChan = make(chan proto.UserRecvMessageReq, 500)
}

func main() {
	conns, err = net.Dial("tcp", "localhost:10000")
	//fmt.Println(conns, "conns")
	//conn, err := net.Dial("tcp", "localhost:10000")
	//fmt.Println(conn, "conn")
	//if conns == conn {
	//	fmt.Println("conns == conn")
	//}
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	tag := true
	for tag {
		fmt.Println("请输入用户名、密码，用空格隔开\n")
		fmt.Scanf("%d %s\n", &userId, &passwd)
		err = login(conns, userId, passwd)
		if err != nil {
			fmt.Println()
			fmt.Println("login frailed", err)
			continue
		}
		tag = false
	}

	//outputUserOnline()            //在线用户
	go processServerMessage(conns) //与服务器交互
	for {
		logic() //逻辑处理
	}

}
