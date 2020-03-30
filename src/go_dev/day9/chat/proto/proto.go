package proto

import "go_dev/day9/chat/chat_server/model"

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
	User model.User `json:"user"`
}

//登录结构体
type LoginCmdRes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
