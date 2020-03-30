package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//客户端结构体
type Client struct {
	conn net.Conn
	buf  [8192]byte
}

//读取信息
func (p *Client) readPackage() (msg Message, err error) {
	n, err := p.conn.Read(p.buf[0:4]) //通信协议
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	fmt.Println("read package:", p.buf[0:4])
	//buffer := bytes.NewBuffer(p.buf[0:8])
	var packLen uint32
	//err = binary.Read(buffer, binary.BigEndian, &packLen)

	packLen = binary.BigEndian.Uint32(p.buf[0:4])
	fmt.Printf("receive len:%d", packLen)

	// if err != nil {
	// 	fmt.Println("read package len failed")
	// 	return
	// }

	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	//获取信息 解压
	fmt.Printf("receive data:%s\n", string(p.buf[0:packLen]))
	fmt.Printf("receive len:%d\n", packLen)
	err = json.Unmarshal(p.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed err", err)
	}
	return
}

//写信息 发送
func (p *Client) writePackage(data []byte) (err error) {
	packLen := uint32(len(data))
	//buffer := bytes.NewBuffer(p.buf[0:4])
	binary.BigEndian.PutUint32(p.buf[0:4], packLen) //整形字符转换
	// err = binary.Write(buffer, binary.BigEndian, packLen) //通信协议
	// if err != nil {
	// 	fmt.Println("write package len failed")
	// 	return
	// }
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
		var msg Message
		msg, err = p.readPackage()
		if err != nil {
			return err
		}
		err = p.processMsg(msg)
		if err != nil {
			return
		}
	}
}

//消息处理  
func (p *Client) processMsg(msg Message) (err error) {
	switch msg.Cmd {
	case UserLogin:
		err = p.login(msg)
	case UserRegister:
		err = p.register(msg)
	default:
		err = errors.New("unsupport msessage")
		return
	}
	return
}

//登录错误 返回
func (p *Client) loginResp(err error) {
	var respMsg Message
	respMsg.Cmd = UserLoginRes
	var loginRes LoginCmdRes
	loginRes.Code = 200
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
func (p *Client) login(msg Message) (err error) {
	defer func() {
		p.loginResp(err)
	}()
	fmt.Printf("recv user login request,data%v", msg)
	var cmd LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	_, err = mgr.Login(cmd.Id, cmd.Passwd) //
	if err != nil {
		return
	}
	return
}

//注册方法
func (p *Client) register(msg Message) (err error) {
	var cmd RegisterCmd
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
