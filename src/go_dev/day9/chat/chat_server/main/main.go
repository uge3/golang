package main

import "time"

func main() {
	initRedis("47.107.49.152:6379", 16, 1004, time.Second*300) //初始化redis
	initUserMgr()                                              //初始化用户
	runServer("0.0.0.0:10000")                                 //运行服务端

}
