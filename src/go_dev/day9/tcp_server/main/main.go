package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("strart server...")
	//设置 连接端口
	listen, err := net.Listen("tcp", "0.0.0.0:5600")
	if err != nil {
		fmt.Println("listen failed,err", err)
		return
	}
	//主进程，接收连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accpet failed,服务器出错", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		//fmt.Println("read:", string(buf))
		//fmt.Printf(strings.TrimSpace(string(buf[0:n])))
		fmt.Printf(string(buf[0:n]))

	}
}
