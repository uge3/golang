package main

import (
	"fmt"

	"go_dev/day11-1/000277_logagent/kafka"
	"go_dev/day11-1/000277_logagent/tailf"

	"github.com/astaxie/beego/logs"
)

func main() { //固定写法，程序的入口

	filename := "./conf/logagent.conf" //配置文件路径
	err := loadConf("ini", filename)   //加载初始化配置文件
	if err != nil {
		fmt.Printf("加载配置文件错误: %v \n", err)
		panic("panic加载配置文件错误")
		return
	} //加载初始化配置文件
	fmt.Println("加载初始化配置文件成功！！")

	err = initLogger() //初始化日志
	if err != nil {
		fmt.Printf("初始化日志失败:%v \n", err)
		panic("panic初始化日志失败")
		return
	} //初始化日志
	fmt.Println("初始化日志成功！！")

	logs.Debug("Debug初始化日志")
	logs.Debug("Debug加载配置文件:%v \n", appConfig)

	//err = tailf.InitTall(appConfig.collectConf)
	err = tailf.InitTall(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("初始化 tail(实时跟踪文件变化) 失败,错误：%v", err)
		return
	}
	fmt.Println("初始化实时跟踪文件变功能成功！！")
	logs.Debug("Debug初始化实时跟踪文件变功能成功!")

	err = kafka.Initkafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("初始化 kafka失败,错误：%v", err)
		return
	}
	fmt.Println("初始化kafka成功！！")
	logs.Debug("初始化kafka成功！！")

	logs.Debug("Debug初始化全部加载成功！")
	/*
		go func() {
			var count int
			for {
				count++
				logs.Debug("测试，循环客户端个数：%d", count)
				time.Sleep(time.Millisecond * 100)
			}
		}() //日志功能调试,注意配置：log_path=D:\project\logs\logagent.log
	*/
	fmt.Println("查看日志文件 log_path=./logs/logagent.log")
	err = serverRun()
	if err != nil {
		logs.Error("服务运行失败,错误：%v", err)
		return
	}

	logs.Info("程序退出!")

}

/*
	b3)启动zookeeper,D:\gorelated\zookeeper\bin>zkServer.cmd
	c4)启动kafka,D:\gorelated\kafka>.\bin\windows\kafka-server-start.bat .\config\server.properties
	必须先保证zookeeper和kafka服务正常启动

	查看消费
	c7)创建消费者（新版本），D:\gorelated\kafka>.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic nginx_log --from-beginning

*/

/*
	fmt.Println("------运行方法,在终端cd(切换)到当前目录：------")
	fmt.Println("注意配置环境变量  新建用户变量 变量名(N)GOPATH  变量值(V)D:/project")

	fmt.Println("编译成二进制的运行方法：d:/project/src/go_dev/day11/000277_logagent> ")
	fmt.Println("go build go_dev/day11/000277_logagent/main")
	fmt.Println("在终端输入：d:/project/src/go_dev/day11/000277_logagent> main.exe ")
	fmt.Println("------退出程序！------")

*/
