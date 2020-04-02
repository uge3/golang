package main

import (
	"fmt"
	"go_dev/day10/chat_mysql/common"
	"go_dev/day10/chat_mysql/proto"
)

var onlineUserMap map[int]*common.User = make(map[int]*common.User, 16)

func outputUserOnline() {
	fmt.Println("在线用户 ：")
	for id, _ := range onlineUserMap {
		if id == userId {
			continue
		}
		fmt.Println("用户ID:", id)
	}
}

//更新用户状态
func updateUserStatus(userStatus proto.UserStatusNotify) {
	user, ok := onlineUserMap[userStatus.UserId]
	if !ok {
		//fmt.Println("user id not exists:", userStatus.UserId)
		user = &common.User{}
		user.UserId = userStatus.UserId //创建一个用户状态
		user.Status = userStatus.Status //更新状态

	}
	user.Status = userStatus.Status //服务端返回的状态
	onlineUserMap[user.UserId] = user
	outputUserOnline()
}
