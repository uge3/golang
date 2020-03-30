package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day9/chat/proto"
	"net"
)

// type Message struct {
// 	Cmd  string `json:"cmd"`
// 	Data string `json:"data"`
// }

// //帐号密码结构体
// type LoginCmd struct {
// 	Id     int    `json:"user_id"`
// 	Passwd string `json:"passwd"`
// }

//读取信息
func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var buf [8192]byte
	n, err := conn.Read(buf[0:4]) //通信协议
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	fmt.Println("read package:", buf[0:4])

	var packLen uint32

	packLen = binary.BigEndian.Uint32(buf[0:4])
	fmt.Printf("receive len:%d\n", packLen)

	n, err = conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	//获取信息 解压
	fmt.Printf("receive data:%s\n", string(buf[0:packLen]))
	fmt.Printf("receive len2:%d\n", packLen)
	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed err", err)
	}
	return
}

func login(conn net.Conn) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLogin

	var loginCmd proto.LoginCmd
	loginCmd.Id = 1
	loginCmd.Passwd = "123456789"

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

	//buffer := bytes.NewBuffer(buf[:])
	// err = binary.Write(buffer, binary.BigEndian, packLen)
	// if err != nil {
	// 	fmt.Println("write package len failed")
	// 	return
	// }
	// fmt.Println("packlen:", buffer.Bytes())
	// n, err := conn.Write(buffer.Bytes())

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
	if err != nil {
		fmt.Println("read package failed err:", err)
	}
	fmt.Println()
	return
}

func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	err = login(conn)
	if err != nil {
		fmt.Println("login frailed", err)
		return
	}
	/*
		defer conn.Close()
		inputReader := bufio.NewReader(os.Stdin)
		for {
			input, _ := inputReader.ReadString('\n')
			trimmedInput := strings.Trim(input, "\r\n")
			if trimmedInput == "Q" {
				return
			}
			_, err = conn.Write([]byte(trimmedInput))
			if err != nil {
				return
			}
		}
	*/
}
