package proto

import "go_dev/day10/chat_mysql/common"

//通讯协议
//通讯消息
type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

//帐号密码结构体
type LoginCmd struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

//注册 帐号 结构体
type RegisterCmd struct {
	User common.User `json:"user"`
}

//登录结构体
type LoginCmdRes struct {
	Code  int    `json:"code"`
	User  []int  `json:"users"`
	Error string `json:"error"`
}

//用户在线协议
type UserListReq struct {
	Id int `json:"user_id"`
}

//用户列表
type UserListRes struct {
	Id []int `json:"users"`
}

//上下线状态  用于广播
type UserStatusNotify struct {
	UserId int `json:"user_id"`
	Status int `json:"user_status"`
}

//发送消息
type UserSendMessageReq struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}

//接收消息
type UserRecvMessageReq struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}
