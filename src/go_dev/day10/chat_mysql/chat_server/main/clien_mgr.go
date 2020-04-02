package main

import "fmt"

//客户端管理 结构体
type ClientMgr struct {
	onlineUsers map[int]*Client
}

var (
	clientMgr *ClientMgr
)

//初始化
func init() {
	clientMgr = &ClientMgr{
		onlineUsers: make(map[int]*Client, 1024),
	}
}

//方法 一
func (p *ClientMgr) AddClient(userId int, client *Client) {
	p.onlineUsers[userId] = client
}

//方法二
func (p *ClientMgr) GetClient(userId int) (client *Client, err error) {
	client, ok := p.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("user %d not exists", userId)
		return
	}
	return
}

//返回在线人员
func (p *ClientMgr) GetAllUsers() map[int]*Client {
	return p.onlineUsers
}

//人员下线
func (p *ClientMgr) DelClient(userId int) {
	delete(p.onlineUsers, userId)
}
