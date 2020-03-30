package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection: close\r\n"
	//msg += "User-Agent:Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"
	msg += "\r\n\r\n"
	nw, err := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed,", err)
		return
	}
	fmt.Println("sendto baidu.com bytes:", nw)
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		fmt.Println("count:", count, "err:", err)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}
