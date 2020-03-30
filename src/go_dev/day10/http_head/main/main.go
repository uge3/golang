package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taoboa.com",
}

func main() {
	for _, v := range url {
		//自定义生成客户端 设置超时
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed,err:%\n", v, err)
			continue
		}
		fmt.Printf("head succ ,status:%v\n", resp.Status)
	}
}
