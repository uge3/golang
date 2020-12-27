package main

import (
	"fmt"
	"go_dev/day11/logagent/kafka"
	"go_dev/day11/logagent/tailf"
	"time"

	"github.com/astaxie/beego/logs" //日志管理
)

func main() {
	filename := "./conf/logagent.conf" //读配置
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("加载配置文件错误: %v \n", err)
		panic("panic加载配置文件错误")
		return
	}
	fmt.Println("加载初始化配置文件成功！！")
	fmt.Println("========|filenmae====", filename, loadCollectConf)

	err = initLogger() //初始化日志
	if err != nil {
		fmt.Printf("load logger failed err :%v\n", err)
		panic("load logger failed")
		return
	}
	logs.Debug("initalize succ")
	logs.Debug("load conf succ,config:%v", appConfig)

	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed,err:%v", err)
		return
	}

	err = kafka.InitKafka(appConfig.kafkaAddr) //kafka 初始化
	logs.Debug("initalize all succ")
	if err != nil {
		logs.Error("init tail failed,err:%v", err)
		return
	}
	logs.Debug("initalize all succ")
	fmt.Println("========|||||||||||||||||||====")
	/*
		err = kafka.Initkafka(appConfig.kafkaAddr)
		if err != nil {
			logs.Error("初始化 kafka失败,错误：%v", err)
			return
		}
		fmt.Println("初始化kafka成功！！")
		logs.Debug("初始化kafka成功！！")

		logs.Debug("Debug初始化全部加载成功！")
	*/
	go func() {
		var count int
		for {
			count++
			logs.Debug("test for logger %d", count)
			time.Sleep(time.Microsecond * 10)
		}
	}()

	fmt.Println("查看日志文件 log_path=./logs/logagent.log")
	err = serverRun()
	fmt.Println("====33333333====")
	if err != nil {
		logs.Error("serverRun failed err:%v", err)
		return
	}
	logs.Info("program exited")
	fmt.Println("============")
}
