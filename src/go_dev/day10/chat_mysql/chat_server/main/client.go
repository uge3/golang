package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day10/chat_mysql/proto"
	"net"
)

//客户端结构体
type Client struct {
	conn   net.Conn
	userId int
	buf    [8192]byte
}

//读取信息
func (p *Client) readPackage() (msg proto.Message, err error) {
	n, err := p.conn.Read(p.buf[0:4]) //通信协议
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	fmt.Println("read package:", p.buf[0:4])
	var packLen uint32

	packLen = binary.BigEndian.Uint32(p.buf[0:4])
	fmt.Printf("receive len:%d\n", packLen)

	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed\n")
		return
	}
	//获取信息 解压
	fmt.Printf("receive data:%s\n", string(p.buf[0:packLen]))
	fmt.Printf("receive len:%d\n", packLen)
	err = json.Unmarshal(p.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed err:", err)
	}
	return
}

//写信息 发送
func (p *Client) writePackage(data []byte) (err error) {
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(p.buf[0:4], packLen) //整形字符转换

	n, err := p.conn.Write(p.buf[0:4])
	if err != nil {
		fmt.Println("write data failed")
		return
	}

	n, err = p.conn.Write(data)
	if err != nil {
		fmt.Println("write data failed")
		return
	}

	if n != int(packLen) {
		fmt.Println("write data not finished")
		err = errors.New("write data not fninshed")
		return
	}
	return
}

//信息接收
func (p *Client) Process() (err error) {
	//长连接
	for {
		var msg proto.Message
		msg, err = p.readPackage()
		if err != nil {
			clientMgr.DelClient(p.userId) //用户下线，从上线列表中删除
			//通知其他用户，其用户下线
			return err
		}
		err = p.processMsg(msg)
		if err != nil {
			fmt.Println("错误日志：", err)
			continue
			// return
		}
	}
}

//消息处理
func (p *Client) processMsg(msg proto.Message) (err error) {
	switch msg.Cmd {
	case proto.UserLogin:
		err = p.login(msg) //用户登陆
	case proto.UserRegister:
		err = p.register(msg)
	case proto.UserSendMessageCmd:
		err = p.processUserSendMessage(msg) //广播消息
	default:
		err = errors.New("unsupport msessage")
		return
	}
	return
}

//消息发送到用户
func (p *Client) SendMessageToUser(text string, sendId int) {
	var respMsg proto.Message              //创建消息变量
	respMsg.Cmd = proto.UserRecvMessageCmd //信息类型

	var recvMsg proto.UserRecvMessageReq //返回值
	recvMsg.UserId = sendId              //发消息的ID
	recvMsg.Data = text                  //消息内容

	data, err := json.Marshal(recvMsg) //打包 json
	if err != nil {
		fmt.Println("marshal failed,", err)
		return
	}
	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal2 failed,", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed", err)
		return
	}
}

//用户聊天消息 广播
func (p *Client) processUserSendMessage(msg proto.Message) (err error) {
	var userReq proto.UserSendMessageReq             //聊天数据包
	err = json.Unmarshal([]byte(msg.Data), &userReq) //解压数据
	if err != nil {
		return
	}
	users := clientMgr.GetAllUsers() //获取所在线用户
	for id, client := range users {
		if id == userReq.UserId { //自己就不转发
			continue
		}
		client.SendMessageToUser(userReq.Data, userReq.UserId) //发送给对方
	}
	return
}

//登录  错误 返回
func (p *Client) loginResp(err error) {
	var respMsg proto.Message
	respMsg.Cmd = proto.UserLoginRes
	var loginRes proto.LoginCmdRes
	loginRes.Code = 200                //正常返回200
	userMap := clientMgr.GetAllUsers() //获取在线用户
	for userId, _ := range userMap {
		fmt.Println("用户ID:", userId)
		loginRes.User = append(loginRes.User, userId)
	}
	if err != nil {
		loginRes.Code = 500
		loginRes.Error = fmt.Sprintf("%V", err)
	}
	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal failed,", err)
		return
	}
	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal2 failed,", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed", err)
		return
	}
}

//登录方法
func (p *Client) login(msg proto.Message) (err error) {
	defer func() {
		p.loginResp(err)
	}()
	fmt.Printf("recv user login request,data%v", msg)
	var cmd proto.LoginCmd //帐号密码
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	_, err = mgr.Login(cmd.Id, cmd.Passwd) //检验用户密码
	if err != nil {
		return
	}
	clientMgr.AddClient(cmd.Id, p)   //加入在线
	p.userId = cmd.Id                //在线状态
	p.NotifyOthersUserOnline(cmd.Id) //通知在线用户
	return
}

//上下线广播
func (p *Client) NotifyOthersUserOnline(userId int) {
	users := clientMgr.GetAllUsers() //获取所在线用户
	for id, client := range users {
		if id == userId {
			continue
		}
		client.NotifyUserOnline(userId) //通知对方
	}
}

//上线通知到用户
func (p *Client) NotifyUserOnline(userId int) {
	var respMsg proto.Message               //创建消息变量
	respMsg.Cmd = proto.UserStatusNotifyRes //信息类型

	var notifyRes proto.UserStatusNotify //返回值
	notifyRes.UserId = userId            //上下线的ID
	notifyRes.Status = proto.UserOnline  //状态           //正常返回200

	data, err := json.Marshal(notifyRes) //打包 json
	if err != nil {
		fmt.Println("marshal failed,", err)
		return
	}
	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal2 failed,", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed", err)
		return
	}
}

//注册方法
func (p *Client) register(msg proto.Message) (err error) {
	var cmd proto.RegisterCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	err = mgr.Register(&cmd.User)
	if err != nil {
		return
	}
	return
}
