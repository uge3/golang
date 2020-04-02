package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_dev/day10/chat_mysql/common"
	"go_dev/day10/chat_mysql/proto"
	"net"
	"os"
)

//登录
func login(conn net.Conn, userid int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLogin

	var loginCmd proto.LoginCmd
	loginCmd.Id = userid
	loginCmd.Passwd = passwd

	data, err := json.Marshal(loginCmd)
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

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data failed")
		return
	}
	fmt.Println("packlen2:", buf)
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
	msg, err = readPackage(conn)
	fmt.Println("消息解压处理完成！")
	if err != nil {
		fmt.Println("read package failed err:", err)
	}
	fmt.Println("接收到的信息", msg)

	var logingResp proto.LoginCmdRes
	json.Unmarshal([]byte(msg.Data), &logingResp)
	if logingResp.Code == 500 {
		fmt.Println("user not register, start register")
		register(conn, userid, passwd) //进行注册
		os.Exit(0)
	}
	fmt.Println("在线用户：")
	for _, v := range logingResp.User {
		if v == userid {
			continue
		}
		fmt.Printf("user %d\n", v)
		user := &common.User{UserId: v}   //构造一个用户
		onlineUserMap[user.UserId] = user //保存在线用户
	}
	return
}
