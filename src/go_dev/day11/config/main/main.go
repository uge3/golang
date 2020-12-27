package main

import (
	"fmt"

	"github.com/astaxie/beego/config" //读取配置文件
)

func main() {

	conf, err := config.NewConfig("ini", "./config.conf") //读取一个新的配置文件

	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}
	fmt.Println("conf内容", conf)
	port, err := conf.Int("server::port") //读取对应的端口

	if err != nil {
		fmt.Println("read server:port failed,err", err)
		return
	}
	fmt.Println("port:", port)

	log_level := conf.String("logs::log_level") //String 只有一个返回值

	fmt.Println("log_level:", log_level)

	log_path := conf.String("collect::log_path")

	fmt.Println("log_path:", log_path)

}
